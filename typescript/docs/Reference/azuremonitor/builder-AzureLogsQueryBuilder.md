---
title: <span class="badge builder"></span> AzureLogsQueryBuilder
---
# <span class="badge builder"></span> AzureLogsQueryBuilder

## Constructor

```typescript
new AzureLogsQueryBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> intersectTime

If set to true the intersection of time ranges specified in the query and Grafana will be used. Otherwise the query time ranges will be used. Defaults to false

```typescript
intersectTime(intersectTime: boolean)
```

### <span class="badge object-method"></span> query

KQL query to be executed.

```typescript
query(query: string)
```

### <span class="badge object-method"></span> resource

@deprecated Use resources instead

```typescript
resource(resource: string)
```

### <span class="badge object-method"></span> resources

Array of resource URIs to be queried.

```typescript
resources(resources: string[])
```

### <span class="badge object-method"></span> resultFormat

Specifies the format results should be returned as.

```typescript
resultFormat(resultFormat: azuremonitor.ResultFormat)
```

### <span class="badge object-method"></span> workspace

Workspace ID. This was removed in Grafana 8, but remains for backwards compat

```typescript
workspace(workspace: string)
```

## See also

 * <span class="badge object-type-interface"></span> [AzureLogsQuery](./object-AzureLogsQuery.md)
