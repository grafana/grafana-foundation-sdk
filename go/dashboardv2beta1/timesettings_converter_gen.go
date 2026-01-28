// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// TimeSettingsConverter accepts a `TimeSettings` object and generates the Go code to build this object using builders.
func TimeSettingsConverter(input TimeSettingsSpec) string {
	calls := []string{
		`dashboardv2beta1.NewTimeSettingsBuilder()`,
	}
	var buffer strings.Builder
	if input.Timezone != nil && *input.Timezone != "" && *input.Timezone != "browser" {

		buffer.WriteString(`Timezone(`)
		arg0 := fmt.Sprintf("%#v", *input.Timezone)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.From != "" && input.From != "now-6h" {

		buffer.WriteString(`From(`)
		arg0 := fmt.Sprintf("%#v", input.From)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.To != "" && input.To != "now" {

		buffer.WriteString(`To(`)
		arg0 := fmt.Sprintf("%#v", input.To)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.AutoRefresh != "" {

		buffer.WriteString(`AutoRefresh(`)
		arg0 := fmt.Sprintf("%#v", input.AutoRefresh)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.AutoRefreshIntervals != nil && len(input.AutoRefreshIntervals) >= 1 {

		buffer.WriteString(`AutoRefreshIntervals(`)
		tmparg0 := []string{}
		for _, arg1 := range input.AutoRefreshIntervals {
			tmpautoRefreshIntervalsarg1 := fmt.Sprintf("%#v", arg1)
			tmparg0 = append(tmparg0, tmpautoRefreshIntervalsarg1)
		}
		arg0 := "[]string{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.QuickRanges != nil && len(input.QuickRanges) >= 1 {

		buffer.WriteString(`QuickRanges(`)
		tmparg0 := []string{}
		for _, arg1 := range input.QuickRanges {
			tmpquickRangesarg1 := TimeRangeOptionConverter(arg1)
			tmparg0 = append(tmparg0, tmpquickRangesarg1)
		}
		arg0 := "[]cog.Builder[dashboardv2beta1.TimeRangeOption]{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.HideTimepicker != false {

		buffer.WriteString(`HideTimepicker(`)
		arg0 := fmt.Sprintf("%#v", input.HideTimepicker)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.WeekStart != nil {

		buffer.WriteString(`WeekStart(`)
		arg0 := cog.Dump(*input.WeekStart)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.FiscalYearStartMonth != 0 {

		buffer.WriteString(`FiscalYearStartMonth(`)
		arg0 := fmt.Sprintf("%#v", input.FiscalYearStartMonth)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.NowDelay != nil && *input.NowDelay != "" {

		buffer.WriteString(`NowDelay(`)
		arg0 := fmt.Sprintf("%#v", *input.NowDelay)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
