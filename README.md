# iglo

API blueprint's formatter.

[![Build Status](https://travis-ci.org/subosito/iglo.svg?branch=master)](https://travis-ci.org/subosito/iglo)

## Writing API documentation

For writing API documentation, the iglo using [API Blueprint](http://apiblueprint.org/) syntax. You can read about its [specification](https://github.com/apiaryio/api-blueprint/blob/master/API%20Blueprint%20Specification.md).

Here's the example:

```markdown
FORMAT: 1A
HOST: https://api.example.com/v1

# Hello API

A simple API demo

# Group People

This section describes about the People

## Person [/people/{id}]

Represent particular Person

+ Parameters

    + id (required, string, `123`) ... The id of the Person.

+ Model (application/json)

    ```
    {"name":"Gesang","birthdate":"01-09-1917"}
    ```

### Retrieve Person [GET]

Return the information for the Person

+ Request (application/json)

    + Headers

        ```
        Authorization: Basic AbcdeFg=
        ```

+ Response 200 (application/json)

    [Person][]

```

## Demo

Make sure you have iglo installed in the `GOPATH`

```bash
$ cd $GOPATH
$ go get github.com/subosito/iglo
```

### Serving as HTTP

You can go to the `examples/api-server` directory and then run the `main.go`.

```bash
$ cd examples/api-server
$ go run main.go
```

Then visit `http://localhost:8080/` to see the output.

Or, you can just visit [this demo page](http://htmlpreview.github.io/?https://gist.github.com/subosito/6725894/raw/523f354769841728ede913e1a6d93bd593ef0a3e/iglo-preview.html) :)

### Exporting as HTML file

You can go to the `examples/api-exporter` directory and then run the `main.go`.

```bash
$ cd examples/api-exporter
$ go run main.go -out "api-output.html"
```

Now you have HTML generated output in the `api-output.html`.

## Dependencies

The iglo `ParseMarkdown` requires [drafter](https://github.com/apiaryio/drafter) to be installed. Refer to the drafter page for the installation details.

## Related Projects

- [github.com/peterhellberg/hiro](https://github.com/peterhellberg/hiro) : Allows you to generate HTML output as a file. Similar with the demo example but with more options.
