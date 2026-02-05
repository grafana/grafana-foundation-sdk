// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package barchart

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
	dashboardv2beta1 "github.com/grafana/grafana-foundation-sdk/go/dashboardv2beta1"
)

// VisualizationConverter accepts a `Visualization` object and generates the Go code to build this object using builders.
func VisualizationConverter(input dashboardv2beta1.VizConfigKind) string {
	calls := []string{
		`barchart.NewVisualizationBuilder()`,
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
	if input.Spec.Options != nil && input.Spec.Options.(*Options).XField != nil && *input.Spec.Options.(*Options).XField != "" {

		buffer.WriteString(`XField(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.Options.(*Options).XField)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.Options != nil && input.Spec.Options.(*Options).ColorByField != nil && *input.Spec.Options.(*Options).ColorByField != "" {

		buffer.WriteString(`ColorByField(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.Options.(*Options).ColorByField)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.Options != nil {

		buffer.WriteString(`Orientation(`)
		arg0 := cog.Dump(input.Spec.Options.(*Options).Orientation)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.Options != nil && input.Spec.Options.(*Options).BarRadius != nil && *input.Spec.Options.(*Options).BarRadius != 0 {

		buffer.WriteString(`BarRadius(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.Options.(*Options).BarRadius)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.Options != nil && input.Spec.Options.(*Options).XTickLabelRotation != 0 {

		buffer.WriteString(`XTickLabelRotation(`)
		arg0 := fmt.Sprintf("%#v", input.Spec.Options.(*Options).XTickLabelRotation)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.Options != nil {

		buffer.WriteString(`XTickLabelMaxLength(`)
		arg0 := fmt.Sprintf("%#v", input.Spec.Options.(*Options).XTickLabelMaxLength)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.Options != nil && input.Spec.Options.(*Options).XTickLabelSpacing != nil && *input.Spec.Options.(*Options).XTickLabelSpacing != 0 {

		buffer.WriteString(`XTickLabelSpacing(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.Options.(*Options).XTickLabelSpacing)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.Options != nil {

		buffer.WriteString(`Stacking(`)
		arg0 := cog.Dump(input.Spec.Options.(*Options).Stacking)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.Options != nil {

		buffer.WriteString(`ShowValue(`)
		arg0 := cog.Dump(input.Spec.Options.(*Options).ShowValue)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.Options != nil && input.Spec.Options.(*Options).BarWidth != 0.97 {

		buffer.WriteString(`BarWidth(`)
		arg0 := fmt.Sprintf("%#v", input.Spec.Options.(*Options).BarWidth)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.Options != nil && input.Spec.Options.(*Options).GroupWidth != 0.7 {

		buffer.WriteString(`GroupWidth(`)
		arg0 := fmt.Sprintf("%#v", input.Spec.Options.(*Options).GroupWidth)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.Options != nil {

		buffer.WriteString(`Legend(`)
		arg0 := common.VizLegendOptionsConverter(input.Spec.Options.(*Options).Legend)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.Options != nil {

		buffer.WriteString(`Tooltip(`)
		arg0 := common.VizTooltipOptionsConverter(input.Spec.Options.(*Options).Tooltip)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.Options != nil && input.Spec.Options.(*Options).Text != nil {

		buffer.WriteString(`Text(`)
		arg0 := common.VizTextDisplayOptionsConverter(*input.Spec.Options.(*Options).Text)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.Options != nil && input.Spec.Options.(*Options).FullHighlight != false {

		buffer.WriteString(`FullHighlight(`)
		arg0 := fmt.Sprintf("%#v", input.Spec.Options.(*Options).FullHighlight)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.FieldConfig.Defaults.Custom != nil && input.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).LineWidth != nil && *input.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).LineWidth != 1 {

		buffer.WriteString(`LineWidth(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).LineWidth)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.FieldConfig.Defaults.Custom != nil && input.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).FillOpacity != nil && *input.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).FillOpacity != 80 {

		buffer.WriteString(`FillOpacity(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).FillOpacity)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.FieldConfig.Defaults.Custom != nil && input.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).GradientMode != nil {

		buffer.WriteString(`GradientMode(`)
		arg0 := cog.Dump(*input.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).GradientMode)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.FieldConfig.Defaults.Custom != nil && input.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).AxisPlacement != nil {

		buffer.WriteString(`AxisPlacement(`)
		arg0 := cog.Dump(*input.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).AxisPlacement)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.FieldConfig.Defaults.Custom != nil && input.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).AxisColorMode != nil {

		buffer.WriteString(`AxisColorMode(`)
		arg0 := cog.Dump(*input.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).AxisColorMode)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.FieldConfig.Defaults.Custom != nil && input.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).AxisLabel != nil && *input.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).AxisLabel != "" {

		buffer.WriteString(`AxisLabel(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).AxisLabel)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.FieldConfig.Defaults.Custom != nil && input.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).AxisWidth != nil {

		buffer.WriteString(`AxisWidth(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).AxisWidth)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.FieldConfig.Defaults.Custom != nil && input.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).AxisSoftMin != nil {

		buffer.WriteString(`AxisSoftMin(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).AxisSoftMin)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.FieldConfig.Defaults.Custom != nil && input.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).AxisSoftMax != nil {

		buffer.WriteString(`AxisSoftMax(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).AxisSoftMax)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.FieldConfig.Defaults.Custom != nil && input.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).AxisGridShow != nil {

		buffer.WriteString(`AxisGridShow(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).AxisGridShow)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.FieldConfig.Defaults.Custom != nil && input.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).ScaleDistribution != nil {

		buffer.WriteString(`ScaleDistribution(`)
		arg0 := common.ScaleDistributionConfigConverter(*input.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).ScaleDistribution)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.FieldConfig.Defaults.Custom != nil && input.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).AxisCenteredZero != nil {

		buffer.WriteString(`AxisCenteredZero(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).AxisCenteredZero)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.FieldConfig.Defaults.Custom != nil && input.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).HideFrom != nil {

		buffer.WriteString(`HideFrom(`)
		arg0 := common.HideSeriesConfigConverter(*input.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).HideFrom)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.FieldConfig.Defaults.Custom != nil && input.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).ThresholdsStyle != nil {

		buffer.WriteString(`ThresholdsStyle(`)
		arg0 := common.GraphThresholdsStyleConfigConverter(*input.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).ThresholdsStyle)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec.FieldConfig.Defaults.Custom != nil && input.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).AxisBorderShow != nil {

		buffer.WriteString(`AxisBorderShow(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).AxisBorderShow)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
