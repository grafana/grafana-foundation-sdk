---
title: <span class="badge object-type-struct"></span> Dataquery
---
# <span class="badge object-type-struct"></span> Dataquery

## Definition

```go
type Dataquery struct {
    // The actual expression/query that will be evaluated by Prometheus
    Expr string `json:"expr"`
    // Returns only the latest value that Prometheus has scraped for the requested time series
    Instant *bool `json:"instant,omitempty"`
    // Returns a Range vector, comprised of a set of time series containing a range of data points over time for each time series
    Range *bool `json:"range,omitempty"`
    // Execute an additional query to identify interesting raw samples relevant for the given expr
    Exemplar *bool `json:"exemplar,omitempty"`
    // Specifies which editor is being used to prepare the query. It can be "code" or "builder"
    EditorMode *prometheus.QueryEditorMode `json:"editorMode,omitempty"`
    // Query format to determine how to display data points in panel. It can be "time_series", "table", "heatmap"
    Format *prometheus.PromQueryFormat `json:"format,omitempty"`
    // Series name override or template. Ex. {{hostname}} will be replaced with label value for hostname
    LegendFormat *string `json:"legendFormat,omitempty"`
    // @deprecated Used to specify how many times to divide max data points by. We use max data points under query options
    // See https://github.com/grafana/grafana/issues/48081
    IntervalFactor *float64 `json:"intervalFactor,omitempty"`
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
