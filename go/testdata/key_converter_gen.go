// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package testdata

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// KeyConverter accepts a `Key` object and generates the Go code to build this object using builders.
func KeyConverter(input Key) string {
	calls := []string{
		`testdata.NewKeyBuilder()`,
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

	{
		buffer.WriteString(`Tick(`)
		arg0 := fmt.Sprintf("%#v", input.Tick)
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
