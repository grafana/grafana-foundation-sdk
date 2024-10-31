// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package candlestick

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// CandlestickFieldMapConverter accepts a `CandlestickFieldMap` object and generates the Go code to build this object using builders.
func CandlestickFieldMapConverter(input CandlestickFieldMap) string {
	calls := []string{
		`candlestick.NewCandlestickFieldMapBuilder()`,
	}
	var buffer strings.Builder
	if input.Open != nil && *input.Open != "" {

		buffer.WriteString(`Open(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Open))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.High != nil && *input.High != "" {

		buffer.WriteString(`High(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.High))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Low != nil && *input.Low != "" {

		buffer.WriteString(`Low(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Low))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Close != nil && *input.Close != "" {

		buffer.WriteString(`Close(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Close))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Volume != nil && *input.Volume != "" {

		buffer.WriteString(`Volume(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Volume))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
