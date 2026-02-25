// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// SwitchVariableConverter accepts a `SwitchVariable` object and generates the Go code to build this object using builders.
func SwitchVariableConverter(input SwitchVariableKind) string {
	calls := []string{
		`dashboardv2beta1.NewSwitchVariableBuilder(` + fmt.Sprintf("%#v", input.Spec.Name) + `)`,
	}
	var buffer strings.Builder
	if input.Spec.Name != "" {

		buffer.WriteString(`Name(`)
		arg0 := fmt.Sprintf("%#v", input.Spec.Name)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.Current != "" && input.Spec.Current != "false" {

		buffer.WriteString(`Current(`)
		arg0 := fmt.Sprintf("%#v", input.Spec.Current)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.EnabledValue != "" && input.Spec.EnabledValue != "true" {

		buffer.WriteString(`EnabledValue(`)
		arg0 := fmt.Sprintf("%#v", input.Spec.EnabledValue)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.DisabledValue != "" && input.Spec.DisabledValue != "false" {

		buffer.WriteString(`DisabledValue(`)
		arg0 := fmt.Sprintf("%#v", input.Spec.DisabledValue)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.Label != nil && *input.Spec.Label != "" {

		buffer.WriteString(`Label(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.Label)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	{
		buffer.WriteString(`Hide(`)
		arg0 := cog.Dump(input.Spec.Hide)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.Spec.SkipUrlSync != false {

		buffer.WriteString(`SkipUrlSync(`)
		arg0 := fmt.Sprintf("%#v", input.Spec.SkipUrlSync)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.Description != nil && *input.Spec.Description != "" {

		buffer.WriteString(`Description(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.Description)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
