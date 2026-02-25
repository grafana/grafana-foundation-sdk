// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// AutoGridConverter accepts a `AutoGrid` object and generates the Go code to build this object using builders.
func AutoGridConverter(input AutoGridLayoutKind) string {
	calls := []string{
		`dashboardv2beta1.NewAutoGridBuilder()`,
	}
	var buffer strings.Builder
	if input.Spec.MaxColumnCount != nil && *input.Spec.MaxColumnCount != 3 {

		buffer.WriteString(`MaxColumnCount(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.MaxColumnCount)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	{
		buffer.WriteString(`ColumnWidthMode(`)
		arg0 := cog.Dump(input.Spec.ColumnWidthMode)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.Spec.ColumnWidth != nil {

		buffer.WriteString(`ColumnWidth(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.ColumnWidth)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	{
		buffer.WriteString(`RowHeightMode(`)
		arg0 := cog.Dump(input.Spec.RowHeightMode)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.Spec.RowHeight != nil {

		buffer.WriteString(`RowHeight(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.RowHeight)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.FillScreen != nil && *input.Spec.FillScreen != false {

		buffer.WriteString(`FillScreen(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.FillScreen)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.Items != nil && len(input.Spec.Items) >= 1 {

		buffer.WriteString(`Items(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Spec.Items {
			tmpitemsarg1 := AutoGridItemConverter(arg1)
			tmparg0 = append(tmparg0, tmpitemsarg1)
		}
		arg0 := "[]cog.Builder[dashboardv2beta1.AutoGridLayoutItemKind]{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.Items != nil && len(input.Spec.Items) >= 1 {
		for _, item := range input.Spec.Items {

			buffer.WriteString(`WithItem(`)
			arg0 := fmt.Sprintf("%#v", item.Kind)
			buffer.WriteString(arg0)
			buffer.WriteString(", ")
			arg1 := cog.Dump(item.Spec)
			buffer.WriteString(arg1)

			buffer.WriteString(")")

			calls = append(calls, buffer.String())
			buffer.Reset()

		}
	}

	return strings.Join(calls, ".\t\n")
}
