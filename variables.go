package simplate

import (
	"html/template"
	"sync"
)

var (
	// public

	// ViewsPath root path of views
	ViewsPath = "views"
	// LayoutFile layout file path
	LayoutFile = "layout/default.html"

	//
	// private variables

	// templates
	simplateViewPathTemplates = make(map[string]*template.Template)
	templatesLock             sync.RWMutex

	// functions
	simplateTplFuncMap = make(template.FuncMap)
)
