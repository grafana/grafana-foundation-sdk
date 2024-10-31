// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alerting

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// TimeIntervalItemConverter accepts a `TimeIntervalItem` object and generates the Go code to build this object using builders.
func TimeIntervalItemConverter(input TimeIntervalItem) string {
	calls := []string{
		`alerting.NewTimeIntervalItemBuilder()`,
	}
	var buffer strings.Builder
	if input.DaysOfMonth != nil && len(input.DaysOfMonth) >= 1 {

		buffer.WriteString(`DaysOfMonth(`)
		tmparg0 := []string{}
		for _, arg1 := range input.DaysOfMonth {
			tmpdays_of_montharg1 := fmt.Sprintf("%#v", arg1)
			tmparg0 = append(tmparg0, tmpdays_of_montharg1)
		}
		arg0 := "[]string{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Location != nil && *input.Location != "" {

		buffer.WriteString(`Location(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Location))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Months != nil && len(input.Months) >= 1 {

		buffer.WriteString(`Months(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Months {
			tmpmonthsarg1 := fmt.Sprintf("%#v", arg1)
			tmparg0 = append(tmparg0, tmpmonthsarg1)
		}
		arg0 := "[]string{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Times != nil && len(input.Times) >= 1 {

		buffer.WriteString(`Times(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Times {
			tmptimesarg1 := TimeIntervalTimeRangeConverter(arg1)
			tmparg0 = append(tmparg0, tmptimesarg1)
		}
		arg0 := "[]cog.Builder[alerting.TimeIntervalTimeRange]{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Weekdays != nil && len(input.Weekdays) >= 1 {

		buffer.WriteString(`Weekdays(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Weekdays {
			tmpweekdaysarg1 := fmt.Sprintf("%#v", arg1)
			tmparg0 = append(tmparg0, tmpweekdaysarg1)
		}
		arg0 := "[]string{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Years != nil && len(input.Years) >= 1 {

		buffer.WriteString(`Years(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Years {
			tmpyearsarg1 := fmt.Sprintf("%#v", arg1)
			tmparg0 = append(tmparg0, tmpyearsarg1)
		}
		arg0 := "[]string{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
