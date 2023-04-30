package html

import (
	"mag"
	"text/template"
)

var files = mag.Content()

func parse(file ...string) *template.Template {
	file = append([]string{"layout.html"}, file...)
	return template.Must(
		template.New("layout.html").ParseFS(files, file...))
}
