package gostache

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"regexp"
	"strings"
)

type Template struct {
	Template string
	Context  interface{}
}

func (t *Template) ParseSection(body string) (string, error) {

	return body, nil
}

// ParsePartial will find all occurences of the partial-mustache pattern
// in body and replace it with the content of the partial template file that
// matches the name used in the partial-mustache pattern if any such file exist
// parsePartial assumes that partials are placed inside the
// templates/partials/ directory in the current-working-directory
func (t *Template) ParsePartial(body string) (string, error) {
	r := regexp.MustCompile(`{{>(\w+)}}`)
	matches := r.FindAllStringSubmatch(body, -1)
	for _, match := range matches {
		if len(match) > 0 {
			cwd := os.Getenv("CWD")
			filename := match[1]
			filepath := cwd + "templates/partials/" + filename + ".mustache"

			f, err := os.Open(filepath)
			f.Close()
			if err != nil {
				return "", err
			}

			partial_template, err := ioutil.ReadFile(filepath)
			if err != nil {
				return "", err
			}

			body = strings.Replace(body, "{{>"+filename+"}}", string(partial_template), -1)
		}
	}
	return body, nil
}

// ParseString will find all occurences of the triple- and double-mustache
// pattern and replace it with the string value of the field in t.Context that
// matches and one exist
func (t *Template) ParseString(body string) (string, error) {
	triple := regexp.MustCompile(`{{{(\w+)}}}`)
	double := regexp.MustCompile(`{{(\w+)}}`)
	rgs := [2]*regexp.Regexp{triple, double}
	for _, r := range rgs {
		matches := r.FindAllStringSubmatch(body, -1)
		for _, match := range matches {
			if len(match) > 0 {
				pattern := match[0]
				fieldname := match[1]

				v := reflect.ValueOf(t.Context)
				value := v.FieldByName(fieldname)
				new_str := fmt.Sprintf("%v", value.Interface())

				var old_str string

				if len(pattern) == (len(fieldname) + 4) {
					old_str = "{{" + fieldname + "}}"
					new_str = HTMLEscape(new_str)
				} else {
					old_str = "{{{" + fieldname + "}}}"
				}

				body = strings.Replace(body, old_str, new_str, -1)
			}
		}
	}
	return body, nil
}

// Render will loop over the content of t.Template as long as it can find the
// mustache prefix `{{`
// when it finds one, it will determine if its a section, partial or a string
// -pattern and then tell either parseSection, parsePartial or parseString
// to parse it

// Render will consecutively call t.ParsePartial, t.ParseSection and
// t.ParseString on t.Template 
// The order of the calls are quite important - if you for example call
// t.ParseString before t.ParseSection, the mustache-sections will lose
// context
func (t *Template) Render() (string, error) {
	body := t.Template
	err := errors.New("")

	body, err = t.ParsePartial(body)
	if err != nil {
		return "", err
	}

	body, err = t.ParseSection(body)
	if err != nil {
		return "", err
	}

	body, err = t.ParseString(body)
	if err != nil {
		return "", err
	}

	return body, nil
}

// RenderString will create a Template structure of the template and context
// parameters and then ask Template.Render to render it. It then returns the
// return value from Template.Render.
func RenderString(template string, context interface{}) string {
	tmpl := Template{template, context}
	body, err := tmpl.Render()
	if err != nil {
		log.Fatal(err)
	}
	return body
}

// RenderFile will look for a mustache-file in the templates directory of the
// current-working-directory, and if it finds it, it reads its content and
// passes that along with the context to RenderString. It will then take
// the result from RenderString and return that.
func RenderFile(filename string, context interface{}) string {
	cwd := os.Getenv("CWD")
	filepath := cwd + "templates/" + filename + ".mustache"
	f, err := os.Open(filepath)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}

	template, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	body := RenderString(string(template), context)
	return body
}

// HTMLEscape replaces all applicable characters to HTML entities
func HTMLEscape(str string) string {
	chars := [5][2]string{
		{"\"", "&quot;"},
		{"'", "&apos;"},
		{"&", "&amp;"},
		{"<", "&lt;"},
		{">", "&gt;"},
	}

	for _, n := range chars {
		str = strings.Replace(str, n[0], n[1], -1)
	}

	return str
}
