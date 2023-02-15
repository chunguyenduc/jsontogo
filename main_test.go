package main

import "testing"

func Test_jsonToGo(t *testing.T) {
	type args struct {
		jsonBytes []byte
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "TC1 - basic title case",
			args: args{
				jsonBytes: []byte(`{"SourceCode": "exampleDataHere"}`),
			},
			want: "type AutoGenerated struct {\n\tSourceCode string `json:\"SourceCode\"`\n}",
		},
		{
			name: "TC2 - basic snake case",
			args: args{
				jsonBytes: []byte(`{"source_code": "exampleDataHere"}`),
			},
			want: "type AutoGenerated struct {\n\tSourceCode string `json:\"source_code\"`\n}",
		},
		{
			name: "TC3 - basic camel case",
			args: args{
				jsonBytes: []byte(`{"sourceCode": "exampleDataHere"}`),
			},
			want: "type AutoGenerated struct {\n\tSourceCode string `json:\"sourceCode\"`\n}",
		},
		{
			name: "TC4 - basic all capital case",
			args: args{
				jsonBytes: []byte(`{"SOURCE_CODE": ""}`),
			},
			want: "type AutoGenerated struct {\n\tSourceCode string `json:\"SOURCE_CODE\"`\n}",
		},
		{
			name: "TC5 - snake case with common initialisms",
			args: args{
				jsonBytes: []byte(`{"public_ip": ""}`),
			},
			want: "type AutoGenerated struct {\n\tPublicIP string `json:\"public_ip\"`\n}",
		},
		{
			name: "TC6 - title case with common initialisms",
			args: args{
				jsonBytes: []byte(`{"PublicIP": ""}`),
			},
			want: "type AutoGenerated struct {\n\tPublicIP string `json:\"PublicIP\"`\n}",
		},
		{
			name: "TC7 - camel case with common initialisms",
			args: args{
				jsonBytes: []byte(`{"publicIP": ""}`),
			},
			want: "type AutoGenerated struct {\n\tPublicIP string `json:\"publicIP\"`\n}",
		},
		{
			name: "TC8 - all cap case with common initialisms",
			args: args{
				jsonBytes: []byte(`{"PUBLIC_IP": ""}`),
			},
			want: "type AutoGenerated struct {\n\tPublicIP string `json:\"PUBLIC_IP\"`\n}",
		},
		{
			name: "TC9 - all cap case with common initialisms",
			args: args{
				jsonBytes: []byte(`{"PUBLIC_IP": ""}`),
			},
			want: "type AutoGenerated struct {\n\tPublicIP string `json:\"PUBLIC_IP\"`\n}",
		},
		{
			name: "TC10 - int field",
			args: args{
				jsonBytes: []byte(`{"age": 46}`),
			},
			want: "type AutoGenerated struct {\n\tAge int `json:\"age\"`\n}",
		},
		{
			name: "TC11 - negative float field",
			args: args{
				jsonBytes: []byte(`{"negativeFloat": -1.00}`),
			},
			want: "type AutoGenerated struct {\n\tNegativeFloat float64 `json:\"negativeFloat\"`\n}",
		},
		{
			name: "basic nested json",
			args: args{
				jsonBytes: []byte(`{"topLevel": { "secondLevel": "exampleDataHere"} }`),
			},
			want: "type AutoGenerated struct {\n\tTopLevel struct {\n\t\tSecondLevel string `json:\"secondLevel\"`\n\t} `json:\"topLevel\"`\n}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := jsonToGo(tt.args.jsonBytes)
			if (err != nil) != tt.wantErr {
				t.Errorf("jsonToGo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("jsonToGo() = %#v, want %#v", got, tt.want)
			}
		})
	}
}