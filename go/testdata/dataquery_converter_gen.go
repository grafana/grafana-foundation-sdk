// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package testdata

import (
	"fmt"
	"strings"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// DataqueryConverter accepts a `Dataquery` object and generates the Go code to build this object using builders.
func DataqueryConverter(input Dataquery) string {
	calls := []string{
		`testdata.NewDataqueryBuilder()`,
	}
	var buffer strings.Builder
	if input.Alias != nil && *input.Alias != "" {

		buffer.WriteString(`Alias(`)
		arg0 := fmt.Sprintf("%#v", *input.Alias)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Channel != nil && *input.Channel != "" {

		buffer.WriteString(`Channel(`)
		arg0 := fmt.Sprintf("%#v", *input.Channel)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.CsvContent != nil && *input.CsvContent != "" {

		buffer.WriteString(`CsvContent(`)
		arg0 := fmt.Sprintf("%#v", *input.CsvContent)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.CsvFileName != nil && *input.CsvFileName != "" {

		buffer.WriteString(`CsvFileName(`)
		arg0 := fmt.Sprintf("%#v", *input.CsvFileName)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.CsvWave != nil && len(input.CsvWave) >= 1 {

		buffer.WriteString(`CsvWave(`)
		tmparg0 := []string{}
		for _, arg1 := range input.CsvWave {
			tmpcsvWavearg1 := CSVWaveConverter(arg1)
			tmparg0 = append(tmparg0, tmpcsvWavearg1)
		}
		arg0 := "[]cog.Builder[testdata.CSVWave]{" + strings.Join(tmparg0, ",\n") + "}"
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
	if input.DropPercent != nil {

		buffer.WriteString(`DropPercent(`)
		arg0 := fmt.Sprintf("%#v", *input.DropPercent)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.ErrorSource != nil {

		buffer.WriteString(`ErrorSource(`)
		arg0 := cog.Dump(*input.ErrorSource)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.ErrorType != nil {

		buffer.WriteString(`ErrorType(`)
		arg0 := cog.Dump(*input.ErrorType)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.FlamegraphDiff != nil {

		buffer.WriteString(`FlamegraphDiff(`)
		arg0 := fmt.Sprintf("%#v", *input.FlamegraphDiff)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Hide != nil {

		buffer.WriteString(`Hide(`)
		arg0 := fmt.Sprintf("%#v", *input.Hide)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.IntervalMs != nil {

		buffer.WriteString(`IntervalMs(`)
		arg0 := fmt.Sprintf("%#v", *input.IntervalMs)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Labels != nil && *input.Labels != "" {

		buffer.WriteString(`Labels(`)
		arg0 := fmt.Sprintf("%#v", *input.Labels)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.LevelColumn != nil {

		buffer.WriteString(`LevelColumn(`)
		arg0 := fmt.Sprintf("%#v", *input.LevelColumn)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Lines != nil {

		buffer.WriteString(`Lines(`)
		arg0 := fmt.Sprintf("%#v", *input.Lines)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Max != nil {

		buffer.WriteString(`Max(`)
		arg0 := fmt.Sprintf("%#v", *input.Max)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.MaxDataPoints != nil {

		buffer.WriteString(`MaxDataPoints(`)
		arg0 := fmt.Sprintf("%#v", *input.MaxDataPoints)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Min != nil {

		buffer.WriteString(`Min(`)
		arg0 := fmt.Sprintf("%#v", *input.Min)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Nodes != nil {

		buffer.WriteString(`Nodes(`)
		arg0 := NodesQueryConverter(*input.Nodes)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Noise != nil {

		buffer.WriteString(`Noise(`)
		arg0 := fmt.Sprintf("%#v", *input.Noise)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Points != nil && len(input.Points) >= 1 {

		buffer.WriteString(`Points(`)
		tmparg0 := []string{}
		for _, arg1 := range input.Points {
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
	if input.PulseWave != nil {

		buffer.WriteString(`PulseWave(`)
		arg0 := PulseWaveQueryConverter(*input.PulseWave)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.QueryType != nil && *input.QueryType != "" {

		buffer.WriteString(`QueryType(`)
		arg0 := fmt.Sprintf("%#v", *input.QueryType)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.RawFrameContent != nil && *input.RawFrameContent != "" {

		buffer.WriteString(`RawFrameContent(`)
		arg0 := fmt.Sprintf("%#v", *input.RawFrameContent)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.RefId != nil && *input.RefId != "" {

		buffer.WriteString(`RefId(`)
		arg0 := fmt.Sprintf("%#v", *input.RefId)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.ResultAssertions != nil {

		buffer.WriteString(`ResultAssertions(`)
		arg0 := ResultAssertionsConverter(*input.ResultAssertions)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.ScenarioId != nil {

		buffer.WriteString(`ScenarioId(`)
		arg0 := cog.Dump(*input.ScenarioId)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.SeriesCount != nil {

		buffer.WriteString(`SeriesCount(`)
		arg0 := fmt.Sprintf("%#v", *input.SeriesCount)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Sim != nil {

		buffer.WriteString(`Sim(`)
		arg0 := SimulationQueryConverter(*input.Sim)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.SpanCount != nil {

		buffer.WriteString(`SpanCount(`)
		arg0 := fmt.Sprintf("%#v", *input.SpanCount)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Spread != nil {

		buffer.WriteString(`Spread(`)
		arg0 := fmt.Sprintf("%#v", *input.Spread)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.StartValue != nil {

		buffer.WriteString(`StartValue(`)
		arg0 := fmt.Sprintf("%#v", *input.StartValue)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Stream != nil {

		buffer.WriteString(`Stream(`)
		arg0 := StreamingQueryConverter(*input.Stream)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.StringInput != nil && *input.StringInput != "" {

		buffer.WriteString(`StringInput(`)
		arg0 := fmt.Sprintf("%#v", *input.StringInput)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.TimeRange != nil {

		buffer.WriteString(`TimeRange(`)
		arg0 := TimeRangeConverter(*input.TimeRange)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Usa != nil {

		buffer.WriteString(`Usa(`)
		arg0 := USAQueryConverter(*input.Usa)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.WithNil != nil {

		buffer.WriteString(`WithNil(`)
		arg0 := fmt.Sprintf("%#v", *input.WithNil)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}

	return strings.Join(calls, ".\t\n")
}
