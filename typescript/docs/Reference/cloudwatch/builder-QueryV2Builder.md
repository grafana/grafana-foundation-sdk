---
title: <span class="badge builder"></span> QueryV2Builder
---
# <span class="badge builder"></span> QueryV2Builder

## Constructor

```typescript
new QueryV2Builder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> cloudWatchAnnotationQuery

```typescript
cloudWatchAnnotationQuery(cloudWatchAnnotationQuery: cog.Builder<cloudwatch.CloudWatchAnnotationQuery>)
```

### <span class="badge object-method"></span> cloudWatchLogsQuery

```typescript
cloudWatchLogsQuery(cloudWatchLogsQuery: cog.Builder<cloudwatch.CloudWatchLogsQuery>)
```

### <span class="badge object-method"></span> cloudWatchMetricsQuery

```typescript
cloudWatchMetricsQuery(cloudWatchMetricsQuery: cog.Builder<cloudwatch.CloudWatchMetricsQuery>)
```

### <span class="badge object-method"></span> datasource

New type for datasource reference

Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.

```typescript
datasource(ref: {
	name?: string;
})
```

### <span class="badge object-method"></span> labels

```typescript
labels(labels: Record<string, string>)
```

### <span class="badge object-method"></span> version

```typescript
version(version: string)
```

## See also

 * <span class="badge object-type-interface"></span> [dashboardv2.DataQueryKind](../dashboardv2/object-DataQueryKind.md)
