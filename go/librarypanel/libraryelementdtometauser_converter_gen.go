// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package librarypanel

import (
	"fmt"
	"strings"
)

// LibraryElementDTOMetaUserConverter accepts a `LibraryElementDTOMetaUser` object and generates the Go code to build this object using builders.
func LibraryElementDTOMetaUserConverter(input LibraryElementDTOMetaUser) string {
	calls := []string{
		`librarypanel.NewLibraryElementDTOMetaUserBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`Id(`)
		arg0 := fmt.Sprintf("%#v", input.Id)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.Name != "" {

		buffer.WriteString(`Name(`)
		arg0 := fmt.Sprintf("%#v", input.Name)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.AvatarUrl != "" {

		buffer.WriteString(`AvatarUrl(`)
		arg0 := fmt.Sprintf("%#v", input.AvatarUrl)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
