// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package playlist

import (
	"fmt"
	"strings"
)

// PlaylistConverter accepts a `Playlist` object and generates the Go code to build this object using builders.
func PlaylistConverter(input Playlist) string {
	calls := []string{
		`playlist.NewPlaylistBuilder()`,
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
	if input.Name != "" {

		buffer.WriteString(`Name(`)
		arg0 := fmt.Sprintf("%#v", input.Name)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Interval != "" && input.Interval != "5m" {

		buffer.WriteString(`Interval(`)
		arg0 := fmt.Sprintf("%#v", input.Interval)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Items != nil && len(input.Items) >= 1 {

		buffer.WriteString(`Items(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Items {
			tmpitemsarg1 := PlaylistItemConverter(arg1)
			tmparg0 = append(tmparg0, tmpitemsarg1)
		}
		arg0 := "[]cog.Builder[playlist.PlaylistItem]{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
