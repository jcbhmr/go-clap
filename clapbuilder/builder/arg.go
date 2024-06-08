package builder

import (
	"errors"

	"github.com/jcbhmr/go-clap/clapbuilder/internal/util"
)

type Arg struct {
	id       util.Id
	help     *string
	longHelp *string
	action   *ArgAction
	// valueParser
	blacklist []util.Id
	// settings ArgFlags
	overrides []util.Id
	groups    []util.Id
	// requires []struct{
	// 	A ArgPredicate
	// 	B Id
	// }
	rIfs []struct {
		A util.Id
		B string
	}
	rIfsAll []struct {
		A []util.Id
		B string
	}
	rUnless    []util.Id
	rUnlessAll []util.Id
	short      *byte
	long       *string
	aliases    []struct {
		A string
		B bool
	}
	shortAliases []struct {
		A byte
		B bool
	}
	dispOrd  *uint
	valNames []string
	// numVals *ValueRange
	valDelim    *byte
	defaultVals []string
	// defaultValsIfs []struct{
	// 	A Id
	// 	B ArgPredicate
	// 	C *string
	// }
	defaultMissingVals []string
	terminator         *string
	index              *uint
	helpHeading        **string
	// valueHint *ValueHint
}

func NewArg(id util.Id) *Arg {
	return (&Arg{
		blacklist: []util.Id{},
		overrides: []util.Id{},
		groups:    []util.Id{},
		rIfs: []struct {
			A util.Id
			B string
		}{},
		rIfsAll: []struct {
			A []util.Id
			B string
		}{},
		rUnless:    []util.Id{},
		rUnlessAll: []util.Id{},
		aliases: []struct {
			A string
			B bool
		}{},
		shortAliases: []struct {
			A byte
			B bool
		}{},
		valNames:           []string{},
		defaultVals:        []string{},
		defaultMissingVals: []string{},
	}).Id(id)
}

func (a *Arg) Id(id util.Id) *Arg {
	a.id = id
	return a
}

func (a *Arg) Short(s *byte) *Arg {
	if s != nil {
		if *s == '-' {
			panic(errors.New("short option name cannot be `-`"))
		}
		a.short = s
	} else {
		a.short = nil
	}
	return a
}

func (a *Arg) Long(s *string) *Arg {
	a.long = s
	return a
}

func (a *Arg) Alias(name *string) *Arg {
	if name != nil {
		a.aliases = append(a.aliases, struct {
			A string
			B bool
		}{
			A: *name,
			B: false,
		})
	} else {
		a.aliases = []struct {
			A string
			B bool
		}{}
	}
	return a
}

func (a *Arg) ShortAlias(name *byte) *Arg {
	if name != nil {
		if *name == '-' {
			panic(errors.New("short alias name cannot be `-`"))
		}
		a.shortAliases = append(a.shortAliases, struct {
			A byte
			B bool
		}{
			A: *name,
			B: false,
		})
	} else {
		a.shortAliases = []struct {
			A byte
			B bool
		}{}
	}
	return a
}

func (a *Arg) Aliases(names []string) *Arg {
	for _, name := range names {
		a.aliases = append(a.aliases, struct {
			A string
			B bool
		}{
			A: name,
			B: false,
		})
	}
	return a
}

func (a *Arg) ShortAliases(names []byte) *Arg {
	for _, name := range names {
		if name == '-' {
			panic(errors.New("short alias name cannot be `-`"))
		}
		a.shortAliases = append(a.shortAliases, struct {
			A byte
			B bool
		}{
			A: name,
			B: false,
		})
	}
	return a
}

func (a *Arg) VisibleAlias(name *string) *Arg {
	if name != nil {
		a.aliases = append(a.aliases, struct {
			A string
			B bool
		}{
			A: *name,
			B: true,
		})
	} else {
		a.aliases = []struct {
			A string
			B bool
		}{}
	}
	return a
}

func (a *Arg) VisibleShortAlias(name *byte) *Arg {
	if name != nil {
		if *name == '-' {
			panic(errors.New("short alias name cannot be `-`"))
		}
		a.shortAliases = append(a.shortAliases, struct {
			A byte
			B bool
		}{
			A: *name,
			B: true,
		})
	} else {
		a.shortAliases = []struct {
			A byte
			B bool
		}{}
	}
	return a
}

func (a *Arg) VisibleAliases(names []string) *Arg {
	for _, name := range names {
		a.aliases = append(a.aliases, struct {
			A string
			B bool
		}{
			A: name,
			B: true,
		})
	}
	return a
}

func (a *Arg) VisibleShortAliases(names []byte) *Arg {
	for _, name := range names {
		if name == '-' {
			panic(errors.New("short alias name cannot be `-`"))
		}
		a.shortAliases = append(a.shortAliases, struct {
			A byte
			B bool
		}{
			A: name,
			B: true,
		})
	}
	return a
}

func (a *Arg) Index(idx *uint) *Arg {
	a.index = idx
	return a
}

func (a *Arg) Required() *Arg {
	return a
}

type ArgAction int

const (
	ArgActionSetTrue ArgAction = iota
)
