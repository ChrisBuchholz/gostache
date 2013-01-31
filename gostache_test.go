package gostache

import (
	"strings"
	"testing"
)

type Person struct {
	Name string
	Age  int
}

func TestRenderString(t *testing.T) {
	p := Person{"Triny", 7}
	result := RenderString("Name: {{Name}}\nAge: {{Age}}", p)

	expected := "Name: Triny\nAge: 7"
	if result != expected {
		t.Error("RenderString did not pass.")
	}
}

func TestRenderFile(t *testing.T) {
	p := Person{"Triny", 7}
	result := RenderFile("a", p)

	expected := "Name: Triny\nAge: 7"
	if !strings.Contains(result, expected) {
		t.Error("RenderFile did not pass.")
	}
}

func TestRenderFileWithPartial(t *testing.T) {
	p := Person{"Triny", 7}
	result := RenderFile("b", p)

	expected := "Partial 7\n\nName: Triny\nAge: 7"
	if !strings.Contains(result, expected) {
		t.Error("RenderFile with a partial template did not pass.")
	}
}
