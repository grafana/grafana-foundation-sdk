// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package folderv1beta1

import (
	"fmt"
	"strings"
)

// FolderConverter accepts a `Folder` object and generates the Go code to build this object using builders.
func FolderConverter(input Folder) string {
	calls := []string{
		`folderv1beta1.NewFolderBuilder(` + fmt.Sprintf("%#v", input.Title) + `)`,
	}
	var buffer strings.Builder
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
		arg0 := fmt.Sprintf("%#v", *input.Description)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
