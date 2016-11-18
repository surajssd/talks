package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/xeipuuv/gojsonschema"
)

var schema = `
{
	"$schema": "http://json-schema.org/draft-04/schema#",
	"type": "object",
	"properties": {
		"name": {"$ref": "#/definitions/name"},
		"company": {"type": "string"},
		"reach_out": {"$ref": "#/definitions/reach_out"},
		"projects": {
			"type": "array",
			"items": {
				"type": "string",
				"enum": [
					"kompose",
					"openshift",
					"kubernetes"
				]
			},
			"uniqueItems": true,
			"minLength": 1
		}

	},
	"required": ["name", "company", "reach_out"],
	"definitions": {
		"name": {
			"id": "#/definitions/name",
			"type": "object",
			"properties": {
				"firstname": {"type": "string", "minLength": 1, "maxLength": 10},
				"lastname": {"type": "string", "minLength": 1, "maxLength": 10}
			}
		},
		"reach_out": {
			"id": "#/definitions/reach_out",
			"type": "object",
			"properties": {
				"twitter": {"type": "string"},
				"irc": {"type": "string"},
				"slack": {"type": "string"},
				"mail": {"type": "string"}
			}
		}
	},
	"additionalProperties": false
}
`

func main() {

	introContents, err := ioutil.ReadFile("intro_example.json")
	if err != nil {
		logrus.Fatalln(err)
	}

	doc := gojsonschema.NewStringLoader(string(introContents))
	schema := gojsonschema.NewStringLoader(schema)

	result, err := gojsonschema.Validate(schema, doc)
	if err != nil {
		logrus.Fatalln(err)
	}

	if !result.Valid() {
		var errs []string
		errs = append(errs, fmt.Sprintln("Input is invalid, following errors found:"))
		for _, desc := range result.Errors() {
			errs = append(errs, fmt.Sprintf("- %s", desc))
		}
		logrus.Fatalln(strings.Join(errs, "\n"))
	}
}
