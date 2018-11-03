package simplate

import (
	"bytes"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	ViewsPath = "views"

	// template function
	dataFormatFunc := func(t time.Time) string {
		return t.UTC().Format("2006年01月02日03时04分05秒UTC")
	}
	AddFuncMap("dataFormat", dataFormatFunc)

	// initial template
	InitTemplate()
	m.Run()
}
func TestSimplateViewPathTemplates(t *testing.T) {
	assert.NotEqual(t, 0, len(simplateViewPathTemplates))

	a := []string{}
	for k, v := range simplateViewPathTemplates {
		a = append(a, k)
		err := v.Execute(os.Stdout, make(map[string]interface{}))
		assert.Nil(t, err)
	}

	assert.Contains(t, a, "home/body.tpl")
	assert.Contains(t, a, "home/head.tpl")
	assert.Contains(t, a, "home/index.html")
}

func TestExecuteTemplate(t *testing.T) {
	timestamp := 1541222924
	tm := time.Unix(int64(timestamp), 0)

	data := make(map[string]interface{})
	data["STime"] = tm

	var buf bytes.Buffer
	err := ExecuteTemplate(&buf, "home/index.html", data)
	assert.Nil(t, err)
	assert.Contains(t, buf.String(), "2018年11月03日05时28分44秒UTC")
}
