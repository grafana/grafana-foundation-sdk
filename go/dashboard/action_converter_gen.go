// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboard

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// ActionConverter accepts a `Action` object and generates the Go code to build this object using builders.
func ActionConverter(input Action) string {
	calls := []string{
		`dashboard.NewActionBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`Type(`)
		arg0 := cog.Dump(input.Type)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.Title != "" {

		buffer.WriteString(`Title(`)
		arg0 := fmt.Sprintf("%#v", input.Title)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Fetch != nil {

		buffer.WriteString(`Fetch(`)
		arg0 := FetchOptionsConverter(*input.Fetch)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Infinity != nil {

		buffer.WriteString(`Infinity(`)
		arg0 := InfinityOptionsConverter(*input.Infinity)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Confirmation != nil && *input.Confirmation != "" {

		buffer.WriteString(`Confirmation(`)
		arg0 := fmt.Sprintf("%#v", *input.Confirmation)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.OneClick != nil {

		buffer.WriteString(`OneClick(`)
		arg0 := fmt.Sprintf("%#v", *input.OneClick)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Variables != nil && len(input.Variables) >= 1 {

		buffer.WriteString(`Variables(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Variables {
			tmpvariablesarg1 := ActionVariableConverter(arg1)
			tmparg0 = append(tmparg0, tmpvariablesarg1)
		}
		arg0 := "[]cog.Builder[dashboard.ActionVariable]{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Style != nil {

		buffer.WriteString(`Style(`)
		arg0 := DashboardActionStyleConverter(*input.Style)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
