// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package expr

import (
	"fmt"
	"strings"
)

// ExprTypeClassicConditionsConditionsQueryConverter accepts a `ExprTypeClassicConditionsConditionsQuery` object and generates the Go code to build this object using builders.
func ExprTypeClassicConditionsConditionsQueryConverter(input ExprTypeClassicConditionsConditionsQuery) string {
	calls := []string{
		`expr.NewExprTypeClassicConditionsConditionsQueryBuilder()`,
	}
	var buffer strings.Builder
	if input.Params != nil && len(input.Params) >= 1 {

		buffer.WriteString(`Params(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Params {
			tmpparamsarg1 := fmt.Sprintf("%#v", arg1)
			tmparg0 = append(tmparg0, tmpparamsarg1)
		}
		arg0 := "[]string{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
