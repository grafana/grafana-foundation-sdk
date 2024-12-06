// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alerting

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// ContactPointConverter accepts a `ContactPoint` object and generates the Go code to build this object using builders.
func ContactPointConverter(input ContactPoint) string {
	calls := []string{
		`alerting.NewContactPointBuilder()`,
	}
	var buffer strings.Builder
	if input.DisableResolveMessage != nil {

		buffer.WriteString(`DisableResolveMessage(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.DisableResolveMessage))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Name != nil && *input.Name != "" {

		buffer.WriteString(`Name(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Name))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Provenance != nil && *input.Provenance != "" {

		buffer.WriteString(`Provenance(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Provenance))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	{
		buffer.WriteString(`Settings(`)
		arg0 := cog.Dump(input.Settings)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	{
		buffer.WriteString(`Type(`)
		arg0 := cog.Dump(input.Type)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.Uid != nil && *input.Uid != "" {

		buffer.WriteString(`Uid(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Uid))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
