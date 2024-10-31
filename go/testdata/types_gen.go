// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package testdata

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
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

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `StreamingQuery` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *StreamingQuery) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "type"
	if fields["type"] != nil {
		if string(fields["type"]) != "null" {
			if err := json.Unmarshal(fields["type"], &resource.Type); err != nil {
				errs = append(errs, cog.MakeBuildErrors("type", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is null"))...)

		}
		delete(fields, "type")
	} else {
		errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is missing from input"))...)
	}
	// Field "speed"
	if fields["speed"] != nil {
		if string(fields["speed"]) != "null" {
			if err := json.Unmarshal(fields["speed"], &resource.Speed); err != nil {
				errs = append(errs, cog.MakeBuildErrors("speed", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("speed", errors.New("required field is null"))...)

		}
		delete(fields, "speed")
	} else {
		errs = append(errs, cog.MakeBuildErrors("speed", errors.New("required field is missing from input"))...)
	}
	// Field "spread"
	if fields["spread"] != nil {
		if string(fields["spread"]) != "null" {
			if err := json.Unmarshal(fields["spread"], &resource.Spread); err != nil {
				errs = append(errs, cog.MakeBuildErrors("spread", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("spread", errors.New("required field is null"))...)

		}
		delete(fields, "spread")
	} else {
		errs = append(errs, cog.MakeBuildErrors("spread", errors.New("required field is missing from input"))...)
	}
	// Field "noise"
	if fields["noise"] != nil {
		if string(fields["noise"]) != "null" {
			if err := json.Unmarshal(fields["noise"], &resource.Noise); err != nil {
				errs = append(errs, cog.MakeBuildErrors("noise", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("noise", errors.New("required field is null"))...)

		}
		delete(fields, "noise")
	} else {
		errs = append(errs, cog.MakeBuildErrors("noise", errors.New("required field is missing from input"))...)
	}
	// Field "bands"
	if fields["bands"] != nil {
		if string(fields["bands"]) != "null" {
			if err := json.Unmarshal(fields["bands"], &resource.Bands); err != nil {
				errs = append(errs, cog.MakeBuildErrors("bands", err)...)
			}

		}
		delete(fields, "bands")

	}
	// Field "url"
	if fields["url"] != nil {
		if string(fields["url"]) != "null" {
			if err := json.Unmarshal(fields["url"], &resource.Url); err != nil {
				errs = append(errs, cog.MakeBuildErrors("url", err)...)
			}

		}
		delete(fields, "url")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("StreamingQuery", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `StreamingQuery` objects.
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

// Validate checks all the validation constraints that may be defined on `StreamingQuery` fields for violations and returns them.
func (resource StreamingQuery) Validate() error {
	return nil
}

type PulseWaveQuery struct {
	TimeStep *int64   `json:"timeStep,omitempty"`
	OnCount  *int64   `json:"onCount,omitempty"`
	OffCount *int64   `json:"offCount,omitempty"`
	OnValue  *float64 `json:"onValue,omitempty"`
	OffValue *float64 `json:"offValue,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `PulseWaveQuery` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *PulseWaveQuery) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "timeStep"
	if fields["timeStep"] != nil {
		if string(fields["timeStep"]) != "null" {
			if err := json.Unmarshal(fields["timeStep"], &resource.TimeStep); err != nil {
				errs = append(errs, cog.MakeBuildErrors("timeStep", err)...)
			}

		}
		delete(fields, "timeStep")

	}
	// Field "onCount"
	if fields["onCount"] != nil {
		if string(fields["onCount"]) != "null" {
			if err := json.Unmarshal(fields["onCount"], &resource.OnCount); err != nil {
				errs = append(errs, cog.MakeBuildErrors("onCount", err)...)
			}

		}
		delete(fields, "onCount")

	}
	// Field "offCount"
	if fields["offCount"] != nil {
		if string(fields["offCount"]) != "null" {
			if err := json.Unmarshal(fields["offCount"], &resource.OffCount); err != nil {
				errs = append(errs, cog.MakeBuildErrors("offCount", err)...)
			}

		}
		delete(fields, "offCount")

	}
	// Field "onValue"
	if fields["onValue"] != nil {
		if string(fields["onValue"]) != "null" {
			if err := json.Unmarshal(fields["onValue"], &resource.OnValue); err != nil {
				errs = append(errs, cog.MakeBuildErrors("onValue", err)...)
			}

		}
		delete(fields, "onValue")

	}
	// Field "offValue"
	if fields["offValue"] != nil {
		if string(fields["offValue"]) != "null" {
			if err := json.Unmarshal(fields["offValue"], &resource.OffValue); err != nil {
				errs = append(errs, cog.MakeBuildErrors("offValue", err)...)
			}

		}
		delete(fields, "offValue")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("PulseWaveQuery", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `PulseWaveQuery` objects.
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

// Validate checks all the validation constraints that may be defined on `PulseWaveQuery` fields for violations and returns them.
func (resource PulseWaveQuery) Validate() error {
	return nil
}

type SimulationQuery struct {
	Key    Key            `json:"key"`
	Config map[string]any `json:"config,omitempty"`
	Stream *bool          `json:"stream,omitempty"`
	Last   *bool          `json:"last,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `SimulationQuery` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *SimulationQuery) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "key"
	if fields["key"] != nil {
		if string(fields["key"]) != "null" {

			resource.Key = Key{}
			if err := resource.Key.UnmarshalJSONStrict(fields["key"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("key", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("key", errors.New("required field is null"))...)

		}
		delete(fields, "key")
	} else {
		errs = append(errs, cog.MakeBuildErrors("key", errors.New("required field is missing from input"))...)
	}
	// Field "config"
	if fields["config"] != nil {
		if string(fields["config"]) != "null" {

			if err := json.Unmarshal(fields["config"], &resource.Config); err != nil {
				errs = append(errs, cog.MakeBuildErrors("config", err)...)
			}

		}
		delete(fields, "config")

	}
	// Field "stream"
	if fields["stream"] != nil {
		if string(fields["stream"]) != "null" {
			if err := json.Unmarshal(fields["stream"], &resource.Stream); err != nil {
				errs = append(errs, cog.MakeBuildErrors("stream", err)...)
			}

		}
		delete(fields, "stream")

	}
	// Field "last"
	if fields["last"] != nil {
		if string(fields["last"]) != "null" {
			if err := json.Unmarshal(fields["last"], &resource.Last); err != nil {
				errs = append(errs, cog.MakeBuildErrors("last", err)...)
			}

		}
		delete(fields, "last")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("SimulationQuery", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `SimulationQuery` objects.
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

// Validate checks all the validation constraints that may be defined on `SimulationQuery` fields for violations and returns them.
func (resource SimulationQuery) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Key.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("key", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type NodesQuery struct {
	Type  *NodesQueryType `json:"type,omitempty"`
	Count *int64          `json:"count,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `NodesQuery` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *NodesQuery) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "type"
	if fields["type"] != nil {
		if string(fields["type"]) != "null" {
			if err := json.Unmarshal(fields["type"], &resource.Type); err != nil {
				errs = append(errs, cog.MakeBuildErrors("type", err)...)
			}

		}
		delete(fields, "type")

	}
	// Field "count"
	if fields["count"] != nil {
		if string(fields["count"]) != "null" {
			if err := json.Unmarshal(fields["count"], &resource.Count); err != nil {
				errs = append(errs, cog.MakeBuildErrors("count", err)...)
			}

		}
		delete(fields, "count")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("NodesQuery", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `NodesQuery` objects.
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

// Validate checks all the validation constraints that may be defined on `NodesQuery` fields for violations and returns them.
func (resource NodesQuery) Validate() error {
	return nil
}

type USAQuery struct {
	Mode   *string  `json:"mode,omitempty"`
	Period *string  `json:"period,omitempty"`
	Fields []string `json:"fields,omitempty"`
	States []string `json:"states,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `USAQuery` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *USAQuery) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "mode"
	if fields["mode"] != nil {
		if string(fields["mode"]) != "null" {
			if err := json.Unmarshal(fields["mode"], &resource.Mode); err != nil {
				errs = append(errs, cog.MakeBuildErrors("mode", err)...)
			}

		}
		delete(fields, "mode")

	}
	// Field "period"
	if fields["period"] != nil {
		if string(fields["period"]) != "null" {
			if err := json.Unmarshal(fields["period"], &resource.Period); err != nil {
				errs = append(errs, cog.MakeBuildErrors("period", err)...)
			}

		}
		delete(fields, "period")

	}
	// Field "fields"
	if fields["fields"] != nil {
		if string(fields["fields"]) != "null" {

			if err := json.Unmarshal(fields["fields"], &resource.Fields); err != nil {
				errs = append(errs, cog.MakeBuildErrors("fields", err)...)
			}

		}
		delete(fields, "fields")

	}
	// Field "states"
	if fields["states"] != nil {
		if string(fields["states"]) != "null" {

			if err := json.Unmarshal(fields["states"], &resource.States); err != nil {
				errs = append(errs, cog.MakeBuildErrors("states", err)...)
			}

		}
		delete(fields, "states")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("USAQuery", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `USAQuery` objects.
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

// Validate checks all the validation constraints that may be defined on `USAQuery` fields for violations and returns them.
func (resource USAQuery) Validate() error {
	return nil
}

type CSVWave struct {
	TimeStep  *int64  `json:"timeStep,omitempty"`
	Name      *string `json:"name,omitempty"`
	ValuesCSV *string `json:"valuesCSV,omitempty"`
	Labels    *string `json:"labels,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `CSVWave` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *CSVWave) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "timeStep"
	if fields["timeStep"] != nil {
		if string(fields["timeStep"]) != "null" {
			if err := json.Unmarshal(fields["timeStep"], &resource.TimeStep); err != nil {
				errs = append(errs, cog.MakeBuildErrors("timeStep", err)...)
			}

		}
		delete(fields, "timeStep")

	}
	// Field "name"
	if fields["name"] != nil {
		if string(fields["name"]) != "null" {
			if err := json.Unmarshal(fields["name"], &resource.Name); err != nil {
				errs = append(errs, cog.MakeBuildErrors("name", err)...)
			}

		}
		delete(fields, "name")

	}
	// Field "valuesCSV"
	if fields["valuesCSV"] != nil {
		if string(fields["valuesCSV"]) != "null" {
			if err := json.Unmarshal(fields["valuesCSV"], &resource.ValuesCSV); err != nil {
				errs = append(errs, cog.MakeBuildErrors("valuesCSV", err)...)
			}

		}
		delete(fields, "valuesCSV")

	}
	// Field "labels"
	if fields["labels"] != nil {
		if string(fields["labels"]) != "null" {
			if err := json.Unmarshal(fields["labels"], &resource.Labels); err != nil {
				errs = append(errs, cog.MakeBuildErrors("labels", err)...)
			}

		}
		delete(fields, "labels")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("CSVWave", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `CSVWave` objects.
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

// Validate checks all the validation constraints that may be defined on `CSVWave` fields for violations and returns them.
func (resource CSVWave) Validate() error {
	return nil
}

// TODO: Should this live here given it's not used in the dataquery?
type Scenario struct {
	Id             string  `json:"id"`
	Name           string  `json:"name"`
	StringInput    string  `json:"stringInput"`
	Description    *string `json:"description,omitempty"`
	HideAliasField *bool   `json:"hideAliasField,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Scenario` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *Scenario) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "id"
	if fields["id"] != nil {
		if string(fields["id"]) != "null" {
			if err := json.Unmarshal(fields["id"], &resource.Id); err != nil {
				errs = append(errs, cog.MakeBuildErrors("id", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("id", errors.New("required field is null"))...)

		}
		delete(fields, "id")
	} else {
		errs = append(errs, cog.MakeBuildErrors("id", errors.New("required field is missing from input"))...)
	}
	// Field "name"
	if fields["name"] != nil {
		if string(fields["name"]) != "null" {
			if err := json.Unmarshal(fields["name"], &resource.Name); err != nil {
				errs = append(errs, cog.MakeBuildErrors("name", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("name", errors.New("required field is null"))...)

		}
		delete(fields, "name")
	} else {
		errs = append(errs, cog.MakeBuildErrors("name", errors.New("required field is missing from input"))...)
	}
	// Field "stringInput"
	if fields["stringInput"] != nil {
		if string(fields["stringInput"]) != "null" {
			if err := json.Unmarshal(fields["stringInput"], &resource.StringInput); err != nil {
				errs = append(errs, cog.MakeBuildErrors("stringInput", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("stringInput", errors.New("required field is null"))...)

		}
		delete(fields, "stringInput")
	} else {
		errs = append(errs, cog.MakeBuildErrors("stringInput", errors.New("required field is missing from input"))...)
	}
	// Field "description"
	if fields["description"] != nil {
		if string(fields["description"]) != "null" {
			if err := json.Unmarshal(fields["description"], &resource.Description); err != nil {
				errs = append(errs, cog.MakeBuildErrors("description", err)...)
			}

		}
		delete(fields, "description")

	}
	// Field "hideAliasField"
	if fields["hideAliasField"] != nil {
		if string(fields["hideAliasField"]) != "null" {
			if err := json.Unmarshal(fields["hideAliasField"], &resource.HideAliasField); err != nil {
				errs = append(errs, cog.MakeBuildErrors("hideAliasField", err)...)
			}

		}
		delete(fields, "hideAliasField")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("Scenario", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `Scenario` objects.
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

// Validate checks all the validation constraints that may be defined on `Scenario` fields for violations and returns them.
func (resource Scenario) Validate() error {
	return nil
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
	DropPercent *float64 `json:"dropPercent,omitempty"`
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

// VariantConfig returns the configuration related to testdata dataqueries.
// This configuration describes how to unmarshal it, convert it to code, …
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
		StrictDataqueryUnmarshaler: func(raw []byte) (variants.Dataquery, error) {
			dataquery := &Dataquery{}

			if err := dataquery.UnmarshalJSONStrict(raw); err != nil {
				return nil, err
			}

			return dataquery, nil
		},
		GoConverter: func(input any) string {
			var dataquery Dataquery
			if cast, ok := input.(*Dataquery); ok {
				dataquery = *cast
			} else {
				dataquery = input.(Dataquery)
			}
			return DataqueryConverter(dataquery)
		},
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Dataquery` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *Dataquery) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "alias"
	if fields["alias"] != nil {
		if string(fields["alias"]) != "null" {
			if err := json.Unmarshal(fields["alias"], &resource.Alias); err != nil {
				errs = append(errs, cog.MakeBuildErrors("alias", err)...)
			}

		}
		delete(fields, "alias")

	}
	// Field "scenarioId"
	if fields["scenarioId"] != nil {
		if string(fields["scenarioId"]) != "null" {
			if err := json.Unmarshal(fields["scenarioId"], &resource.ScenarioId); err != nil {
				errs = append(errs, cog.MakeBuildErrors("scenarioId", err)...)
			}

		}
		delete(fields, "scenarioId")

	}
	// Field "stringInput"
	if fields["stringInput"] != nil {
		if string(fields["stringInput"]) != "null" {
			if err := json.Unmarshal(fields["stringInput"], &resource.StringInput); err != nil {
				errs = append(errs, cog.MakeBuildErrors("stringInput", err)...)
			}

		}
		delete(fields, "stringInput")

	}
	// Field "stream"
	if fields["stream"] != nil {
		if string(fields["stream"]) != "null" {

			resource.Stream = &StreamingQuery{}
			if err := resource.Stream.UnmarshalJSONStrict(fields["stream"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("stream", err)...)
			}

		}
		delete(fields, "stream")

	}
	// Field "pulseWave"
	if fields["pulseWave"] != nil {
		if string(fields["pulseWave"]) != "null" {

			resource.PulseWave = &PulseWaveQuery{}
			if err := resource.PulseWave.UnmarshalJSONStrict(fields["pulseWave"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("pulseWave", err)...)
			}

		}
		delete(fields, "pulseWave")

	}
	// Field "sim"
	if fields["sim"] != nil {
		if string(fields["sim"]) != "null" {

			resource.Sim = &SimulationQuery{}
			if err := resource.Sim.UnmarshalJSONStrict(fields["sim"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("sim", err)...)
			}

		}
		delete(fields, "sim")

	}
	// Field "csvWave"
	if fields["csvWave"] != nil {
		if string(fields["csvWave"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["csvWave"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 CSVWave

				result1 = CSVWave{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("csvWave["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.CsvWave = append(resource.CsvWave, result1)
			}

		}
		delete(fields, "csvWave")

	}
	// Field "labels"
	if fields["labels"] != nil {
		if string(fields["labels"]) != "null" {
			if err := json.Unmarshal(fields["labels"], &resource.Labels); err != nil {
				errs = append(errs, cog.MakeBuildErrors("labels", err)...)
			}

		}
		delete(fields, "labels")

	}
	// Field "lines"
	if fields["lines"] != nil {
		if string(fields["lines"]) != "null" {
			if err := json.Unmarshal(fields["lines"], &resource.Lines); err != nil {
				errs = append(errs, cog.MakeBuildErrors("lines", err)...)
			}

		}
		delete(fields, "lines")

	}
	// Field "levelColumn"
	if fields["levelColumn"] != nil {
		if string(fields["levelColumn"]) != "null" {
			if err := json.Unmarshal(fields["levelColumn"], &resource.LevelColumn); err != nil {
				errs = append(errs, cog.MakeBuildErrors("levelColumn", err)...)
			}

		}
		delete(fields, "levelColumn")

	}
	// Field "channel"
	if fields["channel"] != nil {
		if string(fields["channel"]) != "null" {
			if err := json.Unmarshal(fields["channel"], &resource.Channel); err != nil {
				errs = append(errs, cog.MakeBuildErrors("channel", err)...)
			}

		}
		delete(fields, "channel")

	}
	// Field "nodes"
	if fields["nodes"] != nil {
		if string(fields["nodes"]) != "null" {

			resource.Nodes = &NodesQuery{}
			if err := resource.Nodes.UnmarshalJSONStrict(fields["nodes"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("nodes", err)...)
			}

		}
		delete(fields, "nodes")

	}
	// Field "csvFileName"
	if fields["csvFileName"] != nil {
		if string(fields["csvFileName"]) != "null" {
			if err := json.Unmarshal(fields["csvFileName"], &resource.CsvFileName); err != nil {
				errs = append(errs, cog.MakeBuildErrors("csvFileName", err)...)
			}

		}
		delete(fields, "csvFileName")

	}
	// Field "csvContent"
	if fields["csvContent"] != nil {
		if string(fields["csvContent"]) != "null" {
			if err := json.Unmarshal(fields["csvContent"], &resource.CsvContent); err != nil {
				errs = append(errs, cog.MakeBuildErrors("csvContent", err)...)
			}

		}
		delete(fields, "csvContent")

	}
	// Field "rawFrameContent"
	if fields["rawFrameContent"] != nil {
		if string(fields["rawFrameContent"]) != "null" {
			if err := json.Unmarshal(fields["rawFrameContent"], &resource.RawFrameContent); err != nil {
				errs = append(errs, cog.MakeBuildErrors("rawFrameContent", err)...)
			}

		}
		delete(fields, "rawFrameContent")

	}
	// Field "seriesCount"
	if fields["seriesCount"] != nil {
		if string(fields["seriesCount"]) != "null" {
			if err := json.Unmarshal(fields["seriesCount"], &resource.SeriesCount); err != nil {
				errs = append(errs, cog.MakeBuildErrors("seriesCount", err)...)
			}

		}
		delete(fields, "seriesCount")

	}
	// Field "usa"
	if fields["usa"] != nil {
		if string(fields["usa"]) != "null" {

			resource.Usa = &USAQuery{}
			if err := resource.Usa.UnmarshalJSONStrict(fields["usa"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("usa", err)...)
			}

		}
		delete(fields, "usa")

	}
	// Field "errorType"
	if fields["errorType"] != nil {
		if string(fields["errorType"]) != "null" {
			if err := json.Unmarshal(fields["errorType"], &resource.ErrorType); err != nil {
				errs = append(errs, cog.MakeBuildErrors("errorType", err)...)
			}

		}
		delete(fields, "errorType")

	}
	// Field "spanCount"
	if fields["spanCount"] != nil {
		if string(fields["spanCount"]) != "null" {
			if err := json.Unmarshal(fields["spanCount"], &resource.SpanCount); err != nil {
				errs = append(errs, cog.MakeBuildErrors("spanCount", err)...)
			}

		}
		delete(fields, "spanCount")

	}
	// Field "points"
	if fields["points"] != nil {
		if string(fields["points"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["points"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 []StringOrInt64

				partialArray := []json.RawMessage{}
				if err := json.Unmarshal(partialArray[i1], &partialArray); err != nil {
					return err
				}

				for i2 := range partialArray {
					var result2 StringOrInt64

					result2 = StringOrInt64{}
					if err := result2.UnmarshalJSONStrict(partialArray[i2]); err != nil {
						errs = append(errs, cog.MakeBuildErrors("points["+strconv.Itoa(i1)+"]["+strconv.Itoa(i2)+"]", err)...)
					}
					result1 = append(result1, result2)
				}
				resource.Points = append(resource.Points, result1)
			}

		}
		delete(fields, "points")

	}
	// Field "dropPercent"
	if fields["dropPercent"] != nil {
		if string(fields["dropPercent"]) != "null" {
			if err := json.Unmarshal(fields["dropPercent"], &resource.DropPercent); err != nil {
				errs = append(errs, cog.MakeBuildErrors("dropPercent", err)...)
			}

		}
		delete(fields, "dropPercent")

	}
	// Field "refId"
	if fields["refId"] != nil {
		if string(fields["refId"]) != "null" {
			if err := json.Unmarshal(fields["refId"], &resource.RefId); err != nil {
				errs = append(errs, cog.MakeBuildErrors("refId", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("refId", errors.New("required field is null"))...)

		}
		delete(fields, "refId")
	} else {
		errs = append(errs, cog.MakeBuildErrors("refId", errors.New("required field is missing from input"))...)
	}
	// Field "hide"
	if fields["hide"] != nil {
		if string(fields["hide"]) != "null" {
			if err := json.Unmarshal(fields["hide"], &resource.Hide); err != nil {
				errs = append(errs, cog.MakeBuildErrors("hide", err)...)
			}

		}
		delete(fields, "hide")

	}
	// Field "queryType"
	if fields["queryType"] != nil {
		if string(fields["queryType"]) != "null" {
			if err := json.Unmarshal(fields["queryType"], &resource.QueryType); err != nil {
				errs = append(errs, cog.MakeBuildErrors("queryType", err)...)
			}

		}
		delete(fields, "queryType")

	}
	// Field "datasource"
	if fields["datasource"] != nil {
		if string(fields["datasource"]) != "null" {

			resource.Datasource = &dashboard.DataSourceRef{}
			if err := resource.Datasource.UnmarshalJSONStrict(fields["datasource"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("datasource", err)...)
			}

		}
		delete(fields, "datasource")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("Dataquery", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two dataqueries.
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

// Validate checks all the validation constraints that may be defined on `Dataquery` fields for violations and returns them.
func (resource Dataquery) Validate() error {
	var errs cog.BuildErrors
	if resource.Stream != nil {
		if err := resource.Stream.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("stream", err)...)
		}
	}
	if resource.PulseWave != nil {
		if err := resource.PulseWave.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("pulseWave", err)...)
		}
	}
	if resource.Sim != nil {
		if err := resource.Sim.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("sim", err)...)
		}
	}

	for i1 := range resource.CsvWave {
		if err := resource.CsvWave[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("csvWave["+strconv.Itoa(i1)+"]", err)...)
		}
	}
	if resource.Nodes != nil {
		if err := resource.Nodes.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("nodes", err)...)
		}
	}
	if resource.Usa != nil {
		if err := resource.Usa.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("usa", err)...)
		}
	}

	for i1 := range resource.Points {

		for i2 := range resource.Points[i1] {
			if err := resource.Points[i1][i2].Validate(); err != nil {
				errs = append(errs, cog.MakeBuildErrors("points["+strconv.Itoa(i1)+"]["+strconv.Itoa(i2)+"]", err)...)
			}
		}
	}
	if resource.Datasource != nil {
		if err := resource.Datasource.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("datasource", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type Key struct {
	Type string  `json:"type"`
	Tick float64 `json:"tick"`
	Uid  *string `json:"uid,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Key` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *Key) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "type"
	if fields["type"] != nil {
		if string(fields["type"]) != "null" {
			if err := json.Unmarshal(fields["type"], &resource.Type); err != nil {
				errs = append(errs, cog.MakeBuildErrors("type", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is null"))...)

		}
		delete(fields, "type")
	} else {
		errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is missing from input"))...)
	}
	// Field "tick"
	if fields["tick"] != nil {
		if string(fields["tick"]) != "null" {
			if err := json.Unmarshal(fields["tick"], &resource.Tick); err != nil {
				errs = append(errs, cog.MakeBuildErrors("tick", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("tick", errors.New("required field is null"))...)

		}
		delete(fields, "tick")
	} else {
		errs = append(errs, cog.MakeBuildErrors("tick", errors.New("required field is missing from input"))...)
	}
	// Field "uid"
	if fields["uid"] != nil {
		if string(fields["uid"]) != "null" {
			if err := json.Unmarshal(fields["uid"], &resource.Uid); err != nil {
				errs = append(errs, cog.MakeBuildErrors("uid", err)...)
			}

		}
		delete(fields, "uid")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("Key", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `Key` objects.
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

// Validate checks all the validation constraints that may be defined on `Key` fields for violations and returns them.
func (resource Key) Validate() error {
	return nil
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

// MarshalJSON implements a custom JSON marshalling logic to encode `StringOrInt64` as JSON.
func (resource StringOrInt64) MarshalJSON() ([]byte, error) {
	if resource.String != nil {
		return json.Marshal(resource.String)
	}

	if resource.Int64 != nil {
		return json.Marshal(resource.Int64)
	}

	return nil, fmt.Errorf("no value for disjunction of scalars")
}

// UnmarshalJSON implements a custom JSON unmarshalling logic to decode `StringOrInt64` from JSON.
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

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `StringOrInt64` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *StringOrInt64) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors
	var errList []error

	// String
	var String string

	if err := json.Unmarshal(raw, &String); err != nil {
		errList = append(errList, err)
	} else {
		resource.String = &String
		return nil
	}

	// Int64
	var Int64 int64

	if err := json.Unmarshal(raw, &Int64); err != nil {
		errList = append(errList, err)
	} else {
		resource.Int64 = &Int64
		return nil
	}

	if len(errList) != 0 {
		errs = append(errs, cog.MakeBuildErrors("StringOrInt64", errors.Join(errList...))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `StringOrInt64` objects.
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

// Validate checks all the validation constraints that may be defined on `StringOrInt64` fields for violations and returns them.
func (resource StringOrInt64) Validate() error {
	return nil
}
