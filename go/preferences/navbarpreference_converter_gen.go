// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package preferences

import (
	"fmt"
	"strings"
)

// NavbarPreferenceConverter accepts a `NavbarPreference` object and generates the Go code to build this object using builders.
func NavbarPreferenceConverter(input NavbarPreference) string {
	calls := []string{
		`preferences.NewNavbarPreferenceBuilder()`,
	}
	var buffer strings.Builder
	if input.BookmarkUrls != nil && len(input.BookmarkUrls) >= 1 {

		buffer.WriteString(`BookmarkUrls(`)
		tmparg0 := []string{}
		for _, arg1 := range input.BookmarkUrls {
			tmpbookmarkUrlsarg1 := fmt.Sprintf("%#v", arg1)
			tmparg0 = append(tmparg0, tmpbookmarkUrlsarg1)
		}
		arg0 := "[]string{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
