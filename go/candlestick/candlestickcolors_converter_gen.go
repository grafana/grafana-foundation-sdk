// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package candlestick

import (
	"fmt"
	"strings"
)

// CandlestickColorsConverter accepts a `CandlestickColors` object and generates the Go code to build this object using builders.
func CandlestickColorsConverter(input CandlestickColors) string {
	calls := []string{
		`candlestick.NewCandlestickColorsBuilder()`,
	}
	var buffer strings.Builder
	if input.Up != "" && input.Up != "green" {

		buffer.WriteString(`Up(`)
		arg0 := fmt.Sprintf("%#v", input.Up)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Down != "" && input.Down != "red" {

		buffer.WriteString(`Down(`)
		arg0 := fmt.Sprintf("%#v", input.Down)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Flat != "" && input.Flat != "gray" {

		buffer.WriteString(`Flat(`)
		arg0 := fmt.Sprintf("%#v", input.Flat)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
