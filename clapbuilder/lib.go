package clapbuilder

import "github.com/jcbhmr/go-clap/clapbuilder/builder"
import "github.com/jcbhmr/go-clap/clapbuilder/parser"
import "github.com/jcbhmr/go-clap/clapbuilder/internal/util"
import "github.com/jcbhmr/go-clap/clapbuilder/internal/util/color"

type ArgAction = builder.ArgAction
type Command = builder.Command
type ValueHint = builder.ValueHint
type Arg = builder.Arg
type ArgGroup = builder.ArgGroup
type ArgMatches = parser.ArgMatches
type ColorChoice = color.ColorChoice
type Id = util.Id