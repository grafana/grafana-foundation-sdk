// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package tempo

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	dashboard "github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

type TempoQuery struct {
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
	// TraceQL query or trace ID
	Query *string `json:"query,omitempty"`
	// @deprecated Logfmt query to filter traces by their tags. Example: http.status_code=200 error=true
	Search *string `json:"search,omitempty"`
	// @deprecated Query traces by service name
	ServiceName *string `json:"serviceName,omitempty"`
	// @deprecated Query traces by span name
	SpanName *string `json:"spanName,omitempty"`
	// @deprecated Define the minimum duration to select traces. Use duration format, for example: 1.2s, 100ms
	MinDuration *string `json:"minDuration,omitempty"`
	// @deprecated Define the maximum duration to select traces. Use duration format, for example: 1.2s, 100ms
	MaxDuration *string `json:"maxDuration,omitempty"`
	// Filters to be included in a PromQL query to select data for the service graph. Example: {client="app",service="app"}
	ServiceMapQuery *string `json:"serviceMapQuery,omitempty"`
	// Use service.namespace in addition to service.name to uniquely identify a service.
	ServiceMapIncludeNamespace *bool `json:"serviceMapIncludeNamespace,omitempty"`
	// Defines the maximum number of traces that are returned from Tempo
	Limit *int64 `json:"limit,omitempty"`
	// Defines the maximum number of spans per spanset that are returned from Tempo
	Spss    *int64          `json:"spss,omitempty"`
	Filters []TraceqlFilter `json:"filters"`
	// Filters that are used to query the metrics summary
	GroupBy []TraceqlFilter `json:"groupBy,omitempty"`
	// For mixed data sources the selected datasource is on the query level.
	// For non mixed scenarios this is undefined.
	// TODO find a better way to do this ^ that's friendly to schema
	// TODO this shouldn't be unknown but DataSourceRef | null
	Datasource *dashboard.DataSourceRef `json:"datasource,omitempty"`
	// The type of the table that is used to display the search results
	TableType *SearchTableType `json:"tableType,omitempty"`
}

func (resource TempoQuery) ImplementsDataqueryVariant() {}

func (resource TempoQuery) DataqueryType() string {
	return "tempo"
}

// VariantConfig returns the configuration related to tempo dataqueries.
// This configuration describes how to unmarshal it, convert it to code, …
func VariantConfig() variants.DataqueryConfig {
	return variants.DataqueryConfig{
		Identifier: "tempo",
		DataqueryUnmarshaler: func(raw []byte) (variants.Dataquery, error) {
			dataquery := &TempoQuery{}

			if err := json.Unmarshal(raw, dataquery); err != nil {
				return nil, err
			}

			return dataquery, nil
		},
		StrictDataqueryUnmarshaler: func(raw []byte) (variants.Dataquery, error) {
			dataquery := &TempoQuery{}

			if err := dataquery.UnmarshalJSONStrict(raw); err != nil {
				return nil, err
			}

			return dataquery, nil
		},
		GoConverter: func(input any) string {
			var dataquery TempoQuery
			if cast, ok := input.(*TempoQuery); ok {
				dataquery = *cast
			} else {
				dataquery = input.(TempoQuery)
			}
			return TempoQueryConverter(dataquery)
		},
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TempoQuery` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *TempoQuery) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
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
	// Field "query"
	if fields["query"] != nil {
		if string(fields["query"]) != "null" {
			if err := json.Unmarshal(fields["query"], &resource.Query); err != nil {
				errs = append(errs, cog.MakeBuildErrors("query", err)...)
			}

		}
		delete(fields, "query")

	}
	// Field "search"
	if fields["search"] != nil {
		if string(fields["search"]) != "null" {
			if err := json.Unmarshal(fields["search"], &resource.Search); err != nil {
				errs = append(errs, cog.MakeBuildErrors("search", err)...)
			}

		}
		delete(fields, "search")

	}
	// Field "serviceName"
	if fields["serviceName"] != nil {
		if string(fields["serviceName"]) != "null" {
			if err := json.Unmarshal(fields["serviceName"], &resource.ServiceName); err != nil {
				errs = append(errs, cog.MakeBuildErrors("serviceName", err)...)
			}

		}
		delete(fields, "serviceName")

	}
	// Field "spanName"
	if fields["spanName"] != nil {
		if string(fields["spanName"]) != "null" {
			if err := json.Unmarshal(fields["spanName"], &resource.SpanName); err != nil {
				errs = append(errs, cog.MakeBuildErrors("spanName", err)...)
			}

		}
		delete(fields, "spanName")

	}
	// Field "minDuration"
	if fields["minDuration"] != nil {
		if string(fields["minDuration"]) != "null" {
			if err := json.Unmarshal(fields["minDuration"], &resource.MinDuration); err != nil {
				errs = append(errs, cog.MakeBuildErrors("minDuration", err)...)
			}

		}
		delete(fields, "minDuration")

	}
	// Field "maxDuration"
	if fields["maxDuration"] != nil {
		if string(fields["maxDuration"]) != "null" {
			if err := json.Unmarshal(fields["maxDuration"], &resource.MaxDuration); err != nil {
				errs = append(errs, cog.MakeBuildErrors("maxDuration", err)...)
			}

		}
		delete(fields, "maxDuration")

	}
	// Field "serviceMapQuery"
	if fields["serviceMapQuery"] != nil {
		if string(fields["serviceMapQuery"]) != "null" {
			if err := json.Unmarshal(fields["serviceMapQuery"], &resource.ServiceMapQuery); err != nil {
				errs = append(errs, cog.MakeBuildErrors("serviceMapQuery", err)...)
			}

		}
		delete(fields, "serviceMapQuery")

	}
	// Field "serviceMapIncludeNamespace"
	if fields["serviceMapIncludeNamespace"] != nil {
		if string(fields["serviceMapIncludeNamespace"]) != "null" {
			if err := json.Unmarshal(fields["serviceMapIncludeNamespace"], &resource.ServiceMapIncludeNamespace); err != nil {
				errs = append(errs, cog.MakeBuildErrors("serviceMapIncludeNamespace", err)...)
			}

		}
		delete(fields, "serviceMapIncludeNamespace")

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
	// Field "spss"
	if fields["spss"] != nil {
		if string(fields["spss"]) != "null" {
			if err := json.Unmarshal(fields["spss"], &resource.Spss); err != nil {
				errs = append(errs, cog.MakeBuildErrors("spss", err)...)
			}

		}
		delete(fields, "spss")

	}
	// Field "filters"
	if fields["filters"] != nil {
		if string(fields["filters"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["filters"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 TraceqlFilter

				result1 = TraceqlFilter{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("filters["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.Filters = append(resource.Filters, result1)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("filters", errors.New("required field is null"))...)

		}
		delete(fields, "filters")
	} else {
		errs = append(errs, cog.MakeBuildErrors("filters", errors.New("required field is missing from input"))...)
	}
	// Field "groupBy"
	if fields["groupBy"] != nil {
		if string(fields["groupBy"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["groupBy"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 TraceqlFilter

				result1 = TraceqlFilter{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("groupBy["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.GroupBy = append(resource.GroupBy, result1)
			}

		}
		delete(fields, "groupBy")

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
	// Field "tableType"
	if fields["tableType"] != nil {
		if string(fields["tableType"]) != "null" {
			if err := json.Unmarshal(fields["tableType"], &resource.TableType); err != nil {
				errs = append(errs, cog.MakeBuildErrors("tableType", err)...)
			}

		}
		delete(fields, "tableType")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("TempoQuery", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two dataqueries.
func (resource TempoQuery) Equals(otherCandidate variants.Dataquery) bool {
	if otherCandidate == nil {
		return false
	}

	other, ok := otherCandidate.(TempoQuery)
	if !ok {
		return false
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
	if resource.Query == nil && other.Query != nil || resource.Query != nil && other.Query == nil {
		return false
	}

	if resource.Query != nil {
		if *resource.Query != *other.Query {
			return false
		}
	}
	if resource.Search == nil && other.Search != nil || resource.Search != nil && other.Search == nil {
		return false
	}

	if resource.Search != nil {
		if *resource.Search != *other.Search {
			return false
		}
	}
	if resource.ServiceName == nil && other.ServiceName != nil || resource.ServiceName != nil && other.ServiceName == nil {
		return false
	}

	if resource.ServiceName != nil {
		if *resource.ServiceName != *other.ServiceName {
			return false
		}
	}
	if resource.SpanName == nil && other.SpanName != nil || resource.SpanName != nil && other.SpanName == nil {
		return false
	}

	if resource.SpanName != nil {
		if *resource.SpanName != *other.SpanName {
			return false
		}
	}
	if resource.MinDuration == nil && other.MinDuration != nil || resource.MinDuration != nil && other.MinDuration == nil {
		return false
	}

	if resource.MinDuration != nil {
		if *resource.MinDuration != *other.MinDuration {
			return false
		}
	}
	if resource.MaxDuration == nil && other.MaxDuration != nil || resource.MaxDuration != nil && other.MaxDuration == nil {
		return false
	}

	if resource.MaxDuration != nil {
		if *resource.MaxDuration != *other.MaxDuration {
			return false
		}
	}
	if resource.ServiceMapQuery == nil && other.ServiceMapQuery != nil || resource.ServiceMapQuery != nil && other.ServiceMapQuery == nil {
		return false
	}

	if resource.ServiceMapQuery != nil {
		if *resource.ServiceMapQuery != *other.ServiceMapQuery {
			return false
		}
	}
	if resource.ServiceMapIncludeNamespace == nil && other.ServiceMapIncludeNamespace != nil || resource.ServiceMapIncludeNamespace != nil && other.ServiceMapIncludeNamespace == nil {
		return false
	}

	if resource.ServiceMapIncludeNamespace != nil {
		if *resource.ServiceMapIncludeNamespace != *other.ServiceMapIncludeNamespace {
			return false
		}
	}
	if resource.Limit == nil && other.Limit != nil || resource.Limit != nil && other.Limit == nil {
		return false
	}

	if resource.Limit != nil {
		if *resource.Limit != *other.Limit {
			return false
		}
	}
	if resource.Spss == nil && other.Spss != nil || resource.Spss != nil && other.Spss == nil {
		return false
	}

	if resource.Spss != nil {
		if *resource.Spss != *other.Spss {
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

	if len(resource.GroupBy) != len(other.GroupBy) {
		return false
	}

	for i1 := range resource.GroupBy {
		if !resource.GroupBy[i1].Equals(other.GroupBy[i1]) {
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
	if resource.TableType == nil && other.TableType != nil || resource.TableType != nil && other.TableType == nil {
		return false
	}

	if resource.TableType != nil {
		if *resource.TableType != *other.TableType {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `TempoQuery` fields for violations and returns them.
func (resource TempoQuery) Validate() error {
	var errs cog.BuildErrors

	for i1 := range resource.Filters {
		if err := resource.Filters[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("filters["+strconv.Itoa(i1)+"]", err)...)
		}
	}

	for i1 := range resource.GroupBy {
		if err := resource.GroupBy[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("groupBy["+strconv.Itoa(i1)+"]", err)...)
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

// search = Loki search, nativeSearch = Tempo search for backwards compatibility
type TempoQueryType string

const (
	TempoQueryTypeTraceql       TempoQueryType = "traceql"
	TempoQueryTypeTraceqlSearch TempoQueryType = "traceqlSearch"
	TempoQueryTypeSearch        TempoQueryType = "search"
	TempoQueryTypeServiceMap    TempoQueryType = "serviceMap"
	TempoQueryTypeUpload        TempoQueryType = "upload"
	TempoQueryTypeNativeSearch  TempoQueryType = "nativeSearch"
	TempoQueryTypeTraceId       TempoQueryType = "traceId"
	TempoQueryTypeClear         TempoQueryType = "clear"
)

// The state of the TraceQL streaming search query
type SearchStreamingState string

const (
	SearchStreamingStatePending   SearchStreamingState = "pending"
	SearchStreamingStateStreaming SearchStreamingState = "streaming"
	SearchStreamingStateDone      SearchStreamingState = "done"
	SearchStreamingStateError     SearchStreamingState = "error"
)

// The type of the table that is used to display the search results
type SearchTableType string

const (
	SearchTableTypeTraces SearchTableType = "traces"
	SearchTableTypeSpans  SearchTableType = "spans"
)

// static fields are pre-set in the UI, dynamic fields are added by the user
type TraceqlSearchScope string

const (
	TraceqlSearchScopeIntrinsic TraceqlSearchScope = "intrinsic"
	TraceqlSearchScopeUnscoped  TraceqlSearchScope = "unscoped"
	TraceqlSearchScopeResource  TraceqlSearchScope = "resource"
	TraceqlSearchScopeSpan      TraceqlSearchScope = "span"
)

type TraceqlFilter struct {
	// Uniquely identify the filter, will not be used in the query generation
	Id string `json:"id"`
	// The tag for the search filter, for example: .http.status_code, .service.name, status
	Tag *string `json:"tag,omitempty"`
	// The operator that connects the tag to the value, for example: =, >, !=, =~
	Operator *string `json:"operator,omitempty"`
	// The value for the search filter
	Value *StringOrArrayOfString `json:"value,omitempty"`
	// The type of the value, used for example to check whether we need to wrap the value in quotes when generating the query
	ValueType *string `json:"valueType,omitempty"`
	// The scope of the filter, can either be unscoped/all scopes, resource or span
	Scope *TraceqlSearchScope `json:"scope,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TraceqlFilter` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *TraceqlFilter) UnmarshalJSONStrict(raw []byte) error {
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
	// Field "tag"
	if fields["tag"] != nil {
		if string(fields["tag"]) != "null" {
			if err := json.Unmarshal(fields["tag"], &resource.Tag); err != nil {
				errs = append(errs, cog.MakeBuildErrors("tag", err)...)
			}

		}
		delete(fields, "tag")

	}
	// Field "operator"
	if fields["operator"] != nil {
		if string(fields["operator"]) != "null" {
			if err := json.Unmarshal(fields["operator"], &resource.Operator); err != nil {
				errs = append(errs, cog.MakeBuildErrors("operator", err)...)
			}

		}
		delete(fields, "operator")

	}
	// Field "value"
	if fields["value"] != nil {
		if string(fields["value"]) != "null" {

			resource.Value = &StringOrArrayOfString{}
			if err := resource.Value.UnmarshalJSONStrict(fields["value"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("value", err)...)
			}

		}
		delete(fields, "value")

	}
	// Field "valueType"
	if fields["valueType"] != nil {
		if string(fields["valueType"]) != "null" {
			if err := json.Unmarshal(fields["valueType"], &resource.ValueType); err != nil {
				errs = append(errs, cog.MakeBuildErrors("valueType", err)...)
			}

		}
		delete(fields, "valueType")

	}
	// Field "scope"
	if fields["scope"] != nil {
		if string(fields["scope"]) != "null" {
			if err := json.Unmarshal(fields["scope"], &resource.Scope); err != nil {
				errs = append(errs, cog.MakeBuildErrors("scope", err)...)
			}

		}
		delete(fields, "scope")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("TraceqlFilter", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `TraceqlFilter` objects.
func (resource TraceqlFilter) Equals(other TraceqlFilter) bool {
	if resource.Id != other.Id {
		return false
	}
	if resource.Tag == nil && other.Tag != nil || resource.Tag != nil && other.Tag == nil {
		return false
	}

	if resource.Tag != nil {
		if *resource.Tag != *other.Tag {
			return false
		}
	}
	if resource.Operator == nil && other.Operator != nil || resource.Operator != nil && other.Operator == nil {
		return false
	}

	if resource.Operator != nil {
		if *resource.Operator != *other.Operator {
			return false
		}
	}
	if resource.Value == nil && other.Value != nil || resource.Value != nil && other.Value == nil {
		return false
	}

	if resource.Value != nil {
		if !resource.Value.Equals(*other.Value) {
			return false
		}
	}
	if resource.ValueType == nil && other.ValueType != nil || resource.ValueType != nil && other.ValueType == nil {
		return false
	}

	if resource.ValueType != nil {
		if *resource.ValueType != *other.ValueType {
			return false
		}
	}
	if resource.Scope == nil && other.Scope != nil || resource.Scope != nil && other.Scope == nil {
		return false
	}

	if resource.Scope != nil {
		if *resource.Scope != *other.Scope {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `TraceqlFilter` fields for violations and returns them.
func (resource TraceqlFilter) Validate() error {
	var errs cog.BuildErrors
	if resource.Value != nil {
		if err := resource.Value.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("value", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type StringOrArrayOfString struct {
	String        *string  `json:"String,omitempty"`
	ArrayOfString []string `json:"ArrayOfString,omitempty"`
}

// MarshalJSON implements a custom JSON marshalling logic to encode `StringOrArrayOfString` as JSON.
func (resource StringOrArrayOfString) MarshalJSON() ([]byte, error) {
	if resource.String != nil {
		return json.Marshal(resource.String)
	}

	if resource.ArrayOfString != nil {
		return json.Marshal(resource.ArrayOfString)
	}

	return nil, fmt.Errorf("no value for disjunction of scalars")
}

// UnmarshalJSON implements a custom JSON unmarshalling logic to decode `StringOrArrayOfString` from JSON.
func (resource *StringOrArrayOfString) UnmarshalJSON(raw []byte) error {
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

	// ArrayOfString
	var ArrayOfString []string
	if err := json.Unmarshal(raw, &ArrayOfString); err != nil {
		errList = append(errList, err)
		resource.ArrayOfString = nil
	} else {
		resource.ArrayOfString = ArrayOfString
		return nil
	}

	return errors.Join(errList...)
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `StringOrArrayOfString` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *StringOrArrayOfString) UnmarshalJSONStrict(raw []byte) error {
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

	// ArrayOfString
	var ArrayOfString []string

	if err := json.Unmarshal(raw, &ArrayOfString); err != nil {
		errList = append(errList, err)
	} else {
		resource.ArrayOfString = ArrayOfString
		return nil
	}

	if len(errList) != 0 {
		errs = append(errs, cog.MakeBuildErrors("StringOrArrayOfString", errors.Join(errList...))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `StringOrArrayOfString` objects.
func (resource StringOrArrayOfString) Equals(other StringOrArrayOfString) bool {
	if resource.String == nil && other.String != nil || resource.String != nil && other.String == nil {
		return false
	}

	if resource.String != nil {
		if *resource.String != *other.String {
			return false
		}
	}

	if len(resource.ArrayOfString) != len(other.ArrayOfString) {
		return false
	}

	for i1 := range resource.ArrayOfString {
		if resource.ArrayOfString[i1] != other.ArrayOfString[i1] {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `StringOrArrayOfString` fields for violations and returns them.
func (resource StringOrArrayOfString) Validate() error {
	return nil
}
