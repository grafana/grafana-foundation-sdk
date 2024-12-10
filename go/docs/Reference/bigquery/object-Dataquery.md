---
title: <span class="badge object-type-struct"></span> Dataquery
---
# <span class="badge object-type-struct"></span> Dataquery

Manually converted from https://github.com/grafana/google-bigquery-datasource/blob/18680e42ba557791d109c7c540c2c3f2647592f0/src/types.ts#L75

## Definition

```go
type Dataquery struct {
    Dataset *string `json:"dataset,omitempty"`
    Table *string `json:"table,omitempty"`
    Project *string `json:"project,omitempty"`
    Format bigquery.QueryFormat `json:"format"`
    RawQuery *bool `json:"rawQuery,omitempty"`
    RawSql string `json:"rawSql"`
    Location *string `json:"location,omitempty"`
    Partitioned *bool `json:"partitioned,omitempty"`
    PartitionedField *string `json:"partitionedField,omitempty"`
    ConvertToUTC *bool `json:"convertToUTC,omitempty"`
    Sharded *bool `json:"sharded,omitempty"`
    QueryPriority *bigquery.QueryPriority `json:"queryPriority,omitempty"`
    TimeShift *string `json:"timeShift,omitempty"`
    EditorMode *bigquery.EditorMode `json:"editorMode,omitempty"`
    Sql *bigquery.SQLExpression `json:"sql,omitempty"`
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