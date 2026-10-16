// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// NotebookSpecConverter accepts a `NotebookSpec` object and generates the Go code to build this object using builders.
func NotebookSpecConverter(input NotebookSpec) string {
	calls := []string{
		`dashboardv2beta1.NewNotebookSpecBuilder()`,
	}
	var buffer strings.Builder
	if input.Title != "" {

		buffer.WriteString(`Title(`)
		arg0 := fmt.Sprintf("%#v", input.Title)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Description != nil && *input.Description != "" {

		buffer.WriteString(`Description(`)
		arg0 := fmt.Sprintf("%#v", *input.Description)
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

	{
		buffer.WriteString(`TimeSettings(`)
		arg0 := TimeSettingsConverter(input.TimeSettings)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.Elements != nil {

		buffer.WriteString(`Elements(`)
		arg0 := "map[string]dashboardv2beta1.NotebookElement{"
		for key, arg1 := range input.Elements {
			tmpelementsarg1 := cog.Dump(arg1)
			arg0 += "\t" + fmt.Sprintf("%#v", key) + ": " + tmpelementsarg1 + ","
		}
		arg0 += "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	{
		buffer.WriteString(`Layout(`)
		arg0 := NotebookLayoutKindConverter(input.Layout)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	return strings.Join(calls, ".\t\n")
}
