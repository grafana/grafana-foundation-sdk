// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package testdata

import (
	"fmt"
	"strings"
)

// PulseWaveQueryConverter accepts a `PulseWaveQuery` object and generates the Go code to build this object using builders.
func PulseWaveQueryConverter(input PulseWaveQuery) string {
	calls := []string{
		`testdata.NewPulseWaveQueryBuilder()`,
	}
	var buffer strings.Builder
	if input.TimeStep != nil {

		buffer.WriteString(`TimeStep(`)
		arg0 := fmt.Sprintf("%#v", *input.TimeStep)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.OnCount != nil {

		buffer.WriteString(`OnCount(`)
		arg0 := fmt.Sprintf("%#v", *input.OnCount)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.OffCount != nil {

		buffer.WriteString(`OffCount(`)
		arg0 := fmt.Sprintf("%#v", *input.OffCount)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.OnValue != nil {

		buffer.WriteString(`OnValue(`)
		arg0 := fmt.Sprintf("%#v", *input.OnValue)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.OffValue != nil {

		buffer.WriteString(`OffValue(`)
		arg0 := fmt.Sprintf("%#v", *input.OffValue)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
