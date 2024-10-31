// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package cloudwatch

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// QueryEditorOperatorConverter accepts a `QueryEditorOperator` object and generates the Go code to build this object using builders.
func QueryEditorOperatorConverter(input QueryEditorOperator) string {
	calls := []string{
		`cloudwatch.NewQueryEditorOperatorBuilder()`,
	}
	var buffer strings.Builder
	if input.Name != nil && *input.Name != "" {

		buffer.WriteString(`Name(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Name))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Value != nil && input.Value.ArrayOfQueryEditorOperatorType != nil && len(input.Value.ArrayOfQueryEditorOperatorType) >= 1 {

		buffer.WriteString(`Value(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Value.ArrayOfQueryEditorOperatorType {
			tmpArrayOfQueryEditorOperatorTypearg1 := cog.Dump(arg1)
			tmparg0 = append(tmparg0, tmpArrayOfQueryEditorOperatorTypearg1)
		}
		arg0 := "[]cloudwatch.QueryEditorOperatorType{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
