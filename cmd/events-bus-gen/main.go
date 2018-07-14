package main

import (
	"fmt"
	"github.com/reddec/astools"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
	"path"
	"strings"
	"text/template"
)

func Signature(m atool.Method, f *atool.File) string {
	t := ""
	for i, arg := range m.In {
		if i != 0 {
			t += ", "
		}
		t += arg.Name + " " + f.Printer.ToString(arg.Type)
	}
	return t
}

func Call(m atool.Method) string {
	t := ""
	for i, arg := range m.In {
		if i != 0 {
			t += ", "
		}
		t += arg.Name
	}
	return t
}

func Import(imp, ali string) string {
	if ali == "" {
		return imp
	}

	return ali + " " + imp
}

type params struct {
	Package string
	Imports map[string]string
	Name    string
	File    string
	Events  []atool.Method
}

func (p *params) EventsList() string {
	t := ""
	for _, event := range p.Events {
		if t != "" {
			t += ", "
		}
		t += event.Name
	}
	return t
}

func main() {
	output := kingpin.Flag("output", "Output file name").Short('o').String()
	ifaceName := kingpin.Arg("interface", "Interface name").Required().String()
	file := kingpin.Arg("file", "File with interfaces").Required().String()
	kingpin.Parse()
	info, err := atool.Scan(*file)

	if err != nil {
		panic(err)
	}

	var parentDir string = path.Dir(*file)

	if *output == "" {
		*output = path.Join(parentDir, strings.ToLower(*ifaceName)+".bus.gen.go")
	} else {
		parentDir = path.Dir(*output)
	}

	err = os.MkdirAll(parentDir, 0755)
	if err != nil {
		panic(err)
	}

	var imports = map[string]string{
		"\"reflect\"": "",
		"\"sync\"":    "",
	}

	for imp, ali := range info.Imports {
		imports[imp] = ali
	}

	var busTemplate = template.Must(template.New("").Funcs(template.FuncMap{
		"title": strings.Title,
		"signature": func(m atool.Method) string {
			return Signature(m, info)
		},
		"import": func(imp, ali string) string {
			return Import(imp, ali)
		},
		"call": Call,
	}).Parse(text))
	f, err := os.Create(*output)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	for _, iface := range info.Interfaces {
		if iface.Name == *ifaceName {
			params := params{
				Package: info.Package,
				Name:    iface.Name,
				Imports: imports,
				File:    *file,
				Events:  []atool.Method{},
			}

			for _, method := range iface.Methods {
				if len(method.Out) == 0 {
					params.Events = append(params.Events, *method)
				}
			}
			if len(params.Events) > 0 {
				err := busTemplate.Execute(f, &params)
				if err != nil {
					fmt.Println(err)
				}
			}
			break
		}
	}

}
