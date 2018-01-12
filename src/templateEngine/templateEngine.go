package templateEngine

import (
	"html/template"
	"io"

	"github.com/labstack/echo"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func NewHtmlTemplateEngine() *TemplateRenderer {
	return &TemplateRenderer{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
}
