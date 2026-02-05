---
title: <span class="badge builder"></span> AzureTracesQueryBuilder
---
# <span class="badge builder"></span> AzureTracesQueryBuilder

## Constructor

```java
new AzureTracesQueryBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public AzureTracesQuery build()
```

### <span class="badge object-method"></span> filters

Filters for property values.

```java
public AzureTracesQueryBuilder filters(List<com.grafana.foundation.cog.Builder<AzureTracesFilter>> filters)
```

### <span class="badge object-method"></span> operationId

Operation ID. Used only for Traces queries.

```java
public AzureTracesQueryBuilder operationId(String operationId)
```

### <span class="badge object-method"></span> query

KQL query to be executed.

```java
public AzureTracesQueryBuilder query(String query)
```

### <span class="badge object-method"></span> resources

Array of resource URIs to be queried.

```java
public AzureTracesQueryBuilder resources(List<String> resources)
```

### <span class="badge object-method"></span> resultFormat

Specifies the format results should be returned as.

```java
public AzureTracesQueryBuilder resultFormat(ResultFormat resultFormat)
```

### <span class="badge object-method"></span> traceTypes

Types of events to filter by.

```java
public AzureTracesQueryBuilder traceTypes(List<String> traceTypes)
```

## See also

 * <span class="badge object-type-class"></span> [AzureTracesQuery](./object-AzureTracesQuery.md)
