// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"strings"
)

// RowsLayoutSpecConverter accepts a `RowsLayoutSpec` object and generates the Go code to build this object using builders.
func RowsLayoutSpecConverter(input RowsLayoutSpec) string {
	calls := []string{
		`dashboardv2beta1.NewRowsLayoutSpecBuilder()`,
	}
	var buffer strings.Builder
	if input.Rows != nil && len(input.Rows) >= 1 {

		buffer.WriteString(`Rows(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Rows {
			tmprowsarg1 := RowsLayoutRowConverter(arg1)
			tmparg0 = append(tmparg0, tmprowsarg1)
		}
		arg0 := "[]cog.Builder[dashboardv2beta1.RowsLayoutRowKind]{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
