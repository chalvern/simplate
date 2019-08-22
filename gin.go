package simplate

import (
	"github.com/chalvern/sugar"
	"github.com/gin-gonic/gin/render"
)

// for gin https://github.com/gin-gonic/gin

// GinRenderer html render of gin
type GinRenderer struct {
}

// Instance method of HTMLRender
// https://github.com/gin-gonic/gin/blob/9a820cf0054bcd769f785457b7dbd149a7b29fdd/render/html.go#L21
func (hr *GinRenderer) Instance(name string, data interface{}) render.Render {
	if simplateViewPathTemplates[name] == nil {
		sugar.Warnf("no template of name: %s", name)
	}
	return render.HTML{
		Template: simplateViewPathTemplates[name],
		Data:     data,
	}
}
