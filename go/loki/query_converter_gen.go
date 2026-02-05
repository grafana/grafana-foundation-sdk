// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package loki

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	dashboardv2beta1 "github.com/grafana/grafana-foundation-sdk/go/dashboardv2beta1"
)

// QueryConverter accepts a `Query` object and generates the Go code to build this object using builders.
func QueryConverter(input dashboardv2beta1.DataQueryKind) string {
	calls := []string{
		`loki.NewQueryBuilder()`,
	}
	var buffer strings.Builder
	if input.Version != "" && input.Version != "v0" {

		buffer.WriteString(`Version(`)
		arg0 := fmt.Sprintf("%#v", input.Version)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Datasource != nil {

		buffer.WriteString(`Datasource(`)
		arg0 := dashboardv2beta1.Dashboardv2beta1DataQueryKindDatasourceConverter(*input.Datasource)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).Expr != "" {

		buffer.WriteString(`Expr(`)
		arg0 := fmt.Sprintf("%#v", input.Spec.(*Dataquery).Expr)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).LegendFormat != nil && *input.Spec.(*Dataquery).LegendFormat != "" {

		buffer.WriteString(`LegendFormat(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*Dataquery).LegendFormat)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).MaxLines != nil {

		buffer.WriteString(`MaxLines(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*Dataquery).MaxLines)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).Resolution != nil {

		buffer.WriteString(`Resolution(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*Dataquery).Resolution)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).EditorMode != nil {

		buffer.WriteString(`EditorMode(`)
		arg0 := cog.Dump(*input.Spec.(*Dataquery).EditorMode)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).Range != nil {

		buffer.WriteString(`Range(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*Dataquery).Range)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).Instant != nil {

		buffer.WriteString(`Instant(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*Dataquery).Instant)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).Step != nil && *input.Spec.(*Dataquery).Step != "" {

		buffer.WriteString(`Step(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*Dataquery).Step)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).RefId != nil && *input.Spec.(*Dataquery).RefId != "" {

		buffer.WriteString(`RefId(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*Dataquery).RefId)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).Hide != nil {

		buffer.WriteString(`Hide(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*Dataquery).Hide)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).QueryType != nil && *input.Spec.(*Dataquery).QueryType != "" {

		buffer.WriteString(`QueryType(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*Dataquery).QueryType)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).Direction != nil {

		buffer.WriteString(`Direction(`)
		arg0 := cog.Dump(*input.Spec.(*Dataquery).Direction)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
