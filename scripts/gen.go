package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	"text/template"
)

type Parameter struct {
	Name         string      `json:"name"`
	Description  string      `json:"description"`
	Type         string      `json:"type"`
	Required     bool        `json:"required"`
	DefaultValue interface{} `json:"defaultValue"`
}

type Method struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Path        string      `json:"path"`
	HttpMethod  string      `json:"httpMethod"`
	CnRegion    bool        `json:"cnRegion"`
	Parameters  []Parameter `json:"parameters"`
}

type Resource struct {
	Name    string   `json:"name"`
	Methods []Method `json:"methods"`
}

type Spec struct {
	Resources []Resource `json:"resources"`
}

type TemplateData struct {
	Spec        *Spec
	PackageName string
	SpecFile    string
}

//go:embed api.go.tmpl
var apiTemplate string

var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9 ]+`)
var pathVariableRegex = regexp.MustCompile(`\{([^}]+)\}`)
var htmlTagsRegex = regexp.MustCompile(`<(?:.|\n)*?>`)
var nonSpacedPeriodRegex = regexp.MustCompile(`\.([^ ])`)

func sanitize(str string) string {
	return nonAlphanumericRegex.ReplaceAllString(str, "")
}

func sprintUrl(str string) string {
	variables := make([]string, 0)
	matches := pathVariableRegex.FindAllStringSubmatch(str, -1)
	for _, match := range matches {
		variables = append(variables, match[1])
	}

	result := pathVariableRegex.ReplaceAllString(str, `%v`)
	return fmt.Sprintf("\"%s\", %s", result, strings.Join(variables, ", "))
}

func isPathParam(str string) bool {
	return pathVariableRegex.Match([]byte(str))
}

func removeHtml(str string) string {
	return nonSpacedPeriodRegex.ReplaceAllString(htmlTagsRegex.ReplaceAllString(str, ""), ". $1")
}

func trimRightN(str string, n int) string {
	return str[:len(str)-n]
}

func filterArgParams(params []Parameter) []Parameter {
	args := make([]Parameter, 0)
	for _, p := range params {
		if !(strings.Contains(p.Description, "(example search field)")) &&
			!(p.Name == "namespace" || p.Name == "locale") {
			args = append(args, p)
		}
	}
	return args
}

func main() {
	if len(os.Args) < 4 {
		log.Fatalln("Usage: ./generate <api-spec> <output-file> <package-name>")
	}

	apiSpecFile := os.Args[1]
	file, err := os.Open(apiSpecFile)
	defer file.Close()
	if err != nil {
		log.Fatalln("failed to open spec file", err)
	}

	contents, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalln("failed to read spec file", err)
	}

	spec := &Spec{}
	err = json.Unmarshal(contents, spec)
	if err != nil {
		log.Fatalln("failed to unmarshal spec file", err, string(contents))
	}

	funcs := template.FuncMap{
		"replace":         strings.ReplaceAll,
		"sanitize":        sanitize,
		"sprintUrl":       sprintUrl,
		"isPathParam":     isPathParam,
		"removeHtml":      removeHtml,
		"contains":        strings.Contains,
		"trimRightN":      trimRightN,
		"filterArgParams": filterArgParams,
	}

	tmpl, err := template.New("api").Funcs(funcs).Parse(apiTemplate)
	if err != nil {
		log.Fatalln("failed to parse template file", err)
	}

	file, err = os.OpenFile(os.Args[2], os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	defer file.Close()
	if err != nil {
		log.Fatalln("failed to open output file", err)
	}

	err = tmpl.Execute(file, TemplateData{Spec: spec, PackageName: os.Args[3], SpecFile: os.Args[1]})
	if err != nil {
		log.Fatalln("failed to execute template file", err)
	}
}
