---
title: <span class="badge builder"></span> TracesQueryBuilder
---
# <span class="badge builder"></span> TracesQueryBuilder

Application Insights Traces sub-query properties

## Constructor

```java
new TracesQueryBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public TracesQuery build()
```

### <span class="badge object-method"></span> filters

Filters for property values.

```java
public TracesQueryBuilder filters(List<com.grafana.foundation.cog.Builder<TracesFilter>> filters)
```

### <span class="badge object-method"></span> operationId

Operation ID. Used only for Traces queries.

```java
public TracesQueryBuilder operationId(String operationId)
```

### <span class="badge object-method"></span> query

KQL query to be executed.

```java
public TracesQueryBuilder query(String query)
```

### <span class="badge object-method"></span> resources

Array of resource URIs to be queried.

```java
public TracesQueryBuilder resources(List<String> resources)
```

### <span class="badge object-method"></span> resultFormat

Specifies the format results should be returned as.

```java
public TracesQueryBuilder resultFormat(ResultFormat resultFormat)
```

### <span class="badge object-method"></span> traceTypes

Types of events to filter by.

```java
public TracesQueryBuilder traceTypes(List<String> traceTypes)
```

## See also

 * <span class="badge object-type-class"></span> [TracesQuery](./object-TracesQuery.md)
