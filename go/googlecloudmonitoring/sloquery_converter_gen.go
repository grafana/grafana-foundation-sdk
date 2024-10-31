// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package googlecloudmonitoring

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// SLOQueryConverter accepts a `SLOQuery` object and generates the Go code to build this object using builders.
func SLOQueryConverter(input SLOQuery) string {
	calls := []string{
		`googlecloudmonitoring.NewSLOQueryBuilder()`,
	}
	var buffer strings.Builder
	if input.ProjectName != "" {

		buffer.WriteString(`ProjectName(`)
		arg0 := fmt.Sprintf("%#v", input.ProjectName)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.PerSeriesAligner != nil && *input.PerSeriesAligner != "" {

		buffer.WriteString(`PerSeriesAligner(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.PerSeriesAligner))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.AlignmentPeriod != nil && *input.AlignmentPeriod != "" {

		buffer.WriteString(`AlignmentPeriod(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.AlignmentPeriod))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.SelectorName != "" {

		buffer.WriteString(`SelectorName(`)
		arg0 := fmt.Sprintf("%#v", input.SelectorName)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.ServiceId != "" {

		buffer.WriteString(`ServiceId(`)
		arg0 := fmt.Sprintf("%#v", input.ServiceId)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.ServiceName != "" {

		buffer.WriteString(`ServiceName(`)
		arg0 := fmt.Sprintf("%#v", input.ServiceName)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.SloId != "" {

		buffer.WriteString(`SloId(`)
		arg0 := fmt.Sprintf("%#v", input.SloId)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.SloName != "" {

		buffer.WriteString(`SloName(`)
		arg0 := fmt.Sprintf("%#v", input.SloName)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Goal != nil {

		buffer.WriteString(`Goal(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Goal))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.LookbackPeriod != nil && *input.LookbackPeriod != "" {

		buffer.WriteString(`LookbackPeriod(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.LookbackPeriod))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
