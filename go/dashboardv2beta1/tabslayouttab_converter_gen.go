// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"fmt"
	"strings"
)

// TabsLayoutTabConverter accepts a `TabsLayoutTab` object and generates the Go code to build this object using builders.
func TabsLayoutTabConverter(input TabsLayoutTabKind) string {
	calls := []string{
		`dashboardv2beta1.NewTabsLayoutTabBuilder(` + fmt.Sprintf("%#v", *input.Spec.Title) + `)`,
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
			arg0 := GridLayoutConverter(*input.Spec.Layout.GridLayoutKind)
			buffer.WriteString(arg0)

			buffer.WriteString(")")

			calls = append(calls, buffer.String())
			buffer.Reset()
		}
	}
	if input.Spec.Layout.RowsLayoutKind != nil {
		if input.Spec.Layout.RowsLayoutKind != nil {
			buffer.WriteString(`RowsLayout(`)
			arg0 := RowsLayoutConverter(*input.Spec.Layout.RowsLayoutKind)
			buffer.WriteString(arg0)

			buffer.WriteString(")")

			calls = append(calls, buffer.String())
			buffer.Reset()
		}
	}
	if input.Spec.Layout.AutoGridLayoutKind != nil {
		if input.Spec.Layout.AutoGridLayoutKind != nil {
			buffer.WriteString(`AutoGridLayout(`)
			arg0 := AutoGridLayoutConverter(*input.Spec.Layout.AutoGridLayoutKind)
			buffer.WriteString(arg0)

			buffer.WriteString(")")

			calls = append(calls, buffer.String())
			buffer.Reset()
		}
	}
	if input.Spec.Layout.TabsLayoutKind != nil {
		if input.Spec.Layout.TabsLayoutKind != nil {
			buffer.WriteString(`TabsLayout(`)
			arg0 := TabsLayoutConverter(*input.Spec.Layout.TabsLayoutKind)
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

	return strings.Join(calls, ".\t\n")
}
