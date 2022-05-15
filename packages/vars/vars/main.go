package main

import (
	"fmt"
	"os"
)

func Main(args map[string]interface{}) map[string]interface{} {
	return wrapHTML(fmt.Sprintf(
		"PARENTHESIS_MULTILINE: %q\nINTERNALLY_INDENTED_CURLY_BRACKET_MULTILINE %q\nQUOTED_CURLY_BRACKET %q\nQUOTED_DOUBLE_ESCAPED_CURLY_BRACKET: %q\n",
		os.Getenv("PARENTHESIS_MULTILINE"),
		os.Getenv("INTERNALLY_INDENTED_CURLY_BRACKET_MULTILINE"),
		os.Getenv("QUOTED_CURLY_BRACKET"),
		os.Getenv("QUOTED_DOUBLE_ESCAPED_CURLY_BRACKET"),
	))
}

func wrapHTML(body string) map[string]interface{} {
	return map[string]interface{}{
		"body": "<html><body><pre>" + string(body) + "</pre></body></html>",
	}
}
