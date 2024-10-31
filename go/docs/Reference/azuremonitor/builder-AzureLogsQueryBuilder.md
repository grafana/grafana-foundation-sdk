---
title: <span class="badge builder"></span> AzureLogsQueryBuilder
---
# <span class="badge builder"></span> AzureLogsQueryBuilder

## Constructor

```go
func NewAzureLogsQueryBuilder() *AzureLogsQueryBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *AzureLogsQueryBuilder) Build() (AzureLogsQuery, error)
```

### <span class="badge object-method"></span> IntersectTime

If set to true the intersection of time ranges specified in the query and Grafana will be used. Otherwise the query time ranges will be used. Defaults to false

```go
func (builder *AzureLogsQueryBuilder) IntersectTime(intersectTime bool) *AzureLogsQueryBuilder
```

### <span class="badge object-method"></span> Query

KQL query to be executed.

```go
func (builder *AzureLogsQueryBuilder) Query(query string) *AzureLogsQueryBuilder
```

### <span class="badge object-method"></span> Resource

@deprecated Use resources instead

```go
func (builder *AzureLogsQueryBuilder) Resource(resource string) *AzureLogsQueryBuilder
```

### <span class="badge object-method"></span> Resources

Array of resource URIs to be queried.

```go
func (builder *AzureLogsQueryBuilder) Resources(resources []string) *AzureLogsQueryBuilder
```

### <span class="badge object-method"></span> ResultFormat

Specifies the format results should be returned as.

```go
func (builder *AzureLogsQueryBuilder) ResultFormat(resultFormat azuremonitor.ResultFormat) *AzureLogsQueryBuilder
```

### <span class="badge object-method"></span> Workspace

Workspace ID. This was removed in Grafana 8, but remains for backwards compat

```go
func (builder *AzureLogsQueryBuilder) Workspace(workspace string) *AzureLogsQueryBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [AzureLogsQuery](./object-AzureLogsQuery.md)
