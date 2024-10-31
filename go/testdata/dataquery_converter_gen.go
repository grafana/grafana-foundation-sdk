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
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Alias))
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
	if input.StringInput != nil && *input.StringInput != "" {

		buffer.WriteString(`StringInput(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.StringInput))
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
	if input.PulseWave != nil {

		buffer.WriteString(`PulseWave(`)
		arg0 := PulseWaveQueryConverter(*input.PulseWave)
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
	if input.Labels != nil && *input.Labels != "" {

		buffer.WriteString(`Labels(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Labels))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Lines != nil {

		buffer.WriteString(`Lines(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Lines))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.LevelColumn != nil {

		buffer.WriteString(`LevelColumn(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.LevelColumn))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Channel != nil && *input.Channel != "" {

		buffer.WriteString(`Channel(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Channel))
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
	if input.CsvFileName != nil && *input.CsvFileName != "" {

		buffer.WriteString(`CsvFileName(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.CsvFileName))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.CsvContent != nil && *input.CsvContent != "" {

		buffer.WriteString(`CsvContent(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.CsvContent))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.RawFrameContent != nil && *input.RawFrameContent != "" {

		buffer.WriteString(`RawFrameContent(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.RawFrameContent))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.SeriesCount != nil {

		buffer.WriteString(`SeriesCount(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.SeriesCount))
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
	if input.ErrorType != nil {

		buffer.WriteString(`ErrorType(`)
		arg0 := cog.Dump(*input.ErrorType)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.SpanCount != nil {

		buffer.WriteString(`SpanCount(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.SpanCount))
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
			tmppointsarg1 := "[]testdata.StringOrInt64{" + strings.Join(tmptmppointsarg1, ",\n") + "}"
			tmparg0 = append(tmparg0, tmppointsarg1)
		}
		arg0 := "[][]testdata.StringOrInt64{" + strings.Join(tmparg0, ",\n") + "}"
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.DropPercent != nil {

		buffer.WriteString(`DropPercent(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.DropPercent))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.RefId != "" {

		buffer.WriteString(`RefId(`)
		arg0 := fmt.Sprintf("%#v", input.RefId)
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.Hide != nil {

		buffer.WriteString(`Hide(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.Hide))
		buffer.WriteString(arg0)

		buffer.WriteString(")")

		calls = append(calls, buffer.String())
		buffer.Reset()

	}
	if input.QueryType != nil && *input.QueryType != "" {

		buffer.WriteString(`QueryType(`)
		arg0 := fmt.Sprintf("%#v", cog.Unptr(input.QueryType))
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

	return strings.Join(calls, ".\t\n")
}
