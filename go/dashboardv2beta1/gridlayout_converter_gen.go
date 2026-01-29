// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"strings"
)

// GridLayoutConverter accepts a `GridLayout` object and generates the Go code to build this object using builders.
func GridLayoutConverter(input GridLayoutKind) string {
	calls := []string{
		`dashboardv2beta1.NewGridLayoutBuilder()`,
	}
	var buffer strings.Builder
	if input.Spec.Items != nil && len(input.Spec.Items) >= 1 {

		buffer.WriteString(`Items(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Spec.Items {
			tmpitemsarg1 := GridLayoutItemConverter(arg1)
			tmparg0 = append(tmparg0, tmpitemsarg1)
		}
		arg0 := "[]cog.Builder[dashboardv2beta1.GridLayoutItemKind]{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
