package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	sages := map[string]string{
		"India":    "Gandhi",
		"USA":      "MLK",
		"Meditate": "Buddha",
		"Love":     "Jesus",
	}

	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}

}
