package gen

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
)

var reg *Registry = &Registry{
	Components: make([]string, 0),
}

func RegisterComponent[T any]() {
	typeName := reflect.TypeOf((*T)(nil)).Elem().Name()
	reg.Components = append(reg.Components, typeName)
}

func Generate(location string) {

	templates := []string{"world", "store", "component", "query", "add", "get", "remove", "container"}
	var s string
	for _, v := range templates {
		s += codeTemplates[v]
	}

	reg.Package = filepath.Base(location)
	generated := generate(s, reg, "generated")

	outFile, err := os.Create(fmt.Sprintf("%s/%s.go", location, "generated"))
	if err != nil {
		panic(err)
	}
	outFile.WriteString(generated)
	defer outFile.Close()
}
