package render

import (
	"booking/pkg/config"
	"booking/pkg/models"
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

type templateMap map[string]*template.Template

var functions = template.FuncMap{}

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(templateData *models.TemplateData) *models.TemplateData {

	return templateData
}

func RenderTemplate(w http.ResponseWriter, tmpl string, data *models.TemplateData) {

	var templatesCached templateMap

	if app.UseCache {
		templatesCached = app.TemplateCache
	} else {
		templatesCached, _ = CreateTemplateCache()
	}

	templateCached, ok := templatesCached[tmpl]

	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	data = AddDefaultData(data)

	_ = templateCached.Execute(buf, data)

	_, err := buf.WriteTo(w)

	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}
}

func CreateTemplateCache() (templateMap, error) {
	myCahce := templateMap{}

	pages, err := filepath.Glob("./templates/*.page.html")

	if err != nil {
		return myCahce, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)

		if err != nil {
			return myCahce, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")

		if err != nil {
			return myCahce, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")

			if err != nil {
				return myCahce, err
			}
		}

		myCahce[name] = ts
	}

	return myCahce, nil
}
