package simplate

import (
	"html/template"
	"sync"
)

var (
	// public

	// ViewsPath root path of views
	viewsPath = "views"
	// LayoutFile layout file path
	layoutFile = "layout/default.html"

	//
	// private variables

	// templates
	simplateViewPathTemplates = make(map[string]*template.Template)
	templatesLock             sync.RWMutex

	// functions
	simplateTplFuncMap = make(template.FuncMap)
)

// SetViewsPath set view's path
func SetViewsPath(path string) {
	viewsPath = path
}

// SetLayoutFile set layout's file
func SetLayoutFile(layoutFile string) {
	layoutFile = layoutFile
}
