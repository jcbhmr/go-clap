package anstyle

import "fmt"

type Reset struct{}

func (r Reset) Render() fmt.Stringer {
	return r
}

var _ fmt.Stringer = (*Reset)(nil)

func (r Reset) String() string {
	return "\x1B[0m"
}
