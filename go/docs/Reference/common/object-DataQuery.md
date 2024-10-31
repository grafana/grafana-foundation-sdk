---
title: <span class="badge object-type-struct"></span> DataQuery
---
# <span class="badge object-type-struct"></span> DataQuery

These are the common properties available to all queries in all datasources.

Specific implementations will *extend* this interface, adding the required

properties for the given context.

## Definition

```go
type DataQuery struct {
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
    Datasource any `json:"datasource,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `DataQuery` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (dataQuery *DataQuery) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `DataQuery` objects.

```go
func (dataQuery *DataQuery) Equals(other DataQuery) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `DataQuery` fields for violations and returns them.

```go
func (dataQuery *DataQuery) Validate() error
```

## See also

 * <span class="badge builder"></span> [DataQueryBuilder](./builder-DataQueryBuilder.md)
