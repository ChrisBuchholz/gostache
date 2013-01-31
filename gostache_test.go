package gostache

import "testing"

type Person struct {
	Name string
	Age  int
}

func TestRenderString(t *testing.T) {
	p := Person{"Triny", 7}
	expected := "Name: Triny\nAge: 7"
	result := RenderString("Name: {{Name}}\nAge: {{Age}}", p)
	if result != expected {
		t.Error("RenderString did not pass.")
	}
}
