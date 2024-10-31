---
title: <span class="badge builder"></span> QueryBuilder
---
# <span class="badge builder"></span> QueryBuilder

## Constructor

```typescript
new QueryBuilder(refId: string)
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> datasourceUid

Grafana data source unique identifier; it should be '__expr__' for a Server Side Expression operation.

```typescript
datasourceUid(datasourceUid: string)
```

### <span class="badge object-method"></span> model

JSON is the raw JSON query and includes the above properties as well as custom properties.

```typescript
model(model: cog.Builder<cog.Dataquery>)
```

### <span class="badge object-method"></span> queryType

QueryType is an optional identifier for the type of query.

It can be used to distinguish different types of queries.

```typescript
queryType(queryType: string)
```

### <span class="badge object-method"></span> refId

RefID is the unique identifier of the query, set by the frontend call.

```typescript
refId(refId: string)
```

### <span class="badge object-method"></span> relativeTimeRange

RelativeTimeRange is the per query start and end time

for requests.

```typescript
relativeTimeRange(relativeTimeRange: alerting.RelativeTimeRange)
```

## See also

 * <span class="badge object-type-interface"></span> [Query](./object-Query.md)
