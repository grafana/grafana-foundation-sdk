// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package expr

import (
	"fmt"
	"strings"
)

func ExprTypeClassicConditionsConditionsReducerConverter(input ExprTypeClassicConditionsConditionsReducer) string {
	calls := []string{
		`expr.NewExprTypeClassicConditionsConditionsReducerBuilder()`,
	}
	var buffer strings.Builder
	if input.Type != "" {

		buffer.WriteString(`Type(`)
		arg0 := fmt.Sprintf("%#v", input.Type)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
