// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	"fmt"
	"strings"
)

// MovingAverageSimpleModelSettingsConverter accepts a `MovingAverageSimpleModelSettings` object and generates the Go code to build this object using builders.
func MovingAverageSimpleModelSettingsConverter(input MovingAverageSimpleModelSettings) string {
	calls := []string{
		`elasticsearch.NewMovingAverageSimpleModelSettingsBuilder()`,
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
