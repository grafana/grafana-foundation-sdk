// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package bargauge

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
	dashboard "github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

// PanelConverter accepts a `Panel` object and generates the Go code to build this object using builders.
func PanelConverter(input dashboard.Panel) string {
	calls := []string{
		`bargauge.NewPanelBuilder()`,
	}
	var buffer strings.Builder
	if input.Id != nil {

		buffer.WriteString(`Id(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Id))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Targets != nil && len(input.Targets) >= 1 {

		buffer.WriteString(`Targets(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Targets {
			tmptargetsarg1 := cog.ConvertDataqueryToCode(arg1)
			tmparg0 = append(tmparg0, tmptargetsarg1)
		}
		arg0 := "[]cog.Builder[variants.Dataquery]{" + strings.Join(tmparg0, ",\n") + "}"
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
	if input.Description != nil && *input.Description != "" {

		buffer.WriteString(`Description(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Description))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Transparent != false {

		buffer.WriteString(`Transparent(`)
		arg0 := fmt.Sprintf("%#v", input.Transparent)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Datasource != nil {

		buffer.WriteString(`Datasource(`)
		arg0 := cog.Dump(*input.Datasource)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.GridPos != nil {

		buffer.WriteString(`GridPos(`)
		arg0 := cog.Dump(*input.GridPos)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.GridPos != nil && input.GridPos.H != 9 {

		buffer.WriteString(`Height(`)
		arg0 := fmt.Sprintf("%#v", input.GridPos.H)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.GridPos != nil && input.GridPos.W != 12 {

		buffer.WriteString(`Span(`)
		arg0 := fmt.Sprintf("%#v", input.GridPos.W)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Links != nil && len(input.Links) >= 1 {

		buffer.WriteString(`Links(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Links {
			tmplinksarg1 := dashboard.DashboardLinkConverter(arg1)
			tmparg0 = append(tmparg0, tmplinksarg1)
		}
		arg0 := "[]cog.Builder[dashboard.DashboardLink]{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Repeat != nil && *input.Repeat != "" {

		buffer.WriteString(`Repeat(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Repeat))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.RepeatDirection != nil {

		buffer.WriteString(`RepeatDirection(`)
		arg0 := cog.Dump(*input.RepeatDirection)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.MaxDataPoints != nil {

		buffer.WriteString(`MaxDataPoints(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.MaxDataPoints))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Transformations != nil && len(input.Transformations) >= 1 {

		buffer.WriteString(`Transformations(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Transformations {
			tmptransformationsarg1 := cog.Dump(arg1)
			tmparg0 = append(tmparg0, tmptransformationsarg1)
		}
		arg0 := "[]dashboard.DataTransformerConfig{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Interval != nil && *input.Interval != "" {

		buffer.WriteString(`Interval(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Interval))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.TimeFrom != nil && *input.TimeFrom != "" {

		buffer.WriteString(`TimeFrom(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.TimeFrom))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.TimeShift != nil && *input.TimeShift != "" {

		buffer.WriteString(`TimeShift(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.TimeShift))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.LibraryPanel != nil {

		buffer.WriteString(`LibraryPanel(`)
		arg0 := cog.Dump(*input.LibraryPanel)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.FieldConfig.Defaults.DisplayName != nil && *input.FieldConfig.Defaults.DisplayName != "" {

		buffer.WriteString(`DisplayName(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.FieldConfig.Defaults.DisplayName))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.FieldConfig.Defaults.Unit != nil && *input.FieldConfig.Defaults.Unit != "" {

		buffer.WriteString(`Unit(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.FieldConfig.Defaults.Unit))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.FieldConfig.Defaults.Decimals != nil {

		buffer.WriteString(`Decimals(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.FieldConfig.Defaults.Decimals))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.FieldConfig.Defaults.Min != nil {

		buffer.WriteString(`Min(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.FieldConfig.Defaults.Min))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.FieldConfig.Defaults.Max != nil {

		buffer.WriteString(`Max(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.FieldConfig.Defaults.Max))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.FieldConfig.Defaults.Mappings != nil && len(input.FieldConfig.Defaults.Mappings) >= 1 {

		buffer.WriteString(`Mappings(`)
		tmparg0 := []string{}
		for _, arg1 := range input.FieldConfig.Defaults.Mappings {
			tmpmappingsarg1 := cog.Dump(arg1)
			tmparg0 = append(tmparg0, tmpmappingsarg1)
		}
		arg0 := "[]dashboard.ValueMapping{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.FieldConfig.Defaults.Thresholds != nil {

		buffer.WriteString(`Thresholds(`)
		arg0 := dashboard.ThresholdsConfigConverter(*input.FieldConfig.Defaults.Thresholds)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.FieldConfig.Defaults.Color != nil {

		buffer.WriteString(`ColorScheme(`)
		arg0 := dashboard.FieldColorConverter(*input.FieldConfig.Defaults.Color)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.FieldConfig.Defaults.NoValue != nil && *input.FieldConfig.Defaults.NoValue != "" {

		buffer.WriteString(`NoValue(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.FieldConfig.Defaults.NoValue))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.FieldConfig.Overrides != nil && len(input.FieldConfig.Overrides) >= 1 {

		buffer.WriteString(`Overrides(`)
		tmparg0 := []string{}
		for _, arg1 := range input.FieldConfig.Overrides {
			tmpoverridesarg1 := dashboard.DashboardFieldConfigSourceOverridesConverter(arg1)
			tmparg0 = append(tmparg0, tmpoverridesarg1)
		}
		arg0 := "[]cog.Builder[dashboard.DashboardFieldConfigSourceOverrides]{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.FieldConfig.Overrides != nil && len(input.FieldConfig.Overrides) >= 1 {
		for _, item := range input.FieldConfig.Overrides {

			buffer.WriteString(`WithOverride(`)
			arg0 := cog.Dump(item.Matcher)
			buffer.WriteString(arg0)
			buffer.WriteString(", ")
			tmparg1 := []string{}
			for _, arg1 := range item.Properties {
				tmppropertiesarg1 := cog.Dump(arg1)
				tmparg1 = append(tmparg1, tmppropertiesarg1)
			}
			arg1 := "[]dashboard.DynamicConfigValue{" + strings.Join(tmparg1, ",\n") + "}"
			buffer.WriteString(arg1)

			buffer.WriteString(")")

			calls = append(calls, buffer.String())
			buffer.Reset()

		}
	}
	if input.Options != nil {

		buffer.WriteString(`DisplayMode(`)
		arg0 := cog.Dump(input.Options.(*Options).DisplayMode)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Options != nil {

		buffer.WriteString(`ValueMode(`)
		arg0 := cog.Dump(input.Options.(*Options).ValueMode)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Options != nil && input.Options.(*Options).ShowUnfilled != true {

		buffer.WriteString(`ShowUnfilled(`)
		arg0 := fmt.Sprintf("%#v", input.Options.(*Options).ShowUnfilled)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Options != nil && input.Options.(*Options).MinVizWidth != 0 {

		buffer.WriteString(`MinVizWidth(`)
		arg0 := fmt.Sprintf("%#v", input.Options.(*Options).MinVizWidth)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Options != nil {

		buffer.WriteString(`ReduceOptions(`)
		arg0 := common.ReduceDataOptionsConverter(input.Options.(*Options).ReduceOptions)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Options != nil && input.Options.(*Options).Text != nil {

		buffer.WriteString(`Text(`)
		arg0 := common.VizTextDisplayOptionsConverter(*input.Options.(*Options).Text)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Options != nil && input.Options.(*Options).MinVizHeight != 10 {

		buffer.WriteString(`MinVizHeight(`)
		arg0 := fmt.Sprintf("%#v", input.Options.(*Options).MinVizHeight)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Options != nil {

		buffer.WriteString(`Orientation(`)
		arg0 := cog.Dump(input.Options.(*Options).Orientation)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
