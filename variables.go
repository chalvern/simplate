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

	// private
	simplateViewPathTemplates = make(map[string]*template.Template)
	templatesLock             sync.RWMutex
)
