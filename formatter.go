package iglo

import (
	"bytes"
	"html/template"
	"io"
	"strings"

	"bitbucket.org/pkg/inflect"
	bf "github.com/russross/blackfriday"
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
	return HTMLCustom(Tmpl, w, api)
}

func HTMLCustom(s string, w io.Writer, api *API) error {
	funcMap := template.FuncMap{
		"dasherize":   inflect.Dasherize,
		"trim":        strings.Trim,
		"labelize":    labelize,
		"markdownize": markdownize,
	}

	tmpl, err := template.New("html").Funcs(funcMap).Parse(s)
	if err != nil {
		return err
	}

	err = tmpl.Execute(w, api.AST)
	if err != nil {
		return err
	}

	return nil
}

func MarkdownToHTML(w io.Writer, r io.Reader) error {
	return MarkdownToHTMLCustom(Tmpl, w, r)
}

func MarkdownToHTMLCustom(s string, w io.Writer, r io.Reader) error {
	data, err := ParseMarkdown(r)
	if err != nil {
		return err
	}

	err = JSONToHTMLCustom(s, w, bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	return nil
}

func JSONToHTML(w io.Writer, r io.Reader) error {
	return JSONToHTMLCustom(Tmpl, w, r)
}

func JSONToHTMLCustom(s string, w io.Writer, r io.Reader) error {
	api, err := ParseJSON(r)
	if err != nil {
		return err
	}

	return HTMLCustom(s, w, api)
}
