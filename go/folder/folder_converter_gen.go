// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package folder

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// FolderConverter accepts a `Folder` object and generates the Go code to build this object using builders.
func FolderConverter(input Folder) string {
	calls := []string{
		`folder.NewFolderBuilder()`,
	}
	var buffer strings.Builder
	if input.Uid != "" {

		buffer.WriteString(`Uid(`)
		arg0 := fmt.Sprintf("%#v", input.Uid)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Title != "" {

		buffer.WriteString(`Title(`)
		arg0 := fmt.Sprintf("%#v", input.Title)
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

	return strings.Join(calls, ".\t\n")
}
