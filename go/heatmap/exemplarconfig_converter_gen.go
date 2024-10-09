// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package heatmap

import (
	"fmt"
	"strings"
)

func ExemplarConfigConverter(input ExemplarConfig) string {
	calls := []string{
		`heatmap.NewExemplarConfigBuilder()`,
	}
	var buffer strings.Builder
	if input.Color != "" {

		buffer.WriteString(`Color(`)
		arg0 := fmt.Sprintf("%#v", input.Color)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
