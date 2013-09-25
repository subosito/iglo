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

        {"name":"Gesang"}

### Retrieve Person [GET]

Return the information for the Person

+ Request (application/json)

    + Headers

            Authorization: Basic AbcdeFg=

+ Response 200 (application/json)

        {"name":"Gesang","birthdate":"01-09-1917"}

