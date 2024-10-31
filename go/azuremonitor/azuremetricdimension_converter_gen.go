// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// AzureMetricDimensionConverter accepts a `AzureMetricDimension` object and generates the Go code to build this object using builders.
func AzureMetricDimensionConverter(input AzureMetricDimension) string {
	calls := []string{
		`azuremonitor.NewAzureMetricDimensionBuilder()`,
	}
	var buffer strings.Builder
	if input.Dimension != nil && *input.Dimension != "" {

		buffer.WriteString(`Dimension(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Dimension))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Operator != nil && *input.Operator != "" {

		buffer.WriteString(`Operator(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Operator))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Filters != nil && len(input.Filters) >= 1 {

		buffer.WriteString(`Filters(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Filters {
			tmpfiltersarg1 := fmt.Sprintf("%#v", arg1)
			tmparg0 = append(tmparg0, tmpfiltersarg1)
		}
		arg0 := "[]string{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Filter != nil && *input.Filter != "" {

		buffer.WriteString(`Filter(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Filter))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
