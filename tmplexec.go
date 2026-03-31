package main

import (
	"bytes"
	"html/template"
)

func executeTemplate(r Recipient, t *template.Template) (string, error) {

	var tpl bytes.Buffer // Templates write to io.Writer, Buffer lets us capture output as a string

	err := t.Execute(&tpl, r)
	if err != nil {
		return "", err
	}

	return tpl.String(), nil
}
