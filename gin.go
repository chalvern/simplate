package simplate

import (
	"bytes"
	"html/template"

	"github.com/gin-gonic/gin"

	"github.com/chalvern/sugar"
	"github.com/gin-gonic/gin/render"
)

var (
	// GinRender the render of Gin
	GinRender = &GinRendererS{}
)

// for gin https://github.com/gin-gonic/gin

// GinRendererS html render of gin
type GinRendererS struct {
}

// Instance method of HTMLRender
// https://github.com/gin-gonic/gin/blob/9a820cf0054bcd769f785457b7dbd149a7b29fdd/render/html.go#L21
func (hr *GinRendererS) Instance(name string, data interface{}) render.Render {
	if simplateViewPathTemplates[name] == nil {
		sugar.Warnf("no template of name: %s", name)
	}

	layoutFile := defaultLayoutFile

	// body
	var buf bytes.Buffer
	ExecuteViewPathTemplate(&buf, name, data)
	dataT := make(gin.H)
	dataMap, ok := data.(gin.H)
	if ok {
		dataMap["LayoutContent"] = template.HTML(buf.String())
		dataT = dataMap
		// custom layout
		if layout, ok := dataMap["layout"]; ok {
			layoutFile = layout.(string)
		}
	} else {
		dataT["LayoutContent"] = template.HTML(buf.String())
	}
	return render.HTML{
		Template: simplateViewPathTemplates[layoutFile],
		Data:     dataT,
	}
}
