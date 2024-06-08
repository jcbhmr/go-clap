package anstyle

import (
	"bytes"
	"fmt"
	"io"
)

//	enum Color {
//	  ANSI(ANSIColor),
//	  ANSI256(ANSI256Color),
//	  RGB(RGBColor),
//	}
type Color struct{ A any }

func (c Color) On(background Color) *Style {
	return NewStyle().
		FgColor(&c).
		BgColor(&background)
}

func (c Color) OnDefault() *Style {
	return NewStyle().FgColor(&c)
}

func (c Color) RenderFg() fmt.Stringer {
	switch v := c.A.(type) {
	case ANSIColor:
		return v.asFgBuffer()
	case ANSI256Color:
		return v.asFgBuffer()
	case RGBColor:
		return v.asFgBuffer()
	default:
		panic("unknown color type")
	}
}

func (c Color) writeFgTo(write io.Writer) (int64, error) {
	var buffer io.WriterTo
	switch v := c.A.(type) {
	case ANSIColor:
		buffer = v.asFgBuffer()
	case ANSI256Color:
		buffer = v.asFgBuffer()
	case RGBColor:
		buffer = v.asFgBuffer()
	default:
		panic("unknown color type")
	}
	return buffer.WriteTo(write)
}

func (c Color) RenderBg() fmt.Stringer {
	switch v := c.A.(type) {
	case ANSIColor:
		return v.asBgBuffer()
	case ANSI256Color:
		return v.asBgBuffer()
	case RGBColor:
		return v.asBgBuffer()
	default:
		panic("unknown color type")
	}
}

func (c Color) writeBgTo(write io.Writer) (int64, error) {
	var buffer io.WriterTo
	switch v := c.A.(type) {
	case ANSIColor:
		buffer = v.asBgBuffer()
	case ANSI256Color:
		buffer = v.asBgBuffer()
	case RGBColor:
		buffer = v.asBgBuffer()
	default:
		panic("unknown color type")
	}
	return buffer.WriteTo(write)
}

func (c Color) renderUnderline() fmt.Stringer {
	switch v := c.A.(type) {
	case ANSIColor:
		return v.asUnderlineBuffer()
	case ANSI256Color:
		return v.asUnderlineBuffer()
	case RGBColor:
		return v.asUnderlineBuffer()
	default:
		panic("unknown color type")
	}
}

func (c Color) writeUnderlineTo(write io.Writer) (int64, error) {
	var buffer io.WriterTo
	switch v := c.A.(type) {
	case ANSIColor:
		buffer = v.asUnderlineBuffer()
	case ANSI256Color:
		buffer = v.asUnderlineBuffer()
	case RGBColor:
		buffer = v.asUnderlineBuffer()
	default:
		panic("unknown color type")
	}
	return buffer.WriteTo(write)
}

func ColorFrom(v any) Color {
	switch v := v.(type) {
	case ANSIColor:
		return Color{v}
	case ANSI256Color:
		return Color{v}
	case RGBColor:
		return Color{v}
	case uint8:
		return Color{ANSI256Color(v)}
	case struct {
		A uint8
		B uint8
		C uint8
	}:
		return Color{RGBColor(v)}
	default:
		panic("unknown color type")
	}
}

func (a ANSIColor) IntoColor() Color {
	return Color{a}
}
func (a ANSI256Color) IntoColor() Color {
	return Color{a}
}
func (r RGBColor) IntoColor() Color {
	return Color{r}
}

type ANSIColor uint8

const (
	ANSIColorBlack ANSIColor = iota
	ANSIColorRed
	ANSIColorGreen
	ANSIColorYellow
	ANSIColorBlue
	ANSIColorMagenta
	ANSIColorCyan
	ANSIColorWhite
	ANSIColorBrightBlack
	ANSIColorBrightRed
	ANSIColorBrightGreen
	ANSIColorBrightYellow
	ANSIColorBrightBlue
	ANSIColorBrightMagenta
	ANSIColorBrightCyan
	ANSIColorBrightWhite
)

func (a ANSIColor) On(background Color) *Style {
	return NewStyle().
		FgColor(&Color{a}).
		BgColor(&background)
}

func (a ANSIColor) OnDefault() *Style {
	return NewStyle().FgColor(&Color{a})
}

func (a ANSIColor) RenderFg() fmt.Stringer {
	return a.asFgBuffer()
}

func (a ANSIColor) asFgBuffer() *bytes.Buffer {
	switch a {
	case ANSIColorBlack:
		return bytes.NewBufferString("\x1B[30m")
	case ANSIColorRed:
		return bytes.NewBufferString("\x1B[31m")
	case ANSIColorGreen:
		return bytes.NewBufferString("\x1B[32m")
	case ANSIColorYellow:
		return bytes.NewBufferString("\x1B[33m")
	case ANSIColorBlue:
		return bytes.NewBufferString("\x1B[34m")
	case ANSIColorMagenta:
		return bytes.NewBufferString("\x1B[35m")
	case ANSIColorCyan:
		return bytes.NewBufferString("\x1B[36m")
	case ANSIColorWhite:
		return bytes.NewBufferString("\x1B[37m")
	case ANSIColorBrightBlack:
		return bytes.NewBufferString("\x1B[90m")
	case ANSIColorBrightRed:
		return bytes.NewBufferString("\x1B[91m")
	case ANSIColorBrightGreen:
		return bytes.NewBufferString("\x1B[92m")
	case ANSIColorBrightYellow:
		return bytes.NewBufferString("\x1B[93m")
	case ANSIColorBrightBlue:
		return bytes.NewBufferString("\x1B[94m")
	case ANSIColorBrightMagenta:
		return bytes.NewBufferString("\x1B[95m")
	case ANSIColorBrightCyan:
		return bytes.NewBufferString("\x1B[96m")
	case ANSIColorBrightWhite:
		return bytes.NewBufferString("\x1B[97m")
	default:
		panic("unknown color")
	}
}

func (a ANSIColor) RenderBg() fmt.Stringer {
	return a.asBgBuffer()
}

func (a ANSIColor) asBgBuffer() *bytes.Buffer {
	switch a {
	case ANSIColorBlack:
		return bytes.NewBufferString("\x1B[40m")
	case ANSIColorRed:
		return bytes.NewBufferString("\x1B[41m")
	case ANSIColorGreen:
		return bytes.NewBufferString("\x1B[42m")
	case ANSIColorYellow:
		return bytes.NewBufferString("\x1B[43m")
	case ANSIColorBlue:
		return bytes.NewBufferString("\x1B[44m")
	case ANSIColorMagenta:
		return bytes.NewBufferString("\x1B[45m")
	case ANSIColorCyan:
		return bytes.NewBufferString("\x1B[46m")
	case ANSIColorWhite:
		return bytes.NewBufferString("\x1B[47m")
	case ANSIColorBrightBlack:
		return bytes.NewBufferString("\x1B[100m")
	case ANSIColorBrightRed:
		return bytes.NewBufferString("\x1B[101m")
	case ANSIColorBrightGreen:
		return bytes.NewBufferString("\x1B[102m")
	case ANSIColorBrightYellow:
		return bytes.NewBufferString("\x1B[103m")
	case ANSIColorBrightBlue:
		return bytes.NewBufferString("\x1B[104m")
	case ANSIColorBrightMagenta:
		return bytes.NewBufferString("\x1B[105m")
	case ANSIColorBrightCyan:
		return bytes.NewBufferString("\x1B[106m")
	case ANSIColorBrightWhite:
		return bytes.NewBufferString("\x1B[107m")
	default:
		panic("unknown color")
	}
}

func (a ANSIColor) asUnderlineBuffer() *displayBuffer {
	return ANSI256ColorFrom(a).asUnderlineBuffer()
}

type ANSI256Color uint8

func (a ANSI256Color) On(background Color) *Style {
	return NewStyle().
		FgColor(&Color{a}).
		BgColor(&background)
}

func (a ANSI256Color) OnDefault() *Style {
	return NewStyle().FgColor(&Color{a})
}

func (a ANSI256Color) Index() uint8 {
	return uint8(a)
}

func (a ANSI256Color) IntoANSI() *ANSIColor {
	var ansi ANSIColor
	switch a.Index() {
	case 0:
		ansi = ANSIColorBlack
	case 1:
		ansi = ANSIColorRed
	case 2:
		ansi = ANSIColorGreen
	case 3:
		ansi = ANSIColorYellow
	case 4:
		ansi = ANSIColorBlue
	case 5:
		ansi = ANSIColorMagenta
	case 6:
		ansi = ANSIColorCyan
	case 7:
		ansi = ANSIColorWhite
	case 8:
		ansi = ANSIColorBrightBlack
	case 9:
		ansi = ANSIColorBrightRed
	case 10:
		ansi = ANSIColorBrightGreen
	case 11:
		ansi = ANSIColorBrightYellow
	case 12:
		ansi = ANSIColorBrightBlue
	case 13:
		ansi = ANSIColorBrightMagenta
	case 14:
		ansi = ANSIColorBrightCyan
	case 15:
		ansi = ANSIColorBrightWhite
	default:
		return nil
	}
	return &ansi
}

func ANSI256ColorFromANSI(a ANSIColor) ANSI256Color {
	switch a {
	case ANSIColorBlack:
		return ANSI256Color(0)
	case ANSIColorRed:
		return ANSI256Color(1)
	case ANSIColorGreen:
		return ANSI256Color(2)
	case ANSIColorYellow:
		return ANSI256Color(3)
	case ANSIColorBlue:
		return ANSI256Color(4)
	case ANSIColorMagenta:
		return ANSI256Color(5)
	case ANSIColorCyan:
		return ANSI256Color(6)
	case ANSIColorWhite:
		return ANSI256Color(7)
	case ANSIColorBrightBlack:
		return ANSI256Color(8)
	case ANSIColorBrightRed:
		return ANSI256Color(9)
	case ANSIColorBrightGreen:
		return ANSI256Color(10)
	case ANSIColorBrightYellow:
		return ANSI256Color(11)
	case ANSIColorBrightBlue:
		return ANSI256Color(12)
	case ANSIColorBrightMagenta:
		return ANSI256Color(13)
	case ANSIColorBrightCyan:
		return ANSI256Color(14)
	case ANSIColorBrightWhite:
		return ANSI256Color(15)
	default:
		panic("unknown color")
	}
}

func (a ANSI256Color) RenderFg() fmt.Stringer {
	return a.asFgBuffer()
}

func (a ANSI256Color) asFgBuffer() *displayBuffer {
	var buffer displayBuffer
	buffer.writeString("\x1B[38;5;")
	buffer.writeCode(a.Index())
	buffer.writeString("m")
	return &buffer
}

func (a ANSI256Color) RenderBg() fmt.Stringer {
	return a.asBgBuffer()
}

func (a ANSI256Color) asBgBuffer() *displayBuffer {
	var buffer displayBuffer
	buffer.writeString("\x1B[48;5;")
	buffer.writeCode(a.Index())
	buffer.writeString("m")
	return &buffer
}

func (a ANSI256Color) asUnderlineBuffer() *displayBuffer {
	var buffer displayBuffer
	buffer.writeString("\x1B[58;5;")
	buffer.writeCode(a.Index())
	buffer.writeString("m")
	return &buffer
}

func ANSI256ColorFrom(v any) ANSI256Color {
	switch v := v.(type) {
	case uint8:
		return ANSI256Color(v)
	case ANSIColor:
		return ANSI256ColorFromANSI(v)
	default:
		panic("unknown color type")
	}
}

func (a ANSIColor) IntoANSI256Color() ANSI256Color {
	return ANSI256ColorFromANSI(a)
}

type RGBColor struct {
	A uint8
	B uint8
	C uint8
}

func (r RGBColor) On(background Color) *Style {
	return NewStyle().
		FgColor(&Color{r}).
		BgColor(&background)
}

func (r RGBColor) OnDefault() *Style {
	return NewStyle().FgColor(&Color{r})
}

func (r RGBColor) R() uint8 {
	return r.A
}

func (r RGBColor) G() uint8 {
	return r.B
}

func (r RGBColor) B2() uint8 {
	return r.C
}

func (r RGBColor) RenderFg() fmt.Stringer {
	return r.asFgBuffer()
}

func (r RGBColor) asFgBuffer() *displayBuffer {
	var buffer displayBuffer
	buffer.writeString("\x1B[38;2;")
	buffer.writeCode(r.A)
	buffer.writeString(";")
	buffer.writeCode(r.B)
	buffer.writeString(";")
	buffer.writeCode(r.C)
	buffer.writeString("m")
	return &buffer
}

func (r RGBColor) RenderBg() fmt.Stringer {
	return r.asBgBuffer()
}

func (r RGBColor) asBgBuffer() *displayBuffer {
	var buffer displayBuffer
	buffer.writeString("\x1B[48;2;")
	buffer.writeCode(r.A)
	buffer.writeString(";")
	buffer.writeCode(r.B)
	buffer.writeString(";")
	buffer.writeCode(r.C)
	buffer.writeString("m")
	return &buffer
}

func (r RGBColor) asUnderlineBuffer() *displayBuffer {
	var buffer displayBuffer
	buffer.writeString("\x1B[58;2;")
	buffer.writeCode(r.A)
	buffer.writeString(";")
	buffer.writeCode(r.B)
	buffer.writeString(";")
	buffer.writeCode(r.C)
	buffer.writeString("m")
	return &buffer
}

func RGBColorFrom(inner struct {
	A uint8
	B uint8
	C uint8
}) RGBColor {
	return RGBColor(inner)
}

type displayBuffer struct {
	buffer [19]byte
	len    uint
}

func (d *displayBuffer) writeString(s string) *displayBuffer {
	for i, b := range []byte(s) {
		d.buffer[d.len+uint(i)] = b
	}
	d.len += uint(len(s))
	return d
}

func (d *displayBuffer) writeCode(code uint8) *displayBuffer {
	var c1 uint8 = (code / 100) % 10
	var c2 uint8 = (code / 10) % 10
	var c3 uint8 = code % 10

	printed := true
	if c1 != 0 {
		printed = true
		d.buffer[d.len] = '0' + c1
		d.len++
	}
	if c2 != 0 || printed {
		printed = true
		d.buffer[d.len] = '0' + c2
		d.len++
	}
	d.buffer[d.len] = '0' + c3
	d.len++
	return d
}

func (d *displayBuffer) String() string {
	return string(d.buffer[:d.len])
}

func (d *displayBuffer) WriteTo(write io.Writer) (int64, error) {
	n, err := write.Write(d.buffer[:d.len])
	return int64(n), err
}
