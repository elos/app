package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
	"text/template"

	"github.com/elos/aeolus"
	"github.com/elos/metis"
	"github.com/elos/metis/templates"
)

const (
	Routes templates.Name = iota
	Router
)

func name(s string) string {
	return strings.Title(metis.CamelCase(s))
}

var aToM = map[aeolus.Action]string{
	aeolus.POST:   "POST",
	aeolus.GET:    "GET",
	aeolus.DELETE: "DELETE",
}

func action(s aeolus.Action) string {
	return aToM[s]
}

func argsFor(e *aeolus.Endpoint) string {
	a := "c"

	if userAuth(e) {
		a += ", u"
	}

	if _, ok := e.Requires["db"]; ok {
		a += ", db"
	}

	if _, ok := e.Requires["sessions"]; ok {
		a += ", sessions"
	}

	if _, ok := e.Requires["agents"]; ok {
		a += ", agents"
	}

	return a
}

func userAuth(e *aeolus.Endpoint) bool {
	return e.Auth == "user"
}

func main() {
	h := aeolus.ParseHostFile("./definitions/hosts/app.json")

	e := templates.NewEngine("./", &templates.TemplateSet{
		Routes: []string{"routes.tmpl"},
		Router: []string{"router.tmpl"},
	}).WithFuncMap(template.FuncMap{
		"name":     name,
		"action":   action,
		"argsFor":  argsFor,
		"userAuth": userAuth,
	})

	if err := e.ParseTemplates(); err != nil {
		log.Fatal(err)
	}

	var buf bytes.Buffer
	if err := e.Execute(&buf, Routes, h); err != nil {
		log.Fatal("Templates: %s", err.Error())
	}

	if err := ioutil.WriteFile("../routes/routes.go", buf.Bytes(), 0644); err != nil {
		log.Fatal(err)
	}
	if err := exec.Command("goimports", "-w=true", "../routes/routes.go").Run(); err != nil {
		log.Fatal(err)
	}

	var b bytes.Buffer
	if err := e.Execute(&b, Router, h); err != nil {
		log.Fatal("Templates: %s", err.Error())
	}
	if err := ioutil.WriteFile("../router.go", b.Bytes(), 0644); err != nil {
		log.Fatal(err)
	}
	if err := exec.Command("goimports", "-w=true", "../router.go").Run(); err != nil {
		log.Fatal(err)
	}
}
