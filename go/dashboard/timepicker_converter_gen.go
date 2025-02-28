// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboard

import (
	"fmt"
	"strings"
)

// TimePickerConverter accepts a `TimePicker` object and generates the Go code to build this object using builders.
func TimePickerConverter(input TimePickerConfig) string {
	calls := []string{
		`dashboard.NewTimePickerBuilder()`,
	}
	var buffer strings.Builder
	if input.Hidden != nil && *input.Hidden != false {

		buffer.WriteString(`Hidden(`)
		arg0 := fmt.Sprintf("%#v", *input.Hidden)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.RefreshIntervals != nil && len(input.RefreshIntervals) >= 1 {

		buffer.WriteString(`RefreshIntervals(`)
		tmparg0 := []string{}
		for _, arg1 := range input.RefreshIntervals {
			tmprefresh_intervalsarg1 := fmt.Sprintf("%#v", arg1)
			tmparg0 = append(tmparg0, tmprefresh_intervalsarg1)
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
			tmpquick_rangesarg1 := TimeOptionConverter(arg1)
			tmparg0 = append(tmparg0, tmpquick_rangesarg1)
		}
		arg0 := "[]cog.Builder[dashboard.TimeOption]{" + strings.Join(tmparg0, ",\n") + "}"
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
