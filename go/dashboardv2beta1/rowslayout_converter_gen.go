// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"strings"
)

// RowsLayoutConverter accepts a `RowsLayout` object and generates the Go code to build this object using builders.
func RowsLayoutConverter(input RowsLayoutKind) string {
	calls := []string{
		`dashboardv2beta1.NewRowsLayoutBuilder()`,
	}
	var buffer strings.Builder
	if input.Spec.Rows != nil && len(input.Spec.Rows) >= 1 {

		buffer.WriteString(`Rows(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Spec.Rows {
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
