package main

import (
	"bytes"
	"strings"
	"text/template"
)

var routerTemplate = `
import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
)
{{$svrType := .ServiceType}}
{{$svrName := .ServiceName}}
{{$hasJwt := .HasJwt}}
func Register{{.ServiceType}}Router(e *echo.Echo) {
	{{- if $hasJwt}}
	jwtKey := "dangerous"
	if os.Getenv("JWTKEY") != "" {
		jwtKey = os.Getenv("JWTKEY")
	}
	config := middleware.JWTConfig{
		Claims:     &jwtCustomClaims{},
		SigningKey: []byte(jwtKey),
	}
	{{end}}
	{{- range .JwtRootPaths}}
	{{.RootPath}} := e.Group("/{{.RootPath}}")
	{{.RootPath}}.Use(middleware.JWTWithConfig(config))
	{{end}}
	{{- range .Methods}}
	{{- if .InScope}}
	{{.Scope}}.{{.Method}}("{{.Path}}", _{{$svrType}}_{{.Name}}{{.Num}}_HTTP_Handler)
	{{- else}}
	e.{{.Method}}("{{.Path}}", _{{$svrType}}_{{.Name}}{{.Num}}_HTTP_Handler)
	{{end}}
	{{- end}}
}

{{range .Methods}}
func _{{$svrType}}_{{.Name}}{{.Num}}_HTTP_Handler(c echo.Context) error {
	var req *{{.Request}} = new({{.Request}})

	{{- if .HasBody}}
	if err := c.Bind(req); err != nil {
		return err
	}
	{{- range .Fields}}
	if c.FormValue(strings.ToLower("{{.Name}}")) != "" {
		req.{{.Name}} = c.FormValue(strings.ToLower("{{.Name}}"))
	}
	{{- end}}
	{{- end}}

	{{- range .Fields}}
	if c.QueryParam(strings.ToLower("{{.Name}}")) != "" {
		req.{{.Name}} = c.QueryParam(strings.ToLower("{{.Name}}"))
	}
	if c.Param(strings.ToLower("{{.Name}}")) != "" {
		req.{{.Name}} = c.Param(strings.ToLower("{{.Name}}"))
	}
	{{- end}}

	{{- if .InScope}}
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	username := claims.Name
	fmt.Printf("Got jwt name is: %v\n", username)
	req.Username = username
	{{end}}

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
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"os"
	"time"
)
{{$svrType := .ServiceType}}
{{$svrName := .ServiceName}}
{{$hasJwt := .HasJwt}}
{{range .Methods}}
func {{$svrType}}{{.Name}}BusinessHandler(req *{{.Request}}) ({{.Reply}}, error) {
	{{- if .IsLogin}}
	// Throws unauthorized error
	if req.Username != "hello" || req.Password != "world" {
		return {{.Reply}}{ }, echo.ErrUnauthorized
	}

	// Set custom claims
	claims := &jwtCustomClaims{
		"Hello World",
		true,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	jk := "dangerous"
	if os.Getenv("JWTKEY") != "" {
		jk = os.Getenv("JWTKEY")
	}
	t, err := token.SignedString([]byte(jk))
	if err != nil {
		return {{.Reply}}{ }, err
	}
	{{end}}
	// Here can put your business logic,protoc-gen-ent soon coming
	//Below is example business logic code

	rj, err := json.Marshal(req)
	if err != nil {
		return {{.Reply}}{ }, err
	}
	fmt.Printf("Got {{.Request}} is: %v\n", string(rj))

	{{- if .IsLogin}}
	return {{.Reply}}{Token: "Bearer " + t}, nil
	{{- else}}
	return {{.Reply}}{ }, nil
	{{end}}
}
{{end}}
`
var authTypeTemplate = `

import "github.com/golang-jwt/jwt"

// jwtCustomClaims are custom claims extending default ones.
// See https://github.com/golang-jwt/jwt for more examples
type jwtCustomClaims struct {
	Name  string ` + "`json:\"name\"`" + `
	Admin bool   ` + "`json:\"admin\"`" + `
	jwt.StandardClaims
}

`

type serviceDesc struct {
	ServiceType  string // Greeter
	ServiceName  string // example.Greeter
	Metadata     string // example/v1/greeter.proto
	Methods      []*methodDesc
	MethodSets   map[string]*methodDesc
	LoginUrl     string
	HasJwt       bool
	JwtRootPaths []*JwtRootPath
}

type JwtRootPath struct {
	RootPath string
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
	DefaultHost  string
	InScope      bool
	Scope        string
	IsLogin      bool
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
