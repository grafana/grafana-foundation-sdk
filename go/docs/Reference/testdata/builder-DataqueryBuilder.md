---
title: <span class="badge builder"></span> DataqueryBuilder
---
# <span class="badge builder"></span> DataqueryBuilder

## Constructor

```go
func NewDataqueryBuilder() *DataqueryBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *DataqueryBuilder) Build() (variants.Dataquery, error)
```

### <span class="badge object-method"></span> Alias

```go
func (builder *DataqueryBuilder) Alias(alias string) *DataqueryBuilder
```

### <span class="badge object-method"></span> Channel

```go
func (builder *DataqueryBuilder) Channel(channel string) *DataqueryBuilder
```

### <span class="badge object-method"></span> CsvContent

```go
func (builder *DataqueryBuilder) CsvContent(csvContent string) *DataqueryBuilder
```

### <span class="badge object-method"></span> CsvFileName

```go
func (builder *DataqueryBuilder) CsvFileName(csvFileName string) *DataqueryBuilder
```

### <span class="badge object-method"></span> CsvWave

```go
func (builder *DataqueryBuilder) CsvWave(csvWave []cog.Builder[testdata.CSVWave]) *DataqueryBuilder
```

### <span class="badge object-method"></span> Datasource

For mixed data sources the selected datasource is on the query level.

For non mixed scenarios this is undefined.

TODO find a better way to do this ^ that's friendly to schema

TODO this shouldn't be unknown but DataSourceRef | null

```go
func (builder *DataqueryBuilder) Datasource(datasource dashboard.DataSourceRef) *DataqueryBuilder
```

### <span class="badge object-method"></span> DropPercent

Drop percentage (the chance we will lose a point 0-100)

```go
func (builder *DataqueryBuilder) DropPercent(dropPercent float64) *DataqueryBuilder
```

### <span class="badge object-method"></span> ErrorType

```go
func (builder *DataqueryBuilder) ErrorType(errorType testdata.DataqueryErrorType) *DataqueryBuilder
```

### <span class="badge object-method"></span> FlamegraphDiff

```go
func (builder *DataqueryBuilder) FlamegraphDiff(flamegraphDiff bool) *DataqueryBuilder
```

### <span class="badge object-method"></span> Hide

true if query is disabled (ie should not be returned to the dashboard)

Note this does not always imply that the query should not be executed since

the results from a hidden query may be used as the input to other queries (SSE etc)

```go
func (builder *DataqueryBuilder) Hide(hide bool) *DataqueryBuilder
```

### <span class="badge object-method"></span> Labels

```go
func (builder *DataqueryBuilder) Labels(labels string) *DataqueryBuilder
```

### <span class="badge object-method"></span> LevelColumn

```go
func (builder *DataqueryBuilder) LevelColumn(levelColumn bool) *DataqueryBuilder
```

### <span class="badge object-method"></span> Lines

```go
func (builder *DataqueryBuilder) Lines(lines int64) *DataqueryBuilder
```

### <span class="badge object-method"></span> Nodes

```go
func (builder *DataqueryBuilder) Nodes(nodes cog.Builder[testdata.NodesQuery]) *DataqueryBuilder
```

### <span class="badge object-method"></span> Points

```go
func (builder *DataqueryBuilder) Points(points [][]testdata.StringOrInt64) *DataqueryBuilder
```

### <span class="badge object-method"></span> PulseWave

```go
func (builder *DataqueryBuilder) PulseWave(pulseWave cog.Builder[testdata.PulseWaveQuery]) *DataqueryBuilder
```

### <span class="badge object-method"></span> QueryType

Specify the query flavor

TODO make this required and give it a default

```go
func (builder *DataqueryBuilder) QueryType(queryType string) *DataqueryBuilder
```

### <span class="badge object-method"></span> RawFrameContent

```go
func (builder *DataqueryBuilder) RawFrameContent(rawFrameContent string) *DataqueryBuilder
```

### <span class="badge object-method"></span> RefId

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```go
func (builder *DataqueryBuilder) RefId(refId string) *DataqueryBuilder
```

### <span class="badge object-method"></span> ScenarioId

```go
func (builder *DataqueryBuilder) ScenarioId(scenarioId testdata.TestDataQueryType) *DataqueryBuilder
```

### <span class="badge object-method"></span> SeriesCount

```go
func (builder *DataqueryBuilder) SeriesCount(seriesCount int32) *DataqueryBuilder
```

### <span class="badge object-method"></span> Sim

```go
func (builder *DataqueryBuilder) Sim(sim cog.Builder[testdata.SimulationQuery]) *DataqueryBuilder
```

### <span class="badge object-method"></span> SpanCount

```go
func (builder *DataqueryBuilder) SpanCount(spanCount int32) *DataqueryBuilder
```

### <span class="badge object-method"></span> Stream

```go
func (builder *DataqueryBuilder) Stream(stream cog.Builder[testdata.StreamingQuery]) *DataqueryBuilder
```

### <span class="badge object-method"></span> StringInput

```go
func (builder *DataqueryBuilder) StringInput(stringInput string) *DataqueryBuilder
```

### <span class="badge object-method"></span> Usa

```go
func (builder *DataqueryBuilder) Usa(usa cog.Builder[testdata.USAQuery]) *DataqueryBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Dataquery](./object-Dataquery.md)
