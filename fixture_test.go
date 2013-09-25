package iglo

var dummyJSON = `
{
    "_version": "1.0",
    "name": "Hello API",
    "description": "A simple API demo",
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
        "description": "This section describes about the People",
        "resources": [{
            "name": "Person",
            "description": "Represent particular Person",
            "uriTemplate": "/people/{id}",
            "model": {
                "name": "Person",
                "description": "",
                "headers": {
                    "Content-Type": {
                        "value": "application/json"
                    }
                },
                "body": "{\"name\":\"Gesang\"}",
                "schema": ""
            },
            "parameters": {
                "id": {
                    "description": "The id of the Person.",
                    "type": "string",
                    "required": true,
                    "default": "",
                    "example": "123",
                    "values": []
                }
            },
            "headers": {},
            "actions": [{
                "name": "Retrieve Person",
                "description": "Return the information for the Person",
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
                        "body": "{\"name\":\"Gesang\",\"birthdate\":\"01-09-1917\"}",
                        "schema": ""
                    }]
                }]
            }]
        }]
    }]
}
`
