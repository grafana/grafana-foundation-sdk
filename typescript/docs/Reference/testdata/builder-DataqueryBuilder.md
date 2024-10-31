---
title: <span class="badge builder"></span> DataqueryBuilder
---
# <span class="badge builder"></span> DataqueryBuilder

## Constructor

```typescript
new DataqueryBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> alias

```typescript
alias(alias: string)
```

### <span class="badge object-method"></span> channel

```typescript
channel(channel: string)
```

### <span class="badge object-method"></span> csvContent

```typescript
csvContent(csvContent: string)
```

### <span class="badge object-method"></span> csvFileName

```typescript
csvFileName(csvFileName: string)
```

### <span class="badge object-method"></span> csvWave

```typescript
csvWave(csvWave: cog.Builder<testdata.CSVWave>[])
```

### <span class="badge object-method"></span> datasource

For mixed data sources the selected datasource is on the query level.

For non mixed scenarios this is undefined.

TODO find a better way to do this ^ that's friendly to schema

TODO this shouldn't be unknown but DataSourceRef | null

```typescript
datasource(datasource: dashboard.DataSourceRef)
```

### <span class="badge object-method"></span> dropPercent

Drop percentage (the chance we will lose a point 0-100)

```typescript
dropPercent(dropPercent: number)
```

### <span class="badge object-method"></span> errorType

```typescript
errorType(errorType: "server_panic" | "frontend_exception" | "frontend_observable")
```

### <span class="badge object-method"></span> hide

true if query is disabled (ie should not be returned to the dashboard)

Note this does not always imply that the query should not be executed since

the results from a hidden query may be used as the input to other queries (SSE etc)

```typescript
hide(hide: boolean)
```

### <span class="badge object-method"></span> labels

```typescript
labels(labels: string)
```

### <span class="badge object-method"></span> levelColumn

```typescript
levelColumn(levelColumn: boolean)
```

### <span class="badge object-method"></span> lines

```typescript
lines(lines: number)
```

### <span class="badge object-method"></span> nodes

```typescript
nodes(nodes: cog.Builder<testdata.NodesQuery>)
```

### <span class="badge object-method"></span> points

```typescript
points(points: (string | number)[][])
```

### <span class="badge object-method"></span> pulseWave

```typescript
pulseWave(pulseWave: cog.Builder<testdata.PulseWaveQuery>)
```

### <span class="badge object-method"></span> queryType

Specify the query flavor

TODO make this required and give it a default

```typescript
queryType(queryType: string)
```

### <span class="badge object-method"></span> rawFrameContent

```typescript
rawFrameContent(rawFrameContent: string)
```

### <span class="badge object-method"></span> refId

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```typescript
refId(refId: string)
```

### <span class="badge object-method"></span> scenarioId

```typescript
scenarioId(scenarioId: testdata.TestDataQueryType)
```

### <span class="badge object-method"></span> seriesCount

```typescript
seriesCount(seriesCount: number)
```

### <span class="badge object-method"></span> sim

```typescript
sim(sim: cog.Builder<testdata.SimulationQuery>)
```

### <span class="badge object-method"></span> spanCount

```typescript
spanCount(spanCount: number)
```

### <span class="badge object-method"></span> stream

```typescript
stream(stream: cog.Builder<testdata.StreamingQuery>)
```

### <span class="badge object-method"></span> stringInput

```typescript
stringInput(stringInput: string)
```

### <span class="badge object-method"></span> usa

```typescript
usa(usa: cog.Builder<testdata.USAQuery>)
```

## See also

 * <span class="badge object-type-interface"></span> [dataquery](./object-dataquery.md)
