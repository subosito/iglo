package iglo

import (
	"reflect"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	want := &API{
		Version:     "1.0",
		Name:        "Hello API",
		Description: "A simple API demo",
		Metadata: Metadata{
			Format: Format{"1A"},
			Host:   Host{"https://api.example.com/v1"},
		},
		ResourceGroups: []Group{
			Group{
				Name:        "People",
				Description: "This section describes about the People",
				Resources: []Resource{
					Resource{
						Name:        "Person",
						Description: "Represent particular Person",
						UriTemplate: "/people/{id}",
						Model: Model{
							Name:        "Person",
							Description: "",
							Headers: map[string]Header{
								"Content-Type": Header{Value: "application/json"},
							},
							Body:   "{\"name\":\"Gesang\"}",
							Schema: "",
						},
						Parameters: map[string]Parameter{
							"id": Parameter{
								Description: "The id of the Person.",
								Type:        "string",
								Required:    true,
								Default:     "",
								Example:     "123",
								Values:      []string{},
							},
						},
						Headers: map[string]Header{},
						Actions: []Action{
							Action{
								Name:        "Retrieve Person",
								Description: "Return the information for the Person",
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
												Body:   "{\"name\":\"Gesang\",\"birthdate\":\"01-09-1917\"}",
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

	a, err := Parse(strings.NewReader(dummyJSON))
	if err != nil {
		t.Errorf("Parse() returned an error %s", err)
	}

	if !reflect.DeepEqual(a, want) {
		t.Errorf("Unmarshal Metadata returned %+v, want %+v", a, want)
	}
}
