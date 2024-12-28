package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

func ScanRoutes(folder string) {
	// entries, err := os.ReadDir(folder)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	filePath := "routes.go"
	route, _, err := ScanningRoute(filepath.Join(folder, filePath))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	js, err := json.Marshal(route)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(js))

}
func ScanningRoute(path string) (route map[string]interface{}, group string, err error) {
	var routeList map[string]interface{} = make(map[string]interface{})
	// var groupMap map[string]interface{} = make(map[string]interface{})
	var importedMap map[string]map[string]interface{} = make(map[string]map[string]interface{})
	srcFile, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, "", err
	}
	fs := token.NewFileSet()
	node, err := parser.ParseFile(fs, path, srcFile, parser.AllErrors)
	if err != nil {
		fmt.Println("Error parsing file:", err)
		return nil, "", err
	}
	ginParameterName := ""
	groupVariableName := ""
	currentDir, err := os.Getwd()
	if err != nil {
		return nil, "", err
	}
	dirName := filepath.Base(currentDir)
	// grp := ""
	ast.Inspect(node, func(n ast.Node) bool {
		//Import-----------------------------
		for _, v := range node.Imports {
			if !strings.Contains(v.Path.Value, dirName) {
				continue
			}
			index := strings.Index(v.Path.Value, dirName)
			remainingPath := v.Path.Value[index+len(dirName)+1:]
			combinedPath := filepath.Join(currentDir, remainingPath)
			pck := strings.Trim(filepath.Base(remainingPath), "\"c")
			importedMap[pck] = make(map[string]interface{})
			importedMap[pck]["path"] = remainingPath
			importedMap[pck]["full_path"] = combinedPath
		}
		if funcDecl, ok := n.(*ast.FuncDecl); ok {
			if funcDecl.Name.Name == "Routes" {
				//Get gin Parameter Name-----------------------------
				if funcDecl.Type.Params != nil {
					for _, param := range funcDecl.Type.Params.List {
						for _, name := range param.Names {
							if exprToString(param.Type) == "*gin.Engine" || exprToString(param.Type) == "*gin.RouterGroup" {
								ginParameterName = name.Name
								break
							}
						}
						if ginParameterName != "" {
							break
						}
					}
				}
				//Function Body Section-----------------------------
				if funcDecl.Body != nil {
					ast.Inspect(funcDecl.Body, func(n ast.Node) bool {
						if assign, ok := n.(*ast.AssignStmt); ok {
							if callExpr, ok := assign.Rhs[0].(*ast.CallExpr); ok {
								if selExpr, ok := callExpr.Fun.(*ast.SelectorExpr); ok {
									if _, ok := selExpr.X.(*ast.Ident); ok {
										if ident, ok := assign.Lhs[0].(*ast.Ident); ok {
											if selExpr.Sel.Name == "Group" && groupVariableName == "" {
												// gt := map[string]interface{}{"class": pkgIdent.Name, "path": strings.Trim(exprCallFuncToString(callExpr.Args[0]), "\"")}
												// if _, o := groupMap[pkgIdent.Name]; o {
												// 	gt["parent"] = pkgIdent.Name
												// }
												// groupMap[ident.Name] = gt
												// fmt.Printf("Function %s.%s assigned: %s\n", pkgIdent.Name, selExpr.Sel.Name, ident.Name)
												groupVariableName = ident.Name
												routeList["type"] = "Group"
												routeList["path"] = strings.Trim(exprCallFuncToString(callExpr.Args[0]), "\"")
												routeList["location"] = path
												routeList["route_location"] = path
												routeList["variable_name"] = ident.Name
												absPath, err := filepath.Abs(path)
												if err == nil {
													routeList["full_location"] = absPath

												}
												routeList["filename"] = filepath.Base(path)
											}
										}
									}
								}
							}
						}
						if callExpr, ok := n.(*ast.CallExpr); ok {
							if selExpr, ok := callExpr.Fun.(*ast.SelectorExpr); ok {
								if pkgIdent, ok := selExpr.X.(*ast.Ident); ok {
									v := ginParameterName
									if groupVariableName != "" {
										v = groupVariableName
									}
									if pkgIdent.Name == v {
										var rt map[string]interface{} = map[string]interface{}{}
										if selExpr.Sel.String() == "Group" {

										} else if selExpr.Sel.String() == "Use" {
											t := map[string]interface{}{}
											p, err := findFileWithFunction(filepath.Dir(path), strings.Trim(exprCallFuncToString(callExpr.Args[0]), "\""))
											if err == nil {
												t["location"] = p
												filename := filepath.Base(p)
												t["filename"] = filename
												absolute, err := filepath.Abs(p)
												if err == nil {
													t["full_location"] = absolute
												}

											}
											if _, ok := routeList["middleware"]; !ok {
												routeList["middleware"] = make(map[string]interface{})
											}
											t["route_location"] = path
											routeList["middleware"] = t
										} else {
											rt["method"] = selExpr.Sel.String()
											rt["path"] = strings.Trim(exprCallFuncToString(callExpr.Args[0]), "\"")
											rt["type"] = "Endpoint"
											rt["route_location"] = path

											if len(callExpr.Args) > 1 {
												pckSelector, ok := callExpr.Args[1].(*ast.SelectorExpr)
												if ok {
													p := fmt.Sprintf("./%s", strings.Trim(importedMap[exprToString(pckSelector.X)]["path"].(string), "\""))
													p = filepath.Join(p)
													fi, err := findFileWithFunction(p, pckSelector.Sel.Name)
													if err != nil {
														fmt.Println(err.Error())
													}
													// fmt.Println(filepath.Base(fi))
													fullLocation, isAbsOkay := filepath.Abs(fi)
													// fmt.Println(fullLocation)
													rt["filename"] = filepath.Base(fi)
													rt["location"] = fi
													if isAbsOkay == nil {
														rt["full_location"] = fullLocation
													}
													rt["handler"] = exprCallFuncToString(callExpr.Args[1])
												} else {
													rt["handler"] = exprCallFuncToString(callExpr.Args[1])
												}
											}
										}
										if v, ok := routeList["type"]; ok && v.(string) == "Group" {
											if _, ok := routeList["routes"]; !ok {
												routeList["routes"] = make(map[string]interface{})
											}
											if v, ok := routeList["routes"].(map[string]interface{}); ok {
												p, pok := rt["path"].(string)
												if pok {
													v[p] = rt

												}

											}
										} else {
											routeList[rt["path"].(string)] = rt
										}

									} else {
										v, ok := importedMap[pkgIdent.Name]["path"]
										if ok {
											p, err := findFileWithFunction(strings.Trim(v.(string), "\""), selExpr.Sel.String())
											if err == nil {
												route, group, err := ScanningRoute(p)
												if err != nil {
													fmt.Println(err.Error())
												}
												if err == nil {
													if group != "" {
														if v, ok := routeList["type"]; ok && v.(string) == "Group" {
															if _, exists := routeList["routes"]; !exists {
																routeList["routes"] = make(map[string]interface{})
															}
															groupMap, ok := routeList["routes"].(map[string]interface{})
															if ok {
																groupMap[group] = route
															}

														} else {
															routeList[group] = route

														}

													} else {
														for key, value := range routeList {
															routeList[key] = value
														}
													}
												}
											}

										}
									}

								}

							}

						}

						return true
					})
				}

			}
		}
		return true
	})
	srcFile.Close()
	if v, ok := routeList["type"]; ok && v.(string) == "Group" {
		return routeList, routeList["path"].(string), nil
	}
	return routeList, "", nil

}
func formatNode(n ast.Node, fs *token.FileSet) string {
	var sb strings.Builder
	if err := format.Node(&sb, fs, n); err != nil {
		return fmt.Sprintf("Error formatting node: %v", err)
	}
	return sb.String()
}
func exprCallFuncToString(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name // Variable or identifier
	case *ast.BasicLit:
		return t.Value // Literal values (e.g., strings, numbers)
	case *ast.SelectorExpr:
		return exprToString(t.X) + "." + t.Sel.Name // Qualified identifiers
	case *ast.StarExpr:
		return "*" + exprToString(t.X) // Pointers
	case *ast.CallExpr:
		return exprToString(t.Fun) + "(...)" // Nested function calls
	default:
		return "<unknown>"
	}
}
func exprToString(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.ArrayType:
		return "[]" + exprToString(t.Elt)
	case *ast.StarExpr:
		return "*" + exprToString(t.X)
	case *ast.SelectorExpr:
		return exprToString(t.X) + "." + t.Sel.Name
	default:
		return "<unknown>"
	}
}
func findFileWithFunction(dir, functionName string) (string, error) {
	// Open the directory
	files, err := os.ReadDir(dir)
	if err != nil {
		return "", err
	}

	// Iterate over the files in the directory
	for _, file := range files {
		if file.IsDir() {
			continue // Skip subdirectories
		}

		// Process only .go files
		if filepath.Ext(file.Name()) == ".go" {
			filePath := filepath.Join(dir, file.Name())
			if containsFunctionInFile(filePath, functionName) {
				return filePath, nil // Return the first file that matches
			}
		}
	}

	return "", nil // No file found
}

// containsFunctionInFile checks if a file contains a specific function name
func containsFunctionInFile(filePath, functionName string) bool {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Failed to open file:", filePath, "-", err)
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	target := "func " + functionName + "(" // Looking for function definition
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, target) {
			return true
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", filePath, "-", err)
	}
	return false
}
