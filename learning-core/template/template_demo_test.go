package template

import (
	"os"
	"testing"
	"text/template"
)

func TestDemo(t *testing.T) {
	p := Person{"longshuai", 0}
	tmpl, _ := template.New("test").Parse("Name: {{.Name}}, Age: {{.Age}}")
	_ = tmpl.Execute(os.Stdout, p)
}
