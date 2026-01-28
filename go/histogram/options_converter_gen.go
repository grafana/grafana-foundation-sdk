// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package histogram

import (
	"fmt"
	"strings"

	common "github.com/grafana/grafana-foundation-sdk/go/common"
)

// OptionsConverter accepts a `Options` object and generates the Go code to build this object using builders.
func OptionsConverter(input Options) string {
	calls := []string{
		`histogram.NewOptionsBuilder()`,
	}
	var buffer strings.Builder
	if input.BucketSize != nil {

		buffer.WriteString(`BucketSize(`)
		arg0 := fmt.Sprintf("%#v", *input.BucketSize)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.BucketOffset != nil && *input.BucketOffset != 0 {

		buffer.WriteString(`BucketOffset(`)
		arg0 := fmt.Sprintf("%#v", *input.BucketOffset)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	{
		buffer.WriteString(`Legend(`)
		arg0 := common.VizLegendOptionsConverter(input.Legend)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	{
		buffer.WriteString(`Tooltip(`)
		arg0 := common.VizTooltipOptionsConverter(input.Tooltip)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.Combine != nil {

		buffer.WriteString(`Combine(`)
		arg0 := fmt.Sprintf("%#v", *input.Combine)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
