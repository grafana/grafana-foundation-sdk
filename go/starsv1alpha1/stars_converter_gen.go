// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package starsv1alpha1

import (
	"strings"
)

// StarsConverter accepts a `Stars` object and generates the Go code to build this object using builders.
func StarsConverter(input Stars) string {
	calls := []string{
		`starsv1alpha1.NewStarsBuilder()`,
	}
	var buffer strings.Builder
	if input.Resource != nil && len(input.Resource) >= 1 {

		buffer.WriteString(`Resources(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Resource {
			tmpresourcearg1 := ResourceConverter(arg1)
			tmparg0 = append(tmparg0, tmpresourcearg1)
		}
		arg0 := "[]cog.Builder[starsv1alpha1.Resource]{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
