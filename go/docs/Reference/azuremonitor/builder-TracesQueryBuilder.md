---
title: <span class="badge builder"></span> TracesQueryBuilder
---
# <span class="badge builder"></span> TracesQueryBuilder

## Constructor

```go
func NewTracesQueryBuilder() *TracesQueryBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *TracesQueryBuilder) Build() (TracesQuery, error)
```

### <span class="badge object-method"></span> Filters

Filters for property values.

```go
func (builder *TracesQueryBuilder) Filters(filters []cog.Builder[azuremonitor.TracesFilter]) *TracesQueryBuilder
```

### <span class="badge object-method"></span> OperationId

Operation ID. Used only for Traces queries.

```go
func (builder *TracesQueryBuilder) OperationId(operationId string) *TracesQueryBuilder
```

### <span class="badge object-method"></span> Query

KQL query to be executed.

```go
func (builder *TracesQueryBuilder) Query(query string) *TracesQueryBuilder
```

### <span class="badge object-method"></span> Resources

Array of resource URIs to be queried.

```go
func (builder *TracesQueryBuilder) Resources(resources []string) *TracesQueryBuilder
```

### <span class="badge object-method"></span> ResultFormat

Specifies the format results should be returned as.

```go
func (builder *TracesQueryBuilder) ResultFormat(resultFormat azuremonitor.ResultFormat) *TracesQueryBuilder
```

### <span class="badge object-method"></span> TraceTypes

Types of events to filter by.

```go
func (builder *TracesQueryBuilder) TraceTypes(traceTypes []string) *TracesQueryBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [TracesQuery](./object-TracesQuery.md)
