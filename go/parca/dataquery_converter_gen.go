// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package parca

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// DataqueryConverter accepts a `Dataquery` object and generates the Go code to build this object using builders.
func DataqueryConverter(input Dataquery) string {
	calls := []string{
		`parca.NewDataqueryBuilder()`,
	}
	var buffer strings.Builder
	if input.LabelSelector != "" && input.LabelSelector != "{}" {

		buffer.WriteString(`LabelSelector(`)
		arg0 := fmt.Sprintf("%#v", input.LabelSelector)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.ProfileTypeId != "" {

		buffer.WriteString(`ProfileTypeId(`)
		arg0 := fmt.Sprintf("%#v", input.ProfileTypeId)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.RefId != "" {

		buffer.WriteString(`RefId(`)
		arg0 := fmt.Sprintf("%#v", input.RefId)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Hide != nil {

		buffer.WriteString(`Hide(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Hide))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.QueryType != nil && *input.QueryType != "" {

		buffer.WriteString(`QueryType(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.QueryType))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Datasource != nil {

		buffer.WriteString(`Datasource(`)
		arg0 := cog.Dump(*input.Datasource)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
