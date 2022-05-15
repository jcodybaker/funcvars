package main

import (
	"fmt"
	"os"
)

func Main(args map[string]interface{}) map[string]interface{} {
	return wrapHTML(fmt.Sprintf(
		"MULTILINE: %q\nHEX %q\n",
		os.Getenv("MULTILINE"),
		[]byte(os.Getenv("MULTILINE")),
	))
}

func wrapHTML(body string) map[string]interface{} {
	return map[string]interface{}{
		"body": "<html><body><pre>" + string(body) + "</pre></body></html>",
	}
}
