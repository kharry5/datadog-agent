// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package monitor

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson6151911dDecodeGithubComDataDogDatadogAgentPkgSecurityRulesMonitor(in *jlexer.Lexer, out *RulesetLoadedEvent) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "policies":
			if in.IsNull() {
				in.Skip()
				out.Policies = nil
			} else {
				in.Delim('[')
				if out.Policies == nil {
					if !in.IsDelim(']') {
						out.Policies = make([]*PolicyState, 0, 8)
					} else {
						out.Policies = []*PolicyState{}
					}
				} else {
					out.Policies = (out.Policies)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *PolicyState
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(PolicyState)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.Policies = append(out.Policies, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "date":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Timestamp).UnmarshalJSON(data))
			}
		case "service":
			out.Service = string(in.String())
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
func easyjson6151911dEncodeGithubComDataDogDatadogAgentPkgSecurityRulesMonitor(out *jwriter.Writer, in RulesetLoadedEvent) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"policies\":"
		out.RawString(prefix[1:])
		if in.Policies == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Policies {
				if v2 > 0 {
					out.RawByte(',')
				}
				if v3 == nil {
					out.RawString("null")
				} else {
					(*v3).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"date\":"
		out.RawString(prefix)
		out.Raw((in.Timestamp).MarshalJSON())
	}
	{
		const prefix string = ",\"service\":"
		out.RawString(prefix)
		out.String(string(in.Service))
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RulesetLoadedEvent) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6151911dEncodeGithubComDataDogDatadogAgentPkgSecurityRulesMonitor(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RulesetLoadedEvent) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6151911dDecodeGithubComDataDogDatadogAgentPkgSecurityRulesMonitor(l, v)
}
func easyjson6151911dDecodeGithubComDataDogDatadogAgentPkgSecurityRulesMonitor1(in *jlexer.Lexer, out *RuleState) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = string(in.String())
		case "version":
			out.Version = string(in.String())
		case "expression":
			out.Expression = string(in.String())
		case "status":
			out.Status = string(in.String())
		case "message":
			out.Message = string(in.String())
		case "tags":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Tags = make(map[string]string)
				} else {
					out.Tags = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v4 string
					v4 = string(in.String())
					(out.Tags)[key] = v4
					in.WantComma()
				}
				in.Delim('}')
			}
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
func easyjson6151911dEncodeGithubComDataDogDatadogAgentPkgSecurityRulesMonitor1(out *jwriter.Writer, in RuleState) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	if in.Version != "" {
		const prefix string = ",\"version\":"
		out.RawString(prefix)
		out.String(string(in.Version))
	}
	{
		const prefix string = ",\"expression\":"
		out.RawString(prefix)
		out.String(string(in.Expression))
	}
	{
		const prefix string = ",\"status\":"
		out.RawString(prefix)
		out.String(string(in.Status))
	}
	if in.Message != "" {
		const prefix string = ",\"message\":"
		out.RawString(prefix)
		out.String(string(in.Message))
	}
	if len(in.Tags) != 0 {
		const prefix string = ",\"tags\":"
		out.RawString(prefix)
		{
			out.RawByte('{')
			v5First := true
			for v5Name, v5Value := range in.Tags {
				if v5First {
					v5First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v5Name))
				out.RawByte(':')
				out.String(string(v5Value))
			}
			out.RawByte('}')
		}
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RuleState) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6151911dEncodeGithubComDataDogDatadogAgentPkgSecurityRulesMonitor1(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RuleState) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6151911dDecodeGithubComDataDogDatadogAgentPkgSecurityRulesMonitor1(l, v)
}
func easyjson6151911dDecodeGithubComDataDogDatadogAgentPkgSecurityRulesMonitor2(in *jlexer.Lexer, out *PolicyState) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "name":
			out.Name = string(in.String())
		case "version":
			out.Version = string(in.String())
		case "source":
			out.Source = string(in.String())
		case "rules":
			if in.IsNull() {
				in.Skip()
				out.Rules = nil
			} else {
				in.Delim('[')
				if out.Rules == nil {
					if !in.IsDelim(']') {
						out.Rules = make([]*RuleState, 0, 8)
					} else {
						out.Rules = []*RuleState{}
					}
				} else {
					out.Rules = (out.Rules)[:0]
				}
				for !in.IsDelim(']') {
					var v6 *RuleState
					if in.IsNull() {
						in.Skip()
						v6 = nil
					} else {
						if v6 == nil {
							v6 = new(RuleState)
						}
						(*v6).UnmarshalEasyJSON(in)
					}
					out.Rules = append(out.Rules, v6)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjson6151911dEncodeGithubComDataDogDatadogAgentPkgSecurityRulesMonitor2(out *jwriter.Writer, in PolicyState) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"version\":"
		out.RawString(prefix)
		out.String(string(in.Version))
	}
	{
		const prefix string = ",\"source\":"
		out.RawString(prefix)
		out.String(string(in.Source))
	}
	{
		const prefix string = ",\"rules\":"
		out.RawString(prefix)
		if in.Rules == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v7, v8 := range in.Rules {
				if v7 > 0 {
					out.RawByte(',')
				}
				if v8 == nil {
					out.RawString("null")
				} else {
					(*v8).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PolicyState) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6151911dEncodeGithubComDataDogDatadogAgentPkgSecurityRulesMonitor2(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PolicyState) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6151911dDecodeGithubComDataDogDatadogAgentPkgSecurityRulesMonitor2(l, v)
}
func easyjson6151911dDecodeGithubComDataDogDatadogAgentPkgSecurityRulesMonitor3(in *jlexer.Lexer, out *HeartbeatEvent) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "policy":
			if in.IsNull() {
				in.Skip()
				out.Policy = nil
			} else {
				if out.Policy == nil {
					out.Policy = new(PolicyState)
				}
				(*out.Policy).UnmarshalEasyJSON(in)
			}
		case "date":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Timestamp).UnmarshalJSON(data))
			}
		case "service":
			out.Service = string(in.String())
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
func easyjson6151911dEncodeGithubComDataDogDatadogAgentPkgSecurityRulesMonitor3(out *jwriter.Writer, in HeartbeatEvent) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"policy\":"
		out.RawString(prefix[1:])
		if in.Policy == nil {
			out.RawString("null")
		} else {
			(*in.Policy).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"date\":"
		out.RawString(prefix)
		out.Raw((in.Timestamp).MarshalJSON())
	}
	{
		const prefix string = ",\"service\":"
		out.RawString(prefix)
		out.String(string(in.Service))
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v HeartbeatEvent) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6151911dEncodeGithubComDataDogDatadogAgentPkgSecurityRulesMonitor3(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *HeartbeatEvent) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6151911dDecodeGithubComDataDogDatadogAgentPkgSecurityRulesMonitor3(l, v)
}
