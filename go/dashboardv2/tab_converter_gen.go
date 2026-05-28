// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// TabConverter accepts a `Tab` object and generates the Go code to build this object using builders.
func TabConverter(input TabsLayoutTabKind) string {
	calls := []string{
		`dashboardv2.NewTabBuilder()`,
	}
	var buffer strings.Builder
	if input.Spec.Title != nil && *input.Spec.Title != "" {

		buffer.WriteString(`Title(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.Title)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.Layout.GridLayoutKind != nil {
		if input.Spec.Layout.GridLayoutKind != nil {
			buffer.WriteString(`GridLayout(`)
			arg0 := GridConverter(*input.Spec.Layout.GridLayoutKind)
			buffer.WriteString(arg0)

			buffer.WriteString(")")

			calls = append(calls, buffer.String())
			buffer.Reset()
		}
	}
	if input.Spec.Layout.RowsLayoutKind != nil {
		if input.Spec.Layout.RowsLayoutKind != nil {
			buffer.WriteString(`RowsLayout(`)
			arg0 := RowsConverter(*input.Spec.Layout.RowsLayoutKind)
			buffer.WriteString(arg0)

			buffer.WriteString(")")

			calls = append(calls, buffer.String())
			buffer.Reset()
		}
	}
	if input.Spec.Layout.AutoGridLayoutKind != nil {
		if input.Spec.Layout.AutoGridLayoutKind != nil {
			buffer.WriteString(`AutoGridLayout(`)
			arg0 := AutoGridConverter(*input.Spec.Layout.AutoGridLayoutKind)
			buffer.WriteString(arg0)

			buffer.WriteString(")")

			calls = append(calls, buffer.String())
			buffer.Reset()
		}
	}
	if input.Spec.Layout.TabsLayoutKind != nil {
		if input.Spec.Layout.TabsLayoutKind != nil {
			buffer.WriteString(`TabsLayout(`)
			arg0 := TabsConverter(*input.Spec.Layout.TabsLayoutKind)
			buffer.WriteString(arg0)

			buffer.WriteString(")")

			calls = append(calls, buffer.String())
			buffer.Reset()
		}
	}
	if input.Spec.ConditionalRendering != nil {

		buffer.WriteString(`ConditionalRendering(`)
		arg0 := ConditionalRenderingGroupConverter(*input.Spec.ConditionalRendering)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.Repeat != nil {

		buffer.WriteString(`Repeat(`)
		arg0 := TabRepeatOptionsConverter(*input.Spec.Repeat)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.Variables != nil && len(input.Spec.Variables) >= 1 {

		buffer.WriteString(`Variables(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Spec.Variables {
			tmpvariablesarg1 := cog.Dump(arg1)
			tmparg0 = append(tmparg0, tmpvariablesarg1)
		}
		arg0 := "[]dashboardv2.VariableKind{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
