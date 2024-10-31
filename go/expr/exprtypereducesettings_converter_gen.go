// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package expr

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// ExprTypeReduceSettingsConverter accepts a `ExprTypeReduceSettings` object and generates the Go code to build this object using builders.
func ExprTypeReduceSettingsConverter(input ExprTypeReduceSettings) string {
	calls := []string{
		`expr.NewExprTypeReduceSettingsBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`Mode(`)
		arg0 := cog.Dump(input.Mode)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.ReplaceWithValue != nil {

		buffer.WriteString(`ReplaceWithValue(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.ReplaceWithValue))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
