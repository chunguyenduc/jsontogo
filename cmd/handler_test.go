package cmd

import "testing"

func Test_BuildBuildJSONToGo(t *testing.T) {
	type args struct {
		jsonBytes []byte
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "basic title case",
			args: args{
				jsonBytes: []byte(`{"SourceCode": "exampleDataHere"}`),
			},
			want: "type AutoGenerated struct {\n\tSourceCode string `json:\"SourceCode\"`\n}",
		},
		{
			name: "basic snake case",
			args: args{
				jsonBytes: []byte(`{"source_code": "exampleDataHere"}`),
			},
			want: "type AutoGenerated struct {\n\tSourceCode string `json:\"source_code\"`\n}",
		},
		{
			name: "basic camel case",
			args: args{
				jsonBytes: []byte(`{"sourceCode": "exampleDataHere"}`),
			},
			want: "type AutoGenerated struct {\n\tSourceCode string `json:\"sourceCode\"`\n}",
		},
		{
			name: "basic all capital case",
			args: args{
				jsonBytes: []byte(`{"SOURCE_CODE": ""}`),
			},
			want: "type AutoGenerated struct {\n\tSourceCode string `json:\"SOURCE_CODE\"`\n}",
		},
		{
			name: "snake case with common initialisms",
			args: args{
				jsonBytes: []byte(`{"public_ip": ""}`),
			},
			want: "type AutoGenerated struct {\n\tPublicIP string `json:\"public_ip\"`\n}",
		},
		{
			name: "title case with common initialisms",
			args: args{
				jsonBytes: []byte(`{"PublicIP": ""}`),
			},
			want: "type AutoGenerated struct {\n\tPublicIP string `json:\"PublicIP\"`\n}",
		},
		{
			name: "camel case with common initialisms",
			args: args{
				jsonBytes: []byte(`{"publicIP": ""}`),
			},
			want: "type AutoGenerated struct {\n\tPublicIP string `json:\"publicIP\"`\n}",
		},
		{
			name: "all capital case with common initialisms",
			args: args{
				jsonBytes: []byte(`{"PUBLIC_IP": ""}`),
			},
			want: "type AutoGenerated struct {\n\tPublicIP string `json:\"PUBLIC_IP\"`\n}",
		},
		{
			name: "all capital case with common initialisms",
			args: args{
				jsonBytes: []byte(`{"PUBLIC_IP": ""}`),
			},
			want: "type AutoGenerated struct {\n\tPublicIP string `json:\"PUBLIC_IP\"`\n}",
		},
		{
			name: "int field",
			args: args{
				jsonBytes: []byte(`{"age": 46}`),
			},
			want: "type AutoGenerated struct {\n\tAge int `json:\"age\"`\n}",
		},
		{
			name: "basic nested json",
			args: args{
				jsonBytes: []byte(`{"topLevel": { "secondLevel": "exampleDataHere"} }`),
			},
			want: "type AutoGenerated struct {\n\tTopLevel struct {\n\t\tSecondLevel string `json:\"secondLevel\"`\n\t} `json:\"topLevel\"`\n}",
		},
		{
			name: "nested array json",
			args: args{
				jsonBytes: []byte(`{"people": [{ "name": "Frank"}, {"name": "Dennis"}, {"name": "Dee"}, {"name": "Charley"}, {"name":"Mac"}] }`),
			},
			want: "type AutoGenerated struct {\n\tPeople []struct {\n\t\tName string `json:\"name\"`\n\t} `json:\"people\"`\n}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BuildJSONToGo(tt.args.jsonBytes, "")
			if (err != nil) != tt.wantErr {
				t.Errorf("BuildJSONToGo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("BuildJSONToGo() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
