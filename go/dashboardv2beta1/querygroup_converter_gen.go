// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"strings"
)

// QueryGroupConverter accepts a `QueryGroup` object and generates the Go code to build this object using builders.
func QueryGroupConverter(input QueryGroupKind) string {
	calls := []string{
		`dashboardv2beta1.NewQueryGroupBuilder()`,
	}
	var buffer strings.Builder
	if input.Spec.Queries != nil && len(input.Spec.Queries) >= 1 {

		buffer.WriteString(`Targets(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Spec.Queries {
			tmpqueriesarg1 := TargetConverter(arg1)
			tmparg0 = append(tmparg0, tmpqueriesarg1)
		}
		arg0 := "[]cog.Builder[dashboardv2beta1.PanelQueryKind]{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.Transformations != nil && len(input.Spec.Transformations) >= 1 {

		buffer.WriteString(`Transformations(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Spec.Transformations {
			tmptransformationsarg1 := TransformationConverter(arg1)
			tmparg0 = append(tmparg0, tmptransformationsarg1)
		}
		arg0 := "[]cog.Builder[dashboardv2beta1.TransformationKind]{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	{
		buffer.WriteString(`QueryOptions(`)
		arg0 := QueryOptionsSpecConverter(input.Spec.QueryOptions)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	return strings.Join(calls, ".\t\n")
}
