// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package datagrid

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	dashboardv2beta1 "github.com/grafana/grafana-foundation-sdk/go/dashboardv2beta1"
)

// VisualizationConverter accepts a `Visualization` object and generates the Go code to build this object using builders.
func VisualizationConverter(input dashboardv2beta1.VizConfigKind) string {
	calls := []string{
		`datagrid.NewVisualizationBuilder()`,
	}
	var buffer strings.Builder
	if input.Spec.FieldConfig.Defaults.DisplayName != nil && *input.Spec.FieldConfig.Defaults.DisplayName != "" {

		buffer.WriteString(`DisplayName(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.FieldConfig.Defaults.DisplayName)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.FieldConfig.Defaults.DisplayNameFromDS != nil && *input.Spec.FieldConfig.Defaults.DisplayNameFromDS != "" {

		buffer.WriteString(`DisplayNameFromDS(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.FieldConfig.Defaults.DisplayNameFromDS)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.FieldConfig.Defaults.Description != nil && *input.Spec.FieldConfig.Defaults.Description != "" {

		buffer.WriteString(`Description(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.FieldConfig.Defaults.Description)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.FieldConfig.Defaults.Path != nil && *input.Spec.FieldConfig.Defaults.Path != "" {

		buffer.WriteString(`Path(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.FieldConfig.Defaults.Path)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.FieldConfig.Defaults.Unit != nil && *input.Spec.FieldConfig.Defaults.Unit != "" {

		buffer.WriteString(`Unit(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.FieldConfig.Defaults.Unit)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.FieldConfig.Defaults.Decimals != nil {

		buffer.WriteString(`Decimals(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.FieldConfig.Defaults.Decimals)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.FieldConfig.Defaults.Min != nil {

		buffer.WriteString(`Min(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.FieldConfig.Defaults.Min)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.FieldConfig.Defaults.Max != nil {

		buffer.WriteString(`Max(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.FieldConfig.Defaults.Max)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.FieldConfig.Defaults.Mappings != nil && len(input.Spec.FieldConfig.Defaults.Mappings) >= 1 {

		buffer.WriteString(`Mappings(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Spec.FieldConfig.Defaults.Mappings {
			tmpmappingsarg1 := cog.Dump(arg1)
			tmparg0 = append(tmparg0, tmpmappingsarg1)
		}
		arg0 := "[]dashboardv2beta1.ValueMapping{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.FieldConfig.Defaults.Thresholds != nil {

		buffer.WriteString(`Thresholds(`)
		arg0 := dashboardv2beta1.ThresholdsConfigConverter(*input.Spec.FieldConfig.Defaults.Thresholds)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.FieldConfig.Defaults.Color != nil {

		buffer.WriteString(`ColorScheme(`)
		arg0 := dashboardv2beta1.FieldColorConverter(*input.Spec.FieldConfig.Defaults.Color)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.FieldConfig.Defaults.Links != nil && len(input.Spec.FieldConfig.Defaults.Links) >= 1 {

		buffer.WriteString(`DataLinks(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Spec.FieldConfig.Defaults.Links {
			tmplinksarg1 := cog.Dump(arg1)
			tmparg0 = append(tmparg0, tmplinksarg1)
		}
		arg0 := "[]any{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.FieldConfig.Defaults.Actions != nil && len(input.Spec.FieldConfig.Defaults.Actions) >= 1 {

		buffer.WriteString(`Actions(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Spec.FieldConfig.Defaults.Actions {
			tmpactionsarg1 := dashboardv2beta1.ActionConverter(arg1)
			tmparg0 = append(tmparg0, tmpactionsarg1)
		}
		arg0 := "[]cog.Builder[dashboardv2beta1.Action]{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.FieldConfig.Defaults.NoValue != nil && *input.Spec.FieldConfig.Defaults.NoValue != "" {

		buffer.WriteString(`NoValue(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.FieldConfig.Defaults.NoValue)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.FieldConfig.Overrides != nil && len(input.Spec.FieldConfig.Overrides) >= 1 {

		buffer.WriteString(`Overrides(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Spec.FieldConfig.Overrides {
			tmpoverridesarg1 := dashboardv2beta1.Dashboardv2beta1FieldConfigSourceOverridesConverter(arg1)
			tmparg0 = append(tmparg0, tmpoverridesarg1)
		}
		arg0 := "[]cog.Builder[dashboardv2beta1.Dashboardv2beta1FieldConfigSourceOverrides]{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.FieldConfig.Overrides != nil && len(input.Spec.FieldConfig.Overrides) >= 1 {
		for _, item := range input.Spec.FieldConfig.Overrides {

			buffer.WriteString(`Override(`)
			arg0 := cog.Dump(item.Matcher)
			buffer.WriteString(arg0)
			buffer.WriteString(", ")
			tmparg1 := []string{}
			for _, arg1 := range item.Properties {
				tmppropertiesarg1 := cog.Dump(arg1)
				tmparg1 = append(tmparg1, tmppropertiesarg1)
			}
			arg1 := "[]dashboardv2beta1.DynamicConfigValue{" + strings.Join(tmparg1, ",\n") + "}"
			buffer.WriteString(arg1)

			buffer.WriteString(")")

			calls = append(calls, buffer.String())
			buffer.Reset()

		}
	}
	if input.Spec.Options != nil && input.Spec.Options.(*Options).SelectedSeries != 0 {

		buffer.WriteString(`SelectedSeries(`)
		arg0 := fmt.Sprintf("%#v", input.Spec.Options.(*Options).SelectedSeries)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
