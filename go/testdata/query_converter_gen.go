// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package testdata

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	dashboardv2beta1 "github.com/grafana/grafana-foundation-sdk/go/dashboardv2beta1"
)

// QueryConverter accepts a `Query` object and generates the Go code to build this object using builders.
func QueryConverter(input dashboardv2beta1.DataQueryKind) string {
	calls := []string{
		`testdata.NewQueryBuilder()`,
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
	if input.Spec != nil && input.Spec.(*Dataquery).Alias != nil && *input.Spec.(*Dataquery).Alias != "" {

		buffer.WriteString(`Alias(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*Dataquery).Alias)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).Channel != nil && *input.Spec.(*Dataquery).Channel != "" {

		buffer.WriteString(`Channel(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*Dataquery).Channel)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).CsvContent != nil && *input.Spec.(*Dataquery).CsvContent != "" {

		buffer.WriteString(`CsvContent(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*Dataquery).CsvContent)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).CsvFileName != nil && *input.Spec.(*Dataquery).CsvFileName != "" {

		buffer.WriteString(`CsvFileName(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*Dataquery).CsvFileName)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).CsvWave != nil && len(input.Spec.(*Dataquery).CsvWave) >= 1 {

		buffer.WriteString(`CsvWave(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Spec.(*Dataquery).CsvWave {
			tmpcsvWavearg1 := CSVWaveConverter(arg1)
			tmparg0 = append(tmparg0, tmpcsvWavearg1)
		}
		arg0 := "[]cog.Builder[testdata.CSVWave]{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).Datasource != nil {

		buffer.WriteString(`OldDatasource(`)
		arg0 := cog.Dump(*input.Spec.(*Dataquery).Datasource)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).DropPercent != nil {

		buffer.WriteString(`DropPercent(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*Dataquery).DropPercent)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).ErrorSource != nil {

		buffer.WriteString(`ErrorSource(`)
		arg0 := cog.Dump(*input.Spec.(*Dataquery).ErrorSource)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).ErrorType != nil {

		buffer.WriteString(`ErrorType(`)
		arg0 := cog.Dump(*input.Spec.(*Dataquery).ErrorType)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).FlamegraphDiff != nil {

		buffer.WriteString(`FlamegraphDiff(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*Dataquery).FlamegraphDiff)
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
	if input.Spec != nil && input.Spec.(*Dataquery).IntervalMs != nil {

		buffer.WriteString(`IntervalMs(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*Dataquery).IntervalMs)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).Labels != nil && *input.Spec.(*Dataquery).Labels != "" {

		buffer.WriteString(`Labels(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*Dataquery).Labels)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).LevelColumn != nil {

		buffer.WriteString(`LevelColumn(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*Dataquery).LevelColumn)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).Lines != nil {

		buffer.WriteString(`Lines(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*Dataquery).Lines)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).Max != nil {

		buffer.WriteString(`Max(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*Dataquery).Max)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).MaxDataPoints != nil {

		buffer.WriteString(`MaxDataPoints(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*Dataquery).MaxDataPoints)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).Min != nil {

		buffer.WriteString(`Min(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*Dataquery).Min)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).Nodes != nil {

		buffer.WriteString(`Nodes(`)
		arg0 := NodesQueryConverter(*input.Spec.(*Dataquery).Nodes)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).Noise != nil {

		buffer.WriteString(`Noise(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*Dataquery).Noise)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).Points != nil && len(input.Spec.(*Dataquery).Points) >= 1 {

		buffer.WriteString(`Points(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Spec.(*Dataquery).Points {
			tmptmppointsarg1 := []string{}
			for _, arg1Value := range arg1 {
				tmparg1arg1Value := cog.Dump(arg1Value)
				tmptmppointsarg1 = append(tmptmppointsarg1, tmparg1arg1Value)
			}
			tmppointsarg1 := "[]any{" + strings.Join(tmptmppointsarg1, ",\n") + "}"
			tmparg0 = append(tmparg0, tmppointsarg1)
		}
		arg0 := "[][]any{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).PulseWave != nil {

		buffer.WriteString(`PulseWave(`)
		arg0 := PulseWaveQueryConverter(*input.Spec.(*Dataquery).PulseWave)
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
	if input.Spec != nil && input.Spec.(*Dataquery).RawFrameContent != nil && *input.Spec.(*Dataquery).RawFrameContent != "" {

		buffer.WriteString(`RawFrameContent(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*Dataquery).RawFrameContent)
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
	if input.Spec != nil && input.Spec.(*Dataquery).ResultAssertions != nil {

		buffer.WriteString(`ResultAssertions(`)
		arg0 := ResultAssertionsConverter(*input.Spec.(*Dataquery).ResultAssertions)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).ScenarioId != nil {

		buffer.WriteString(`ScenarioId(`)
		arg0 := cog.Dump(*input.Spec.(*Dataquery).ScenarioId)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).SeriesCount != nil {

		buffer.WriteString(`SeriesCount(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*Dataquery).SeriesCount)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).Sim != nil {

		buffer.WriteString(`Sim(`)
		arg0 := SimulationQueryConverter(*input.Spec.(*Dataquery).Sim)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).SpanCount != nil {

		buffer.WriteString(`SpanCount(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*Dataquery).SpanCount)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).Spread != nil {

		buffer.WriteString(`Spread(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*Dataquery).Spread)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).StartValue != nil {

		buffer.WriteString(`StartValue(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*Dataquery).StartValue)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).Stream != nil {

		buffer.WriteString(`Stream(`)
		arg0 := StreamingQueryConverter(*input.Spec.(*Dataquery).Stream)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).StringInput != nil && *input.Spec.(*Dataquery).StringInput != "" {

		buffer.WriteString(`StringInput(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*Dataquery).StringInput)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).TimeRange != nil {

		buffer.WriteString(`TimeRange(`)
		arg0 := TimeRangeConverter(*input.Spec.(*Dataquery).TimeRange)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).Usa != nil {

		buffer.WriteString(`Usa(`)
		arg0 := USAQueryConverter(*input.Spec.(*Dataquery).Usa)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spec != nil && input.Spec.(*Dataquery).WithNil != nil {

		buffer.WriteString(`WithNil(`)
		arg0 := fmt.Sprintf("%#v", *input.Spec.(*Dataquery).WithNil)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
