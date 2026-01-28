// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package athena

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// DataqueryConverter accepts a `Dataquery` object and generates the Go code to build this object using builders.
func DataqueryConverter(input Dataquery) string {
	calls := []string{
		`athena.NewDataqueryBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`Format(`)
		arg0 := cog.Dump(input.Format)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	{
		buffer.WriteString(`ConnectionArgs(`)
		arg0 := ConnectionArgsConverter(input.ConnectionArgs)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.Table != nil && *input.Table != "" {

		buffer.WriteString(`Table(`)
		arg0 := fmt.Sprintf("%#v", *input.Table)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Column != nil && *input.Column != "" {

		buffer.WriteString(`Column(`)
		arg0 := fmt.Sprintf("%#v", *input.Column)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.QueryID != nil && *input.QueryID != "" {

		buffer.WriteString(`QueryID(`)
		arg0 := fmt.Sprintf("%#v", *input.QueryID)
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
	if input.QueryType != nil && *input.QueryType != "" {

		buffer.WriteString(`QueryType(`)
		arg0 := fmt.Sprintf("%#v", *input.QueryType)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.RawSQL != "" {

		buffer.WriteString(`RawSQL(`)
		arg0 := fmt.Sprintf("%#v", input.RawSQL)
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
