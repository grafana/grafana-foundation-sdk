// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// AzureMetricQueryConverter accepts a `AzureMetricQuery` object and generates the Go code to build this object using builders.
func AzureMetricQueryConverter(input AzureMetricQuery) string {
	calls := []string{
		`azuremonitor.NewAzureMetricQueryBuilder()`,
	}
	var buffer strings.Builder
	if input.Resources != nil && len(input.Resources) >= 1 {

		buffer.WriteString(`Resources(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Resources {
			tmpresourcesarg1 := AzureMonitorResourceConverter(arg1)
			tmparg0 = append(tmparg0, tmpresourcesarg1)
		}
		arg0 := "[]cog.Builder[azuremonitor.AzureMonitorResource]{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.MetricNamespace != nil && *input.MetricNamespace != "" {

		buffer.WriteString(`MetricNamespace(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.MetricNamespace))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.CustomNamespace != nil && *input.CustomNamespace != "" {

		buffer.WriteString(`CustomNamespace(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.CustomNamespace))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.MetricName != nil && *input.MetricName != "" {

		buffer.WriteString(`MetricName(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.MetricName))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Region != nil && *input.Region != "" {

		buffer.WriteString(`Region(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Region))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.TimeGrain != nil && *input.TimeGrain != "" {

		buffer.WriteString(`TimeGrain(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.TimeGrain))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Aggregation != nil && *input.Aggregation != "" {

		buffer.WriteString(`Aggregation(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Aggregation))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.DimensionFilters != nil && len(input.DimensionFilters) >= 1 {

		buffer.WriteString(`DimensionFilters(`)
		tmparg0 := []string{}
		for _, arg1 := range input.DimensionFilters {
			tmpdimensionFiltersarg1 := AzureMetricDimensionConverter(arg1)
			tmparg0 = append(tmparg0, tmpdimensionFiltersarg1)
		}
		arg0 := "[]cog.Builder[azuremonitor.AzureMetricDimension]{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Top != nil && *input.Top != "" {

		buffer.WriteString(`Top(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Top))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.AllowedTimeGrainsMs != nil && len(input.AllowedTimeGrainsMs) >= 1 {

		buffer.WriteString(`AllowedTimeGrainsMs(`)
		tmparg0 := []string{}
		for _, arg1 := range input.AllowedTimeGrainsMs {
			tmpallowedTimeGrainsMsarg1 := fmt.Sprintf("%#v", arg1)
			tmparg0 = append(tmparg0, tmpallowedTimeGrainsMsarg1)
		}
		arg0 := "[]int64{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Alias != nil && *input.Alias != "" {

		buffer.WriteString(`Alias(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Alias))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.TimeGrainUnit != nil && *input.TimeGrainUnit != "" {

		buffer.WriteString(`TimeGrainUnit(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.TimeGrainUnit))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Dimension != nil && *input.Dimension != "" {

		buffer.WriteString(`Dimension(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Dimension))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.DimensionFilter != nil && *input.DimensionFilter != "" {

		buffer.WriteString(`DimensionFilter(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.DimensionFilter))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.MetricDefinition != nil && *input.MetricDefinition != "" {

		buffer.WriteString(`MetricDefinition(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.MetricDefinition))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.ResourceUri != nil && *input.ResourceUri != "" {

		buffer.WriteString(`ResourceUri(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.ResourceUri))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.ResourceGroup != nil && *input.ResourceGroup != "" {

		buffer.WriteString(`ResourceGroup(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.ResourceGroup))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.ResourceName != nil && *input.ResourceName != "" {

		buffer.WriteString(`ResourceName(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.ResourceName))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
