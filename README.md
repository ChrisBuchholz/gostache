gostache [![Build Status](https://travis-ci.org/ChrisBuchholz/gostache.png?branch=master)](https://travis-ci.org/ChrisBuchholz/gostache)
======

gostache is a Go implementation of [mustache](https://github.com/defunkt/mustache). It is heavily inspired by the [JavaScript](https://github.com/janl/mustache.js) and [Python](https://github.com/defunkt/pystache) implementations.

There is still a lot to do before gostache is ready for use, so be mindful.

## Usage

Quick example:

    package main

    import "github.com/ChrisBuchholz/gostache"

    type Person struct {
        Name string
        Age  int
    }

    func main() {
        p := Person{"Triny", 7}
        result := gostache.RenderString("Name: {{Name}}, Age: {{Age}}", p)
        println(result); // Name: Triny, Age: 7
    }

gostache simply looks for mustaches in the string (the first argument to
gostache.RenderString) and when it finds one, it will look for an exported
field in the structure (the second argument) with same name, and if it finds
one, it will use that as the value container of the mustache. 
