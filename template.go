package main

import (
	"bytes"
	"strings"
	"text/template"
)

var routerTemplate = `

import (
	"net/http"

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
	var payload *{{.Request}}
	payload = nil
	var pathParam map[string]string
	pathParam = nil

	{{- if .HasBody}}
	payload = new({{.Request}})
	if err := c.Bind(payload); err != nil {
		return err
	}
	{{- end}}

	{{- if .HasVars}}
	pathParam = make(map[string]string)
	{{- range .PathParams}}
	pathParam["{{.PathName}}"] = c.Param("{{.PathValue}}")
	{{- end}}
	reply, err := {{$svrType}}{{.Name}}BusinessHandler(&pathParam, payload)
	if err != nil {
		return err
	}
	{{- end}}

	//sm, _ := json.Marshal(payload)
	//fmt.Fprintf(os.Stderr, "payload = %v\n", string(sm))
	//sm1, _ := json.Marshal(pathParam)
	//fmt.Fprintf(os.Stderr, "pathParam = %v\n", string(sm1))

	return c.JSON(http.StatusOK, &reply)
}
{{end}}
`

var handlerTemplate = `

{{$svrType := .ServiceType}}
{{$svrName := .ServiceName}}

{{range .Methods}}
func {{$svrType}}{{.Name}}BusinessHandler(pathParam *map[string]string, payload *{{.Request}}) ({{.Reply}}, error) {
	// Here can put your business logic

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
	PathParams   []*PathParam
	HasBody      bool
	Body         string
	ResponseBody string
}

type PathParam struct {
	PathName  string
	PathValue string
}

func (s *serviceDesc) execute(tpl string) string {
	s.MethodSets = make(map[string]*methodDesc)
	for _, m := range s.Methods {
		s.MethodSets[m.Name] = m
	}
	//sm, _ := json.Marshal(s.MethodSets)
	//fmt.Fprintf(os.Stderr, "1====s.MethodSets = \n%v\n", string(sm))

	buf := new(bytes.Buffer)
	tmpl, err := template.New("http").Parse(strings.TrimSpace(tpl))
	if err != nil {
		panic(err)
	}
	if err := tmpl.Execute(buf, s); err != nil {
		panic(err)
	}
	//fmt.Fprintf(os.Stderr, "2======buf=\n%v", strings.Trim(buf.String(), "\r\n"))
	return strings.Trim(buf.String(), "\r\n")
}
