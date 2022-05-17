package utils

import (
	"html/template"
	"testing"
	"time"
)

func TestAddThousandsSeparator(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TestAddThousandsSeparator1",
			args: args{n: 6},
			want: "6",
		},
		{
			name: "TestAddThousandsSeparator100",
			args: args{n: 600},
			want: "600",
		},
		{
			name: "TestAddThousandsSeparator1000",
			args: args{n: 6000},
			want: "6,000",
		},
		{
			name: "TestAddThousandsSeparator10000",
			args: args{n: 60000},
			want: "60,000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddThousandsSeparator(tt.args.n); got != tt.want {
				t.Errorf("AddThousandsSeparator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatBool(t *testing.T) {
	type args struct {
		b bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TestFormatBoolYes",
			args: args{b: true},
			want: "Yes",
		},
		{
			name: "TestFormatBoolNo",
			args: args{b: false},
			want: "No",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatBool(tt.args.b); got != tt.want {
				t.Errorf("FormatBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatDate(t *testing.T) {
	type args struct {
		t time.Time
	}

	// testDate := time.Date(1999, 2, 28, 11, 11, 11, 0, time.Local)
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TestFormatDate",
			args: args{t: time.Date(1999, 2, 28, 11, 11, 11, 0, time.Local)},
			want: "28/02/1999",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatDate(tt.args.t); got != tt.want {
				t.Errorf("FormatDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatDateTimeInDefaultZone(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TestFormatDateTimeInDefaultZone",
			args: args{t: time.Date(1999, 2, 28, 11, 11, 11, 0, time.Local)},
			want: "11:11 AM, 28/02/1999 AEDT",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatDateTimeInDefaultZone(tt.args.t); got != tt.want {
				t.Errorf("FormatDateTimeInDefaultZone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatTime(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name     string
		testTime time.Time
		want     string
	}{
		{
			name:     "TestFormatTime",
			testTime: time.Date(1999, 2, 28, 11, 11, 11, 0, time.Local),
			want:     "11:11 AM",
		},
		{
			name:     "TestFormatTimeZero",
			testTime: time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
			want:     "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatTime(tt.testTime); got != tt.want {
				t.Errorf("FormatTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseTemplate(t *testing.T) {
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
			got, err := ParseTemplate(tt.templateName, tt.templatePath, tt.values)
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
