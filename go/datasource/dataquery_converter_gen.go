// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package datasource

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// DataqueryConverter accepts a `Dataquery` object and generates the Go code to build this object using builders.
func DataqueryConverter(input Dataquery) string {
	calls := []string{
		`datasource.NewDataqueryBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`PanelId(`)
		arg0 := fmt.Sprintf("%#v", input.PanelId)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.RefId != nil && *input.RefId != "" {

		buffer.WriteString(`RefId(`)
		arg0 := fmt.Sprintf("%#v", *input.RefId)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Hide != nil {

		buffer.WriteString(`Hide(`)
		arg0 := fmt.Sprintf("%#v", *input.Hide)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	{
		buffer.WriteString(`WithTransforms(`)
		arg0 := fmt.Sprintf("%#v", input.WithTransforms)
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
