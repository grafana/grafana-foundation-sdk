// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package expr

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

func TypeThresholdConverter(input TypeThreshold) string {
	calls := []string{
		`expr.NewTypeThresholdBuilder()`,
	}
	var buffer strings.Builder
	if input.Conditions != nil && len(input.Conditions) >= 1 {

		buffer.WriteString(`Conditions(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Conditions {
			tmpconditionsarg1 := ExprTypeThresholdConditionsConverter(arg1)
			tmparg0 = append(tmparg0, tmpconditionsarg1)
		}
		arg0 := "[]cog.Builder[expr.ExprTypeThresholdConditions]{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Datasource != nil {

		buffer.WriteString(`Datasource(`)
		arg0 := cog.Dump(*input.Datasource)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Expression != "" {

		buffer.WriteString(`Expression(`)
		arg0 := fmt.Sprintf("%#v", input.Expression)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Hide != nil {

		buffer.WriteString(`Hide(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Hide))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.IntervalMs != nil {

		buffer.WriteString(`IntervalMs(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.IntervalMs))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.MaxDataPoints != nil {

		buffer.WriteString(`MaxDataPoints(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.MaxDataPoints))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.QueryType != nil && *input.QueryType != "" {

		buffer.WriteString(`QueryType(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.QueryType))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.RefId != "" {

		buffer.WriteString(`RefId(`)
		arg0 := fmt.Sprintf("%#v", input.RefId)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.ResultAssertions != nil {

		buffer.WriteString(`ResultAssertions(`)
		arg0 := ExprTypeThresholdResultAssertionsConverter(*input.ResultAssertions)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.TimeRange != nil {

		buffer.WriteString(`TimeRange(`)
		arg0 := ExprTypeThresholdTimeRangeConverter(*input.TimeRange)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
