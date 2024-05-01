package gen

import (
	"bytes"
	"strconv"
	"strings"
	"text/template"
)

type Registry struct {
	Package    string
	Components []string
}

func generate(templateStr string, data any, outName string) string {
	tmpl, err := template.New(outName).Parse(templateStr)
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		panic(err)
	}

	s := buf.String()
	s = strings.Replace(s, "&lt;", "<", len(s))
	s = strings.Replace(s, "&gt;", ">", len(s))
	return s
}

func Present(list []int, value int) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}

func IntSliceToString(slice []int) string {
	var s string
	for _, v := range slice {
		strconv.Itoa(v)
	}
	return s
}
