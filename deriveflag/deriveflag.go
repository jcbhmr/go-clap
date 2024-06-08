package deriveflag

import (
	"flag"
	"fmt"
	"go/ast"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/tools/go/packages"
)

type StringsValue struct {A *[]string}

func (c StringsValue) String() string {
	return strings.Join(*c.A, ",")
}

func (c StringsValue) Set(value string) error {
	*c.A = strings.Split(value, ",")
	return nil
}

var Usage = func() {
	w := flag.CommandLine.Output()
	fmt.Fprintf(w, "#[derive(%s)] for Go\n", Name)
	fmt.Fprintf(w, "Usage of %s:\n", Name)
	flag.PrintDefaults()
}

var Name = filepath.Base(os.Args[0])
var typeFlag = flag.String("type", "", "comma-separated list of type names; must be set")
var Type []*ast.GenDecl
var outputFlag = flag.String("output", "", "output file name; default srcdir/<yourtype>_derive_<derivename>.go")
var Output *os.File
var Tags []string

func init() {
	flag.Usage = func() {
		Usage()
	}
	flag.Var(StringsValue{&Tags}, "tags", "comma-separated list of build tags to apply")
}

func fatalUsagef(format string, args ...any) {
	w := flag.CommandLine.Output()
	fmt.Fprintf(w, format, args...)
	fmt.Fprintln(w)
	flag.Usage()
	os.Exit(2)
}

func Parse() {
	flag.Parse()

	typeNames := strings.Split(*typeFlag, ",")
	if len(typeNames) == 0 {
		fatalUsagef("flag -type must be set")
	}

	var patterns []string
	if flag.NArg() == 0 {
		patterns = []string{"."}
	} else {
		patterns = flag.Args()
	}

	canHaveTags := false
	if len(patterns) == 1 {
		stats, err := os.Stat(patterns[0])
		if err != nil {
			log.Fatalf("os.Stat() %s: %v", patterns[0], err)
		}
		if stats.IsDir() {
			canHaveTags = true
		}
	}
	if !canHaveTags && len(Tags) > 0 {
		log.Fatalf("-tags option applies only to directories, not when files are specified")
	}

	var buildFlags []string
	if len(Tags) > 0 {
		buildFlags = []string{"-tags=" + strings.Join(Tags, " ")}
	}
	cfg := &packages.Config{
		Mode:       packages.NeedName | packages.NeedTypes | packages.NeedTypesInfo | packages.NeedSyntax,
		Tests:      false,
		BuildFlags: buildFlags,
		Logf:       log.Printf,
	}
	pkgs, err := packages.Load(cfg, patterns...)
	if err != nil {
		log.Fatalf("packages.Load() %v: %v", patterns, err)
	}
	if len(pkgs) != 1 {
		log.Fatalf("len(pkgs) expected 1 got %d", len(pkgs))
	}

	Type = []*ast.GenDecl{}
	f := func(n ast.Node) bool {
		genDecl, ok := n.(*ast.GenDecl)
		if !ok {
			return true
		}
		if genDecl.Tok != token.TYPE {
			return true
		}
		if len(genDecl.Specs) != 1 {
			return true
		}
		typeSpec, ok := genDecl.Specs[0].(*ast.TypeSpec)
		if !ok {
			return true
		}
		for _, typeName := range typeNames {
			if typeName == typeSpec.Name.Name {
				Type = append(Type, genDecl)
				break
			}
		}
		return true
	}
	for _, file := range pkgs[0].Syntax {
		ast.Inspect(file, f)
	}

	var path string
	if *outputFlag == "" {
		dirname := filepath.Dir(patterns[0])
		filename := strings.ToLower(typeNames[0]) + "_derive_" + strings.ToLower(Name) + ".go"
		path = filepath.Join(dirname, filename)
	} else {
		path = *outputFlag
	}
	Output, err = os.Create(path)
	if err != nil {
		log.Fatalf("os.Create() %s: %v", path, err)
	}
}
