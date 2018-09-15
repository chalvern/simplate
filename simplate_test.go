package simplate

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	ViewsPath = "views"
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
	data := make(map[string]interface{})
	err := ExecuteTemplate(os.Stdout, "home/index.html", data)
	assert.Nil(t, err)
}
