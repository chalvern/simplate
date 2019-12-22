package simplate

import (
	"html/template"
	"sync"
)

var (

	// defaultViewsPath default root path of views
	defaultViewsPath = "views"
	// defaultLayoutFile default layout file path
	defaultLayoutFile = "layout/default.html"

	// templates
	simplateViewPathTemplates = make(map[string]*template.Template)
	templatesLock             sync.RWMutex

	// functions
	simplateTplFuncMap = make(template.FuncMap)
)

// SetViewsPath set view's path
func SetViewsPath(path string) {
	defaultViewsPath = path
}

// SetLayoutFile set layout's file
func SetLayoutFile(layout string) {
	defaultLayoutFile = layout
}
