// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// ExtendedStatConverter accepts a `ExtendedStat` object and generates the Go code to build this object using builders.
func ExtendedStatConverter(input ExtendedStat) string {
	calls := []string{
		`elasticsearch.NewExtendedStatBuilder()`,
	}
	var buffer strings.Builder
	if input.Label != "" {

		buffer.WriteString(`Label(`)
		arg0 := fmt.Sprintf("%#v", input.Label)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	{
		buffer.WriteString(`Value(`)
		arg0 := cog.Dump(input.Value)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	return strings.Join(calls, ".\t\n")
}
