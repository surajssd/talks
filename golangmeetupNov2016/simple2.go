package main

import (
	"fmt"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/xeipuuv/gojsonschema"
)

func main() {

	docContents := `
		["kompose", "kubernetes", "openshift", "somefoo"]
	`
	schemaContents := `{
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
	}`

	doc := gojsonschema.NewStringLoader(docContents)
	schema := gojsonschema.NewStringLoader(schemaContents)

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
