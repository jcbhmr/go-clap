package anstyle_test

import (
	"fmt"

	"github.com/jcbhmr/go-clap/anstyle"
)

func Example() {
	fmt.Printf("%q\n", anstyle.NewStyle().Bold())
	fmt.Printf("%q\n", anstyle.NewStyle().FgColor(&anstyle.Color{anstyle.ANSIColorBlue}))
	fmt.Printf("%q\n", anstyle.NewStyle().BgColor(&anstyle.Color{anstyle.ANSIColorBlue}))
	// Output:
	// "\x1b[1m"
	// "\x1b[34m"
	// "\x1b[44m"
}
