// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package expr

import (
	"strings"
)

// ExprConverter accepts a `Expr` object and generates the Go code to build this object using builders.
func ExprConverter(input Expr) string {
	calls := []string{
		`expr.NewExprBuilder()`,
	}
	var buffer strings.Builder
	if input.TypeMath != nil {

		buffer.WriteString(`TypeMath(`)
		arg0 := TypeMathConverter(*input.TypeMath)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.TypeReduce != nil {

		buffer.WriteString(`TypeReduce(`)
		arg0 := TypeReduceConverter(*input.TypeReduce)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.TypeResample != nil {

		buffer.WriteString(`TypeResample(`)
		arg0 := TypeResampleConverter(*input.TypeResample)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.TypeClassicConditions != nil {

		buffer.WriteString(`TypeClassicConditions(`)
		arg0 := TypeClassicConditionsConverter(*input.TypeClassicConditions)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.TypeThreshold != nil {

		buffer.WriteString(`TypeThreshold(`)
		arg0 := TypeThresholdConverter(*input.TypeThreshold)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.TypeSql != nil {

		buffer.WriteString(`TypeSql(`)
		arg0 := TypeSqlConverter(*input.TypeSql)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
