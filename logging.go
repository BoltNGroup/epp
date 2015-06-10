package epp

import (
	"fmt"
	"io"
)

var DebugLogger io.Writer

func logRequest(xml []byte) {
	if DebugLogger == nil {
		return
	}
	fmt.Printf("<!-- REQUEST -->\n%s\n\n", string(xml))
}

func logResponse(xml []byte) {
	if DebugLogger == nil {
		return
	}
	fmt.Printf("<!-- RESPONSE -->\n%s\n\n", string(xml))
}