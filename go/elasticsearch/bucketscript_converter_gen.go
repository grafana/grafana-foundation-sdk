// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

func BucketScriptConverter(input BucketScript) string {
	calls := []string{
		`elasticsearch.NewBucketScriptBuilder()`,
	}
	var buffer strings.Builder
	if input.PipelineVariables != nil && len(input.PipelineVariables) >= 1 {

		buffer.WriteString(`PipelineVariables(`)
		tmparg0 := []string{}
		for _, arg1 := range input.PipelineVariables {
			tmppipelineVariablesarg1 := PipelineVariableConverter(arg1)
			tmparg0 = append(tmparg0, tmppipelineVariablesarg1)
		}
		arg0 := "[]cog.Builder[elasticsearch.PipelineVariable]{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Id != "" {

		buffer.WriteString(`Id(`)
		arg0 := fmt.Sprintf("%#v", input.Id)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Settings != nil {

		buffer.WriteString(`Settings(`)
		arg0 := ElasticsearchBucketScriptSettingsConverter(*input.Settings)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Hide != nil {

		buffer.WriteString(`Hide(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Hide))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
