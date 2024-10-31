# testdata

## Objects

 * <span class="badge object-type-struct"></span> [CSVWave](./object-CSVWave.md)
 * <span class="badge object-type-enum"></span> [DataqueryErrorType](./object-DataqueryErrorType.md)
 * <span class="badge object-type-struct"></span> [Key](./object-Key.md)
 * <span class="badge object-type-struct"></span> [NodesQuery](./object-NodesQuery.md)
 * <span class="badge object-type-enum"></span> [NodesQueryType](./object-NodesQueryType.md)
 * <span class="badge object-type-struct"></span> [PulseWaveQuery](./object-PulseWaveQuery.md)
 * <span class="badge object-type-struct"></span> [Scenario](./object-Scenario.md)
 * <span class="badge object-type-struct"></span> [SimulationQuery](./object-SimulationQuery.md)
 * <span class="badge object-type-struct"></span> [StreamingQuery](./object-StreamingQuery.md)
 * <span class="badge object-type-enum"></span> [StreamingQueryType](./object-StreamingQueryType.md)
 * <span class="badge object-type-struct"></span> [StringOrInt64](./object-StringOrInt64.md)
 * <span class="badge object-type-enum"></span> [TestDataQueryType](./object-TestDataQueryType.md)
 * <span class="badge object-type-struct"></span> [USAQuery](./object-USAQuery.md)
 * <span class="badge object-type-struct"></span> [Dataquery](./object-Dataquery.md)
## Builders

 * <span class="badge builder"></span> [CSVWaveBuilder](./builder-CSVWaveBuilder.md)
 * <span class="badge builder"></span> [KeyBuilder](./builder-KeyBuilder.md)
 * <span class="badge builder"></span> [NodesQueryBuilder](./builder-NodesQueryBuilder.md)
 * <span class="badge builder"></span> [PulseWaveQueryBuilder](./builder-PulseWaveQueryBuilder.md)
 * <span class="badge builder"></span> [ScenarioBuilder](./builder-ScenarioBuilder.md)
 * <span class="badge builder"></span> [SimulationQueryBuilder](./builder-SimulationQueryBuilder.md)
 * <span class="badge builder"></span> [StreamingQueryBuilder](./builder-StreamingQueryBuilder.md)
 * <span class="badge builder"></span> [USAQueryBuilder](./builder-USAQueryBuilder.md)
 * <span class="badge builder"></span> [DataqueryBuilder](./builder-DataqueryBuilder.md)
## Functions

### <span class="badge function"></span> VariantConfig

VariantConfig returns the configuration related to testdata dataqueries.

This configuration describes how to unmarshal it, convert it to code, …

```go
func VariantConfig() variants.DataqueryConfig
```

### <span class="badge function"></span> StreamingQueryConverter

StreamingQueryConverter accepts a `StreamingQuery` object and generates the Go code to build this object using builders.

```go
func StreamingQueryConverter(input StreamingQuery) string
```

### <span class="badge function"></span> PulseWaveQueryConverter

PulseWaveQueryConverter accepts a `PulseWaveQuery` object and generates the Go code to build this object using builders.

```go
func PulseWaveQueryConverter(input PulseWaveQuery) string
```

### <span class="badge function"></span> SimulationQueryConverter

SimulationQueryConverter accepts a `SimulationQuery` object and generates the Go code to build this object using builders.

```go
func SimulationQueryConverter(input SimulationQuery) string
```

### <span class="badge function"></span> NodesQueryConverter

NodesQueryConverter accepts a `NodesQuery` object and generates the Go code to build this object using builders.

```go
func NodesQueryConverter(input NodesQuery) string
```

### <span class="badge function"></span> USAQueryConverter

USAQueryConverter accepts a `USAQuery` object and generates the Go code to build this object using builders.

```go
func USAQueryConverter(input USAQuery) string
```

### <span class="badge function"></span> CSVWaveConverter

CSVWaveConverter accepts a `CSVWave` object and generates the Go code to build this object using builders.

```go
func CSVWaveConverter(input CSVWave) string
```

### <span class="badge function"></span> ScenarioConverter

ScenarioConverter accepts a `Scenario` object and generates the Go code to build this object using builders.

```go
func ScenarioConverter(input Scenario) string
```

### <span class="badge function"></span> DataqueryConverter

DataqueryConverter accepts a `Dataquery` object and generates the Go code to build this object using builders.

```go
func DataqueryConverter(input Dataquery) string
```

### <span class="badge function"></span> KeyConverter

KeyConverter accepts a `Key` object and generates the Go code to build this object using builders.

```go
func KeyConverter(input Key) string
```

