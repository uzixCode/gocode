package utils

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
)

// StructInfo stores the file name and a single struct name
type StructInfo struct {
	FileName   string
	StructName string
}

func getFirstStructNameFromFile(filePath string) (string, error) {
	// Create a new file set
	fset := token.NewFileSet()

	// Parse the file
	node, err := parser.ParseFile(fset, filePath, nil, parser.AllErrors)
	if err != nil {
		return "", err
	}

	// Walk through the AST and find the first struct declaration
	for _, decl := range node.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.TYPE {
			continue
		}

		for _, spec := range genDecl.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}

			_, ok = typeSpec.Type.(*ast.StructType)
			if ok {
				return typeSpec.Name.Name, nil
			}
		}
	}

	return "", nil // No struct found
}

func GetAllStructNamesInFolder(folderPath string) ([]StructInfo, error) {
	var results []StructInfo

	// Walk through the folder
	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if it's a Go file
		if !info.IsDir() && filepath.Ext(info.Name()) == ".go" {
			structName, err := getFirstStructNameFromFile(path)
			if err != nil {
				return err
			}
			if structName != "" {
				results = append(results, StructInfo{
					FileName:   info.Name(), // Use only the file name
					StructName: structName,
				})
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return results, nil
}
