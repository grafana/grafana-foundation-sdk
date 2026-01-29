---
title: <span class="badge builder"></span> QueryBuilder
---
# <span class="badge builder"></span> QueryBuilder

## Constructor

```typescript
new QueryBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> convertToUTC

```typescript
convertToUTC(convertToUTC: boolean)
```

### <span class="badge object-method"></span> dataset

```typescript
dataset(dataset: string)
```

### <span class="badge object-method"></span> datasource

New type for datasource reference

Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.

```typescript
datasource(ref: {
	name?: string;
})
```

### <span class="badge object-method"></span> editorMode

```typescript
editorMode(editorMode: bigquery.EditorMode)
```

### <span class="badge object-method"></span> format

```typescript
format(format: bigquery.QueryFormat)
```

### <span class="badge object-method"></span> hide

If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.

```typescript
hide(hide: boolean)
```

### <span class="badge object-method"></span> location

```typescript
location(location: string)
```

### <span class="badge object-method"></span> oldDatasource

For mixed data sources the selected datasource is on the query level.

For non mixed scenarios this is undefined.

TODO find a better way to do this ^ that's friendly to schema

TODO this shouldn't be unknown but DataSourceRef | null

```typescript
oldDatasource(datasource: common.DataSourceRef)
```

### <span class="badge object-method"></span> partitioned

```typescript
partitioned(partitioned: boolean)
```

### <span class="badge object-method"></span> partitionedField

```typescript
partitionedField(partitionedField: string)
```

### <span class="badge object-method"></span> project

```typescript
project(project: string)
```

### <span class="badge object-method"></span> queryPriority

```typescript
queryPriority(queryPriority: bigquery.QueryPriority)
```

### <span class="badge object-method"></span> queryType

Specify the query flavor

TODO make this required and give it a default

```typescript
queryType(queryType: string)
```

### <span class="badge object-method"></span> rawQuery

```typescript
rawQuery(rawQuery: boolean)
```

### <span class="badge object-method"></span> rawSql

```typescript
rawSql(rawSql: string)
```

### <span class="badge object-method"></span> refId

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```typescript
refId(refId: string)
```

### <span class="badge object-method"></span> sharded

```typescript
sharded(sharded: boolean)
```

### <span class="badge object-method"></span> sql

```typescript
sql(sql: cog.Builder<bigquery.SQLExpression>)
```

### <span class="badge object-method"></span> table

```typescript
table(table: string)
```

### <span class="badge object-method"></span> timeShift

```typescript
timeShift(timeShift: string)
```

### <span class="badge object-method"></span> version

```typescript
version(version: string)
```

## See also

 * <span class="badge object-type-interface"></span> [dashboardv2beta1.DataQueryKind](../dashboardv2beta1/object-DataQueryKind.md)
