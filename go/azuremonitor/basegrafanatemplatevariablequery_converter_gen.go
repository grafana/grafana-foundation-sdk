// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	"fmt"
	"strings"
)

// BaseGrafanaTemplateVariableQueryConverter accepts a `BaseGrafanaTemplateVariableQuery` object and generates the Go code to build this object using builders.
func BaseGrafanaTemplateVariableQueryConverter(input BaseGrafanaTemplateVariableQuery) string {
	calls := []string{
		`azuremonitor.NewBaseGrafanaTemplateVariableQueryBuilder()`,
	}
	var buffer strings.Builder
	if input.RawQuery != nil && *input.RawQuery != "" {

		buffer.WriteString(`RawQuery(`)
		arg0 := fmt.Sprintf("%#v", *input.RawQuery)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
