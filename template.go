package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"
)

// ParseTemplate parses a single template filename.
func ParseTemplate(fileName string, data interface{}) (string, error) {
	t, err := template.ParseFiles(fileName)
	if err != nil {
		return "", err
	}

	buffer := new(bytes.Buffer)
	if err := t.Execute(buffer, data); err != nil {
		return "", err
	}

	return buffer.String(), nil
}

// ParseTemplateAndFuncs based on html.
func ParseTemplateAndFuncs(templateName, templatePath string, data interface{}) (string, error) {
	templateFuncs := template.FuncMap{
		"formatTime": FormatTime,
		"formatDate": FormatDate,
		"formatBool": FormatBool,
	}

	return ParseTemplateWithFuncs(templateName, templatePath, data, templateFuncs)
}

// ParseTemplateWithFuncs parses a html template with the given template funcs.
func ParseTemplateWithFuncs(templateName, templatePath string, data interface{}, templateFuncs template.FuncMap) (string, error) {
	if !strings.HasSuffix(templateName, ".gohtml") {
		templateName += ".gohtml"
	}

	templateFile := fmt.Sprintf("%s/%s", templatePath, templateName)

	t, err := template.New(templateName).Funcs(templateFuncs).ParseFiles(templateFile)
	if err != nil {
		return "", err
	}
	buffer := new(bytes.Buffer)
	if err = t.Execute(buffer, data); err != nil {
		return "", err
	}
	return buffer.String(), nil
}
