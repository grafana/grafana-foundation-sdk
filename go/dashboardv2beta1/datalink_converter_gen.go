// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"fmt"
	"strings"
)

// DataLinkConverter accepts a `DataLink` object and generates the Go code to build this object using builders.
func DataLinkConverter(input DataLink) string {
	calls := []string{
		`dashboardv2beta1.NewDataLinkBuilder()`,
	}
	var buffer strings.Builder
	if input.Title != "" {

		buffer.WriteString(`Title(`)
		arg0 := fmt.Sprintf("%#v", input.Title)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Url != "" {

		buffer.WriteString(`Url(`)
		arg0 := fmt.Sprintf("%#v", input.Url)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.TargetBlank != nil {

		buffer.WriteString(`TargetBlank(`)
		arg0 := fmt.Sprintf("%#v", *input.TargetBlank)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
