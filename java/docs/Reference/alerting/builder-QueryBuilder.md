---
title: <span class="badge builder"></span> QueryBuilder
---
# <span class="badge builder"></span> QueryBuilder

## Constructor

```java
new QueryBuilder(String refId)
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public Query build()
```

### <span class="badge object-method"></span> datasourceUid

Grafana data source unique identifier; it should be '__expr__' for a Server Side Expression operation.

```java
public QueryBuilder datasourceUid(String datasourceUid)
```

### <span class="badge object-method"></span> model

JSON is the raw JSON query and includes the above properties as well as custom properties.

```java
public QueryBuilder model(com.grafana.foundation.cog.Builder<Dataquery> model)
```

### <span class="badge object-method"></span> queryType

QueryType is an optional identifier for the type of query.

It can be used to distinguish different types of queries.

```java
public QueryBuilder queryType(String queryType)
```

### <span class="badge object-method"></span> refId

RefID is the unique identifier of the query, set by the frontend call.

```java
public QueryBuilder refId(String refId)
```

### <span class="badge object-method"></span> relativeTimeRange

RelativeTimeRange is the per query start and end time

for requests.

```java
public QueryBuilder relativeTimeRange(RelativeTimeRange relativeTimeRange)
```

## See also

 * <span class="badge object-type-class"></span> [Query](./object-Query.md)
