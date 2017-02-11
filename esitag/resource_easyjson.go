// AUTOGENERATED FILE: easyjson marshaller/unmarshallers.

package esitag

import (
	http "net/http"
	url "net/url"

	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *jlexer.Lexer
	_ *jwriter.Writer
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
		case "max_body_size":
			out.Tag.MaxBodySize = uint64(in.Uint64())
		case "key":
			out.Tag.Key = string(in.String())
		case "return_headers":
			if in.IsNull() {
				in.Skip()
				out.Tag.ReturnHeaders = nil
			} else {
				in.Delim('[')
				if !in.IsDelim(']') {
					out.Tag.ReturnHeaders = make([]string, 0, 4)
				} else {
					out.Tag.ReturnHeaders = []string{}
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Tag.ReturnHeaders = append(out.Tag.ReturnHeaders, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "return_headers_all":
			out.Tag.ReturnHeadersAll = bool(in.Bool())
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
	if in.Tag.MaxBodySize != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"max_body_size\":")
		out.Uint64(uint64(in.Tag.MaxBodySize))
	}
	if in.Tag.Key != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"key\":")
		out.String(string(in.Tag.Key))
	}
	if len(in.Tag.ReturnHeaders) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"return_headers\":")
		if in.Tag.ReturnHeaders == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Tag.ReturnHeaders {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	if in.Tag.ReturnHeadersAll {
		if !first {
			out.RawByte(',')
		}
		//first = false
		out.RawString("\"return_headers_all\":")
		out.Bool(in.Tag.ReturnHeadersAll)
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
				easyjson1688e6a4DecodeNetUrl(in, out.URL)
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
					var v4 []string
					if in.IsNull() {
						in.Skip()
						v4 = nil
					} else {
						in.Delim('[')
						if !in.IsDelim(']') {
							v4 = make([]string, 0, 4)
						} else {
							v4 = []string{}
						}
						for !in.IsDelim(']') {
							var v5 string
							v5 = string(in.String())
							v4 = append(v4, v5)
							in.WantComma()
						}
						in.Delim(']')
					}
					(out.Header)[key] = v4
					in.WantComma()
				}
				in.Delim('}')
			}
		case "content_length":
			out.ContentLength = in.Int64()
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
					var v6 string
					v6 = string(in.String())
					out.TransferEncoding = append(out.TransferEncoding, v6)
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
					key := in.String()
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
					(out.Form)[key] = v7
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
					var v9 []string
					if in.IsNull() {
						in.Skip()
						v9 = nil
					} else {
						in.Delim('[')
						if !in.IsDelim(']') {
							v9 = make([]string, 0, 4)
						} else {
							v9 = []string{}
						}
						for !in.IsDelim(']') {
							var v10 string
							v10 = string(in.String())
							v9 = append(v9, v10)
							in.WantComma()
						}
						in.Delim(']')
					}
					(out.PostForm)[key] = v9
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
					var v11 []string
					if in.IsNull() {
						in.Skip()
						v11 = nil
					} else {
						in.Delim('[')
						if !in.IsDelim(']') {
							v11 = make([]string, 0, 4)
						} else {
							v11 = []string{}
						}
						for !in.IsDelim(']') {
							var v12 string
							v12 = string(in.String())
							v11 = append(v11, v12)
							in.WantComma()
						}
						in.Delim(']')
					}
					(out.Trailer)[key] = v11
					in.WantComma()
				}
				in.Delim('}')
			}
		case "remote_addr":
			out.RemoteAddr = string(in.String())
		case "request_uri":
			out.RequestURI = string(in.String())
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
		out.String(in.Method)
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
		out.String(in.Proto)
	}
	if in.ProtoMajor != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"proto_major\":")
		out.Int(in.ProtoMajor)
	}
	if in.ProtoMinor != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"proto_minor\":")
		out.Int(in.ProtoMinor)
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
			v13First := true
			for v13Name, v13Value := range in.Header {
				if !v13First {
					out.RawByte(',')
				}
				v13First = false
				out.String(string(v13Name))
				out.RawByte(':')
				if v13Value == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
					out.RawString("null")
				} else {
					out.RawByte('[')
					for v14, v15 := range v13Value {
						if v14 > 0 {
							out.RawByte(',')
						}
						out.String(string(v15))
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
		out.Int64(in.ContentLength)
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
			for v16, v17 := range in.TransferEncoding {
				if v16 > 0 {
					out.RawByte(',')
				}
				out.String(string(v17))
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
			v18First := true
			for v18Name, v18Value := range in.Form {
				if !v18First {
					out.RawByte(',')
				}
				v18First = false
				out.String(string(v18Name))
				out.RawByte(':')
				if v18Value == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
					out.RawString("null")
				} else {
					out.RawByte('[')
					for v19, v20 := range v18Value {
						if v19 > 0 {
							out.RawByte(',')
						}
						out.String(string(v20))
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
			v21First := true
			for v21Name, v21Value := range in.PostForm {
				if !v21First {
					out.RawByte(',')
				}
				v21First = false
				out.String(string(v21Name))
				out.RawByte(':')
				if v21Value == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
					out.RawString("null")
				} else {
					out.RawByte('[')
					for v22, v23 := range v21Value {
						if v22 > 0 {
							out.RawByte(',')
						}
						out.String(string(v23))
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
			v24First := true
			for v24Name, v24Value := range in.Trailer {
				if !v24First {
					out.RawByte(',')
				}
				v24First = false
				out.String(string(v24Name))
				out.RawByte(':')
				if v24Value == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
					out.RawString("null")
				} else {
					out.RawByte('[')
					for v25, v26 := range v24Value {
						if v25 > 0 {
							out.RawByte(',')
						}
						out.String(string(v26))
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
		out.String(in.RemoteAddr)
	}
	if in.RequestURI != "" {
		if !first {
			out.RawByte(',')
		}
		//first = false
		out.RawString("\"request_uri\":")
		out.String(string(in.RequestURI))
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
			out.Scheme = in.String()
		case "opaque":
			out.Opaque = in.String()
		case "user":
			if in.IsNull() {
				in.Skip()
				out.User = nil
			} else {
				if out.User == nil {
					out.User = new(url.Userinfo)
				}
				easyjson1688e6a4DecodeNetUrl1(in, out.User)
			}
		case "host":
			out.Host = in.String()
		case "path":
			out.Path = in.String()
		case "raw_path":
			out.RawPath = in.String()
		case "force_query":
			out.ForceQuery = in.Bool()
		case "raw_query":
			out.RawQuery = in.String()
		case "fragment":
			out.Fragment = in.String()
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
		out.String(in.Scheme)
	}
	if in.Opaque != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"opaque\":")
		out.String(in.Opaque)
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
		out.String(in.Host)
	}
	if in.Path != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"path\":")
		out.String(in.Path)
	}
	if in.RawPath != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"raw_path\":")
		out.String(in.RawPath)
	}
	if in.ForceQuery {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"force_query\":")
		out.Bool(in.ForceQuery)
	}
	if in.RawQuery != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"raw_query\":")
		out.String(in.RawQuery)
	}
	if in.Fragment != "" {
		if !first {
			out.RawByte(',')
		}
		//first = false
		out.RawString("\"fragment\":")
		out.String(in.Fragment)
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

		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}

		in.SkipRecursive()

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
