// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alerting

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// MatcherConverter accepts a `Matcher` object and generates the Go code to build this object using builders.
func MatcherConverter(input Matcher) string {
	calls := []string{
		`alerting.NewMatcherBuilder()`,
	}
	var buffer strings.Builder
	if input.Name != nil && *input.Name != "" {

		buffer.WriteString(`Name(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Name))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Type != nil {

		buffer.WriteString(`Type(`)
		arg0 := cog.Dump(*input.Type)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Value != nil && *input.Value != "" {

		buffer.WriteString(`Value(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Value))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
