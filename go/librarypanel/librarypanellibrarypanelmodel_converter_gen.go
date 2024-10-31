// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package librarypanel

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	dashboard "github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

// LibrarypanelLibraryPanelModelConverter accepts a `LibrarypanelLibraryPanelModel` object and generates the Go code to build this object using builders.
func LibrarypanelLibraryPanelModelConverter(input LibrarypanelLibraryPanelModel) string {
	calls := []string{
		`librarypanel.NewLibrarypanelLibraryPanelModelBuilder()`,
	}
	var buffer strings.Builder
	if input.Type != "" {

		buffer.WriteString(`Type(`)
		arg0 := fmt.Sprintf("%#v", input.Type)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.PluginVersion != nil && *input.PluginVersion != "" {

		buffer.WriteString(`PluginVersion(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.PluginVersion))
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
	if input.Transparent != nil && *input.Transparent != false {

		buffer.WriteString(`Transparent(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Transparent))
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
	if input.MaxPerRow != nil {

		buffer.WriteString(`MaxPerRow(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.MaxPerRow))
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
	if input.HideTimeOverride != nil {

		buffer.WriteString(`HideTimeOverride(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.HideTimeOverride))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.CacheTimeout != nil && *input.CacheTimeout != "" {

		buffer.WriteString(`CacheTimeout(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.CacheTimeout))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.QueryCachingTTL != nil {

		buffer.WriteString(`QueryCachingTTL(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.QueryCachingTTL))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Options != nil {

		buffer.WriteString(`Options(`)
		arg0 := cog.Dump(input.Options)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.FieldConfig != nil {

		buffer.WriteString(`FieldConfig(`)
		arg0 := cog.Dump(*input.FieldConfig)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
