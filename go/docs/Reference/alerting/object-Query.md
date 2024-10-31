---
title: <span class="badge object-type-struct"></span> Query
---
# <span class="badge object-type-struct"></span> Query

## Definition

```go
type Query struct {
    // Grafana data source unique identifier; it should be '__expr__' for a Server Side Expression operation.
    DatasourceUid *string `json:"datasourceUid,omitempty"`
    // JSON is the raw JSON query and includes the above properties as well as custom properties.
    Model cog/variants.Dataquery `json:"model,omitempty"`
    // QueryType is an optional identifier for the type of query.
    // It can be used to distinguish different types of queries.
    QueryType *string `json:"queryType,omitempty"`
    // RefID is the unique identifier of the query, set by the frontend call.
    RefId *string `json:"refId,omitempty"`
    // RelativeTimeRange is the per query start and end time
    // for requests.
    RelativeTimeRange *alerting.RelativeTimeRange `json:"relativeTimeRange,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSON

UnmarshalJSON implements a custom JSON unmarshalling logic to decode `Query` from JSON.

```go
func (query *Query) UnmarshalJSON(raw []byte) error
```

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Query` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (query *Query) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `Query` objects.

```go
func (query *Query) Equals(other Query) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `Query` fields for violations and returns them.

```go
func (query *Query) Validate() error
```

## See also

 * <span class="badge builder"></span> [QueryBuilder](./builder-QueryBuilder.md)
