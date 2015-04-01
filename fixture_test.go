package iglo

var dummyMarkdown = `
FORMAT: 1A
HOST: https://api.example.com/v1

# Hello API

A simple API demo

# Group People

This section describes about the People

## Person [/people/{id}]

Represent particular Person

+ Parameters

    + id (required, string, ` + "`123`" + `) ... The id of the Person.

+ Model (application/json)

    ` + "```" + `
    {"name":"Gesang","birthdate":"01-09-1917"}
    ` + "```" + `

### Retrieve Person [GET]

Return the information for the Person

+ Request (application/json)

    + Headers

        ` + "```" + `
        Authorization: Basic AbcdeFg=
        ` + "```" + `

+ Response 200

    [Person][]

`

var dummyAPI = &API{
	Version:     "3.0",
	Name:        "Hello API",
	Description: "A simple API demo\n\n",
	Metadata: []Metadata{
		Metadata{
			Name:  "FORMAT",
			Value: "1A",
		},
		Metadata{
			Name:  "HOST",
			Value: "https://api.example.com/v1",
		},
	},
	ResourceGroups: []ResourceGroup{
		ResourceGroup{
			Name:        "People",
			Description: "This section describes about the People\n\n",
			Resources: []Resource{
				Resource{
					Name:        "Person",
					Description: "Represent particular Person\n\n",
					UriTemplate: "/people/{id}",
					Model: Model{
						Name:        "Person",
						Description: "",
						Headers: []Header{
							Header{
								Name:  "Content-Type",
								Value: "application/json",
							},
						},
						Body:   "{\"name\":\"Gesang\",\"birthdate\":\"01-09-1917\"}\n",
						Schema: "",
					},
					Parameters: []Parameter{
						Parameter{
							Name:        "id",
							Description: "The id of the Person.",
							Type:        "string",
							Required:    true,
							Default:     "",
							Example:     "123",
							Values:      []Value{},
						},
					},
					Actions: []Action{
						Action{
							Name:        "Retrieve Person",
							Description: "Return the information for the Person\n\n",
							Method:      "GET",
							Parameters:  []Parameter{},
							Examples: []Example{
								Example{
									Name:        "",
									Description: "",
									Requests: []Request{
										Request{
											Name:        "",
											Description: "",
											Headers: []Header{
												Header{
													Name:  "Content-Type",
													Value: "application/json",
												},
												Header{
													Name:  "Authorization",
													Value: "Basic AbcdeFg=",
												},
											},
											Body:   "",
											Schema: "",
										},
									},
									Responses: []Response{
										Response{
											Name:        "200",
											Description: "",
											Headers: []Header{
												Header{
													Name:  "Content-Type",
													Value: "application/json",
												},
											},
											Body:   "{\"name\":\"Gesang\",\"birthdate\":\"01-09-1917\"}\n",
											Schema: "",
										},
									},
								},
							},
						},
					},
				},
			},
		},
	},
}
