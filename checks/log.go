package checks

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var LogAnalyzer = &analysis.Analyzer{
	Name:     "CheckLogMessage",
	Doc:      "reports lack of log messages in functions",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	filter := []ast.Node{
		(*ast.FuncDecl)(nil),
	}
	inspect.Preorder(filter, func(n ast.Node) {
		fn, ok := n.(*ast.FuncDecl)
		if !ok {
			return
		}

		debugLogExists := false
		for _, line := range fn.Body.List {
			expr, ok := line.(*ast.ExprStmt)
			if !ok {
				continue
			}

			call, ok := expr.X.(*ast.CallExpr)
			if !ok {
				continue
			}

			sel, ok := call.Fun.(*ast.SelectorExpr)
			if !ok {
				continue
			}
			ident, ok := sel.X.(*ast.Ident)
			if !ok {
				return
			}

			if ident.Name == "log" {
				debugLogExists = true
			}
		}

		if !debugLogExists {
			pass.Report(analysis.Diagnostic{
				Pos:     fn.Pos(),
				End:     fn.End(),
				Message: "function should have at least one log function",
			})
		}
		return
	})

	return nil, nil

}
