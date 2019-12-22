package simplate

import (
	"bytes"
	"html/template"
	"io"

	"github.com/chalvern/sugar"
)

// InitTemplate init templates
func InitTemplate() error {

	simplateViewPathTemplates = make(map[string]*template.Template)

	return BuildTemplate(defaultViewsPath)
}

// ExecuteTemplate execute template with default layout file.
func ExecuteTemplate(wr io.Writer, bodyName string, data map[string]interface{}) error {
	if layout, ok := data["layout"]; ok {
		return ExecuteViewPathTemplateWithLayout(wr, layout.(string), bodyName, data)
	}
	return ExecuteViewPathTemplateWithLayout(wr, defaultLayoutFile, bodyName, data)
}

// ExecuteViewPathTemplateWithLayout excute template with layout
func ExecuteViewPathTemplateWithLayout(wr io.Writer, layoutName, bodyName string, data map[string]interface{}) error {
	// body
	var buf bytes.Buffer
	ExecuteViewPathTemplate(&buf, bodyName, data)
	data["LayoutContent"] = template.HTML(buf.String())

	// layout
	return ExecuteViewPathTemplate(wr, layoutName, data)
}

// ExecuteViewPathTemplate applies the template with name and from specific viewPath to the specified data object,
// writing the output to wr.
func ExecuteViewPathTemplate(wr io.Writer, name string, data interface{}) error {
	if t, ok := simplateViewPathTemplates[name]; ok {
		var err error
		if t.Lookup(name) != nil {
			err = t.ExecuteTemplate(wr, name, data)
		} else {
			err = t.Execute(wr, data)
		}
		if err != nil {
			sugar.Error("template Execute err:", err)
		}
		return err
	}
	panic("can't find templatefile in the path:" + defaultViewsPath + "/" + name)
}
