package iglo

import (
	"bitbucket.org/pkg/inflect"
	"html/template"
	"io"
	"strings"
)

func labelize(method string) string {
	switch method {
	case "GET":
		return "primary"
	case "POST":
		return "success"
	case "PUT":
		return "info"
	case "PATCH":
		return "warning"
	case "DELETE":
		return "danger"
	}

	return "default"
}

func HTML(w io.Writer, api *API) error {
	// template functions
	funcMap := template.FuncMap{
		"dasherize": inflect.Dasherize,
		"trim":      strings.Trim,
		"labelize":  labelize,
	}

	tmpl, err := template.New("html").Funcs(funcMap).Parse(Tmpl)
	if err != nil {
		return err
	}

	err = tmpl.Execute(w, api)
	if err != nil {
		return err
	}

	return nil
}

var Tmpl = `
{{define "Headers"}}
<table class="table">
	{{range $index, $element := .}}
		<tbody>
		<tr>
			<th>{{$index}}</th>
			<td>{{.Value}}</td>
		</tr>
		</tbody>
	{{end}}
</table>
{{end}}

{{define "Examples"}}
{{end}}

{{define "Parameters"}}
{{end}}

{{define "Resources"}}
	{{range .Resources}}
		{{$UriTemplate := .UriTemplate}}

		<div class="page-header">
			<h2 id="{{.Name | dasherize}}">{{.Name}}</h2>
		</div>

		<p>{{trim .Description "\n"}}</p>

		{{range .Actions}}
		<div class="panel panel-{{labelize .Method}}">
			<div class="panel-heading">
				{{.Method}}
				{{$UriTemplate}}
			</div>

			<div class="panel-body">
				{{.Description}}
			</div>
		</div>
		{{end}}
	{{end}}
{{end}}

{{define "ResourceGroups"}}
	{{range .}}
		<div class="page-header">
			<h1>{{.Name}}</h1>
		</div>

		{{template "Resources" .}}
	{{end}}
{{end}}

{{define "NavResourceGroups"}}
<ul class="nav">
	{{range .}}
	<li>
		<a href="#{{.Name | dasherize}}">{{.Name}}</a>
		<ul class="nav">
			{{range .Resources}}
				<li><a href="#{{.Name | dasherize }}">{{.Name}}</a></li>
			{{end}}
		</ul>
	</li>
	{{end}}
</ul>
{{end}}

<!DOCTYPE html>
	<head>
		<meta charset="utf-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>{{.Name}}</title>
		<meta name="description" content="{{.Description}}">
		<link rel="stylesheet" href="//netdna.bootstrapcdn.com/bootstrap/3.0.0/css/bootstrap.min.css">
	</head>
	<body>
		<div class="container">
			<div class="row">
				<div class="col-md-3">
					{{template "NavResourceGroups" .ResourceGroups}}
				</div>

				<div class="col-md-9">
					{{template "ResourceGroups" .ResourceGroups}}
				</div>
			</div>
		</div>
		<script src="https://google-code-prettify.googlecode.com/svn/loader/run_prettify.js"></script>
	</body>
</html>
`
