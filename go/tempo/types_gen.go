// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package tempo

import (
	"encoding/json"
	"errors"
	"fmt"

	cogvariants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
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
	Datasource any `json:"datasource,omitempty"`
	// The type of the table that is used to display the search results
	TableType *SearchTableType `json:"tableType,omitempty"`
}

func (resource TempoQuery) ImplementsDataqueryVariant() {}

func VariantConfig() cogvariants.DataqueryConfig {
	return cogvariants.DataqueryConfig{
		Identifier: "tempo",
		DataqueryUnmarshaler: func(raw []byte) (cogvariants.Dataquery, error) {
			dataquery := TempoQuery{}

			if err := json.Unmarshal(raw, &dataquery); err != nil {
				return nil, err
			}

			return dataquery, nil
		},
	}
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

type StringOrArrayOfString struct {
	String        *string  `json:"String,omitempty"`
	ArrayOfString []string `json:"ArrayOfString,omitempty"`
}

func (resource StringOrArrayOfString) MarshalJSON() ([]byte, error) {
	if resource.String != nil {
		return json.Marshal(resource.String)
	}

	if resource.ArrayOfString != nil {
		return json.Marshal(resource.ArrayOfString)
	}

	return nil, fmt.Errorf("no value for disjunction of scalars")
}

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
