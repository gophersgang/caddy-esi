// partially AUTOGENERATED FILE: easyjson marshaller/unmarshallers.

package esitag

import (
	http "net/http"
	url "net/url"
	time "time"

	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

func easyjson1688e6a4DecodeGithubComSchumacherFMCaddyesiBackend(in *jlexer.Lexer, out *ResourceArgs) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "external_req":
			if in.IsNull() {
				in.Skip()
				out.ExternalReq = nil
			} else {
				if out.ExternalReq == nil {
					out.ExternalReq = new(http.Request)
				}
				easyjson1688e6a4DecodeGithubComSchumacherFMCaddyesiBackend1(in, out.ExternalReq)
			}
		case "url":
			out.URL = string(in.String())
		case "tag":
			easyjson1688e6a4DecodeGithubComSchumacherFMCaddyesiBackend2(in, &out.Tag)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson1688e6a4EncodeGithubComSchumacherFMCaddyesiBackend(out *jwriter.Writer, in ResourceArgs) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ExternalReq != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"external_req\":")
		if in.ExternalReq == nil {
			out.RawString("null")
		} else {
			easyjson1688e6a4EncodeGithubComSchumacherFMCaddyesiBackend1(out, in.ExternalReq)
		}
	}
	if in.URL != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"url\":")
		out.String(string(in.URL))
	}
	if true {
		if !first {
			out.RawByte(',')
		}
		out.RawString("\"tag\":")
		easyjson1688e6a4EncodeGithubComSchumacherFMCaddyesiEsitagTestdata2(out, in.Tag)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ResourceArgs) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1688e6a4EncodeGithubComSchumacherFMCaddyesiBackend(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ResourceArgs) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1688e6a4EncodeGithubComSchumacherFMCaddyesiBackend(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ResourceArgs) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1688e6a4DecodeGithubComSchumacherFMCaddyesiBackend(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ResourceArgs) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1688e6a4DecodeGithubComSchumacherFMCaddyesiBackend(l, v)
}
func easyjson1688e6a4DecodeGithubComSchumacherFMCaddyesiBackend2(in *jlexer.Lexer, out *Config) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "forward_headers":
			if in.IsNull() {
				in.Skip()
				out.ForwardHeaders = nil
			} else {
				in.Delim('[')
				if !in.IsDelim(']') {
					out.ForwardHeaders = make([]string, 0, 4)
				} else {
					out.ForwardHeaders = []string{}
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.ForwardHeaders = append(out.ForwardHeaders, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "return_headers":
			if in.IsNull() {
				in.Skip()
				out.ReturnHeaders = nil
			} else {
				in.Delim('[')
				if !in.IsDelim(']') {
					out.ReturnHeaders = make([]string, 0, 4)
				} else {
					out.ReturnHeaders = []string{}
				}
				for !in.IsDelim(']') {
					var v2 string
					v2 = string(in.String())
					out.ReturnHeaders = append(out.ReturnHeaders, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "forward_post_data":
			out.ForwardPostData = bool(in.Bool())
		case "forward_headers_all":
			out.ForwardHeadersAll = bool(in.Bool())
		case "return_headers_all":
			out.ReturnHeadersAll = bool(in.Bool())
		case "timeout":
			out.Timeout = time.Duration(in.Int64())
		case "ttl":
			out.TTL = time.Duration(in.Int64())
		case "max_body_size":
			out.MaxBodySize = uint64(in.Uint64())
		case "key":
			out.Key = string(in.String())
		case "coalesce":
			out.Coalesce = bool(in.Bool())
		case "print_debug":
			out.PrintDebug = bool(in.Bool())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson1688e6a4EncodeGithubComSchumacherFMCaddyesiEsitagTestdata2(out *jwriter.Writer, in Config) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.ForwardHeaders) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"forward_headers\":")
		if in.ForwardHeaders == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v3, v4 := range in.ForwardHeaders {
				if v3 > 0 {
					out.RawByte(',')
				}
				out.String(string(v4))
			}
			out.RawByte(']')
		}
	}
	if len(in.ReturnHeaders) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"return_headers\":")
		if in.ReturnHeaders == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.ReturnHeaders {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.String(string(v6))
			}
			out.RawByte(']')
		}
	}
	if in.ForwardPostData {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"forward_post_data\":")
		out.Bool(bool(in.ForwardPostData))
	}
	if in.ForwardHeadersAll {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"forward_headers_all\":")
		out.Bool(bool(in.ForwardHeadersAll))
	}
	if in.ReturnHeadersAll {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"return_headers_all\":")
		out.Bool(bool(in.ReturnHeadersAll))
	}
	if in.Timeout != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"timeout\":")
		out.Int64(int64(in.Timeout))
	}
	if in.TTL != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"ttl\":")
		out.Int64(int64(in.TTL))
	}
	if in.MaxBodySize != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"max_body_size\":")
		out.Uint64(uint64(in.MaxBodySize))
	}
	if in.Key != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"key\":")
		out.String(string(in.Key))
	}
	if in.Coalesce {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"coalesce\":")
		out.Bool(bool(in.Coalesce))
	}
	if in.PrintDebug {
		if !first {
			out.RawByte(',')
		}
		out.RawString("\"print_debug\":")
		out.Bool(bool(in.PrintDebug))
	}
	out.RawByte('}')
}
func easyjson1688e6a4DecodeGithubComSchumacherFMCaddyesiBackend1(in *jlexer.Lexer, out *http.Request) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "method":
			out.Method = in.String()
		case "url":
			if in.IsNull() {
				in.Skip()
				out.URL = nil
			} else {
				if out.URL == nil {
					out.URL = new(url.URL)
				}
				easyjson1688e6a4DecodeNetUrl(in, &*out.URL)
			}
		case "proto":
			out.Proto = in.String()
		case "proto_major":
			out.ProtoMajor = in.Int()
		case "proto_minor":
			out.ProtoMinor = in.Int()
		case "header":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Header = make(http.Header)
				} else {
					out.Header = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v7 []string
					if in.IsNull() {
						in.Skip()
						v7 = nil
					} else {
						in.Delim('[')
						if !in.IsDelim(']') {
							v7 = make([]string, 0, 4)
						} else {
							v7 = []string{}
						}
						for !in.IsDelim(']') {
							var v8 string
							v8 = string(in.String())
							v7 = append(v7, v8)
							in.WantComma()
						}
						in.Delim(']')
					}
					(out.Header)[key] = v7
					in.WantComma()
				}
				in.Delim('}')
			}
		case "content_length":
			out.ContentLength = int64(in.Int64())
		case "transfer_encoding":
			if in.IsNull() {
				in.Skip()
				out.TransferEncoding = nil
			} else {
				in.Delim('[')
				if !in.IsDelim(']') {
					out.TransferEncoding = make([]string, 0, 4)
				} else {
					out.TransferEncoding = []string{}
				}
				for !in.IsDelim(']') {
					var v9 string
					v9 = string(in.String())
					out.TransferEncoding = append(out.TransferEncoding, v9)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "close":
			out.Close = bool(in.Bool())
		case "host":
			out.Host = string(in.String())
		case "form":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Form = make(url.Values)
				} else {
					out.Form = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v10 []string
					if in.IsNull() {
						in.Skip()
						v10 = nil
					} else {
						in.Delim('[')
						if !in.IsDelim(']') {
							v10 = make([]string, 0, 4)
						} else {
							v10 = []string{}
						}
						for !in.IsDelim(']') {
							var v11 string
							v11 = string(in.String())
							v10 = append(v10, v11)
							in.WantComma()
						}
						in.Delim(']')
					}
					(out.Form)[key] = v10
					in.WantComma()
				}
				in.Delim('}')
			}
		case "post_form":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.PostForm = make(url.Values)
				} else {
					out.PostForm = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v12 []string
					if in.IsNull() {
						in.Skip()
						v12 = nil
					} else {
						in.Delim('[')
						if !in.IsDelim(']') {
							v12 = make([]string, 0, 4)
						} else {
							v12 = []string{}
						}
						for !in.IsDelim(']') {
							var v13 string
							v13 = string(in.String())
							v12 = append(v12, v13)
							in.WantComma()
						}
						in.Delim(']')
					}
					(out.PostForm)[key] = v12
					in.WantComma()
				}
				in.Delim('}')
			}
		case "trailer":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Trailer = make(http.Header)
				} else {
					out.Trailer = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v14 []string
					if in.IsNull() {
						in.Skip()
						v14 = nil
					} else {
						in.Delim('[')
						if !in.IsDelim(']') {
							v14 = make([]string, 0, 4)
						} else {
							v14 = []string{}
						}
						for !in.IsDelim(']') {
							var v15 string
							v15 = string(in.String())
							v14 = append(v14, v15)
							in.WantComma()
						}
						in.Delim(']')
					}
					(out.Trailer)[key] = v14
					in.WantComma()
				}
				in.Delim('}')
			}
		case "remote_addr":
			out.RemoteAddr = string(in.String())
		case "request_uri":
			out.RequestURI = string(in.String())
		case "body":
			//if in.IsNull() {
			in.Skip()
			out.Body = nil
			//} else {
			//	out.Body = ioutil.NopCloser(bytes.NewReader(in.Bytes()))
			//}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson1688e6a4EncodeGithubComSchumacherFMCaddyesiBackend1(out *jwriter.Writer, in *http.Request) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Method != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"method\":")
		out.String(string(in.Method))
	}
	if in.URL != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"url\":")
		if in.URL == nil {
			out.RawString("null")
		} else {
			easyjson1688e6a4EncodeNetUrl(out, *in.URL)
		}
	}
	if in.Proto != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"proto\":")
		out.String(string(in.Proto))
	}
	if in.ProtoMajor != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"proto_major\":")
		out.Int(int(in.ProtoMajor))
	}
	if in.ProtoMinor != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"proto_minor\":")
		out.Int(int(in.ProtoMinor))
	}
	if len(in.Header) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"header\":")
		if in.Header == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v17First := true
			for v17Name, v17Value := range in.Header {
				if !v17First {
					out.RawByte(',')
				}
				v17First = false
				out.String(string(v17Name))
				out.RawByte(':')
				if v17Value == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
					out.RawString("null")
				} else {
					out.RawByte('[')
					for v18, v19 := range v17Value {
						if v18 > 0 {
							out.RawByte(',')
						}
						out.String(string(v19))
					}
					out.RawByte(']')
				}
			}
			out.RawByte('}')
		}
	}
	if in.ContentLength != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"content_length\":")
		out.Int64(int64(in.ContentLength))
	}
	if len(in.TransferEncoding) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"transfer_encoding\":")
		if in.TransferEncoding == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v20, v21 := range in.TransferEncoding {
				if v20 > 0 {
					out.RawByte(',')
				}
				out.String(string(v21))
			}
			out.RawByte(']')
		}
	}
	if in.Close {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"close\":")
		out.Bool(bool(in.Close))
	}
	if in.Host != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"host\":")
		out.String(string(in.Host))
	}
	if len(in.Form) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"form\":")
		if in.Form == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v22First := true
			for v22Name, v22Value := range in.Form {
				if !v22First {
					out.RawByte(',')
				}
				v22First = false
				out.String(string(v22Name))
				out.RawByte(':')
				if v22Value == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
					out.RawString("null")
				} else {
					out.RawByte('[')
					for v23, v24 := range v22Value {
						if v23 > 0 {
							out.RawByte(',')
						}
						out.String(string(v24))
					}
					out.RawByte(']')
				}
			}
			out.RawByte('}')
		}
	}
	if len(in.PostForm) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"post_form\":")
		if in.PostForm == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v25First := true
			for v25Name, v25Value := range in.PostForm {
				if !v25First {
					out.RawByte(',')
				}
				v25First = false
				out.String(string(v25Name))
				out.RawByte(':')
				if v25Value == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
					out.RawString("null")
				} else {
					out.RawByte('[')
					for v26, v27 := range v25Value {
						if v26 > 0 {
							out.RawByte(',')
						}
						out.String(string(v27))
					}
					out.RawByte(']')
				}
			}
			out.RawByte('}')
		}
	}
	if len(in.Trailer) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"trailer\":")
		if in.Trailer == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v28First := true
			for v28Name, v28Value := range in.Trailer {
				if !v28First {
					out.RawByte(',')
				}
				v28First = false
				out.String(string(v28Name))
				out.RawByte(':')
				if v28Value == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
					out.RawString("null")
				} else {
					out.RawByte('[')
					for v29, v30 := range v28Value {
						if v29 > 0 {
							out.RawByte(',')
						}
						out.String(string(v30))
					}
					out.RawByte(']')
				}
			}
			out.RawByte('}')
		}
	}
	if in.RemoteAddr != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"remote_addr\":")
		out.String(string(in.RemoteAddr))
	}
	if in.RequestURI != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"request_uri\":")
		out.String(string(in.RequestURI))
	}
	if in.Body != nil && (in.Method == "POST" || in.Method == "PUT" || in.Method == "PATCH") { // todo check proper flag
		if !first {
			out.RawByte(',')
		}
		out.RawString("\"body\":")
		// out.Base64Bytes(in.Body)
		out.Base64Bytes([]byte(`TODO read body`))
	}
	out.RawByte('}')
}
func easyjson1688e6a4DecodeNetUrl(in *jlexer.Lexer, out *url.URL) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "scheme":
			out.Scheme = string(in.String())
		case "opaque":
			out.Opaque = string(in.String())
		case "user":
			if in.IsNull() {
				in.Skip()
				out.User = nil
			} else {
				if out.User == nil {
					out.User = new(url.Userinfo)
				}
				easyjson1688e6a4DecodeNetUrl1(in, &*out.User)
			}
		case "host":
			out.Host = string(in.String())
		case "path":
			out.Path = string(in.String())
		case "raw_path":
			out.RawPath = string(in.String())
		case "force_query":
			out.ForceQuery = bool(in.Bool())
		case "raw_query":
			out.RawQuery = string(in.String())
		case "fragment":
			out.Fragment = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson1688e6a4EncodeNetUrl(out *jwriter.Writer, in url.URL) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Scheme != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"scheme\":")
		out.String(string(in.Scheme))
	}
	if in.Opaque != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"opaque\":")
		out.String(string(in.Opaque))
	}
	if in.User != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"user\":")
		if in.User == nil {
			out.RawString("null")
		} else {
			easyjson1688e6a4EncodeNetUrl1(out, *in.User)
		}
	}
	if in.Host != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"host\":")
		out.String(string(in.Host))
	}
	if in.Path != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"path\":")
		out.String(string(in.Path))
	}
	if in.RawPath != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"raw_path\":")
		out.String(string(in.RawPath))
	}
	if in.ForceQuery {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"force_query\":")
		out.Bool(bool(in.ForceQuery))
	}
	if in.RawQuery != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"raw_query\":")
		out.String(string(in.RawQuery))
	}
	if in.Fragment != "" {
		if !first {
			out.RawByte(',')
		}
		out.RawString("\"fragment\":")
		out.String(string(in.Fragment))
	}
	out.RawByte('}')
}
func easyjson1688e6a4DecodeNetUrl1(in *jlexer.Lexer, out *url.Userinfo) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson1688e6a4EncodeNetUrl1(out *jwriter.Writer, in url.Userinfo) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}
