// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package athena

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	dashboardv2beta1 "github.com/grafana/grafana-foundation-sdk/go/dashboardv2beta1"
)

// QueryConverter accepts a `Query` object and generates the Go code to build this object using builders.
func QueryConverter(input dashboardv2beta1.DataQueryKind) string {
	calls := []string{
		`athena.NewQueryBuilder()`,
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
	if input.Spec != nil {

		buffer.WriteString(`Format(`)
		arg0 := cog.Dump(input.Spec.(*Dataquery).Format)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil {

		buffer.WriteString(`ConnectionArgs(`)
		arg0 := ConnectionArgsConverter(input.Spec.(*Dataquery).ConnectionArgs)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).Table != nil && *input.Spec.(*Dataquery).Table != "" {

		buffer.WriteString(`Table(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*Dataquery).Table)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).Column != nil && *input.Spec.(*Dataquery).Column != "" {

		buffer.WriteString(`Column(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*Dataquery).Column)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).QueryID != nil && *input.Spec.(*Dataquery).QueryID != "" {

		buffer.WriteString(`QueryID(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*Dataquery).QueryID)
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
	if input.Spec != nil && input.Spec.(*Dataquery).RawSQL != "" {

		buffer.WriteString(`RawSQL(`)
		arg0 := fmt.Sprintf("%#v", input.Spec.(*Dataquery).RawSQL)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
