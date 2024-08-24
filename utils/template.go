package utils

import (
	"html/template"
)

func NewTemplate(templateName string) (*template.Template, error) {
	return template.New(templateName).Funcs(
		template.FuncMap{
			"IsString": func(i interface{}) bool {
				_, isString := i.(string)
				return isString
			},
			"arr": func(els ...any) []any {
				return els
			},
		}).ParseFiles(
		"templates/index.html",
		"templates/head.html",
		"templates/header.html",
		"templates/features/add-feature-form.html",
		"templates/features/list.html",
		"templates/features/list-item.html",
		"templates/features/list-item-status.html",
		"templates/features/flag-boolean-list-item.html",
		"templates/features/flag-string-list-item.html",
	)
}
