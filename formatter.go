package iglo

import (
	"bitbucket.org/pkg/inflect"
	"bytes"
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

func MarkdownToHTML(w io.Writer, r io.Reader) error {
	data, err := ParseMarkdown(r)
	if err != nil {
		return err
	}

	err = JSONToHTML(w, bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	return nil
}

func JSONToHTML(w io.Writer, r io.Reader) error {
	api, err := ParseJSON(r)
	if err != nil {
		return err
	}

	err = HTML(w, api)
	if err != nil {
		return err
	}

	return nil
}

var Tmpl = `
{{define "Responses"}}
	{{range .}}
		<li class="list-group-item bg-default response">
			<strong>Response <code>{{.Name}}</code></strong>
			<a href="javascript:;" class="pull-right btn btn-link btn-sm snippet-toggle">SHOW</a>
		</li>
		<li class="list-group-item snippet">
			{{if .Headers}}
				{{range $index := .Headers}}<code>&lt; {{.Name}}: {{.Value}}</code><br>{{end}}
			{{end}}
			{{if .Body}}
				<pre class="prettyprint">{{.Body}}</pre>
			{{end}}
		</li>
	{{end}}
{{end}}

{{define "Requests"}}
	{{range .}}
		<li class="list-group-item bg-default response">
			<strong>Requests</strong>
			<a href="javascript:;" class="pull-right btn btn-link btn-sm snippet-toggle">SHOW</a>
		</li>
		<li class="list-group-item snippet">
			{{if .Headers}}
				{{range $index := .Headers}}<code>&gt; {{.Name}}: {{.Value}}</code><br>{{end}}
			{{end}}
			{{if .Body}}
				<pre class="prettyprint">{{.Body}}</pre>
			{{end}}
		</li>
	{{end}}
{{end}}

{{define "Examples"}}
	{{range .}}
		{{template "Requests" .Requests}}
		{{template "Responses" .Responses}}
	{{end}}
{{end}}

{{define "Parameters"}}
<dl class="dl-horizontal">
	{{range $index := .}}
		<dt>{{.Name}}</dt>
		<dd>
		{{if .Required}}
			<strong>(required)</strong>
		{{end}}
		<code>{{.Type}}</code> {{.Description}}
		</dd>
	{{end}}
</dl>
{{end}}

{{define "Resources"}}
{{range .}}
	{{$UriTemplate := .UriTemplate}}
	{{$Parameters := .Parameters}}

	{{range .Actions}}
	<div class="panel panel-info">
		<div class="panel-heading">
			<span class="btn btn-{{.Method | labelize}}">{{.Method}}</span>
			<code>{{$UriTemplate}}</code>
		</div>

		<div class="panel-body">
			{{.Description | markdownize}}
		</div>

		<ul class="list-group">
			{{if $Parameters}}
				<li class="list-group-item bg-default"><strong>Parameters</strong></li>
				<li class="list-group-item">{{template "Parameters" $Parameters}}</li>
			{{end}}
			{{if .Examples}}{{template "Examples" .Examples}}{{end}}
		</ul>
	</div>
	{{end}}
{{end}}
{{end}}

{{define "ResourceGroups"}}
{{range .}}
	<div class="tab-pane" id="{{.Name | dasherize}}">
		<div class="panel panel-default">
			<div class="panel-heading">
				<h2 id="{{.Name | dasherize}}">{{.Name}}</h2>
			</div>
			<div class="panel-body">
				<p class="lead"><small>{{.Description | markdownize}}</small></p>
				{{template "Resources" .Resources}}
			</div>
		</div>
	</div>
{{end}}
{{end}}

{{define "NavResourceGroups"}}
<ul class="nav nav-pills nav-stacked nav-rg affix" id="group-tab">
	{{range .}}
		<li><a href="#{{.Name | dasherize}}" data-toggle="tab"><strong>{{.Name}}</strong></a></li>
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
		<link href='//fonts.googleapis.com/css?family=Open+Sans' rel='stylesheet' type='text/css'>
		<style>
			body {
				font-family: 'Open Sans', sans-serif;
			}

			tt, pre, code {
				font-family: Consolas, "Liberation Mono", Courier, monospace;
				background-color: transparent !important;
			}

			pre.prettyprint {
				border: 0px !important;
				background-color: #fff;
			}

			.bg-default {
				background-color: #F8F8F8;
			}

			.snippet {
				list-style: none;
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
					<div class="page-header">
						<h1>{{.Name}}</h1>
						<h2 class="lead"><small>{{.Description | markdownize}}</small></h2>
					</div>
				</div>

				<div class="col-md-3">
					{{template "NavResourceGroups" .ResourceGroups}}
				</div>

				<div class="col-md-9">
					<div class="tab-content">
						{{template "ResourceGroups" .ResourceGroups}}
					</div>
				</div>
			</div>
		</div>
		<script src="//cdnjs.cloudflare.com/ajax/libs/jquery/1.10.2/jquery.min.js"></script>
		<script src="//netdna.bootstrapcdn.com/bootstrap/3.0.0/js/bootstrap.min.js"></script>
		<script src="//google-code-prettify.googlecode.com/svn/loader/run_prettify.js"></script>
		<script>
			jQuery(function($) {
				$('#group-tab a:first').tab('show');
				$('.snippet-toggle').on("click", function(e) {
					e.preventDefault();

					var target = $(this).data('target');
					$(this).parent().next().toggle();
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
