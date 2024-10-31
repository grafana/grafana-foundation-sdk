// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package expr

import (
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// ExprTypeClassicConditionsConditionsOperatorConverter accepts a `ExprTypeClassicConditionsConditionsOperator` object and generates the Go code to build this object using builders.
func ExprTypeClassicConditionsConditionsOperatorConverter(input ExprTypeClassicConditionsConditionsOperator) string {
	calls := []string{
		`expr.NewExprTypeClassicConditionsConditionsOperatorBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`Type(`)
		arg0 := cog.Dump(input.Type)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	return strings.Join(calls, ".\t\n")
}
