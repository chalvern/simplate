package simplate

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimplateViewPathTemplates(t *testing.T) {
	ViewsPath = "views"
	InitTemplate()
	assert.NotEqual(t, 0, len(simplateViewPathTemplates))

	keys := []string{"home/body.tpl", "home/head.tpl", "home/index.html"}
	for k, v := range simplateViewPathTemplates {
		assert.Contains(t, keys, k)
		v.Execute(os.Stdout, k)
	}
}
