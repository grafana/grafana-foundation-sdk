// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package testdata

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	dashboardv2 "github.com/grafana/grafana-foundation-sdk/go/dashboardv2"
)

var _ cog.Builder[dashboardv2.DataQueryKind] = (*QueryV2Builder)(nil)

type QueryV2Builder struct {
	internal *dashboardv2.DataQueryKind
	errors   cog.BuildErrors
}

func NewQueryV2Builder() *QueryV2Builder {
	resource := dashboardv2.NewDataQueryKind()
	builder := &QueryV2Builder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Kind = "DataQuery"
	builder.internal.Group = "testdata"

	return builder
}

func (builder *QueryV2Builder) Build() (dashboardv2.DataQueryKind, error) {
	if err := builder.internal.Validate(); err != nil {
		return dashboardv2.DataQueryKind{}, err
	}

	if len(builder.errors) > 0 {
		return dashboardv2.DataQueryKind{}, cog.MakeBuildErrors("testdata.queryV2", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *QueryV2Builder) RecordError(path string, err error) *QueryV2Builder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *QueryV2Builder) Version(version string) *QueryV2Builder {
	builder.internal.Version = version

	return builder
}

func (builder *QueryV2Builder) Labels(labels map[string]string) *QueryV2Builder {
	builder.internal.Labels = labels

	return builder
}

// New type for datasource reference
// Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.
func (builder *QueryV2Builder) Datasource(datasource cog.Builder[dashboardv2.Dashboardv2DataQueryKindDatasource]) *QueryV2Builder {
	datasourceResource, err := datasource.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Datasource = &datasourceResource

	return builder
}

func (builder *QueryV2Builder) Alias(alias string) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).Alias = &alias

	return builder
}

// Used for live query
func (builder *QueryV2Builder) Channel(channel string) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).Channel = &channel

	return builder
}

func (builder *QueryV2Builder) CsvContent(csvContent string) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).CsvContent = &csvContent

	return builder
}

func (builder *QueryV2Builder) CsvFileName(csvFileName string) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).CsvFileName = &csvFileName

	return builder
}

func (builder *QueryV2Builder) CsvWave(csvWave []cog.Builder[CSVWave]) *QueryV2Builder {
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

// Drop percentage (the chance we will lose a point 0-100)
func (builder *QueryV2Builder) DropPercent(dropPercent float64) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).DropPercent = &dropPercent

	return builder
}

// Possible enum values:
//   - `"plugin"`
//   - `"downstream"`
func (builder *QueryV2Builder) ErrorSource(errorSource DataqueryErrorSource) *QueryV2Builder {
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
func (builder *QueryV2Builder) ErrorType(errorType DataqueryErrorType) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).ErrorType = &errorType

	return builder
}

func (builder *QueryV2Builder) FlamegraphDiff(flamegraphDiff bool) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).FlamegraphDiff = &flamegraphDiff

	return builder
}

// true if query is disabled (ie should not be returned to the dashboard)
// NOTE: this does not always imply that the query should not be executed since
// the results from a hidden query may be used as the input to other queries (SSE etc)
func (builder *QueryV2Builder) Hide(hide bool) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).Hide = &hide

	return builder
}

// Interval is the suggested duration between time points in a time series query.
// NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
// from the interval required to fill a pixels in the visualization
func (builder *QueryV2Builder) IntervalMs(intervalMs float64) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).IntervalMs = &intervalMs

	return builder
}

func (builder *QueryV2Builder) DataqueryLabels(labels string) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).Labels = &labels

	return builder
}

func (builder *QueryV2Builder) LevelColumn(levelColumn bool) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).LevelColumn = &levelColumn

	return builder
}

func (builder *QueryV2Builder) Lines(lines int64) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).Lines = &lines

	return builder
}

func (builder *QueryV2Builder) Max(max float64) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).Max = &max

	return builder
}

// MaxDataPoints is the maximum number of data points that should be returned from a time series query.
// NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
// from the number of pixels visible in a visualization
func (builder *QueryV2Builder) MaxDataPoints(maxDataPoints int64) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).MaxDataPoints = &maxDataPoints

	return builder
}

func (builder *QueryV2Builder) Min(min float64) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).Min = &min

	return builder
}

func (builder *QueryV2Builder) Nodes(nodes cog.Builder[NodesQuery]) *QueryV2Builder {
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

func (builder *QueryV2Builder) Noise(noise float64) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).Noise = &noise

	return builder
}

func (builder *QueryV2Builder) Points(points [][]any) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).Points = points

	return builder
}

func (builder *QueryV2Builder) PulseWave(pulseWave cog.Builder[PulseWaveQuery]) *QueryV2Builder {
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
func (builder *QueryV2Builder) QueryType(queryType string) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).QueryType = &queryType

	return builder
}

func (builder *QueryV2Builder) RawFrameContent(rawFrameContent string) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).RawFrameContent = &rawFrameContent

	return builder
}

// RefID is the unique identifier of the query, set by the frontend call.
func (builder *QueryV2Builder) RefId(refId string) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).RefId = &refId

	return builder
}

// Optionally define expected query result behavior
func (builder *QueryV2Builder) ResultAssertions(resultAssertions cog.Builder[ResultAssertions]) *QueryV2Builder {
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
//   - `"random_walk"`
//   - `"random_walk_table"`
//   - `"random_walk_with_error"`
//   - `"raw_frame"`
//   - `"server_error_500"`
//   - `"simulation"`
//   - `"slow_query"`
//   - `"streaming_client"`
//   - `"table_static"`
//   - `"trace"`
//   - `"usa"`
//   - `"variables-query"`
func (builder *QueryV2Builder) ScenarioId(scenarioId DataqueryScenarioId) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).ScenarioId = &scenarioId

	return builder
}

func (builder *QueryV2Builder) SeriesCount(seriesCount int64) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).SeriesCount = &seriesCount

	return builder
}

func (builder *QueryV2Builder) Sim(sim cog.Builder[SimulationQuery]) *QueryV2Builder {
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

func (builder *QueryV2Builder) SpanCount(spanCount int64) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).SpanCount = &spanCount

	return builder
}

func (builder *QueryV2Builder) Spread(spread float64) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).Spread = &spread

	return builder
}

func (builder *QueryV2Builder) StartValue(startValue float64) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).StartValue = &startValue

	return builder
}

func (builder *QueryV2Builder) Stream(stream cog.Builder[StreamingQuery]) *QueryV2Builder {
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
func (builder *QueryV2Builder) StringInput(stringInput string) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).StringInput = &stringInput

	return builder
}

// TimeRange represents the query range
// NOTE: unlike generic /ds/query, we can now send explicit time values in each query
// NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
func (builder *QueryV2Builder) TimeRange(timeRange cog.Builder[TimeRange]) *QueryV2Builder {
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

func (builder *QueryV2Builder) Usa(usa cog.Builder[USAQuery]) *QueryV2Builder {
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

func (builder *QueryV2Builder) WithNil(withNil bool) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).WithNil = &withNil

	return builder
}
