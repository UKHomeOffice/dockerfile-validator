package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	//GET displays the upload form.
	case "GET":
		display(w, "upload", nil)

	//POST takes the uploaded file(s) and saves it to disk.
	case "POST":

		dockerfile, _, _ := r.FormFile("dockerfile")
		dfile, _ := DockerfileRead(dockerfile)
		defer dockerfile.Close()

		rulesfile, _, _ := r.FormFile("rules")
		defer rulesfile.Close()

		rulesdata, _ := ioutil.ReadAll(rulesfile)
		rules, _ := loadRules(rulesdata)
		v := Validation{rules, dfile}
		valid, msg := v.validate()

		if valid {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Docker file is valid")
		} else {
			w.WriteHeader(http.StatusConflict)
			fmt.Fprintf(w, msg)
		}

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}

//Compile templates on start
var templates = template.Must(template.ParseFiles("upload.html"))

//Display the named template
func display(w http.ResponseWriter, tmpl string, data interface{}) {
	templates.ExecuteTemplate(w, tmpl+".html", data)
}
