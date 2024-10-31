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

type CSVWave struct {
	Labels    *string `json:"labels,omitempty"`
	Name      *string `json:"name,omitempty"`
	TimeStep  *int64  `json:"timeStep,omitempty"`
	ValuesCSV *string `json:"valuesCSV,omitempty"`
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
	// Field "labels"
	if fields["labels"] != nil {
		if string(fields["labels"]) != "null" {
			if err := json.Unmarshal(fields["labels"], &resource.Labels); err != nil {
				errs = append(errs, cog.MakeBuildErrors("labels", err)...)
			}

		}
		delete(fields, "labels")

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
	// Field "timeStep"
	if fields["timeStep"] != nil {
		if string(fields["timeStep"]) != "null" {
			if err := json.Unmarshal(fields["timeStep"], &resource.TimeStep); err != nil {
				errs = append(errs, cog.MakeBuildErrors("timeStep", err)...)
			}

		}
		delete(fields, "timeStep")

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
	if resource.Labels == nil && other.Labels != nil || resource.Labels != nil && other.Labels == nil {
		return false
	}

	if resource.Labels != nil {
		if *resource.Labels != *other.Labels {
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
	if resource.TimeStep == nil && other.TimeStep != nil || resource.TimeStep != nil && other.TimeStep == nil {
		return false
	}

	if resource.TimeStep != nil {
		if *resource.TimeStep != *other.TimeStep {
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

	return true
}

// Validate checks all the validation constraints that may be defined on `CSVWave` fields for violations and returns them.
func (resource CSVWave) Validate() error {
	return nil
}

type NodesQuery struct {
	Count *int64 `json:"count,omitempty"`
	Seed  *int64 `json:"seed,omitempty"`
	// Possible enum values:
	//  - `"random"`
	//  - `"random edges"`
	//  - `"response_medium"`
	//  - `"response_small"`
	//  - `"feature_showcase"`
	Type *NodesQueryType `json:"type,omitempty"`
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
	// Field "count"
	if fields["count"] != nil {
		if string(fields["count"]) != "null" {
			if err := json.Unmarshal(fields["count"], &resource.Count); err != nil {
				errs = append(errs, cog.MakeBuildErrors("count", err)...)
			}

		}
		delete(fields, "count")

	}
	// Field "seed"
	if fields["seed"] != nil {
		if string(fields["seed"]) != "null" {
			if err := json.Unmarshal(fields["seed"], &resource.Seed); err != nil {
				errs = append(errs, cog.MakeBuildErrors("seed", err)...)
			}

		}
		delete(fields, "seed")

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
	if resource.Count == nil && other.Count != nil || resource.Count != nil && other.Count == nil {
		return false
	}

	if resource.Count != nil {
		if *resource.Count != *other.Count {
			return false
		}
	}
	if resource.Seed == nil && other.Seed != nil || resource.Seed != nil && other.Seed == nil {
		return false
	}

	if resource.Seed != nil {
		if *resource.Seed != *other.Seed {
			return false
		}
	}
	if resource.Type == nil && other.Type != nil || resource.Type != nil && other.Type == nil {
		return false
	}

	if resource.Type != nil {
		if *resource.Type != *other.Type {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `NodesQuery` fields for violations and returns them.
func (resource NodesQuery) Validate() error {
	return nil
}

type PulseWaveQuery struct {
	OffCount *int64   `json:"offCount,omitempty"`
	OffValue *float64 `json:"offValue,omitempty"`
	OnCount  *int64   `json:"onCount,omitempty"`
	OnValue  *float64 `json:"onValue,omitempty"`
	TimeStep *int64   `json:"timeStep,omitempty"`
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
	// Field "offCount"
	if fields["offCount"] != nil {
		if string(fields["offCount"]) != "null" {
			if err := json.Unmarshal(fields["offCount"], &resource.OffCount); err != nil {
				errs = append(errs, cog.MakeBuildErrors("offCount", err)...)
			}

		}
		delete(fields, "offCount")

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
	// Field "onCount"
	if fields["onCount"] != nil {
		if string(fields["onCount"]) != "null" {
			if err := json.Unmarshal(fields["onCount"], &resource.OnCount); err != nil {
				errs = append(errs, cog.MakeBuildErrors("onCount", err)...)
			}

		}
		delete(fields, "onCount")

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
	// Field "timeStep"
	if fields["timeStep"] != nil {
		if string(fields["timeStep"]) != "null" {
			if err := json.Unmarshal(fields["timeStep"], &resource.TimeStep); err != nil {
				errs = append(errs, cog.MakeBuildErrors("timeStep", err)...)
			}

		}
		delete(fields, "timeStep")

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
	if resource.OffCount == nil && other.OffCount != nil || resource.OffCount != nil && other.OffCount == nil {
		return false
	}

	if resource.OffCount != nil {
		if *resource.OffCount != *other.OffCount {
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
	if resource.OnCount == nil && other.OnCount != nil || resource.OnCount != nil && other.OnCount == nil {
		return false
	}

	if resource.OnCount != nil {
		if *resource.OnCount != *other.OnCount {
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
	if resource.TimeStep == nil && other.TimeStep != nil || resource.TimeStep != nil && other.TimeStep == nil {
		return false
	}

	if resource.TimeStep != nil {
		if *resource.TimeStep != *other.TimeStep {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `PulseWaveQuery` fields for violations and returns them.
func (resource PulseWaveQuery) Validate() error {
	return nil
}

type ResultAssertions struct {
	// Maximum frame count
	MaxFrames *int64 `json:"maxFrames,omitempty"`
	// Type asserts that the frame matches a known type structure.
	// Possible enum values:
	//  - `""`
	//  - `"timeseries-wide"`
	//  - `"timeseries-long"`
	//  - `"timeseries-many"`
	//  - `"timeseries-multi"`
	//  - `"directory-listing"`
	//  - `"table"`
	//  - `"numeric-wide"`
	//  - `"numeric-multi"`
	//  - `"numeric-long"`
	//  - `"log-lines"`
	Type *ResultAssertionsType `json:"type,omitempty"`
	// TypeVersion is the version of the Type property. Versions greater than 0.0 correspond to the dataplane
	// contract documentation https://grafana.github.io/dataplane/contract/.
	TypeVersion []int64 `json:"typeVersion"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ResultAssertions` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ResultAssertions) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "maxFrames"
	if fields["maxFrames"] != nil {
		if string(fields["maxFrames"]) != "null" {
			if err := json.Unmarshal(fields["maxFrames"], &resource.MaxFrames); err != nil {
				errs = append(errs, cog.MakeBuildErrors("maxFrames", err)...)
			}

		}
		delete(fields, "maxFrames")

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
	// Field "typeVersion"
	if fields["typeVersion"] != nil {
		if string(fields["typeVersion"]) != "null" {

			if err := json.Unmarshal(fields["typeVersion"], &resource.TypeVersion); err != nil {
				errs = append(errs, cog.MakeBuildErrors("typeVersion", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("typeVersion", errors.New("required field is null"))...)

		}
		delete(fields, "typeVersion")
	} else {
		errs = append(errs, cog.MakeBuildErrors("typeVersion", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ResultAssertions", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ResultAssertions` objects.
func (resource ResultAssertions) Equals(other ResultAssertions) bool {
	if resource.MaxFrames == nil && other.MaxFrames != nil || resource.MaxFrames != nil && other.MaxFrames == nil {
		return false
	}

	if resource.MaxFrames != nil {
		if *resource.MaxFrames != *other.MaxFrames {
			return false
		}
	}
	if resource.Type == nil && other.Type != nil || resource.Type != nil && other.Type == nil {
		return false
	}

	if resource.Type != nil {
		if *resource.Type != *other.Type {
			return false
		}
	}

	if len(resource.TypeVersion) != len(other.TypeVersion) {
		return false
	}

	for i1 := range resource.TypeVersion {
		if resource.TypeVersion[i1] != other.TypeVersion[i1] {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `ResultAssertions` fields for violations and returns them.
func (resource ResultAssertions) Validate() error {
	return nil
}

type Key struct {
	Tick float64 `json:"tick"`
	Type string  `json:"type"`
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
	if resource.Tick != other.Tick {
		return false
	}
	if resource.Type != other.Type {
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

type SimulationQuery struct {
	Config any   `json:"config,omitempty"`
	Key    Key   `json:"key"`
	Last   *bool `json:"last,omitempty"`
	Stream *bool `json:"stream,omitempty"`
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
	// Field "config"
	if fields["config"] != nil {
		if string(fields["config"]) != "null" {
			if err := json.Unmarshal(fields["config"], &resource.Config); err != nil {
				errs = append(errs, cog.MakeBuildErrors("config", err)...)
			}

		}
		delete(fields, "config")

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
	// Field "last"
	if fields["last"] != nil {
		if string(fields["last"]) != "null" {
			if err := json.Unmarshal(fields["last"], &resource.Last); err != nil {
				errs = append(errs, cog.MakeBuildErrors("last", err)...)
			}

		}
		delete(fields, "last")

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
	// is DeepEqual good enough here?
	if !reflect.DeepEqual(resource.Config, other.Config) {
		return false
	}
	if !resource.Key.Equals(other.Key) {
		return false
	}
	if resource.Last == nil && other.Last != nil || resource.Last != nil && other.Last == nil {
		return false
	}

	if resource.Last != nil {
		if *resource.Last != *other.Last {
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

type StreamingQuery struct {
	Bands  *int64  `json:"bands,omitempty"`
	Noise  float64 `json:"noise"`
	Speed  float64 `json:"speed"`
	Spread float64 `json:"spread"`
	// Possible enum values:
	//  - `"fetch"`
	//  - `"logs"`
	//  - `"signal"`
	//  - `"traces"`
	Type StreamingQueryType `json:"type"`
	Url  *string            `json:"url,omitempty"`
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
	// Field "bands"
	if fields["bands"] != nil {
		if string(fields["bands"]) != "null" {
			if err := json.Unmarshal(fields["bands"], &resource.Bands); err != nil {
				errs = append(errs, cog.MakeBuildErrors("bands", err)...)
			}

		}
		delete(fields, "bands")

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
	if resource.Bands == nil && other.Bands != nil || resource.Bands != nil && other.Bands == nil {
		return false
	}

	if resource.Bands != nil {
		if *resource.Bands != *other.Bands {
			return false
		}
	}
	if resource.Noise != other.Noise {
		return false
	}
	if resource.Speed != other.Speed {
		return false
	}
	if resource.Spread != other.Spread {
		return false
	}
	if resource.Type != other.Type {
		return false
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

type TimeRange struct {
	// From is the start time of the query.
	From string `json:"from"`
	// To is the end time of the query.
	To string `json:"to"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TimeRange` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *TimeRange) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "from"
	if fields["from"] != nil {
		if string(fields["from"]) != "null" {
			if err := json.Unmarshal(fields["from"], &resource.From); err != nil {
				errs = append(errs, cog.MakeBuildErrors("from", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("from", errors.New("required field is null"))...)

		}
		delete(fields, "from")
	} else {
		errs = append(errs, cog.MakeBuildErrors("from", errors.New("required field is missing from input"))...)
	}
	// Field "to"
	if fields["to"] != nil {
		if string(fields["to"]) != "null" {
			if err := json.Unmarshal(fields["to"], &resource.To); err != nil {
				errs = append(errs, cog.MakeBuildErrors("to", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("to", errors.New("required field is null"))...)

		}
		delete(fields, "to")
	} else {
		errs = append(errs, cog.MakeBuildErrors("to", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("TimeRange", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `TimeRange` objects.
func (resource TimeRange) Equals(other TimeRange) bool {
	if resource.From != other.From {
		return false
	}
	if resource.To != other.To {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `TimeRange` fields for violations and returns them.
func (resource TimeRange) Validate() error {
	return nil
}

type USAQuery struct {
	Fields []string `json:"fields,omitempty"`
	Mode   *string  `json:"mode,omitempty"`
	Period *string  `json:"period,omitempty"`
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
	// Field "fields"
	if fields["fields"] != nil {
		if string(fields["fields"]) != "null" {

			if err := json.Unmarshal(fields["fields"], &resource.Fields); err != nil {
				errs = append(errs, cog.MakeBuildErrors("fields", err)...)
			}

		}
		delete(fields, "fields")

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

	if len(resource.Fields) != len(other.Fields) {
		return false
	}

	for i1 := range resource.Fields {
		if resource.Fields[i1] != other.Fields[i1] {
			return false
		}
	}
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

type Dataquery struct {
	Alias *string `json:"alias,omitempty"`
	// Used for live query
	Channel     *string   `json:"channel,omitempty"`
	CsvContent  *string   `json:"csvContent,omitempty"`
	CsvFileName *string   `json:"csvFileName,omitempty"`
	CsvWave     []CSVWave `json:"csvWave,omitempty"`
	// The datasource
	Datasource *dashboard.DataSourceRef `json:"datasource,omitempty"`
	// Drop percentage (the chance we will lose a point 0-100)
	DropPercent *float64 `json:"dropPercent,omitempty"`
	// Possible enum values:
	//  - `"plugin"`
	//  - `"downstream"`
	ErrorSource *DataqueryErrorSource `json:"errorSource,omitempty"`
	// Possible enum values:
	//  - `"frontend_exception"`
	//  - `"frontend_observable"`
	//  - `"server_panic"`
	ErrorType      *DataqueryErrorType `json:"errorType,omitempty"`
	FlamegraphDiff *bool               `json:"flamegraphDiff,omitempty"`
	// true if query is disabled (ie should not be returned to the dashboard)
	// NOTE: this does not always imply that the query should not be executed since
	// the results from a hidden query may be used as the input to other queries (SSE etc)
	Hide *bool `json:"hide,omitempty"`
	// Interval is the suggested duration between time points in a time series query.
	// NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
	// from the interval required to fill a pixels in the visualization
	IntervalMs  *float64 `json:"intervalMs,omitempty"`
	Labels      *string  `json:"labels,omitempty"`
	LevelColumn *bool    `json:"levelColumn,omitempty"`
	Lines       *int64   `json:"lines,omitempty"`
	Max         *float64 `json:"max,omitempty"`
	// MaxDataPoints is the maximum number of data points that should be returned from a time series query.
	// NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
	// from the number of pixels visible in a visualization
	MaxDataPoints *int64          `json:"maxDataPoints,omitempty"`
	Min           *float64        `json:"min,omitempty"`
	Nodes         *NodesQuery     `json:"nodes,omitempty"`
	Noise         *float64        `json:"noise,omitempty"`
	Points        [][]any         `json:"points,omitempty"`
	PulseWave     *PulseWaveQuery `json:"pulseWave,omitempty"`
	// QueryType is an optional identifier for the type of query.
	// It can be used to distinguish different types of queries.
	QueryType       *string `json:"queryType,omitempty"`
	RawFrameContent *string `json:"rawFrameContent,omitempty"`
	// RefID is the unique identifier of the query, set by the frontend call.
	RefId *string `json:"refId,omitempty"`
	// Optionally define expected query result behavior
	ResultAssertions *ResultAssertions `json:"resultAssertions,omitempty"`
	// Possible enum values:
	//  - `"annotations"`
	//  - `"arrow"`
	//  - `"csv_content"`
	//  - `"csv_file"`
	//  - `"csv_metric_values"`
	//  - `"datapoints_outside_range"`
	//  - `"error_with_source"`
	//  - `"exponential_heatmap_bucket_data"`
	//  - `"flame_graph"`
	//  - `"grafana_api"`
	//  - `"linear_heatmap_bucket_data"`
	//  - `"live"`
	//  - `"logs"`
	//  - `"manual_entry"`
	//  - `"no_data_points"`
	//  - `"node_graph"`
	//  - `"predictable_csv_wave"`
	//  - `"predictable_pulse"`
	//  - `"random_walk"`
	//  - `"random_walk_table"`
	//  - `"random_walk_with_error"`
	//  - `"raw_frame"`
	//  - `"server_error_500"`
	//  - `"simulation"`
	//  - `"slow_query"`
	//  - `"streaming_client"`
	//  - `"table_static"`
	//  - `"trace"`
	//  - `"usa"`
	//  - `"variables-query"`
	ScenarioId  *DataqueryScenarioId `json:"scenarioId,omitempty"`
	SeriesCount *int64               `json:"seriesCount,omitempty"`
	Sim         *SimulationQuery     `json:"sim,omitempty"`
	SpanCount   *int64               `json:"spanCount,omitempty"`
	Spread      *float64             `json:"spread,omitempty"`
	StartValue  *float64             `json:"startValue,omitempty"`
	Stream      *StreamingQuery      `json:"stream,omitempty"`
	// common parameter used by many query types
	StringInput *string `json:"stringInput,omitempty"`
	// TimeRange represents the query range
	// NOTE: unlike generic /ds/query, we can now send explicit time values in each query
	// NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
	TimeRange *TimeRange `json:"timeRange,omitempty"`
	Usa       *USAQuery  `json:"usa,omitempty"`
	WithNil   *bool      `json:"withNil,omitempty"`
}

func (resource Dataquery) ImplementsDataqueryVariant() {}

func (resource Dataquery) DataqueryType() string {
	return ""
}

// VariantConfig returns the configuration related to  dataqueries.
// This configuration describes how to unmarshal it, convert it to code, …
func VariantConfig() variants.DataqueryConfig {
	return variants.DataqueryConfig{
		Identifier: "",
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
	// Field "channel"
	if fields["channel"] != nil {
		if string(fields["channel"]) != "null" {
			if err := json.Unmarshal(fields["channel"], &resource.Channel); err != nil {
				errs = append(errs, cog.MakeBuildErrors("channel", err)...)
			}

		}
		delete(fields, "channel")

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
	// Field "csvFileName"
	if fields["csvFileName"] != nil {
		if string(fields["csvFileName"]) != "null" {
			if err := json.Unmarshal(fields["csvFileName"], &resource.CsvFileName); err != nil {
				errs = append(errs, cog.MakeBuildErrors("csvFileName", err)...)
			}

		}
		delete(fields, "csvFileName")

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
	// Field "dropPercent"
	if fields["dropPercent"] != nil {
		if string(fields["dropPercent"]) != "null" {
			if err := json.Unmarshal(fields["dropPercent"], &resource.DropPercent); err != nil {
				errs = append(errs, cog.MakeBuildErrors("dropPercent", err)...)
			}

		}
		delete(fields, "dropPercent")

	}
	// Field "errorSource"
	if fields["errorSource"] != nil {
		if string(fields["errorSource"]) != "null" {
			if err := json.Unmarshal(fields["errorSource"], &resource.ErrorSource); err != nil {
				errs = append(errs, cog.MakeBuildErrors("errorSource", err)...)
			}

		}
		delete(fields, "errorSource")

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
	// Field "flamegraphDiff"
	if fields["flamegraphDiff"] != nil {
		if string(fields["flamegraphDiff"]) != "null" {
			if err := json.Unmarshal(fields["flamegraphDiff"], &resource.FlamegraphDiff); err != nil {
				errs = append(errs, cog.MakeBuildErrors("flamegraphDiff", err)...)
			}

		}
		delete(fields, "flamegraphDiff")

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
	// Field "intervalMs"
	if fields["intervalMs"] != nil {
		if string(fields["intervalMs"]) != "null" {
			if err := json.Unmarshal(fields["intervalMs"], &resource.IntervalMs); err != nil {
				errs = append(errs, cog.MakeBuildErrors("intervalMs", err)...)
			}

		}
		delete(fields, "intervalMs")

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
	// Field "levelColumn"
	if fields["levelColumn"] != nil {
		if string(fields["levelColumn"]) != "null" {
			if err := json.Unmarshal(fields["levelColumn"], &resource.LevelColumn); err != nil {
				errs = append(errs, cog.MakeBuildErrors("levelColumn", err)...)
			}

		}
		delete(fields, "levelColumn")

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
	// Field "max"
	if fields["max"] != nil {
		if string(fields["max"]) != "null" {
			if err := json.Unmarshal(fields["max"], &resource.Max); err != nil {
				errs = append(errs, cog.MakeBuildErrors("max", err)...)
			}

		}
		delete(fields, "max")

	}
	// Field "maxDataPoints"
	if fields["maxDataPoints"] != nil {
		if string(fields["maxDataPoints"]) != "null" {
			if err := json.Unmarshal(fields["maxDataPoints"], &resource.MaxDataPoints); err != nil {
				errs = append(errs, cog.MakeBuildErrors("maxDataPoints", err)...)
			}

		}
		delete(fields, "maxDataPoints")

	}
	// Field "min"
	if fields["min"] != nil {
		if string(fields["min"]) != "null" {
			if err := json.Unmarshal(fields["min"], &resource.Min); err != nil {
				errs = append(errs, cog.MakeBuildErrors("min", err)...)
			}

		}
		delete(fields, "min")

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
	// Field "noise"
	if fields["noise"] != nil {
		if string(fields["noise"]) != "null" {
			if err := json.Unmarshal(fields["noise"], &resource.Noise); err != nil {
				errs = append(errs, cog.MakeBuildErrors("noise", err)...)
			}

		}
		delete(fields, "noise")

	}
	// Field "points"
	if fields["points"] != nil {
		if string(fields["points"]) != "null" {

			if err := json.Unmarshal(fields["points"], &resource.Points); err != nil {
				errs = append(errs, cog.MakeBuildErrors("points", err)...)
			}

		}
		delete(fields, "points")

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
	// Field "queryType"
	if fields["queryType"] != nil {
		if string(fields["queryType"]) != "null" {
			if err := json.Unmarshal(fields["queryType"], &resource.QueryType); err != nil {
				errs = append(errs, cog.MakeBuildErrors("queryType", err)...)
			}

		}
		delete(fields, "queryType")

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
	// Field "refId"
	if fields["refId"] != nil {
		if string(fields["refId"]) != "null" {
			if err := json.Unmarshal(fields["refId"], &resource.RefId); err != nil {
				errs = append(errs, cog.MakeBuildErrors("refId", err)...)
			}

		}
		delete(fields, "refId")

	}
	// Field "resultAssertions"
	if fields["resultAssertions"] != nil {
		if string(fields["resultAssertions"]) != "null" {

			resource.ResultAssertions = &ResultAssertions{}
			if err := resource.ResultAssertions.UnmarshalJSONStrict(fields["resultAssertions"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("resultAssertions", err)...)
			}

		}
		delete(fields, "resultAssertions")

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
	// Field "seriesCount"
	if fields["seriesCount"] != nil {
		if string(fields["seriesCount"]) != "null" {
			if err := json.Unmarshal(fields["seriesCount"], &resource.SeriesCount); err != nil {
				errs = append(errs, cog.MakeBuildErrors("seriesCount", err)...)
			}

		}
		delete(fields, "seriesCount")

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
	// Field "spanCount"
	if fields["spanCount"] != nil {
		if string(fields["spanCount"]) != "null" {
			if err := json.Unmarshal(fields["spanCount"], &resource.SpanCount); err != nil {
				errs = append(errs, cog.MakeBuildErrors("spanCount", err)...)
			}

		}
		delete(fields, "spanCount")

	}
	// Field "spread"
	if fields["spread"] != nil {
		if string(fields["spread"]) != "null" {
			if err := json.Unmarshal(fields["spread"], &resource.Spread); err != nil {
				errs = append(errs, cog.MakeBuildErrors("spread", err)...)
			}

		}
		delete(fields, "spread")

	}
	// Field "startValue"
	if fields["startValue"] != nil {
		if string(fields["startValue"]) != "null" {
			if err := json.Unmarshal(fields["startValue"], &resource.StartValue); err != nil {
				errs = append(errs, cog.MakeBuildErrors("startValue", err)...)
			}

		}
		delete(fields, "startValue")

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
	// Field "stringInput"
	if fields["stringInput"] != nil {
		if string(fields["stringInput"]) != "null" {
			if err := json.Unmarshal(fields["stringInput"], &resource.StringInput); err != nil {
				errs = append(errs, cog.MakeBuildErrors("stringInput", err)...)
			}

		}
		delete(fields, "stringInput")

	}
	// Field "timeRange"
	if fields["timeRange"] != nil {
		if string(fields["timeRange"]) != "null" {

			resource.TimeRange = &TimeRange{}
			if err := resource.TimeRange.UnmarshalJSONStrict(fields["timeRange"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("timeRange", err)...)
			}

		}
		delete(fields, "timeRange")

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
	// Field "withNil"
	if fields["withNil"] != nil {
		if string(fields["withNil"]) != "null" {
			if err := json.Unmarshal(fields["withNil"], &resource.WithNil); err != nil {
				errs = append(errs, cog.MakeBuildErrors("withNil", err)...)
			}

		}
		delete(fields, "withNil")

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
	if resource.Channel == nil && other.Channel != nil || resource.Channel != nil && other.Channel == nil {
		return false
	}

	if resource.Channel != nil {
		if *resource.Channel != *other.Channel {
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
	if resource.CsvFileName == nil && other.CsvFileName != nil || resource.CsvFileName != nil && other.CsvFileName == nil {
		return false
	}

	if resource.CsvFileName != nil {
		if *resource.CsvFileName != *other.CsvFileName {
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
	if resource.Datasource == nil && other.Datasource != nil || resource.Datasource != nil && other.Datasource == nil {
		return false
	}

	if resource.Datasource != nil {
		if !resource.Datasource.Equals(*other.Datasource) {
			return false
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
	if resource.ErrorSource == nil && other.ErrorSource != nil || resource.ErrorSource != nil && other.ErrorSource == nil {
		return false
	}

	if resource.ErrorSource != nil {
		if *resource.ErrorSource != *other.ErrorSource {
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
	if resource.FlamegraphDiff == nil && other.FlamegraphDiff != nil || resource.FlamegraphDiff != nil && other.FlamegraphDiff == nil {
		return false
	}

	if resource.FlamegraphDiff != nil {
		if *resource.FlamegraphDiff != *other.FlamegraphDiff {
			return false
		}
	}
	if resource.Hide == nil && other.Hide != nil || resource.Hide != nil && other.Hide == nil {
		return false
	}

	if resource.Hide != nil {
		if *resource.Hide != *other.Hide {
			return false
		}
	}
	if resource.IntervalMs == nil && other.IntervalMs != nil || resource.IntervalMs != nil && other.IntervalMs == nil {
		return false
	}

	if resource.IntervalMs != nil {
		if *resource.IntervalMs != *other.IntervalMs {
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
	if resource.LevelColumn == nil && other.LevelColumn != nil || resource.LevelColumn != nil && other.LevelColumn == nil {
		return false
	}

	if resource.LevelColumn != nil {
		if *resource.LevelColumn != *other.LevelColumn {
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
	if resource.Max == nil && other.Max != nil || resource.Max != nil && other.Max == nil {
		return false
	}

	if resource.Max != nil {
		if *resource.Max != *other.Max {
			return false
		}
	}
	if resource.MaxDataPoints == nil && other.MaxDataPoints != nil || resource.MaxDataPoints != nil && other.MaxDataPoints == nil {
		return false
	}

	if resource.MaxDataPoints != nil {
		if *resource.MaxDataPoints != *other.MaxDataPoints {
			return false
		}
	}
	if resource.Min == nil && other.Min != nil || resource.Min != nil && other.Min == nil {
		return false
	}

	if resource.Min != nil {
		if *resource.Min != *other.Min {
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
	if resource.Noise == nil && other.Noise != nil || resource.Noise != nil && other.Noise == nil {
		return false
	}

	if resource.Noise != nil {
		if *resource.Noise != *other.Noise {
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
			// is DeepEqual good enough here?
			if !reflect.DeepEqual(resource.Points[i1][i2], other.Points[i1][i2]) {
				return false
			}
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
	if resource.QueryType == nil && other.QueryType != nil || resource.QueryType != nil && other.QueryType == nil {
		return false
	}

	if resource.QueryType != nil {
		if *resource.QueryType != *other.QueryType {
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
	if resource.RefId == nil && other.RefId != nil || resource.RefId != nil && other.RefId == nil {
		return false
	}

	if resource.RefId != nil {
		if *resource.RefId != *other.RefId {
			return false
		}
	}
	if resource.ResultAssertions == nil && other.ResultAssertions != nil || resource.ResultAssertions != nil && other.ResultAssertions == nil {
		return false
	}

	if resource.ResultAssertions != nil {
		if !resource.ResultAssertions.Equals(*other.ResultAssertions) {
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
	if resource.SeriesCount == nil && other.SeriesCount != nil || resource.SeriesCount != nil && other.SeriesCount == nil {
		return false
	}

	if resource.SeriesCount != nil {
		if *resource.SeriesCount != *other.SeriesCount {
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
	if resource.SpanCount == nil && other.SpanCount != nil || resource.SpanCount != nil && other.SpanCount == nil {
		return false
	}

	if resource.SpanCount != nil {
		if *resource.SpanCount != *other.SpanCount {
			return false
		}
	}
	if resource.Spread == nil && other.Spread != nil || resource.Spread != nil && other.Spread == nil {
		return false
	}

	if resource.Spread != nil {
		if *resource.Spread != *other.Spread {
			return false
		}
	}
	if resource.StartValue == nil && other.StartValue != nil || resource.StartValue != nil && other.StartValue == nil {
		return false
	}

	if resource.StartValue != nil {
		if *resource.StartValue != *other.StartValue {
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
	if resource.StringInput == nil && other.StringInput != nil || resource.StringInput != nil && other.StringInput == nil {
		return false
	}

	if resource.StringInput != nil {
		if *resource.StringInput != *other.StringInput {
			return false
		}
	}
	if resource.TimeRange == nil && other.TimeRange != nil || resource.TimeRange != nil && other.TimeRange == nil {
		return false
	}

	if resource.TimeRange != nil {
		if !resource.TimeRange.Equals(*other.TimeRange) {
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
	if resource.WithNil == nil && other.WithNil != nil || resource.WithNil != nil && other.WithNil == nil {
		return false
	}

	if resource.WithNil != nil {
		if *resource.WithNil != *other.WithNil {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `Dataquery` fields for violations and returns them.
func (resource Dataquery) Validate() error {
	var errs cog.BuildErrors

	for i1 := range resource.CsvWave {
		if err := resource.CsvWave[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("csvWave["+strconv.Itoa(i1)+"]", err)...)
		}
	}
	if resource.Datasource != nil {
		if err := resource.Datasource.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("datasource", err)...)
		}
	}
	if resource.Nodes != nil {
		if err := resource.Nodes.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("nodes", err)...)
		}
	}
	if resource.PulseWave != nil {
		if err := resource.PulseWave.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("pulseWave", err)...)
		}
	}
	if resource.ResultAssertions != nil {
		if err := resource.ResultAssertions.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("resultAssertions", err)...)
		}
	}
	if resource.Sim != nil {
		if err := resource.Sim.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("sim", err)...)
		}
	}
	if resource.Stream != nil {
		if err := resource.Stream.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("stream", err)...)
		}
	}
	if resource.TimeRange != nil {
		if err := resource.TimeRange.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("timeRange", err)...)
		}
	}
	if resource.Usa != nil {
		if err := resource.Usa.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("usa", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type NodesQueryType string

const (
	NodesQueryTypeRandom          NodesQueryType = "random"
	NodesQueryTypeRandomEdges     NodesQueryType = "random edges"
	NodesQueryTypeResponseMedium  NodesQueryType = "response_medium"
	NodesQueryTypeResponseSmall   NodesQueryType = "response_small"
	NodesQueryTypeFeatureShowcase NodesQueryType = "feature_showcase"
)

type ResultAssertionsType string

const (
	ResultAssertionsTypeNone             ResultAssertionsType = ""
	ResultAssertionsTypeTimeseriesWide   ResultAssertionsType = "timeseries-wide"
	ResultAssertionsTypeTimeseriesLong   ResultAssertionsType = "timeseries-long"
	ResultAssertionsTypeTimeseriesMany   ResultAssertionsType = "timeseries-many"
	ResultAssertionsTypeTimeseriesMulti  ResultAssertionsType = "timeseries-multi"
	ResultAssertionsTypeDirectoryListing ResultAssertionsType = "directory-listing"
	ResultAssertionsTypeTable            ResultAssertionsType = "table"
	ResultAssertionsTypeNumericWide      ResultAssertionsType = "numeric-wide"
	ResultAssertionsTypeNumericMulti     ResultAssertionsType = "numeric-multi"
	ResultAssertionsTypeNumericLong      ResultAssertionsType = "numeric-long"
	ResultAssertionsTypeLogLines         ResultAssertionsType = "log-lines"
)

type StreamingQueryType string

const (
	StreamingQueryTypeFetch  StreamingQueryType = "fetch"
	StreamingQueryTypeLogs   StreamingQueryType = "logs"
	StreamingQueryTypeSignal StreamingQueryType = "signal"
	StreamingQueryTypeTraces StreamingQueryType = "traces"
)

type DataqueryErrorSource string

const (
	DataqueryErrorSourcePlugin     DataqueryErrorSource = "plugin"
	DataqueryErrorSourceDownstream DataqueryErrorSource = "downstream"
)

type DataqueryErrorType string

const (
	DataqueryErrorTypeFrontendException  DataqueryErrorType = "frontend_exception"
	DataqueryErrorTypeFrontendObservable DataqueryErrorType = "frontend_observable"
	DataqueryErrorTypeServerPanic        DataqueryErrorType = "server_panic"
)

type DataqueryScenarioId string

const (
	DataqueryScenarioIdAnnotations                  DataqueryScenarioId = "annotations"
	DataqueryScenarioIdArrow                        DataqueryScenarioId = "arrow"
	DataqueryScenarioIdCsvContent                   DataqueryScenarioId = "csv_content"
	DataqueryScenarioIdCsvFile                      DataqueryScenarioId = "csv_file"
	DataqueryScenarioIdCsvMetricValues              DataqueryScenarioId = "csv_metric_values"
	DataqueryScenarioIdDatapointsOutsideRange       DataqueryScenarioId = "datapoints_outside_range"
	DataqueryScenarioIdErrorWithSource              DataqueryScenarioId = "error_with_source"
	DataqueryScenarioIdExponentialHeatmapBucketData DataqueryScenarioId = "exponential_heatmap_bucket_data"
	DataqueryScenarioIdFlameGraph                   DataqueryScenarioId = "flame_graph"
	DataqueryScenarioIdGrafanaApi                   DataqueryScenarioId = "grafana_api"
	DataqueryScenarioIdLinearHeatmapBucketData      DataqueryScenarioId = "linear_heatmap_bucket_data"
	DataqueryScenarioIdLive                         DataqueryScenarioId = "live"
	DataqueryScenarioIdLogs                         DataqueryScenarioId = "logs"
	DataqueryScenarioIdManualEntry                  DataqueryScenarioId = "manual_entry"
	DataqueryScenarioIdNoDataPoints                 DataqueryScenarioId = "no_data_points"
	DataqueryScenarioIdNodeGraph                    DataqueryScenarioId = "node_graph"
	DataqueryScenarioIdPredictableCsvWave           DataqueryScenarioId = "predictable_csv_wave"
	DataqueryScenarioIdPredictablePulse             DataqueryScenarioId = "predictable_pulse"
	DataqueryScenarioIdRandomWalk                   DataqueryScenarioId = "random_walk"
	DataqueryScenarioIdRandomWalkTable              DataqueryScenarioId = "random_walk_table"
	DataqueryScenarioIdRandomWalkWithError          DataqueryScenarioId = "random_walk_with_error"
	DataqueryScenarioIdRawFrame                     DataqueryScenarioId = "raw_frame"
	DataqueryScenarioIdServerError500               DataqueryScenarioId = "server_error_500"
	DataqueryScenarioIdSimulation                   DataqueryScenarioId = "simulation"
	DataqueryScenarioIdSlowQuery                    DataqueryScenarioId = "slow_query"
	DataqueryScenarioIdStreamingClient              DataqueryScenarioId = "streaming_client"
	DataqueryScenarioIdTableStatic                  DataqueryScenarioId = "table_static"
	DataqueryScenarioIdTrace                        DataqueryScenarioId = "trace"
	DataqueryScenarioIdUsa                          DataqueryScenarioId = "usa"
	DataqueryScenarioIdVariablesQuery               DataqueryScenarioId = "variables-query"
)
