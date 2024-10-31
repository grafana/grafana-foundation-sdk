// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package librarypanel

import (
	"fmt"
	"strings"
)

// LibraryElementDTOMetaConverter accepts a `LibraryElementDTOMeta` object and generates the Go code to build this object using builders.
func LibraryElementDTOMetaConverter(input LibraryElementDTOMeta) string {
	calls := []string{
		`librarypanel.NewLibraryElementDTOMetaBuilder()`,
	}
	var buffer strings.Builder
	if input.FolderName != "" {

		buffer.WriteString(`FolderName(`)
		arg0 := fmt.Sprintf("%#v", input.FolderName)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.FolderUid != "" {

		buffer.WriteString(`FolderUid(`)
		arg0 := fmt.Sprintf("%#v", input.FolderUid)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	{
		buffer.WriteString(`ConnectedDashboards(`)
		arg0 := fmt.Sprintf("%#v", input.ConnectedDashboards)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	{
		buffer.WriteString(`Created(`)
		arg0 := fmt.Sprintf("%#v", input.Created)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	{
		buffer.WriteString(`Updated(`)
		arg0 := fmt.Sprintf("%#v", input.Updated)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	{
		buffer.WriteString(`CreatedBy(`)
		arg0 := LibraryElementDTOMetaUserConverter(input.CreatedBy)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	{
		buffer.WriteString(`UpdatedBy(`)
		arg0 := LibraryElementDTOMetaUserConverter(input.UpdatedBy)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	return strings.Join(calls, ".\t\n")
}
