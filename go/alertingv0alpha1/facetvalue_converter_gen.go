// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alertingv0alpha1

import (
	"fmt"
	"strings"
)

// FacetValueConverter accepts a `FacetValue` object and generates the Go code to build this object using builders.
func FacetValueConverter(input FacetValue) string {
	calls := []string{
		`alertingv0alpha1.NewFacetValueBuilder()`,
	}
	var buffer strings.Builder
	if input.Value != "" {

		buffer.WriteString(`Value(`)
		arg0 := fmt.Sprintf("%#v", input.Value)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	{
		buffer.WriteString(`Count(`)
		arg0 := fmt.Sprintf("%#v", input.Count)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	return strings.Join(calls, ".\t\n")
}
