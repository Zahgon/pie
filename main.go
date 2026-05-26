//go:generate go run generate.go

package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/elliotchance/pie/functions"
)

func check(err error) { _ = "STUB: not implemented"; return }

func getIdentName(e ast.Expr) string { _ = "STUB: not implemented"; return "" }

func getKeyAndElementType(pkg *ast.Package, name string, typeSpec *ast.TypeSpec) (string, string, string, *TypeExplorer) {
	_ = "STUB: not implemented"
	return "", "", "", nil
}

func findType(pkgs map[string]*ast.Package, name string) (packageName, keyType, elementType string, explorer *TypeExplorer) {
	_ = "STUB: not implemented"
	return "", "", "", nil
}

func getType(keyType, elementType string) int { _ = "STUB: not implemented"; return 0 }

func getImports(packageName, s string) (imports []string) { _ = "STUB: not implemented"; return nil }

func getAllImports(packageName string, files []string, explorer *TypeExplorer) (imports []string) {
	_ = "STUB: not implemented"
	return nil
}

// We have to generate imports slightly differently when we are building code
// that will go into its own packages vs an external package.
func isSelfPackage(packageName string) bool { _ = "STUB: not implemented"; return false }

func main() {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, ".", nil, 0)
	check(err)

	for _, arg := range os.Args[1:] {
		mapOrSliceType, fns := getFunctionsFromArg(arg)
		packageName, keyType, elementType, explorer := findType(pkgs, mapOrSliceType)
		kind := getType(keyType, elementType)

		var templates []string
		for _, function := range functions.Functions {
			if fns[0] != "*" && !stringSliceContains(fns, function.Name) {
				continue
			}

			if function.For&kind != 0 {
				templates = append(templates, pieTemplates[function.Name])
			}
		}

		// Aggregate imports.
		t := fmt.Sprintf("package %s\n\n", packageName)

		imports := getAllImports(packageName, templates, explorer)
		if len(imports) > 0 {
			t += "import ("
			for _, imp := range imports {
				t += fmt.Sprintf("\n\t%s", imp)
			}
			t += "\n)\n\n"
		}

		for _, tmpl := range templates {
			i := strings.Index(tmpl, "//")
			t += tmpl[i:] + "\n"
		}

		t = strings.Replace(t, "StringSliceType", mapOrSliceType, -1)
		t = strings.Replace(t, "StringElementType", elementType, -1)
		t = strings.Replace(t, "ElementType", elementType, -1)
		t = strings.Replace(t, "MapType", mapOrSliceType, -1)
		t = strings.Replace(t, "KeyType", elementType, -1)
		t = strings.Replace(t, "KeySliceType", "[]"+keyType, -1)
		t = strings.Replace(t, "SliceType", mapOrSliceType, -1)

		if !explorer.HasEquals() {
			re := regexp.MustCompile(`([\w_]+|[\w_]+\[\w+\])\.Equals\(([^)]+)\)`)
			t = ReplaceAllStringSubmatchFunc(re, t, func(groups []string) string {
				return fmt.Sprintf("%s == %s", groups[1], groups[2])
			})
		}

		if !explorer.HasString() {
			t = strings.Replace(t, "mightBeString.String()", `fmt.Sprintf("%v", mightBeString)`, -1)
		}

		switch kind {
		case functions.ForNumbers:
			t = strings.Replace(t, "ElementZeroValue", "0", -1)

		case functions.ForStrings:
			t = strings.Replace(t, "ElementZeroValue", `""`, -1)

		case functions.ForStructs:
			zeroValue := fmt.Sprintf("%s{}", elementType)

			// If its a pointer we need to replace '*' -> '&' when
			// instantiating.
			if elementType[0] == '*' || explorer.IsInterface {
				zeroValue = "nil"
			}

			t = strings.Replace(t, "ElementZeroValue", zeroValue, -1)
		}

		if isSelfPackage(packageName) {
			t = strings.Replace(t, "pie.Strings", "Strings", -1)
			t = strings.Replace(t, "pie.Ints", "Ints", -1)
			t = strings.Replace(t, "pie.Float64s", "Float64s", -1)
		}

		// The TrimRight is important to remove an extra new line that conflicts
		// with go fmt.
		t = strings.TrimRight(t, "\n") + "\n"

		err := ioutil.WriteFile(strings.ToLower(mapOrSliceType)+"_pie.go", []byte(t), 0755)
		check(err)
	}
}

func getFunctionsFromArg(arg string) (mapOrSliceType string, fns []string) {
	_ = "STUB: not implemented"
	return "", nil
}

func stringSliceContains(haystack []string, needle string) bool {
	_ = "STUB: not implemented"
	return false
}

// http://elliot.land/post/go-replace-string-with-regular-expression-callback
func ReplaceAllStringSubmatchFunc(re *regexp.Regexp, str string, repl func([]string) string) string {
	_ = "STUB: not implemented"
	return ""
}

func isNegative(b byte) bool { _ = "STUB: not implemented"; return false }

func addBrackets(str string) string { _ = "STUB: not implemented"; return "" }
