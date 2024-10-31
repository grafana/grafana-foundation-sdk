---
title: <span class="badge object-type-struct"></span> Dataquery
---
# <span class="badge object-type-struct"></span> Dataquery

## Definition

```go
type Dataquery struct {
    // The LogQL query.
    Expr string `json:"expr"`
    // Used to override the name of the series.
    LegendFormat *string `json:"legendFormat,omitempty"`
    // Used to limit the number of log rows returned.
    MaxLines *int64 `json:"maxLines,omitempty"`
    // @deprecated, now use step.
    Resolution *int64 `json:"resolution,omitempty"`
    EditorMode *loki.QueryEditorMode `json:"editorMode,omitempty"`
    // @deprecated, now use queryType.
    Range *bool `json:"range,omitempty"`
    // @deprecated, now use queryType.
    Instant *bool `json:"instant,omitempty"`
    // Used to set step value for range queries.
    Step *string `json:"step,omitempty"`
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
