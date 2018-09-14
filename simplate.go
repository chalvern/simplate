package simplate

import (
	"html/template"
)

// InitTemplate init templates
func InitTemplate() error {

	simplateViewPathTemplates = make(map[string]*template.Template)

	return BuildTemplate(ViewsPath)
}
