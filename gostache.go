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
	template string
	context  interface{}
}

func (t *Template) parseBlock(body string) (string, error) {
	return body, nil
}

func (t *Template) parsePartial(body string) (string, error) {
	r := regexp.MustCompile(`{{>(\w+)}}`)
	match := r.FindStringSubmatch(body)
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

		body = strings.Replace(body, "{{>"+filename+"}}", string(partial_template), 1)
	}
	return body, nil
}

func (t *Template) parseString(body string) (string, error) {
	r := regexp.MustCompile(`{{(\w+)}}`)
	match := r.FindStringSubmatch(body)
	if len(match) > 0 {
		fieldname := match[1]
		v := reflect.ValueOf(t.context)
		value := v.FieldByName(fieldname)
		str_value := fmt.Sprintf("%v", value.Interface())
		body = strings.Replace(body, "{{"+fieldname+"}}", str_value, 1)
	}
	return body, nil
}

func (t *Template) Render() (string, error) {
	body := t.template
	err := errors.New("")
	for {
		index := strings.Index(body, "{{")
		if index < 0 {
			break
		}
		switch {
		case body[index+2:index+3] == "#" || body[index+2:index+3] == "^":
			body, err = t.parseBlock(body)
		case body[index+2:index+3] == ">":
			body, err = t.parsePartial(body)
		default:
			body, err = t.parseString(body)
		}
	}

	if err != nil {
		return "", err
	}

	return body, nil
}

func RenderString(template string, context interface{}) string {
	tmpl := Template{template, context}
	body, err := tmpl.Render()
	if err != nil {
		log.Fatal(err)
	}
	return body
}

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
