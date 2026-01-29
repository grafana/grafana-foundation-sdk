// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// AdHocFilterWithLabelsConverter accepts a `AdHocFilterWithLabels` object and generates the Go code to build this object using builders.
func AdHocFilterWithLabelsConverter(input AdHocFilterWithLabels) string {
	calls := []string{
		`dashboardv2beta1.NewAdHocFilterWithLabelsBuilder()`,
	}
	var buffer strings.Builder
	if input.Key != "" {

		buffer.WriteString(`Key(`)
		arg0 := fmt.Sprintf("%#v", input.Key)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Operator != "" {

		buffer.WriteString(`Operator(`)
		arg0 := fmt.Sprintf("%#v", input.Operator)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Value != "" {

		buffer.WriteString(`Value(`)
		arg0 := fmt.Sprintf("%#v", input.Value)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Values != nil && len(input.Values) >= 1 {

		buffer.WriteString(`Values(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Values {
			tmpvaluesarg1 := fmt.Sprintf("%#v", arg1)
			tmparg0 = append(tmparg0, tmpvaluesarg1)
		}
		arg0 := "[]string{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.KeyLabel != nil && *input.KeyLabel != "" {

		buffer.WriteString(`KeyLabel(`)
		arg0 := fmt.Sprintf("%#v", *input.KeyLabel)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.ValueLabels != nil && len(input.ValueLabels) >= 1 {

		buffer.WriteString(`ValueLabels(`)
		tmparg0 := []string{}
		for _, arg1 := range input.ValueLabels {
			tmpvalueLabelsarg1 := fmt.Sprintf("%#v", arg1)
			tmparg0 = append(tmparg0, tmpvalueLabelsarg1)
		}
		arg0 := "[]string{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.ForceEdit != nil {

		buffer.WriteString(`ForceEdit(`)
		arg0 := fmt.Sprintf("%#v", *input.ForceEdit)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Origin != nil {

		buffer.WriteString(`Origin(`)
		arg0 := cog.Dump(*input.Origin)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Condition != nil && *input.Condition != "" {

		buffer.WriteString(`Condition(`)
		arg0 := fmt.Sprintf("%#v", *input.Condition)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
