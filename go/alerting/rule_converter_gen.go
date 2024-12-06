// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alerting

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// RuleConverter accepts a `Rule` object and generates the Go code to build this object using builders.
func RuleConverter(input Rule) string {
	calls := []string{
		`alerting.NewRuleBuilder(` + fmt.Sprintf("%#v", input.Title) + `)`,
	}
	var buffer strings.Builder
	if input.Annotations != nil {

		buffer.WriteString(`Annotations(`)
		arg0 := "map[string]string{"
		for key, arg1 := range input.Annotations {
			tmpannotationsarg1 := fmt.Sprintf("%#v", arg1)
			arg0 += "\t" + fmt.Sprintf("%#v", key) + ": " + tmpannotationsarg1 + ","
		}
		arg0 += "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Condition != "" {

		buffer.WriteString(`Condition(`)
		arg0 := fmt.Sprintf("%#v", input.Condition)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Data != nil && len(input.Data) >= 1 {

		buffer.WriteString(`Queries(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Data {
			tmpdataarg1 := QueryConverter(arg1)
			tmparg0 = append(tmparg0, tmpdataarg1)
		}
		arg0 := "[]cog.Builder[alerting.Query]{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	{
		buffer.WriteString(`ExecErrState(`)
		arg0 := cog.Dump(input.ExecErrState)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.FolderUID != "" {

		buffer.WriteString(`FolderUID(`)
		arg0 := fmt.Sprintf("%#v", input.FolderUID)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.For != "" {

		buffer.WriteString(`For(`)
		arg0 := fmt.Sprintf("%#v", input.For)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Id != nil {

		buffer.WriteString(`Id(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Id))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.IsPaused != nil {

		buffer.WriteString(`IsPaused(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.IsPaused))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Labels != nil {

		buffer.WriteString(`Labels(`)
		arg0 := "map[string]string{"
		for key, arg1 := range input.Labels {
			tmplabelsarg1 := fmt.Sprintf("%#v", arg1)
			arg0 += "\t" + fmt.Sprintf("%#v", key) + ": " + tmplabelsarg1 + ","
		}
		arg0 += "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	{
		buffer.WriteString(`NoDataState(`)
		arg0 := cog.Dump(input.NoDataState)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.NotificationSettings != nil {

		buffer.WriteString(`NotificationSettings(`)
		arg0 := NotificationSettingsConverter(*input.NotificationSettings)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	{
		buffer.WriteString(`OrgID(`)
		arg0 := fmt.Sprintf("%#v", input.OrgID)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.Provenance != nil {

		buffer.WriteString(`Provenance(`)
		arg0 := cog.Dump(*input.Provenance)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Record != nil {

		buffer.WriteString(`Record(`)
		arg0 := RecordRuleConverter(*input.Record)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.RuleGroup != "" {

		buffer.WriteString(`RuleGroup(`)
		arg0 := fmt.Sprintf("%#v", input.RuleGroup)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Title != "" {

		buffer.WriteString(`Title(`)
		arg0 := fmt.Sprintf("%#v", input.Title)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Uid != nil && *input.Uid != "" {

		buffer.WriteString(`Uid(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Uid))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Updated != nil {

		buffer.WriteString(`Updated(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Updated))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
