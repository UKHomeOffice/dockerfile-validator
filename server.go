package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func uploadRulesHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	//GET displays the upload form.
	case "GET":
		message := "Rules currently defined: \n\n" + rules.String()
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
		message := "Post your Dockerfile to validate it against the rules: \n\n" + rules.String()
		w.Write([]byte(message))

	//POST takes the uploaded file(s) and saves it to disk.
	case "POST":

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

		// w.WriteHeader(http.StatusConflict)
		// fmt.Fprintf(w, "No Dockerfile found in request")

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	//GET displays the upload form.
	case "GET":
		message := "Post your Dockerfile and Rules to check if it's valid"
		w.Write([]byte(message))

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

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	message := "Dockerfile Validator. Upload your Dockerfile to test if it's complaiant with the rules"
	w.Write([]byte(message))
}
