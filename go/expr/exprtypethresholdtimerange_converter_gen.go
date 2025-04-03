// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package expr

import (
	"fmt"
	"strings"
)

// ExprTypeThresholdTimeRangeConverter accepts a `ExprTypeThresholdTimeRange` object and generates the Go code to build this object using builders.
func ExprTypeThresholdTimeRangeConverter(input ExprTypeThresholdTimeRange) string {
	calls := []string{
		`expr.NewExprTypeThresholdTimeRangeBuilder()`,
	}
	var buffer strings.Builder
	if input.From != "" && input.From != "now-6h" {

		buffer.WriteString(`From(`)
		arg0 := fmt.Sprintf("%#v", input.From)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.To != "" && input.To != "now" {

		buffer.WriteString(`To(`)
		arg0 := fmt.Sprintf("%#v", input.To)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
