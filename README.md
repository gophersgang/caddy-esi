# ESI Tags for Caddy Server (Beta)

This plugin implements partial ESI [Edge Side Includes](https://en.wikipedia.org/wiki/Edge_Side_Includes) 
support for the [Caddy Webserver](https://caddyserver.com) to allow data
retrieving from heterogeneous backends aka. micro services.

[![Build Status](https://travis-ci.org/SchumacherFM/caddyesi.svg?branch=master)](https://travis-ci.org/SchumacherFM/caddyesi)

#### Some features:

- Query data from HTTP, HTTPS, gRPC, Redis, Memcache and soon SQL.
- Multiple incoming requests trigger only one single parsing of the ESI tags per
page
- Querying multiple backend server parallel and concurrent.
- Coalesce multiple incoming requests into one single request to a backend
server
- Forwarding and returning (todo) of HTTP headers from backend servers
- Query multiple backend servers sequentially as a fall back mechanism
- Query multiple backend servers parallel and use the first returned result and
discard other responses, an obvious race condition. (todo)
- Support for NoSQL Server to query a key and simply display its value
- Variables support based on Server, Cookie, Header or GET/POST form parameters
- Error handling and fail over. Either display a text from a string or a static
file content when a backend server is unavailable.
- If an HTTP/S backend request takes longer than the specified timeout, flip the
source URL into an XHR request or an HTTP2 push. (todo)
- No full ESI support, if desired use a scripting language ;-)
- Supports Go >= 1.8

## High level overview

![Architectural overview](caddy-esi-archi.png)

## Plugin configuration (optional)

```
https://cyrillschumacher.local:2718 {
    ...
    other caddy directives
    ...
    esi [/path_optional] {
        [timeout 5ms|100us|1m|...]
        [ttl 5ms|100us|1m|...]
        [max_body_size 500kib|5MB|10GB|2EB|etc]
        [page_id_source [host,path,ip, etc]]
        [allowed_methods [GET,POST,etc]]
        [cmd_header_name [X-What-Ever]]
        [cache redis://localhost:6379/0]
        [cache redis://localhost:6380/0]
        [cache memcache://localhost:11211/2]
        [cache inmemory]
        [on_error (filename|"any text")]
        [log_file (filename|stdout|stderr)]
        [log_level (fatal|info|debug)]
        [resources path/to/url_configrations.(xml|json)]
    }
}
```

Most of the **global optional configuration values** can be overwritten with a
different value in a single ESI tag.

`esi` defines the namespace in the Caddy configuration file. Reserved keys are:

| Config Name |  Default | Support in ESI tag | Description |
| ----------- |  ------- | ----------- |  ----------- |
| `[path]`   | `/`    | n/a | Under this path all pages gets parsed for ESI tags. Default path sets to slash. |
| `timeout`   | 20s    | Yes | Time when a request to a resource should be canceled. [time.Duration](https://golang.org/pkg/time/#Duration) |
| `ttl`      | disabled  | Yes | Time-to-live value in the NoSQL cache for data returned from the backend resources. |
| `max_body_size` | 5MB       | Yes |  Limits the size of the returned body from a backend resource. |
| `cache` | disabled | No | Defines a cache service which stores the retrieved data from a backend resource but only when the ttl (within an ESI tag) has been set. Can only occur multiple times! |
| `page_id_source` | `host`, `path` | No | Special directive on how to identify a request to the same page. The following settings can be used to calculate the hash value. Available settings: `remoteaddr`, `realip`, `scheme`, `host`, `path`, `rawpath`, `rawquery` and `url`. Special feature to access cookies and headers: Prefix with `cookie-` or `header-` to access the appropriate value. Attention: The more granular you define the higher possibility occurs that your RAM will be filled up (will be fixed ...). |
| `allowed_methods` | `GET` | No | Any method listed here triggers the ESI middleware |
| `cmd_header_name` | Disabled | No | Specify here a header name to enable certain commands for e.g. purging the internal ESI tag cache, setting log-level |
| `log_file` | disabled | No | Put in here either a file name or the wordings `stderr` or `stdout` to write to those file descriptors. |
| `log_level` | disabled | No | Available key words `debug` the most verbose and `info`, less verbose. |
| `log_format` | n/a | No | Not yet supported. Ideas? |
| `resources` | n/a | No | Additional configuration file in XML or JSON to refer to more backend resource services. |

`cmd_header_name` current supported values are:

- `purge` use the value `purge` with your defined `cmd_header_name` to purge the
ESI tag cache. `X-Esi-Cmd: purge`
- `log-debug` enables debug logging. Costs heavily performance.
- `log-info` enables info logging. Logs some errors and other note worthy informations.
- `log-none` disables logging.

`resources` defines the path to a configuration file for more backend resource
services. An example on how the XML or JSON must be coded:

```xml
<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<items>
    <item>
        <alias>redis01</alias>
        <url><![CDATA[redis://127.0.0.1:6379/?db=0&max_active=10&max_idle=4]]></url>
        <!--<query>Unused and hence optional</query>-->
    </item>
    <item>
        <alias>grpc01</alias>
        <url><![CDATA[grpc://127.0.0.1:53044/?timeout=60s&tls=1&ca_file=../path/to/root.pem&server_host_override=my.domain.kom]]></url>
        <!--<query>Unused and hence optional</query>-->
    </item>
    <item>
        <alias>mysql01</alias>
        <url><![CDATA[mysql://user:password@tcp(localhost:5555)/dbname?charset=utf8mb4,utf8&tls=skip-verify]]></url>
        <query><![CDATA[SELECT `value` FROM tableX WHERE key='?']]></query> <!--Creates a long-lived prepared statement-->
    </item>
    <item>
        <alias>mysql02</alias>
        <url>mysql01</url><!--Uses the connection of mysql01-->
        <query><![CDATA[SELECT `value` FROM tableY WHERE another_key=?]]></query> <!--Creates a long-lived prepared statement-->
    </item>
</items>
```

```json
[
  {
    "alias": "redis01",
    "url": "redis://127.0.0.1:6379/?db=0&max_active=10&max_idle=4"
  },
  {
    "alias": "grpc01",
    "url": "grpc://127.0.0.1:53044/?timeout=60s&tls=1&ca_file=../path/to/root.pem&server_host_override=my.domain.kom"
  },
  {
    "alias": "mysql01",
    "url": "mysql://user:password@tcp(localhost:5555)/dbname?charset=utf8mb4,utf8&tls=skip-verify",
    "query": "SELECT `value` FROM tableX WHERE key='?'"
  },
  {
    "alias": "mysql02",
    "url": "mysql01",
    "query": "SELECT `value` FROM tableY WHERE another_key=?"
  }
]
```

Please be aware that this file must be stored securely on the hard drive and
that the owner of the Caddy process can read it. Other resource credential
loading options might be added in the future.

The further usage of configuration gets explained in the ESI tag documentation.

## Supported ESI Tags and their attributes

Implemented:

- [x] Caddy configuration parsing
- [x] ESI tag parsing
- [x] ESI tag replacement
- [x] Background Fetcher workers
- [x] Basic ESI Tag
- [x] With timeout
- [ ] With ttl
- [x] Load local file after timeout
- [ ] Flip src to AJAX call after timeout
- [x] Forward all headers
- [x] Forward some headers
- [ ] Forward POST,PATCH, PUT data
- [ ] Return all headers
- [ ] Return some headers
- [ ] Forward QUERY STRING and/or POST form data
- [x] Multiple sources
- [ ] Multiple sources with `race="true"`
- [x] Dynamic sources/keys (string replacement)
- [ ] Conditional tag loading
- [x] Redis access
- [x] Memcache access
- [ ] MySQL access
- [ ] PGSQL access
- [x] gRPC access
- [x] Shell scripts/programs access (stderr|out|in) communication
- [x] Handle compressed content from backends (Go http.Client)
- [x] Query HTTP/S backend servers
- [ ] Coalesce multiple requests into one backend request

Databases, NoSQL or gRPC access in package `backend` must be enabled via Go
build tags. Please see the source files for the correct build tag.

### TL;DR ESI tag options

Quick overview which options are available for two different kinds of ESI tags.

1. Querying an HTTP backend

```
<esi:include src="https://micro1.service/esi/foo" src="https://microN.service/esi/foo" 
    timeout="time.Duration" ttl="time.Duration" 
    onerror="text or path to file" maxbodysize="bytes"
    forwardheaders="all or specific comma separated list of header names"
    returnheaders="all or specific comma separated list of header names"
    coalesce="true|false" printdebug="true|false"
/>
```

2. Querying a NOSQL database or gRPC service

```
<esi:include src="alias_name1" src="alias_nameN" key="key name" 
    onerror="text or path to file" timeout="time.Duration" coalesce="true|false" printdebug="true|false"/>
```

`key` or `src` can contain variables which gets replaced to their corresponding
values. See section "Dynamic sources and keys".

### Basic tag

The basic tag waits on the `src` until `esi.timeout` cancels the loading. If src
contains an invalid URL or has not been specified at all, nothing happens.

An unreachable `src` gets marked as "failed" in the internal cache by a circuit
breaker with an exponential back-off strategy: 2ms, 4ms, 8ms, 16ms, 32ms ... or
as defined in the `esi.timeout` or tag based `timeout` attribute.

ESI tags are getting internally cached after they have been parsed. Changes to
ESI tags during development must trigger a cache clearance via restarting or
sending the correct `cmd_header_name` header key.

A page ID defines the internal cache key for a set of ESI tags listed in an HTML
page. For each incoming request the ESI processor knows beforehand which ESI
tags to load and to process instead of parsing the HTML page sequentially. All
ESI `src` will be loaded parallel before the page output starts. Once the
content has been received the Content-Length header gets calculated and the
output starts. 
TODO: Chunked transfer implementation with an immediately starting output.

```
<esi:include src="https://micro.service/esi/foo" />
```

### With timeout (optional)

The basic tag with the attribute `timeout` waits for the src until the timeout
occurs. After the timeout, the ESI tag gets not rendered and hence displays
nothing. The attribute `timeout` overwrites the default `esi.timeout`.

```
<esi:include src="https://micro.service/esi/foo" timeout="time.Duration" />
```

100% Support with http/s requests to backend services.

### With ttl (optional) (TODO)

The basic tag with the attribute `ttl` stores the returned data from the `src`
in the specified `esi.backend`. The attribute `ttl` overwrites the default
`esi.ttl`. If `esi.backend` has not been set or `ttl` set to empty, caching is
disabled.

`esi.backend` uses the `src` as a cache key.

```
<esi:include src="https://micro.service/esi/foo" ttl="time.Duration" />
```
 
### Load local file after timeout (optional)

The basic tag with the attribute `timeout` waits for the src until the timeout
occurs. After the timeout, the ESI processor loads the local file or text from
the attribute `onerror` and renders it into the page. If the file does not
exists an ugly error gets shown.

Supported file extensions: `"html", "htm", "xml", "txt", "json"`

```
<esi:include src="https://micro.service/esi/foo" timeout="time.Duration" onerror="path/to/mylocalFile.html"/>

<esi:include src="https://micro.service/esi/foo" timeout="time.Duration" onerror="Cannot load weather service"/>
```

### Max body size to limit the size of the body returned from a backend (optional)

The basic tag with the attribute `maxbodysize="size"` takes care that the
returned body size has been limited to the provided value. Default body size is
5 MB. Other attributes can be additionally defined.
Available size identifiers: [https://github.com/dustin/go-humanize/blob/master/bytes.go#L34](https://github.com/dustin/go-humanize/blob/master/bytes.go#L34)

```
<esi:include src="https://micro.service/esi/foo" maxbodysize="3MB"/>
```

### Flip src to AJAX call after timeout (optional) (TODO)

The basic tag with the attribute `timeout` waits for the src until the timeout
occurs. After the timeout, the ESI processor converts the URI in the `src`
attribute to an AJAX call in the frontend. The attribute `onerror="ajax"` must
exists or feature is disabled. TODO: JS code of the template ...

```
<esi:include src="https://micro.service/esi/foo" timeout="time.Duration" onerror="ajax"/>
```

### Forward POST, PATCH or PUT data (optional)

The basic tag with the attribute `forwardpostdata` forwards all incoming request
POST, PATCH, PUT data to the `src`. Other attributes can be additionally
defined. Default value `false` which does not forward anything. POST data gets
only forwarded to those backend resources which can handle the data.

```
<esi:include src="https://micro.service/esi/foo" forwardpostdata="true"/>
```

### Forward all headers (optional)

The basic tag with the attribute `forwardheaders` forwards all incoming request
headers to the `src`. Other attributes can be additionally defined. POST data
gets only forwarded to those backend resources which can handle the data.

```
<esi:include src="https://micro.service/esi/foo" forwardheaders="all"/>
```

### Forward some headers (optional)

The basic tag with the attribute `forwardheaders` forwards only the specified
headers of the incoming request to the `src`. Other attributes can be
additionally defined. POST data gets only forwarded to those backend resources
which can handle the data.

```
<esi:include src="https://micro.service/esi/foo" forwardheaders="Cookie,Accept-Language,Authorization"/>
```

### Return all headers (optional) (TODO)

The basic tag with the attribute `returnheaders` returns all `src` headers to
the final response. Other attributes can be additionally defined. If duplicate
headers from multiple sources occurs, they are getting appended to the response.

```
<esi:include src="https://micro.service/esi/foo" returnheaders="all"/>
```

### Return some headers (optional) (TODO)

The basic tag with the attribute `returnheaders` returns the listed headers of
the `src` to the final response. Other attributes can be additionally defined. If
duplicate headers from multiple sources occurs, they are getting appended to the
response.

```
<esi:include src="https://micro.service/esi/foo" returnheaders="Set-Cookie"/>
```

### Coalesce multiple requests into one backend request (optional)

The basic tag with the attribute `coalesce="boolean"` takes care that for
multiple incoming requests only one backend request gets fired. Other parallel
incoming requests waits until the backend returns from the initializing request.
Any returned data from a backend resource gets shared with an unknown number of
external requests. This feature limits the pressure onto a backend resource.
Other attributes can be additionally defined. Default value `false`, hence
disabled.
`coalesce="false|true|1|0"`

```
<esi:include src="https://micro.service/esi/foo" coalesce="true"/>
```

### Printdebug (optional)

The basic tag with the attribute `printdebug="boolean"` allows to print
debugging information as an HTML comment. Debugging contains the time taken to
complete the backend request, any possible error and the raw tag itself. Be
aware that if you store sensitive information in the raw tag they will leak
towards the evil internet. Other attributes can be additionally defined. Default
value `false`, hence disabled.
`printdebug="false|true|1|0"`

```
<esi:include src="https://micro.service/esi/foo" printdebug="true"/>
```

### Multiple sources

The basic ESI tag can contain multiple sources. The ESI processor tries to load
`src` attributes in its specified order. The next `src` gets called after the
`esi.timeout` or `timeout` occurs. Other attributes can be additionally defined.
Add the attribute `race="true"` (TODO) to fire all requests at once and the one which
is the fastest gets served and the others dropped.

```
<esi:include 
    src="https://micro1.service/esi/foo" 
    src="http://micro2.service/esi/foo"
    src="https://micro3.service/esi/foo" 
    timeout="time.Duration" 
    race="boolean" />
```

### Dynamic sources and keys (string replacement)

The basic ESI tag can extend all `src` URLs and `key` attributes with additional
parameters from the `http.Request` object. A replacement string must be one of
the listed below and must consists of the pattern: `{[^\}]+}`. Replacement
strings can occur anywhere and n-times in the `src` or `key` attribute.

```
<esi:include src="http://micro.service/search?query={Fsearch_term}"/>
<esi:include src="https://micro.service/{dir}/{file}/?id={HMy-Header-Key}"/>
```

The following string replacements are possible:

- {method}
- {scheme}
- {hostname}
- {host} (incl. port)
- {hostonly} (excl. port)
- {path}         (header Caddy-Rewrite-Original-URI supported)
- {path_escaped} (header Caddy-Rewrite-Original-URI supported)
- {rewrite_path}
- {rewrite_path_escaped}
- {query}
- {query_escaped}
- {fragment}
- {proto}
- {remote} IP address, might not be the real IP address
- {port}
- {uri} the RequestURI
- {uri_escaped} the RequestURI
- {when}
- {when_iso}
- {file}
- {dir}

To access any value in the header, cookie, GET form or POST form you can write a
special crafted replacement string.

- {HMy-Header-Key-Name} must start with `H` followed by the header key name
(case in-sensitive), in this case: `My-Header-Key-Name`.
- {CMy-Cookie-Name} must start with `C` followed by the cookie name (case
sensitive), in this case: `My-Cookie-Name`.
- {FMy-Input-Name} must start with `F` followed by the name of the field (case
sensitive), in this case: `My-Input-Name`.

### Conditional tag loading (TODO)

The basic ESI tag can contain the attribute `condition`. A `condition` must
evaluate an expression to `true` to trigger the loading of the `src`.

```
<esi:include src="http://micro.service/search?query={{ .Req.Form.Encode }}"
    condition="{hostname} eq 'customer.micro.service'"/>
```

### Access via src aliases

The ESI processor can access NoSQL, gRPC and SQL resources which are specified
in the Caddy resources configuration file. The `src` attribute must contain the
valid alias name as defined in resources file. The ESI tag must contain a `key`
attribute which accesses the value in the NoSQL server or forwards that value to
gRPC or uses it in the SELECT query for a database. The returned value will be
rendered unmodified into the HTML page output. You are responsible for secure
escaping of your data. If multiple `src` tags are defined the next `src` gets
used once the key cannot be found in the previous source or the `timeout`
applies. Return and forward headers might not be supported.

```
<esi:include src="redisAWS1" src="redisLocal1" key="my_redis_key_x" timeout="time.Duration" />
<esi:include src="redisAWS1" key="my_redis_key_x" onerror="myLocalFile.html" timeout="time.Duration" />
<esi:include src="redis1" key="prefix_{host}_{HX-Whatever}" timeout="time.Duration" />
```

### Registered schemes or aliases in the src attribute

The `src` attribute can contain or refer to the following aliases (defined in
the resource configuration file) to fetch content from a resource backend:

- http://
- https://
- sh://
- Name of the aliases as defined in the resources file

To access the next resources you must specify them as an alias in the resources
`xml|json` file. Because those connection lives the whole Caddy process and
might never get closed.

- redis://
- memcache://
- sql:// a SQL query (TODO)
- grpc:// a remote procedure call with gRPC

#### HTTP / HTTPS

`http|s` allows the usual HTTP requests to a resource.

Always enabled.

#### sh shell

`sh` starts a shell program or any other executable script. The incoming HTTP
request plus the current arguments gets transformed into a JSON string and will
be passed as first argument to the program via stdin. If the program writes to
stderr CaddyESI recognizes that and triggers it as an usual error. Too many
errors and the circuit breaker opens. To output towards the HTTP response gets written
to stdout. The src attribute must look like:

```html
<esi:include src="sh:///path/to/myGoBinary" />
<esi:include src="sh:///path/to/my/slow/php/script.php" />
<esi:include src="sh:///path/to/my/slow/php/script.php --arg1 --arg2=x --argN=y" />

<esi:include src="sh://php /path/to/my/slow/php/script.php" />
```

*ProTip:* Provide always the full path to the binary or script because the
lookup time in the operation systems environment PATH variable will take long.
In general shell execution is pretty slow no matter which kind of
program/script you call.

Disabled by default and must be enabled with a new compiled Caddy binary and the
build tag `esishell`.

#### SQL queries TODO

`sql` Uses a prepared statement once the ESI tag has been parsed.

```html
<esi:include src="alias_name" key="{FMyFormInputName}" />
```

`key` will be used as parameter in the WHERE query part.

Disabled by default and must be enabled with a new compiled Caddy binary and the
build tag `esisql`.

#### gRPC Remote Procedure Calls

`grpc` queries another gRPC endpoint. A struct of arguments gets encoded and
passed over the wire.

```html
<esi:include src="alias_name" key="cart1" />
```

You are responsible for building the gRPC server. The `.proto` file and the
corresponding Go file have been placed in package `backend/esigrpc`. Your gRPC
server can run in any programming language you like and is supported by gRPC.
[https://www.grpc.io](https://www.grpc.io)

Disabled by default and must be enabled with a new compiled Caddy binary and the
build tag `esigrpc`.

## Unsupported ESI Tags

All other tags, as defined in
[https://www.w3.org/TR/esi-lang](https://www.w3.org/TR/esi-lang), won't be
supported. You should switch to a server side scripting language ;-).

# Building

Use build tags to enable the different backend resource handlers in package
`esitag/backend`. Supported tags are:

- esiall: enables all handlers
- esigrpc: compiles gRPC handler only
- esiredis: compiles Redis handler only
- esimemcache: compiles Memcache handler only
- esishell: compiles shell handler only
- .... ?

See file `integration.sh` for examples.

# Contribute

Send me a pull request or open an issue if you encounter a bug or something can
be improved!

Multi-time pull request senders gets collaborator access.

# Attribution

### https://godoc.org/golang.org/x/sync/syncmap

In package `esitag` the type `EntitiesMap` and its files starting with
`syncmap_*.go`.

```text
Copyright 2016 The Go Authors. All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.

Copyright (c) 2009 The Go Authors. All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are
met:

   * Redistributions of source code must retain the above copyright
notice, this list of conditions and the following disclaimer.
   * Redistributions in binary form must reproduce the above
copyright notice, this list of conditions and the following disclaimer
in the documentation and/or other materials provided with the
distribution.
   * Neither the name of Google Inc. nor the names of its
contributors may be used to endorse or promote products derived from
this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
"AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
```

### github.com/dustin/go-humanize

```text
Copyright (c) 2005-2008  Dustin Sallings <dustin@spy.net>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

<http://www.opensource.org/licenses/mit-license.php>
```

### github.com/uber-go/zap

```text
Copyright (c) 2016 Uber Technologies, Inc.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
```

### github.com/bradfitz/gomemcache

```text
Copyright 2011 Google Inc.

Licensed under the Apache License, Version 2.0 (the "License"); you may not use
this file except in compliance with the License. You may obtain a copy of the
License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed
under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
CONDITIONS OF ANY KIND, either express or implied. See the License for the
specific language governing permissions and limitations under the License.
```

### github.com/garyburd/redigo

```text
Copyright 2012 Gary Burd

Licensed under the Apache License, Version 2.0 (the "License"): you may not use
this file except in compliance with the License. You may obtain a copy of the
License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed
under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
CONDITIONS OF ANY KIND, either express or implied. See the License for the
specific language governing permissions and limitations under the License.
```

### google.golang.org/grpc

```text
Copyright 2016, Google Inc.
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are
met:

    * Redistributions of source code must retain the above copyright
notice, this list of conditions and the following disclaimer.
    * Redistributions in binary form must reproduce the above
copyright notice, this list of conditions and the following disclaimer
in the documentation and/or other materials provided with the
distribution.
    * Neither the name of Google Inc. nor the names of its
contributors may be used to endorse or promote products derived from
this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
"AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
```

# License

[Cyrill Schumacher](https://github.com/SchumacherFM) - [My pgp public key](https://www.schumacher.fm/cyrill.asc)

Copyright 2016 Cyrill Schumacher All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License"); you may not use
this file except in compliance with the License. You may obtain a copy of the
License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed
under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
CONDITIONS OF ANY KIND, either express or implied. See the License for the
specific language governing permissions and limitations under the License.
