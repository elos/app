package main

import (
	"bytes"
	"fmt"
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
	RoutesContext
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

func isDynamic(e *aeolus.Endpoint) bool {
	return strings.Contains(e.Path, ":")
}

func signatureFor(e *aeolus.Endpoint) string {
	var buf bytes.Buffer
	tokens := strings.Split(e.Path, "/")
	args := make([]string, 0)
	for i := range tokens {
		if strings.Contains(tokens[i], ":") {
			args = append(args, string(tokens[i][1:]))
		}
	}

	fmt.Fprint(&buf, "func (")
	for i, arg := range args {
		if i != 0 {
			fmt.Fprint(&buf, ",")
		}
		fmt.Fprintf(&buf, "%s", arg)
	}
	if len(args) > 0 {
		fmt.Fprint(&buf, " string) string")
	} else {
		fmt.Fprint(&buf, ") string")
	}
	return buf.String()
}

func interpolatorFor(e *aeolus.Endpoint) string {
	var buf bytes.Buffer
	tokens := strings.Split(e.Path, "/")
	args := make([]string, 0)
	for i := range tokens {
		if strings.Contains(tokens[i], ":") {
			args = append(args, string(tokens[i][1:]))
		}
	}

	fmt.Fprintf(&buf, "func (r *RoutesContext) %s(", name(e.Name))
	for i, arg := range args {
		if i != 0 {
			fmt.Fprint(&buf, " ,")
		}
		fmt.Fprintf(&buf, "%s", arg)
	}
	if len(args) > 0 {
		fmt.Fprint(&buf, " string) string {")
	} else {
		fmt.Fprint(&buf, ") string {")
	}
	for i, token := range tokens {
		if strings.Contains(token, ":") {
			tokens[i] = "%s"
		}
	}
	fmt.Fprintf(&buf, "return fmt.Sprintf(\"%s\"", strings.Join(tokens, "/"))
	for _, arg := range args {
		fmt.Fprint(&buf, ",")
		fmt.Fprintf(&buf, "%s", arg)
	}
	fmt.Fprint(&buf, ")\n}")

	return buf.String()
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
		Routes:        []string{"routes.tmpl"},
		Router:        []string{"router.tmpl"},
		RoutesContext: []string{"routes_context.tmpl"},
	}).WithFuncMap(template.FuncMap{
		"name":            name,
		"action":          action,
		"argsFor":         argsFor,
		"userAuth":        userAuth,
		"signatureFor":    signatureFor,
		"interpolatorFor": interpolatorFor,
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

	var bb bytes.Buffer
	if err := e.Execute(&bb, RoutesContext, h); err != nil {
		log.Fatal("Templates: %s", err.Error())
	}
	if err := ioutil.WriteFile("../views/routes_context.go", bb.Bytes(), 0644); err != nil {
		log.Fatal(err)
	}
	if err := exec.Command("goimports", "-w=true", "../views/routes_context.go").Run(); err != nil {
		log.Fatal(err)
	}
}
