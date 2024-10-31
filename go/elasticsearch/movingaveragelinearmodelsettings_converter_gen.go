// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	"fmt"
	"strings"
)

// MovingAverageLinearModelSettingsConverter accepts a `MovingAverageLinearModelSettings` object and generates the Go code to build this object using builders.
func MovingAverageLinearModelSettingsConverter(input MovingAverageLinearModelSettings) string {
	calls := []string{
		`elasticsearch.NewMovingAverageLinearModelSettingsBuilder()`,
	}
	var buffer strings.Builder
	if input.Window != "" {

		buffer.WriteString(`Window(`)
		arg0 := fmt.Sprintf("%#v", input.Window)
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
