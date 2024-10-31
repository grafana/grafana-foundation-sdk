// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package team

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// TeamConverter accepts a `Team` object and generates the Go code to build this object using builders.
func TeamConverter(input Team) string {
	calls := []string{
		`team.NewTeamBuilder(` + fmt.Sprintf("%#v", input.Name) + `)`,
	}
	var buffer strings.Builder
	if input.Name != "" {

		buffer.WriteString(`Name(`)
		arg0 := fmt.Sprintf("%#v", input.Name)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Email != nil && *input.Email != "" {

		buffer.WriteString(`Email(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Email))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
