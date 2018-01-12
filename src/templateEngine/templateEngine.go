package templateEngine

import (
	"html/template"
	"io"
	"omt-project/config"

	"github.com/labstack/echo"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func NewHtmlTemplateEngine() *TemplateRenderer {
	c := config.Setting()
	// FIXME
	// Gin is started at omt-project/src
	// normal server is started at omt-project/
	var templatePath string
	if c.AppEnv == "PROD" {
		templatePath = "public/views/*.html"
	} else {
		templatePath = "../public/views/*.html"
	}
	return &TemplateRenderer{
		templates: template.Must(
			template.ParseGlob(templatePath),
		),
	}
}
