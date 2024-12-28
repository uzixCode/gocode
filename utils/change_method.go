package utils

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
)

func ModifyMethodFile(filePath, filterMethod, filterPath, changeToMethod string) error {
	// Read the source file
	source, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	// Parse the source code
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filePath, source, parser.AllErrors)
	if err != nil {
		return fmt.Errorf("error parsing source: %w", err)
	}

	// Traverse and modify the AST
	ast.Inspect(node, func(n ast.Node) bool {
		// Look for expression statements
		if callExpr, ok := n.(*ast.CallExpr); ok {
			if selExpr, ok := callExpr.Fun.(*ast.SelectorExpr); ok {
				// Check for the filter method (e.g., POST)
				if selExpr.Sel.Name == filterMethod {
					// Find the first argument (route path)
					if len(callExpr.Args) > 0 {
						if basicLit, ok := callExpr.Args[0].(*ast.BasicLit); ok && basicLit.Value == fmt.Sprintf(`"%s"`, filterPath) {
							// Change the method name to the desired method
							selExpr.Sel.Name = changeToMethod
						}
					}
				}
			}
		}
		return true
	})

	// Print the modified AST back to source code
	var buf bytes.Buffer
	err = printer.Fprint(&buf, fset, node)
	if err != nil {
		return fmt.Errorf("error printing modified code: %w", err)
	}

	// Overwrite the original file with the modified code
	err = os.WriteFile(filePath, buf.Bytes(), 0644)
	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}

	return nil
}
