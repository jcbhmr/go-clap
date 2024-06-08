package styling

import "github.com/jcbhmr/go-clap/anstyle"

type Styles struct {
	header anstyle.Style
	error anstyle.Style
	usage anstyle.Style
	literal anstyle.Style
	placeholder anstyle.Style
	valid anstyle.Style
	invalid anstyle.Style
}

func (s Styles) Plain() Styles {
	return Styles{}
}
