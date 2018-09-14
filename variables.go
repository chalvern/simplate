package simplate

import (
	"html/template"
	"sync"

	"github.com/sirupsen/logrus"
)

var (
	// public

	// ViewsPath root path of views
	ViewsPath = "views"
	// Logs logs
	Logs = logrus.New

	// private
	simplateViewPathTemplates = make(map[string]*template.Template)
	templatesLock             sync.RWMutex
)
