// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alerting

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// NotificationTemplateConverter accepts a `NotificationTemplate` object and generates the Go code to build this object using builders.
func NotificationTemplateConverter(input NotificationTemplate) string {
	calls := []string{
		`alerting.NewNotificationTemplateBuilder()`,
	}
	var buffer strings.Builder
	if input.Name != nil && *input.Name != "" {

		buffer.WriteString(`Name(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Name))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Provenance != nil {

		buffer.WriteString(`Provenance(`)
		arg0 := cog.Dump(*input.Provenance)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Template != nil && *input.Template != "" {

		buffer.WriteString(`Template(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Template))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Version != nil && *input.Version != "" {

		buffer.WriteString(`Version(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Version))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
