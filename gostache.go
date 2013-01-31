package gostache

import (
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strings"
)

type Template struct {
	template string
	dir      string
	context  interface{}
}

func (t *Template) parseBlock(body string) string {
	return body
}

func (t *Template) parsePartial(body string) string {
	return body
}

func (t *Template) parseString(body string) string {
	r := regexp.MustCompile(`{{(\w+)}}`)
	match := r.FindStringSubmatch(body)
	if len(match) > 0 {
		fieldname := match[1]
		v := reflect.ValueOf(t.context)
		value := v.FieldByName(fieldname)
		str_value := fmt.Sprintf("%v", value.Interface())
		body = strings.Replace(body, "{{"+fieldname+"}}", str_value, 1)
	}
	return body
}

func (t *Template) Render() string {
	body := t.template
	for {
		index := strings.Index(body, "{{")
		if index < 0 {
			break
		}
		switch {
		case t.template[index+2:index+3] == "#" || t.template[index+2:index+3] == "^":
			body = t.parseBlock(body)
		case t.template[index+2:index+3] == ">":
			body = t.parsePartial(body)
		default:
			body = t.parseString(body)
		}
	}
	return body
}

func RenderString(template string, context interface{}) string {
	cwd := os.Getenv("CWD")
	tmpl := Template{template, cwd, context}
	body := tmpl.Render()
	return body
}
