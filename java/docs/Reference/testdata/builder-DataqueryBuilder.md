---
title: <span class="badge builder"></span> DataqueryBuilder
---
# <span class="badge builder"></span> DataqueryBuilder

## Constructor

```java
new DataqueryBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public Dataquery build()
```

### <span class="badge object-method"></span> alias

```java
public dataqueryBuilder alias(String alias)
```

### <span class="badge object-method"></span> channel

```java
public dataqueryBuilder channel(String channel)
```

### <span class="badge object-method"></span> csvContent

```java
public dataqueryBuilder csvContent(String csvContent)
```

### <span class="badge object-method"></span> csvFileName

```java
public dataqueryBuilder csvFileName(String csvFileName)
```

### <span class="badge object-method"></span> csvWave

```java
public dataqueryBuilder csvWave(List<com.grafana.foundation.cog.Builder<CSVWave>> csvWave)
```

### <span class="badge object-method"></span> datasource

For mixed data sources the selected datasource is on the query level.

For non mixed scenarios this is undefined.

TODO find a better way to do this ^ that's friendly to schema

TODO this shouldn't be unknown but DataSourceRef | null

```java
public dataqueryBuilder datasource(Object datasource)
```

### <span class="badge object-method"></span> dropPercent

Drop percentage (the chance we will lose a point 0-100)

```java
public dataqueryBuilder dropPercent(Double dropPercent)
```

### <span class="badge object-method"></span> errorType

```java
public dataqueryBuilder errorType(DataqueryErrorType errorType)
```

### <span class="badge object-method"></span> hide

true if query is disabled (ie should not be returned to the dashboard)

Note this does not always imply that the query should not be executed since

the results from a hidden query may be used as the input to other queries (SSE etc)

```java
public dataqueryBuilder hide(Boolean hide)
```

### <span class="badge object-method"></span> labels

```java
public dataqueryBuilder labels(String labels)
```

### <span class="badge object-method"></span> levelColumn

```java
public dataqueryBuilder levelColumn(Boolean levelColumn)
```

### <span class="badge object-method"></span> lines

```java
public dataqueryBuilder lines(Long lines)
```

### <span class="badge object-method"></span> nodes

```java
public dataqueryBuilder nodes(com.grafana.foundation.cog.Builder<NodesQuery> nodes)
```

### <span class="badge object-method"></span> points

```java
public dataqueryBuilder points(List<List<StringOrInt64>> points)
```

### <span class="badge object-method"></span> pulseWave

```java
public dataqueryBuilder pulseWave(com.grafana.foundation.cog.Builder<PulseWaveQuery> pulseWave)
```

### <span class="badge object-method"></span> queryType

Specify the query flavor

TODO make this required and give it a default

```java
public dataqueryBuilder queryType(String queryType)
```

### <span class="badge object-method"></span> rawFrameContent

```java
public dataqueryBuilder rawFrameContent(String rawFrameContent)
```

### <span class="badge object-method"></span> refId

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```java
public dataqueryBuilder refId(String refId)
```

### <span class="badge object-method"></span> scenarioId

```java
public dataqueryBuilder scenarioId(TestDataQueryType scenarioId)
```

### <span class="badge object-method"></span> seriesCount

```java
public dataqueryBuilder seriesCount(Integer seriesCount)
```

### <span class="badge object-method"></span> sim

```java
public dataqueryBuilder sim(com.grafana.foundation.cog.Builder<SimulationQuery> sim)
```

### <span class="badge object-method"></span> spanCount

```java
public dataqueryBuilder spanCount(Integer spanCount)
```

### <span class="badge object-method"></span> stream

```java
public dataqueryBuilder stream(com.grafana.foundation.cog.Builder<StreamingQuery> stream)
```

### <span class="badge object-method"></span> stringInput

```java
public dataqueryBuilder stringInput(String stringInput)
```

### <span class="badge object-method"></span> usa

```java
public dataqueryBuilder usa(com.grafana.foundation.cog.Builder<USAQuery> usa)
```

## See also

 * <span class="badge object-type-class"></span> [Dataquery](./object-Dataquery.md)
