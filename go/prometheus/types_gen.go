// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package prometheus

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
	dashboardv2 "github.com/grafana/grafana-foundation-sdk/go/dashboardv2"
	dashboardv2beta1 "github.com/grafana/grafana-foundation-sdk/go/dashboardv2beta1"
)

type AdhocFilters struct {
	Key      string   `json:"key"`
	Operator string   `json:"operator"`
	Value    string   `json:"value"`
	Values   []string `json:"values,omitempty"`
}

// NewAdhocFilters creates a new AdhocFilters object.
func NewAdhocFilters() *AdhocFilters {
	return &AdhocFilters{}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `AdhocFilters` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *AdhocFilters) UnmarshalJSONStrict(raw []byte) error {
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
			if err := json.Unmarshal(fields["key"], &resource.Key); err != nil {
				errs = append(errs, cog.MakeBuildErrors("key", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("key", errors.New("required field is null"))...)

		}
		delete(fields, "key")
	} else {
		errs = append(errs, cog.MakeBuildErrors("key", errors.New("required field is missing from input"))...)
	}
	// Field "operator"
	if fields["operator"] != nil {
		if string(fields["operator"]) != "null" {
			if err := json.Unmarshal(fields["operator"], &resource.Operator); err != nil {
				errs = append(errs, cog.MakeBuildErrors("operator", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("operator", errors.New("required field is null"))...)

		}
		delete(fields, "operator")
	} else {
		errs = append(errs, cog.MakeBuildErrors("operator", errors.New("required field is missing from input"))...)
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
	// Field "values"
	if fields["values"] != nil {
		if string(fields["values"]) != "null" {

			if err := json.Unmarshal(fields["values"], &resource.Values); err != nil {
				errs = append(errs, cog.MakeBuildErrors("values", err)...)
			}

		}
		delete(fields, "values")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("AdhocFilters", fmt.Errorf("unexpected field '%s'", field))...)
	}

	return errs
}

// Equals tests the equality of two `AdhocFilters` objects.
func (resource AdhocFilters) Equals(other AdhocFilters) bool {
	if resource.Key != other.Key {
		return false
	}
	if resource.Operator != other.Operator {
		return false
	}
	if resource.Value != other.Value {
		return false
	}

	if len(resource.Values) != len(other.Values) {
		return false
	}

	for i1 := range resource.Values {
		if resource.Values[i1] != other.Values[i1] {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `AdhocFilters` fields for violations and returns them.
func (resource AdhocFilters) Validate() error {
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

// NewResultAssertions creates a new ResultAssertions object.
func NewResultAssertions() *ResultAssertions {
	return &ResultAssertions{
		TypeVersion: []int64{},
	}
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

type ScopesFilters struct {
	Key      string   `json:"key"`
	Operator string   `json:"operator"`
	Value    string   `json:"value"`
	Values   []string `json:"values,omitempty"`
}

// NewScopesFilters creates a new ScopesFilters object.
func NewScopesFilters() *ScopesFilters {
	return &ScopesFilters{}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ScopesFilters` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ScopesFilters) UnmarshalJSONStrict(raw []byte) error {
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
			if err := json.Unmarshal(fields["key"], &resource.Key); err != nil {
				errs = append(errs, cog.MakeBuildErrors("key", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("key", errors.New("required field is null"))...)

		}
		delete(fields, "key")
	} else {
		errs = append(errs, cog.MakeBuildErrors("key", errors.New("required field is missing from input"))...)
	}
	// Field "operator"
	if fields["operator"] != nil {
		if string(fields["operator"]) != "null" {
			if err := json.Unmarshal(fields["operator"], &resource.Operator); err != nil {
				errs = append(errs, cog.MakeBuildErrors("operator", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("operator", errors.New("required field is null"))...)

		}
		delete(fields, "operator")
	} else {
		errs = append(errs, cog.MakeBuildErrors("operator", errors.New("required field is missing from input"))...)
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
	// Field "values"
	if fields["values"] != nil {
		if string(fields["values"]) != "null" {

			if err := json.Unmarshal(fields["values"], &resource.Values); err != nil {
				errs = append(errs, cog.MakeBuildErrors("values", err)...)
			}

		}
		delete(fields, "values")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ScopesFilters", fmt.Errorf("unexpected field '%s'", field))...)
	}

	return errs
}

// Equals tests the equality of two `ScopesFilters` objects.
func (resource ScopesFilters) Equals(other ScopesFilters) bool {
	if resource.Key != other.Key {
		return false
	}
	if resource.Operator != other.Operator {
		return false
	}
	if resource.Value != other.Value {
		return false
	}

	if len(resource.Values) != len(other.Values) {
		return false
	}

	for i1 := range resource.Values {
		if resource.Values[i1] != other.Values[i1] {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `ScopesFilters` fields for violations and returns them.
func (resource ScopesFilters) Validate() error {
	return nil
}

type Scopes struct {
	DefaultPath []string        `json:"defaultPath,omitempty"`
	Filters     []ScopesFilters `json:"filters,omitempty"`
	Title       string          `json:"title"`
}

// NewScopes creates a new Scopes object.
func NewScopes() *Scopes {
	return &Scopes{}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Scopes` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *Scopes) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "defaultPath"
	if fields["defaultPath"] != nil {
		if string(fields["defaultPath"]) != "null" {

			if err := json.Unmarshal(fields["defaultPath"], &resource.DefaultPath); err != nil {
				errs = append(errs, cog.MakeBuildErrors("defaultPath", err)...)
			}

		}
		delete(fields, "defaultPath")

	}
	// Field "filters"
	if fields["filters"] != nil {
		if string(fields["filters"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["filters"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 ScopesFilters

				result1 = ScopesFilters{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("filters["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.Filters = append(resource.Filters, result1)
			}

		}
		delete(fields, "filters")

	}
	// Field "title"
	if fields["title"] != nil {
		if string(fields["title"]) != "null" {
			if err := json.Unmarshal(fields["title"], &resource.Title); err != nil {
				errs = append(errs, cog.MakeBuildErrors("title", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("title", errors.New("required field is null"))...)

		}
		delete(fields, "title")
	} else {
		errs = append(errs, cog.MakeBuildErrors("title", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("Scopes", fmt.Errorf("unexpected field '%s'", field))...)
	}

	return errs
}

// Equals tests the equality of two `Scopes` objects.
func (resource Scopes) Equals(other Scopes) bool {

	if len(resource.DefaultPath) != len(other.DefaultPath) {
		return false
	}

	for i1 := range resource.DefaultPath {
		if resource.DefaultPath[i1] != other.DefaultPath[i1] {
			return false
		}
	}

	if len(resource.Filters) != len(other.Filters) {
		return false
	}

	for i1 := range resource.Filters {
		if !resource.Filters[i1].Equals(other.Filters[i1]) {
			return false
		}
	}
	if resource.Title != other.Title {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `Scopes` fields for violations and returns them.
func (resource Scopes) Validate() error {
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

type TimeRange struct {
	// From is the start time of the query.
	From string `json:"from"`
	// To is the end time of the query.
	To string `json:"to"`
}

// NewTimeRange creates a new TimeRange object.
func NewTimeRange() *TimeRange {
	return &TimeRange{
		From: "now-6h",
		To:   "now",
	}
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

type Dataquery struct {
	// Additional Ad-hoc filters that take precedence over Scope on conflict.
	AdhocFilters []AdhocFilters `json:"adhocFilters,omitempty"`
	// The datasource
	Datasource *common.DataSourceRef `json:"datasource,omitempty"`
	// what we should show in the editor
	// Possible enum values:
	//  - `"builder"`
	//  - `"code"`
	EditorMode *QueryEditorMode `json:"editorMode,omitempty"`
	// Execute an additional query to identify interesting raw samples relevant for the given expr
	Exemplar *bool `json:"exemplar,omitempty"`
	// The actual expression/query that will be evaluated by Prometheus
	Expr string `json:"expr"`
	// The response format
	// Possible enum values:
	//  - `"time_series"`
	//  - `"table"`
	//  - `"heatmap"`
	Format *PromQueryFormat `json:"format,omitempty"`
	// Group By parameters to apply to aggregate expressions in the query
	GroupByKeys []string `json:"groupByKeys,omitempty"`
	// true if query is disabled (ie should not be returned to the dashboard)
	// NOTE: this does not always imply that the query should not be executed since
	// the results from a hidden query may be used as the input to other queries (SSE etc)
	Hide *bool `json:"hide,omitempty"`
	// Returns only the latest value that Prometheus has scraped for the requested time series
	Instant *bool `json:"instant,omitempty"`
	// Used to specify how many times to divide max data points by. We use max data points under query options
	// See https://github.com/grafana/grafana/issues/48081
	// Deprecated: use interval
	IntervalFactor *int64 `json:"intervalFactor,omitempty"`
	// Interval is the suggested duration between time points in a time series query.
	// NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
	// from the interval required to fill a pixels in the visualization
	IntervalMs *float64 `json:"intervalMs,omitempty"`
	// Series name override or template. Ex. {{hostname}} will be replaced with label value for hostname
	LegendFormat *string `json:"legendFormat,omitempty"`
	// MaxDataPoints is the maximum number of data points that should be returned from a time series query.
	// NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
	// from the number of pixels visible in a visualization
	MaxDataPoints *int64 `json:"maxDataPoints,omitempty"`
	// QueryType is an optional identifier for the type of query.
	// It can be used to distinguish different types of queries.
	QueryType *string `json:"queryType,omitempty"`
	// Returns a Range vector, comprised of a set of time series containing a range of data points over time for each time series
	Range *bool `json:"range,omitempty"`
	// RefID is the unique identifier of the query, set by the frontend call.
	RefId *string `json:"refId,omitempty"`
	// Optionally define expected query result behavior
	ResultAssertions *ResultAssertions `json:"resultAssertions,omitempty"`
	// A set of filters applied to apply to the query
	Scopes []Scopes `json:"scopes,omitempty"`
	// TimeRange represents the query range
	// NOTE: unlike generic /ds/query, we can now send explicit time values in each query
	// NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
	TimeRange *TimeRange `json:"timeRange,omitempty"`
	// An additional lower limit for the step parameter of the Prometheus query and for the
	// `$__interval` and `$__rate_interval` variables.
	Interval *string `json:"interval,omitempty"`
}

func (resource Dataquery) ImplementsDataqueryVariant() {}

func (resource Dataquery) DataqueryType() string {
	return "prometheus"
}

// NewDataquery creates a new Dataquery object.
func NewDataquery() *Dataquery {
	return &Dataquery{}
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
	// Field "adhocFilters"
	if fields["adhocFilters"] != nil {
		if string(fields["adhocFilters"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["adhocFilters"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 AdhocFilters

				result1 = AdhocFilters{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("adhocFilters["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.AdhocFilters = append(resource.AdhocFilters, result1)
			}

		}
		delete(fields, "adhocFilters")

	}
	// Field "datasource"
	if fields["datasource"] != nil {
		if string(fields["datasource"]) != "null" {

			resource.Datasource = &common.DataSourceRef{}
			if err := resource.Datasource.UnmarshalJSONStrict(fields["datasource"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("datasource", err)...)
			}

		}
		delete(fields, "datasource")

	}
	// Field "editorMode"
	if fields["editorMode"] != nil {
		if string(fields["editorMode"]) != "null" {
			if err := json.Unmarshal(fields["editorMode"], &resource.EditorMode); err != nil {
				errs = append(errs, cog.MakeBuildErrors("editorMode", err)...)
			}

		}
		delete(fields, "editorMode")

	}
	// Field "exemplar"
	if fields["exemplar"] != nil {
		if string(fields["exemplar"]) != "null" {
			if err := json.Unmarshal(fields["exemplar"], &resource.Exemplar); err != nil {
				errs = append(errs, cog.MakeBuildErrors("exemplar", err)...)
			}

		}
		delete(fields, "exemplar")

	}
	// Field "expr"
	if fields["expr"] != nil {
		if string(fields["expr"]) != "null" {
			if err := json.Unmarshal(fields["expr"], &resource.Expr); err != nil {
				errs = append(errs, cog.MakeBuildErrors("expr", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("expr", errors.New("required field is null"))...)

		}
		delete(fields, "expr")
	} else {
		errs = append(errs, cog.MakeBuildErrors("expr", errors.New("required field is missing from input"))...)
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
	// Field "groupByKeys"
	if fields["groupByKeys"] != nil {
		if string(fields["groupByKeys"]) != "null" {

			if err := json.Unmarshal(fields["groupByKeys"], &resource.GroupByKeys); err != nil {
				errs = append(errs, cog.MakeBuildErrors("groupByKeys", err)...)
			}

		}
		delete(fields, "groupByKeys")

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
	// Field "instant"
	if fields["instant"] != nil {
		if string(fields["instant"]) != "null" {
			if err := json.Unmarshal(fields["instant"], &resource.Instant); err != nil {
				errs = append(errs, cog.MakeBuildErrors("instant", err)...)
			}

		}
		delete(fields, "instant")

	}
	// Field "intervalFactor"
	if fields["intervalFactor"] != nil {
		if string(fields["intervalFactor"]) != "null" {
			if err := json.Unmarshal(fields["intervalFactor"], &resource.IntervalFactor); err != nil {
				errs = append(errs, cog.MakeBuildErrors("intervalFactor", err)...)
			}

		}
		delete(fields, "intervalFactor")

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
	// Field "legendFormat"
	if fields["legendFormat"] != nil {
		if string(fields["legendFormat"]) != "null" {
			if err := json.Unmarshal(fields["legendFormat"], &resource.LegendFormat); err != nil {
				errs = append(errs, cog.MakeBuildErrors("legendFormat", err)...)
			}

		}
		delete(fields, "legendFormat")

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
	// Field "queryType"
	if fields["queryType"] != nil {
		if string(fields["queryType"]) != "null" {
			if err := json.Unmarshal(fields["queryType"], &resource.QueryType); err != nil {
				errs = append(errs, cog.MakeBuildErrors("queryType", err)...)
			}

		}
		delete(fields, "queryType")

	}
	// Field "range"
	if fields["range"] != nil {
		if string(fields["range"]) != "null" {
			if err := json.Unmarshal(fields["range"], &resource.Range); err != nil {
				errs = append(errs, cog.MakeBuildErrors("range", err)...)
			}

		}
		delete(fields, "range")

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
	// Field "scopes"
	if fields["scopes"] != nil {
		if string(fields["scopes"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["scopes"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 Scopes

				result1 = Scopes{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("scopes["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.Scopes = append(resource.Scopes, result1)
			}

		}
		delete(fields, "scopes")

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
	// Field "interval"
	if fields["interval"] != nil {
		if string(fields["interval"]) != "null" {
			if err := json.Unmarshal(fields["interval"], &resource.Interval); err != nil {
				errs = append(errs, cog.MakeBuildErrors("interval", err)...)
			}

		}
		delete(fields, "interval")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("Dataquery", fmt.Errorf("unexpected field '%s'", field))...)
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

	if len(resource.AdhocFilters) != len(other.AdhocFilters) {
		return false
	}

	for i1 := range resource.AdhocFilters {
		if !resource.AdhocFilters[i1].Equals(other.AdhocFilters[i1]) {
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
	if resource.EditorMode == nil && other.EditorMode != nil || resource.EditorMode != nil && other.EditorMode == nil {
		return false
	}

	if resource.EditorMode != nil {
		if *resource.EditorMode != *other.EditorMode {
			return false
		}
	}
	if resource.Exemplar == nil && other.Exemplar != nil || resource.Exemplar != nil && other.Exemplar == nil {
		return false
	}

	if resource.Exemplar != nil {
		if *resource.Exemplar != *other.Exemplar {
			return false
		}
	}
	if resource.Expr != other.Expr {
		return false
	}
	if resource.Format == nil && other.Format != nil || resource.Format != nil && other.Format == nil {
		return false
	}

	if resource.Format != nil {
		if *resource.Format != *other.Format {
			return false
		}
	}

	if len(resource.GroupByKeys) != len(other.GroupByKeys) {
		return false
	}

	for i1 := range resource.GroupByKeys {
		if resource.GroupByKeys[i1] != other.GroupByKeys[i1] {
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
	if resource.Instant == nil && other.Instant != nil || resource.Instant != nil && other.Instant == nil {
		return false
	}

	if resource.Instant != nil {
		if *resource.Instant != *other.Instant {
			return false
		}
	}
	if resource.IntervalFactor == nil && other.IntervalFactor != nil || resource.IntervalFactor != nil && other.IntervalFactor == nil {
		return false
	}

	if resource.IntervalFactor != nil {
		if *resource.IntervalFactor != *other.IntervalFactor {
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
	if resource.LegendFormat == nil && other.LegendFormat != nil || resource.LegendFormat != nil && other.LegendFormat == nil {
		return false
	}

	if resource.LegendFormat != nil {
		if *resource.LegendFormat != *other.LegendFormat {
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
	if resource.QueryType == nil && other.QueryType != nil || resource.QueryType != nil && other.QueryType == nil {
		return false
	}

	if resource.QueryType != nil {
		if *resource.QueryType != *other.QueryType {
			return false
		}
	}
	if resource.Range == nil && other.Range != nil || resource.Range != nil && other.Range == nil {
		return false
	}

	if resource.Range != nil {
		if *resource.Range != *other.Range {
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

	if len(resource.Scopes) != len(other.Scopes) {
		return false
	}

	for i1 := range resource.Scopes {
		if !resource.Scopes[i1].Equals(other.Scopes[i1]) {
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
	if resource.Interval == nil && other.Interval != nil || resource.Interval != nil && other.Interval == nil {
		return false
	}

	if resource.Interval != nil {
		if *resource.Interval != *other.Interval {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `Dataquery` fields for violations and returns them.
func (resource Dataquery) Validate() error {
	var errs cog.BuildErrors

	for i1 := range resource.AdhocFilters {
		if err := resource.AdhocFilters[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("adhocFilters["+strconv.Itoa(i1)+"]", err)...)
		}
	}
	if resource.Datasource != nil {
		if err := resource.Datasource.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("datasource", err)...)
		}
	}
	if resource.ResultAssertions != nil {
		if err := resource.ResultAssertions.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("resultAssertions", err)...)
		}
	}

	for i1 := range resource.Scopes {
		if err := resource.Scopes[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("scopes["+strconv.Itoa(i1)+"]", err)...)
		}
	}
	if resource.TimeRange != nil {
		if err := resource.TimeRange.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("timeRange", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

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

type QueryEditorMode string

const (
	QueryEditorModeBuilder QueryEditorMode = "builder"
	QueryEditorModeCode    QueryEditorMode = "code"
)

type PromQueryFormat string

const (
	PromQueryFormatTimeSeries PromQueryFormat = "time_series"
	PromQueryFormatTable      PromQueryFormat = "table"
	PromQueryFormatHeatmap    PromQueryFormat = "heatmap"
)

// VariantConfig returns the configuration related to prometheus dataqueries.
// This configuration describes how to unmarshal it, convert it to code, …
func VariantConfig() variants.DataqueryConfig {
	return variants.DataqueryConfig{
		Identifier: "prometheus",
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
		GoV2Converter: func(input any) string {
			if cast, ok := input.(*dashboardv2beta1.DataQueryKind); ok {
				return QueryConverter(*cast)
			} else if cast, ok := input.(dashboardv2beta1.DataQueryKind); ok {
				return QueryConverter(cast)
			}
			if cast, ok := input.(*dashboardv2.DataQueryKind); ok {
				return QueryV2Converter(*cast)
			} else if cast, ok := input.(dashboardv2.DataQueryKind); ok {
				return QueryV2Converter(cast)
			}

			return "/* could not convert DataQueryKind */"
		},
	}
}
