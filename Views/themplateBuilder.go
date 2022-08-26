package views

import (
	"html/template"
	"os"
)

type TemplateBuilder func() (*template.Template, error)

type TemplateVeriables struct {
	Builder TemplateBuilder
}

func BuildThemplate(layout string, page string) TemplateVeriables {
	dir, _ := os.Getwd()
	layout = dir + "/Views/Layouts/" + layout
	page = dir + "/Views/Pages/" + page
	return TemplateVeriables{func() (*template.Template, error) {
		tmpl, err := template.New("").ParseFiles(layout, page)
		if err != nil {
			return template.New(""), err
		}

		return tmpl, nil
	}}
}
