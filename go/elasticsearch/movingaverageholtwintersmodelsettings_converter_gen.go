// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	"fmt"
	"strings"
)

// MovingAverageHoltWintersModelSettingsConverter accepts a `MovingAverageHoltWintersModelSettings` object and generates the Go code to build this object using builders.
func MovingAverageHoltWintersModelSettingsConverter(input MovingAverageHoltWintersModelSettings) string {
	calls := []string{
		`elasticsearch.NewMovingAverageHoltWintersModelSettingsBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`Settings(`)
		arg0 := ElasticsearchMovingAverageHoltWintersModelSettingsSettingsConverter(input.Settings)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.Window != "" {

		buffer.WriteString(`Window(`)
		arg0 := fmt.Sprintf("%#v", input.Window)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	{
		buffer.WriteString(`Minimize(`)
		arg0 := fmt.Sprintf("%#v", input.Minimize)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.Predict != "" {

		buffer.WriteString(`Predict(`)
		arg0 := fmt.Sprintf("%#v", input.Predict)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
