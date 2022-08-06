package main

import (
	"bytes"
	"strings"
	"text/template"
)

var routerTemplate = `

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

{{$svrType := .ServiceType}}
{{$svrName := .ServiceName}}

func Register{{.ServiceType}}Router(e *echo.Echo) {
	{{- range .Methods}}
	e.{{.Method}}("{{.Path}}", _{{$svrType}}_{{.Name}}{{.Num}}_HTTP_Handler)
	{{- end}}
}

{{range .Methods}}
func _{{$svrType}}_{{.Name}}{{.Num}}_HTTP_Handler(c echo.Context) error {
	var req *{{.Request}} = new({{.Request}})

	{{- if .HasBody}}
	if err := c.Bind(req); err != nil {
		return err
	}
	{{- end}}

	{{- range .Fields}}
	if c.QueryParam(strings.ToLower("{{.Name}}")) != "" {
		req.{{.Name}} = c.QueryParam(strings.ToLower("{{.Name}}"))
	}
	if c.Param(strings.ToLower("{{.Name}}")) != "" {
		req.{{.Name}} = c.Param(strings.ToLower("{{.Name}}"))
	}
	{{- end}}

	reply, err := {{$svrType}}{{.Name}}BusinessHandler(req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &reply)
}
{{end}}
`

var handlerTemplate = `

import (
	"encoding/json"
	"fmt"
)
{{$svrType := .ServiceType}}
{{$svrName := .ServiceName}}

{{range .Methods}}
func {{$svrType}}{{.Name}}BusinessHandler(req *{{.Request}}) ({{.Reply}}, error) {
	// Here can put your business logic,protoc-gen-ent soon coming
	//Below is example business logic code

	reqJson, err := json.Marshal(req)
	if err != nil {
		return {{.Reply}}{}, err
	}
	fmt.Printf("Got {{.Request}} is: %v\n", string(reqJson))

	return {{.Reply}}{}, nil
}
{{end}}
`

type serviceDesc struct {
	ServiceType string // Greeter
	ServiceName string // example.Greeter
	Metadata    string // example/v1/greeter.proto
	Methods     []*methodDesc
	MethodSets  map[string]*methodDesc
}

type methodDesc struct {
	Name         string
	OriginalName string // The parsed original name
	Num          int
	Request      string
	Reply        string
	Path         string
	Method       string
	HasVars      bool
	HasBody      bool
	Body         string
	ResponseBody string
	Fields       []*RequestField
}

type RequestField struct {
	Name string
}

func (s *serviceDesc) execute(tpl string) string {
	s.MethodSets = make(map[string]*methodDesc)
	for _, m := range s.Methods {
		s.MethodSets[m.Name] = m
	}

	buf := new(bytes.Buffer)
	tmpl, err := template.New("http").Parse(strings.TrimSpace(tpl))
	if err != nil {
		panic(err)
	}
	if err := tmpl.Execute(buf, s); err != nil {
		panic(err)
	}
	return strings.Trim(buf.String(), "\r\n")
}
