// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alertingv0alpha1

import (
	"fmt"
	"strings"
)

// RuleSearchHitFieldsConverter accepts a `RuleSearchHitFields` object and generates the Go code to build this object using builders.
func RuleSearchHitFieldsConverter(input RuleSearchHitFields) string {
	calls := []string{
		`alertingv0alpha1.NewRuleSearchHitFieldsBuilder()`,
	}
	var buffer strings.Builder
	if input.Title != nil && *input.Title != "" {

		buffer.WriteString(`Title(`)
		arg0 := fmt.Sprintf("%#v", *input.Title)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Folder != nil && *input.Folder != "" {

		buffer.WriteString(`Folder(`)
		arg0 := fmt.Sprintf("%#v", *input.Folder)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Type != nil && *input.Type != "" {

		buffer.WriteString(`Type(`)
		arg0 := fmt.Sprintf("%#v", *input.Type)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Interval != nil && *input.Interval != "" {

		buffer.WriteString(`Interval(`)
		arg0 := fmt.Sprintf("%#v", *input.Interval)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Paused != nil {

		buffer.WriteString(`Paused(`)
		arg0 := fmt.Sprintf("%#v", *input.Paused)
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
	if input.DatasourceUIDs != nil && len(input.DatasourceUIDs) >= 1 {

		buffer.WriteString(`DatasourceUIDs(`)
		tmparg0 := []string{}
		for _, arg1 := range input.DatasourceUIDs {
			tmpdatasourceUIDsarg1 := fmt.Sprintf("%#v", arg1)
			tmparg0 = append(tmparg0, tmpdatasourceUIDsarg1)
		}
		arg0 := "[]string{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
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
	if input.For != nil && *input.For != "" {

		buffer.WriteString(`For(`)
		arg0 := fmt.Sprintf("%#v", *input.For)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.KeepFiringFor != nil && *input.KeepFiringFor != "" {

		buffer.WriteString(`KeepFiringFor(`)
		arg0 := fmt.Sprintf("%#v", *input.KeepFiringFor)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.DashboardUID != nil && *input.DashboardUID != "" {

		buffer.WriteString(`DashboardUID(`)
		arg0 := fmt.Sprintf("%#v", *input.DashboardUID)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.PanelID != nil {

		buffer.WriteString(`PanelID(`)
		arg0 := fmt.Sprintf("%#v", *input.PanelID)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Receiver != nil && *input.Receiver != "" {

		buffer.WriteString(`Receiver(`)
		arg0 := fmt.Sprintf("%#v", *input.Receiver)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.NotificationType != nil && *input.NotificationType != "" {

		buffer.WriteString(`NotificationType(`)
		arg0 := fmt.Sprintf("%#v", *input.NotificationType)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.RoutingTree != nil && *input.RoutingTree != "" {

		buffer.WriteString(`RoutingTree(`)
		arg0 := fmt.Sprintf("%#v", *input.RoutingTree)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Metric != nil && *input.Metric != "" {

		buffer.WriteString(`Metric(`)
		arg0 := fmt.Sprintf("%#v", *input.Metric)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.TargetDatasourceUID != nil && *input.TargetDatasourceUID != "" {

		buffer.WriteString(`TargetDatasourceUID(`)
		arg0 := fmt.Sprintf("%#v", *input.TargetDatasourceUID)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
