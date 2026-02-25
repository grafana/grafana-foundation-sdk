// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"strings"
)

// RowsConverter accepts a `Rows` object and generates the Go code to build this object using builders.
func RowsConverter(input RowsLayoutKind) string {
	calls := []string{
		`dashboardv2beta1.NewRowsBuilder()`,
	}
	var buffer strings.Builder
	if input.Spec.Rows != nil && len(input.Spec.Rows) >= 1 {

		buffer.WriteString(`Rows(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Spec.Rows {
			tmprowsarg1 := RowConverter(arg1)
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
