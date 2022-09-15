package render

import (
	"fmt"
	"html/template"
	"net/http"
)

// renderTemplate a function to render template

func RendersTemplate(w http.ResponseWriter, html string) {
	parseTemplate, _ := template.ParseFiles("./templates/" + html)
	err := parseTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error running template", err)
		return
	}
}