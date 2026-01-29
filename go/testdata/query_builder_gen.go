// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package testdata

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
	dashboardv2beta1 "github.com/grafana/grafana-foundation-sdk/go/dashboardv2beta1"
)

var _ cog.Builder[dashboardv2beta1.DataQueryKind] = (*QueryBuilder)(nil)

type QueryBuilder struct {
	internal *dashboardv2beta1.DataQueryKind
	errors   cog.BuildErrors
}

func NewQueryBuilder() *QueryBuilder {
	resource := dashboardv2beta1.NewDataQueryKind()
	builder := &QueryBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Kind = "DataQuery"
	builder.internal.Group = "testdata"

	return builder
}

func (builder *QueryBuilder) Build() (dashboardv2beta1.DataQueryKind, error) {
	if err := builder.internal.Validate(); err != nil {
		return dashboardv2beta1.DataQueryKind{}, err
	}

	if len(builder.errors) > 0 {
		return dashboardv2beta1.DataQueryKind{}, cog.MakeBuildErrors("testdata.query", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *QueryBuilder) Version(version string) *QueryBuilder {
	builder.internal.Version = version

	return builder
}

// New type for datasource reference
// Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.
func (builder *QueryBuilder) Datasource(datasource cog.Builder[dashboardv2beta1.Dashboardv2beta1DataQueryKindDatasource]) *QueryBuilder {
	datasourceResource, err := datasource.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Datasource = &datasourceResource

	return builder
}

func (builder *QueryBuilder) Alias(alias string) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).Alias = &alias

	return builder
}

// Used for live query
func (builder *QueryBuilder) Channel(channel string) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).Channel = &channel

	return builder
}

func (builder *QueryBuilder) CsvContent(csvContent string) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).CsvContent = &csvContent

	return builder
}

func (builder *QueryBuilder) CsvFileName(csvFileName string) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).CsvFileName = &csvFileName

	return builder
}

func (builder *QueryBuilder) CsvWave(csvWave []cog.Builder[CSVWave]) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	csvWaveResources := make([]CSVWave, 0, len(csvWave))
	for _, r1 := range csvWave {
		csvWaveDepth1, err := r1.Build()
		if err != nil {
			builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
			return builder
		}
		csvWaveResources = append(csvWaveResources, csvWaveDepth1)
	}
	builder.internal.Spec.(*Dataquery).CsvWave = csvWaveResources

	return builder
}

// The datasource
func (builder *QueryBuilder) OldDatasource(datasource common.DataSourceRef) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).Datasource = &datasource

	return builder
}

// Drop percentage (the chance we will lose a point 0-100)
func (builder *QueryBuilder) DropPercent(dropPercent float64) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).DropPercent = &dropPercent

	return builder
}

// Possible enum values:
//   - `"plugin"`
//   - `"downstream"`
func (builder *QueryBuilder) ErrorSource(errorSource DataqueryErrorSource) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).ErrorSource = &errorSource

	return builder
}

// Possible enum values:
//   - `"frontend_exception"`
//   - `"frontend_observable"`
//   - `"server_panic"`
func (builder *QueryBuilder) ErrorType(errorType DataqueryErrorType) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).ErrorType = &errorType

	return builder
}

func (builder *QueryBuilder) FlamegraphDiff(flamegraphDiff bool) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).FlamegraphDiff = &flamegraphDiff

	return builder
}

// true if query is disabled (ie should not be returned to the dashboard)
// NOTE: this does not always imply that the query should not be executed since
// the results from a hidden query may be used as the input to other queries (SSE etc)
func (builder *QueryBuilder) Hide(hide bool) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).Hide = &hide

	return builder
}

// Interval is the suggested duration between time points in a time series query.
// NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
// from the interval required to fill a pixels in the visualization
func (builder *QueryBuilder) IntervalMs(intervalMs float64) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).IntervalMs = &intervalMs

	return builder
}

func (builder *QueryBuilder) Labels(labels string) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).Labels = &labels

	return builder
}

func (builder *QueryBuilder) LevelColumn(levelColumn bool) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).LevelColumn = &levelColumn

	return builder
}

func (builder *QueryBuilder) Lines(lines int64) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).Lines = &lines

	return builder
}

func (builder *QueryBuilder) Max(max float64) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).Max = &max

	return builder
}

// MaxDataPoints is the maximum number of data points that should be returned from a time series query.
// NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
// from the number of pixels visible in a visualization
func (builder *QueryBuilder) MaxDataPoints(maxDataPoints int64) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).MaxDataPoints = &maxDataPoints

	return builder
}

func (builder *QueryBuilder) Min(min float64) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).Min = &min

	return builder
}

func (builder *QueryBuilder) Nodes(nodes cog.Builder[NodesQuery]) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	nodesResource, err := nodes.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.(*Dataquery).Nodes = &nodesResource

	return builder
}

func (builder *QueryBuilder) Noise(noise float64) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).Noise = &noise

	return builder
}

func (builder *QueryBuilder) Points(points [][]any) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).Points = points

	return builder
}

func (builder *QueryBuilder) PulseWave(pulseWave cog.Builder[PulseWaveQuery]) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	pulseWaveResource, err := pulseWave.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.(*Dataquery).PulseWave = &pulseWaveResource

	return builder
}

// QueryType is an optional identifier for the type of query.
// It can be used to distinguish different types of queries.
func (builder *QueryBuilder) QueryType(queryType string) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).QueryType = &queryType

	return builder
}

func (builder *QueryBuilder) RawFrameContent(rawFrameContent string) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).RawFrameContent = &rawFrameContent

	return builder
}

// RefID is the unique identifier of the query, set by the frontend call.
func (builder *QueryBuilder) RefId(refId string) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).RefId = &refId

	return builder
}

// Optionally define expected query result behavior
func (builder *QueryBuilder) ResultAssertions(resultAssertions cog.Builder[ResultAssertions]) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	resultAssertionsResource, err := resultAssertions.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.(*Dataquery).ResultAssertions = &resultAssertionsResource

	return builder
}

// Possible enum values:
//   - `"annotations"`
//   - `"arrow"`
//   - `"csv_content"`
//   - `"csv_file"`
//   - `"csv_metric_values"`
//   - `"datapoints_outside_range"`
//   - `"error_with_source"`
//   - `"exponential_heatmap_bucket_data"`
//   - `"flame_graph"`
//   - `"grafana_api"`
//   - `"linear_heatmap_bucket_data"`
//   - `"live"`
//   - `"logs"`
//   - `"manual_entry"`
//   - `"no_data_points"`
//   - `"node_graph"`
//   - `"predictable_csv_wave"`
//   - `"predictable_pulse"`
//   - `"query_meta"`
//   - `"random_walk"`
//   - `"random_walk_table"`
//   - `"random_walk_with_error"`
//   - `"raw_frame"`
//   - `"server_error_500"`
//   - `"steps"`
//   - `"simulation"`
//   - `"slow_query"`
//   - `"streaming_client"`
//   - `"table_static"`
//   - `"trace"`
//   - `"usa"`
//   - `"variables-query"`
func (builder *QueryBuilder) ScenarioId(scenarioId DataqueryScenarioId) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).ScenarioId = &scenarioId

	return builder
}

func (builder *QueryBuilder) SeriesCount(seriesCount int64) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).SeriesCount = &seriesCount

	return builder
}

func (builder *QueryBuilder) Sim(sim cog.Builder[SimulationQuery]) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	simResource, err := sim.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.(*Dataquery).Sim = &simResource

	return builder
}

func (builder *QueryBuilder) SpanCount(spanCount int64) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).SpanCount = &spanCount

	return builder
}

func (builder *QueryBuilder) Spread(spread float64) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).Spread = &spread

	return builder
}

func (builder *QueryBuilder) StartValue(startValue float64) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).StartValue = &startValue

	return builder
}

func (builder *QueryBuilder) Stream(stream cog.Builder[StreamingQuery]) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	streamResource, err := stream.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.(*Dataquery).Stream = &streamResource

	return builder
}

// common parameter used by many query types
func (builder *QueryBuilder) StringInput(stringInput string) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).StringInput = &stringInput

	return builder
}

// TimeRange represents the query range
// NOTE: unlike generic /ds/query, we can now send explicit time values in each query
// NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
func (builder *QueryBuilder) TimeRange(timeRange cog.Builder[TimeRange]) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	timeRangeResource, err := timeRange.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.(*Dataquery).TimeRange = &timeRangeResource

	return builder
}

func (builder *QueryBuilder) Usa(usa cog.Builder[USAQuery]) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	usaResource, err := usa.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.(*Dataquery).Usa = &usaResource

	return builder
}

func (builder *QueryBuilder) WithNil(withNil bool) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).WithNil = &withNil

	return builder
}
