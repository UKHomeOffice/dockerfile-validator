package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
)

var debug = false
var rules Rules

func (v Validation) validFrom() bool {
	for _, entry := range v.Rules.From {
		if entry == v.Dockerfile.From() {
			return true
		}
	}
	return false
}

func (v Validation) isRootUser() bool {
	if v.Dockerfile.User() == "root" || v.Dockerfile.User() == "" {
		return true
	}
	return false
}

func (v Validation) validate() (bool, string) {

	from := v.Dockerfile.From()

	if !v.validFrom() {
		return false, "FROM not valid"
	}

	if v.isRootUser() {
		if v.Rules.NotAllowRootUser {
			return false, "Running as root"
		}
	}
	return true, ""

}

func main() {

	debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	rules, _ = loadRulesFromFile("rules.yaml")
	log.Println("listening in port 8080")
	http.HandleFunc("/validate", validateHandler)
	http.HandleFunc("/setup", uploadHandler)
	http.HandleFunc("/rules", uploadRulesHandler)
	http.ListenAndServe(":8080", nil)
}
