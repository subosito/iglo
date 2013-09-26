package iglo

import (
	"bitbucket.org/pkg/inflect"
	bf "github.com/russross/blackfriday"
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

func markdownize(str string) template.HTML {
	b := bf.MarkdownCommon([]byte(str))
	return template.HTML(string(b))
}

func HTML(w io.Writer, api *API) error {
	// template functions
	funcMap := template.FuncMap{
		"dasherize":   inflect.Dasherize,
		"trim":        strings.Trim,
		"labelize":    labelize,
		"markdownize": markdownize,
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
<dl class="dl-horizontal">
	{{range $index, $element := .}}
		<dt>{{$index}}</dt>
		<dd>{{.Value}}</dd>
	{{end}}
</dl>
{{end}}

{{define "Responses"}}
	{{range .}}
		{{if .Body}}
			<li class="list-group-item bg-default response">
				<strong>Response</strong>
				<a href="javascript:;" class="pull-right btn btn-default response-toggle"><small>SHOW</small></a>
			</li>
			<li class="list-group-item response-snippet">
				<pre class="prettyprint">{{.Body}}</pre>
			</li>
		{{end}}
		{{if .Headers}}
			<li class="list-group-item bg-default"><strong>Headers</strong></li>
			<li class="list-group-item">
				{{template "Headers" .Headers}}
			</li>
		{{end}}
	{{end}}
{{end}}

{{define "Examples"}}
	{{range .}}
		{{template "Responses" .Responses}}
	{{end}}
{{end}}

{{define "Parameters"}}
<table class="table table-bordered">
	<thead>
		<tr>
			<th>Name</th>
			<th>Required</th>
			<th>Type</th>
			<th>Description</th>
		</tr>
	</thead>
	{{range $index, $element := .}}
		<tbody>
		<tr>
			<th>{{$index}}</th>
			<td>{{.Required}}</td>
			<td>{{.Type}}</td>
			<td>{{.Description}}</td>
		</tr>
		</tbody>
	{{end}}
</table>
{{end}}

{{define "Resources"}}
	{{range .Resources}}
		{{$UriTemplate := .UriTemplate}}
		{{$Parameters := .Parameters}}

		<div class="page-header">
			<h2 id="{{.Name | dasherize}}">{{.Name}}</h2>
			<p class="lead"><small>{{trim .Description "\n"}}</small></p>
		</div>

		{{range .Actions}}
		<div class="panel panel-{{labelize .Method}}">
			<div class="panel-heading">
				{{.Method}}
				{{$UriTemplate}}
			</div>

			<div class="panel-body">
				{{.Description}}
			</div>

			<ul class="list-group">
				{{if .Examples}}{{template "Examples" .Examples}}{{end}}
				{{if $Parameters}}
					<li class="list-group-item bg-default"><strong>Parameters</strong></li>
					<li class="list-group-item">{{template "Parameters" $Parameters}}</li>
				{{end}}
			</ul>
		</div>
		{{end}}
	{{end}}
{{end}}

{{define "ResourceGroups"}}
	{{range .}}
		{{template "Resources" .}}
	{{end}}
{{end}}

{{define "NavResourceGroups"}}
<div class="nav-rg list-group affix">
	{{range .}}
		<a href="#" class="list-group-item active"><strong>{{.Name}}</strong></a>
		{{range .Resources}}
			<a class="list-group-item" href="#{{.Name | dasherize }}">{{.Name}}</a>
		{{end}}
	{{end}}
</div>
{{end}}

<!DOCTYPE html>
	<head>
		<meta charset="utf-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>{{.Name}}</title>
		<meta name="description" content="{{.Description}}">
		<link rel="stylesheet" href="//netdna.bootstrapcdn.com/bootstrap/3.0.0/css/bootstrap.min.css">
		<link href='http://fonts.googleapis.com/css?family=Open+Sans' rel='stylesheet' type='text/css'>
		<style>
			body {
				font-family: 'Open Sans', sans-serif;
				margin-top: 20px;
			}

			tt, pre, code {
				font-family: Consolas, "Liberation Mono", Courier, monospace;
			}

			pre.prettyprint {
				border: 0px !important;
				background-color: #fff;
			}

			.bg-default {
				background-color: #F8F8F8;
			}

			.response-snippet {
				display: none;
			}

			.nav-rg {
				width: 262.5px;
			}
		</style>
	</head>
	<body>
		<div class="container">
			<div class="row">
				<div class="col-md-12">
					<nav class="navbar navbar-default">
						<div class="navbar-header">
							<a class="navbar-brand">{{.Name}}</a>
						</div>
					</nav>

					<h3 class="lead">{{markdownize .Description}}</h3>
				</div>

				<div class="col-md-3">
					{{template "NavResourceGroups" .ResourceGroups}}
				</div>

				<div class="col-md-9">
					{{template "ResourceGroups" .ResourceGroups}}
				</div>
			</div>
		</div>
		<script src="//cdnjs.cloudflare.com/ajax/libs/jquery/1.10.2/jquery.min.js"></script>
		<script src="https://google-code-prettify.googlecode.com/svn/loader/run_prettify.js"></script>
		<script>
			jQuery(function($) {
				$('.response-toggle').on("click", function(e) {
					e.preventDefault();

					$(this).parent().parent().find(".response-snippet").toggle();
					if ($(this).text() == "SHOW") {
						$(this).text("HIDE");
					} else {
						$(this).text("SHOW");
					}
				});
			});
		</script>
	</body>
</html>
`
