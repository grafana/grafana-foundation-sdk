---
title: <span class="badge object-type-struct"></span> Dataquery
---
# <span class="badge object-type-struct"></span> Dataquery

## Definition

```go
type Dataquery struct {
    // Additional Ad-hoc filters that take precedence over Scope on conflict.
    AdhocFilters []prometheus.AdhocFilters `json:"adhocFilters,omitempty"`
    // The datasource
    Datasource *common.DataSourceRef `json:"datasource,omitempty"`
    // what we should show in the editor
    // Possible enum values:
    //  - `"builder"` 
    //  - `"code"` 
    EditorMode *prometheus.QueryEditorMode `json:"editorMode,omitempty"`
    // Execute an additional query to identify interesting raw samples relevant for the given expr
    Exemplar *bool `json:"exemplar,omitempty"`
    // The actual expression/query that will be evaluated by Prometheus
    Expr string `json:"expr"`
    // The response format
    // Possible enum values:
    //  - `"time_series"` 
    //  - `"table"` 
    //  - `"heatmap"` 
    Format *prometheus.PromQueryFormat `json:"format,omitempty"`
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
    ResultAssertions *prometheus.ResultAssertions `json:"resultAssertions,omitempty"`
    // A set of filters applied to apply to the query
    Scopes []prometheus.Scopes `json:"scopes,omitempty"`
    // TimeRange represents the query range
    // NOTE: unlike generic /ds/query, we can now send explicit time values in each query
    // NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
    TimeRange *prometheus.TimeRange `json:"timeRange,omitempty"`
    // An additional lower limit for the step parameter of the Prometheus query and for the
    // `$__interval` and `$__rate_interval` variables.
    Interval *string `json:"interval,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Dataquery` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (dataquery *Dataquery) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `Dataquery` objects.

```go
func (dataquery *Dataquery) Equals(other Dataquery) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `Dataquery` fields for violations and returns them.

```go
func (dataquery *Dataquery) Validate() error
```

## See also

 * <span class="badge builder"></span> [DataqueryBuilder](./builder-DataqueryBuilder.md)
