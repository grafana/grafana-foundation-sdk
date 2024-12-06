// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// BaseMovingAverageModelSettingsConverter accepts a `BaseMovingAverageModelSettings` object and generates the Go code to build this object using builders.
func BaseMovingAverageModelSettingsConverter(input BaseMovingAverageModelSettings) string {
	calls := []string{
		`elasticsearch.NewBaseMovingAverageModelSettingsBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`Model(`)
		arg0 := cog.Dump(input.Model)
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
