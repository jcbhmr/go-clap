package clapbuilder

import (
	"fmt"
	"runtime/debug"
	"strings"
)

var PseudoMacroCallerSkip = 0

var biCache *debug.BuildInfo

func bi() *debug.BuildInfo {
	if biCache == nil {
		var ok bool
		biCache, ok = debug.ReadBuildInfo()
		if !ok {
			panic(fmt.Errorf("debug.ReadBuildInfo() not ok"))
		}
	}
	return biCache
}

func GoModName1() string {
	parts := strings.Split(bi().Main.Path, "/")
	return parts[len(parts)-1]
}

func GoModVersion1() string {
	return bi().Main.Version
}

func GoModAuthors1() string {
	return ""
}

func GoModDescription1() string {
	return ""
}

func Command1(args ...any) *Command {
	if len(args) == 0 {
		return Command1(GoModName1())
	} else if len(args) == 1 {
		name := args[0].(string)
		cmd := NewCommand(name).Version(GoModVersion1())
		author := GoModAuthors1()
		if author != "" {
			cmd = cmd.Author(author)
		}
		about := GoModDescription1()
		if about != "" {
			cmd = cmd.About(about)
		}
		return cmd
	} else {
		panic(fmt.Errorf("len(args) expected [0,1] got %d", len(args)))
	}
}

func Arg1(args ...any) *Arg {
	return &Arg{}
}
