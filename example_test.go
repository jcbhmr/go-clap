package clap_test

import (
	"fmt"
	"log"
	"os"

	"github.com/jcbhmr/go-clap"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func ExampleQuickStart() {
	os.Args = []string{"01_quick", "--help"}

	matches := clap.Command1().
		Arg(clap.Arg1("[name]", "Optional name to operate on")).
		Arg(clap.Arg1("-c", "--config", "<FILE>", "Sets a custom config file").Required()).
		Arg(clap.Arg1("-d", "--debug", "...", "Turn debugging information on")).
		Subcommand(clap.NewCommand("test").
			About("does testing things").
			Arg(clap.Arg1("-l", "--list", "lists test values")).Action(clap.ArgActionSetTrue)).
		GetMatches()

	if name, ok := matches.GetOneString("name"); ok {
		fmt.Printf("Value for name: %s", name)
	}

	if configPath, ok := matches.GetOneString("config"); ok {
		fmt.Printf("Config file path: %s", configPath)
	}

	x, ok := matches.GetOneU8("debug")
	if !ok {
		log.Fatal("Counts are defaulted")
	}
	switch x {
	case 0:
		fmt.Println("Debug mode is off")
	case 1:
		fmt.Println("Debug mode is kind of on")
	case 2:
		fmt.Println("Debug mode is on")
	default:
		fmt.Println("Don't be crazy")
	}

	if matches2, ok := matches.SubcommandMatches("test"); ok {
		if matches2.GetFlag("list") {
			fmt.Println("Printing testing lists...")
		} else {
			fmt.Println("Not printing testing lists...")
		}
	}

	// Output:
	// A simple to use, efficient, and full-featured Command Line Argument Parser
	//
	// Usage: 01_quick[EXE] [OPTIONS] [name] [COMMAND]
	//
	// Commands:
	//   test  does testing things
	//   help  Print this message or the help of the given subcommand(s)
	//
	// Arguments:
	//   [name]  Optional name to operate on
	//
	// Options:
	//   -c, --config <FILE>  Sets a custom config file
	//   -d, --debug...       Turn debugging information on
	//   -h, --help           Print help
	//   -V, --version        Print version
}

func ExampleConfiguringTheParser1() {
	os.Args = []string{"02_apps", "--help"}

	matches := clap.NewCommand("MyApp").
		Version("1.0").
		About("Does awesome things").
		Arg(clap.Arg1("--two", "<VALUE>").Required()).
		Arg(clap.Arg1("--one", "<VALUE>").Required()).
		GetMatches()

	x, ok := matches.GetOneString("two")
	if !ok {
		log.Fatal("required")
	}
	fmt.Printf("two: %#+v", x)

	x, ok = matches.GetOneString("one")
	if !ok {
		log.Fatal("required")
	}
	fmt.Printf("one: %#+v", x)

	// Output:
	// Does awesome things
	//
	// Usage: 02_apps[EXE] --two <VALUE> --one <VALUE>
	//
	// Options:
	//       --two <VALUE>
	//       --one <VALUE>
	//   -h, --help         Print help
	//   -V, --version      Print version
}
