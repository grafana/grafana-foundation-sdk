// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package annotationslist

import (
	"fmt"
	"strings"
)

// OptionsConverter accepts a `Options` object and generates the Go code to build this object using builders.
func OptionsConverter(input Options) string {
	calls := []string{
		`annotationslist.NewOptionsBuilder()`,
	}
	var buffer strings.Builder
	if input.OnlyFromThisDashboard != false {

		buffer.WriteString(`OnlyFromThisDashboard(`)
		arg0 := fmt.Sprintf("%#v", input.OnlyFromThisDashboard)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.OnlyInTimeRange != false {

		buffer.WriteString(`OnlyInTimeRange(`)
		arg0 := fmt.Sprintf("%#v", input.OnlyInTimeRange)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Tags != nil && len(input.Tags) >= 1 {

		buffer.WriteString(`Tags(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Tags {
			tmptagsarg1 := fmt.Sprintf("%#v", arg1)
			tmparg0 = append(tmparg0, tmptagsarg1)
		}
		arg0 := "[]string{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Limit != 10 {

		buffer.WriteString(`Limit(`)
		arg0 := fmt.Sprintf("%#v", input.Limit)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.ShowUser != true {

		buffer.WriteString(`ShowUser(`)
		arg0 := fmt.Sprintf("%#v", input.ShowUser)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.ShowTime != true {

		buffer.WriteString(`ShowTime(`)
		arg0 := fmt.Sprintf("%#v", input.ShowTime)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.ShowTags != true {

		buffer.WriteString(`ShowTags(`)
		arg0 := fmt.Sprintf("%#v", input.ShowTags)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.NavigateToPanel != true {

		buffer.WriteString(`NavigateToPanel(`)
		arg0 := fmt.Sprintf("%#v", input.NavigateToPanel)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.NavigateBefore != "" && input.NavigateBefore != "10m" {

		buffer.WriteString(`NavigateBefore(`)
		arg0 := fmt.Sprintf("%#v", input.NavigateBefore)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.NavigateAfter != "" && input.NavigateAfter != "10m" {

		buffer.WriteString(`NavigateAfter(`)
		arg0 := fmt.Sprintf("%#v", input.NavigateAfter)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
