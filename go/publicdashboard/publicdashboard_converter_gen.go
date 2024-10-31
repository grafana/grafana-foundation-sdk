// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package publicdashboard

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// PublicDashboardConverter accepts a `PublicDashboard` object and generates the Go code to build this object using builders.
func PublicDashboardConverter(input PublicDashboard) string {
	calls := []string{
		`publicdashboard.NewPublicDashboardBuilder()`,
	}
	var buffer strings.Builder
	if input.Uid != "" {

		buffer.WriteString(`Uid(`)
		arg0 := fmt.Sprintf("%#v", input.Uid)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.DashboardUid != "" {

		buffer.WriteString(`DashboardUid(`)
		arg0 := fmt.Sprintf("%#v", input.DashboardUid)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.AccessToken != nil && *input.AccessToken != "" {

		buffer.WriteString(`AccessToken(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.AccessToken))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	{
		buffer.WriteString(`IsEnabled(`)
		arg0 := fmt.Sprintf("%#v", input.IsEnabled)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	{
		buffer.WriteString(`AnnotationsEnabled(`)
		arg0 := fmt.Sprintf("%#v", input.AnnotationsEnabled)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	{
		buffer.WriteString(`TimeSelectionEnabled(`)
		arg0 := fmt.Sprintf("%#v", input.TimeSelectionEnabled)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	return strings.Join(calls, ".\t\n")
}
