// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alerting

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// RuleGroupConverter accepts a `RuleGroup` object and generates the Go code to build this object using builders.
func RuleGroupConverter(input RuleGroup) string {
	calls := []string{
		`alerting.NewRuleGroupBuilder(` + fmt.Sprintf("%#v", cog.Unptr(input.Title)) + `)`,
	}
	var buffer strings.Builder
	if input.FolderUid != nil && *input.FolderUid != "" {

		buffer.WriteString(`FolderUid(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.FolderUid))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Interval != nil {

		buffer.WriteString(`Interval(`)
		arg0 := cog.Dump(*input.Interval)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Rules != nil && len(input.Rules) >= 1 {

		buffer.WriteString(`Rules(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Rules {
			tmprulesarg1 := RuleConverter(arg1)
			tmparg0 = append(tmparg0, tmprulesarg1)
		}
		arg0 := "[]cog.Builder[alerting.Rule]{" + strings.Join(tmparg0, ",\n") + "}"
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
