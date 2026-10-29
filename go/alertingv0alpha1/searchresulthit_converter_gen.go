// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alertingv0alpha1

import (
	"fmt"
	"strings"
)

// SearchResultHitConverter accepts a `SearchResultHit` object and generates the Go code to build this object using builders.
func SearchResultHitConverter(input SearchResultHit) string {
	calls := []string{
		`alertingv0alpha1.NewSearchResultHitBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`Resource(`)
		arg0 := SearchResultResourceConverter(input.Resource)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.Score != nil {

		buffer.WriteString(`Score(`)
		arg0 := fmt.Sprintf("%#v", *input.Score)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	{
		buffer.WriteString(`Fields(`)
		arg0 := RuleSearchHitFieldsConverter(input.Fields)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	return strings.Join(calls, ".\t\n")
}
