package main

import (
	"os"
	"text/template"
)

/*
func main() {
	templ := `
package {{.Package}}

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(rand.Intn(100))
	{{.FunctionName}}()
}

func {{.FunctionName}}() {
	fmt.Println("Hello From {{.MyName}}")
}
	`

	data := map[string]string{
		"Package":      "main",
		"FunctionName": "Hello",
		"MyName":       "World",
	}

	t := template.Must(template.New("main").Parse(templ))
	t.Execute(os.Stdout, data)
}
*/

func main() {
	templ := `package {{.Package}}

import "fmt"

type {{.StructName}} struct {
	{{range .Fields}} {{.Name}} {{.Type}} {{end}}
}

func printStruct(s {{.StructName}}) {
	fmt.Println(s)
}

func main() {
	s := {{.StructName}}{
		{{range .Fields}} {{.Name}}: {{.Value}}, {{end}}
	}

	printStruct(s)
}
`
	type Field struct {
		Name  string
		Type  string
		Value string
	}

	type Model struct {
		Package   string
		StructNme string
		Fields    []Field
	}

	data := Model{
		Package:   "main",
		StructNme: "User",
		Fields: []Field{
			{Name: "Name", Type: "string", Value: `"John Doe"`},
			{Name: "Age", Type: "int", Value: "30"},
		},
	}

	t := template.Must(template.New("main").Parse(templ))
	t.Execute(os.Stdout, data)
}

// go run main.go > main2.go && go run main2.go
