// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package googlecloudmonitoring

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// FilterConverter accepts a `Filter` object and generates the Go code to build this object using builders.
func FilterConverter(input Filter) string {
	calls := []string{
		`googlecloudmonitoring.NewFilterBuilder()`,
	}
	var buffer strings.Builder
	if input.Key != "" {

		buffer.WriteString(`Key(`)
		arg0 := fmt.Sprintf("%#v", input.Key)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Operator != "" {

		buffer.WriteString(`Operator(`)
		arg0 := fmt.Sprintf("%#v", input.Operator)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Value != "" {

		buffer.WriteString(`Value(`)
		arg0 := fmt.Sprintf("%#v", input.Value)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Condition != nil && *input.Condition != "" {

		buffer.WriteString(`Condition(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Condition))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
