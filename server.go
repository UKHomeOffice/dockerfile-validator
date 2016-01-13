package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

func uploadRulesHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	//GET displays the upload form.
	case "GET":
		message := "Rules currently defined: \n" + rules.String()
		w.Write([]byte(message))

	//POST takes the uploaded file(s) and saves it to disk.
	case "POST":
		rulesfile, _, _ := r.FormFile("rules")
		defer rulesfile.Close()

		rulesdata, _ := ioutil.ReadAll(rulesfile)
		rules, _ = loadRules(rulesdata)
		fmt.Fprintf(w, "Rules file uploaded")
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}

func validateHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	//GET displays the upload form.
	case "GET":
		display(w, "upload", nil)

	//POST takes the uploaded file(s) and saves it to disk.
	case "POST":
		if r.Header.Get("dockerfile") != "" {
			dockerfile, _, _ := r.FormFile("dockerfile")
			dfile, _ := DockerfileRead(dockerfile)
			defer dockerfile.Close()

			v := Validation{rules, dfile}
			valid, msg := v.validate()
			if valid {
				w.WriteHeader(http.StatusOK)
				return
			} else {
				w.WriteHeader(http.StatusConflict)
				fmt.Fprintf(w, msg)
			}
		}

		w.WriteHeader(http.StatusConflict)
		fmt.Fprintf(w, "No Dockerfile found in request")

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

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
