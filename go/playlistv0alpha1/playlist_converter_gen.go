// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package playlistv0alpha1

import (
	"fmt"
	"strings"
)

// PlaylistConverter accepts a `Playlist` object and generates the Go code to build this object using builders.
func PlaylistConverter(input Playlist) string {
	calls := []string{
		`playlistv0alpha1.NewPlaylistBuilder(` + fmt.Sprintf("%#v", input.Title) + `)`,
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
	if input.Interval != "" {

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
			tmpitemsarg1 := ItemConverter(arg1)
			tmparg0 = append(tmparg0, tmpitemsarg1)
		}
		arg0 := "[]cog.Builder[playlistv0alpha1.Item]{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
