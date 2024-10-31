// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// ElasticsearchPercentilesSettingsConverter accepts a `ElasticsearchPercentilesSettings` object and generates the Go code to build this object using builders.
func ElasticsearchPercentilesSettingsConverter(input ElasticsearchPercentilesSettings) string {
	calls := []string{
		`elasticsearch.NewElasticsearchPercentilesSettingsBuilder()`,
	}
	var buffer strings.Builder
	if input.Script != nil {

		buffer.WriteString(`Script(`)
		arg0 := InlineScriptConverter(*input.Script)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Missing != nil && *input.Missing != "" {

		buffer.WriteString(`Missing(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Missing))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Percents != nil && len(input.Percents) >= 1 {

		buffer.WriteString(`Percents(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Percents {
			tmppercentsarg1 := fmt.Sprintf("%#v", arg1)
			tmparg0 = append(tmparg0, tmppercentsarg1)
		}
		arg0 := "[]string{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
