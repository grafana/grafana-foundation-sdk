// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package testdata

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// SimulationQueryConverter accepts a `SimulationQuery` object and generates the Go code to build this object using builders.
func SimulationQueryConverter(input SimulationQuery) string {
	calls := []string{
		`testdata.NewSimulationQueryBuilder()`,
	}
	var buffer strings.Builder

	{
		buffer.WriteString(`Key(`)
		arg0 := KeyConverter(input.Key)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.Config != nil {

		buffer.WriteString(`Config(`)
		arg0 := "map[string]any{"
		for key, arg1 := range input.Config {
			tmpconfigarg1 := cog.Dump(arg1)
			arg0 += "\t" + fmt.Sprintf("%#v", key) + ": " + tmpconfigarg1 + ","
		}
		arg0 += "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Stream != nil {

		buffer.WriteString(`Stream(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Stream))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Last != nil {

		buffer.WriteString(`Last(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Last))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
