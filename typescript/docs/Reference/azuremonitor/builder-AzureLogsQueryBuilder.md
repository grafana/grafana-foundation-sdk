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

### <span class="badge object-method"></span> basicLogsQuery

If set to true the query will be run as a basic logs query

```typescript
basicLogsQuery(basicLogsQuery: boolean)
```

### <span class="badge object-method"></span> dashboardTime

If set to true the dashboard time range will be used as a filter for the query. Otherwise the query time ranges will be used. Defaults to false.

```typescript
dashboardTime(dashboardTime: boolean)
```

### <span class="badge object-method"></span> intersectTime

@deprecated Use dashboardTime instead

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

### <span class="badge object-method"></span> timeColumn

If dashboardTime is set to true this value dictates which column the time filter will be applied to. Defaults to the first tables timeSpan column, the first datetime column found, or TimeGenerated

```typescript
timeColumn(timeColumn: string)
```

### <span class="badge object-method"></span> workspace

Workspace ID. This was removed in Grafana 8, but remains for backwards compat.

```typescript
workspace(workspace: string)
```

## See also

 * <span class="badge object-type-interface"></span> [AzureLogsQuery](./object-AzureLogsQuery.md)
