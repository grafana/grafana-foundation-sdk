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

### <span class="badge object-method"></span> datasource

New type for datasource reference

Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.

```typescript
datasource(ref: {
	name?: string;
})
```

### <span class="badge object-method"></span> direction

```typescript
direction(direction: loki.LokiQueryDirection)
```

### <span class="badge object-method"></span> editorMode

```typescript
editorMode(editorMode: loki.QueryEditorMode)
```

### <span class="badge object-method"></span> expr

The LogQL query.

```typescript
expr(expr: string)
```

### <span class="badge object-method"></span> hide

If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.

```typescript
hide(hide: boolean)
```

### <span class="badge object-method"></span> instant

@deprecated, now use queryType.

```typescript
instant(instant: boolean)
```

### <span class="badge object-method"></span> legendFormat

Used to override the name of the series.

```typescript
legendFormat(legendFormat: string)
```

### <span class="badge object-method"></span> maxLines

Used to limit the number of log rows returned.

```typescript
maxLines(maxLines: number)
```

### <span class="badge object-method"></span> queryType

Specify the query flavor

TODO make this required and give it a default

```typescript
queryType(queryType: string)
```

### <span class="badge object-method"></span> range

@deprecated, now use queryType.

```typescript
range(range: boolean)
```

### <span class="badge object-method"></span> refId

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```typescript
refId(refId: string)
```

### <span class="badge object-method"></span> resolution

@deprecated, now use step.

```typescript
resolution(resolution: number)
```

### <span class="badge object-method"></span> step

Used to set step value for range queries.

```typescript
step(step: string)
```

### <span class="badge object-method"></span> version

```typescript
version(version: string)
```

## See also

 * <span class="badge object-type-interface"></span> [dashboardv2beta1.DataQueryKind](../dashboardv2beta1/object-DataQueryKind.md)
