package clapbuilder

type Arg struct {
	id       Id
	help     *string
	longHelp *string
	action   *ArgAction
	// valueParser
	blacklist []Id
	// settings ArgFlags
	overrides []Id
	groups    []Id
	// requires []struct{
	// 	A ArgPredicate
	// 	B Id
	// }
	rIfs []struct {
		A Id
		B string
	}
	rIfsAll []struct {
		A []Id
		B string
	}
	rUnless    []Id
	rUnlessAll []Id
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

func NewArg(id Id) *Arg {
	return (&Arg{}).Id(id)
}

func (a *Arg) Id(id Id) *Arg {
	a.id = id
	return a
}

func (a *Arg) Short(s byte) *Arg {
	a.short = &s
	return a
}

func (a *Arg) Required() *Arg {
	return a
}

type ArgAction int

const (
	ArgActionSetTrue ArgAction = iota
)
