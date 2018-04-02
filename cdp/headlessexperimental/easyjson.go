// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package headlessexperimental

import (
	json "encoding/json"
	runtime "github.com/rjeczalik/chromedp/cdp/runtime"
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

func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpHeadlessexperimental(in *jlexer.Lexer, out *ScreenshotParams) {
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
		case "format":
			(out.Format).UnmarshalEasyJSON(in)
		case "quality":
			out.Quality = int64(in.Int64())
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
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpHeadlessexperimental(out *jwriter.Writer, in ScreenshotParams) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Format != "" {
		const prefix string = ",\"format\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Format).MarshalEasyJSON(out)
	}
	if in.Quality != 0 {
		const prefix string = ",\"quality\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.Quality))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ScreenshotParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpHeadlessexperimental(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ScreenshotParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpHeadlessexperimental(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ScreenshotParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpHeadlessexperimental(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ScreenshotParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpHeadlessexperimental(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpHeadlessexperimental1(in *jlexer.Lexer, out *EventNeedsBeginFramesChanged) {
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
		case "needsBeginFrames":
			out.NeedsBeginFrames = bool(in.Bool())
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
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpHeadlessexperimental1(out *jwriter.Writer, in EventNeedsBeginFramesChanged) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"needsBeginFrames\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.NeedsBeginFrames))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EventNeedsBeginFramesChanged) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpHeadlessexperimental1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EventNeedsBeginFramesChanged) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpHeadlessexperimental1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EventNeedsBeginFramesChanged) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpHeadlessexperimental1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EventNeedsBeginFramesChanged) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpHeadlessexperimental1(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpHeadlessexperimental2(in *jlexer.Lexer, out *EventMainFrameReadyForScreenshots) {
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
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpHeadlessexperimental2(out *jwriter.Writer, in EventMainFrameReadyForScreenshots) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EventMainFrameReadyForScreenshots) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpHeadlessexperimental2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EventMainFrameReadyForScreenshots) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpHeadlessexperimental2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EventMainFrameReadyForScreenshots) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpHeadlessexperimental2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EventMainFrameReadyForScreenshots) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpHeadlessexperimental2(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpHeadlessexperimental3(in *jlexer.Lexer, out *EnableParams) {
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
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpHeadlessexperimental3(out *jwriter.Writer, in EnableParams) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EnableParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpHeadlessexperimental3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EnableParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpHeadlessexperimental3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EnableParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpHeadlessexperimental3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EnableParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpHeadlessexperimental3(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpHeadlessexperimental4(in *jlexer.Lexer, out *DisableParams) {
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
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpHeadlessexperimental4(out *jwriter.Writer, in DisableParams) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DisableParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpHeadlessexperimental4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DisableParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpHeadlessexperimental4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DisableParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpHeadlessexperimental4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DisableParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpHeadlessexperimental4(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpHeadlessexperimental5(in *jlexer.Lexer, out *BeginFrameReturns) {
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
		case "hasDamage":
			out.HasDamage = bool(in.Bool())
		case "mainFrameContentUpdated":
			out.MainFrameContentUpdated = bool(in.Bool())
		case "screenshotData":
			out.ScreenshotData = string(in.String())
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
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpHeadlessexperimental5(out *jwriter.Writer, in BeginFrameReturns) {
	out.RawByte('{')
	first := true
	_ = first
	if in.HasDamage {
		const prefix string = ",\"hasDamage\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.HasDamage))
	}
	if in.MainFrameContentUpdated {
		const prefix string = ",\"mainFrameContentUpdated\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.MainFrameContentUpdated))
	}
	if in.ScreenshotData != "" {
		const prefix string = ",\"screenshotData\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ScreenshotData))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v BeginFrameReturns) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpHeadlessexperimental5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v BeginFrameReturns) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpHeadlessexperimental5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *BeginFrameReturns) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpHeadlessexperimental5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *BeginFrameReturns) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpHeadlessexperimental5(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpHeadlessexperimental6(in *jlexer.Lexer, out *BeginFrameParams) {
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
		case "frameTime":
			if in.IsNull() {
				in.Skip()
				out.FrameTime = nil
			} else {
				if out.FrameTime == nil {
					out.FrameTime = new(runtime.Timestamp)
				}
				(*out.FrameTime).UnmarshalEasyJSON(in)
			}
		case "deadline":
			if in.IsNull() {
				in.Skip()
				out.Deadline = nil
			} else {
				if out.Deadline == nil {
					out.Deadline = new(runtime.Timestamp)
				}
				(*out.Deadline).UnmarshalEasyJSON(in)
			}
		case "interval":
			out.Interval = float64(in.Float64())
		case "screenshot":
			if in.IsNull() {
				in.Skip()
				out.Screenshot = nil
			} else {
				if out.Screenshot == nil {
					out.Screenshot = new(ScreenshotParams)
				}
				(*out.Screenshot).UnmarshalEasyJSON(in)
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
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpHeadlessexperimental6(out *jwriter.Writer, in BeginFrameParams) {
	out.RawByte('{')
	first := true
	_ = first
	if in.FrameTime != nil {
		const prefix string = ",\"frameTime\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.FrameTime).MarshalEasyJSON(out)
	}
	if in.Deadline != nil {
		const prefix string = ",\"deadline\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Deadline).MarshalEasyJSON(out)
	}
	if in.Interval != 0 {
		const prefix string = ",\"interval\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.Interval))
	}
	if in.Screenshot != nil {
		const prefix string = ",\"screenshot\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Screenshot).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v BeginFrameParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpHeadlessexperimental6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v BeginFrameParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpHeadlessexperimental6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *BeginFrameParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpHeadlessexperimental6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *BeginFrameParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpHeadlessexperimental6(l, v)
}
