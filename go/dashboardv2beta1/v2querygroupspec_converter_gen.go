// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"strings"
)

// V2QueryGroupSpecConverter accepts a `V2QueryGroupSpec` object and generates the Go code to build this object using builders.
func V2QueryGroupSpecConverter(input V2QueryGroupSpec) string {
	calls := []string{
		`dashboardv2beta1.NewV2QueryGroupSpecBuilder()`,
	}
	var buffer strings.Builder
	if input.Queries != nil && len(input.Queries) >= 1 {

		buffer.WriteString(`Queries(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Queries {
			tmpqueriesarg1 := TargetConverter(arg1)
			tmparg0 = append(tmparg0, tmpqueriesarg1)
		}
		arg0 := "[]cog.Builder[dashboardv2beta1.PanelQueryKind]{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Transformations != nil && len(input.Transformations) >= 1 {

		buffer.WriteString(`Transformations(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Transformations {
			tmptransformationsarg1 := V2TransformationKindConverter(arg1)
			tmparg0 = append(tmparg0, tmptransformationsarg1)
		}
		arg0 := "[]cog.Builder[dashboardv2beta1.V2TransformationKind]{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	{
		buffer.WriteString(`QueryOptions(`)
		arg0 := QueryOptionsSpecConverter(input.QueryOptions)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	return strings.Join(calls, ".\t\n")
}
