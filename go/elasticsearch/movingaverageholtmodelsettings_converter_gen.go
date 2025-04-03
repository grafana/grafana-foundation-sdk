// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	"fmt"
	"strings"
)

// MovingAverageHoltModelSettingsConverter accepts a `MovingAverageHoltModelSettings` object and generates the Go code to build this object using builders.
func MovingAverageHoltModelSettingsConverter(input MovingAverageHoltModelSettings) string {
	calls := []string{
		`elasticsearch.NewMovingAverageHoltModelSettingsBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`Settings(`)
		arg0 := ElasticsearchMovingAverageHoltModelSettingsSettingsConverter(input.Settings)
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
