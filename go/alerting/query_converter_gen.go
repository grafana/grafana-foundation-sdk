// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alerting

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// QueryConverter accepts a `Query` object and generates the Go code to build this object using builders.
func QueryConverter(input Query) string {
	calls := []string{
		`alerting.NewQueryBuilder(` + fmt.Sprintf("%#v", cog.Unptr(input.RefId)) + `)`,
	}
	var buffer strings.Builder
	if input.DatasourceUid != nil && *input.DatasourceUid != "" {

		buffer.WriteString(`DatasourceUid(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.DatasourceUid))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Model != nil {

		buffer.WriteString(`Model(`)
		arg0 := cog.ConvertDataqueryToCode(input.Model)
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
	if input.RefId != nil && *input.RefId != "" {

		buffer.WriteString(`RefId(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.RefId))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.RelativeTimeRange != nil && input.RelativeTimeRange.From != nil && input.RelativeTimeRange.To != nil {

		buffer.WriteString(`RelativeTimeRange(`)
		arg0 := cog.Dump(*input.RelativeTimeRange.From)
		buffer.WriteString(arg0)
		buffer.WriteString(", ")
		arg1 := cog.Dump(*input.RelativeTimeRange.To)
		buffer.WriteString(arg1)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
