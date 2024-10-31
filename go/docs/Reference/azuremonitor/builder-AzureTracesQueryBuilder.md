---
title: <span class="badge builder"></span> AzureTracesQueryBuilder
---
# <span class="badge builder"></span> AzureTracesQueryBuilder

## Constructor

```go
func NewAzureTracesQueryBuilder() *AzureTracesQueryBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *AzureTracesQueryBuilder) Build() (AzureTracesQuery, error)
```

### <span class="badge object-method"></span> Filters

Filters for property values.

```go
func (builder *AzureTracesQueryBuilder) Filters(filters []cog.Builder[azuremonitor.AzureTracesFilter]) *AzureTracesQueryBuilder
```

### <span class="badge object-method"></span> OperationId

Operation ID. Used only for Traces queries.

```go
func (builder *AzureTracesQueryBuilder) OperationId(operationId string) *AzureTracesQueryBuilder
```

### <span class="badge object-method"></span> Query

KQL query to be executed.

```go
func (builder *AzureTracesQueryBuilder) Query(query string) *AzureTracesQueryBuilder
```

### <span class="badge object-method"></span> Resources

Array of resource URIs to be queried.

```go
func (builder *AzureTracesQueryBuilder) Resources(resources []string) *AzureTracesQueryBuilder
```

### <span class="badge object-method"></span> ResultFormat

Specifies the format results should be returned as.

```go
func (builder *AzureTracesQueryBuilder) ResultFormat(resultFormat azuremonitor.ResultFormat) *AzureTracesQueryBuilder
```

### <span class="badge object-method"></span> TraceTypes

Types of events to filter by.

```go
func (builder *AzureTracesQueryBuilder) TraceTypes(traceTypes []string) *AzureTracesQueryBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [AzureTracesQuery](./object-AzureTracesQuery.md)
