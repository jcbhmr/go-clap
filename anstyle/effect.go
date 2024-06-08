package anstyle

import (
	"fmt"
	"io"
)

type seq[V any] func(yield func(V) bool)
type seq2[K, V any] func(yield func(K, V) bool)

type Effects uint16

const EffectsPlain Effects = Effects(0)
const EffectsBold Effects = Effects(1 << 0)
const EffectsDimmed Effects = Effects(1 << 1)
const EffectsItalic Effects = Effects(1 << 2)
const EffectsUnderline Effects = Effects(1 << 3)
const EffectsDoubleUnderline Effects = Effects(1 << 4)
const EffectsCurlyUnderline Effects = Effects(1 << 5)
const EffectsDottedUnderline Effects = Effects(1 << 6)
const EffectsDashedUnderline Effects = Effects(1 << 7)
const EffectsBlink Effects = Effects(1 << 8)
const EffectsInvert Effects = Effects(1 << 9)
const EffectsHidden Effects = Effects(1 << 10)
const EffectsStrikethrough Effects = Effects(1 << 11)

func NewEffects() *Effects {
	e := EffectsPlain
	return &e
}

func (e Effects) IsPlain() bool {
	return e == EffectsPlain
}

func (e Effects) Contains(other Effects) bool {
	return e&other == other
}

func (e *Effects) Insert(other Effects) *Effects {
	*e |= other
	return e
}

func (e *Effects) Remove(other Effects) *Effects {
	*e &= ^other
	return e
}

func (e *Effects) Clear() *Effects {
	*e = *NewEffects()
	return e
}

func (e *Effects) Set(other Effects, enable bool) *Effects {
	if enable {
		return e.Insert(other)
	} else {
		return e.Remove(other)
	}
}

func (e Effects) Render() fmt.Stringer {
	return effectsDisplay(e)
}

func (e Effects) Iter() seq[Effects] {
	var index int = 0
	return func(yield func(Effects) bool) {
		for index < len(metadata2) {
			i := index
			index += 1
			effect := Effects(1 << i)
			if e.Contains(effect) {
				if !yield(effect) {
					break
				}
			}
		}
	}

}

func (e Effects) indexIter() seq[int] {
	var index int = 0
	return func(yield func(int) bool) {
		for index < len(metadata2) {
			i := index
			index += 1
			effect := Effects(1 << i)
			if e.Contains(effect) {
				if !yield(i) {
					break
				}
			}
		}
	}
}

func (e Effects) writeTo(write io.Writer) (int64, error) {
	var n int64 = 0
	var err error
	e.indexIter()(func(index int) bool {
		n2, err2 := write.Write([]byte(metadata2[index].Escape))
		n += int64(n2)
		if err != nil {
			err = err2
			return false
		}
		return true
	})
	return n, err
}

func (e Effects) GoString() string {
	s := ""
	s += "Effects("
	var i int = 0
	Effects(e).indexIter()(func(index int) bool {
		if i != 0 {
			s += " | "
		}
		s += metadata2[index].Name
		i++
		return true
	})
	s += ")"
	return s
}

type metadata struct {
	Name   string
	Escape string
}

var metadata2 = [12]metadata{
	{Name: "BOLD", Escape: "\x1B[1m"},
	{Name: "DIMMED", Escape: "\x1B[2m"},
	{Name: "ITALIC", Escape: "\x1B[3m"},
	{Name: "UNDERLINE", Escape: "\x1B[4m"},
	{Name: "DOUBLE_UNDERLINE", Escape: "\x1B[21m"},
	{Name: "CURLY_UNDERLINE", Escape: "\x1B[4:3m"},
	{Name: "DOTTED_UNDERLINE", Escape: "\x1B[4:4m"},
	{Name: "DASHED_UNDERLINE", Escape: "\x1B[4:5m"},
	{Name: "BLINK", Escape: "\x1B[5m"},
	{Name: "INVERT", Escape: "\x1B[7m"},
	{Name: "HIDDEN", Escape: "\x1B[8m"},
	{Name: "STRIKETHROUGH", Escape: "\x1B[9m"},
}

type effectsDisplay Effects

func (e effectsDisplay) String() string {
	s := ""
	Effects(e).indexIter()(func(index int) bool {
		escape := metadata2[index].Escape
		s += escape
		return true
	})
	return s
}
