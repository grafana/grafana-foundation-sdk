// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package librarypanel

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// LibraryPanelConverter accepts a `LibraryPanel` object and generates the Go code to build this object using builders.
func LibraryPanelConverter(input LibraryPanel) string {
	calls := []string{
		`librarypanel.NewLibraryPanelBuilder()`,
	}
	var buffer strings.Builder
	if input.FolderUid != nil && *input.FolderUid != "" {

		buffer.WriteString(`FolderUid(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.FolderUid))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Uid != "" {

		buffer.WriteString(`Uid(`)
		arg0 := fmt.Sprintf("%#v", input.Uid)
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
	if input.Description != nil && *input.Description != "" {

		buffer.WriteString(`Description(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Description))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Type != "" {

		buffer.WriteString(`Type(`)
		arg0 := fmt.Sprintf("%#v", input.Type)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.SchemaVersion != nil {

		buffer.WriteString(`SchemaVersion(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.SchemaVersion))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	{
		buffer.WriteString(`Version(`)
		arg0 := fmt.Sprintf("%#v", input.Version)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	{
		buffer.WriteString(`Model(`)
		arg0 := LibrarypanelLibraryPanelModelConverter(input.Model)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.Meta != nil {

		buffer.WriteString(`Meta(`)
		arg0 := LibraryElementDTOMetaConverter(*input.Meta)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
