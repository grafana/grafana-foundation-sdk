// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// DashboardConverter accepts a `Dashboard` object and generates the Go code to build this object using builders.
func DashboardConverter(input Dashboard) string {
	calls := []string{
		`dashboardv2beta1.NewDashboardBuilder(` + fmt.Sprintf("%#v", input.Title) + `)`,
	}
	var buffer strings.Builder
	if input.Annotations != nil && len(input.Annotations) >= 1 {

		buffer.WriteString(`Annotations(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Annotations {
			tmpannotationsarg1 := AnnotationQueryConverter(arg1)
			tmparg0 = append(tmparg0, tmpannotationsarg1)
		}
		arg0 := "[]cog.Builder[dashboardv2beta1.AnnotationQueryKind]{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	{
		buffer.WriteString(`CursorSync(`)
		arg0 := cog.Dump(input.CursorSync)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()
	}

	if input.Description != nil && *input.Description != "" {

		buffer.WriteString(`Description(`)
		arg0 := fmt.Sprintf("%#v", *input.Description)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Editable != nil && *input.Editable != true {

		buffer.WriteString(`Editable(`)
		arg0 := fmt.Sprintf("%#v", *input.Editable)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Elements != nil {

		buffer.WriteString(`Elements(`)
		arg0 := "map[string]dashboardv2beta1.Element{"
		for key, arg1 := range input.Elements {
			tmpelementsarg1 := cog.Dump(arg1)
			arg0 += "\t" + fmt.Sprintf("%#v", key) + ": " + tmpelementsarg1 + ","
		}
		arg0 += "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Layout.GridLayoutKind != nil {
		if input.Layout.GridLayoutKind != nil {
			buffer.WriteString(`GridLayout(`)
			arg0 := GridLayoutConverter(*input.Layout.GridLayoutKind)
			buffer.WriteString(arg0)

			buffer.WriteString(")")

			calls = append(calls, buffer.String())
			buffer.Reset()
		}
	}
	if input.Layout.RowsLayoutKind != nil {
		if input.Layout.RowsLayoutKind != nil {
			buffer.WriteString(`RowsLayout(`)
			arg0 := RowsLayoutConverter(*input.Layout.RowsLayoutKind)
			buffer.WriteString(arg0)

			buffer.WriteString(")")

			calls = append(calls, buffer.String())
			buffer.Reset()
		}
	}
	if input.Layout.AutoGridLayoutKind != nil {
		if input.Layout.AutoGridLayoutKind != nil {
			buffer.WriteString(`AutoGridLayout(`)
			arg0 := AutoGridLayoutConverter(*input.Layout.AutoGridLayoutKind)
			buffer.WriteString(arg0)

			buffer.WriteString(")")

			calls = append(calls, buffer.String())
			buffer.Reset()
		}
	}
	if input.Layout.TabsLayoutKind != nil {
		if input.Layout.TabsLayoutKind != nil {
			buffer.WriteString(`TabsLayout(`)
			arg0 := TabsLayoutConverter(*input.Layout.TabsLayoutKind)
			buffer.WriteString(arg0)

			buffer.WriteString(")")

			calls = append(calls, buffer.String())
			buffer.Reset()
		}
	}
	if input.Links != nil && len(input.Links) >= 1 {

		buffer.WriteString(`Links(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Links {
			tmplinksarg1 := DashboardLinkConverter(arg1)
			tmparg0 = append(tmparg0, tmplinksarg1)
		}
		arg0 := "[]cog.Builder[dashboardv2beta1.DashboardLink]{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.LiveNow != nil {

		buffer.WriteString(`LiveNow(`)
		arg0 := fmt.Sprintf("%#v", *input.LiveNow)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Preload != false {

		buffer.WriteString(`Preload(`)
		arg0 := fmt.Sprintf("%#v", input.Preload)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Revision != nil {

		buffer.WriteString(`Revision(`)
		arg0 := fmt.Sprintf("%#v", *input.Revision)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Tags != nil && len(input.Tags) >= 1 {

		buffer.WriteString(`Tags(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Tags {
			tmptagsarg1 := fmt.Sprintf("%#v", arg1)
			tmparg0 = append(tmparg0, tmptagsarg1)
		}
		arg0 := "[]string{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	{
		buffer.WriteString(`TimeSettings(`)
		arg0 := TimeSettingsConverter(input.TimeSettings)
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
	if input.Variables != nil && len(input.Variables) >= 1 {

		buffer.WriteString(`Variables(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Variables {
			tmpvariablesarg1 := cog.Dump(arg1)
			tmparg0 = append(tmparg0, tmpvariablesarg1)
		}
		arg0 := "[]dashboardv2beta1.VariableKind{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Elements != nil {
		for key, value := range input.Elements {
			if value.PanelKind != nil {
				buffer.WriteString(`Panel(`)
				arg0 := fmt.Sprintf("%#v", key)
				buffer.WriteString(arg0)
				buffer.WriteString(", ")
				arg1 := PanelConverter(*value.PanelKind)
				buffer.WriteString(arg1)

				buffer.WriteString(")")

				calls = append(calls, buffer.String())
				buffer.Reset()
			}
			if value.LibraryPanelKind != nil {
				buffer.WriteString(`LibraryPanel(`)
				arg0 := fmt.Sprintf("%#v", key)
				buffer.WriteString(arg0)
				buffer.WriteString(", ")
				arg1 := LibraryPanelConverter(*value.LibraryPanelKind)
				buffer.WriteString(arg1)

				buffer.WriteString(")")

				calls = append(calls, buffer.String())
				buffer.Reset()
			}
		}
	}
	if input.Variables != nil && len(input.Variables) >= 1 {
		for _, item := range input.Variables {
			if item.QueryVariableKind != nil {
				buffer.WriteString(`QueryVariable(`)
				arg0 := QueryVariableConverter(*item.QueryVariableKind)
				buffer.WriteString(arg0)

				buffer.WriteString(")")

				calls = append(calls, buffer.String())
				buffer.Reset()
			}
			if item.TextVariableKind != nil {
				buffer.WriteString(`TextVariable(`)
				arg0 := TextVariableConverter(*item.TextVariableKind)
				buffer.WriteString(arg0)

				buffer.WriteString(")")

				calls = append(calls, buffer.String())
				buffer.Reset()
			}
			if item.ConstantVariableKind != nil {
				buffer.WriteString(`ConstantVariable(`)
				arg0 := ConstantVariableConverter(*item.ConstantVariableKind)
				buffer.WriteString(arg0)

				buffer.WriteString(")")

				calls = append(calls, buffer.String())
				buffer.Reset()
			}
			if item.DatasourceVariableKind != nil {
				buffer.WriteString(`DatasourceVariable(`)
				arg0 := DatasourceVariableConverter(*item.DatasourceVariableKind)
				buffer.WriteString(arg0)

				buffer.WriteString(")")

				calls = append(calls, buffer.String())
				buffer.Reset()
			}
			if item.IntervalVariableKind != nil {
				buffer.WriteString(`IntervalVariable(`)
				arg0 := IntervalVariableConverter(*item.IntervalVariableKind)
				buffer.WriteString(arg0)

				buffer.WriteString(")")

				calls = append(calls, buffer.String())
				buffer.Reset()
			}
			if item.CustomVariableKind != nil {
				buffer.WriteString(`CustomVariable(`)
				arg0 := CustomVariableConverter(*item.CustomVariableKind)
				buffer.WriteString(arg0)

				buffer.WriteString(")")

				calls = append(calls, buffer.String())
				buffer.Reset()
			}
			if item.GroupByVariableKind != nil {
				buffer.WriteString(`GroupByVariable(`)
				arg0 := GroupByVariableConverter(*item.GroupByVariableKind)
				buffer.WriteString(arg0)

				buffer.WriteString(")")

				calls = append(calls, buffer.String())
				buffer.Reset()
			}
			if item.AdhocVariableKind != nil {
				buffer.WriteString(`AdhocVariable(`)
				arg0 := AdhocVariableConverter(*item.AdhocVariableKind)
				buffer.WriteString(arg0)

				buffer.WriteString(")")

				calls = append(calls, buffer.String())
				buffer.Reset()
			}
			if item.SwitchVariableKind != nil {
				buffer.WriteString(`SwitchVariableKind(`)
				arg0 := SwitchVariableKindConverter(*item.SwitchVariableKind)
				buffer.WriteString(arg0)

				buffer.WriteString(")")

				calls = append(calls, buffer.String())
				buffer.Reset()
			}
		}
	}

	return strings.Join(calls, ".\t\n")
}
