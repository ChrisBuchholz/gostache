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

## Escaping

gostache escapes all values when using the double-mustache syntax. Characters
that gets escaped are `" ' & < >`. To disable escaping, simply use
triple-mustaches like `{{{unescaped_variable}}}`.

## Templates

You can put your templates in files instead of keeping them as strings in your
program.

    result := gostache.RenderFile("mytemplate", p)

gostache will look for a file templates/mytemplate.mustache inside the
directory that gostache is executed from (CWD).

## Partials

Using partials with gostache is a no-fuzz deal.

    {{>top}}

    {{content}}

    {{>bottom}}

gostache will look for templates/partials/top.mustache and templates/partials/
bottom.mustache inside CWD and replace `{{>top}}` and `{{>bottom}}` with the
content of those files respectively.
