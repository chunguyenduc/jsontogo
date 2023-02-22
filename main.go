package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/golang/glog"
)

const (
	SPACE   = " "
	TAB     = "\t"
	NEWLINE = "\n"
)

func jsonToGo(jsonBytes []byte) (string, error) {
	result := ""
	dummy := ""
	tabs := 0
	isValid := json.Valid(jsonBytes)
	if !isValid {
		return dummy, fmt.Errorf("json invalid")
	}

	var data interface{}
	err := json.Unmarshal(jsonBytes, &data)
	if err != nil {
		return dummy, err
	}

	structName := "AutoGenerated"
	result += "type " + structName

	glog.Infoln(data)

	return result + parseScope(data, tabs), nil
}

func parseScope(data interface{}, tabs int) string {
	glog.Errorf("parseSCope: %#v", data)
	result := ""
	scope := getType(data)
	if scope == "[]interface {}" {
		dataArray := data.([]interface{})
		sliceType := ""
		for _, d := range dataArray {
			_type := getType(d)
			if len(sliceType) == 0 {
				sliceType = _type
			} else if _type != sliceType {
				sliceType = "interface{}"
				break
			}
		}
		slice := SPACE + "[]"
		result += slice

		if sliceType == "map[string]interface {}" { // struct in array
			type ValueCount struct {
				Value interface{}
				Count int
			}
			mapAllFields := make(map[string]*ValueCount)
			length := len(dataArray)
			for _, d := range dataArray {
				for k, v := range d.(map[string]interface{}) {
					valueCount, ok := mapAllFields[k]
					glog.Infoln(valueCount, ok)
					if !ok {
						mapAllFields[k] = &ValueCount{
							Value: v,
							Count: 1,
						}
					} else {
						valueCount.Count++
					}
				}
			}

			omitEmpty := make(map[string]bool)
			structFields := make(map[string]interface{})

			// log debug
			for k, valueCount := range mapAllFields {
				glog.Infoln(k, valueCount)
			}

			for keyname, valueCount := range mapAllFields {
				structFields[keyname] = valueCount.Value
				omitEmpty[keyname] = valueCount.Count != length
			}
			result += parseStruct(structFields, omitEmpty, tabs)
		}

	} else if scope == "map[string]interface {}" {
		result += SPACE + parseStruct(data.(map[string]interface{}), nil, tabs)
		glog.Errorf("%#v", result)
	} else {
		result += SPACE + scope
	}

	return result
}

func parseStruct(structFields map[string]interface{}, omitEmpty map[string]bool, tabs int) string {
	glog.Infoln("parseStruct: ", structFields)
	result := "struct {" + NEWLINE
	glog.Infoln(omitEmpty)

	for key, value := range structFields {
		tabs++
		result += indent(tabs)
		result += makeFieldName(key)
		temp := parseScope(value, tabs)
		temp += " `json:\"" + key
		if isOmitEmpty, ok := omitEmpty[key]; ok && isOmitEmpty {
			temp += ",omitempty"
		}
		temp += "\"`"
		result += temp + NEWLINE
	}
	tabs--
	result += indent(tabs) + "}"
	glog.Errorf("%#v", result)
	return result
}

func indent(tabs int) string {
	result := ""
	for i := 0; i < tabs; i++ {
		result += TAB
	}
	return result
}

func getType(data interface{}) string {
	_type := fmt.Sprintf("%T", data)
	if _type == "float64" {
		value, errInt := strconv.ParseInt(fmt.Sprintf("%v", data.(float64)), 10, 64)
		if errInt == nil {
			if value < 2147483647 || value > -2147483648 {
				return "int"
			} else {
				return "int64"
			}
		}
	}
	return _type
}

func makeFieldName(field string) string {
	var commonInitialisms = map[string]bool{
		"ACL":   true,
		"API":   true,
		"ASCII": true,
		"CPU":   true,
		"CSS":   true,
		"DNS":   true,
		"EOF":   true,
		"GUID":  true,
		"HTML":  true,
		"HTTP":  true,
		"HTTPS": true,
		"ID":    true,
		"IP":    true,
		"JSON":  true,
		"LHS":   true,
		"QPS":   true,
		"RAM":   true,
		"RHS":   true,
		"RPC":   true,
		"SLA":   true,
		"SMTP":  true,
		"SQL":   true,
		"SSH":   true,
		"TCP":   true,
		"TLS":   true,
		"TTL":   true,
		"UDP":   true,
		"UI":    true,
		"UID":   true,
		"UUID":  true,
		"URI":   true,
		"URL":   true,
		"UTF8":  true,
		"VM":    true,
		"XML":   true,
		"XMPP":  true,
		"XSRF":  true,
		"XSS":   true,
	}
	result := ""
	fields := strings.Split(field, "_")
	for _, f := range fields {
		if ok := commonInitialisms[strings.ToUpper(f)]; ok {
			f = strings.ToUpper(f)
		} else if f == strings.ToUpper(f) {
			f = strings.ToLower(f)
		}
		result += strings.Title(f)
	}
	return result
}

func main() {
	flag.Parse()
	var jsonBytes = []byte(`
	{"people": [{ "name": "Frank"}, {"name": "Dennis"}, {"name": "Dee"}, {"name": "Charley"}, {"name":"Mac"}] }

	`)
	tempStruct, err := jsonToGo(jsonBytes)
	if err != nil {
		glog.Infoln(err)
	}
	glog.Infoln(tempStruct)
}
