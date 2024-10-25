package utils

import (
	"fmt"
	"html/template"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseTemplate(t *testing.T) {
	type args struct {
		data     interface{}
		fileName string
	}

	tests := []struct {
		args    args
		wantErr assert.ErrorAssertionFunc
		name    string
		want    string
	}{
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseTemplate(tt.args.fileName, tt.args.data)
			if !tt.wantErr(t, err, fmt.Sprintf("ParseTemplate(%v, %v)", tt.args.fileName, tt.args.data)) {
				return
			}

			assert.Equalf(t, tt.want, got, "ParseTemplate(%v, %v)", tt.args.fileName, tt.args.data)
		})
	}
}

func TestParseTemplateAndFuncs(t *testing.T) {
	valuesPass := struct {
		Name     string
		Password string
	}{Name: "firstName", Password: "Password"}

	valuesFail := struct {
		First string
		Last  string
	}{First: "firstName", Last: "LastName"}

	tests := []struct {
		name         string
		templateName string
		templatePath string
		values       interface{}
		want         string
		wantErr      bool
	}{
		{
			name:         "TestParseTemplatePass",
			templateName: "testtemplate.gohtml",
			templatePath: "./testdata",
			values:       valuesPass,
			wantErr:      false,
			want:         "Hi firstName. Your password is Password",
		},
		{
			name:         "TestParseTemplateFail",
			templateName: "testtemplate.gohtml",
			templatePath: "./testdata",
			values:       valuesFail,
			wantErr:      true,
			want:         "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseTemplateAndFuncs(tt.templateName, tt.templatePath, tt.values)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseTemplate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseTemplate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseTemplateWithFuncs(t *testing.T) {
	type args struct {
		templateName  string
		templatePath  string
		data          interface{}
		templateFuncs template.FuncMap
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseTemplateWithFuncs(tt.args.templateName, tt.args.templatePath, tt.args.data, tt.args.templateFuncs)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseTemplateWithFuncs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseTemplateWithFuncs() got = %v, want %v", got, tt.want)
			}
		})
	}
}
