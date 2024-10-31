// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package playlist

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// PlaylistItemConverter accepts a `PlaylistItem` object and generates the Go code to build this object using builders.
func PlaylistItemConverter(input PlaylistItem) string {
	calls := []string{
		`playlist.NewPlaylistItemBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`Type(`)
		arg0 := cog.Dump(input.Type)
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
	if input.Title != nil && *input.Title != "" {

		buffer.WriteString(`Title(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Title))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
