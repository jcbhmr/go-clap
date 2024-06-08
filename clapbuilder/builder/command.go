package clapbuilder

type Command struct{
	name string
	longFlag *string
	shortFlag *string
	displayName *string
	binName *string
	author *string
	version *string
	longVersion *string
	about *string
	longAbout *string
	beforeHelp *string
	beforeLongHelp *string
	afterHelp *string
	afterLongHelp *string
	aliases []struct{
		A string
		B bool
	}
	shortFlagAliases []struct{
		A byte
		B bool
	}
	longFlagAliases []struct{
		A string
		B bool
	}
	usageStr *string
	usageName *string
	helpStr *string
	dispOrd *uint
	// settings AppFlags
	// gSettings AppFlags
	// args MKeyMap
	subcommands []*Command
	// groups []ArgGroup
	currentHelpHeading *string
	currentDispOrd *uint
	subcommandValueName *string
	subcommandHeading *string
	// externalValueParser
	longHelpExists bool
	deferred func(*Command)*Command
	// appExt Extensions
}

func NewCommand(name string) *Command {
	return &Command{}
}

func (c *Command) Version(version string) *Command {
	return c
}

func (c *Command) Author(author string) *Command {
	return c
}

func (c *Command) About(about string) *Command {
	return c
}

func (c *Command) Arg(arg *Arg) *Command {
	if c.currentDispOrd != nil {
		if !arg.IsPositional() {
			current := *c.currentDispOrd
			if arg.dispOrd == nil {
				*arg.dispOrd = current
			}
			*c.currentDispOrd = current + 1
		}
	}
}

func (c *Command) Subcommand(subcommand *Command) *Command {
	return c
}

func (c *Command) Action(action ArgAction) *Command {
	return c
}

func (c *Command) GetMatches() *Matches {
	return &Matches{}
}