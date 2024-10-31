// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

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

type BucketAggregation = DateHistogramOrHistogramOrTermsOrFiltersOrGeoHashGridOrNested

type MetricAggregation = CountOrMovingAverageOrDerivativeOrCumulativeSumOrBucketScriptOrSerialDiffOrRawDataOrRawDocumentOrUniqueCountOrPercentilesOrExtendedStatsOrMinOrMaxOrSumOrAverageOrMovingFunctionOrLogsOrRateOrTopMetrics

type BucketAggregationType string

const (
	BucketAggregationTypeTerms         BucketAggregationType = "terms"
	BucketAggregationTypeFilters       BucketAggregationType = "filters"
	BucketAggregationTypeGeohashGrid   BucketAggregationType = "geohash_grid"
	BucketAggregationTypeDateHistogram BucketAggregationType = "date_histogram"
	BucketAggregationTypeHistogram     BucketAggregationType = "histogram"
	BucketAggregationTypeNested        BucketAggregationType = "nested"
)

type BaseBucketAggregation struct {
	Id       string                `json:"id"`
	Type     BucketAggregationType `json:"type"`
	Settings any                   `json:"settings,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `BaseBucketAggregation` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *BaseBucketAggregation) UnmarshalJSONStrict(raw []byte) error {
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
	// Field "settings"
	if fields["settings"] != nil {
		if string(fields["settings"]) != "null" {
			if err := json.Unmarshal(fields["settings"], &resource.Settings); err != nil {
				errs = append(errs, cog.MakeBuildErrors("settings", err)...)
			}

		}
		delete(fields, "settings")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("BaseBucketAggregation", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `BaseBucketAggregation` objects.
func (resource BaseBucketAggregation) Equals(other BaseBucketAggregation) bool {
	if resource.Id != other.Id {
		return false
	}
	if resource.Type != other.Type {
		return false
	}
	// is DeepEqual good enough here?
	if !reflect.DeepEqual(resource.Settings, other.Settings) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `BaseBucketAggregation` fields for violations and returns them.
func (resource BaseBucketAggregation) Validate() error {
	return nil
}

type BucketAggregationWithField struct {
	Field    *string               `json:"field,omitempty"`
	Id       string                `json:"id"`
	Type     BucketAggregationType `json:"type"`
	Settings any                   `json:"settings,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `BucketAggregationWithField` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *BucketAggregationWithField) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "field"
	if fields["field"] != nil {
		if string(fields["field"]) != "null" {
			if err := json.Unmarshal(fields["field"], &resource.Field); err != nil {
				errs = append(errs, cog.MakeBuildErrors("field", err)...)
			}

		}
		delete(fields, "field")

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
	// Field "settings"
	if fields["settings"] != nil {
		if string(fields["settings"]) != "null" {
			if err := json.Unmarshal(fields["settings"], &resource.Settings); err != nil {
				errs = append(errs, cog.MakeBuildErrors("settings", err)...)
			}

		}
		delete(fields, "settings")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("BucketAggregationWithField", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `BucketAggregationWithField` objects.
func (resource BucketAggregationWithField) Equals(other BucketAggregationWithField) bool {
	if resource.Field == nil && other.Field != nil || resource.Field != nil && other.Field == nil {
		return false
	}

	if resource.Field != nil {
		if *resource.Field != *other.Field {
			return false
		}
	}
	if resource.Id != other.Id {
		return false
	}
	if resource.Type != other.Type {
		return false
	}
	// is DeepEqual good enough here?
	if !reflect.DeepEqual(resource.Settings, other.Settings) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `BucketAggregationWithField` fields for violations and returns them.
func (resource BucketAggregationWithField) Validate() error {
	return nil
}

type DateHistogram struct {
	Field    *string                             `json:"field,omitempty"`
	Id       string                              `json:"id"`
	Type     string                              `json:"type"`
	Settings *ElasticsearchDateHistogramSettings `json:"settings,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `DateHistogram` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *DateHistogram) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "field"
	if fields["field"] != nil {
		if string(fields["field"]) != "null" {
			if err := json.Unmarshal(fields["field"], &resource.Field); err != nil {
				errs = append(errs, cog.MakeBuildErrors("field", err)...)
			}

		}
		delete(fields, "field")

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
	// Field "settings"
	if fields["settings"] != nil {
		if string(fields["settings"]) != "null" {

			resource.Settings = &ElasticsearchDateHistogramSettings{}
			if err := resource.Settings.UnmarshalJSONStrict(fields["settings"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("settings", err)...)
			}

		}
		delete(fields, "settings")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("DateHistogram", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `DateHistogram` objects.
func (resource DateHistogram) Equals(other DateHistogram) bool {
	if resource.Field == nil && other.Field != nil || resource.Field != nil && other.Field == nil {
		return false
	}

	if resource.Field != nil {
		if *resource.Field != *other.Field {
			return false
		}
	}
	if resource.Id != other.Id {
		return false
	}
	if resource.Type != other.Type {
		return false
	}
	if resource.Settings == nil && other.Settings != nil || resource.Settings != nil && other.Settings == nil {
		return false
	}

	if resource.Settings != nil {
		if !resource.Settings.Equals(*other.Settings) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `DateHistogram` fields for violations and returns them.
func (resource DateHistogram) Validate() error {
	var errs cog.BuildErrors
	if !(resource.Type == "date_histogram") {
		errs = append(errs, cog.MakeBuildErrors(
			"type",
			errors.New("must be == date_histogram"),
		)...)
	}
	if resource.Settings != nil {
		if err := resource.Settings.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("settings", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type DateHistogramSettings struct {
	Interval    *string `json:"interval,omitempty"`
	MinDocCount *string `json:"min_doc_count,omitempty"`
	TrimEdges   *string `json:"trimEdges,omitempty"`
	Offset      *string `json:"offset,omitempty"`
	TimeZone    *string `json:"timeZone,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `DateHistogramSettings` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *DateHistogramSettings) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "interval"
	if fields["interval"] != nil {
		if string(fields["interval"]) != "null" {
			if err := json.Unmarshal(fields["interval"], &resource.Interval); err != nil {
				errs = append(errs, cog.MakeBuildErrors("interval", err)...)
			}

		}
		delete(fields, "interval")

	}
	// Field "min_doc_count"
	if fields["min_doc_count"] != nil {
		if string(fields["min_doc_count"]) != "null" {
			if err := json.Unmarshal(fields["min_doc_count"], &resource.MinDocCount); err != nil {
				errs = append(errs, cog.MakeBuildErrors("min_doc_count", err)...)
			}

		}
		delete(fields, "min_doc_count")

	}
	// Field "trimEdges"
	if fields["trimEdges"] != nil {
		if string(fields["trimEdges"]) != "null" {
			if err := json.Unmarshal(fields["trimEdges"], &resource.TrimEdges); err != nil {
				errs = append(errs, cog.MakeBuildErrors("trimEdges", err)...)
			}

		}
		delete(fields, "trimEdges")

	}
	// Field "offset"
	if fields["offset"] != nil {
		if string(fields["offset"]) != "null" {
			if err := json.Unmarshal(fields["offset"], &resource.Offset); err != nil {
				errs = append(errs, cog.MakeBuildErrors("offset", err)...)
			}

		}
		delete(fields, "offset")

	}
	// Field "timeZone"
	if fields["timeZone"] != nil {
		if string(fields["timeZone"]) != "null" {
			if err := json.Unmarshal(fields["timeZone"], &resource.TimeZone); err != nil {
				errs = append(errs, cog.MakeBuildErrors("timeZone", err)...)
			}

		}
		delete(fields, "timeZone")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("DateHistogramSettings", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `DateHistogramSettings` objects.
func (resource DateHistogramSettings) Equals(other DateHistogramSettings) bool {
	if resource.Interval == nil && other.Interval != nil || resource.Interval != nil && other.Interval == nil {
		return false
	}

	if resource.Interval != nil {
		if *resource.Interval != *other.Interval {
			return false
		}
	}
	if resource.MinDocCount == nil && other.MinDocCount != nil || resource.MinDocCount != nil && other.MinDocCount == nil {
		return false
	}

	if resource.MinDocCount != nil {
		if *resource.MinDocCount != *other.MinDocCount {
			return false
		}
	}
	if resource.TrimEdges == nil && other.TrimEdges != nil || resource.TrimEdges != nil && other.TrimEdges == nil {
		return false
	}

	if resource.TrimEdges != nil {
		if *resource.TrimEdges != *other.TrimEdges {
			return false
		}
	}
	if resource.Offset == nil && other.Offset != nil || resource.Offset != nil && other.Offset == nil {
		return false
	}

	if resource.Offset != nil {
		if *resource.Offset != *other.Offset {
			return false
		}
	}
	if resource.TimeZone == nil && other.TimeZone != nil || resource.TimeZone != nil && other.TimeZone == nil {
		return false
	}

	if resource.TimeZone != nil {
		if *resource.TimeZone != *other.TimeZone {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `DateHistogramSettings` fields for violations and returns them.
func (resource DateHistogramSettings) Validate() error {
	return nil
}

type Histogram struct {
	Field    *string                         `json:"field,omitempty"`
	Id       string                          `json:"id"`
	Type     string                          `json:"type"`
	Settings *ElasticsearchHistogramSettings `json:"settings,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Histogram` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *Histogram) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "field"
	if fields["field"] != nil {
		if string(fields["field"]) != "null" {
			if err := json.Unmarshal(fields["field"], &resource.Field); err != nil {
				errs = append(errs, cog.MakeBuildErrors("field", err)...)
			}

		}
		delete(fields, "field")

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
	// Field "settings"
	if fields["settings"] != nil {
		if string(fields["settings"]) != "null" {

			resource.Settings = &ElasticsearchHistogramSettings{}
			if err := resource.Settings.UnmarshalJSONStrict(fields["settings"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("settings", err)...)
			}

		}
		delete(fields, "settings")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("Histogram", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `Histogram` objects.
func (resource Histogram) Equals(other Histogram) bool {
	if resource.Field == nil && other.Field != nil || resource.Field != nil && other.Field == nil {
		return false
	}

	if resource.Field != nil {
		if *resource.Field != *other.Field {
			return false
		}
	}
	if resource.Id != other.Id {
		return false
	}
	if resource.Type != other.Type {
		return false
	}
	if resource.Settings == nil && other.Settings != nil || resource.Settings != nil && other.Settings == nil {
		return false
	}

	if resource.Settings != nil {
		if !resource.Settings.Equals(*other.Settings) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `Histogram` fields for violations and returns them.
func (resource Histogram) Validate() error {
	var errs cog.BuildErrors
	if !(resource.Type == "histogram") {
		errs = append(errs, cog.MakeBuildErrors(
			"type",
			errors.New("must be == histogram"),
		)...)
	}
	if resource.Settings != nil {
		if err := resource.Settings.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("settings", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type HistogramSettings struct {
	Interval    *string `json:"interval,omitempty"`
	MinDocCount *string `json:"min_doc_count,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `HistogramSettings` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *HistogramSettings) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "interval"
	if fields["interval"] != nil {
		if string(fields["interval"]) != "null" {
			if err := json.Unmarshal(fields["interval"], &resource.Interval); err != nil {
				errs = append(errs, cog.MakeBuildErrors("interval", err)...)
			}

		}
		delete(fields, "interval")

	}
	// Field "min_doc_count"
	if fields["min_doc_count"] != nil {
		if string(fields["min_doc_count"]) != "null" {
			if err := json.Unmarshal(fields["min_doc_count"], &resource.MinDocCount); err != nil {
				errs = append(errs, cog.MakeBuildErrors("min_doc_count", err)...)
			}

		}
		delete(fields, "min_doc_count")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("HistogramSettings", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `HistogramSettings` objects.
func (resource HistogramSettings) Equals(other HistogramSettings) bool {
	if resource.Interval == nil && other.Interval != nil || resource.Interval != nil && other.Interval == nil {
		return false
	}

	if resource.Interval != nil {
		if *resource.Interval != *other.Interval {
			return false
		}
	}
	if resource.MinDocCount == nil && other.MinDocCount != nil || resource.MinDocCount != nil && other.MinDocCount == nil {
		return false
	}

	if resource.MinDocCount != nil {
		if *resource.MinDocCount != *other.MinDocCount {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `HistogramSettings` fields for violations and returns them.
func (resource HistogramSettings) Validate() error {
	return nil
}

type TermsOrder string

const (
	TermsOrderDesc TermsOrder = "desc"
	TermsOrderAsc  TermsOrder = "asc"
)

type Nested struct {
	Field    *string `json:"field,omitempty"`
	Id       string  `json:"id"`
	Type     string  `json:"type"`
	Settings any     `json:"settings,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Nested` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *Nested) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "field"
	if fields["field"] != nil {
		if string(fields["field"]) != "null" {
			if err := json.Unmarshal(fields["field"], &resource.Field); err != nil {
				errs = append(errs, cog.MakeBuildErrors("field", err)...)
			}

		}
		delete(fields, "field")

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
	// Field "settings"
	if fields["settings"] != nil {
		if string(fields["settings"]) != "null" {
			if err := json.Unmarshal(fields["settings"], &resource.Settings); err != nil {
				errs = append(errs, cog.MakeBuildErrors("settings", err)...)
			}

		}
		delete(fields, "settings")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("Nested", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `Nested` objects.
func (resource Nested) Equals(other Nested) bool {
	if resource.Field == nil && other.Field != nil || resource.Field != nil && other.Field == nil {
		return false
	}

	if resource.Field != nil {
		if *resource.Field != *other.Field {
			return false
		}
	}
	if resource.Id != other.Id {
		return false
	}
	if resource.Type != other.Type {
		return false
	}
	// is DeepEqual good enough here?
	if !reflect.DeepEqual(resource.Settings, other.Settings) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `Nested` fields for violations and returns them.
func (resource Nested) Validate() error {
	var errs cog.BuildErrors
	if !(resource.Type == "nested") {
		errs = append(errs, cog.MakeBuildErrors(
			"type",
			errors.New("must be == nested"),
		)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type Terms struct {
	Field    *string                     `json:"field,omitempty"`
	Id       string                      `json:"id"`
	Type     string                      `json:"type"`
	Settings *ElasticsearchTermsSettings `json:"settings,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Terms` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *Terms) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "field"
	if fields["field"] != nil {
		if string(fields["field"]) != "null" {
			if err := json.Unmarshal(fields["field"], &resource.Field); err != nil {
				errs = append(errs, cog.MakeBuildErrors("field", err)...)
			}

		}
		delete(fields, "field")

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
	// Field "settings"
	if fields["settings"] != nil {
		if string(fields["settings"]) != "null" {

			resource.Settings = &ElasticsearchTermsSettings{}
			if err := resource.Settings.UnmarshalJSONStrict(fields["settings"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("settings", err)...)
			}

		}
		delete(fields, "settings")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("Terms", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `Terms` objects.
func (resource Terms) Equals(other Terms) bool {
	if resource.Field == nil && other.Field != nil || resource.Field != nil && other.Field == nil {
		return false
	}

	if resource.Field != nil {
		if *resource.Field != *other.Field {
			return false
		}
	}
	if resource.Id != other.Id {
		return false
	}
	if resource.Type != other.Type {
		return false
	}
	if resource.Settings == nil && other.Settings != nil || resource.Settings != nil && other.Settings == nil {
		return false
	}

	if resource.Settings != nil {
		if !resource.Settings.Equals(*other.Settings) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `Terms` fields for violations and returns them.
func (resource Terms) Validate() error {
	var errs cog.BuildErrors
	if !(resource.Type == "terms") {
		errs = append(errs, cog.MakeBuildErrors(
			"type",
			errors.New("must be == terms"),
		)...)
	}
	if resource.Settings != nil {
		if err := resource.Settings.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("settings", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type TermsSettings struct {
	Order       *TermsOrder `json:"order,omitempty"`
	Size        *string     `json:"size,omitempty"`
	MinDocCount *string     `json:"min_doc_count,omitempty"`
	OrderBy     *string     `json:"orderBy,omitempty"`
	Missing     *string     `json:"missing,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TermsSettings` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *TermsSettings) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "order"
	if fields["order"] != nil {
		if string(fields["order"]) != "null" {
			if err := json.Unmarshal(fields["order"], &resource.Order); err != nil {
				errs = append(errs, cog.MakeBuildErrors("order", err)...)
			}

		}
		delete(fields, "order")

	}
	// Field "size"
	if fields["size"] != nil {
		if string(fields["size"]) != "null" {
			if err := json.Unmarshal(fields["size"], &resource.Size); err != nil {
				errs = append(errs, cog.MakeBuildErrors("size", err)...)
			}

		}
		delete(fields, "size")

	}
	// Field "min_doc_count"
	if fields["min_doc_count"] != nil {
		if string(fields["min_doc_count"]) != "null" {
			if err := json.Unmarshal(fields["min_doc_count"], &resource.MinDocCount); err != nil {
				errs = append(errs, cog.MakeBuildErrors("min_doc_count", err)...)
			}

		}
		delete(fields, "min_doc_count")

	}
	// Field "orderBy"
	if fields["orderBy"] != nil {
		if string(fields["orderBy"]) != "null" {
			if err := json.Unmarshal(fields["orderBy"], &resource.OrderBy); err != nil {
				errs = append(errs, cog.MakeBuildErrors("orderBy", err)...)
			}

		}
		delete(fields, "orderBy")

	}
	// Field "missing"
	if fields["missing"] != nil {
		if string(fields["missing"]) != "null" {
			if err := json.Unmarshal(fields["missing"], &resource.Missing); err != nil {
				errs = append(errs, cog.MakeBuildErrors("missing", err)...)
			}

		}
		delete(fields, "missing")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("TermsSettings", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `TermsSettings` objects.
func (resource TermsSettings) Equals(other TermsSettings) bool {
	if resource.Order == nil && other.Order != nil || resource.Order != nil && other.Order == nil {
		return false
	}

	if resource.Order != nil {
		if *resource.Order != *other.Order {
			return false
		}
	}
	if resource.Size == nil && other.Size != nil || resource.Size != nil && other.Size == nil {
		return false
	}

	if resource.Size != nil {
		if *resource.Size != *other.Size {
			return false
		}
	}
	if resource.MinDocCount == nil && other.MinDocCount != nil || resource.MinDocCount != nil && other.MinDocCount == nil {
		return false
	}

	if resource.MinDocCount != nil {
		if *resource.MinDocCount != *other.MinDocCount {
			return false
		}
	}
	if resource.OrderBy == nil && other.OrderBy != nil || resource.OrderBy != nil && other.OrderBy == nil {
		return false
	}

	if resource.OrderBy != nil {
		if *resource.OrderBy != *other.OrderBy {
			return false
		}
	}
	if resource.Missing == nil && other.Missing != nil || resource.Missing != nil && other.Missing == nil {
		return false
	}

	if resource.Missing != nil {
		if *resource.Missing != *other.Missing {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `TermsSettings` fields for violations and returns them.
func (resource TermsSettings) Validate() error {
	return nil
}

type Filters struct {
	Id       string                        `json:"id"`
	Type     string                        `json:"type"`
	Settings *ElasticsearchFiltersSettings `json:"settings,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Filters` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *Filters) UnmarshalJSONStrict(raw []byte) error {
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
	// Field "settings"
	if fields["settings"] != nil {
		if string(fields["settings"]) != "null" {

			resource.Settings = &ElasticsearchFiltersSettings{}
			if err := resource.Settings.UnmarshalJSONStrict(fields["settings"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("settings", err)...)
			}

		}
		delete(fields, "settings")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("Filters", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `Filters` objects.
func (resource Filters) Equals(other Filters) bool {
	if resource.Id != other.Id {
		return false
	}
	if resource.Type != other.Type {
		return false
	}
	if resource.Settings == nil && other.Settings != nil || resource.Settings != nil && other.Settings == nil {
		return false
	}

	if resource.Settings != nil {
		if !resource.Settings.Equals(*other.Settings) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `Filters` fields for violations and returns them.
func (resource Filters) Validate() error {
	var errs cog.BuildErrors
	if !(resource.Type == "filters") {
		errs = append(errs, cog.MakeBuildErrors(
			"type",
			errors.New("must be == filters"),
		)...)
	}
	if resource.Settings != nil {
		if err := resource.Settings.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("settings", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type Filter struct {
	Query string `json:"query"`
	Label string `json:"label"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Filter` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *Filter) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "query"
	if fields["query"] != nil {
		if string(fields["query"]) != "null" {
			if err := json.Unmarshal(fields["query"], &resource.Query); err != nil {
				errs = append(errs, cog.MakeBuildErrors("query", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("query", errors.New("required field is null"))...)

		}
		delete(fields, "query")
	} else {
		errs = append(errs, cog.MakeBuildErrors("query", errors.New("required field is missing from input"))...)
	}
	// Field "label"
	if fields["label"] != nil {
		if string(fields["label"]) != "null" {
			if err := json.Unmarshal(fields["label"], &resource.Label); err != nil {
				errs = append(errs, cog.MakeBuildErrors("label", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("label", errors.New("required field is null"))...)

		}
		delete(fields, "label")
	} else {
		errs = append(errs, cog.MakeBuildErrors("label", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("Filter", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `Filter` objects.
func (resource Filter) Equals(other Filter) bool {
	if resource.Query != other.Query {
		return false
	}
	if resource.Label != other.Label {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `Filter` fields for violations and returns them.
func (resource Filter) Validate() error {
	return nil
}

type FiltersSettings struct {
	Filters []Filter `json:"filters,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `FiltersSettings` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *FiltersSettings) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "filters"
	if fields["filters"] != nil {
		if string(fields["filters"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["filters"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 Filter

				result1 = Filter{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("filters["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.Filters = append(resource.Filters, result1)
			}

		}
		delete(fields, "filters")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("FiltersSettings", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `FiltersSettings` objects.
func (resource FiltersSettings) Equals(other FiltersSettings) bool {

	if len(resource.Filters) != len(other.Filters) {
		return false
	}

	for i1 := range resource.Filters {
		if !resource.Filters[i1].Equals(other.Filters[i1]) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `FiltersSettings` fields for violations and returns them.
func (resource FiltersSettings) Validate() error {
	var errs cog.BuildErrors

	for i1 := range resource.Filters {
		if err := resource.Filters[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("filters["+strconv.Itoa(i1)+"]", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type GeoHashGrid struct {
	Field    *string                           `json:"field,omitempty"`
	Id       string                            `json:"id"`
	Type     string                            `json:"type"`
	Settings *ElasticsearchGeoHashGridSettings `json:"settings,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `GeoHashGrid` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *GeoHashGrid) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "field"
	if fields["field"] != nil {
		if string(fields["field"]) != "null" {
			if err := json.Unmarshal(fields["field"], &resource.Field); err != nil {
				errs = append(errs, cog.MakeBuildErrors("field", err)...)
			}

		}
		delete(fields, "field")

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
	// Field "settings"
	if fields["settings"] != nil {
		if string(fields["settings"]) != "null" {

			resource.Settings = &ElasticsearchGeoHashGridSettings{}
			if err := resource.Settings.UnmarshalJSONStrict(fields["settings"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("settings", err)...)
			}

		}
		delete(fields, "settings")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("GeoHashGrid", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `GeoHashGrid` objects.
func (resource GeoHashGrid) Equals(other GeoHashGrid) bool {
	if resource.Field == nil && other.Field != nil || resource.Field != nil && other.Field == nil {
		return false
	}

	if resource.Field != nil {
		if *resource.Field != *other.Field {
			return false
		}
	}
	if resource.Id != other.Id {
		return false
	}
	if resource.Type != other.Type {
		return false
	}
	if resource.Settings == nil && other.Settings != nil || resource.Settings != nil && other.Settings == nil {
		return false
	}

	if resource.Settings != nil {
		if !resource.Settings.Equals(*other.Settings) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `GeoHashGrid` fields for violations and returns them.
func (resource GeoHashGrid) Validate() error {
	var errs cog.BuildErrors
	if !(resource.Type == "geohash_grid") {
		errs = append(errs, cog.MakeBuildErrors(
			"type",
			errors.New("must be == geohash_grid"),
		)...)
	}
	if resource.Settings != nil {
		if err := resource.Settings.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("settings", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type GeoHashGridSettings struct {
	Precision *string `json:"precision,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `GeoHashGridSettings` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *GeoHashGridSettings) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "precision"
	if fields["precision"] != nil {
		if string(fields["precision"]) != "null" {
			if err := json.Unmarshal(fields["precision"], &resource.Precision); err != nil {
				errs = append(errs, cog.MakeBuildErrors("precision", err)...)
			}

		}
		delete(fields, "precision")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("GeoHashGridSettings", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `GeoHashGridSettings` objects.
func (resource GeoHashGridSettings) Equals(other GeoHashGridSettings) bool {
	if resource.Precision == nil && other.Precision != nil || resource.Precision != nil && other.Precision == nil {
		return false
	}

	if resource.Precision != nil {
		if *resource.Precision != *other.Precision {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `GeoHashGridSettings` fields for violations and returns them.
func (resource GeoHashGridSettings) Validate() error {
	return nil
}

type PipelineMetricAggregationType string

const (
	PipelineMetricAggregationTypeMovingAvg     PipelineMetricAggregationType = "moving_avg"
	PipelineMetricAggregationTypeMovingFn      PipelineMetricAggregationType = "moving_fn"
	PipelineMetricAggregationTypeDerivative    PipelineMetricAggregationType = "derivative"
	PipelineMetricAggregationTypeSerialDiff    PipelineMetricAggregationType = "serial_diff"
	PipelineMetricAggregationTypeCumulativeSum PipelineMetricAggregationType = "cumulative_sum"
	PipelineMetricAggregationTypeBucketScript  PipelineMetricAggregationType = "bucket_script"
)

type MetricAggregationType = StringOrPipelineMetricAggregationType

type BaseMetricAggregation struct {
	Type MetricAggregationType `json:"type"`
	Id   string                `json:"id"`
	Hide *bool                 `json:"hide,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `BaseMetricAggregation` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *BaseMetricAggregation) UnmarshalJSONStrict(raw []byte) error {
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

			resource.Type = MetricAggregationType{}
			if err := resource.Type.UnmarshalJSONStrict(fields["type"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("type", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is null"))...)

		}
		delete(fields, "type")
	} else {
		errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is missing from input"))...)
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
	// Field "hide"
	if fields["hide"] != nil {
		if string(fields["hide"]) != "null" {
			if err := json.Unmarshal(fields["hide"], &resource.Hide); err != nil {
				errs = append(errs, cog.MakeBuildErrors("hide", err)...)
			}

		}
		delete(fields, "hide")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("BaseMetricAggregation", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `BaseMetricAggregation` objects.
func (resource BaseMetricAggregation) Equals(other BaseMetricAggregation) bool {
	if !resource.Type.Equals(other.Type) {
		return false
	}
	if resource.Id != other.Id {
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

	return true
}

// Validate checks all the validation constraints that may be defined on `BaseMetricAggregation` fields for violations and returns them.
func (resource BaseMetricAggregation) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Type.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("type", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type PipelineVariable struct {
	Name        string `json:"name"`
	PipelineAgg string `json:"pipelineAgg"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `PipelineVariable` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *PipelineVariable) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
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
	// Field "pipelineAgg"
	if fields["pipelineAgg"] != nil {
		if string(fields["pipelineAgg"]) != "null" {
			if err := json.Unmarshal(fields["pipelineAgg"], &resource.PipelineAgg); err != nil {
				errs = append(errs, cog.MakeBuildErrors("pipelineAgg", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("pipelineAgg", errors.New("required field is null"))...)

		}
		delete(fields, "pipelineAgg")
	} else {
		errs = append(errs, cog.MakeBuildErrors("pipelineAgg", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("PipelineVariable", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `PipelineVariable` objects.
func (resource PipelineVariable) Equals(other PipelineVariable) bool {
	if resource.Name != other.Name {
		return false
	}
	if resource.PipelineAgg != other.PipelineAgg {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `PipelineVariable` fields for violations and returns them.
func (resource PipelineVariable) Validate() error {
	return nil
}

type MetricAggregationWithField struct {
	Field *string               `json:"field,omitempty"`
	Type  MetricAggregationType `json:"type"`
	Id    string                `json:"id"`
	Hide  *bool                 `json:"hide,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `MetricAggregationWithField` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *MetricAggregationWithField) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "field"
	if fields["field"] != nil {
		if string(fields["field"]) != "null" {
			if err := json.Unmarshal(fields["field"], &resource.Field); err != nil {
				errs = append(errs, cog.MakeBuildErrors("field", err)...)
			}

		}
		delete(fields, "field")

	}
	// Field "type"
	if fields["type"] != nil {
		if string(fields["type"]) != "null" {

			resource.Type = MetricAggregationType{}
			if err := resource.Type.UnmarshalJSONStrict(fields["type"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("type", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is null"))...)

		}
		delete(fields, "type")
	} else {
		errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is missing from input"))...)
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
	// Field "hide"
	if fields["hide"] != nil {
		if string(fields["hide"]) != "null" {
			if err := json.Unmarshal(fields["hide"], &resource.Hide); err != nil {
				errs = append(errs, cog.MakeBuildErrors("hide", err)...)
			}

		}
		delete(fields, "hide")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("MetricAggregationWithField", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `MetricAggregationWithField` objects.
func (resource MetricAggregationWithField) Equals(other MetricAggregationWithField) bool {
	if resource.Field == nil && other.Field != nil || resource.Field != nil && other.Field == nil {
		return false
	}

	if resource.Field != nil {
		if *resource.Field != *other.Field {
			return false
		}
	}
	if !resource.Type.Equals(other.Type) {
		return false
	}
	if resource.Id != other.Id {
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

	return true
}

// Validate checks all the validation constraints that may be defined on `MetricAggregationWithField` fields for violations and returns them.
func (resource MetricAggregationWithField) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Type.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("type", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type MetricAggregationWithMissingSupport struct {
	Settings *ElasticsearchMetricAggregationWithMissingSupportSettings `json:"settings,omitempty"`
	Type     MetricAggregationType                                     `json:"type"`
	Id       string                                                    `json:"id"`
	Hide     *bool                                                     `json:"hide,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `MetricAggregationWithMissingSupport` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *MetricAggregationWithMissingSupport) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "settings"
	if fields["settings"] != nil {
		if string(fields["settings"]) != "null" {

			resource.Settings = &ElasticsearchMetricAggregationWithMissingSupportSettings{}
			if err := resource.Settings.UnmarshalJSONStrict(fields["settings"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("settings", err)...)
			}

		}
		delete(fields, "settings")

	}
	// Field "type"
	if fields["type"] != nil {
		if string(fields["type"]) != "null" {

			resource.Type = MetricAggregationType{}
			if err := resource.Type.UnmarshalJSONStrict(fields["type"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("type", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is null"))...)

		}
		delete(fields, "type")
	} else {
		errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is missing from input"))...)
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
	// Field "hide"
	if fields["hide"] != nil {
		if string(fields["hide"]) != "null" {
			if err := json.Unmarshal(fields["hide"], &resource.Hide); err != nil {
				errs = append(errs, cog.MakeBuildErrors("hide", err)...)
			}

		}
		delete(fields, "hide")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("MetricAggregationWithMissingSupport", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `MetricAggregationWithMissingSupport` objects.
func (resource MetricAggregationWithMissingSupport) Equals(other MetricAggregationWithMissingSupport) bool {
	if resource.Settings == nil && other.Settings != nil || resource.Settings != nil && other.Settings == nil {
		return false
	}

	if resource.Settings != nil {
		if !resource.Settings.Equals(*other.Settings) {
			return false
		}
	}
	if !resource.Type.Equals(other.Type) {
		return false
	}
	if resource.Id != other.Id {
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

	return true
}

// Validate checks all the validation constraints that may be defined on `MetricAggregationWithMissingSupport` fields for violations and returns them.
func (resource MetricAggregationWithMissingSupport) Validate() error {
	var errs cog.BuildErrors
	if resource.Settings != nil {
		if err := resource.Settings.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("settings", err)...)
		}
	}
	if err := resource.Type.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("type", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type InlineScript = StringOrElasticsearchInlineScript

type MetricAggregationWithInlineScript struct {
	Settings *ElasticsearchMetricAggregationWithInlineScriptSettings `json:"settings,omitempty"`
	Type     MetricAggregationType                                   `json:"type"`
	Id       string                                                  `json:"id"`
	Hide     *bool                                                   `json:"hide,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `MetricAggregationWithInlineScript` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *MetricAggregationWithInlineScript) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "settings"
	if fields["settings"] != nil {
		if string(fields["settings"]) != "null" {

			resource.Settings = &ElasticsearchMetricAggregationWithInlineScriptSettings{}
			if err := resource.Settings.UnmarshalJSONStrict(fields["settings"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("settings", err)...)
			}

		}
		delete(fields, "settings")

	}
	// Field "type"
	if fields["type"] != nil {
		if string(fields["type"]) != "null" {

			resource.Type = MetricAggregationType{}
			if err := resource.Type.UnmarshalJSONStrict(fields["type"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("type", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is null"))...)

		}
		delete(fields, "type")
	} else {
		errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is missing from input"))...)
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
	// Field "hide"
	if fields["hide"] != nil {
		if string(fields["hide"]) != "null" {
			if err := json.Unmarshal(fields["hide"], &resource.Hide); err != nil {
				errs = append(errs, cog.MakeBuildErrors("hide", err)...)
			}

		}
		delete(fields, "hide")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("MetricAggregationWithInlineScript", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `MetricAggregationWithInlineScript` objects.
func (resource MetricAggregationWithInlineScript) Equals(other MetricAggregationWithInlineScript) bool {
	if resource.Settings == nil && other.Settings != nil || resource.Settings != nil && other.Settings == nil {
		return false
	}

	if resource.Settings != nil {
		if !resource.Settings.Equals(*other.Settings) {
			return false
		}
	}
	if !resource.Type.Equals(other.Type) {
		return false
	}
	if resource.Id != other.Id {
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

	return true
}

// Validate checks all the validation constraints that may be defined on `MetricAggregationWithInlineScript` fields for violations and returns them.
func (resource MetricAggregationWithInlineScript) Validate() error {
	var errs cog.BuildErrors
	if resource.Settings != nil {
		if err := resource.Settings.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("settings", err)...)
		}
	}
	if err := resource.Type.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("type", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type Count struct {
	Type string `json:"type"`
	Id   string `json:"id"`
	Hide *bool  `json:"hide,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Count` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *Count) UnmarshalJSONStrict(raw []byte) error {
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
	// Field "hide"
	if fields["hide"] != nil {
		if string(fields["hide"]) != "null" {
			if err := json.Unmarshal(fields["hide"], &resource.Hide); err != nil {
				errs = append(errs, cog.MakeBuildErrors("hide", err)...)
			}

		}
		delete(fields, "hide")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("Count", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `Count` objects.
func (resource Count) Equals(other Count) bool {
	if resource.Type != other.Type {
		return false
	}
	if resource.Id != other.Id {
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

	return true
}

// Validate checks all the validation constraints that may be defined on `Count` fields for violations and returns them.
func (resource Count) Validate() error {
	var errs cog.BuildErrors
	if !(resource.Type == "count") {
		errs = append(errs, cog.MakeBuildErrors(
			"type",
			errors.New("must be == count"),
		)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type Average struct {
	Type     string                        `json:"type"`
	Field    *string                       `json:"field,omitempty"`
	Id       string                        `json:"id"`
	Settings *ElasticsearchAverageSettings `json:"settings,omitempty"`
	Hide     *bool                         `json:"hide,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Average` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *Average) UnmarshalJSONStrict(raw []byte) error {
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
	// Field "field"
	if fields["field"] != nil {
		if string(fields["field"]) != "null" {
			if err := json.Unmarshal(fields["field"], &resource.Field); err != nil {
				errs = append(errs, cog.MakeBuildErrors("field", err)...)
			}

		}
		delete(fields, "field")

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
	// Field "settings"
	if fields["settings"] != nil {
		if string(fields["settings"]) != "null" {

			resource.Settings = &ElasticsearchAverageSettings{}
			if err := resource.Settings.UnmarshalJSONStrict(fields["settings"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("settings", err)...)
			}

		}
		delete(fields, "settings")

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

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("Average", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `Average` objects.
func (resource Average) Equals(other Average) bool {
	if resource.Type != other.Type {
		return false
	}
	if resource.Field == nil && other.Field != nil || resource.Field != nil && other.Field == nil {
		return false
	}

	if resource.Field != nil {
		if *resource.Field != *other.Field {
			return false
		}
	}
	if resource.Id != other.Id {
		return false
	}
	if resource.Settings == nil && other.Settings != nil || resource.Settings != nil && other.Settings == nil {
		return false
	}

	if resource.Settings != nil {
		if !resource.Settings.Equals(*other.Settings) {
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

	return true
}

// Validate checks all the validation constraints that may be defined on `Average` fields for violations and returns them.
func (resource Average) Validate() error {
	var errs cog.BuildErrors
	if !(resource.Type == "avg") {
		errs = append(errs, cog.MakeBuildErrors(
			"type",
			errors.New("must be == avg"),
		)...)
	}
	if resource.Settings != nil {
		if err := resource.Settings.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("settings", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type Sum struct {
	Type     string                    `json:"type"`
	Field    *string                   `json:"field,omitempty"`
	Id       string                    `json:"id"`
	Settings *ElasticsearchSumSettings `json:"settings,omitempty"`
	Hide     *bool                     `json:"hide,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Sum` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *Sum) UnmarshalJSONStrict(raw []byte) error {
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
	// Field "field"
	if fields["field"] != nil {
		if string(fields["field"]) != "null" {
			if err := json.Unmarshal(fields["field"], &resource.Field); err != nil {
				errs = append(errs, cog.MakeBuildErrors("field", err)...)
			}

		}
		delete(fields, "field")

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
	// Field "settings"
	if fields["settings"] != nil {
		if string(fields["settings"]) != "null" {

			resource.Settings = &ElasticsearchSumSettings{}
			if err := resource.Settings.UnmarshalJSONStrict(fields["settings"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("settings", err)...)
			}

		}
		delete(fields, "settings")

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

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("Sum", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `Sum` objects.
func (resource Sum) Equals(other Sum) bool {
	if resource.Type != other.Type {
		return false
	}
	if resource.Field == nil && other.Field != nil || resource.Field != nil && other.Field == nil {
		return false
	}

	if resource.Field != nil {
		if *resource.Field != *other.Field {
			return false
		}
	}
	if resource.Id != other.Id {
		return false
	}
	if resource.Settings == nil && other.Settings != nil || resource.Settings != nil && other.Settings == nil {
		return false
	}

	if resource.Settings != nil {
		if !resource.Settings.Equals(*other.Settings) {
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

	return true
}

// Validate checks all the validation constraints that may be defined on `Sum` fields for violations and returns them.
func (resource Sum) Validate() error {
	var errs cog.BuildErrors
	if !(resource.Type == "sum") {
		errs = append(errs, cog.MakeBuildErrors(
			"type",
			errors.New("must be == sum"),
		)...)
	}
	if resource.Settings != nil {
		if err := resource.Settings.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("settings", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type Max struct {
	Type     string                    `json:"type"`
	Field    *string                   `json:"field,omitempty"`
	Id       string                    `json:"id"`
	Settings *ElasticsearchMaxSettings `json:"settings,omitempty"`
	Hide     *bool                     `json:"hide,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Max` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *Max) UnmarshalJSONStrict(raw []byte) error {
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
	// Field "field"
	if fields["field"] != nil {
		if string(fields["field"]) != "null" {
			if err := json.Unmarshal(fields["field"], &resource.Field); err != nil {
				errs = append(errs, cog.MakeBuildErrors("field", err)...)
			}

		}
		delete(fields, "field")

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
	// Field "settings"
	if fields["settings"] != nil {
		if string(fields["settings"]) != "null" {

			resource.Settings = &ElasticsearchMaxSettings{}
			if err := resource.Settings.UnmarshalJSONStrict(fields["settings"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("settings", err)...)
			}

		}
		delete(fields, "settings")

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

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("Max", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `Max` objects.
func (resource Max) Equals(other Max) bool {
	if resource.Type != other.Type {
		return false
	}
	if resource.Field == nil && other.Field != nil || resource.Field != nil && other.Field == nil {
		return false
	}

	if resource.Field != nil {
		if *resource.Field != *other.Field {
			return false
		}
	}
	if resource.Id != other.Id {
		return false
	}
	if resource.Settings == nil && other.Settings != nil || resource.Settings != nil && other.Settings == nil {
		return false
	}

	if resource.Settings != nil {
		if !resource.Settings.Equals(*other.Settings) {
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

	return true
}

// Validate checks all the validation constraints that may be defined on `Max` fields for violations and returns them.
func (resource Max) Validate() error {
	var errs cog.BuildErrors
	if !(resource.Type == "max") {
		errs = append(errs, cog.MakeBuildErrors(
			"type",
			errors.New("must be == max"),
		)...)
	}
	if resource.Settings != nil {
		if err := resource.Settings.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("settings", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type Min struct {
	Type     string                    `json:"type"`
	Field    *string                   `json:"field,omitempty"`
	Id       string                    `json:"id"`
	Settings *ElasticsearchMinSettings `json:"settings,omitempty"`
	Hide     *bool                     `json:"hide,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Min` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *Min) UnmarshalJSONStrict(raw []byte) error {
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
	// Field "field"
	if fields["field"] != nil {
		if string(fields["field"]) != "null" {
			if err := json.Unmarshal(fields["field"], &resource.Field); err != nil {
				errs = append(errs, cog.MakeBuildErrors("field", err)...)
			}

		}
		delete(fields, "field")

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
	// Field "settings"
	if fields["settings"] != nil {
		if string(fields["settings"]) != "null" {

			resource.Settings = &ElasticsearchMinSettings{}
			if err := resource.Settings.UnmarshalJSONStrict(fields["settings"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("settings", err)...)
			}

		}
		delete(fields, "settings")

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

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("Min", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `Min` objects.
func (resource Min) Equals(other Min) bool {
	if resource.Type != other.Type {
		return false
	}
	if resource.Field == nil && other.Field != nil || resource.Field != nil && other.Field == nil {
		return false
	}

	if resource.Field != nil {
		if *resource.Field != *other.Field {
			return false
		}
	}
	if resource.Id != other.Id {
		return false
	}
	if resource.Settings == nil && other.Settings != nil || resource.Settings != nil && other.Settings == nil {
		return false
	}

	if resource.Settings != nil {
		if !resource.Settings.Equals(*other.Settings) {
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

	return true
}

// Validate checks all the validation constraints that may be defined on `Min` fields for violations and returns them.
func (resource Min) Validate() error {
	var errs cog.BuildErrors
	if !(resource.Type == "min") {
		errs = append(errs, cog.MakeBuildErrors(
			"type",
			errors.New("must be == min"),
		)...)
	}
	if resource.Settings != nil {
		if err := resource.Settings.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("settings", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type ExtendedStatMetaType string

const (
	ExtendedStatMetaTypeAvg                     ExtendedStatMetaType = "avg"
	ExtendedStatMetaTypeMin                     ExtendedStatMetaType = "min"
	ExtendedStatMetaTypeMax                     ExtendedStatMetaType = "max"
	ExtendedStatMetaTypeSum                     ExtendedStatMetaType = "sum"
	ExtendedStatMetaTypeCount                   ExtendedStatMetaType = "count"
	ExtendedStatMetaTypeStdDeviation            ExtendedStatMetaType = "std_deviation"
	ExtendedStatMetaTypeStdDeviationBoundsUpper ExtendedStatMetaType = "std_deviation_bounds_upper"
	ExtendedStatMetaTypeStdDeviationBoundsLower ExtendedStatMetaType = "std_deviation_bounds_lower"
)

type ExtendedStat struct {
	Label string               `json:"label"`
	Value ExtendedStatMetaType `json:"value"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ExtendedStat` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ExtendedStat) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "label"
	if fields["label"] != nil {
		if string(fields["label"]) != "null" {
			if err := json.Unmarshal(fields["label"], &resource.Label); err != nil {
				errs = append(errs, cog.MakeBuildErrors("label", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("label", errors.New("required field is null"))...)

		}
		delete(fields, "label")
	} else {
		errs = append(errs, cog.MakeBuildErrors("label", errors.New("required field is missing from input"))...)
	}
	// Field "value"
	if fields["value"] != nil {
		if string(fields["value"]) != "null" {
			if err := json.Unmarshal(fields["value"], &resource.Value); err != nil {
				errs = append(errs, cog.MakeBuildErrors("value", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("value", errors.New("required field is null"))...)

		}
		delete(fields, "value")
	} else {
		errs = append(errs, cog.MakeBuildErrors("value", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ExtendedStat", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ExtendedStat` objects.
func (resource ExtendedStat) Equals(other ExtendedStat) bool {
	if resource.Label != other.Label {
		return false
	}
	if resource.Value != other.Value {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `ExtendedStat` fields for violations and returns them.
func (resource ExtendedStat) Validate() error {
	return nil
}

type ExtendedStats struct {
	Type     string                              `json:"type"`
	Settings *ElasticsearchExtendedStatsSettings `json:"settings,omitempty"`
	Field    *string                             `json:"field,omitempty"`
	Id       string                              `json:"id"`
	Meta     any                                 `json:"meta,omitempty"`
	Hide     *bool                               `json:"hide,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ExtendedStats` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ExtendedStats) UnmarshalJSONStrict(raw []byte) error {
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
	// Field "settings"
	if fields["settings"] != nil {
		if string(fields["settings"]) != "null" {

			resource.Settings = &ElasticsearchExtendedStatsSettings{}
			if err := resource.Settings.UnmarshalJSONStrict(fields["settings"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("settings", err)...)
			}

		}
		delete(fields, "settings")

	}
	// Field "field"
	if fields["field"] != nil {
		if string(fields["field"]) != "null" {
			if err := json.Unmarshal(fields["field"], &resource.Field); err != nil {
				errs = append(errs, cog.MakeBuildErrors("field", err)...)
			}

		}
		delete(fields, "field")

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
	// Field "meta"
	if fields["meta"] != nil {
		if string(fields["meta"]) != "null" {
			if err := json.Unmarshal(fields["meta"], &resource.Meta); err != nil {
				errs = append(errs, cog.MakeBuildErrors("meta", err)...)
			}

		}
		delete(fields, "meta")

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

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ExtendedStats", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ExtendedStats` objects.
func (resource ExtendedStats) Equals(other ExtendedStats) bool {
	if resource.Type != other.Type {
		return false
	}
	if resource.Settings == nil && other.Settings != nil || resource.Settings != nil && other.Settings == nil {
		return false
	}

	if resource.Settings != nil {
		if !resource.Settings.Equals(*other.Settings) {
			return false
		}
	}
	if resource.Field == nil && other.Field != nil || resource.Field != nil && other.Field == nil {
		return false
	}

	if resource.Field != nil {
		if *resource.Field != *other.Field {
			return false
		}
	}
	if resource.Id != other.Id {
		return false
	}
	// is DeepEqual good enough here?
	if !reflect.DeepEqual(resource.Meta, other.Meta) {
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

	return true
}

// Validate checks all the validation constraints that may be defined on `ExtendedStats` fields for violations and returns them.
func (resource ExtendedStats) Validate() error {
	var errs cog.BuildErrors
	if !(resource.Type == "extended_stats") {
		errs = append(errs, cog.MakeBuildErrors(
			"type",
			errors.New("must be == extended_stats"),
		)...)
	}
	if resource.Settings != nil {
		if err := resource.Settings.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("settings", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type Percentiles struct {
	Type     string                            `json:"type"`
	Field    *string                           `json:"field,omitempty"`
	Id       string                            `json:"id"`
	Settings *ElasticsearchPercentilesSettings `json:"settings,omitempty"`
	Hide     *bool                             `json:"hide,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Percentiles` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *Percentiles) UnmarshalJSONStrict(raw []byte) error {
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
	// Field "field"
	if fields["field"] != nil {
		if string(fields["field"]) != "null" {
			if err := json.Unmarshal(fields["field"], &resource.Field); err != nil {
				errs = append(errs, cog.MakeBuildErrors("field", err)...)
			}

		}
		delete(fields, "field")

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
	// Field "settings"
	if fields["settings"] != nil {
		if string(fields["settings"]) != "null" {

			resource.Settings = &ElasticsearchPercentilesSettings{}
			if err := resource.Settings.UnmarshalJSONStrict(fields["settings"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("settings", err)...)
			}

		}
		delete(fields, "settings")

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

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("Percentiles", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `Percentiles` objects.
func (resource Percentiles) Equals(other Percentiles) bool {
	if resource.Type != other.Type {
		return false
	}
	if resource.Field == nil && other.Field != nil || resource.Field != nil && other.Field == nil {
		return false
	}

	if resource.Field != nil {
		if *resource.Field != *other.Field {
			return false
		}
	}
	if resource.Id != other.Id {
		return false
	}
	if resource.Settings == nil && other.Settings != nil || resource.Settings != nil && other.Settings == nil {
		return false
	}

	if resource.Settings != nil {
		if !resource.Settings.Equals(*other.Settings) {
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

	return true
}

// Validate checks all the validation constraints that may be defined on `Percentiles` fields for violations and returns them.
func (resource Percentiles) Validate() error {
	var errs cog.BuildErrors
	if !(resource.Type == "percentiles") {
		errs = append(errs, cog.MakeBuildErrors(
			"type",
			errors.New("must be == percentiles"),
		)...)
	}
	if resource.Settings != nil {
		if err := resource.Settings.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("settings", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type UniqueCount struct {
	Type     string                            `json:"type"`
	Field    *string                           `json:"field,omitempty"`
	Id       string                            `json:"id"`
	Settings *ElasticsearchUniqueCountSettings `json:"settings,omitempty"`
	Hide     *bool                             `json:"hide,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `UniqueCount` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *UniqueCount) UnmarshalJSONStrict(raw []byte) error {
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
	// Field "field"
	if fields["field"] != nil {
		if string(fields["field"]) != "null" {
			if err := json.Unmarshal(fields["field"], &resource.Field); err != nil {
				errs = append(errs, cog.MakeBuildErrors("field", err)...)
			}

		}
		delete(fields, "field")

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
	// Field "settings"
	if fields["settings"] != nil {
		if string(fields["settings"]) != "null" {

			resource.Settings = &ElasticsearchUniqueCountSettings{}
			if err := resource.Settings.UnmarshalJSONStrict(fields["settings"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("settings", err)...)
			}

		}
		delete(fields, "settings")

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

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("UniqueCount", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `UniqueCount` objects.
func (resource UniqueCount) Equals(other UniqueCount) bool {
	if resource.Type != other.Type {
		return false
	}
	if resource.Field == nil && other.Field != nil || resource.Field != nil && other.Field == nil {
		return false
	}

	if resource.Field != nil {
		if *resource.Field != *other.Field {
			return false
		}
	}
	if resource.Id != other.Id {
		return false
	}
	if resource.Settings == nil && other.Settings != nil || resource.Settings != nil && other.Settings == nil {
		return false
	}

	if resource.Settings != nil {
		if !resource.Settings.Equals(*other.Settings) {
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

	return true
}

// Validate checks all the validation constraints that may be defined on `UniqueCount` fields for violations and returns them.
func (resource UniqueCount) Validate() error {
	var errs cog.BuildErrors
	if !(resource.Type == "cardinality") {
		errs = append(errs, cog.MakeBuildErrors(
			"type",
			errors.New("must be == cardinality"),
		)...)
	}
	if resource.Settings != nil {
		if err := resource.Settings.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("settings", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type RawDocument struct {
	Type     string                            `json:"type"`
	Id       string                            `json:"id"`
	Settings *ElasticsearchRawDocumentSettings `json:"settings,omitempty"`
	Hide     *bool                             `json:"hide,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `RawDocument` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *RawDocument) UnmarshalJSONStrict(raw []byte) error {
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
	// Field "settings"
	if fields["settings"] != nil {
		if string(fields["settings"]) != "null" {

			resource.Settings = &ElasticsearchRawDocumentSettings{}
			if err := resource.Settings.UnmarshalJSONStrict(fields["settings"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("settings", err)...)
			}

		}
		delete(fields, "settings")

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

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("RawDocument", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `RawDocument` objects.
func (resource RawDocument) Equals(other RawDocument) bool {
	if resource.Type != other.Type {
		return false
	}
	if resource.Id != other.Id {
		return false
	}
	if resource.Settings == nil && other.Settings != nil || resource.Settings != nil && other.Settings == nil {
		return false
	}

	if resource.Settings != nil {
		if !resource.Settings.Equals(*other.Settings) {
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

	return true
}

// Validate checks all the validation constraints that may be defined on `RawDocument` fields for violations and returns them.
func (resource RawDocument) Validate() error {
	var errs cog.BuildErrors
	if !(resource.Type == "raw_document") {
		errs = append(errs, cog.MakeBuildErrors(
			"type",
			errors.New("must be == raw_document"),
		)...)
	}
	if resource.Settings != nil {
		if err := resource.Settings.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("settings", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type RawData struct {
	Type     string                        `json:"type"`
	Id       string                        `json:"id"`
	Settings *ElasticsearchRawDataSettings `json:"settings,omitempty"`
	Hide     *bool                         `json:"hide,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `RawData` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *RawData) UnmarshalJSONStrict(raw []byte) error {
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
	// Field "settings"
	if fields["settings"] != nil {
		if string(fields["settings"]) != "null" {

			resource.Settings = &ElasticsearchRawDataSettings{}
			if err := resource.Settings.UnmarshalJSONStrict(fields["settings"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("settings", err)...)
			}

		}
		delete(fields, "settings")

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

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("RawData", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `RawData` objects.
func (resource RawData) Equals(other RawData) bool {
	if resource.Type != other.Type {
		return false
	}
	if resource.Id != other.Id {
		return false
	}
	if resource.Settings == nil && other.Settings != nil || resource.Settings != nil && other.Settings == nil {
		return false
	}

	if resource.Settings != nil {
		if !resource.Settings.Equals(*other.Settings) {
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

	return true
}

// Validate checks all the validation constraints that may be defined on `RawData` fields for violations and returns them.
func (resource RawData) Validate() error {
	var errs cog.BuildErrors
	if !(resource.Type == "raw_data") {
		errs = append(errs, cog.MakeBuildErrors(
			"type",
			errors.New("must be == raw_data"),
		)...)
	}
	if resource.Settings != nil {
		if err := resource.Settings.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("settings", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type Logs struct {
	Type     string                     `json:"type"`
	Id       string                     `json:"id"`
	Settings *ElasticsearchLogsSettings `json:"settings,omitempty"`
	Hide     *bool                      `json:"hide,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Logs` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *Logs) UnmarshalJSONStrict(raw []byte) error {
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
	// Field "settings"
	if fields["settings"] != nil {
		if string(fields["settings"]) != "null" {

			resource.Settings = &ElasticsearchLogsSettings{}
			if err := resource.Settings.UnmarshalJSONStrict(fields["settings"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("settings", err)...)
			}

		}
		delete(fields, "settings")

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

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("Logs", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `Logs` objects.
func (resource Logs) Equals(other Logs) bool {
	if resource.Type != other.Type {
		return false
	}
	if resource.Id != other.Id {
		return false
	}
	if resource.Settings == nil && other.Settings != nil || resource.Settings != nil && other.Settings == nil {
		return false
	}

	if resource.Settings != nil {
		if !resource.Settings.Equals(*other.Settings) {
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

	return true
}

// Validate checks all the validation constraints that may be defined on `Logs` fields for violations and returns them.
func (resource Logs) Validate() error {
	var errs cog.BuildErrors
	if !(resource.Type == "logs") {
		errs = append(errs, cog.MakeBuildErrors(
			"type",
			errors.New("must be == logs"),
		)...)
	}
	if resource.Settings != nil {
		if err := resource.Settings.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("settings", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type Rate struct {
	Type     string                     `json:"type"`
	Field    *string                    `json:"field,omitempty"`
	Id       string                     `json:"id"`
	Settings *ElasticsearchRateSettings `json:"settings,omitempty"`
	Hide     *bool                      `json:"hide,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Rate` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *Rate) UnmarshalJSONStrict(raw []byte) error {
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
	// Field "field"
	if fields["field"] != nil {
		if string(fields["field"]) != "null" {
			if err := json.Unmarshal(fields["field"], &resource.Field); err != nil {
				errs = append(errs, cog.MakeBuildErrors("field", err)...)
			}

		}
		delete(fields, "field")

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
	// Field "settings"
	if fields["settings"] != nil {
		if string(fields["settings"]) != "null" {

			resource.Settings = &ElasticsearchRateSettings{}
			if err := resource.Settings.UnmarshalJSONStrict(fields["settings"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("settings", err)...)
			}

		}
		delete(fields, "settings")

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

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("Rate", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `Rate` objects.
func (resource Rate) Equals(other Rate) bool {
	if resource.Type != other.Type {
		return false
	}
	if resource.Field == nil && other.Field != nil || resource.Field != nil && other.Field == nil {
		return false
	}

	if resource.Field != nil {
		if *resource.Field != *other.Field {
			return false
		}
	}
	if resource.Id != other.Id {
		return false
	}
	if resource.Settings == nil && other.Settings != nil || resource.Settings != nil && other.Settings == nil {
		return false
	}

	if resource.Settings != nil {
		if !resource.Settings.Equals(*other.Settings) {
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

	return true
}

// Validate checks all the validation constraints that may be defined on `Rate` fields for violations and returns them.
func (resource Rate) Validate() error {
	var errs cog.BuildErrors
	if !(resource.Type == "rate") {
		errs = append(errs, cog.MakeBuildErrors(
			"type",
			errors.New("must be == rate"),
		)...)
	}
	if resource.Settings != nil {
		if err := resource.Settings.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("settings", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type BasePipelineMetricAggregation struct {
	PipelineAgg *string `json:"pipelineAgg,omitempty"`
	Field       *string `json:"field,omitempty"`
	Type        string  `json:"type"`
	Id          string  `json:"id"`
	Hide        *bool   `json:"hide,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `BasePipelineMetricAggregation` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *BasePipelineMetricAggregation) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "pipelineAgg"
	if fields["pipelineAgg"] != nil {
		if string(fields["pipelineAgg"]) != "null" {
			if err := json.Unmarshal(fields["pipelineAgg"], &resource.PipelineAgg); err != nil {
				errs = append(errs, cog.MakeBuildErrors("pipelineAgg", err)...)
			}

		}
		delete(fields, "pipelineAgg")

	}
	// Field "field"
	if fields["field"] != nil {
		if string(fields["field"]) != "null" {
			if err := json.Unmarshal(fields["field"], &resource.Field); err != nil {
				errs = append(errs, cog.MakeBuildErrors("field", err)...)
			}

		}
		delete(fields, "field")

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
	// Field "hide"
	if fields["hide"] != nil {
		if string(fields["hide"]) != "null" {
			if err := json.Unmarshal(fields["hide"], &resource.Hide); err != nil {
				errs = append(errs, cog.MakeBuildErrors("hide", err)...)
			}

		}
		delete(fields, "hide")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("BasePipelineMetricAggregation", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `BasePipelineMetricAggregation` objects.
func (resource BasePipelineMetricAggregation) Equals(other BasePipelineMetricAggregation) bool {
	if resource.PipelineAgg == nil && other.PipelineAgg != nil || resource.PipelineAgg != nil && other.PipelineAgg == nil {
		return false
	}

	if resource.PipelineAgg != nil {
		if *resource.PipelineAgg != *other.PipelineAgg {
			return false
		}
	}
	if resource.Field == nil && other.Field != nil || resource.Field != nil && other.Field == nil {
		return false
	}

	if resource.Field != nil {
		if *resource.Field != *other.Field {
			return false
		}
	}
	if resource.Type != other.Type {
		return false
	}
	if resource.Id != other.Id {
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

	return true
}

// Validate checks all the validation constraints that may be defined on `BasePipelineMetricAggregation` fields for violations and returns them.
func (resource BasePipelineMetricAggregation) Validate() error {
	return nil
}

type PipelineMetricAggregationWithMultipleBucketPaths struct {
	PipelineVariables []PipelineVariable    `json:"pipelineVariables,omitempty"`
	Type              MetricAggregationType `json:"type"`
	Id                string                `json:"id"`
	Hide              *bool                 `json:"hide,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `PipelineMetricAggregationWithMultipleBucketPaths` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *PipelineMetricAggregationWithMultipleBucketPaths) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "pipelineVariables"
	if fields["pipelineVariables"] != nil {
		if string(fields["pipelineVariables"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["pipelineVariables"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 PipelineVariable

				result1 = PipelineVariable{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("pipelineVariables["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.PipelineVariables = append(resource.PipelineVariables, result1)
			}

		}
		delete(fields, "pipelineVariables")

	}
	// Field "type"
	if fields["type"] != nil {
		if string(fields["type"]) != "null" {

			resource.Type = MetricAggregationType{}
			if err := resource.Type.UnmarshalJSONStrict(fields["type"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("type", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is null"))...)

		}
		delete(fields, "type")
	} else {
		errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is missing from input"))...)
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
	// Field "hide"
	if fields["hide"] != nil {
		if string(fields["hide"]) != "null" {
			if err := json.Unmarshal(fields["hide"], &resource.Hide); err != nil {
				errs = append(errs, cog.MakeBuildErrors("hide", err)...)
			}

		}
		delete(fields, "hide")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("PipelineMetricAggregationWithMultipleBucketPaths", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `PipelineMetricAggregationWithMultipleBucketPaths` objects.
func (resource PipelineMetricAggregationWithMultipleBucketPaths) Equals(other PipelineMetricAggregationWithMultipleBucketPaths) bool {

	if len(resource.PipelineVariables) != len(other.PipelineVariables) {
		return false
	}

	for i1 := range resource.PipelineVariables {
		if !resource.PipelineVariables[i1].Equals(other.PipelineVariables[i1]) {
			return false
		}
	}
	if !resource.Type.Equals(other.Type) {
		return false
	}
	if resource.Id != other.Id {
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

	return true
}

// Validate checks all the validation constraints that may be defined on `PipelineMetricAggregationWithMultipleBucketPaths` fields for violations and returns them.
func (resource PipelineMetricAggregationWithMultipleBucketPaths) Validate() error {
	var errs cog.BuildErrors

	for i1 := range resource.PipelineVariables {
		if err := resource.PipelineVariables[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("pipelineVariables["+strconv.Itoa(i1)+"]", err)...)
		}
	}
	if err := resource.Type.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("type", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type MovingAverageModel string

const (
	MovingAverageModelSimple      MovingAverageModel = "simple"
	MovingAverageModelLinear      MovingAverageModel = "linear"
	MovingAverageModelEwma        MovingAverageModel = "ewma"
	MovingAverageModelHolt        MovingAverageModel = "holt"
	MovingAverageModelHoltWinters MovingAverageModel = "holt_winters"
)

type MovingAverageModelOption struct {
	Label string             `json:"label"`
	Value MovingAverageModel `json:"value"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `MovingAverageModelOption` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *MovingAverageModelOption) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "label"
	if fields["label"] != nil {
		if string(fields["label"]) != "null" {
			if err := json.Unmarshal(fields["label"], &resource.Label); err != nil {
				errs = append(errs, cog.MakeBuildErrors("label", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("label", errors.New("required field is null"))...)

		}
		delete(fields, "label")
	} else {
		errs = append(errs, cog.MakeBuildErrors("label", errors.New("required field is missing from input"))...)
	}
	// Field "value"
	if fields["value"] != nil {
		if string(fields["value"]) != "null" {
			if err := json.Unmarshal(fields["value"], &resource.Value); err != nil {
				errs = append(errs, cog.MakeBuildErrors("value", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("value", errors.New("required field is null"))...)

		}
		delete(fields, "value")
	} else {
		errs = append(errs, cog.MakeBuildErrors("value", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("MovingAverageModelOption", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `MovingAverageModelOption` objects.
func (resource MovingAverageModelOption) Equals(other MovingAverageModelOption) bool {
	if resource.Label != other.Label {
		return false
	}
	if resource.Value != other.Value {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `MovingAverageModelOption` fields for violations and returns them.
func (resource MovingAverageModelOption) Validate() error {
	return nil
}

type BaseMovingAverageModelSettings struct {
	Model   MovingAverageModel `json:"model"`
	Window  string             `json:"window"`
	Predict string             `json:"predict"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `BaseMovingAverageModelSettings` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *BaseMovingAverageModelSettings) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "model"
	if fields["model"] != nil {
		if string(fields["model"]) != "null" {
			if err := json.Unmarshal(fields["model"], &resource.Model); err != nil {
				errs = append(errs, cog.MakeBuildErrors("model", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("model", errors.New("required field is null"))...)

		}
		delete(fields, "model")
	} else {
		errs = append(errs, cog.MakeBuildErrors("model", errors.New("required field is missing from input"))...)
	}
	// Field "window"
	if fields["window"] != nil {
		if string(fields["window"]) != "null" {
			if err := json.Unmarshal(fields["window"], &resource.Window); err != nil {
				errs = append(errs, cog.MakeBuildErrors("window", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("window", errors.New("required field is null"))...)

		}
		delete(fields, "window")
	} else {
		errs = append(errs, cog.MakeBuildErrors("window", errors.New("required field is missing from input"))...)
	}
	// Field "predict"
	if fields["predict"] != nil {
		if string(fields["predict"]) != "null" {
			if err := json.Unmarshal(fields["predict"], &resource.Predict); err != nil {
				errs = append(errs, cog.MakeBuildErrors("predict", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("predict", errors.New("required field is null"))...)

		}
		delete(fields, "predict")
	} else {
		errs = append(errs, cog.MakeBuildErrors("predict", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("BaseMovingAverageModelSettings", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `BaseMovingAverageModelSettings` objects.
func (resource BaseMovingAverageModelSettings) Equals(other BaseMovingAverageModelSettings) bool {
	if resource.Model != other.Model {
		return false
	}
	if resource.Window != other.Window {
		return false
	}
	if resource.Predict != other.Predict {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `BaseMovingAverageModelSettings` fields for violations and returns them.
func (resource BaseMovingAverageModelSettings) Validate() error {
	return nil
}

type MovingAverageSimpleModelSettings struct {
	Model   string `json:"model"`
	Window  string `json:"window"`
	Predict string `json:"predict"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `MovingAverageSimpleModelSettings` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *MovingAverageSimpleModelSettings) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "model"
	if fields["model"] != nil {
		if string(fields["model"]) != "null" {
			if err := json.Unmarshal(fields["model"], &resource.Model); err != nil {
				errs = append(errs, cog.MakeBuildErrors("model", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("model", errors.New("required field is null"))...)

		}
		delete(fields, "model")
	} else {
		errs = append(errs, cog.MakeBuildErrors("model", errors.New("required field is missing from input"))...)
	}
	// Field "window"
	if fields["window"] != nil {
		if string(fields["window"]) != "null" {
			if err := json.Unmarshal(fields["window"], &resource.Window); err != nil {
				errs = append(errs, cog.MakeBuildErrors("window", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("window", errors.New("required field is null"))...)

		}
		delete(fields, "window")
	} else {
		errs = append(errs, cog.MakeBuildErrors("window", errors.New("required field is missing from input"))...)
	}
	// Field "predict"
	if fields["predict"] != nil {
		if string(fields["predict"]) != "null" {
			if err := json.Unmarshal(fields["predict"], &resource.Predict); err != nil {
				errs = append(errs, cog.MakeBuildErrors("predict", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("predict", errors.New("required field is null"))...)

		}
		delete(fields, "predict")
	} else {
		errs = append(errs, cog.MakeBuildErrors("predict", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("MovingAverageSimpleModelSettings", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `MovingAverageSimpleModelSettings` objects.
func (resource MovingAverageSimpleModelSettings) Equals(other MovingAverageSimpleModelSettings) bool {
	if resource.Model != other.Model {
		return false
	}
	if resource.Window != other.Window {
		return false
	}
	if resource.Predict != other.Predict {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `MovingAverageSimpleModelSettings` fields for violations and returns them.
func (resource MovingAverageSimpleModelSettings) Validate() error {
	var errs cog.BuildErrors
	if !(resource.Model == "simple") {
		errs = append(errs, cog.MakeBuildErrors(
			"model",
			errors.New("must be == simple"),
		)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type MovingAverageLinearModelSettings struct {
	Model   string `json:"model"`
	Window  string `json:"window"`
	Predict string `json:"predict"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `MovingAverageLinearModelSettings` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *MovingAverageLinearModelSettings) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "model"
	if fields["model"] != nil {
		if string(fields["model"]) != "null" {
			if err := json.Unmarshal(fields["model"], &resource.Model); err != nil {
				errs = append(errs, cog.MakeBuildErrors("model", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("model", errors.New("required field is null"))...)

		}
		delete(fields, "model")
	} else {
		errs = append(errs, cog.MakeBuildErrors("model", errors.New("required field is missing from input"))...)
	}
	// Field "window"
	if fields["window"] != nil {
		if string(fields["window"]) != "null" {
			if err := json.Unmarshal(fields["window"], &resource.Window); err != nil {
				errs = append(errs, cog.MakeBuildErrors("window", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("window", errors.New("required field is null"))...)

		}
		delete(fields, "window")
	} else {
		errs = append(errs, cog.MakeBuildErrors("window", errors.New("required field is missing from input"))...)
	}
	// Field "predict"
	if fields["predict"] != nil {
		if string(fields["predict"]) != "null" {
			if err := json.Unmarshal(fields["predict"], &resource.Predict); err != nil {
				errs = append(errs, cog.MakeBuildErrors("predict", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("predict", errors.New("required field is null"))...)

		}
		delete(fields, "predict")
	} else {
		errs = append(errs, cog.MakeBuildErrors("predict", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("MovingAverageLinearModelSettings", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `MovingAverageLinearModelSettings` objects.
func (resource MovingAverageLinearModelSettings) Equals(other MovingAverageLinearModelSettings) bool {
	if resource.Model != other.Model {
		return false
	}
	if resource.Window != other.Window {
		return false
	}
	if resource.Predict != other.Predict {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `MovingAverageLinearModelSettings` fields for violations and returns them.
func (resource MovingAverageLinearModelSettings) Validate() error {
	var errs cog.BuildErrors
	if !(resource.Model == "linear") {
		errs = append(errs, cog.MakeBuildErrors(
			"model",
			errors.New("must be == linear"),
		)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type MovingAverageEWMAModelSettings struct {
	Model    string                                               `json:"model"`
	Settings *ElasticsearchMovingAverageEWMAModelSettingsSettings `json:"settings,omitempty"`
	Window   string                                               `json:"window"`
	Minimize bool                                                 `json:"minimize"`
	Predict  string                                               `json:"predict"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `MovingAverageEWMAModelSettings` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *MovingAverageEWMAModelSettings) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "model"
	if fields["model"] != nil {
		if string(fields["model"]) != "null" {
			if err := json.Unmarshal(fields["model"], &resource.Model); err != nil {
				errs = append(errs, cog.MakeBuildErrors("model", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("model", errors.New("required field is null"))...)

		}
		delete(fields, "model")
	} else {
		errs = append(errs, cog.MakeBuildErrors("model", errors.New("required field is missing from input"))...)
	}
	// Field "settings"
	if fields["settings"] != nil {
		if string(fields["settings"]) != "null" {

			resource.Settings = &ElasticsearchMovingAverageEWMAModelSettingsSettings{}
			if err := resource.Settings.UnmarshalJSONStrict(fields["settings"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("settings", err)...)
			}

		}
		delete(fields, "settings")

	}
	// Field "window"
	if fields["window"] != nil {
		if string(fields["window"]) != "null" {
			if err := json.Unmarshal(fields["window"], &resource.Window); err != nil {
				errs = append(errs, cog.MakeBuildErrors("window", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("window", errors.New("required field is null"))...)

		}
		delete(fields, "window")
	} else {
		errs = append(errs, cog.MakeBuildErrors("window", errors.New("required field is missing from input"))...)
	}
	// Field "minimize"
	if fields["minimize"] != nil {
		if string(fields["minimize"]) != "null" {
			if err := json.Unmarshal(fields["minimize"], &resource.Minimize); err != nil {
				errs = append(errs, cog.MakeBuildErrors("minimize", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("minimize", errors.New("required field is null"))...)

		}
		delete(fields, "minimize")
	} else {
		errs = append(errs, cog.MakeBuildErrors("minimize", errors.New("required field is missing from input"))...)
	}
	// Field "predict"
	if fields["predict"] != nil {
		if string(fields["predict"]) != "null" {
			if err := json.Unmarshal(fields["predict"], &resource.Predict); err != nil {
				errs = append(errs, cog.MakeBuildErrors("predict", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("predict", errors.New("required field is null"))...)

		}
		delete(fields, "predict")
	} else {
		errs = append(errs, cog.MakeBuildErrors("predict", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("MovingAverageEWMAModelSettings", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `MovingAverageEWMAModelSettings` objects.
func (resource MovingAverageEWMAModelSettings) Equals(other MovingAverageEWMAModelSettings) bool {
	if resource.Model != other.Model {
		return false
	}
	if resource.Settings == nil && other.Settings != nil || resource.Settings != nil && other.Settings == nil {
		return false
	}

	if resource.Settings != nil {
		if !resource.Settings.Equals(*other.Settings) {
			return false
		}
	}
	if resource.Window != other.Window {
		return false
	}
	if resource.Minimize != other.Minimize {
		return false
	}
	if resource.Predict != other.Predict {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `MovingAverageEWMAModelSettings` fields for violations and returns them.
func (resource MovingAverageEWMAModelSettings) Validate() error {
	var errs cog.BuildErrors
	if !(resource.Model == "ewma") {
		errs = append(errs, cog.MakeBuildErrors(
			"model",
			errors.New("must be == ewma"),
		)...)
	}
	if resource.Settings != nil {
		if err := resource.Settings.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("settings", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type MovingAverageHoltModelSettings struct {
	Model    string                                              `json:"model"`
	Settings ElasticsearchMovingAverageHoltModelSettingsSettings `json:"settings"`
	Window   string                                              `json:"window"`
	Minimize bool                                                `json:"minimize"`
	Predict  string                                              `json:"predict"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `MovingAverageHoltModelSettings` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *MovingAverageHoltModelSettings) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "model"
	if fields["model"] != nil {
		if string(fields["model"]) != "null" {
			if err := json.Unmarshal(fields["model"], &resource.Model); err != nil {
				errs = append(errs, cog.MakeBuildErrors("model", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("model", errors.New("required field is null"))...)

		}
		delete(fields, "model")
	} else {
		errs = append(errs, cog.MakeBuildErrors("model", errors.New("required field is missing from input"))...)
	}
	// Field "settings"
	if fields["settings"] != nil {
		if string(fields["settings"]) != "null" {

			resource.Settings = ElasticsearchMovingAverageHoltModelSettingsSettings{}
			if err := resource.Settings.UnmarshalJSONStrict(fields["settings"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("settings", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("settings", errors.New("required field is null"))...)

		}
		delete(fields, "settings")
	} else {
		errs = append(errs, cog.MakeBuildErrors("settings", errors.New("required field is missing from input"))...)
	}
	// Field "window"
	if fields["window"] != nil {
		if string(fields["window"]) != "null" {
			if err := json.Unmarshal(fields["window"], &resource.Window); err != nil {
				errs = append(errs, cog.MakeBuildErrors("window", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("window", errors.New("required field is null"))...)

		}
		delete(fields, "window")
	} else {
		errs = append(errs, cog.MakeBuildErrors("window", errors.New("required field is missing from input"))...)
	}
	// Field "minimize"
	if fields["minimize"] != nil {
		if string(fields["minimize"]) != "null" {
			if err := json.Unmarshal(fields["minimize"], &resource.Minimize); err != nil {
				errs = append(errs, cog.MakeBuildErrors("minimize", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("minimize", errors.New("required field is null"))...)

		}
		delete(fields, "minimize")
	} else {
		errs = append(errs, cog.MakeBuildErrors("minimize", errors.New("required field is missing from input"))...)
	}
	// Field "predict"
	if fields["predict"] != nil {
		if string(fields["predict"]) != "null" {
			if err := json.Unmarshal(fields["predict"], &resource.Predict); err != nil {
				errs = append(errs, cog.MakeBuildErrors("predict", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("predict", errors.New("required field is null"))...)

		}
		delete(fields, "predict")
	} else {
		errs = append(errs, cog.MakeBuildErrors("predict", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("MovingAverageHoltModelSettings", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `MovingAverageHoltModelSettings` objects.
func (resource MovingAverageHoltModelSettings) Equals(other MovingAverageHoltModelSettings) bool {
	if resource.Model != other.Model {
		return false
	}
	if !resource.Settings.Equals(other.Settings) {
		return false
	}
	if resource.Window != other.Window {
		return false
	}
	if resource.Minimize != other.Minimize {
		return false
	}
	if resource.Predict != other.Predict {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `MovingAverageHoltModelSettings` fields for violations and returns them.
func (resource MovingAverageHoltModelSettings) Validate() error {
	var errs cog.BuildErrors
	if !(resource.Model == "holt") {
		errs = append(errs, cog.MakeBuildErrors(
			"model",
			errors.New("must be == holt"),
		)...)
	}
	if err := resource.Settings.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("settings", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type MovingAverageHoltWintersModelSettings struct {
	Model    string                                                     `json:"model"`
	Settings ElasticsearchMovingAverageHoltWintersModelSettingsSettings `json:"settings"`
	Window   string                                                     `json:"window"`
	Minimize bool                                                       `json:"minimize"`
	Predict  string                                                     `json:"predict"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `MovingAverageHoltWintersModelSettings` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *MovingAverageHoltWintersModelSettings) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "model"
	if fields["model"] != nil {
		if string(fields["model"]) != "null" {
			if err := json.Unmarshal(fields["model"], &resource.Model); err != nil {
				errs = append(errs, cog.MakeBuildErrors("model", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("model", errors.New("required field is null"))...)

		}
		delete(fields, "model")
	} else {
		errs = append(errs, cog.MakeBuildErrors("model", errors.New("required field is missing from input"))...)
	}
	// Field "settings"
	if fields["settings"] != nil {
		if string(fields["settings"]) != "null" {

			resource.Settings = ElasticsearchMovingAverageHoltWintersModelSettingsSettings{}
			if err := resource.Settings.UnmarshalJSONStrict(fields["settings"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("settings", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("settings", errors.New("required field is null"))...)

		}
		delete(fields, "settings")
	} else {
		errs = append(errs, cog.MakeBuildErrors("settings", errors.New("required field is missing from input"))...)
	}
	// Field "window"
	if fields["window"] != nil {
		if string(fields["window"]) != "null" {
			if err := json.Unmarshal(fields["window"], &resource.Window); err != nil {
				errs = append(errs, cog.MakeBuildErrors("window", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("window", errors.New("required field is null"))...)

		}
		delete(fields, "window")
	} else {
		errs = append(errs, cog.MakeBuildErrors("window", errors.New("required field is missing from input"))...)
	}
	// Field "minimize"
	if fields["minimize"] != nil {
		if string(fields["minimize"]) != "null" {
			if err := json.Unmarshal(fields["minimize"], &resource.Minimize); err != nil {
				errs = append(errs, cog.MakeBuildErrors("minimize", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("minimize", errors.New("required field is null"))...)

		}
		delete(fields, "minimize")
	} else {
		errs = append(errs, cog.MakeBuildErrors("minimize", errors.New("required field is missing from input"))...)
	}
	// Field "predict"
	if fields["predict"] != nil {
		if string(fields["predict"]) != "null" {
			if err := json.Unmarshal(fields["predict"], &resource.Predict); err != nil {
				errs = append(errs, cog.MakeBuildErrors("predict", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("predict", errors.New("required field is null"))...)

		}
		delete(fields, "predict")
	} else {
		errs = append(errs, cog.MakeBuildErrors("predict", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("MovingAverageHoltWintersModelSettings", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `MovingAverageHoltWintersModelSettings` objects.
func (resource MovingAverageHoltWintersModelSettings) Equals(other MovingAverageHoltWintersModelSettings) bool {
	if resource.Model != other.Model {
		return false
	}
	if !resource.Settings.Equals(other.Settings) {
		return false
	}
	if resource.Window != other.Window {
		return false
	}
	if resource.Minimize != other.Minimize {
		return false
	}
	if resource.Predict != other.Predict {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `MovingAverageHoltWintersModelSettings` fields for violations and returns them.
func (resource MovingAverageHoltWintersModelSettings) Validate() error {
	var errs cog.BuildErrors
	if !(resource.Model == "holt_winters") {
		errs = append(errs, cog.MakeBuildErrors(
			"model",
			errors.New("must be == holt_winters"),
		)...)
	}
	if err := resource.Settings.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("settings", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// #MovingAverage's settings are overridden in types.ts
type MovingAverage struct {
	PipelineAgg *string        `json:"pipelineAgg,omitempty"`
	Field       *string        `json:"field,omitempty"`
	Type        string         `json:"type"`
	Id          string         `json:"id"`
	Settings    map[string]any `json:"settings,omitempty"`
	Hide        *bool          `json:"hide,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `MovingAverage` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *MovingAverage) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "pipelineAgg"
	if fields["pipelineAgg"] != nil {
		if string(fields["pipelineAgg"]) != "null" {
			if err := json.Unmarshal(fields["pipelineAgg"], &resource.PipelineAgg); err != nil {
				errs = append(errs, cog.MakeBuildErrors("pipelineAgg", err)...)
			}

		}
		delete(fields, "pipelineAgg")

	}
	// Field "field"
	if fields["field"] != nil {
		if string(fields["field"]) != "null" {
			if err := json.Unmarshal(fields["field"], &resource.Field); err != nil {
				errs = append(errs, cog.MakeBuildErrors("field", err)...)
			}

		}
		delete(fields, "field")

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
	// Field "settings"
	if fields["settings"] != nil {
		if string(fields["settings"]) != "null" {

			if err := json.Unmarshal(fields["settings"], &resource.Settings); err != nil {
				errs = append(errs, cog.MakeBuildErrors("settings", err)...)
			}

		}
		delete(fields, "settings")

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

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("MovingAverage", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `MovingAverage` objects.
func (resource MovingAverage) Equals(other MovingAverage) bool {
	if resource.PipelineAgg == nil && other.PipelineAgg != nil || resource.PipelineAgg != nil && other.PipelineAgg == nil {
		return false
	}

	if resource.PipelineAgg != nil {
		if *resource.PipelineAgg != *other.PipelineAgg {
			return false
		}
	}
	if resource.Field == nil && other.Field != nil || resource.Field != nil && other.Field == nil {
		return false
	}

	if resource.Field != nil {
		if *resource.Field != *other.Field {
			return false
		}
	}
	if resource.Type != other.Type {
		return false
	}
	if resource.Id != other.Id {
		return false
	}

	if len(resource.Settings) != len(other.Settings) {
		return false
	}

	for key1 := range resource.Settings {
		// is DeepEqual good enough here?
		if !reflect.DeepEqual(resource.Settings[key1], other.Settings[key1]) {
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

	return true
}

// Validate checks all the validation constraints that may be defined on `MovingAverage` fields for violations and returns them.
func (resource MovingAverage) Validate() error {
	var errs cog.BuildErrors
	if !(resource.Type == "moving_avg") {
		errs = append(errs, cog.MakeBuildErrors(
			"type",
			errors.New("must be == moving_avg"),
		)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type MovingFunction struct {
	PipelineAgg *string                              `json:"pipelineAgg,omitempty"`
	Field       *string                              `json:"field,omitempty"`
	Type        string                               `json:"type"`
	Id          string                               `json:"id"`
	Settings    *ElasticsearchMovingFunctionSettings `json:"settings,omitempty"`
	Hide        *bool                                `json:"hide,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `MovingFunction` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *MovingFunction) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "pipelineAgg"
	if fields["pipelineAgg"] != nil {
		if string(fields["pipelineAgg"]) != "null" {
			if err := json.Unmarshal(fields["pipelineAgg"], &resource.PipelineAgg); err != nil {
				errs = append(errs, cog.MakeBuildErrors("pipelineAgg", err)...)
			}

		}
		delete(fields, "pipelineAgg")

	}
	// Field "field"
	if fields["field"] != nil {
		if string(fields["field"]) != "null" {
			if err := json.Unmarshal(fields["field"], &resource.Field); err != nil {
				errs = append(errs, cog.MakeBuildErrors("field", err)...)
			}

		}
		delete(fields, "field")

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
	// Field "settings"
	if fields["settings"] != nil {
		if string(fields["settings"]) != "null" {

			resource.Settings = &ElasticsearchMovingFunctionSettings{}
			if err := resource.Settings.UnmarshalJSONStrict(fields["settings"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("settings", err)...)
			}

		}
		delete(fields, "settings")

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

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("MovingFunction", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `MovingFunction` objects.
func (resource MovingFunction) Equals(other MovingFunction) bool {
	if resource.PipelineAgg == nil && other.PipelineAgg != nil || resource.PipelineAgg != nil && other.PipelineAgg == nil {
		return false
	}

	if resource.PipelineAgg != nil {
		if *resource.PipelineAgg != *other.PipelineAgg {
			return false
		}
	}
	if resource.Field == nil && other.Field != nil || resource.Field != nil && other.Field == nil {
		return false
	}

	if resource.Field != nil {
		if *resource.Field != *other.Field {
			return false
		}
	}
	if resource.Type != other.Type {
		return false
	}
	if resource.Id != other.Id {
		return false
	}
	if resource.Settings == nil && other.Settings != nil || resource.Settings != nil && other.Settings == nil {
		return false
	}

	if resource.Settings != nil {
		if !resource.Settings.Equals(*other.Settings) {
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

	return true
}

// Validate checks all the validation constraints that may be defined on `MovingFunction` fields for violations and returns them.
func (resource MovingFunction) Validate() error {
	var errs cog.BuildErrors
	if !(resource.Type == "moving_fn") {
		errs = append(errs, cog.MakeBuildErrors(
			"type",
			errors.New("must be == moving_fn"),
		)...)
	}
	if resource.Settings != nil {
		if err := resource.Settings.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("settings", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type Derivative struct {
	PipelineAgg *string                          `json:"pipelineAgg,omitempty"`
	Field       *string                          `json:"field,omitempty"`
	Type        string                           `json:"type"`
	Id          string                           `json:"id"`
	Settings    *ElasticsearchDerivativeSettings `json:"settings,omitempty"`
	Hide        *bool                            `json:"hide,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Derivative` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *Derivative) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "pipelineAgg"
	if fields["pipelineAgg"] != nil {
		if string(fields["pipelineAgg"]) != "null" {
			if err := json.Unmarshal(fields["pipelineAgg"], &resource.PipelineAgg); err != nil {
				errs = append(errs, cog.MakeBuildErrors("pipelineAgg", err)...)
			}

		}
		delete(fields, "pipelineAgg")

	}
	// Field "field"
	if fields["field"] != nil {
		if string(fields["field"]) != "null" {
			if err := json.Unmarshal(fields["field"], &resource.Field); err != nil {
				errs = append(errs, cog.MakeBuildErrors("field", err)...)
			}

		}
		delete(fields, "field")

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
	// Field "settings"
	if fields["settings"] != nil {
		if string(fields["settings"]) != "null" {

			resource.Settings = &ElasticsearchDerivativeSettings{}
			if err := resource.Settings.UnmarshalJSONStrict(fields["settings"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("settings", err)...)
			}

		}
		delete(fields, "settings")

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

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("Derivative", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `Derivative` objects.
func (resource Derivative) Equals(other Derivative) bool {
	if resource.PipelineAgg == nil && other.PipelineAgg != nil || resource.PipelineAgg != nil && other.PipelineAgg == nil {
		return false
	}

	if resource.PipelineAgg != nil {
		if *resource.PipelineAgg != *other.PipelineAgg {
			return false
		}
	}
	if resource.Field == nil && other.Field != nil || resource.Field != nil && other.Field == nil {
		return false
	}

	if resource.Field != nil {
		if *resource.Field != *other.Field {
			return false
		}
	}
	if resource.Type != other.Type {
		return false
	}
	if resource.Id != other.Id {
		return false
	}
	if resource.Settings == nil && other.Settings != nil || resource.Settings != nil && other.Settings == nil {
		return false
	}

	if resource.Settings != nil {
		if !resource.Settings.Equals(*other.Settings) {
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

	return true
}

// Validate checks all the validation constraints that may be defined on `Derivative` fields for violations and returns them.
func (resource Derivative) Validate() error {
	var errs cog.BuildErrors
	if !(resource.Type == "derivative") {
		errs = append(errs, cog.MakeBuildErrors(
			"type",
			errors.New("must be == derivative"),
		)...)
	}
	if resource.Settings != nil {
		if err := resource.Settings.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("settings", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type SerialDiff struct {
	PipelineAgg *string                          `json:"pipelineAgg,omitempty"`
	Field       *string                          `json:"field,omitempty"`
	Type        string                           `json:"type"`
	Id          string                           `json:"id"`
	Settings    *ElasticsearchSerialDiffSettings `json:"settings,omitempty"`
	Hide        *bool                            `json:"hide,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `SerialDiff` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *SerialDiff) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "pipelineAgg"
	if fields["pipelineAgg"] != nil {
		if string(fields["pipelineAgg"]) != "null" {
			if err := json.Unmarshal(fields["pipelineAgg"], &resource.PipelineAgg); err != nil {
				errs = append(errs, cog.MakeBuildErrors("pipelineAgg", err)...)
			}

		}
		delete(fields, "pipelineAgg")

	}
	// Field "field"
	if fields["field"] != nil {
		if string(fields["field"]) != "null" {
			if err := json.Unmarshal(fields["field"], &resource.Field); err != nil {
				errs = append(errs, cog.MakeBuildErrors("field", err)...)
			}

		}
		delete(fields, "field")

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
	// Field "settings"
	if fields["settings"] != nil {
		if string(fields["settings"]) != "null" {

			resource.Settings = &ElasticsearchSerialDiffSettings{}
			if err := resource.Settings.UnmarshalJSONStrict(fields["settings"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("settings", err)...)
			}

		}
		delete(fields, "settings")

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

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("SerialDiff", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `SerialDiff` objects.
func (resource SerialDiff) Equals(other SerialDiff) bool {
	if resource.PipelineAgg == nil && other.PipelineAgg != nil || resource.PipelineAgg != nil && other.PipelineAgg == nil {
		return false
	}

	if resource.PipelineAgg != nil {
		if *resource.PipelineAgg != *other.PipelineAgg {
			return false
		}
	}
	if resource.Field == nil && other.Field != nil || resource.Field != nil && other.Field == nil {
		return false
	}

	if resource.Field != nil {
		if *resource.Field != *other.Field {
			return false
		}
	}
	if resource.Type != other.Type {
		return false
	}
	if resource.Id != other.Id {
		return false
	}
	if resource.Settings == nil && other.Settings != nil || resource.Settings != nil && other.Settings == nil {
		return false
	}

	if resource.Settings != nil {
		if !resource.Settings.Equals(*other.Settings) {
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

	return true
}

// Validate checks all the validation constraints that may be defined on `SerialDiff` fields for violations and returns them.
func (resource SerialDiff) Validate() error {
	var errs cog.BuildErrors
	if !(resource.Type == "serial_diff") {
		errs = append(errs, cog.MakeBuildErrors(
			"type",
			errors.New("must be == serial_diff"),
		)...)
	}
	if resource.Settings != nil {
		if err := resource.Settings.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("settings", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type CumulativeSum struct {
	PipelineAgg *string                             `json:"pipelineAgg,omitempty"`
	Field       *string                             `json:"field,omitempty"`
	Type        string                              `json:"type"`
	Id          string                              `json:"id"`
	Settings    *ElasticsearchCumulativeSumSettings `json:"settings,omitempty"`
	Hide        *bool                               `json:"hide,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `CumulativeSum` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *CumulativeSum) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "pipelineAgg"
	if fields["pipelineAgg"] != nil {
		if string(fields["pipelineAgg"]) != "null" {
			if err := json.Unmarshal(fields["pipelineAgg"], &resource.PipelineAgg); err != nil {
				errs = append(errs, cog.MakeBuildErrors("pipelineAgg", err)...)
			}

		}
		delete(fields, "pipelineAgg")

	}
	// Field "field"
	if fields["field"] != nil {
		if string(fields["field"]) != "null" {
			if err := json.Unmarshal(fields["field"], &resource.Field); err != nil {
				errs = append(errs, cog.MakeBuildErrors("field", err)...)
			}

		}
		delete(fields, "field")

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
	// Field "settings"
	if fields["settings"] != nil {
		if string(fields["settings"]) != "null" {

			resource.Settings = &ElasticsearchCumulativeSumSettings{}
			if err := resource.Settings.UnmarshalJSONStrict(fields["settings"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("settings", err)...)
			}

		}
		delete(fields, "settings")

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

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("CumulativeSum", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `CumulativeSum` objects.
func (resource CumulativeSum) Equals(other CumulativeSum) bool {
	if resource.PipelineAgg == nil && other.PipelineAgg != nil || resource.PipelineAgg != nil && other.PipelineAgg == nil {
		return false
	}

	if resource.PipelineAgg != nil {
		if *resource.PipelineAgg != *other.PipelineAgg {
			return false
		}
	}
	if resource.Field == nil && other.Field != nil || resource.Field != nil && other.Field == nil {
		return false
	}

	if resource.Field != nil {
		if *resource.Field != *other.Field {
			return false
		}
	}
	if resource.Type != other.Type {
		return false
	}
	if resource.Id != other.Id {
		return false
	}
	if resource.Settings == nil && other.Settings != nil || resource.Settings != nil && other.Settings == nil {
		return false
	}

	if resource.Settings != nil {
		if !resource.Settings.Equals(*other.Settings) {
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

	return true
}

// Validate checks all the validation constraints that may be defined on `CumulativeSum` fields for violations and returns them.
func (resource CumulativeSum) Validate() error {
	var errs cog.BuildErrors
	if !(resource.Type == "cumulative_sum") {
		errs = append(errs, cog.MakeBuildErrors(
			"type",
			errors.New("must be == cumulative_sum"),
		)...)
	}
	if resource.Settings != nil {
		if err := resource.Settings.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("settings", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type BucketScript struct {
	Type              string                             `json:"type"`
	PipelineVariables []PipelineVariable                 `json:"pipelineVariables,omitempty"`
	Id                string                             `json:"id"`
	Settings          *ElasticsearchBucketScriptSettings `json:"settings,omitempty"`
	Hide              *bool                              `json:"hide,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `BucketScript` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *BucketScript) UnmarshalJSONStrict(raw []byte) error {
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
	// Field "pipelineVariables"
	if fields["pipelineVariables"] != nil {
		if string(fields["pipelineVariables"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["pipelineVariables"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 PipelineVariable

				result1 = PipelineVariable{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("pipelineVariables["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.PipelineVariables = append(resource.PipelineVariables, result1)
			}

		}
		delete(fields, "pipelineVariables")

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
	// Field "settings"
	if fields["settings"] != nil {
		if string(fields["settings"]) != "null" {

			resource.Settings = &ElasticsearchBucketScriptSettings{}
			if err := resource.Settings.UnmarshalJSONStrict(fields["settings"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("settings", err)...)
			}

		}
		delete(fields, "settings")

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

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("BucketScript", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `BucketScript` objects.
func (resource BucketScript) Equals(other BucketScript) bool {
	if resource.Type != other.Type {
		return false
	}

	if len(resource.PipelineVariables) != len(other.PipelineVariables) {
		return false
	}

	for i1 := range resource.PipelineVariables {
		if !resource.PipelineVariables[i1].Equals(other.PipelineVariables[i1]) {
			return false
		}
	}
	if resource.Id != other.Id {
		return false
	}
	if resource.Settings == nil && other.Settings != nil || resource.Settings != nil && other.Settings == nil {
		return false
	}

	if resource.Settings != nil {
		if !resource.Settings.Equals(*other.Settings) {
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

	return true
}

// Validate checks all the validation constraints that may be defined on `BucketScript` fields for violations and returns them.
func (resource BucketScript) Validate() error {
	var errs cog.BuildErrors
	if !(resource.Type == "bucket_script") {
		errs = append(errs, cog.MakeBuildErrors(
			"type",
			errors.New("must be == bucket_script"),
		)...)
	}

	for i1 := range resource.PipelineVariables {
		if err := resource.PipelineVariables[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("pipelineVariables["+strconv.Itoa(i1)+"]", err)...)
		}
	}
	if resource.Settings != nil {
		if err := resource.Settings.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("settings", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type TopMetrics struct {
	Type     string                           `json:"type"`
	Id       string                           `json:"id"`
	Settings *ElasticsearchTopMetricsSettings `json:"settings,omitempty"`
	Hide     *bool                            `json:"hide,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TopMetrics` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *TopMetrics) UnmarshalJSONStrict(raw []byte) error {
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
	// Field "settings"
	if fields["settings"] != nil {
		if string(fields["settings"]) != "null" {

			resource.Settings = &ElasticsearchTopMetricsSettings{}
			if err := resource.Settings.UnmarshalJSONStrict(fields["settings"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("settings", err)...)
			}

		}
		delete(fields, "settings")

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

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("TopMetrics", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `TopMetrics` objects.
func (resource TopMetrics) Equals(other TopMetrics) bool {
	if resource.Type != other.Type {
		return false
	}
	if resource.Id != other.Id {
		return false
	}
	if resource.Settings == nil && other.Settings != nil || resource.Settings != nil && other.Settings == nil {
		return false
	}

	if resource.Settings != nil {
		if !resource.Settings.Equals(*other.Settings) {
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

	return true
}

// Validate checks all the validation constraints that may be defined on `TopMetrics` fields for violations and returns them.
func (resource TopMetrics) Validate() error {
	var errs cog.BuildErrors
	if !(resource.Type == "top_metrics") {
		errs = append(errs, cog.MakeBuildErrors(
			"type",
			errors.New("must be == top_metrics"),
		)...)
	}
	if resource.Settings != nil {
		if err := resource.Settings.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("settings", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type PipelineMetricAggregation = MovingAverageOrDerivativeOrCumulativeSumOrBucketScript

type MetricAggregationWithSettings = BucketScriptOrCumulativeSumOrDerivativeOrSerialDiffOrRawDataOrRawDocumentOrUniqueCountOrPercentilesOrExtendedStatsOrMinOrMaxOrSumOrAverageOrMovingAverageOrMovingFunctionOrLogsOrRateOrTopMetrics

type Dataquery struct {
	// Alias pattern
	Alias *string `json:"alias,omitempty"`
	// Lucene query
	Query *string `json:"query,omitempty"`
	// Name of time field
	TimeField *string `json:"timeField,omitempty"`
	// List of bucket aggregations
	BucketAggs []BucketAggregation `json:"bucketAggs,omitempty"`
	// List of metric aggregations
	Metrics []MetricAggregation `json:"metrics,omitempty"`
	// A unique identifier for the query within the list of targets.
	// In server side expressions, the refId is used as a variable name to identify results.
	// By default, the UI will assign A->Z; however setting meaningful names may be useful.
	RefId string `json:"refId"`
	// If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
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
	return "elasticsearch"
}

// VariantConfig returns the configuration related to elasticsearch dataqueries.
// This configuration describes how to unmarshal it, convert it to code, …
func VariantConfig() variants.DataqueryConfig {
	return variants.DataqueryConfig{
		Identifier: "elasticsearch",
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
	// Field "query"
	if fields["query"] != nil {
		if string(fields["query"]) != "null" {
			if err := json.Unmarshal(fields["query"], &resource.Query); err != nil {
				errs = append(errs, cog.MakeBuildErrors("query", err)...)
			}

		}
		delete(fields, "query")

	}
	// Field "timeField"
	if fields["timeField"] != nil {
		if string(fields["timeField"]) != "null" {
			if err := json.Unmarshal(fields["timeField"], &resource.TimeField); err != nil {
				errs = append(errs, cog.MakeBuildErrors("timeField", err)...)
			}

		}
		delete(fields, "timeField")

	}
	// Field "bucketAggs"
	if fields["bucketAggs"] != nil {
		if string(fields["bucketAggs"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["bucketAggs"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 BucketAggregation

				result1 = BucketAggregation{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("bucketAggs["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.BucketAggs = append(resource.BucketAggs, result1)
			}

		}
		delete(fields, "bucketAggs")

	}
	// Field "metrics"
	if fields["metrics"] != nil {
		if string(fields["metrics"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["metrics"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 MetricAggregation

				result1 = MetricAggregation{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("metrics["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.Metrics = append(resource.Metrics, result1)
			}

		}
		delete(fields, "metrics")

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
	if resource.Query == nil && other.Query != nil || resource.Query != nil && other.Query == nil {
		return false
	}

	if resource.Query != nil {
		if *resource.Query != *other.Query {
			return false
		}
	}
	if resource.TimeField == nil && other.TimeField != nil || resource.TimeField != nil && other.TimeField == nil {
		return false
	}

	if resource.TimeField != nil {
		if *resource.TimeField != *other.TimeField {
			return false
		}
	}

	if len(resource.BucketAggs) != len(other.BucketAggs) {
		return false
	}

	for i1 := range resource.BucketAggs {
		if !resource.BucketAggs[i1].Equals(other.BucketAggs[i1]) {
			return false
		}
	}

	if len(resource.Metrics) != len(other.Metrics) {
		return false
	}

	for i1 := range resource.Metrics {
		if !resource.Metrics[i1].Equals(other.Metrics[i1]) {
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

	for i1 := range resource.BucketAggs {
		if err := resource.BucketAggs[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("bucketAggs["+strconv.Itoa(i1)+"]", err)...)
		}
	}

	for i1 := range resource.Metrics {
		if err := resource.Metrics[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("metrics["+strconv.Itoa(i1)+"]", err)...)
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

type ElasticsearchDateHistogramSettings struct {
	Interval    *string `json:"interval,omitempty"`
	MinDocCount *string `json:"min_doc_count,omitempty"`
	TrimEdges   *string `json:"trimEdges,omitempty"`
	Offset      *string `json:"offset,omitempty"`
	TimeZone    *string `json:"timeZone,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ElasticsearchDateHistogramSettings` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ElasticsearchDateHistogramSettings) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "interval"
	if fields["interval"] != nil {
		if string(fields["interval"]) != "null" {
			if err := json.Unmarshal(fields["interval"], &resource.Interval); err != nil {
				errs = append(errs, cog.MakeBuildErrors("interval", err)...)
			}

		}
		delete(fields, "interval")

	}
	// Field "min_doc_count"
	if fields["min_doc_count"] != nil {
		if string(fields["min_doc_count"]) != "null" {
			if err := json.Unmarshal(fields["min_doc_count"], &resource.MinDocCount); err != nil {
				errs = append(errs, cog.MakeBuildErrors("min_doc_count", err)...)
			}

		}
		delete(fields, "min_doc_count")

	}
	// Field "trimEdges"
	if fields["trimEdges"] != nil {
		if string(fields["trimEdges"]) != "null" {
			if err := json.Unmarshal(fields["trimEdges"], &resource.TrimEdges); err != nil {
				errs = append(errs, cog.MakeBuildErrors("trimEdges", err)...)
			}

		}
		delete(fields, "trimEdges")

	}
	// Field "offset"
	if fields["offset"] != nil {
		if string(fields["offset"]) != "null" {
			if err := json.Unmarshal(fields["offset"], &resource.Offset); err != nil {
				errs = append(errs, cog.MakeBuildErrors("offset", err)...)
			}

		}
		delete(fields, "offset")

	}
	// Field "timeZone"
	if fields["timeZone"] != nil {
		if string(fields["timeZone"]) != "null" {
			if err := json.Unmarshal(fields["timeZone"], &resource.TimeZone); err != nil {
				errs = append(errs, cog.MakeBuildErrors("timeZone", err)...)
			}

		}
		delete(fields, "timeZone")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ElasticsearchDateHistogramSettings", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ElasticsearchDateHistogramSettings` objects.
func (resource ElasticsearchDateHistogramSettings) Equals(other ElasticsearchDateHistogramSettings) bool {
	if resource.Interval == nil && other.Interval != nil || resource.Interval != nil && other.Interval == nil {
		return false
	}

	if resource.Interval != nil {
		if *resource.Interval != *other.Interval {
			return false
		}
	}
	if resource.MinDocCount == nil && other.MinDocCount != nil || resource.MinDocCount != nil && other.MinDocCount == nil {
		return false
	}

	if resource.MinDocCount != nil {
		if *resource.MinDocCount != *other.MinDocCount {
			return false
		}
	}
	if resource.TrimEdges == nil && other.TrimEdges != nil || resource.TrimEdges != nil && other.TrimEdges == nil {
		return false
	}

	if resource.TrimEdges != nil {
		if *resource.TrimEdges != *other.TrimEdges {
			return false
		}
	}
	if resource.Offset == nil && other.Offset != nil || resource.Offset != nil && other.Offset == nil {
		return false
	}

	if resource.Offset != nil {
		if *resource.Offset != *other.Offset {
			return false
		}
	}
	if resource.TimeZone == nil && other.TimeZone != nil || resource.TimeZone != nil && other.TimeZone == nil {
		return false
	}

	if resource.TimeZone != nil {
		if *resource.TimeZone != *other.TimeZone {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `ElasticsearchDateHistogramSettings` fields for violations and returns them.
func (resource ElasticsearchDateHistogramSettings) Validate() error {
	return nil
}

type ElasticsearchHistogramSettings struct {
	Interval    *string `json:"interval,omitempty"`
	MinDocCount *string `json:"min_doc_count,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ElasticsearchHistogramSettings` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ElasticsearchHistogramSettings) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "interval"
	if fields["interval"] != nil {
		if string(fields["interval"]) != "null" {
			if err := json.Unmarshal(fields["interval"], &resource.Interval); err != nil {
				errs = append(errs, cog.MakeBuildErrors("interval", err)...)
			}

		}
		delete(fields, "interval")

	}
	// Field "min_doc_count"
	if fields["min_doc_count"] != nil {
		if string(fields["min_doc_count"]) != "null" {
			if err := json.Unmarshal(fields["min_doc_count"], &resource.MinDocCount); err != nil {
				errs = append(errs, cog.MakeBuildErrors("min_doc_count", err)...)
			}

		}
		delete(fields, "min_doc_count")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ElasticsearchHistogramSettings", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ElasticsearchHistogramSettings` objects.
func (resource ElasticsearchHistogramSettings) Equals(other ElasticsearchHistogramSettings) bool {
	if resource.Interval == nil && other.Interval != nil || resource.Interval != nil && other.Interval == nil {
		return false
	}

	if resource.Interval != nil {
		if *resource.Interval != *other.Interval {
			return false
		}
	}
	if resource.MinDocCount == nil && other.MinDocCount != nil || resource.MinDocCount != nil && other.MinDocCount == nil {
		return false
	}

	if resource.MinDocCount != nil {
		if *resource.MinDocCount != *other.MinDocCount {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `ElasticsearchHistogramSettings` fields for violations and returns them.
func (resource ElasticsearchHistogramSettings) Validate() error {
	return nil
}

type ElasticsearchTermsSettings struct {
	Order       *TermsOrder `json:"order,omitempty"`
	Size        *string     `json:"size,omitempty"`
	MinDocCount *string     `json:"min_doc_count,omitempty"`
	OrderBy     *string     `json:"orderBy,omitempty"`
	Missing     *string     `json:"missing,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ElasticsearchTermsSettings` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ElasticsearchTermsSettings) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "order"
	if fields["order"] != nil {
		if string(fields["order"]) != "null" {
			if err := json.Unmarshal(fields["order"], &resource.Order); err != nil {
				errs = append(errs, cog.MakeBuildErrors("order", err)...)
			}

		}
		delete(fields, "order")

	}
	// Field "size"
	if fields["size"] != nil {
		if string(fields["size"]) != "null" {
			if err := json.Unmarshal(fields["size"], &resource.Size); err != nil {
				errs = append(errs, cog.MakeBuildErrors("size", err)...)
			}

		}
		delete(fields, "size")

	}
	// Field "min_doc_count"
	if fields["min_doc_count"] != nil {
		if string(fields["min_doc_count"]) != "null" {
			if err := json.Unmarshal(fields["min_doc_count"], &resource.MinDocCount); err != nil {
				errs = append(errs, cog.MakeBuildErrors("min_doc_count", err)...)
			}

		}
		delete(fields, "min_doc_count")

	}
	// Field "orderBy"
	if fields["orderBy"] != nil {
		if string(fields["orderBy"]) != "null" {
			if err := json.Unmarshal(fields["orderBy"], &resource.OrderBy); err != nil {
				errs = append(errs, cog.MakeBuildErrors("orderBy", err)...)
			}

		}
		delete(fields, "orderBy")

	}
	// Field "missing"
	if fields["missing"] != nil {
		if string(fields["missing"]) != "null" {
			if err := json.Unmarshal(fields["missing"], &resource.Missing); err != nil {
				errs = append(errs, cog.MakeBuildErrors("missing", err)...)
			}

		}
		delete(fields, "missing")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ElasticsearchTermsSettings", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ElasticsearchTermsSettings` objects.
func (resource ElasticsearchTermsSettings) Equals(other ElasticsearchTermsSettings) bool {
	if resource.Order == nil && other.Order != nil || resource.Order != nil && other.Order == nil {
		return false
	}

	if resource.Order != nil {
		if *resource.Order != *other.Order {
			return false
		}
	}
	if resource.Size == nil && other.Size != nil || resource.Size != nil && other.Size == nil {
		return false
	}

	if resource.Size != nil {
		if *resource.Size != *other.Size {
			return false
		}
	}
	if resource.MinDocCount == nil && other.MinDocCount != nil || resource.MinDocCount != nil && other.MinDocCount == nil {
		return false
	}

	if resource.MinDocCount != nil {
		if *resource.MinDocCount != *other.MinDocCount {
			return false
		}
	}
	if resource.OrderBy == nil && other.OrderBy != nil || resource.OrderBy != nil && other.OrderBy == nil {
		return false
	}

	if resource.OrderBy != nil {
		if *resource.OrderBy != *other.OrderBy {
			return false
		}
	}
	if resource.Missing == nil && other.Missing != nil || resource.Missing != nil && other.Missing == nil {
		return false
	}

	if resource.Missing != nil {
		if *resource.Missing != *other.Missing {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `ElasticsearchTermsSettings` fields for violations and returns them.
func (resource ElasticsearchTermsSettings) Validate() error {
	return nil
}

type ElasticsearchFiltersSettings struct {
	Filters []Filter `json:"filters,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ElasticsearchFiltersSettings` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ElasticsearchFiltersSettings) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "filters"
	if fields["filters"] != nil {
		if string(fields["filters"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["filters"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 Filter

				result1 = Filter{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("filters["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.Filters = append(resource.Filters, result1)
			}

		}
		delete(fields, "filters")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ElasticsearchFiltersSettings", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ElasticsearchFiltersSettings` objects.
func (resource ElasticsearchFiltersSettings) Equals(other ElasticsearchFiltersSettings) bool {

	if len(resource.Filters) != len(other.Filters) {
		return false
	}

	for i1 := range resource.Filters {
		if !resource.Filters[i1].Equals(other.Filters[i1]) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `ElasticsearchFiltersSettings` fields for violations and returns them.
func (resource ElasticsearchFiltersSettings) Validate() error {
	var errs cog.BuildErrors

	for i1 := range resource.Filters {
		if err := resource.Filters[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("filters["+strconv.Itoa(i1)+"]", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type ElasticsearchGeoHashGridSettings struct {
	Precision *string `json:"precision,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ElasticsearchGeoHashGridSettings` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ElasticsearchGeoHashGridSettings) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "precision"
	if fields["precision"] != nil {
		if string(fields["precision"]) != "null" {
			if err := json.Unmarshal(fields["precision"], &resource.Precision); err != nil {
				errs = append(errs, cog.MakeBuildErrors("precision", err)...)
			}

		}
		delete(fields, "precision")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ElasticsearchGeoHashGridSettings", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ElasticsearchGeoHashGridSettings` objects.
func (resource ElasticsearchGeoHashGridSettings) Equals(other ElasticsearchGeoHashGridSettings) bool {
	if resource.Precision == nil && other.Precision != nil || resource.Precision != nil && other.Precision == nil {
		return false
	}

	if resource.Precision != nil {
		if *resource.Precision != *other.Precision {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `ElasticsearchGeoHashGridSettings` fields for violations and returns them.
func (resource ElasticsearchGeoHashGridSettings) Validate() error {
	return nil
}

type ElasticsearchMetricAggregationWithMissingSupportSettings struct {
	Missing *string `json:"missing,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ElasticsearchMetricAggregationWithMissingSupportSettings` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ElasticsearchMetricAggregationWithMissingSupportSettings) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "missing"
	if fields["missing"] != nil {
		if string(fields["missing"]) != "null" {
			if err := json.Unmarshal(fields["missing"], &resource.Missing); err != nil {
				errs = append(errs, cog.MakeBuildErrors("missing", err)...)
			}

		}
		delete(fields, "missing")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ElasticsearchMetricAggregationWithMissingSupportSettings", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ElasticsearchMetricAggregationWithMissingSupportSettings` objects.
func (resource ElasticsearchMetricAggregationWithMissingSupportSettings) Equals(other ElasticsearchMetricAggregationWithMissingSupportSettings) bool {
	if resource.Missing == nil && other.Missing != nil || resource.Missing != nil && other.Missing == nil {
		return false
	}

	if resource.Missing != nil {
		if *resource.Missing != *other.Missing {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `ElasticsearchMetricAggregationWithMissingSupportSettings` fields for violations and returns them.
func (resource ElasticsearchMetricAggregationWithMissingSupportSettings) Validate() error {
	return nil
}

type ElasticsearchInlineScript struct {
	Inline *string `json:"inline,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ElasticsearchInlineScript` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ElasticsearchInlineScript) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "inline"
	if fields["inline"] != nil {
		if string(fields["inline"]) != "null" {
			if err := json.Unmarshal(fields["inline"], &resource.Inline); err != nil {
				errs = append(errs, cog.MakeBuildErrors("inline", err)...)
			}

		}
		delete(fields, "inline")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ElasticsearchInlineScript", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ElasticsearchInlineScript` objects.
func (resource ElasticsearchInlineScript) Equals(other ElasticsearchInlineScript) bool {
	if resource.Inline == nil && other.Inline != nil || resource.Inline != nil && other.Inline == nil {
		return false
	}

	if resource.Inline != nil {
		if *resource.Inline != *other.Inline {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `ElasticsearchInlineScript` fields for violations and returns them.
func (resource ElasticsearchInlineScript) Validate() error {
	return nil
}

type ElasticsearchMetricAggregationWithInlineScriptSettings struct {
	Script *InlineScript `json:"script,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ElasticsearchMetricAggregationWithInlineScriptSettings` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ElasticsearchMetricAggregationWithInlineScriptSettings) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "script"
	if fields["script"] != nil {
		if string(fields["script"]) != "null" {

			resource.Script = &InlineScript{}
			if err := resource.Script.UnmarshalJSONStrict(fields["script"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("script", err)...)
			}

		}
		delete(fields, "script")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ElasticsearchMetricAggregationWithInlineScriptSettings", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ElasticsearchMetricAggregationWithInlineScriptSettings` objects.
func (resource ElasticsearchMetricAggregationWithInlineScriptSettings) Equals(other ElasticsearchMetricAggregationWithInlineScriptSettings) bool {
	if resource.Script == nil && other.Script != nil || resource.Script != nil && other.Script == nil {
		return false
	}

	if resource.Script != nil {
		if !resource.Script.Equals(*other.Script) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `ElasticsearchMetricAggregationWithInlineScriptSettings` fields for violations and returns them.
func (resource ElasticsearchMetricAggregationWithInlineScriptSettings) Validate() error {
	var errs cog.BuildErrors
	if resource.Script != nil {
		if err := resource.Script.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("script", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type ElasticsearchAverageSettings struct {
	Script  *InlineScript `json:"script,omitempty"`
	Missing *string       `json:"missing,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ElasticsearchAverageSettings` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ElasticsearchAverageSettings) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "script"
	if fields["script"] != nil {
		if string(fields["script"]) != "null" {

			resource.Script = &InlineScript{}
			if err := resource.Script.UnmarshalJSONStrict(fields["script"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("script", err)...)
			}

		}
		delete(fields, "script")

	}
	// Field "missing"
	if fields["missing"] != nil {
		if string(fields["missing"]) != "null" {
			if err := json.Unmarshal(fields["missing"], &resource.Missing); err != nil {
				errs = append(errs, cog.MakeBuildErrors("missing", err)...)
			}

		}
		delete(fields, "missing")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ElasticsearchAverageSettings", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ElasticsearchAverageSettings` objects.
func (resource ElasticsearchAverageSettings) Equals(other ElasticsearchAverageSettings) bool {
	if resource.Script == nil && other.Script != nil || resource.Script != nil && other.Script == nil {
		return false
	}

	if resource.Script != nil {
		if !resource.Script.Equals(*other.Script) {
			return false
		}
	}
	if resource.Missing == nil && other.Missing != nil || resource.Missing != nil && other.Missing == nil {
		return false
	}

	if resource.Missing != nil {
		if *resource.Missing != *other.Missing {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `ElasticsearchAverageSettings` fields for violations and returns them.
func (resource ElasticsearchAverageSettings) Validate() error {
	var errs cog.BuildErrors
	if resource.Script != nil {
		if err := resource.Script.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("script", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type ElasticsearchSumSettings struct {
	Script  *InlineScript `json:"script,omitempty"`
	Missing *string       `json:"missing,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ElasticsearchSumSettings` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ElasticsearchSumSettings) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "script"
	if fields["script"] != nil {
		if string(fields["script"]) != "null" {

			resource.Script = &InlineScript{}
			if err := resource.Script.UnmarshalJSONStrict(fields["script"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("script", err)...)
			}

		}
		delete(fields, "script")

	}
	// Field "missing"
	if fields["missing"] != nil {
		if string(fields["missing"]) != "null" {
			if err := json.Unmarshal(fields["missing"], &resource.Missing); err != nil {
				errs = append(errs, cog.MakeBuildErrors("missing", err)...)
			}

		}
		delete(fields, "missing")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ElasticsearchSumSettings", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ElasticsearchSumSettings` objects.
func (resource ElasticsearchSumSettings) Equals(other ElasticsearchSumSettings) bool {
	if resource.Script == nil && other.Script != nil || resource.Script != nil && other.Script == nil {
		return false
	}

	if resource.Script != nil {
		if !resource.Script.Equals(*other.Script) {
			return false
		}
	}
	if resource.Missing == nil && other.Missing != nil || resource.Missing != nil && other.Missing == nil {
		return false
	}

	if resource.Missing != nil {
		if *resource.Missing != *other.Missing {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `ElasticsearchSumSettings` fields for violations and returns them.
func (resource ElasticsearchSumSettings) Validate() error {
	var errs cog.BuildErrors
	if resource.Script != nil {
		if err := resource.Script.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("script", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type ElasticsearchMaxSettings struct {
	Script  *InlineScript `json:"script,omitempty"`
	Missing *string       `json:"missing,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ElasticsearchMaxSettings` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ElasticsearchMaxSettings) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "script"
	if fields["script"] != nil {
		if string(fields["script"]) != "null" {

			resource.Script = &InlineScript{}
			if err := resource.Script.UnmarshalJSONStrict(fields["script"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("script", err)...)
			}

		}
		delete(fields, "script")

	}
	// Field "missing"
	if fields["missing"] != nil {
		if string(fields["missing"]) != "null" {
			if err := json.Unmarshal(fields["missing"], &resource.Missing); err != nil {
				errs = append(errs, cog.MakeBuildErrors("missing", err)...)
			}

		}
		delete(fields, "missing")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ElasticsearchMaxSettings", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ElasticsearchMaxSettings` objects.
func (resource ElasticsearchMaxSettings) Equals(other ElasticsearchMaxSettings) bool {
	if resource.Script == nil && other.Script != nil || resource.Script != nil && other.Script == nil {
		return false
	}

	if resource.Script != nil {
		if !resource.Script.Equals(*other.Script) {
			return false
		}
	}
	if resource.Missing == nil && other.Missing != nil || resource.Missing != nil && other.Missing == nil {
		return false
	}

	if resource.Missing != nil {
		if *resource.Missing != *other.Missing {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `ElasticsearchMaxSettings` fields for violations and returns them.
func (resource ElasticsearchMaxSettings) Validate() error {
	var errs cog.BuildErrors
	if resource.Script != nil {
		if err := resource.Script.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("script", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type ElasticsearchMinSettings struct {
	Script  *InlineScript `json:"script,omitempty"`
	Missing *string       `json:"missing,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ElasticsearchMinSettings` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ElasticsearchMinSettings) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "script"
	if fields["script"] != nil {
		if string(fields["script"]) != "null" {

			resource.Script = &InlineScript{}
			if err := resource.Script.UnmarshalJSONStrict(fields["script"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("script", err)...)
			}

		}
		delete(fields, "script")

	}
	// Field "missing"
	if fields["missing"] != nil {
		if string(fields["missing"]) != "null" {
			if err := json.Unmarshal(fields["missing"], &resource.Missing); err != nil {
				errs = append(errs, cog.MakeBuildErrors("missing", err)...)
			}

		}
		delete(fields, "missing")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ElasticsearchMinSettings", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ElasticsearchMinSettings` objects.
func (resource ElasticsearchMinSettings) Equals(other ElasticsearchMinSettings) bool {
	if resource.Script == nil && other.Script != nil || resource.Script != nil && other.Script == nil {
		return false
	}

	if resource.Script != nil {
		if !resource.Script.Equals(*other.Script) {
			return false
		}
	}
	if resource.Missing == nil && other.Missing != nil || resource.Missing != nil && other.Missing == nil {
		return false
	}

	if resource.Missing != nil {
		if *resource.Missing != *other.Missing {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `ElasticsearchMinSettings` fields for violations and returns them.
func (resource ElasticsearchMinSettings) Validate() error {
	var errs cog.BuildErrors
	if resource.Script != nil {
		if err := resource.Script.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("script", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type ElasticsearchExtendedStatsSettings struct {
	Script  *InlineScript `json:"script,omitempty"`
	Missing *string       `json:"missing,omitempty"`
	Sigma   *string       `json:"sigma,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ElasticsearchExtendedStatsSettings` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ElasticsearchExtendedStatsSettings) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "script"
	if fields["script"] != nil {
		if string(fields["script"]) != "null" {

			resource.Script = &InlineScript{}
			if err := resource.Script.UnmarshalJSONStrict(fields["script"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("script", err)...)
			}

		}
		delete(fields, "script")

	}
	// Field "missing"
	if fields["missing"] != nil {
		if string(fields["missing"]) != "null" {
			if err := json.Unmarshal(fields["missing"], &resource.Missing); err != nil {
				errs = append(errs, cog.MakeBuildErrors("missing", err)...)
			}

		}
		delete(fields, "missing")

	}
	// Field "sigma"
	if fields["sigma"] != nil {
		if string(fields["sigma"]) != "null" {
			if err := json.Unmarshal(fields["sigma"], &resource.Sigma); err != nil {
				errs = append(errs, cog.MakeBuildErrors("sigma", err)...)
			}

		}
		delete(fields, "sigma")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ElasticsearchExtendedStatsSettings", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ElasticsearchExtendedStatsSettings` objects.
func (resource ElasticsearchExtendedStatsSettings) Equals(other ElasticsearchExtendedStatsSettings) bool {
	if resource.Script == nil && other.Script != nil || resource.Script != nil && other.Script == nil {
		return false
	}

	if resource.Script != nil {
		if !resource.Script.Equals(*other.Script) {
			return false
		}
	}
	if resource.Missing == nil && other.Missing != nil || resource.Missing != nil && other.Missing == nil {
		return false
	}

	if resource.Missing != nil {
		if *resource.Missing != *other.Missing {
			return false
		}
	}
	if resource.Sigma == nil && other.Sigma != nil || resource.Sigma != nil && other.Sigma == nil {
		return false
	}

	if resource.Sigma != nil {
		if *resource.Sigma != *other.Sigma {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `ElasticsearchExtendedStatsSettings` fields for violations and returns them.
func (resource ElasticsearchExtendedStatsSettings) Validate() error {
	var errs cog.BuildErrors
	if resource.Script != nil {
		if err := resource.Script.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("script", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type ElasticsearchPercentilesSettings struct {
	Script   *InlineScript `json:"script,omitempty"`
	Missing  *string       `json:"missing,omitempty"`
	Percents []string      `json:"percents,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ElasticsearchPercentilesSettings` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ElasticsearchPercentilesSettings) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "script"
	if fields["script"] != nil {
		if string(fields["script"]) != "null" {

			resource.Script = &InlineScript{}
			if err := resource.Script.UnmarshalJSONStrict(fields["script"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("script", err)...)
			}

		}
		delete(fields, "script")

	}
	// Field "missing"
	if fields["missing"] != nil {
		if string(fields["missing"]) != "null" {
			if err := json.Unmarshal(fields["missing"], &resource.Missing); err != nil {
				errs = append(errs, cog.MakeBuildErrors("missing", err)...)
			}

		}
		delete(fields, "missing")

	}
	// Field "percents"
	if fields["percents"] != nil {
		if string(fields["percents"]) != "null" {

			if err := json.Unmarshal(fields["percents"], &resource.Percents); err != nil {
				errs = append(errs, cog.MakeBuildErrors("percents", err)...)
			}

		}
		delete(fields, "percents")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ElasticsearchPercentilesSettings", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ElasticsearchPercentilesSettings` objects.
func (resource ElasticsearchPercentilesSettings) Equals(other ElasticsearchPercentilesSettings) bool {
	if resource.Script == nil && other.Script != nil || resource.Script != nil && other.Script == nil {
		return false
	}

	if resource.Script != nil {
		if !resource.Script.Equals(*other.Script) {
			return false
		}
	}
	if resource.Missing == nil && other.Missing != nil || resource.Missing != nil && other.Missing == nil {
		return false
	}

	if resource.Missing != nil {
		if *resource.Missing != *other.Missing {
			return false
		}
	}

	if len(resource.Percents) != len(other.Percents) {
		return false
	}

	for i1 := range resource.Percents {
		if resource.Percents[i1] != other.Percents[i1] {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `ElasticsearchPercentilesSettings` fields for violations and returns them.
func (resource ElasticsearchPercentilesSettings) Validate() error {
	var errs cog.BuildErrors
	if resource.Script != nil {
		if err := resource.Script.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("script", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type ElasticsearchUniqueCountSettings struct {
	PrecisionThreshold *string `json:"precision_threshold,omitempty"`
	Missing            *string `json:"missing,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ElasticsearchUniqueCountSettings` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ElasticsearchUniqueCountSettings) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "precision_threshold"
	if fields["precision_threshold"] != nil {
		if string(fields["precision_threshold"]) != "null" {
			if err := json.Unmarshal(fields["precision_threshold"], &resource.PrecisionThreshold); err != nil {
				errs = append(errs, cog.MakeBuildErrors("precision_threshold", err)...)
			}

		}
		delete(fields, "precision_threshold")

	}
	// Field "missing"
	if fields["missing"] != nil {
		if string(fields["missing"]) != "null" {
			if err := json.Unmarshal(fields["missing"], &resource.Missing); err != nil {
				errs = append(errs, cog.MakeBuildErrors("missing", err)...)
			}

		}
		delete(fields, "missing")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ElasticsearchUniqueCountSettings", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ElasticsearchUniqueCountSettings` objects.
func (resource ElasticsearchUniqueCountSettings) Equals(other ElasticsearchUniqueCountSettings) bool {
	if resource.PrecisionThreshold == nil && other.PrecisionThreshold != nil || resource.PrecisionThreshold != nil && other.PrecisionThreshold == nil {
		return false
	}

	if resource.PrecisionThreshold != nil {
		if *resource.PrecisionThreshold != *other.PrecisionThreshold {
			return false
		}
	}
	if resource.Missing == nil && other.Missing != nil || resource.Missing != nil && other.Missing == nil {
		return false
	}

	if resource.Missing != nil {
		if *resource.Missing != *other.Missing {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `ElasticsearchUniqueCountSettings` fields for violations and returns them.
func (resource ElasticsearchUniqueCountSettings) Validate() error {
	return nil
}

type ElasticsearchRawDocumentSettings struct {
	Size *string `json:"size,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ElasticsearchRawDocumentSettings` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ElasticsearchRawDocumentSettings) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "size"
	if fields["size"] != nil {
		if string(fields["size"]) != "null" {
			if err := json.Unmarshal(fields["size"], &resource.Size); err != nil {
				errs = append(errs, cog.MakeBuildErrors("size", err)...)
			}

		}
		delete(fields, "size")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ElasticsearchRawDocumentSettings", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ElasticsearchRawDocumentSettings` objects.
func (resource ElasticsearchRawDocumentSettings) Equals(other ElasticsearchRawDocumentSettings) bool {
	if resource.Size == nil && other.Size != nil || resource.Size != nil && other.Size == nil {
		return false
	}

	if resource.Size != nil {
		if *resource.Size != *other.Size {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `ElasticsearchRawDocumentSettings` fields for violations and returns them.
func (resource ElasticsearchRawDocumentSettings) Validate() error {
	return nil
}

type ElasticsearchRawDataSettings struct {
	Size *string `json:"size,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ElasticsearchRawDataSettings` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ElasticsearchRawDataSettings) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "size"
	if fields["size"] != nil {
		if string(fields["size"]) != "null" {
			if err := json.Unmarshal(fields["size"], &resource.Size); err != nil {
				errs = append(errs, cog.MakeBuildErrors("size", err)...)
			}

		}
		delete(fields, "size")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ElasticsearchRawDataSettings", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ElasticsearchRawDataSettings` objects.
func (resource ElasticsearchRawDataSettings) Equals(other ElasticsearchRawDataSettings) bool {
	if resource.Size == nil && other.Size != nil || resource.Size != nil && other.Size == nil {
		return false
	}

	if resource.Size != nil {
		if *resource.Size != *other.Size {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `ElasticsearchRawDataSettings` fields for violations and returns them.
func (resource ElasticsearchRawDataSettings) Validate() error {
	return nil
}

type ElasticsearchLogsSettings struct {
	Limit *string `json:"limit,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ElasticsearchLogsSettings` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ElasticsearchLogsSettings) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "limit"
	if fields["limit"] != nil {
		if string(fields["limit"]) != "null" {
			if err := json.Unmarshal(fields["limit"], &resource.Limit); err != nil {
				errs = append(errs, cog.MakeBuildErrors("limit", err)...)
			}

		}
		delete(fields, "limit")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ElasticsearchLogsSettings", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ElasticsearchLogsSettings` objects.
func (resource ElasticsearchLogsSettings) Equals(other ElasticsearchLogsSettings) bool {
	if resource.Limit == nil && other.Limit != nil || resource.Limit != nil && other.Limit == nil {
		return false
	}

	if resource.Limit != nil {
		if *resource.Limit != *other.Limit {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `ElasticsearchLogsSettings` fields for violations and returns them.
func (resource ElasticsearchLogsSettings) Validate() error {
	return nil
}

type ElasticsearchRateSettings struct {
	Unit *string `json:"unit,omitempty"`
	Mode *string `json:"mode,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ElasticsearchRateSettings` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ElasticsearchRateSettings) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "unit"
	if fields["unit"] != nil {
		if string(fields["unit"]) != "null" {
			if err := json.Unmarshal(fields["unit"], &resource.Unit); err != nil {
				errs = append(errs, cog.MakeBuildErrors("unit", err)...)
			}

		}
		delete(fields, "unit")

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

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ElasticsearchRateSettings", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ElasticsearchRateSettings` objects.
func (resource ElasticsearchRateSettings) Equals(other ElasticsearchRateSettings) bool {
	if resource.Unit == nil && other.Unit != nil || resource.Unit != nil && other.Unit == nil {
		return false
	}

	if resource.Unit != nil {
		if *resource.Unit != *other.Unit {
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

	return true
}

// Validate checks all the validation constraints that may be defined on `ElasticsearchRateSettings` fields for violations and returns them.
func (resource ElasticsearchRateSettings) Validate() error {
	return nil
}

type ElasticsearchMovingAverageEWMAModelSettingsSettings struct {
	Alpha *string `json:"alpha,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ElasticsearchMovingAverageEWMAModelSettingsSettings` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ElasticsearchMovingAverageEWMAModelSettingsSettings) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "alpha"
	if fields["alpha"] != nil {
		if string(fields["alpha"]) != "null" {
			if err := json.Unmarshal(fields["alpha"], &resource.Alpha); err != nil {
				errs = append(errs, cog.MakeBuildErrors("alpha", err)...)
			}

		}
		delete(fields, "alpha")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ElasticsearchMovingAverageEWMAModelSettingsSettings", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ElasticsearchMovingAverageEWMAModelSettingsSettings` objects.
func (resource ElasticsearchMovingAverageEWMAModelSettingsSettings) Equals(other ElasticsearchMovingAverageEWMAModelSettingsSettings) bool {
	if resource.Alpha == nil && other.Alpha != nil || resource.Alpha != nil && other.Alpha == nil {
		return false
	}

	if resource.Alpha != nil {
		if *resource.Alpha != *other.Alpha {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `ElasticsearchMovingAverageEWMAModelSettingsSettings` fields for violations and returns them.
func (resource ElasticsearchMovingAverageEWMAModelSettingsSettings) Validate() error {
	return nil
}

type ElasticsearchMovingAverageHoltModelSettingsSettings struct {
	Alpha *string `json:"alpha,omitempty"`
	Beta  *string `json:"beta,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ElasticsearchMovingAverageHoltModelSettingsSettings` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ElasticsearchMovingAverageHoltModelSettingsSettings) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "alpha"
	if fields["alpha"] != nil {
		if string(fields["alpha"]) != "null" {
			if err := json.Unmarshal(fields["alpha"], &resource.Alpha); err != nil {
				errs = append(errs, cog.MakeBuildErrors("alpha", err)...)
			}

		}
		delete(fields, "alpha")

	}
	// Field "beta"
	if fields["beta"] != nil {
		if string(fields["beta"]) != "null" {
			if err := json.Unmarshal(fields["beta"], &resource.Beta); err != nil {
				errs = append(errs, cog.MakeBuildErrors("beta", err)...)
			}

		}
		delete(fields, "beta")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ElasticsearchMovingAverageHoltModelSettingsSettings", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ElasticsearchMovingAverageHoltModelSettingsSettings` objects.
func (resource ElasticsearchMovingAverageHoltModelSettingsSettings) Equals(other ElasticsearchMovingAverageHoltModelSettingsSettings) bool {
	if resource.Alpha == nil && other.Alpha != nil || resource.Alpha != nil && other.Alpha == nil {
		return false
	}

	if resource.Alpha != nil {
		if *resource.Alpha != *other.Alpha {
			return false
		}
	}
	if resource.Beta == nil && other.Beta != nil || resource.Beta != nil && other.Beta == nil {
		return false
	}

	if resource.Beta != nil {
		if *resource.Beta != *other.Beta {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `ElasticsearchMovingAverageHoltModelSettingsSettings` fields for violations and returns them.
func (resource ElasticsearchMovingAverageHoltModelSettingsSettings) Validate() error {
	return nil
}

type ElasticsearchMovingAverageHoltWintersModelSettingsSettings struct {
	Alpha  *string `json:"alpha,omitempty"`
	Beta   *string `json:"beta,omitempty"`
	Gamma  *string `json:"gamma,omitempty"`
	Period *string `json:"period,omitempty"`
	Pad    *bool   `json:"pad,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ElasticsearchMovingAverageHoltWintersModelSettingsSettings` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ElasticsearchMovingAverageHoltWintersModelSettingsSettings) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "alpha"
	if fields["alpha"] != nil {
		if string(fields["alpha"]) != "null" {
			if err := json.Unmarshal(fields["alpha"], &resource.Alpha); err != nil {
				errs = append(errs, cog.MakeBuildErrors("alpha", err)...)
			}

		}
		delete(fields, "alpha")

	}
	// Field "beta"
	if fields["beta"] != nil {
		if string(fields["beta"]) != "null" {
			if err := json.Unmarshal(fields["beta"], &resource.Beta); err != nil {
				errs = append(errs, cog.MakeBuildErrors("beta", err)...)
			}

		}
		delete(fields, "beta")

	}
	// Field "gamma"
	if fields["gamma"] != nil {
		if string(fields["gamma"]) != "null" {
			if err := json.Unmarshal(fields["gamma"], &resource.Gamma); err != nil {
				errs = append(errs, cog.MakeBuildErrors("gamma", err)...)
			}

		}
		delete(fields, "gamma")

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
	// Field "pad"
	if fields["pad"] != nil {
		if string(fields["pad"]) != "null" {
			if err := json.Unmarshal(fields["pad"], &resource.Pad); err != nil {
				errs = append(errs, cog.MakeBuildErrors("pad", err)...)
			}

		}
		delete(fields, "pad")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ElasticsearchMovingAverageHoltWintersModelSettingsSettings", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ElasticsearchMovingAverageHoltWintersModelSettingsSettings` objects.
func (resource ElasticsearchMovingAverageHoltWintersModelSettingsSettings) Equals(other ElasticsearchMovingAverageHoltWintersModelSettingsSettings) bool {
	if resource.Alpha == nil && other.Alpha != nil || resource.Alpha != nil && other.Alpha == nil {
		return false
	}

	if resource.Alpha != nil {
		if *resource.Alpha != *other.Alpha {
			return false
		}
	}
	if resource.Beta == nil && other.Beta != nil || resource.Beta != nil && other.Beta == nil {
		return false
	}

	if resource.Beta != nil {
		if *resource.Beta != *other.Beta {
			return false
		}
	}
	if resource.Gamma == nil && other.Gamma != nil || resource.Gamma != nil && other.Gamma == nil {
		return false
	}

	if resource.Gamma != nil {
		if *resource.Gamma != *other.Gamma {
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
	if resource.Pad == nil && other.Pad != nil || resource.Pad != nil && other.Pad == nil {
		return false
	}

	if resource.Pad != nil {
		if *resource.Pad != *other.Pad {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `ElasticsearchMovingAverageHoltWintersModelSettingsSettings` fields for violations and returns them.
func (resource ElasticsearchMovingAverageHoltWintersModelSettingsSettings) Validate() error {
	return nil
}

type ElasticsearchMovingFunctionSettings struct {
	Window *string       `json:"window,omitempty"`
	Script *InlineScript `json:"script,omitempty"`
	Shift  *string       `json:"shift,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ElasticsearchMovingFunctionSettings` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ElasticsearchMovingFunctionSettings) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "window"
	if fields["window"] != nil {
		if string(fields["window"]) != "null" {
			if err := json.Unmarshal(fields["window"], &resource.Window); err != nil {
				errs = append(errs, cog.MakeBuildErrors("window", err)...)
			}

		}
		delete(fields, "window")

	}
	// Field "script"
	if fields["script"] != nil {
		if string(fields["script"]) != "null" {

			resource.Script = &InlineScript{}
			if err := resource.Script.UnmarshalJSONStrict(fields["script"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("script", err)...)
			}

		}
		delete(fields, "script")

	}
	// Field "shift"
	if fields["shift"] != nil {
		if string(fields["shift"]) != "null" {
			if err := json.Unmarshal(fields["shift"], &resource.Shift); err != nil {
				errs = append(errs, cog.MakeBuildErrors("shift", err)...)
			}

		}
		delete(fields, "shift")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ElasticsearchMovingFunctionSettings", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ElasticsearchMovingFunctionSettings` objects.
func (resource ElasticsearchMovingFunctionSettings) Equals(other ElasticsearchMovingFunctionSettings) bool {
	if resource.Window == nil && other.Window != nil || resource.Window != nil && other.Window == nil {
		return false
	}

	if resource.Window != nil {
		if *resource.Window != *other.Window {
			return false
		}
	}
	if resource.Script == nil && other.Script != nil || resource.Script != nil && other.Script == nil {
		return false
	}

	if resource.Script != nil {
		if !resource.Script.Equals(*other.Script) {
			return false
		}
	}
	if resource.Shift == nil && other.Shift != nil || resource.Shift != nil && other.Shift == nil {
		return false
	}

	if resource.Shift != nil {
		if *resource.Shift != *other.Shift {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `ElasticsearchMovingFunctionSettings` fields for violations and returns them.
func (resource ElasticsearchMovingFunctionSettings) Validate() error {
	var errs cog.BuildErrors
	if resource.Script != nil {
		if err := resource.Script.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("script", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type ElasticsearchDerivativeSettings struct {
	Unit *string `json:"unit,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ElasticsearchDerivativeSettings` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ElasticsearchDerivativeSettings) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "unit"
	if fields["unit"] != nil {
		if string(fields["unit"]) != "null" {
			if err := json.Unmarshal(fields["unit"], &resource.Unit); err != nil {
				errs = append(errs, cog.MakeBuildErrors("unit", err)...)
			}

		}
		delete(fields, "unit")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ElasticsearchDerivativeSettings", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ElasticsearchDerivativeSettings` objects.
func (resource ElasticsearchDerivativeSettings) Equals(other ElasticsearchDerivativeSettings) bool {
	if resource.Unit == nil && other.Unit != nil || resource.Unit != nil && other.Unit == nil {
		return false
	}

	if resource.Unit != nil {
		if *resource.Unit != *other.Unit {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `ElasticsearchDerivativeSettings` fields for violations and returns them.
func (resource ElasticsearchDerivativeSettings) Validate() error {
	return nil
}

type ElasticsearchSerialDiffSettings struct {
	Lag *string `json:"lag,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ElasticsearchSerialDiffSettings` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ElasticsearchSerialDiffSettings) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "lag"
	if fields["lag"] != nil {
		if string(fields["lag"]) != "null" {
			if err := json.Unmarshal(fields["lag"], &resource.Lag); err != nil {
				errs = append(errs, cog.MakeBuildErrors("lag", err)...)
			}

		}
		delete(fields, "lag")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ElasticsearchSerialDiffSettings", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ElasticsearchSerialDiffSettings` objects.
func (resource ElasticsearchSerialDiffSettings) Equals(other ElasticsearchSerialDiffSettings) bool {
	if resource.Lag == nil && other.Lag != nil || resource.Lag != nil && other.Lag == nil {
		return false
	}

	if resource.Lag != nil {
		if *resource.Lag != *other.Lag {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `ElasticsearchSerialDiffSettings` fields for violations and returns them.
func (resource ElasticsearchSerialDiffSettings) Validate() error {
	return nil
}

type ElasticsearchCumulativeSumSettings struct {
	Format *string `json:"format,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ElasticsearchCumulativeSumSettings` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ElasticsearchCumulativeSumSettings) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "format"
	if fields["format"] != nil {
		if string(fields["format"]) != "null" {
			if err := json.Unmarshal(fields["format"], &resource.Format); err != nil {
				errs = append(errs, cog.MakeBuildErrors("format", err)...)
			}

		}
		delete(fields, "format")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ElasticsearchCumulativeSumSettings", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ElasticsearchCumulativeSumSettings` objects.
func (resource ElasticsearchCumulativeSumSettings) Equals(other ElasticsearchCumulativeSumSettings) bool {
	if resource.Format == nil && other.Format != nil || resource.Format != nil && other.Format == nil {
		return false
	}

	if resource.Format != nil {
		if *resource.Format != *other.Format {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `ElasticsearchCumulativeSumSettings` fields for violations and returns them.
func (resource ElasticsearchCumulativeSumSettings) Validate() error {
	return nil
}

type ElasticsearchBucketScriptSettings struct {
	Script *InlineScript `json:"script,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ElasticsearchBucketScriptSettings` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ElasticsearchBucketScriptSettings) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "script"
	if fields["script"] != nil {
		if string(fields["script"]) != "null" {

			resource.Script = &InlineScript{}
			if err := resource.Script.UnmarshalJSONStrict(fields["script"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("script", err)...)
			}

		}
		delete(fields, "script")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ElasticsearchBucketScriptSettings", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ElasticsearchBucketScriptSettings` objects.
func (resource ElasticsearchBucketScriptSettings) Equals(other ElasticsearchBucketScriptSettings) bool {
	if resource.Script == nil && other.Script != nil || resource.Script != nil && other.Script == nil {
		return false
	}

	if resource.Script != nil {
		if !resource.Script.Equals(*other.Script) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `ElasticsearchBucketScriptSettings` fields for violations and returns them.
func (resource ElasticsearchBucketScriptSettings) Validate() error {
	var errs cog.BuildErrors
	if resource.Script != nil {
		if err := resource.Script.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("script", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type ElasticsearchTopMetricsSettings struct {
	Order   *string  `json:"order,omitempty"`
	OrderBy *string  `json:"orderBy,omitempty"`
	Metrics []string `json:"metrics,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ElasticsearchTopMetricsSettings` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ElasticsearchTopMetricsSettings) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "order"
	if fields["order"] != nil {
		if string(fields["order"]) != "null" {
			if err := json.Unmarshal(fields["order"], &resource.Order); err != nil {
				errs = append(errs, cog.MakeBuildErrors("order", err)...)
			}

		}
		delete(fields, "order")

	}
	// Field "orderBy"
	if fields["orderBy"] != nil {
		if string(fields["orderBy"]) != "null" {
			if err := json.Unmarshal(fields["orderBy"], &resource.OrderBy); err != nil {
				errs = append(errs, cog.MakeBuildErrors("orderBy", err)...)
			}

		}
		delete(fields, "orderBy")

	}
	// Field "metrics"
	if fields["metrics"] != nil {
		if string(fields["metrics"]) != "null" {

			if err := json.Unmarshal(fields["metrics"], &resource.Metrics); err != nil {
				errs = append(errs, cog.MakeBuildErrors("metrics", err)...)
			}

		}
		delete(fields, "metrics")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ElasticsearchTopMetricsSettings", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ElasticsearchTopMetricsSettings` objects.
func (resource ElasticsearchTopMetricsSettings) Equals(other ElasticsearchTopMetricsSettings) bool {
	if resource.Order == nil && other.Order != nil || resource.Order != nil && other.Order == nil {
		return false
	}

	if resource.Order != nil {
		if *resource.Order != *other.Order {
			return false
		}
	}
	if resource.OrderBy == nil && other.OrderBy != nil || resource.OrderBy != nil && other.OrderBy == nil {
		return false
	}

	if resource.OrderBy != nil {
		if *resource.OrderBy != *other.OrderBy {
			return false
		}
	}

	if len(resource.Metrics) != len(other.Metrics) {
		return false
	}

	for i1 := range resource.Metrics {
		if resource.Metrics[i1] != other.Metrics[i1] {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `ElasticsearchTopMetricsSettings` fields for violations and returns them.
func (resource ElasticsearchTopMetricsSettings) Validate() error {
	return nil
}

type DateHistogramOrHistogramOrTermsOrFiltersOrGeoHashGridOrNested struct {
	DateHistogram *DateHistogram `json:"DateHistogram,omitempty"`
	Histogram     *Histogram     `json:"Histogram,omitempty"`
	Terms         *Terms         `json:"Terms,omitempty"`
	Filters       *Filters       `json:"Filters,omitempty"`
	GeoHashGrid   *GeoHashGrid   `json:"GeoHashGrid,omitempty"`
	Nested        *Nested        `json:"Nested,omitempty"`
}

// MarshalJSON implements a custom JSON marshalling logic to encode `DateHistogramOrHistogramOrTermsOrFiltersOrGeoHashGridOrNested` as JSON.
func (resource DateHistogramOrHistogramOrTermsOrFiltersOrGeoHashGridOrNested) MarshalJSON() ([]byte, error) {
	if resource.DateHistogram != nil {
		return json.Marshal(resource.DateHistogram)
	}
	if resource.Histogram != nil {
		return json.Marshal(resource.Histogram)
	}
	if resource.Terms != nil {
		return json.Marshal(resource.Terms)
	}
	if resource.Filters != nil {
		return json.Marshal(resource.Filters)
	}
	if resource.GeoHashGrid != nil {
		return json.Marshal(resource.GeoHashGrid)
	}
	if resource.Nested != nil {
		return json.Marshal(resource.Nested)
	}

	return nil, fmt.Errorf("no value for disjunction of refs")
}

// UnmarshalJSON implements a custom JSON unmarshalling logic to decode `DateHistogramOrHistogramOrTermsOrFiltersOrGeoHashGridOrNested` from JSON.
func (resource *DateHistogramOrHistogramOrTermsOrFiltersOrGeoHashGridOrNested) UnmarshalJSON(raw []byte) error {
	if raw == nil {
		return nil
	}

	// FIXME: this is wasteful, we need to find a more efficient way to unmarshal this.
	parsedAsMap := make(map[string]any)
	if err := json.Unmarshal(raw, &parsedAsMap); err != nil {
		return err
	}

	discriminator, found := parsedAsMap["type"]
	if !found {
		return errors.New("discriminator field 'type' not found in payload")
	}

	switch discriminator {
	case "date_histogram":
		var dateHistogram DateHistogram
		if err := json.Unmarshal(raw, &dateHistogram); err != nil {
			return err
		}

		resource.DateHistogram = &dateHistogram
		return nil
	case "filters":
		var filters Filters
		if err := json.Unmarshal(raw, &filters); err != nil {
			return err
		}

		resource.Filters = &filters
		return nil
	case "geohash_grid":
		var geoHashGrid GeoHashGrid
		if err := json.Unmarshal(raw, &geoHashGrid); err != nil {
			return err
		}

		resource.GeoHashGrid = &geoHashGrid
		return nil
	case "histogram":
		var histogram Histogram
		if err := json.Unmarshal(raw, &histogram); err != nil {
			return err
		}

		resource.Histogram = &histogram
		return nil
	case "nested":
		var nested Nested
		if err := json.Unmarshal(raw, &nested); err != nil {
			return err
		}

		resource.Nested = &nested
		return nil
	case "terms":
		var terms Terms
		if err := json.Unmarshal(raw, &terms); err != nil {
			return err
		}

		resource.Terms = &terms
		return nil
	}

	return fmt.Errorf("could not unmarshal resource with `type = %v`", discriminator)
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `DateHistogramOrHistogramOrTermsOrFiltersOrGeoHashGridOrNested` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *DateHistogramOrHistogramOrTermsOrFiltersOrGeoHashGridOrNested) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	// FIXME: this is wasteful, we need to find a more efficient way to unmarshal this.
	parsedAsMap := make(map[string]any)
	if err := json.Unmarshal(raw, &parsedAsMap); err != nil {
		return err
	}

	discriminator, found := parsedAsMap["type"]
	if !found {
		return fmt.Errorf("discriminator field 'type' not found in payload")
	}

	switch discriminator {
	case "date_histogram":
		dateHistogram := &DateHistogram{}
		if err := dateHistogram.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.DateHistogram = dateHistogram
		return nil
	case "filters":
		filters := &Filters{}
		if err := filters.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.Filters = filters
		return nil
	case "geohash_grid":
		geoHashGrid := &GeoHashGrid{}
		if err := geoHashGrid.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.GeoHashGrid = geoHashGrid
		return nil
	case "histogram":
		histogram := &Histogram{}
		if err := histogram.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.Histogram = histogram
		return nil
	case "nested":
		nested := &Nested{}
		if err := nested.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.Nested = nested
		return nil
	case "terms":
		terms := &Terms{}
		if err := terms.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.Terms = terms
		return nil
	}

	return fmt.Errorf("could not unmarshal resource with `type = %v`", discriminator)
}

// Equals tests the equality of two `DateHistogramOrHistogramOrTermsOrFiltersOrGeoHashGridOrNested` objects.
func (resource DateHistogramOrHistogramOrTermsOrFiltersOrGeoHashGridOrNested) Equals(other DateHistogramOrHistogramOrTermsOrFiltersOrGeoHashGridOrNested) bool {
	if resource.DateHistogram == nil && other.DateHistogram != nil || resource.DateHistogram != nil && other.DateHistogram == nil {
		return false
	}

	if resource.DateHistogram != nil {
		if !resource.DateHistogram.Equals(*other.DateHistogram) {
			return false
		}
	}
	if resource.Histogram == nil && other.Histogram != nil || resource.Histogram != nil && other.Histogram == nil {
		return false
	}

	if resource.Histogram != nil {
		if !resource.Histogram.Equals(*other.Histogram) {
			return false
		}
	}
	if resource.Terms == nil && other.Terms != nil || resource.Terms != nil && other.Terms == nil {
		return false
	}

	if resource.Terms != nil {
		if !resource.Terms.Equals(*other.Terms) {
			return false
		}
	}
	if resource.Filters == nil && other.Filters != nil || resource.Filters != nil && other.Filters == nil {
		return false
	}

	if resource.Filters != nil {
		if !resource.Filters.Equals(*other.Filters) {
			return false
		}
	}
	if resource.GeoHashGrid == nil && other.GeoHashGrid != nil || resource.GeoHashGrid != nil && other.GeoHashGrid == nil {
		return false
	}

	if resource.GeoHashGrid != nil {
		if !resource.GeoHashGrid.Equals(*other.GeoHashGrid) {
			return false
		}
	}
	if resource.Nested == nil && other.Nested != nil || resource.Nested != nil && other.Nested == nil {
		return false
	}

	if resource.Nested != nil {
		if !resource.Nested.Equals(*other.Nested) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `DateHistogramOrHistogramOrTermsOrFiltersOrGeoHashGridOrNested` fields for violations and returns them.
func (resource DateHistogramOrHistogramOrTermsOrFiltersOrGeoHashGridOrNested) Validate() error {
	var errs cog.BuildErrors
	if resource.DateHistogram != nil {
		if err := resource.DateHistogram.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("DateHistogram", err)...)
		}
	}
	if resource.Histogram != nil {
		if err := resource.Histogram.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("Histogram", err)...)
		}
	}
	if resource.Terms != nil {
		if err := resource.Terms.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("Terms", err)...)
		}
	}
	if resource.Filters != nil {
		if err := resource.Filters.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("Filters", err)...)
		}
	}
	if resource.GeoHashGrid != nil {
		if err := resource.GeoHashGrid.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("GeoHashGrid", err)...)
		}
	}
	if resource.Nested != nil {
		if err := resource.Nested.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("Nested", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type CountOrMovingAverageOrDerivativeOrCumulativeSumOrBucketScriptOrSerialDiffOrRawDataOrRawDocumentOrUniqueCountOrPercentilesOrExtendedStatsOrMinOrMaxOrSumOrAverageOrMovingFunctionOrLogsOrRateOrTopMetrics struct {
	Count          *Count          `json:"Count,omitempty"`
	MovingAverage  *MovingAverage  `json:"MovingAverage,omitempty"`
	Derivative     *Derivative     `json:"Derivative,omitempty"`
	CumulativeSum  *CumulativeSum  `json:"CumulativeSum,omitempty"`
	BucketScript   *BucketScript   `json:"BucketScript,omitempty"`
	SerialDiff     *SerialDiff     `json:"SerialDiff,omitempty"`
	RawData        *RawData        `json:"RawData,omitempty"`
	RawDocument    *RawDocument    `json:"RawDocument,omitempty"`
	UniqueCount    *UniqueCount    `json:"UniqueCount,omitempty"`
	Percentiles    *Percentiles    `json:"Percentiles,omitempty"`
	ExtendedStats  *ExtendedStats  `json:"ExtendedStats,omitempty"`
	Min            *Min            `json:"Min,omitempty"`
	Max            *Max            `json:"Max,omitempty"`
	Sum            *Sum            `json:"Sum,omitempty"`
	Average        *Average        `json:"Average,omitempty"`
	MovingFunction *MovingFunction `json:"MovingFunction,omitempty"`
	Logs           *Logs           `json:"Logs,omitempty"`
	Rate           *Rate           `json:"Rate,omitempty"`
	TopMetrics     *TopMetrics     `json:"TopMetrics,omitempty"`
}

// MarshalJSON implements a custom JSON marshalling logic to encode `CountOrMovingAverageOrDerivativeOrCumulativeSumOrBucketScriptOrSerialDiffOrRawDataOrRawDocumentOrUniqueCountOrPercentilesOrExtendedStatsOrMinOrMaxOrSumOrAverageOrMovingFunctionOrLogsOrRateOrTopMetrics` as JSON.
func (resource CountOrMovingAverageOrDerivativeOrCumulativeSumOrBucketScriptOrSerialDiffOrRawDataOrRawDocumentOrUniqueCountOrPercentilesOrExtendedStatsOrMinOrMaxOrSumOrAverageOrMovingFunctionOrLogsOrRateOrTopMetrics) MarshalJSON() ([]byte, error) {
	if resource.Count != nil {
		return json.Marshal(resource.Count)
	}
	if resource.MovingAverage != nil {
		return json.Marshal(resource.MovingAverage)
	}
	if resource.Derivative != nil {
		return json.Marshal(resource.Derivative)
	}
	if resource.CumulativeSum != nil {
		return json.Marshal(resource.CumulativeSum)
	}
	if resource.BucketScript != nil {
		return json.Marshal(resource.BucketScript)
	}
	if resource.SerialDiff != nil {
		return json.Marshal(resource.SerialDiff)
	}
	if resource.RawData != nil {
		return json.Marshal(resource.RawData)
	}
	if resource.RawDocument != nil {
		return json.Marshal(resource.RawDocument)
	}
	if resource.UniqueCount != nil {
		return json.Marshal(resource.UniqueCount)
	}
	if resource.Percentiles != nil {
		return json.Marshal(resource.Percentiles)
	}
	if resource.ExtendedStats != nil {
		return json.Marshal(resource.ExtendedStats)
	}
	if resource.Min != nil {
		return json.Marshal(resource.Min)
	}
	if resource.Max != nil {
		return json.Marshal(resource.Max)
	}
	if resource.Sum != nil {
		return json.Marshal(resource.Sum)
	}
	if resource.Average != nil {
		return json.Marshal(resource.Average)
	}
	if resource.MovingFunction != nil {
		return json.Marshal(resource.MovingFunction)
	}
	if resource.Logs != nil {
		return json.Marshal(resource.Logs)
	}
	if resource.Rate != nil {
		return json.Marshal(resource.Rate)
	}
	if resource.TopMetrics != nil {
		return json.Marshal(resource.TopMetrics)
	}

	return nil, fmt.Errorf("no value for disjunction of refs")
}

// UnmarshalJSON implements a custom JSON unmarshalling logic to decode `CountOrMovingAverageOrDerivativeOrCumulativeSumOrBucketScriptOrSerialDiffOrRawDataOrRawDocumentOrUniqueCountOrPercentilesOrExtendedStatsOrMinOrMaxOrSumOrAverageOrMovingFunctionOrLogsOrRateOrTopMetrics` from JSON.
func (resource *CountOrMovingAverageOrDerivativeOrCumulativeSumOrBucketScriptOrSerialDiffOrRawDataOrRawDocumentOrUniqueCountOrPercentilesOrExtendedStatsOrMinOrMaxOrSumOrAverageOrMovingFunctionOrLogsOrRateOrTopMetrics) UnmarshalJSON(raw []byte) error {
	if raw == nil {
		return nil
	}

	// FIXME: this is wasteful, we need to find a more efficient way to unmarshal this.
	parsedAsMap := make(map[string]any)
	if err := json.Unmarshal(raw, &parsedAsMap); err != nil {
		return err
	}

	discriminator, found := parsedAsMap["type"]
	if !found {
		return errors.New("discriminator field 'type' not found in payload")
	}

	switch discriminator {
	case "avg":
		var average Average
		if err := json.Unmarshal(raw, &average); err != nil {
			return err
		}

		resource.Average = &average
		return nil
	case "bucket_script":
		var bucketScript BucketScript
		if err := json.Unmarshal(raw, &bucketScript); err != nil {
			return err
		}

		resource.BucketScript = &bucketScript
		return nil
	case "cardinality":
		var uniqueCount UniqueCount
		if err := json.Unmarshal(raw, &uniqueCount); err != nil {
			return err
		}

		resource.UniqueCount = &uniqueCount
		return nil
	case "count":
		var count Count
		if err := json.Unmarshal(raw, &count); err != nil {
			return err
		}

		resource.Count = &count
		return nil
	case "cumulative_sum":
		var cumulativeSum CumulativeSum
		if err := json.Unmarshal(raw, &cumulativeSum); err != nil {
			return err
		}

		resource.CumulativeSum = &cumulativeSum
		return nil
	case "derivative":
		var derivative Derivative
		if err := json.Unmarshal(raw, &derivative); err != nil {
			return err
		}

		resource.Derivative = &derivative
		return nil
	case "extended_stats":
		var extendedStats ExtendedStats
		if err := json.Unmarshal(raw, &extendedStats); err != nil {
			return err
		}

		resource.ExtendedStats = &extendedStats
		return nil
	case "logs":
		var logs Logs
		if err := json.Unmarshal(raw, &logs); err != nil {
			return err
		}

		resource.Logs = &logs
		return nil
	case "max":
		var max Max
		if err := json.Unmarshal(raw, &max); err != nil {
			return err
		}

		resource.Max = &max
		return nil
	case "min":
		var min Min
		if err := json.Unmarshal(raw, &min); err != nil {
			return err
		}

		resource.Min = &min
		return nil
	case "moving_avg":
		var movingAverage MovingAverage
		if err := json.Unmarshal(raw, &movingAverage); err != nil {
			return err
		}

		resource.MovingAverage = &movingAverage
		return nil
	case "moving_fn":
		var movingFunction MovingFunction
		if err := json.Unmarshal(raw, &movingFunction); err != nil {
			return err
		}

		resource.MovingFunction = &movingFunction
		return nil
	case "percentiles":
		var percentiles Percentiles
		if err := json.Unmarshal(raw, &percentiles); err != nil {
			return err
		}

		resource.Percentiles = &percentiles
		return nil
	case "rate":
		var rate Rate
		if err := json.Unmarshal(raw, &rate); err != nil {
			return err
		}

		resource.Rate = &rate
		return nil
	case "raw_data":
		var rawData RawData
		if err := json.Unmarshal(raw, &rawData); err != nil {
			return err
		}

		resource.RawData = &rawData
		return nil
	case "raw_document":
		var rawDocument RawDocument
		if err := json.Unmarshal(raw, &rawDocument); err != nil {
			return err
		}

		resource.RawDocument = &rawDocument
		return nil
	case "serial_diff":
		var serialDiff SerialDiff
		if err := json.Unmarshal(raw, &serialDiff); err != nil {
			return err
		}

		resource.SerialDiff = &serialDiff
		return nil
	case "sum":
		var sum Sum
		if err := json.Unmarshal(raw, &sum); err != nil {
			return err
		}

		resource.Sum = &sum
		return nil
	case "top_metrics":
		var topMetrics TopMetrics
		if err := json.Unmarshal(raw, &topMetrics); err != nil {
			return err
		}

		resource.TopMetrics = &topMetrics
		return nil
	}

	return fmt.Errorf("could not unmarshal resource with `type = %v`", discriminator)
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `CountOrMovingAverageOrDerivativeOrCumulativeSumOrBucketScriptOrSerialDiffOrRawDataOrRawDocumentOrUniqueCountOrPercentilesOrExtendedStatsOrMinOrMaxOrSumOrAverageOrMovingFunctionOrLogsOrRateOrTopMetrics` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *CountOrMovingAverageOrDerivativeOrCumulativeSumOrBucketScriptOrSerialDiffOrRawDataOrRawDocumentOrUniqueCountOrPercentilesOrExtendedStatsOrMinOrMaxOrSumOrAverageOrMovingFunctionOrLogsOrRateOrTopMetrics) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	// FIXME: this is wasteful, we need to find a more efficient way to unmarshal this.
	parsedAsMap := make(map[string]any)
	if err := json.Unmarshal(raw, &parsedAsMap); err != nil {
		return err
	}

	discriminator, found := parsedAsMap["type"]
	if !found {
		return fmt.Errorf("discriminator field 'type' not found in payload")
	}

	switch discriminator {
	case "avg":
		average := &Average{}
		if err := average.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.Average = average
		return nil
	case "bucket_script":
		bucketScript := &BucketScript{}
		if err := bucketScript.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.BucketScript = bucketScript
		return nil
	case "cardinality":
		uniqueCount := &UniqueCount{}
		if err := uniqueCount.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.UniqueCount = uniqueCount
		return nil
	case "count":
		count := &Count{}
		if err := count.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.Count = count
		return nil
	case "cumulative_sum":
		cumulativeSum := &CumulativeSum{}
		if err := cumulativeSum.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.CumulativeSum = cumulativeSum
		return nil
	case "derivative":
		derivative := &Derivative{}
		if err := derivative.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.Derivative = derivative
		return nil
	case "extended_stats":
		extendedStats := &ExtendedStats{}
		if err := extendedStats.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.ExtendedStats = extendedStats
		return nil
	case "logs":
		logs := &Logs{}
		if err := logs.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.Logs = logs
		return nil
	case "max":
		max := &Max{}
		if err := max.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.Max = max
		return nil
	case "min":
		min := &Min{}
		if err := min.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.Min = min
		return nil
	case "moving_avg":
		movingAverage := &MovingAverage{}
		if err := movingAverage.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.MovingAverage = movingAverage
		return nil
	case "moving_fn":
		movingFunction := &MovingFunction{}
		if err := movingFunction.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.MovingFunction = movingFunction
		return nil
	case "percentiles":
		percentiles := &Percentiles{}
		if err := percentiles.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.Percentiles = percentiles
		return nil
	case "rate":
		rate := &Rate{}
		if err := rate.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.Rate = rate
		return nil
	case "raw_data":
		rawData := &RawData{}
		if err := rawData.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.RawData = rawData
		return nil
	case "raw_document":
		rawDocument := &RawDocument{}
		if err := rawDocument.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.RawDocument = rawDocument
		return nil
	case "serial_diff":
		serialDiff := &SerialDiff{}
		if err := serialDiff.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.SerialDiff = serialDiff
		return nil
	case "sum":
		sum := &Sum{}
		if err := sum.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.Sum = sum
		return nil
	case "top_metrics":
		topMetrics := &TopMetrics{}
		if err := topMetrics.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.TopMetrics = topMetrics
		return nil
	}

	return fmt.Errorf("could not unmarshal resource with `type = %v`", discriminator)
}

// Equals tests the equality of two `CountOrMovingAverageOrDerivativeOrCumulativeSumOrBucketScriptOrSerialDiffOrRawDataOrRawDocumentOrUniqueCountOrPercentilesOrExtendedStatsOrMinOrMaxOrSumOrAverageOrMovingFunctionOrLogsOrRateOrTopMetrics` objects.
func (resource CountOrMovingAverageOrDerivativeOrCumulativeSumOrBucketScriptOrSerialDiffOrRawDataOrRawDocumentOrUniqueCountOrPercentilesOrExtendedStatsOrMinOrMaxOrSumOrAverageOrMovingFunctionOrLogsOrRateOrTopMetrics) Equals(other CountOrMovingAverageOrDerivativeOrCumulativeSumOrBucketScriptOrSerialDiffOrRawDataOrRawDocumentOrUniqueCountOrPercentilesOrExtendedStatsOrMinOrMaxOrSumOrAverageOrMovingFunctionOrLogsOrRateOrTopMetrics) bool {
	if resource.Count == nil && other.Count != nil || resource.Count != nil && other.Count == nil {
		return false
	}

	if resource.Count != nil {
		if !resource.Count.Equals(*other.Count) {
			return false
		}
	}
	if resource.MovingAverage == nil && other.MovingAverage != nil || resource.MovingAverage != nil && other.MovingAverage == nil {
		return false
	}

	if resource.MovingAverage != nil {
		if !resource.MovingAverage.Equals(*other.MovingAverage) {
			return false
		}
	}
	if resource.Derivative == nil && other.Derivative != nil || resource.Derivative != nil && other.Derivative == nil {
		return false
	}

	if resource.Derivative != nil {
		if !resource.Derivative.Equals(*other.Derivative) {
			return false
		}
	}
	if resource.CumulativeSum == nil && other.CumulativeSum != nil || resource.CumulativeSum != nil && other.CumulativeSum == nil {
		return false
	}

	if resource.CumulativeSum != nil {
		if !resource.CumulativeSum.Equals(*other.CumulativeSum) {
			return false
		}
	}
	if resource.BucketScript == nil && other.BucketScript != nil || resource.BucketScript != nil && other.BucketScript == nil {
		return false
	}

	if resource.BucketScript != nil {
		if !resource.BucketScript.Equals(*other.BucketScript) {
			return false
		}
	}
	if resource.SerialDiff == nil && other.SerialDiff != nil || resource.SerialDiff != nil && other.SerialDiff == nil {
		return false
	}

	if resource.SerialDiff != nil {
		if !resource.SerialDiff.Equals(*other.SerialDiff) {
			return false
		}
	}
	if resource.RawData == nil && other.RawData != nil || resource.RawData != nil && other.RawData == nil {
		return false
	}

	if resource.RawData != nil {
		if !resource.RawData.Equals(*other.RawData) {
			return false
		}
	}
	if resource.RawDocument == nil && other.RawDocument != nil || resource.RawDocument != nil && other.RawDocument == nil {
		return false
	}

	if resource.RawDocument != nil {
		if !resource.RawDocument.Equals(*other.RawDocument) {
			return false
		}
	}
	if resource.UniqueCount == nil && other.UniqueCount != nil || resource.UniqueCount != nil && other.UniqueCount == nil {
		return false
	}

	if resource.UniqueCount != nil {
		if !resource.UniqueCount.Equals(*other.UniqueCount) {
			return false
		}
	}
	if resource.Percentiles == nil && other.Percentiles != nil || resource.Percentiles != nil && other.Percentiles == nil {
		return false
	}

	if resource.Percentiles != nil {
		if !resource.Percentiles.Equals(*other.Percentiles) {
			return false
		}
	}
	if resource.ExtendedStats == nil && other.ExtendedStats != nil || resource.ExtendedStats != nil && other.ExtendedStats == nil {
		return false
	}

	if resource.ExtendedStats != nil {
		if !resource.ExtendedStats.Equals(*other.ExtendedStats) {
			return false
		}
	}
	if resource.Min == nil && other.Min != nil || resource.Min != nil && other.Min == nil {
		return false
	}

	if resource.Min != nil {
		if !resource.Min.Equals(*other.Min) {
			return false
		}
	}
	if resource.Max == nil && other.Max != nil || resource.Max != nil && other.Max == nil {
		return false
	}

	if resource.Max != nil {
		if !resource.Max.Equals(*other.Max) {
			return false
		}
	}
	if resource.Sum == nil && other.Sum != nil || resource.Sum != nil && other.Sum == nil {
		return false
	}

	if resource.Sum != nil {
		if !resource.Sum.Equals(*other.Sum) {
			return false
		}
	}
	if resource.Average == nil && other.Average != nil || resource.Average != nil && other.Average == nil {
		return false
	}

	if resource.Average != nil {
		if !resource.Average.Equals(*other.Average) {
			return false
		}
	}
	if resource.MovingFunction == nil && other.MovingFunction != nil || resource.MovingFunction != nil && other.MovingFunction == nil {
		return false
	}

	if resource.MovingFunction != nil {
		if !resource.MovingFunction.Equals(*other.MovingFunction) {
			return false
		}
	}
	if resource.Logs == nil && other.Logs != nil || resource.Logs != nil && other.Logs == nil {
		return false
	}

	if resource.Logs != nil {
		if !resource.Logs.Equals(*other.Logs) {
			return false
		}
	}
	if resource.Rate == nil && other.Rate != nil || resource.Rate != nil && other.Rate == nil {
		return false
	}

	if resource.Rate != nil {
		if !resource.Rate.Equals(*other.Rate) {
			return false
		}
	}
	if resource.TopMetrics == nil && other.TopMetrics != nil || resource.TopMetrics != nil && other.TopMetrics == nil {
		return false
	}

	if resource.TopMetrics != nil {
		if !resource.TopMetrics.Equals(*other.TopMetrics) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `CountOrMovingAverageOrDerivativeOrCumulativeSumOrBucketScriptOrSerialDiffOrRawDataOrRawDocumentOrUniqueCountOrPercentilesOrExtendedStatsOrMinOrMaxOrSumOrAverageOrMovingFunctionOrLogsOrRateOrTopMetrics` fields for violations and returns them.
func (resource CountOrMovingAverageOrDerivativeOrCumulativeSumOrBucketScriptOrSerialDiffOrRawDataOrRawDocumentOrUniqueCountOrPercentilesOrExtendedStatsOrMinOrMaxOrSumOrAverageOrMovingFunctionOrLogsOrRateOrTopMetrics) Validate() error {
	var errs cog.BuildErrors
	if resource.Count != nil {
		if err := resource.Count.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("Count", err)...)
		}
	}
	if resource.MovingAverage != nil {
		if err := resource.MovingAverage.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("MovingAverage", err)...)
		}
	}
	if resource.Derivative != nil {
		if err := resource.Derivative.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("Derivative", err)...)
		}
	}
	if resource.CumulativeSum != nil {
		if err := resource.CumulativeSum.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("CumulativeSum", err)...)
		}
	}
	if resource.BucketScript != nil {
		if err := resource.BucketScript.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("BucketScript", err)...)
		}
	}
	if resource.SerialDiff != nil {
		if err := resource.SerialDiff.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("SerialDiff", err)...)
		}
	}
	if resource.RawData != nil {
		if err := resource.RawData.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("RawData", err)...)
		}
	}
	if resource.RawDocument != nil {
		if err := resource.RawDocument.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("RawDocument", err)...)
		}
	}
	if resource.UniqueCount != nil {
		if err := resource.UniqueCount.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("UniqueCount", err)...)
		}
	}
	if resource.Percentiles != nil {
		if err := resource.Percentiles.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("Percentiles", err)...)
		}
	}
	if resource.ExtendedStats != nil {
		if err := resource.ExtendedStats.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("ExtendedStats", err)...)
		}
	}
	if resource.Min != nil {
		if err := resource.Min.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("Min", err)...)
		}
	}
	if resource.Max != nil {
		if err := resource.Max.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("Max", err)...)
		}
	}
	if resource.Sum != nil {
		if err := resource.Sum.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("Sum", err)...)
		}
	}
	if resource.Average != nil {
		if err := resource.Average.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("Average", err)...)
		}
	}
	if resource.MovingFunction != nil {
		if err := resource.MovingFunction.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("MovingFunction", err)...)
		}
	}
	if resource.Logs != nil {
		if err := resource.Logs.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("Logs", err)...)
		}
	}
	if resource.Rate != nil {
		if err := resource.Rate.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("Rate", err)...)
		}
	}
	if resource.TopMetrics != nil {
		if err := resource.TopMetrics.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("TopMetrics", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type StringOrPipelineMetricAggregationType struct {
	String                        *string                        `json:"String,omitempty"`
	PipelineMetricAggregationType *PipelineMetricAggregationType `json:"PipelineMetricAggregationType,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `StringOrPipelineMetricAggregationType` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *StringOrPipelineMetricAggregationType) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "String"
	if fields["String"] != nil {
		if string(fields["String"]) != "null" {
			if err := json.Unmarshal(fields["String"], &resource.String); err != nil {
				errs = append(errs, cog.MakeBuildErrors("String", err)...)
			}

		}
		delete(fields, "String")

	}
	// Field "PipelineMetricAggregationType"
	if fields["PipelineMetricAggregationType"] != nil {
		if string(fields["PipelineMetricAggregationType"]) != "null" {
			if err := json.Unmarshal(fields["PipelineMetricAggregationType"], &resource.PipelineMetricAggregationType); err != nil {
				errs = append(errs, cog.MakeBuildErrors("PipelineMetricAggregationType", err)...)
			}

		}
		delete(fields, "PipelineMetricAggregationType")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("StringOrPipelineMetricAggregationType", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `StringOrPipelineMetricAggregationType` objects.
func (resource StringOrPipelineMetricAggregationType) Equals(other StringOrPipelineMetricAggregationType) bool {
	if resource.String == nil && other.String != nil || resource.String != nil && other.String == nil {
		return false
	}

	if resource.String != nil {
		if *resource.String != *other.String {
			return false
		}
	}
	if resource.PipelineMetricAggregationType == nil && other.PipelineMetricAggregationType != nil || resource.PipelineMetricAggregationType != nil && other.PipelineMetricAggregationType == nil {
		return false
	}

	if resource.PipelineMetricAggregationType != nil {
		if *resource.PipelineMetricAggregationType != *other.PipelineMetricAggregationType {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `StringOrPipelineMetricAggregationType` fields for violations and returns them.
func (resource StringOrPipelineMetricAggregationType) Validate() error {
	return nil
}

type StringOrElasticsearchInlineScript struct {
	String                    *string                    `json:"String,omitempty"`
	ElasticsearchInlineScript *ElasticsearchInlineScript `json:"ElasticsearchInlineScript,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `StringOrElasticsearchInlineScript` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *StringOrElasticsearchInlineScript) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "String"
	if fields["String"] != nil {
		if string(fields["String"]) != "null" {
			if err := json.Unmarshal(fields["String"], &resource.String); err != nil {
				errs = append(errs, cog.MakeBuildErrors("String", err)...)
			}

		}
		delete(fields, "String")

	}
	// Field "ElasticsearchInlineScript"
	if fields["ElasticsearchInlineScript"] != nil {
		if string(fields["ElasticsearchInlineScript"]) != "null" {

			resource.ElasticsearchInlineScript = &ElasticsearchInlineScript{}
			if err := resource.ElasticsearchInlineScript.UnmarshalJSONStrict(fields["ElasticsearchInlineScript"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("ElasticsearchInlineScript", err)...)
			}

		}
		delete(fields, "ElasticsearchInlineScript")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("StringOrElasticsearchInlineScript", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `StringOrElasticsearchInlineScript` objects.
func (resource StringOrElasticsearchInlineScript) Equals(other StringOrElasticsearchInlineScript) bool {
	if resource.String == nil && other.String != nil || resource.String != nil && other.String == nil {
		return false
	}

	if resource.String != nil {
		if *resource.String != *other.String {
			return false
		}
	}
	if resource.ElasticsearchInlineScript == nil && other.ElasticsearchInlineScript != nil || resource.ElasticsearchInlineScript != nil && other.ElasticsearchInlineScript == nil {
		return false
	}

	if resource.ElasticsearchInlineScript != nil {
		if !resource.ElasticsearchInlineScript.Equals(*other.ElasticsearchInlineScript) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `StringOrElasticsearchInlineScript` fields for violations and returns them.
func (resource StringOrElasticsearchInlineScript) Validate() error {
	var errs cog.BuildErrors
	if resource.ElasticsearchInlineScript != nil {
		if err := resource.ElasticsearchInlineScript.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("ElasticsearchInlineScript", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type MovingAverageOrDerivativeOrCumulativeSumOrBucketScript struct {
	MovingAverage *MovingAverage `json:"MovingAverage,omitempty"`
	Derivative    *Derivative    `json:"Derivative,omitempty"`
	CumulativeSum *CumulativeSum `json:"CumulativeSum,omitempty"`
	BucketScript  *BucketScript  `json:"BucketScript,omitempty"`
}

// MarshalJSON implements a custom JSON marshalling logic to encode `MovingAverageOrDerivativeOrCumulativeSumOrBucketScript` as JSON.
func (resource MovingAverageOrDerivativeOrCumulativeSumOrBucketScript) MarshalJSON() ([]byte, error) {
	if resource.MovingAverage != nil {
		return json.Marshal(resource.MovingAverage)
	}
	if resource.Derivative != nil {
		return json.Marshal(resource.Derivative)
	}
	if resource.CumulativeSum != nil {
		return json.Marshal(resource.CumulativeSum)
	}
	if resource.BucketScript != nil {
		return json.Marshal(resource.BucketScript)
	}

	return nil, fmt.Errorf("no value for disjunction of refs")
}

// UnmarshalJSON implements a custom JSON unmarshalling logic to decode `MovingAverageOrDerivativeOrCumulativeSumOrBucketScript` from JSON.
func (resource *MovingAverageOrDerivativeOrCumulativeSumOrBucketScript) UnmarshalJSON(raw []byte) error {
	if raw == nil {
		return nil
	}

	// FIXME: this is wasteful, we need to find a more efficient way to unmarshal this.
	parsedAsMap := make(map[string]any)
	if err := json.Unmarshal(raw, &parsedAsMap); err != nil {
		return err
	}

	discriminator, found := parsedAsMap["type"]
	if !found {
		return errors.New("discriminator field 'type' not found in payload")
	}

	switch discriminator {
	case "bucket_script":
		var bucketScript BucketScript
		if err := json.Unmarshal(raw, &bucketScript); err != nil {
			return err
		}

		resource.BucketScript = &bucketScript
		return nil
	case "cumulative_sum":
		var cumulativeSum CumulativeSum
		if err := json.Unmarshal(raw, &cumulativeSum); err != nil {
			return err
		}

		resource.CumulativeSum = &cumulativeSum
		return nil
	case "derivative":
		var derivative Derivative
		if err := json.Unmarshal(raw, &derivative); err != nil {
			return err
		}

		resource.Derivative = &derivative
		return nil
	case "moving_avg":
		var movingAverage MovingAverage
		if err := json.Unmarshal(raw, &movingAverage); err != nil {
			return err
		}

		resource.MovingAverage = &movingAverage
		return nil
	}

	return fmt.Errorf("could not unmarshal resource with `type = %v`", discriminator)
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `MovingAverageOrDerivativeOrCumulativeSumOrBucketScript` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *MovingAverageOrDerivativeOrCumulativeSumOrBucketScript) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	// FIXME: this is wasteful, we need to find a more efficient way to unmarshal this.
	parsedAsMap := make(map[string]any)
	if err := json.Unmarshal(raw, &parsedAsMap); err != nil {
		return err
	}

	discriminator, found := parsedAsMap["type"]
	if !found {
		return fmt.Errorf("discriminator field 'type' not found in payload")
	}

	switch discriminator {
	case "bucket_script":
		bucketScript := &BucketScript{}
		if err := bucketScript.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.BucketScript = bucketScript
		return nil
	case "cumulative_sum":
		cumulativeSum := &CumulativeSum{}
		if err := cumulativeSum.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.CumulativeSum = cumulativeSum
		return nil
	case "derivative":
		derivative := &Derivative{}
		if err := derivative.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.Derivative = derivative
		return nil
	case "moving_avg":
		movingAverage := &MovingAverage{}
		if err := movingAverage.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.MovingAverage = movingAverage
		return nil
	}

	return fmt.Errorf("could not unmarshal resource with `type = %v`", discriminator)
}

// Equals tests the equality of two `MovingAverageOrDerivativeOrCumulativeSumOrBucketScript` objects.
func (resource MovingAverageOrDerivativeOrCumulativeSumOrBucketScript) Equals(other MovingAverageOrDerivativeOrCumulativeSumOrBucketScript) bool {
	if resource.MovingAverage == nil && other.MovingAverage != nil || resource.MovingAverage != nil && other.MovingAverage == nil {
		return false
	}

	if resource.MovingAverage != nil {
		if !resource.MovingAverage.Equals(*other.MovingAverage) {
			return false
		}
	}
	if resource.Derivative == nil && other.Derivative != nil || resource.Derivative != nil && other.Derivative == nil {
		return false
	}

	if resource.Derivative != nil {
		if !resource.Derivative.Equals(*other.Derivative) {
			return false
		}
	}
	if resource.CumulativeSum == nil && other.CumulativeSum != nil || resource.CumulativeSum != nil && other.CumulativeSum == nil {
		return false
	}

	if resource.CumulativeSum != nil {
		if !resource.CumulativeSum.Equals(*other.CumulativeSum) {
			return false
		}
	}
	if resource.BucketScript == nil && other.BucketScript != nil || resource.BucketScript != nil && other.BucketScript == nil {
		return false
	}

	if resource.BucketScript != nil {
		if !resource.BucketScript.Equals(*other.BucketScript) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `MovingAverageOrDerivativeOrCumulativeSumOrBucketScript` fields for violations and returns them.
func (resource MovingAverageOrDerivativeOrCumulativeSumOrBucketScript) Validate() error {
	var errs cog.BuildErrors
	if resource.MovingAverage != nil {
		if err := resource.MovingAverage.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("MovingAverage", err)...)
		}
	}
	if resource.Derivative != nil {
		if err := resource.Derivative.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("Derivative", err)...)
		}
	}
	if resource.CumulativeSum != nil {
		if err := resource.CumulativeSum.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("CumulativeSum", err)...)
		}
	}
	if resource.BucketScript != nil {
		if err := resource.BucketScript.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("BucketScript", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type BucketScriptOrCumulativeSumOrDerivativeOrSerialDiffOrRawDataOrRawDocumentOrUniqueCountOrPercentilesOrExtendedStatsOrMinOrMaxOrSumOrAverageOrMovingAverageOrMovingFunctionOrLogsOrRateOrTopMetrics struct {
	BucketScript   *BucketScript   `json:"BucketScript,omitempty"`
	CumulativeSum  *CumulativeSum  `json:"CumulativeSum,omitempty"`
	Derivative     *Derivative     `json:"Derivative,omitempty"`
	SerialDiff     *SerialDiff     `json:"SerialDiff,omitempty"`
	RawData        *RawData        `json:"RawData,omitempty"`
	RawDocument    *RawDocument    `json:"RawDocument,omitempty"`
	UniqueCount    *UniqueCount    `json:"UniqueCount,omitempty"`
	Percentiles    *Percentiles    `json:"Percentiles,omitempty"`
	ExtendedStats  *ExtendedStats  `json:"ExtendedStats,omitempty"`
	Min            *Min            `json:"Min,omitempty"`
	Max            *Max            `json:"Max,omitempty"`
	Sum            *Sum            `json:"Sum,omitempty"`
	Average        *Average        `json:"Average,omitempty"`
	MovingAverage  *MovingAverage  `json:"MovingAverage,omitempty"`
	MovingFunction *MovingFunction `json:"MovingFunction,omitempty"`
	Logs           *Logs           `json:"Logs,omitempty"`
	Rate           *Rate           `json:"Rate,omitempty"`
	TopMetrics     *TopMetrics     `json:"TopMetrics,omitempty"`
}

// MarshalJSON implements a custom JSON marshalling logic to encode `BucketScriptOrCumulativeSumOrDerivativeOrSerialDiffOrRawDataOrRawDocumentOrUniqueCountOrPercentilesOrExtendedStatsOrMinOrMaxOrSumOrAverageOrMovingAverageOrMovingFunctionOrLogsOrRateOrTopMetrics` as JSON.
func (resource BucketScriptOrCumulativeSumOrDerivativeOrSerialDiffOrRawDataOrRawDocumentOrUniqueCountOrPercentilesOrExtendedStatsOrMinOrMaxOrSumOrAverageOrMovingAverageOrMovingFunctionOrLogsOrRateOrTopMetrics) MarshalJSON() ([]byte, error) {
	if resource.BucketScript != nil {
		return json.Marshal(resource.BucketScript)
	}
	if resource.CumulativeSum != nil {
		return json.Marshal(resource.CumulativeSum)
	}
	if resource.Derivative != nil {
		return json.Marshal(resource.Derivative)
	}
	if resource.SerialDiff != nil {
		return json.Marshal(resource.SerialDiff)
	}
	if resource.RawData != nil {
		return json.Marshal(resource.RawData)
	}
	if resource.RawDocument != nil {
		return json.Marshal(resource.RawDocument)
	}
	if resource.UniqueCount != nil {
		return json.Marshal(resource.UniqueCount)
	}
	if resource.Percentiles != nil {
		return json.Marshal(resource.Percentiles)
	}
	if resource.ExtendedStats != nil {
		return json.Marshal(resource.ExtendedStats)
	}
	if resource.Min != nil {
		return json.Marshal(resource.Min)
	}
	if resource.Max != nil {
		return json.Marshal(resource.Max)
	}
	if resource.Sum != nil {
		return json.Marshal(resource.Sum)
	}
	if resource.Average != nil {
		return json.Marshal(resource.Average)
	}
	if resource.MovingAverage != nil {
		return json.Marshal(resource.MovingAverage)
	}
	if resource.MovingFunction != nil {
		return json.Marshal(resource.MovingFunction)
	}
	if resource.Logs != nil {
		return json.Marshal(resource.Logs)
	}
	if resource.Rate != nil {
		return json.Marshal(resource.Rate)
	}
	if resource.TopMetrics != nil {
		return json.Marshal(resource.TopMetrics)
	}

	return nil, fmt.Errorf("no value for disjunction of refs")
}

// UnmarshalJSON implements a custom JSON unmarshalling logic to decode `BucketScriptOrCumulativeSumOrDerivativeOrSerialDiffOrRawDataOrRawDocumentOrUniqueCountOrPercentilesOrExtendedStatsOrMinOrMaxOrSumOrAverageOrMovingAverageOrMovingFunctionOrLogsOrRateOrTopMetrics` from JSON.
func (resource *BucketScriptOrCumulativeSumOrDerivativeOrSerialDiffOrRawDataOrRawDocumentOrUniqueCountOrPercentilesOrExtendedStatsOrMinOrMaxOrSumOrAverageOrMovingAverageOrMovingFunctionOrLogsOrRateOrTopMetrics) UnmarshalJSON(raw []byte) error {
	if raw == nil {
		return nil
	}

	// FIXME: this is wasteful, we need to find a more efficient way to unmarshal this.
	parsedAsMap := make(map[string]any)
	if err := json.Unmarshal(raw, &parsedAsMap); err != nil {
		return err
	}

	discriminator, found := parsedAsMap["type"]
	if !found {
		return errors.New("discriminator field 'type' not found in payload")
	}

	switch discriminator {
	case "avg":
		var average Average
		if err := json.Unmarshal(raw, &average); err != nil {
			return err
		}

		resource.Average = &average
		return nil
	case "bucket_script":
		var bucketScript BucketScript
		if err := json.Unmarshal(raw, &bucketScript); err != nil {
			return err
		}

		resource.BucketScript = &bucketScript
		return nil
	case "cardinality":
		var uniqueCount UniqueCount
		if err := json.Unmarshal(raw, &uniqueCount); err != nil {
			return err
		}

		resource.UniqueCount = &uniqueCount
		return nil
	case "cumulative_sum":
		var cumulativeSum CumulativeSum
		if err := json.Unmarshal(raw, &cumulativeSum); err != nil {
			return err
		}

		resource.CumulativeSum = &cumulativeSum
		return nil
	case "derivative":
		var derivative Derivative
		if err := json.Unmarshal(raw, &derivative); err != nil {
			return err
		}

		resource.Derivative = &derivative
		return nil
	case "extended_stats":
		var extendedStats ExtendedStats
		if err := json.Unmarshal(raw, &extendedStats); err != nil {
			return err
		}

		resource.ExtendedStats = &extendedStats
		return nil
	case "logs":
		var logs Logs
		if err := json.Unmarshal(raw, &logs); err != nil {
			return err
		}

		resource.Logs = &logs
		return nil
	case "max":
		var max Max
		if err := json.Unmarshal(raw, &max); err != nil {
			return err
		}

		resource.Max = &max
		return nil
	case "min":
		var min Min
		if err := json.Unmarshal(raw, &min); err != nil {
			return err
		}

		resource.Min = &min
		return nil
	case "moving_avg":
		var movingAverage MovingAverage
		if err := json.Unmarshal(raw, &movingAverage); err != nil {
			return err
		}

		resource.MovingAverage = &movingAverage
		return nil
	case "moving_fn":
		var movingFunction MovingFunction
		if err := json.Unmarshal(raw, &movingFunction); err != nil {
			return err
		}

		resource.MovingFunction = &movingFunction
		return nil
	case "percentiles":
		var percentiles Percentiles
		if err := json.Unmarshal(raw, &percentiles); err != nil {
			return err
		}

		resource.Percentiles = &percentiles
		return nil
	case "rate":
		var rate Rate
		if err := json.Unmarshal(raw, &rate); err != nil {
			return err
		}

		resource.Rate = &rate
		return nil
	case "raw_data":
		var rawData RawData
		if err := json.Unmarshal(raw, &rawData); err != nil {
			return err
		}

		resource.RawData = &rawData
		return nil
	case "raw_document":
		var rawDocument RawDocument
		if err := json.Unmarshal(raw, &rawDocument); err != nil {
			return err
		}

		resource.RawDocument = &rawDocument
		return nil
	case "serial_diff":
		var serialDiff SerialDiff
		if err := json.Unmarshal(raw, &serialDiff); err != nil {
			return err
		}

		resource.SerialDiff = &serialDiff
		return nil
	case "sum":
		var sum Sum
		if err := json.Unmarshal(raw, &sum); err != nil {
			return err
		}

		resource.Sum = &sum
		return nil
	case "top_metrics":
		var topMetrics TopMetrics
		if err := json.Unmarshal(raw, &topMetrics); err != nil {
			return err
		}

		resource.TopMetrics = &topMetrics
		return nil
	}

	return fmt.Errorf("could not unmarshal resource with `type = %v`", discriminator)
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `BucketScriptOrCumulativeSumOrDerivativeOrSerialDiffOrRawDataOrRawDocumentOrUniqueCountOrPercentilesOrExtendedStatsOrMinOrMaxOrSumOrAverageOrMovingAverageOrMovingFunctionOrLogsOrRateOrTopMetrics` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *BucketScriptOrCumulativeSumOrDerivativeOrSerialDiffOrRawDataOrRawDocumentOrUniqueCountOrPercentilesOrExtendedStatsOrMinOrMaxOrSumOrAverageOrMovingAverageOrMovingFunctionOrLogsOrRateOrTopMetrics) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	// FIXME: this is wasteful, we need to find a more efficient way to unmarshal this.
	parsedAsMap := make(map[string]any)
	if err := json.Unmarshal(raw, &parsedAsMap); err != nil {
		return err
	}

	discriminator, found := parsedAsMap["type"]
	if !found {
		return fmt.Errorf("discriminator field 'type' not found in payload")
	}

	switch discriminator {
	case "avg":
		average := &Average{}
		if err := average.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.Average = average
		return nil
	case "bucket_script":
		bucketScript := &BucketScript{}
		if err := bucketScript.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.BucketScript = bucketScript
		return nil
	case "cardinality":
		uniqueCount := &UniqueCount{}
		if err := uniqueCount.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.UniqueCount = uniqueCount
		return nil
	case "cumulative_sum":
		cumulativeSum := &CumulativeSum{}
		if err := cumulativeSum.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.CumulativeSum = cumulativeSum
		return nil
	case "derivative":
		derivative := &Derivative{}
		if err := derivative.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.Derivative = derivative
		return nil
	case "extended_stats":
		extendedStats := &ExtendedStats{}
		if err := extendedStats.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.ExtendedStats = extendedStats
		return nil
	case "logs":
		logs := &Logs{}
		if err := logs.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.Logs = logs
		return nil
	case "max":
		max := &Max{}
		if err := max.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.Max = max
		return nil
	case "min":
		min := &Min{}
		if err := min.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.Min = min
		return nil
	case "moving_avg":
		movingAverage := &MovingAverage{}
		if err := movingAverage.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.MovingAverage = movingAverage
		return nil
	case "moving_fn":
		movingFunction := &MovingFunction{}
		if err := movingFunction.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.MovingFunction = movingFunction
		return nil
	case "percentiles":
		percentiles := &Percentiles{}
		if err := percentiles.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.Percentiles = percentiles
		return nil
	case "rate":
		rate := &Rate{}
		if err := rate.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.Rate = rate
		return nil
	case "raw_data":
		rawData := &RawData{}
		if err := rawData.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.RawData = rawData
		return nil
	case "raw_document":
		rawDocument := &RawDocument{}
		if err := rawDocument.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.RawDocument = rawDocument
		return nil
	case "serial_diff":
		serialDiff := &SerialDiff{}
		if err := serialDiff.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.SerialDiff = serialDiff
		return nil
	case "sum":
		sum := &Sum{}
		if err := sum.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.Sum = sum
		return nil
	case "top_metrics":
		topMetrics := &TopMetrics{}
		if err := topMetrics.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.TopMetrics = topMetrics
		return nil
	}

	return fmt.Errorf("could not unmarshal resource with `type = %v`", discriminator)
}

// Equals tests the equality of two `BucketScriptOrCumulativeSumOrDerivativeOrSerialDiffOrRawDataOrRawDocumentOrUniqueCountOrPercentilesOrExtendedStatsOrMinOrMaxOrSumOrAverageOrMovingAverageOrMovingFunctionOrLogsOrRateOrTopMetrics` objects.
func (resource BucketScriptOrCumulativeSumOrDerivativeOrSerialDiffOrRawDataOrRawDocumentOrUniqueCountOrPercentilesOrExtendedStatsOrMinOrMaxOrSumOrAverageOrMovingAverageOrMovingFunctionOrLogsOrRateOrTopMetrics) Equals(other BucketScriptOrCumulativeSumOrDerivativeOrSerialDiffOrRawDataOrRawDocumentOrUniqueCountOrPercentilesOrExtendedStatsOrMinOrMaxOrSumOrAverageOrMovingAverageOrMovingFunctionOrLogsOrRateOrTopMetrics) bool {
	if resource.BucketScript == nil && other.BucketScript != nil || resource.BucketScript != nil && other.BucketScript == nil {
		return false
	}

	if resource.BucketScript != nil {
		if !resource.BucketScript.Equals(*other.BucketScript) {
			return false
		}
	}
	if resource.CumulativeSum == nil && other.CumulativeSum != nil || resource.CumulativeSum != nil && other.CumulativeSum == nil {
		return false
	}

	if resource.CumulativeSum != nil {
		if !resource.CumulativeSum.Equals(*other.CumulativeSum) {
			return false
		}
	}
	if resource.Derivative == nil && other.Derivative != nil || resource.Derivative != nil && other.Derivative == nil {
		return false
	}

	if resource.Derivative != nil {
		if !resource.Derivative.Equals(*other.Derivative) {
			return false
		}
	}
	if resource.SerialDiff == nil && other.SerialDiff != nil || resource.SerialDiff != nil && other.SerialDiff == nil {
		return false
	}

	if resource.SerialDiff != nil {
		if !resource.SerialDiff.Equals(*other.SerialDiff) {
			return false
		}
	}
	if resource.RawData == nil && other.RawData != nil || resource.RawData != nil && other.RawData == nil {
		return false
	}

	if resource.RawData != nil {
		if !resource.RawData.Equals(*other.RawData) {
			return false
		}
	}
	if resource.RawDocument == nil && other.RawDocument != nil || resource.RawDocument != nil && other.RawDocument == nil {
		return false
	}

	if resource.RawDocument != nil {
		if !resource.RawDocument.Equals(*other.RawDocument) {
			return false
		}
	}
	if resource.UniqueCount == nil && other.UniqueCount != nil || resource.UniqueCount != nil && other.UniqueCount == nil {
		return false
	}

	if resource.UniqueCount != nil {
		if !resource.UniqueCount.Equals(*other.UniqueCount) {
			return false
		}
	}
	if resource.Percentiles == nil && other.Percentiles != nil || resource.Percentiles != nil && other.Percentiles == nil {
		return false
	}

	if resource.Percentiles != nil {
		if !resource.Percentiles.Equals(*other.Percentiles) {
			return false
		}
	}
	if resource.ExtendedStats == nil && other.ExtendedStats != nil || resource.ExtendedStats != nil && other.ExtendedStats == nil {
		return false
	}

	if resource.ExtendedStats != nil {
		if !resource.ExtendedStats.Equals(*other.ExtendedStats) {
			return false
		}
	}
	if resource.Min == nil && other.Min != nil || resource.Min != nil && other.Min == nil {
		return false
	}

	if resource.Min != nil {
		if !resource.Min.Equals(*other.Min) {
			return false
		}
	}
	if resource.Max == nil && other.Max != nil || resource.Max != nil && other.Max == nil {
		return false
	}

	if resource.Max != nil {
		if !resource.Max.Equals(*other.Max) {
			return false
		}
	}
	if resource.Sum == nil && other.Sum != nil || resource.Sum != nil && other.Sum == nil {
		return false
	}

	if resource.Sum != nil {
		if !resource.Sum.Equals(*other.Sum) {
			return false
		}
	}
	if resource.Average == nil && other.Average != nil || resource.Average != nil && other.Average == nil {
		return false
	}

	if resource.Average != nil {
		if !resource.Average.Equals(*other.Average) {
			return false
		}
	}
	if resource.MovingAverage == nil && other.MovingAverage != nil || resource.MovingAverage != nil && other.MovingAverage == nil {
		return false
	}

	if resource.MovingAverage != nil {
		if !resource.MovingAverage.Equals(*other.MovingAverage) {
			return false
		}
	}
	if resource.MovingFunction == nil && other.MovingFunction != nil || resource.MovingFunction != nil && other.MovingFunction == nil {
		return false
	}

	if resource.MovingFunction != nil {
		if !resource.MovingFunction.Equals(*other.MovingFunction) {
			return false
		}
	}
	if resource.Logs == nil && other.Logs != nil || resource.Logs != nil && other.Logs == nil {
		return false
	}

	if resource.Logs != nil {
		if !resource.Logs.Equals(*other.Logs) {
			return false
		}
	}
	if resource.Rate == nil && other.Rate != nil || resource.Rate != nil && other.Rate == nil {
		return false
	}

	if resource.Rate != nil {
		if !resource.Rate.Equals(*other.Rate) {
			return false
		}
	}
	if resource.TopMetrics == nil && other.TopMetrics != nil || resource.TopMetrics != nil && other.TopMetrics == nil {
		return false
	}

	if resource.TopMetrics != nil {
		if !resource.TopMetrics.Equals(*other.TopMetrics) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `BucketScriptOrCumulativeSumOrDerivativeOrSerialDiffOrRawDataOrRawDocumentOrUniqueCountOrPercentilesOrExtendedStatsOrMinOrMaxOrSumOrAverageOrMovingAverageOrMovingFunctionOrLogsOrRateOrTopMetrics` fields for violations and returns them.
func (resource BucketScriptOrCumulativeSumOrDerivativeOrSerialDiffOrRawDataOrRawDocumentOrUniqueCountOrPercentilesOrExtendedStatsOrMinOrMaxOrSumOrAverageOrMovingAverageOrMovingFunctionOrLogsOrRateOrTopMetrics) Validate() error {
	var errs cog.BuildErrors
	if resource.BucketScript != nil {
		if err := resource.BucketScript.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("BucketScript", err)...)
		}
	}
	if resource.CumulativeSum != nil {
		if err := resource.CumulativeSum.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("CumulativeSum", err)...)
		}
	}
	if resource.Derivative != nil {
		if err := resource.Derivative.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("Derivative", err)...)
		}
	}
	if resource.SerialDiff != nil {
		if err := resource.SerialDiff.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("SerialDiff", err)...)
		}
	}
	if resource.RawData != nil {
		if err := resource.RawData.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("RawData", err)...)
		}
	}
	if resource.RawDocument != nil {
		if err := resource.RawDocument.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("RawDocument", err)...)
		}
	}
	if resource.UniqueCount != nil {
		if err := resource.UniqueCount.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("UniqueCount", err)...)
		}
	}
	if resource.Percentiles != nil {
		if err := resource.Percentiles.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("Percentiles", err)...)
		}
	}
	if resource.ExtendedStats != nil {
		if err := resource.ExtendedStats.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("ExtendedStats", err)...)
		}
	}
	if resource.Min != nil {
		if err := resource.Min.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("Min", err)...)
		}
	}
	if resource.Max != nil {
		if err := resource.Max.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("Max", err)...)
		}
	}
	if resource.Sum != nil {
		if err := resource.Sum.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("Sum", err)...)
		}
	}
	if resource.Average != nil {
		if err := resource.Average.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("Average", err)...)
		}
	}
	if resource.MovingAverage != nil {
		if err := resource.MovingAverage.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("MovingAverage", err)...)
		}
	}
	if resource.MovingFunction != nil {
		if err := resource.MovingFunction.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("MovingFunction", err)...)
		}
	}
	if resource.Logs != nil {
		if err := resource.Logs.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("Logs", err)...)
		}
	}
	if resource.Rate != nil {
		if err := resource.Rate.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("Rate", err)...)
		}
	}
	if resource.TopMetrics != nil {
		if err := resource.TopMetrics.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("TopMetrics", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}
