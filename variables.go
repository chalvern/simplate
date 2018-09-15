package simplate

import (
	"html/template"
	"sync"
)

var (
	// public

	// ViewsPath root path of views
	ViewsPath = "views"

	// private
	simplateViewPathTemplates = make(map[string]*template.Template)
	templatesLock             sync.RWMutex
)
