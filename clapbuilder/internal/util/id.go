package util

type Id string

const IdHelp Id = "help"
const IdVersion Id = "version"
const IdExternal Id = "external"

func (i Id) String() string {
	return string(i)
}
