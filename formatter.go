package iglo

import (
	"bitbucket.org/pkg/inflect"
	"html/template"
	"io"
	"strings"
)

func HTML(w io.Writer, api *API) error {
	// template functions
	funcMap := template.FuncMap{
		"dasherize": inflect.Dasherize,
		"trim":      strings.Trim,
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
				<div class="col-md-4">
					<ul class="nav nav-pills nav-stacked">
						{{range .ResourceGroups}}
						<li><a href="#{{.Name | dasherize }}">{{.Name}}</a></li>
						{{end}}
					</ul>
				</div>

				<div class="col-md-8">
					{{range .ResourceGroups}}
						<h1>{{.Name}}</h1>

						{{range .Resources}}
							<h3>{{.Name}}</h3>
							<p>{{trim .Description "\n"}}</p>
							{{$uri := .UriTemplate}}

							{{range .Actions}}
								<div class="panel panel-primary">
									<div class="panel-heading">
										{{.Method}}
										{{$uri}}
									</div>
									<div class="panel-body">
										{{.Description}}

										<h4>Parameters</h4>
										{{range $index, $element := .Parameters}}
											{{$index}}
											{{$element}}
											{{.Description}}
										{{end}}
									</div>
								</div>
							{{end}}
						{{end}}
					{{end}}
				</div>
			</div>
		</div>
	</body>
</html>
`
