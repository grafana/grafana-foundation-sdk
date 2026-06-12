// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package prometheus

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// DataqueryConverter accepts a `Dataquery` object and generates the Go code to build this object using builders.
func DataqueryConverter(input Dataquery) string {
	calls := []string{
		`prometheus.NewDataqueryBuilder()`,
	}
	var buffer strings.Builder
	if input.AdhocFilters != nil && len(input.AdhocFilters) >= 1 {

		buffer.WriteString(`AdhocFilters(`)
		tmparg0 := []string{}
		for _, arg1 := range input.AdhocFilters {
			tmpadhocFiltersarg1 := AdhocFiltersConverter(arg1)
			tmparg0 = append(tmparg0, tmpadhocFiltersarg1)
		}
		arg0 := "[]cog.Builder[prometheus.AdhocFilters]{" + strings.Join(tmparg0, ",\n") + "}"
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
	if input.EditorMode != nil {

		buffer.WriteString(`EditorMode(`)
		arg0 := cog.Dump(*input.EditorMode)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Exemplar != nil {

		buffer.WriteString(`Exemplar(`)
		arg0 := fmt.Sprintf("%#v", *input.Exemplar)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Expr != "" {

		buffer.WriteString(`Expr(`)
		arg0 := fmt.Sprintf("%#v", input.Expr)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Format != nil {

		buffer.WriteString(`Format(`)
		arg0 := cog.Dump(*input.Format)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.GroupByKeys != nil && len(input.GroupByKeys) >= 1 {

		buffer.WriteString(`GroupByKeys(`)
		tmparg0 := []string{}
		for _, arg1 := range input.GroupByKeys {
			tmpgroupByKeysarg1 := fmt.Sprintf("%#v", arg1)
			tmparg0 = append(tmparg0, tmpgroupByKeysarg1)
		}
		arg0 := "[]string{" + strings.Join(tmparg0, ",\n") + "}"
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
	if input.Instant != nil && *input.Instant == true && input.Range != nil && *input.Range == false {

		buffer.WriteString(`Instant(`)
		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.IntervalFactor != nil {

		buffer.WriteString(`IntervalFactor(`)
		arg0 := fmt.Sprintf("%#v", *input.IntervalFactor)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.IntervalMs != nil {

		buffer.WriteString(`IntervalMs(`)
		arg0 := fmt.Sprintf("%#v", *input.IntervalMs)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.LegendFormat != nil && *input.LegendFormat != "" {

		buffer.WriteString(`LegendFormat(`)
		arg0 := fmt.Sprintf("%#v", *input.LegendFormat)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.MaxDataPoints != nil {

		buffer.WriteString(`MaxDataPoints(`)
		arg0 := fmt.Sprintf("%#v", *input.MaxDataPoints)
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
	if input.Range != nil && *input.Range == true && input.Instant != nil && *input.Instant == false {

		buffer.WriteString(`Range(`)
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
	if input.ResultAssertions != nil {

		buffer.WriteString(`ResultAssertions(`)
		arg0 := ResultAssertionsConverter(*input.ResultAssertions)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Scopes != nil && len(input.Scopes) >= 1 {

		buffer.WriteString(`Scopes(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Scopes {
			tmpscopesarg1 := ScopesConverter(arg1)
			tmparg0 = append(tmparg0, tmpscopesarg1)
		}
		arg0 := "[]cog.Builder[prometheus.Scopes]{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.TimeRange != nil {

		buffer.WriteString(`TimeRange(`)
		arg0 := TimeRangeConverter(*input.TimeRange)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Interval != nil && *input.Interval != "" {

		buffer.WriteString(`Interval(`)
		arg0 := fmt.Sprintf("%#v", *input.Interval)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
