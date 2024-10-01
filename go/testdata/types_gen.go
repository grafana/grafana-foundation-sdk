// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package testdata

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	dashboard "github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

type TestDataQueryType string

const (
	TestDataQueryTypeRandomWalk                   TestDataQueryType = "random_walk"
	TestDataQueryTypeSlowQuery                    TestDataQueryType = "slow_query"
	TestDataQueryTypeRandomWalkWithError          TestDataQueryType = "random_walk_with_error"
	TestDataQueryTypeRandomWalkTable              TestDataQueryType = "random_walk_table"
	TestDataQueryTypeExponentialHeatmapBucketData TestDataQueryType = "exponential_heatmap_bucket_data"
	TestDataQueryTypeLinearHeatmapBucketData      TestDataQueryType = "linear_heatmap_bucket_data"
	TestDataQueryTypeNoDataPoints                 TestDataQueryType = "no_data_points"
	TestDataQueryTypeDataPointsOutsideRange       TestDataQueryType = "datapoints_outside_range"
	TestDataQueryTypeCSVMetricValues              TestDataQueryType = "csv_metric_values"
	TestDataQueryTypePredictablePulse             TestDataQueryType = "predictable_pulse"
	TestDataQueryTypePredictableCSVWave           TestDataQueryType = "predictable_csv_wave"
	TestDataQueryTypeStreamingClient              TestDataQueryType = "streaming_client"
	TestDataQueryTypeSimulation                   TestDataQueryType = "simulation"
	TestDataQueryTypeUSA                          TestDataQueryType = "usa"
	TestDataQueryTypeLive                         TestDataQueryType = "live"
	TestDataQueryTypeGrafanaAPI                   TestDataQueryType = "grafana_api"
	TestDataQueryTypeArrow                        TestDataQueryType = "arrow"
	TestDataQueryTypeAnnotations                  TestDataQueryType = "annotations"
	TestDataQueryTypeTableStatic                  TestDataQueryType = "table_static"
	TestDataQueryTypeServerError500               TestDataQueryType = "server_error_500"
	TestDataQueryTypeLogs                         TestDataQueryType = "logs"
	TestDataQueryTypeNodeGraph                    TestDataQueryType = "node_graph"
	TestDataQueryTypeFlameGraph                   TestDataQueryType = "flame_graph"
	TestDataQueryTypeRawFrame                     TestDataQueryType = "raw_frame"
	TestDataQueryTypeCSVFile                      TestDataQueryType = "csv_file"
	TestDataQueryTypeCSVContent                   TestDataQueryType = "csv_content"
	TestDataQueryTypeTrace                        TestDataQueryType = "trace"
	TestDataQueryTypeManualEntry                  TestDataQueryType = "manual_entry"
	TestDataQueryTypeVariablesQuery               TestDataQueryType = "variables-query"
)

type StreamingQuery struct {
	Type   StreamingQueryType `json:"type"`
	Speed  int32              `json:"speed"`
	Spread int32              `json:"spread"`
	Noise  int32              `json:"noise"`
	Bands  *int32             `json:"bands,omitempty"`
	Url    *string            `json:"url,omitempty"`
}

func (resource StreamingQuery) Equals(other StreamingQuery) bool {
	if resource.Type != other.Type {
		return false
	}
	if resource.Speed != other.Speed {
		return false
	}
	if resource.Spread != other.Spread {
		return false
	}
	if resource.Noise != other.Noise {
		return false
	}
	if resource.Bands == nil && other.Bands != nil || resource.Bands != nil && other.Bands == nil {
		return false
	}

	if resource.Bands != nil {
		if *resource.Bands != *other.Bands {
			return false
		}
	}
	if resource.Url == nil && other.Url != nil || resource.Url != nil && other.Url == nil {
		return false
	}

	if resource.Url != nil {
		if *resource.Url != *other.Url {
			return false
		}
	}

	return true
}

type PulseWaveQuery struct {
	TimeStep *int64   `json:"timeStep,omitempty"`
	OnCount  *int64   `json:"onCount,omitempty"`
	OffCount *int64   `json:"offCount,omitempty"`
	OnValue  *float64 `json:"onValue,omitempty"`
	OffValue *float64 `json:"offValue,omitempty"`
}

func (resource PulseWaveQuery) Equals(other PulseWaveQuery) bool {
	if resource.TimeStep == nil && other.TimeStep != nil || resource.TimeStep != nil && other.TimeStep == nil {
		return false
	}

	if resource.TimeStep != nil {
		if *resource.TimeStep != *other.TimeStep {
			return false
		}
	}
	if resource.OnCount == nil && other.OnCount != nil || resource.OnCount != nil && other.OnCount == nil {
		return false
	}

	if resource.OnCount != nil {
		if *resource.OnCount != *other.OnCount {
			return false
		}
	}
	if resource.OffCount == nil && other.OffCount != nil || resource.OffCount != nil && other.OffCount == nil {
		return false
	}

	if resource.OffCount != nil {
		if *resource.OffCount != *other.OffCount {
			return false
		}
	}
	if resource.OnValue == nil && other.OnValue != nil || resource.OnValue != nil && other.OnValue == nil {
		return false
	}

	if resource.OnValue != nil {
		if *resource.OnValue != *other.OnValue {
			return false
		}
	}
	if resource.OffValue == nil && other.OffValue != nil || resource.OffValue != nil && other.OffValue == nil {
		return false
	}

	if resource.OffValue != nil {
		if *resource.OffValue != *other.OffValue {
			return false
		}
	}

	return true
}

type SimulationQuery struct {
	Key    Key            `json:"key"`
	Config map[string]any `json:"config,omitempty"`
	Stream *bool          `json:"stream,omitempty"`
	Last   *bool          `json:"last,omitempty"`
}

func (resource SimulationQuery) Equals(other SimulationQuery) bool {
	if !resource.Key.Equals(other.Key) {
		return false
	}

	if len(resource.Config) != len(other.Config) {
		return false
	}

	for key1 := range resource.Config {
		// is DeepEqual good enough here?
		if !reflect.DeepEqual(resource.Config[key1], other.Config[key1]) {
			return false
		}
	}
	if resource.Stream == nil && other.Stream != nil || resource.Stream != nil && other.Stream == nil {
		return false
	}

	if resource.Stream != nil {
		if *resource.Stream != *other.Stream {
			return false
		}
	}
	if resource.Last == nil && other.Last != nil || resource.Last != nil && other.Last == nil {
		return false
	}

	if resource.Last != nil {
		if *resource.Last != *other.Last {
			return false
		}
	}

	return true
}

type NodesQuery struct {
	Type  *NodesQueryType `json:"type,omitempty"`
	Count *int64          `json:"count,omitempty"`
}

func (resource NodesQuery) Equals(other NodesQuery) bool {
	if resource.Type == nil && other.Type != nil || resource.Type != nil && other.Type == nil {
		return false
	}

	if resource.Type != nil {
		if *resource.Type != *other.Type {
			return false
		}
	}
	if resource.Count == nil && other.Count != nil || resource.Count != nil && other.Count == nil {
		return false
	}

	if resource.Count != nil {
		if *resource.Count != *other.Count {
			return false
		}
	}

	return true
}

type USAQuery struct {
	Mode   *string  `json:"mode,omitempty"`
	Period *string  `json:"period,omitempty"`
	Fields []string `json:"fields,omitempty"`
	States []string `json:"states,omitempty"`
}

func (resource USAQuery) Equals(other USAQuery) bool {
	if resource.Mode == nil && other.Mode != nil || resource.Mode != nil && other.Mode == nil {
		return false
	}

	if resource.Mode != nil {
		if *resource.Mode != *other.Mode {
			return false
		}
	}
	if resource.Period == nil && other.Period != nil || resource.Period != nil && other.Period == nil {
		return false
	}

	if resource.Period != nil {
		if *resource.Period != *other.Period {
			return false
		}
	}

	if len(resource.Fields) != len(other.Fields) {
		return false
	}

	for i1 := range resource.Fields {
		if resource.Fields[i1] != other.Fields[i1] {
			return false
		}
	}

	if len(resource.States) != len(other.States) {
		return false
	}

	for i1 := range resource.States {
		if resource.States[i1] != other.States[i1] {
			return false
		}
	}

	return true
}

type CSVWave struct {
	TimeStep  *int64  `json:"timeStep,omitempty"`
	Name      *string `json:"name,omitempty"`
	ValuesCSV *string `json:"valuesCSV,omitempty"`
	Labels    *string `json:"labels,omitempty"`
}

func (resource CSVWave) Equals(other CSVWave) bool {
	if resource.TimeStep == nil && other.TimeStep != nil || resource.TimeStep != nil && other.TimeStep == nil {
		return false
	}

	if resource.TimeStep != nil {
		if *resource.TimeStep != *other.TimeStep {
			return false
		}
	}
	if resource.Name == nil && other.Name != nil || resource.Name != nil && other.Name == nil {
		return false
	}

	if resource.Name != nil {
		if *resource.Name != *other.Name {
			return false
		}
	}
	if resource.ValuesCSV == nil && other.ValuesCSV != nil || resource.ValuesCSV != nil && other.ValuesCSV == nil {
		return false
	}

	if resource.ValuesCSV != nil {
		if *resource.ValuesCSV != *other.ValuesCSV {
			return false
		}
	}
	if resource.Labels == nil && other.Labels != nil || resource.Labels != nil && other.Labels == nil {
		return false
	}

	if resource.Labels != nil {
		if *resource.Labels != *other.Labels {
			return false
		}
	}

	return true
}

// TODO: Should this live here given it's not used in the dataquery?
type Scenario struct {
	Id             string  `json:"id"`
	Name           string  `json:"name"`
	StringInput    string  `json:"stringInput"`
	Description    *string `json:"description,omitempty"`
	HideAliasField *bool   `json:"hideAliasField,omitempty"`
}

func (resource Scenario) Equals(other Scenario) bool {
	if resource.Id != other.Id {
		return false
	}
	if resource.Name != other.Name {
		return false
	}
	if resource.StringInput != other.StringInput {
		return false
	}
	if resource.Description == nil && other.Description != nil || resource.Description != nil && other.Description == nil {
		return false
	}

	if resource.Description != nil {
		if *resource.Description != *other.Description {
			return false
		}
	}
	if resource.HideAliasField == nil && other.HideAliasField != nil || resource.HideAliasField != nil && other.HideAliasField == nil {
		return false
	}

	if resource.HideAliasField != nil {
		if *resource.HideAliasField != *other.HideAliasField {
			return false
		}
	}

	return true
}

type Dataquery struct {
	Alias           *string             `json:"alias,omitempty"`
	ScenarioId      *TestDataQueryType  `json:"scenarioId,omitempty"`
	StringInput     *string             `json:"stringInput,omitempty"`
	Stream          *StreamingQuery     `json:"stream,omitempty"`
	PulseWave       *PulseWaveQuery     `json:"pulseWave,omitempty"`
	Sim             *SimulationQuery    `json:"sim,omitempty"`
	CsvWave         []CSVWave           `json:"csvWave,omitempty"`
	Labels          *string             `json:"labels,omitempty"`
	Lines           *int64              `json:"lines,omitempty"`
	LevelColumn     *bool               `json:"levelColumn,omitempty"`
	Channel         *string             `json:"channel,omitempty"`
	Nodes           *NodesQuery         `json:"nodes,omitempty"`
	CsvFileName     *string             `json:"csvFileName,omitempty"`
	CsvContent      *string             `json:"csvContent,omitempty"`
	RawFrameContent *string             `json:"rawFrameContent,omitempty"`
	SeriesCount     *int32              `json:"seriesCount,omitempty"`
	Usa             *USAQuery           `json:"usa,omitempty"`
	ErrorType       *DataqueryErrorType `json:"errorType,omitempty"`
	SpanCount       *int32              `json:"spanCount,omitempty"`
	Points          [][]StringOrInt64   `json:"points,omitempty"`
	// Drop percentage (the chance we will lose a point 0-100)
	DropPercent    *float64 `json:"dropPercent,omitempty"`
	FlamegraphDiff *bool    `json:"flamegraphDiff,omitempty"`
	// A unique identifier for the query within the list of targets.
	// In server side expressions, the refId is used as a variable name to identify results.
	// By default, the UI will assign A->Z; however setting meaningful names may be useful.
	RefId string `json:"refId"`
	// true if query is disabled (ie should not be returned to the dashboard)
	// Note this does not always imply that the query should not be executed since
	// the results from a hidden query may be used as the input to other queries (SSE etc)
	Hide *bool `json:"hide,omitempty"`
	// Specify the query flavor
	// TODO make this required and give it a default
	QueryType *string `json:"queryType,omitempty"`
	// For mixed data sources the selected datasource is on the query level.
	// For non mixed scenarios this is undefined.
	// TODO find a better way to do this ^ that's friendly to schema
	// TODO this shouldn't be unknown but DataSourceRef | null
	Datasource *dashboard.DataSourceRef `json:"datasource,omitempty"`
}

func (resource Dataquery) ImplementsDataqueryVariant() {}

func (resource Dataquery) DataqueryType() string {
	return "testdata"
}

func VariantConfig() variants.DataqueryConfig {
	return variants.DataqueryConfig{
		Identifier: "testdata",
		DataqueryUnmarshaler: func(raw []byte) (variants.Dataquery, error) {
			dataquery := &Dataquery{}

			if err := json.Unmarshal(raw, dataquery); err != nil {
				return nil, err
			}

			return dataquery, nil
		},
	}
}

func (resource Dataquery) Equals(otherCandidate variants.Dataquery) bool {
	if otherCandidate == nil {
		return false
	}

	other, ok := otherCandidate.(Dataquery)
	if !ok {
		return false
	}
	if resource.Alias == nil && other.Alias != nil || resource.Alias != nil && other.Alias == nil {
		return false
	}

	if resource.Alias != nil {
		if *resource.Alias != *other.Alias {
			return false
		}
	}
	if resource.ScenarioId == nil && other.ScenarioId != nil || resource.ScenarioId != nil && other.ScenarioId == nil {
		return false
	}

	if resource.ScenarioId != nil {
		if *resource.ScenarioId != *other.ScenarioId {
			return false
		}
	}
	if resource.StringInput == nil && other.StringInput != nil || resource.StringInput != nil && other.StringInput == nil {
		return false
	}

	if resource.StringInput != nil {
		if *resource.StringInput != *other.StringInput {
			return false
		}
	}
	if resource.Stream == nil && other.Stream != nil || resource.Stream != nil && other.Stream == nil {
		return false
	}

	if resource.Stream != nil {
		if !resource.Stream.Equals(*other.Stream) {
			return false
		}
	}
	if resource.PulseWave == nil && other.PulseWave != nil || resource.PulseWave != nil && other.PulseWave == nil {
		return false
	}

	if resource.PulseWave != nil {
		if !resource.PulseWave.Equals(*other.PulseWave) {
			return false
		}
	}
	if resource.Sim == nil && other.Sim != nil || resource.Sim != nil && other.Sim == nil {
		return false
	}

	if resource.Sim != nil {
		if !resource.Sim.Equals(*other.Sim) {
			return false
		}
	}

	if len(resource.CsvWave) != len(other.CsvWave) {
		return false
	}

	for i1 := range resource.CsvWave {
		if !resource.CsvWave[i1].Equals(other.CsvWave[i1]) {
			return false
		}
	}
	if resource.Labels == nil && other.Labels != nil || resource.Labels != nil && other.Labels == nil {
		return false
	}

	if resource.Labels != nil {
		if *resource.Labels != *other.Labels {
			return false
		}
	}
	if resource.Lines == nil && other.Lines != nil || resource.Lines != nil && other.Lines == nil {
		return false
	}

	if resource.Lines != nil {
		if *resource.Lines != *other.Lines {
			return false
		}
	}
	if resource.LevelColumn == nil && other.LevelColumn != nil || resource.LevelColumn != nil && other.LevelColumn == nil {
		return false
	}

	if resource.LevelColumn != nil {
		if *resource.LevelColumn != *other.LevelColumn {
			return false
		}
	}
	if resource.Channel == nil && other.Channel != nil || resource.Channel != nil && other.Channel == nil {
		return false
	}

	if resource.Channel != nil {
		if *resource.Channel != *other.Channel {
			return false
		}
	}
	if resource.Nodes == nil && other.Nodes != nil || resource.Nodes != nil && other.Nodes == nil {
		return false
	}

	if resource.Nodes != nil {
		if !resource.Nodes.Equals(*other.Nodes) {
			return false
		}
	}
	if resource.CsvFileName == nil && other.CsvFileName != nil || resource.CsvFileName != nil && other.CsvFileName == nil {
		return false
	}

	if resource.CsvFileName != nil {
		if *resource.CsvFileName != *other.CsvFileName {
			return false
		}
	}
	if resource.CsvContent == nil && other.CsvContent != nil || resource.CsvContent != nil && other.CsvContent == nil {
		return false
	}

	if resource.CsvContent != nil {
		if *resource.CsvContent != *other.CsvContent {
			return false
		}
	}
	if resource.RawFrameContent == nil && other.RawFrameContent != nil || resource.RawFrameContent != nil && other.RawFrameContent == nil {
		return false
	}

	if resource.RawFrameContent != nil {
		if *resource.RawFrameContent != *other.RawFrameContent {
			return false
		}
	}
	if resource.SeriesCount == nil && other.SeriesCount != nil || resource.SeriesCount != nil && other.SeriesCount == nil {
		return false
	}

	if resource.SeriesCount != nil {
		if *resource.SeriesCount != *other.SeriesCount {
			return false
		}
	}
	if resource.Usa == nil && other.Usa != nil || resource.Usa != nil && other.Usa == nil {
		return false
	}

	if resource.Usa != nil {
		if !resource.Usa.Equals(*other.Usa) {
			return false
		}
	}
	if resource.ErrorType == nil && other.ErrorType != nil || resource.ErrorType != nil && other.ErrorType == nil {
		return false
	}

	if resource.ErrorType != nil {
		if *resource.ErrorType != *other.ErrorType {
			return false
		}
	}
	if resource.SpanCount == nil && other.SpanCount != nil || resource.SpanCount != nil && other.SpanCount == nil {
		return false
	}

	if resource.SpanCount != nil {
		if *resource.SpanCount != *other.SpanCount {
			return false
		}
	}

	if len(resource.Points) != len(other.Points) {
		return false
	}

	for i1 := range resource.Points {

		if len(resource.Points[i1]) != len(other.Points[i1]) {
			return false
		}

		for i2 := range resource.Points[i1] {
			if !resource.Points[i1][i2].Equals(other.Points[i1][i2]) {
				return false
			}
		}
	}
	if resource.DropPercent == nil && other.DropPercent != nil || resource.DropPercent != nil && other.DropPercent == nil {
		return false
	}

	if resource.DropPercent != nil {
		if *resource.DropPercent != *other.DropPercent {
			return false
		}
	}
	if resource.FlamegraphDiff == nil && other.FlamegraphDiff != nil || resource.FlamegraphDiff != nil && other.FlamegraphDiff == nil {
		return false
	}

	if resource.FlamegraphDiff != nil {
		if *resource.FlamegraphDiff != *other.FlamegraphDiff {
			return false
		}
	}
	if resource.RefId != other.RefId {
		return false
	}
	if resource.Hide == nil && other.Hide != nil || resource.Hide != nil && other.Hide == nil {
		return false
	}

	if resource.Hide != nil {
		if *resource.Hide != *other.Hide {
			return false
		}
	}
	if resource.QueryType == nil && other.QueryType != nil || resource.QueryType != nil && other.QueryType == nil {
		return false
	}

	if resource.QueryType != nil {
		if *resource.QueryType != *other.QueryType {
			return false
		}
	}
	if resource.Datasource == nil && other.Datasource != nil || resource.Datasource != nil && other.Datasource == nil {
		return false
	}

	if resource.Datasource != nil {
		if !resource.Datasource.Equals(*other.Datasource) {
			return false
		}
	}

	return true
}

type Key struct {
	Type string  `json:"type"`
	Tick float64 `json:"tick"`
	Uid  *string `json:"uid,omitempty"`
}

func (resource Key) Equals(other Key) bool {
	if resource.Type != other.Type {
		return false
	}
	if resource.Tick != other.Tick {
		return false
	}
	if resource.Uid == nil && other.Uid != nil || resource.Uid != nil && other.Uid == nil {
		return false
	}

	if resource.Uid != nil {
		if *resource.Uid != *other.Uid {
			return false
		}
	}

	return true
}

type StreamingQueryType string

const (
	StreamingQueryTypeSignal StreamingQueryType = "signal"
	StreamingQueryTypeLogs   StreamingQueryType = "logs"
	StreamingQueryTypeFetch  StreamingQueryType = "fetch"
)

type NodesQueryType string

const (
	NodesQueryTypeRandom      NodesQueryType = "random"
	NodesQueryTypeResponse    NodesQueryType = "response"
	NodesQueryTypeRandomEdges NodesQueryType = "random edges"
)

type DataqueryErrorType string

const (
	DataqueryErrorTypeServerPanic        DataqueryErrorType = "server_panic"
	DataqueryErrorTypeFrontendException  DataqueryErrorType = "frontend_exception"
	DataqueryErrorTypeFrontendObservable DataqueryErrorType = "frontend_observable"
)

type StringOrInt64 struct {
	String *string `json:"String,omitempty"`
	Int64  *int64  `json:"Int64,omitempty"`
}

func (resource StringOrInt64) MarshalJSON() ([]byte, error) {
	if resource.String != nil {
		return json.Marshal(resource.String)
	}

	if resource.Int64 != nil {
		return json.Marshal(resource.Int64)
	}

	return nil, fmt.Errorf("no value for disjunction of scalars")
}

func (resource *StringOrInt64) UnmarshalJSON(raw []byte) error {
	if raw == nil {
		return nil
	}

	var errList []error

	// String
	var String string
	if err := json.Unmarshal(raw, &String); err != nil {
		errList = append(errList, err)
		resource.String = nil
	} else {
		resource.String = &String
		return nil
	}

	// Int64
	var Int64 int64
	if err := json.Unmarshal(raw, &Int64); err != nil {
		errList = append(errList, err)
		resource.Int64 = nil
	} else {
		resource.Int64 = &Int64
		return nil
	}

	return errors.Join(errList...)
}

func (resource StringOrInt64) Equals(other StringOrInt64) bool {
	if resource.String == nil && other.String != nil || resource.String != nil && other.String == nil {
		return false
	}

	if resource.String != nil {
		if *resource.String != *other.String {
			return false
		}
	}
	if resource.Int64 == nil && other.Int64 != nil || resource.Int64 != nil && other.Int64 == nil {
		return false
	}

	if resource.Int64 != nil {
		if *resource.Int64 != *other.Int64 {
			return false
		}
	}

	return true
}
