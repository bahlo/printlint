// Package printcheck defines an Analyzer that reports usage of `fmt.Println`
// and friends.
package printcheck

import (
	"bytes"
	"go/ast"
	"go/printer"
	"go/token"

	"golang.org/x/tools/go/analysis"
)

// Analyzer is the analyzer for printcheck.
var Analyzer = &analysis.Analyzer{
	Name: "printlint",
	Doc:  "reports usage of fmt.Print",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	badFuncs := []string{"Print", "Println", "Printf"}

	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			ce, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			se, ok := ce.Fun.(*ast.SelectorExpr)
			if !ok {
				return true
			}

			id, ok := se.X.(*ast.Ident)
			if !ok {
				return true
			}

			// We only want to look at fmt.* calls.
			if id.Name != "fmt" {
				return true
			}

			// We only want to warn for Print{,f,ln}.
			for _, fn := range badFuncs {
				if se.Sel.Name == fn {
					pass.Reportf(ce.Pos(), "fmt.%s found %q", fn,
						prettyPrint(pass.Fset, ce))
					return false
				}
			}

			return true
		})
	}

	return nil, nil
}

// prettyPrint returns the pretty-print of the given node.
func prettyPrint(fset *token.FileSet, x interface{}) string {
	var buf bytes.Buffer
	if err := printer.Fprint(&buf, fset, x); err != nil {
		panic(err)
	}
	return buf.String()
}
