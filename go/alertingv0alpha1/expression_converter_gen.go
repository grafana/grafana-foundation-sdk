// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alertingv0alpha1

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// ExpressionConverter accepts a `Expression` object and generates the Go code to build this object using builders.
func ExpressionConverter(input Expression) string {
	calls := []string{
		`alertingv0alpha1.NewExpressionBuilder()`,
	}
	var buffer strings.Builder
	if input.QueryType != nil && *input.QueryType != "" {

		buffer.WriteString(`QueryType(`)
		arg0 := fmt.Sprintf("%#v", *input.QueryType)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.RelativeTimeRange != nil {

		buffer.WriteString(`RelativeTimeRange(`)
		arg0 := RelativeTimeRangeConverter(*input.RelativeTimeRange)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.DatasourceUID != nil {

		buffer.WriteString(`DatasourceUID(`)
		arg0 := cog.Dump(*input.DatasourceUID)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Model != nil {

		buffer.WriteString(`Model(`)
		arg0 := cog.Dump(input.Model)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Source != nil {

		buffer.WriteString(`Source(`)
		arg0 := fmt.Sprintf("%#v", *input.Source)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
