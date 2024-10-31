// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package expr

import (
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// ExprTypeThresholdConditionsConverter accepts a `ExprTypeThresholdConditions` object and generates the Go code to build this object using builders.
func ExprTypeThresholdConditionsConverter(input ExprTypeThresholdConditions) string {
	calls := []string{
		`expr.NewExprTypeThresholdConditionsBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`Evaluator(`)
		arg0 := ExprTypeThresholdConditionsEvaluatorConverter(input.Evaluator)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.LoadedDimensions != nil {

		buffer.WriteString(`LoadedDimensions(`)
		arg0 := cog.Dump(input.LoadedDimensions)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.UnloadEvaluator != nil {

		buffer.WriteString(`UnloadEvaluator(`)
		arg0 := ExprTypeThresholdConditionsUnloadEvaluatorConverter(*input.UnloadEvaluator)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
