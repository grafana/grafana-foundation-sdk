// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package testdata

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	dashboard "github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

var _ cog.Builder[variants.Dataquery] = (*DataqueryBuilder)(nil)

type DataqueryBuilder struct {
	internal *Dataquery
	errors   map[string]cog.BuildErrors
}

func NewDataqueryBuilder() *DataqueryBuilder {
	resource := &Dataquery{}
	builder := &DataqueryBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *DataqueryBuilder) Build() (variants.Dataquery, error) {
	if err := builder.internal.Validate(); err != nil {
		return Dataquery{}, err
	}

	return *builder.internal, nil
}

func (builder *DataqueryBuilder) Alias(alias string) *DataqueryBuilder {
	builder.internal.Alias = &alias

	return builder
}

func (builder *DataqueryBuilder) ScenarioId(scenarioId TestDataQueryType) *DataqueryBuilder {
	builder.internal.ScenarioId = &scenarioId

	return builder
}

func (builder *DataqueryBuilder) StringInput(stringInput string) *DataqueryBuilder {
	builder.internal.StringInput = &stringInput

	return builder
}

func (builder *DataqueryBuilder) Stream(stream cog.Builder[StreamingQuery]) *DataqueryBuilder {
	streamResource, err := stream.Build()
	if err != nil {
		builder.errors["stream"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Stream = &streamResource

	return builder
}

func (builder *DataqueryBuilder) PulseWave(pulseWave cog.Builder[PulseWaveQuery]) *DataqueryBuilder {
	pulseWaveResource, err := pulseWave.Build()
	if err != nil {
		builder.errors["pulseWave"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.PulseWave = &pulseWaveResource

	return builder
}

func (builder *DataqueryBuilder) Sim(sim cog.Builder[SimulationQuery]) *DataqueryBuilder {
	simResource, err := sim.Build()
	if err != nil {
		builder.errors["sim"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Sim = &simResource

	return builder
}

func (builder *DataqueryBuilder) CsvWave(csvWave []cog.Builder[CSVWave]) *DataqueryBuilder {
	csvWaveResources := make([]CSVWave, 0, len(csvWave))
	for _, r1 := range csvWave {
		csvWaveDepth1, err := r1.Build()
		if err != nil {
			builder.errors["csvWave"] = err.(cog.BuildErrors)
			return builder
		}
		csvWaveResources = append(csvWaveResources, csvWaveDepth1)
	}
	builder.internal.CsvWave = csvWaveResources

	return builder
}

func (builder *DataqueryBuilder) Labels(labels string) *DataqueryBuilder {
	builder.internal.Labels = &labels

	return builder
}

func (builder *DataqueryBuilder) Lines(lines int64) *DataqueryBuilder {
	builder.internal.Lines = &lines

	return builder
}

func (builder *DataqueryBuilder) LevelColumn(levelColumn bool) *DataqueryBuilder {
	builder.internal.LevelColumn = &levelColumn

	return builder
}

func (builder *DataqueryBuilder) Channel(channel string) *DataqueryBuilder {
	builder.internal.Channel = &channel

	return builder
}

func (builder *DataqueryBuilder) Nodes(nodes cog.Builder[NodesQuery]) *DataqueryBuilder {
	nodesResource, err := nodes.Build()
	if err != nil {
		builder.errors["nodes"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Nodes = &nodesResource

	return builder
}

func (builder *DataqueryBuilder) CsvFileName(csvFileName string) *DataqueryBuilder {
	builder.internal.CsvFileName = &csvFileName

	return builder
}

func (builder *DataqueryBuilder) CsvContent(csvContent string) *DataqueryBuilder {
	builder.internal.CsvContent = &csvContent

	return builder
}

func (builder *DataqueryBuilder) RawFrameContent(rawFrameContent string) *DataqueryBuilder {
	builder.internal.RawFrameContent = &rawFrameContent

	return builder
}

func (builder *DataqueryBuilder) SeriesCount(seriesCount int32) *DataqueryBuilder {
	builder.internal.SeriesCount = &seriesCount

	return builder
}

func (builder *DataqueryBuilder) Usa(usa cog.Builder[USAQuery]) *DataqueryBuilder {
	usaResource, err := usa.Build()
	if err != nil {
		builder.errors["usa"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Usa = &usaResource

	return builder
}

func (builder *DataqueryBuilder) ErrorType(errorType DataqueryErrorType) *DataqueryBuilder {
	builder.internal.ErrorType = &errorType

	return builder
}

func (builder *DataqueryBuilder) SpanCount(spanCount int32) *DataqueryBuilder {
	builder.internal.SpanCount = &spanCount

	return builder
}

func (builder *DataqueryBuilder) Points(points [][]StringOrInt64) *DataqueryBuilder {
	builder.internal.Points = points

	return builder
}

// Drop percentage (the chance we will lose a point 0-100)
func (builder *DataqueryBuilder) DropPercent(dropPercent float64) *DataqueryBuilder {
	builder.internal.DropPercent = &dropPercent

	return builder
}

// A unique identifier for the query within the list of targets.
// In server side expressions, the refId is used as a variable name to identify results.
// By default, the UI will assign A->Z; however setting meaningful names may be useful.
func (builder *DataqueryBuilder) RefId(refId string) *DataqueryBuilder {
	builder.internal.RefId = refId

	return builder
}

// true if query is disabled (ie should not be returned to the dashboard)
// Note this does not always imply that the query should not be executed since
// the results from a hidden query may be used as the input to other queries (SSE etc)
func (builder *DataqueryBuilder) Hide(hide bool) *DataqueryBuilder {
	builder.internal.Hide = &hide

	return builder
}

// Specify the query flavor
// TODO make this required and give it a default
func (builder *DataqueryBuilder) QueryType(queryType string) *DataqueryBuilder {
	builder.internal.QueryType = &queryType

	return builder
}

// For mixed data sources the selected datasource is on the query level.
// For non mixed scenarios this is undefined.
// TODO find a better way to do this ^ that's friendly to schema
// TODO this shouldn't be unknown but DataSourceRef | null
func (builder *DataqueryBuilder) Datasource(datasource dashboard.DataSourceRef) *DataqueryBuilder {
	builder.internal.Datasource = &datasource

	return builder
}

func (builder *DataqueryBuilder) applyDefaults() {
	builder.ScenarioId("random_walk")
}
