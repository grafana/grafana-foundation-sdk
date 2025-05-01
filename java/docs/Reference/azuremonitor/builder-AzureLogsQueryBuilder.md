---
title: <span class="badge builder"></span> AzureLogsQueryBuilder
---
# <span class="badge builder"></span> AzureLogsQueryBuilder

## Constructor

```java
new AzureLogsQueryBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public AzureLogsQuery build()
```

### <span class="badge object-method"></span> basicLogsQuery

If set to true the query will be run as a basic logs query

```java
public AzureLogsQueryBuilder basicLogsQuery(Boolean basicLogsQuery)
```

### <span class="badge object-method"></span> dashboardTime

If set to true the dashboard time range will be used as a filter for the query. Otherwise the query time ranges will be used. Defaults to false.

```java
public AzureLogsQueryBuilder dashboardTime(Boolean dashboardTime)
```

### <span class="badge object-method"></span> intersectTime

@deprecated Use dashboardTime instead

```java
public AzureLogsQueryBuilder intersectTime(Boolean intersectTime)
```

### <span class="badge object-method"></span> query

KQL query to be executed.

```java
public AzureLogsQueryBuilder query(String query)
```

### <span class="badge object-method"></span> resource

@deprecated Use resources instead

```java
public AzureLogsQueryBuilder resource(String resource)
```

### <span class="badge object-method"></span> resources

Array of resource URIs to be queried.

```java
public AzureLogsQueryBuilder resources(List<String> resources)
```

### <span class="badge object-method"></span> resultFormat

Specifies the format results should be returned as.

```java
public AzureLogsQueryBuilder resultFormat(ResultFormat resultFormat)
```

### <span class="badge object-method"></span> timeColumn

If dashboardTime is set to true this value dictates which column the time filter will be applied to. Defaults to the first tables timeSpan column, the first datetime column found, or TimeGenerated

```java
public AzureLogsQueryBuilder timeColumn(String timeColumn)
```

### <span class="badge object-method"></span> workspace

Workspace ID. This was removed in Grafana 8, but remains for backwards compat.

```java
public AzureLogsQueryBuilder workspace(String workspace)
```

## See also

 * <span class="badge object-type-class"></span> [AzureLogsQuery](./object-AzureLogsQuery.md)
