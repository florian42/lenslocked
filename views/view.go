package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var (
	/*
		LayoutDir defines the dir where the layouts are stored.
	*/
	LayoutDir = "views/layouts/"
	/*
		TemplateExt defines the extension of the templates for the layouts.
	*/
	TemplateExt = ".html"
	TemplateDir = "views/"
)

func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}
	return files
}

/*
NewView handles appending common template files, parses them and constructs a View and returns it.
*/
func NewView(layout string, files ...string) *View {
	addTemplatePath(files)
	addTemplateExt(files)
	files = append(files, layoutFiles()...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
		Layout:   layout,
	}
}

/*
View stores template that we want to execute.
*/
type View struct {
	Template *template.Template
	Layout   string
}

/*
Render renders a view.
*/
func (v *View) Render(w http.ResponseWriter,
	data interface{}) error {
	w.Header().Set("Content-Type", "text/html")
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := v.Render(w, nil); err != nil {
		panic(err)
	}
}

func addTemplatePath(files []string) {
	for i, f := range files {
		files[i] = TemplateDir + f
	}
}

func addTemplateExt(files []string) {
	for i, f := range files {
		files[i] = f + TemplateExt
	}
}
