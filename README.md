# iglo

API blueprint's formatter.

[![Build Status](https://drone.io/github.com/subosito/iglo/status.png)](https://drone.io/github.com/subosito/iglo/latest)

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

You can go to the `examples` directory and then run the `api-server.go`.

```bash
$ cd examples
$ go run api-server.go
```

Then visit `http://localhost:8080/` to see the output.

Or, you can just visit [this demo page](http://htmlpreview.github.io/?https://gist.github.com/subosito/6725894/raw/523f354769841728ede913e1a6d93bd593ef0a3e/iglo-preview.html) :)

## Dependencies

The iglo `ParseMarkdown` requires [snowcrash](https://github.com/apiaryio/snowcrash) to be installed. Refer to the snowcrash page for the installation details.

