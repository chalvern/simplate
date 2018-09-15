## simplate = simple template of golang

abstract from [beego](https://github.com/astaxie/beego), part of its template.

SIMPLATE would load all template files in ViewsPath, and keep all template.Template, 
indexed by VIEW-PATH-NAME, in map. SIMPLATE can avoid loading template file repeatly
into memory, to save time.

## Quick Start

### import the package

```golang
import "github.com/chalvern/simplate"
```

### create template dir and files

SIMPLATE has a variable named `ViewsPath` keep the template's directory. 
SIMPLATE will load all template files in the `ViewsPath` into memory.

### init simplate

```golang
import "github.com/chalvern/simplate"

func InitIsmplate(){
  simplate.ViewsPath  = "your-templates-dir" // default is "views"
  simplate.LayoutFile = "your-layout-file" // default is "layout/default.html"

  // initial
  simplate.InitTemplate()
}

func YourCode() error {
  data := make(map[string]interface{}
  // data["Jingwei"] = "https://jingwei.link"
  return simplate.ExecuteTemplate(os.Stdout,"home/body.html", data)
}

```

## more doc

SIMPLATE keeps all template.Template in private variable `simplateViewPathTemplates`.
As your template dir hierarchy as:

```bash

views
|--home
    |--index.html
    |--head.tpl
    |--body.tpl
|--layout
    |--default.html

```

the keys of `simplateViewPathTemplates` would be `[]string{"home/body.tpl","home/head.tpl","home/index.html","layout/default.html"}`.

Finally, enjoy your coding.
