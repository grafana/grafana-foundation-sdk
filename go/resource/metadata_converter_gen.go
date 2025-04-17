// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package resource

import (
	"fmt"
	"strings"
)

// MetadataConverter accepts a `Metadata` object and generates the Go code to build this object using builders.
func MetadataConverter(input Metadata) string {
	calls := []string{
		`resource.NewMetadataBuilder()`,
	}
	var buffer strings.Builder
	if input.Name != "" {

		buffer.WriteString(`Name(`)
		arg0 := fmt.Sprintf("%#v", input.Name)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Namespace != nil && *input.Namespace != "" {

		buffer.WriteString(`Namespace(`)
		arg0 := fmt.Sprintf("%#v", *input.Namespace)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Labels != nil {

		buffer.WriteString(`Labels(`)
		arg0 := "map[string]string{"
		for key, arg1 := range input.Labels {
			tmplabelsarg1 := fmt.Sprintf("%#v", arg1)
			arg0 += "\t" + fmt.Sprintf("%#v", key) + ": " + tmplabelsarg1 + ","
		}
		arg0 += "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Annotations != nil {

		buffer.WriteString(`Annotations(`)
		arg0 := "map[string]string{"
		for key, arg1 := range input.Annotations {
			tmpannotationsarg1 := fmt.Sprintf("%#v", arg1)
			arg0 += "\t" + fmt.Sprintf("%#v", key) + ": " + tmpannotationsarg1 + ","
		}
		arg0 += "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Uid != nil && *input.Uid != "" {

		buffer.WriteString(`Uid(`)
		arg0 := fmt.Sprintf("%#v", *input.Uid)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.ResourceVersion != nil && *input.ResourceVersion != "" {

		buffer.WriteString(`ResourceVersion(`)
		arg0 := fmt.Sprintf("%#v", *input.ResourceVersion)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Generation != nil {

		buffer.WriteString(`Generation(`)
		arg0 := fmt.Sprintf("%#v", *input.Generation)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.CreationTimestamp != nil && *input.CreationTimestamp != "" {

		buffer.WriteString(`CreationTimestamp(`)
		arg0 := fmt.Sprintf("%#v", *input.CreationTimestamp)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.UpdateTimestamp != nil && *input.UpdateTimestamp != "" {

		buffer.WriteString(`UpdateTimestamp(`)
		arg0 := fmt.Sprintf("%#v", *input.UpdateTimestamp)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.DeletionTimestamp != nil && *input.DeletionTimestamp != "" {

		buffer.WriteString(`DeletionTimestamp(`)
		arg0 := fmt.Sprintf("%#v", *input.DeletionTimestamp)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
