// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package news

import (
	"fmt"
	"strings"
)

// OptionsConverter accepts a `Options` object and generates the Go code to build this object using builders.
func OptionsConverter(input Options) string {
	calls := []string{
		`news.NewOptionsBuilder()`,
	}
	var buffer strings.Builder
	if input.FeedUrl != nil && *input.FeedUrl != "" {

		buffer.WriteString(`FeedUrl(`)
		arg0 := fmt.Sprintf("%#v", *input.FeedUrl)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.ShowImage != nil && *input.ShowImage != true {

		buffer.WriteString(`ShowImage(`)
		arg0 := fmt.Sprintf("%#v", *input.ShowImage)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
