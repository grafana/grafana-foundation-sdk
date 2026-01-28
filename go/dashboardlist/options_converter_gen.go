// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardlist

import (
	"fmt"
	"strings"
)

// OptionsConverter accepts a `Options` object and generates the Go code to build this object using builders.
func OptionsConverter(input Options) string {
	calls := []string{
		`dashboardlist.NewOptionsBuilder()`,
	}
	var buffer strings.Builder
	if input.KeepTime != false {

		buffer.WriteString(`KeepTime(`)
		arg0 := fmt.Sprintf("%#v", input.KeepTime)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.IncludeVars != false {

		buffer.WriteString(`IncludeVars(`)
		arg0 := fmt.Sprintf("%#v", input.IncludeVars)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.ShowStarred != true {

		buffer.WriteString(`ShowStarred(`)
		arg0 := fmt.Sprintf("%#v", input.ShowStarred)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.ShowRecentlyViewed != false {

		buffer.WriteString(`ShowRecentlyViewed(`)
		arg0 := fmt.Sprintf("%#v", input.ShowRecentlyViewed)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.ShowSearch != false {

		buffer.WriteString(`ShowSearch(`)
		arg0 := fmt.Sprintf("%#v", input.ShowSearch)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.ShowHeadings != true {

		buffer.WriteString(`ShowHeadings(`)
		arg0 := fmt.Sprintf("%#v", input.ShowHeadings)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.ShowFolderNames != true {

		buffer.WriteString(`ShowFolderNames(`)
		arg0 := fmt.Sprintf("%#v", input.ShowFolderNames)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.MaxItems != 10 {

		buffer.WriteString(`MaxItems(`)
		arg0 := fmt.Sprintf("%#v", input.MaxItems)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Query != "" {

		buffer.WriteString(`Query(`)
		arg0 := fmt.Sprintf("%#v", input.Query)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Tags != nil && len(input.Tags) >= 1 {

		buffer.WriteString(`Tags(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Tags {
			tmptagsarg1 := fmt.Sprintf("%#v", arg1)
			tmparg0 = append(tmparg0, tmptagsarg1)
		}
		arg0 := "[]string{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.FolderId != nil {

		buffer.WriteString(`FolderId(`)
		arg0 := fmt.Sprintf("%#v", *input.FolderId)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.FolderUID != nil && *input.FolderUID != "" {

		buffer.WriteString(`FolderUID(`)
		arg0 := fmt.Sprintf("%#v", *input.FolderUID)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
