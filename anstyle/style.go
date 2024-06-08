package anstyle

import (
	"bytes"
	"fmt"
	"io"
)

type Style struct {
	fg        *Color
	bg        *Color
	underline *Color
	effects   Effects
}

func NewStyle() *Style {
	return &Style{
		fg:        nil,
		bg:        nil,
		underline: nil,
		effects:   Effects(0),
	}
}

func (s *Style) FgColor(color *Color) *Style {
	s.fg = color
	return s
}

func (s *Style) BgColor(color *Color) *Style {
	s.bg = color
	return s
}

func (s *Style) UnderlineColor(color *Color) *Style {
	s.underline = color
	return s
}

func (s *Style) Effects(effects Effects) *Style {
	s.effects = effects
	return s
}

func (s Style) Render() fmt.Stringer {
	return styleDisplay(s)
}

func (s Style) WriteTo(write io.Writer) (int64, error) {
	var i int64
	n, err := s.effects.writeTo(write)
	i += n
	if err != nil {
		return i, err
	}
	if s.fg != nil {
		n, err := s.fg.writeFgTo(write)
		i += n
		if err != nil {
			return i, err
		}
	}
	if s.bg != nil {
		n, err := s.bg.writeBgTo(write)
		i += n
		if err != nil {
			return i, err
		}
	}
	if s.underline != nil {
		n, err := s.underline.writeUnderlineTo(write)
		i += n
		if err != nil {
			return i, err
		}
	}
	return i, nil
}

type stringStringer string

func (s stringStringer) String() string {
	return string(s)
}

func (s Style) RenderReset() fmt.Stringer {
	if s != *NewStyle() {
		return stringStringer("\x1B[0m")
	} else {
		return stringStringer("")
	}
}

func (s Style) WriteResetTo(write io.Writer) (int64, error) {
	if s != *NewStyle() {
		n, err := write.Write([]byte("\x1B[0m"))
		return int64(n), err
	} else {
		return 0, nil
	}
}

func (s *Style) Bold() *Style {
	s.effects = *s.effects.Insert(EffectsBold)
	return s
}

func (s *Style) Dimmed() *Style {
	s.effects = *s.effects.Insert(EffectsDimmed)
	return s
}

func (s *Style) Italic() *Style {
	s.effects = *s.effects.Insert(EffectsItalic)
	return s
}

func (s *Style) Underline() *Style {
	s.effects = *s.effects.Insert(EffectsUnderline)
	return s
}

func (s *Style) Blink() *Style {
	s.effects = *s.effects.Insert(EffectsBlink)
	return s
}

func (s *Style) Invert() *Style {
	s.effects = *s.effects.Insert(EffectsInvert)
	return s
}

func (s *Style) Hidden() *Style {
	s.effects = *s.effects.Insert(EffectsHidden)
	return s
}

func (s *Style) Strikethrough() *Style {
	s.effects = *s.effects.Insert(EffectsStrikethrough)
	return s
}

func (s Style) GetFgColor() *Color {
	return s.fg
}

func (s Style) GetBgColor() *Color {
	return s.bg
}

func (s Style) GetUnderlineColor() *Color {
	return s.underline
}

func (s Style) GetEffects() Effects {
	return s.effects
}

func (s Style) IsPlain() bool {
	return s.fg == nil && s.bg == nil && s.underline == nil && s.effects.IsPlain()
}

func StyleFrom(effects Effects) *Style {
	return NewStyle().Effects(effects)
}

func (s Style) String() string {
	var buffer bytes.Buffer
	s.WriteTo(&buffer)
	return buffer.String()
}

type styleDisplay Style

func (s styleDisplay) String() string {
	var buffer bytes.Buffer
	Style(s).WriteTo(&buffer)
	return buffer.String()
}
