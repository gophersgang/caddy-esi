// Copyright 2016-2017, Cyrill @ Schumacher.fm and the CaddyESI Contributors
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not
// use this file except in compliance with the License. You may obtain a copy of
// the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations under
// the License.

// +build esiall esimemcache

package backend

import (
	"context"
	"net/http"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/corestoreio/errors"
)

func init() {
	RegisterResourceHandlerFactory("memcache", NewMemCache)
}

type esiMemCache struct {
	isCancellable bool
	url           string
	pool          *memcache.Client
}

// NewMemCache creates a new memcache resource handler supporting n-memcache server.
func NewMemCache(cfg *ConfigItem) (ResourceHandler, error) {
	addr, _, params, err := ParseNoSQLURL(cfg.URL)
	if err != nil {
		return nil, errors.NewNotValidf("[backend] NewMemCache error parsing URL %q => %s", cfg.URL, err)
	}

	idleTimeout, err := time.ParseDuration(params.Get("idle_timeout"))
	if err != nil {
		return nil, errors.NewNotValidf("[backend] NewMemCache.ParseNoSQLURL. Parameter idle_timeout not valid in  %q", cfg.URL)
	}

	servers := []string{addr}
	if len(params["server"]) > 0 {
		servers = append(servers, params["server"]...)
	}

	mc := &esiMemCache{
		isCancellable: params.Get("cancellable") == "1",
		url:           cfg.URL,
		pool:          memcache.New(servers...),
	}
	mc.pool.Timeout = idleTimeout

	if params.Get("lazy") == "1" {
		return mc, nil
	}

	// some pseudo ping
	if _, err := mc.pool.Get("caddyesi_key_not_found"); err != nil && err != memcache.ErrCacheMiss {
		return nil, errors.NewFatalf("[backend] MemCache ping failed: %s", err)
	}

	return mc, nil
}

// Closes closes the resource when Caddy restarts or reloads. If supported
// by the resource.
func (mc *esiMemCache) Close() error {
	return nil
}

// DoRequest returns a value from the field Key in the args argument. Header is
// not supported. Request cancellation through a timeout (when the client
// request gets cancelled) is supported.
func (mc *esiMemCache) DoRequest(args *ResourceArgs) (_ http.Header, _ []byte, err error) {
	if mc.isCancellable {
		// 50000	     28794 ns/op	    1026 B/op	      33 allocs/op
		return mc.doRequestCancel(args)
	}
	// 50000	     25071 ns/op	     529 B/op	      25 allocs/op
	return mc.doRequest(args)
}

func (mc *esiMemCache) validateArgs(args *ResourceArgs) (err error) {
	switch {
	case args.Key == "":
		err = errors.NewEmptyf("[esibackend] For ResourceArgs %#v the URL value is empty", args)
	case args.ExternalReq == nil:
		err = errors.NewEmptyf("[esibackend] For ResourceArgs %q => %q the ExternalReq value is nil", mc.url, args.Key)
	case args.Timeout < 1:
		err = errors.NewEmptyf("[esibackend] For ResourceArgs %q => %q the timeout value is empty", mc.url, args.Key)
	case args.MaxBodySize == 0:
		err = errors.NewEmptyf("[esibackend] For ResourceArgs %q => %q the maxBodySize value is empty", mc.url, args.Key)
	}
	return
}

func (mc *esiMemCache) doRequest(args *ResourceArgs) (_ http.Header, _ []byte, err error) {
	if err := mc.validateArgs(args); err != nil {
		return nil, nil, errors.Wrap(err, "[backend] doRequest.validateArgs")
	}

	key := args.Key

	nKey, err := args.TemplateToURL(args.KeyTemplate)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "[backend] MemCache.Get.TemplateToURL %q => %q", mc.url, args.Key)
	}
	if nKey != "" {
		key = nKey
	}

	itm, err := mc.pool.Get(key)
	if err == memcache.ErrCacheMiss {
		return nil, nil, errors.NewNotFoundf("[backend] URL %q: Key %q not found", mc.url, args.Key)
	}
	if err != nil {
		return nil, nil, errors.Wrapf(err, "[backend] Memcache.Get %q => %q", mc.url, args.Key)
	}

	if mbs := int(args.MaxBodySize); len(itm.Value) > mbs && mbs > 0 {
		itm.Value = itm.Value[:mbs]
	}
	// safe to return the slice header because each memcache query creates a new pointer
	return nil, itm.Value, err
}

// DoRequest returns a value from the field Key in the args argument. Header is
// not supported. Request cancellation through a timeout (when the client
// request gets cancelled) is supported.
func (mc *esiMemCache) doRequestCancel(args *ResourceArgs) (_ http.Header, _ []byte, err error) {
	if err := mc.validateArgs(args); err != nil {
		return nil, nil, errors.Wrap(err, "[backend] doRequestCancel.validateArgs")
	}

	key := args.Key

	// See git history for a version without context.WithTimeout. A bit faster and less allocs.
	ctx, cancel := context.WithTimeout(args.ExternalReq.Context(), args.Timeout)
	defer cancel()

	content := make(chan []byte)
	retErr := make(chan error)
	go func() {

		nKey, err := args.TemplateToURL(args.KeyTemplate)
		if err != nil {
			retErr <- errors.Wrapf(err, "[backend] MemCache.Get.TemplateToURL %q => %q", mc.url, args.Key)
			return
		}
		if nKey != "" {
			key = nKey
		}

		itm, err := mc.pool.Get(key)
		if err == memcache.ErrCacheMiss {
			retErr <- errors.NewNotFoundf("[backend] URL %q: Key %q not found", mc.url, args.Key)
			return
		}

		if err != nil {
			retErr <- errors.Wrapf(err, "[backend] MemCache.Get %q => %q", mc.url, args.Key)
			return
		}

		if mbs := int(args.MaxBodySize); len(itm.Value) > mbs && mbs > 0 {
			itm.Value = itm.Value[:mbs]
		}
		content <- itm.Value
	}()

	var value []byte
	select {
	case <-ctx.Done():
		err = errors.Wrapf(ctx.Err(), "[backend] MemCache Get Context cancelled. Previous possible error: %+v", retErr)
	case value = <-content:
	case err = <-retErr:
	}
	return nil, value, err
}
