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

+ Response 200 (application/json)

    [Person][]

`

var dummyJSON = `
{
    "_version": "1.0",
    "name": "Hello API",
    "description": "A simple API demo\n\n",
    "metadata": {
        "FORMAT": {
            "value": "1A"
        },
        "HOST": {
            "value": "https://api.example.com/v1"
        }
    },
    "resourceGroups": [{
        "name": "People",
        "description": "This section describes about the People\n\n",
        "resources": [{
            "name": "Person",
            "description": "Represent particular Person\n\n",
            "uriTemplate": "/people/{id}",
            "model": {
                "name": "Person",
                "description": "",
                "headers": {
                    "Content-Type": {
                        "value": "application/json"
                    }
                },
                "body": "{\"name\":\"Gesang\",\"birthdate\":\"01-09-1917\"}\n",
                "schema": ""
            },
            "parameters": {
                "id": {
                    "description": "The id of the Person.",
                    "type": "string",
                    "required": false,
                    "default": "",
                    "example": "123",
                    "values": []
                }
            },
            "headers": {},
            "actions": [{
                "name": "Retrieve Person",
                "description": "Return the information for the Person\n\n",
                "method": "GET",
                "parameters": {},
                "headers": {},
                "examples": [{
                    "name": "",
                    "description": "",
                    "requests": [{
                        "name": "",
                        "description": "",
                        "headers": {
                            "Content-Type": {
                                "value": "application/json"
                            },
                            "Authorization": {
                                "value": "Basic AbcdeFg="
                            }
                        },
                        "body": "",
                        "schema": ""
                    }],
                    "responses": [{
                        "name": "200",
                        "description": "",
                        "headers": {
                            "Content-Type": {
                                "value": "application/json"
                            }
                        },
                        "body": "{\"name\":\"Gesang\",\"birthdate\":\"01-09-1917\"}\n",
                        "schema": ""
                    }]
                }]
            }]
        }]
    }]
}
`

var dummyAPI = &API{
	Version:     "1.0",
	Name:        "Hello API",
	Description: "A simple API demo\n\n",
	Metadata: Metadata{
		Format: Format{"1A"},
		Host:   Host{"https://api.example.com/v1"},
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
						Headers: map[string]Header{
							"Content-Type": Header{Value: "application/json"},
						},
						Body:   "{\"name\":\"Gesang\",\"birthdate\":\"01-09-1917\"}\n",
						Schema: "",
					},
					Parameters: map[string]Parameter{
						"id": Parameter{
							Description: "The id of the Person.",
							Type:        "string",
							Required:    false,
							Default:     "",
							Example:     "123",
							Values:      []string{},
						},
					},
					Headers: map[string]Header{},
					Actions: []Action{
						Action{
							Name:        "Retrieve Person",
							Description: "Return the information for the Person\n\n",
							Method:      "GET",
							Parameters:  map[string]Parameter{},
							Headers:     map[string]Header{},
							Examples: []Example{
								Example{
									Name:        "",
									Description: "",
									Requests: []Request{
										Request{
											Name:        "",
											Description: "",
											Headers: map[string]Header{
												"Content-Type":  Header{Value: "application/json"},
												"Authorization": Header{Value: "Basic AbcdeFg="},
											},
											Body:   "",
											Schema: "",
										},
									},
									Responses: []Response{
										Response{
											Name:        "200",
											Description: "",
											Headers: map[string]Header{
												"Content-Type": Header{Value: "application/json"},
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
