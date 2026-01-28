// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package table

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
)

// OptionsConverter accepts a `Options` object and generates the Go code to build this object using builders.
func OptionsConverter(input Options) string {
	calls := []string{
		`table.NewOptionsBuilder()`,
	}
	var buffer strings.Builder
	if input.FrameIndex != 0 {

		buffer.WriteString(`FrameIndex(`)
		arg0 := fmt.Sprintf("%#v", input.FrameIndex)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.ShowHeader != true {

		buffer.WriteString(`ShowHeader(`)
		arg0 := fmt.Sprintf("%#v", input.ShowHeader)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.ShowTypeIcons != nil && *input.ShowTypeIcons != false {

		buffer.WriteString(`ShowTypeIcons(`)
		arg0 := fmt.Sprintf("%#v", *input.ShowTypeIcons)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.SortBy != nil && len(input.SortBy) >= 1 {

		buffer.WriteString(`SortBy(`)
		tmparg0 := []string{}
		for _, arg1 := range input.SortBy {
			tmpsortByarg1 := common.TableSortByFieldStateConverter(arg1)
			tmparg0 = append(tmparg0, tmpsortByarg1)
		}
		arg0 := "[]cog.Builder[common.TableSortByFieldState]{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Footer != nil {

		buffer.WriteString(`Footer(`)
		arg0 := common.TableFooterOptionsConverter(*input.Footer)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.CellHeight != nil {

		buffer.WriteString(`CellHeight(`)
		arg0 := cog.Dump(*input.CellHeight)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
