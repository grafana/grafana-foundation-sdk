// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// AutoGridLayoutSpecConverter accepts a `AutoGridLayoutSpec` object and generates the Go code to build this object using builders.
func AutoGridLayoutSpecConverter(input AutoGridLayoutSpec) string {
	calls := []string{
		`dashboardv2beta1.NewAutoGridLayoutSpecBuilder()`,
	}
	var buffer strings.Builder
	if input.MaxColumnCount != nil && *input.MaxColumnCount != 3 {

		buffer.WriteString(`MaxColumnCount(`)
		arg0 := fmt.Sprintf("%#v", *input.MaxColumnCount)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	{
		buffer.WriteString(`ColumnWidthMode(`)
		arg0 := cog.Dump(input.ColumnWidthMode)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.ColumnWidth != nil {

		buffer.WriteString(`ColumnWidth(`)
		arg0 := fmt.Sprintf("%#v", *input.ColumnWidth)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	{
		buffer.WriteString(`RowHeightMode(`)
		arg0 := cog.Dump(input.RowHeightMode)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.RowHeight != nil {

		buffer.WriteString(`RowHeight(`)
		arg0 := fmt.Sprintf("%#v", *input.RowHeight)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.FillScreen != nil && *input.FillScreen != false {

		buffer.WriteString(`FillScreen(`)
		arg0 := fmt.Sprintf("%#v", *input.FillScreen)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Items != nil && len(input.Items) >= 1 {

		buffer.WriteString(`Items(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Items {
			tmpitemsarg1 := AutoGridLayoutItemConverter(arg1)
			tmparg0 = append(tmparg0, tmpitemsarg1)
		}
		arg0 := "[]cog.Builder[dashboardv2beta1.AutoGridLayoutItemKind]{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
