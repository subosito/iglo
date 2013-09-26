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
					<ul class="list-group">
						{{range .ResourceGroups}}
						<li class="list-group-item">
							<a href="#{{.Name | dasherize }}">{{.Name}}</a>
						</li>
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
									</div>

									<div class="panel-footer">
										{{range .Examples}}
											<ul class="list-group">
												{{range .Requests}}
													<li class="list-group-item">
														<h4 class="list-group-item-heading">Headers</h4>
														<table class="table">
															{{range $index, $element := .Headers}}
															<tbody>
																<tr>
																	<th>{{$index}}</th>
																	<td>{{.Value}}</td>
																</tr>
															</tbody>
															{{end}}
														</table>
													</li>
													<li class="list-group-item">
														<h4 class="list-group-item-heading">Body</h4>
														<pre class="prettyprint">{{.Body}}</pre>
													</li>
												{{end}}
											</ul>
										{{end}}
									</div>

								</div>
							{{end}}
						{{end}}
					{{end}}
				</div>
			</div>
		</div>
		<script src="https://google-code-prettify.googlecode.com/svn/loader/run_prettify.js"></script>
	</body>
</html>
`
