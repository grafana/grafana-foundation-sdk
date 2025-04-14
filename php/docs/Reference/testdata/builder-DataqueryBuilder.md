---
title: <span class="badge builder"></span> DataqueryBuilder
---
# <span class="badge builder"></span> DataqueryBuilder

## Constructor

```php
new DataqueryBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> alias

```php
alias(string $alias)
```

### <span class="badge object-method"></span> channel

Used for live query

```php
channel(string $channel)
```

### <span class="badge object-method"></span> csvContent

```php
csvContent(string $csvContent)
```

### <span class="badge object-method"></span> csvFileName

```php
csvFileName(string $csvFileName)
```

### <span class="badge object-method"></span> csvWave

@param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Testdata\CSVWave>> $csvWave

```php
csvWave(array $csvWave)
```

### <span class="badge object-method"></span> datasource

The datasource

```php
datasource(\Grafana\Foundation\Dashboard\DataSourceRef $datasource)
```

### <span class="badge object-method"></span> dropPercent

Drop percentage (the chance we will lose a point 0-100)

```php
dropPercent(float $dropPercent)
```

### <span class="badge object-method"></span> errorSource

Possible enum values:

 - `"plugin"` 

 - `"downstream"` 

```php
errorSource(\Grafana\Foundation\Testdata\DataqueryErrorSource $errorSource)
```

### <span class="badge object-method"></span> errorType

Possible enum values:

 - `"frontend_exception"` 

 - `"frontend_observable"` 

 - `"server_panic"` 

```php
errorType(\Grafana\Foundation\Testdata\DataqueryErrorType $errorType)
```

### <span class="badge object-method"></span> flamegraphDiff

```php
flamegraphDiff(bool $flamegraphDiff)
```

### <span class="badge object-method"></span> hide

true if query is disabled (ie should not be returned to the dashboard)

NOTE: this does not always imply that the query should not be executed since

the results from a hidden query may be used as the input to other queries (SSE etc)

```php
hide(bool $hide)
```

### <span class="badge object-method"></span> intervalMs

Interval is the suggested duration between time points in a time series query.

NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated

from the interval required to fill a pixels in the visualization

```php
intervalMs(float $intervalMs)
```

### <span class="badge object-method"></span> labels

```php
labels(string $labels)
```

### <span class="badge object-method"></span> levelColumn

```php
levelColumn(bool $levelColumn)
```

### <span class="badge object-method"></span> lines

```php
lines(int $lines)
```

### <span class="badge object-method"></span> max

```php
max(float $max)
```

### <span class="badge object-method"></span> maxDataPoints

MaxDataPoints is the maximum number of data points that should be returned from a time series query.

NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated

from the number of pixels visible in a visualization

```php
maxDataPoints(int $maxDataPoints)
```

### <span class="badge object-method"></span> min

```php
min(float $min)
```

### <span class="badge object-method"></span> nodes

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Testdata\NodesQuery> $nodes

```php
nodes(\Grafana\Foundation\Cog\Builder $nodes)
```

### <span class="badge object-method"></span> noise

```php
noise(float $noise)
```

### <span class="badge object-method"></span> points

@param array<array<mixed>> $points

```php
points(array $points)
```

### <span class="badge object-method"></span> pulseWave

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Testdata\PulseWaveQuery> $pulseWave

```php
pulseWave(\Grafana\Foundation\Cog\Builder $pulseWave)
```

### <span class="badge object-method"></span> queryType

QueryType is an optional identifier for the type of query.

It can be used to distinguish different types of queries.

```php
queryType(string $queryType)
```

### <span class="badge object-method"></span> rawFrameContent

```php
rawFrameContent(string $rawFrameContent)
```

### <span class="badge object-method"></span> refId

RefID is the unique identifier of the query, set by the frontend call.

```php
refId(string $refId)
```

### <span class="badge object-method"></span> resultAssertions

Optionally define expected query result behavior

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Testdata\ResultAssertions> $resultAssertions

```php
resultAssertions(\Grafana\Foundation\Cog\Builder $resultAssertions)
```

### <span class="badge object-method"></span> scenarioId

Possible enum values:

 - `"annotations"` 

 - `"arrow"` 

 - `"csv_content"` 

 - `"csv_file"` 

 - `"csv_metric_values"` 

 - `"datapoints_outside_range"` 

 - `"error_with_source"` 

 - `"exponential_heatmap_bucket_data"` 

 - `"flame_graph"` 

 - `"grafana_api"` 

 - `"linear_heatmap_bucket_data"` 

 - `"live"` 

 - `"logs"` 

 - `"manual_entry"` 

 - `"no_data_points"` 

 - `"node_graph"` 

 - `"predictable_csv_wave"` 

 - `"predictable_pulse"` 

 - `"random_walk"` 

 - `"random_walk_table"` 

 - `"random_walk_with_error"` 

 - `"raw_frame"` 

 - `"server_error_500"` 

 - `"simulation"` 

 - `"slow_query"` 

 - `"streaming_client"` 

 - `"table_static"` 

 - `"trace"` 

 - `"usa"` 

 - `"variables-query"` 

```php
scenarioId(\Grafana\Foundation\Testdata\DataqueryScenarioId $scenarioId)
```

### <span class="badge object-method"></span> seriesCount

```php
seriesCount(int $seriesCount)
```

### <span class="badge object-method"></span> sim

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Testdata\SimulationQuery> $sim

```php
sim(\Grafana\Foundation\Cog\Builder $sim)
```

### <span class="badge object-method"></span> spanCount

```php
spanCount(int $spanCount)
```

### <span class="badge object-method"></span> spread

```php
spread(float $spread)
```

### <span class="badge object-method"></span> startValue

```php
startValue(float $startValue)
```

### <span class="badge object-method"></span> stream

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Testdata\StreamingQuery> $stream

```php
stream(\Grafana\Foundation\Cog\Builder $stream)
```

### <span class="badge object-method"></span> stringInput

common parameter used by many query types

```php
stringInput(string $stringInput)
```

### <span class="badge object-method"></span> timeRange

TimeRange represents the query range

NOTE: unlike generic /ds/query, we can now send explicit time values in each query

NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Testdata\TimeRange> $timeRange

```php
timeRange(\Grafana\Foundation\Cog\Builder $timeRange)
```

### <span class="badge object-method"></span> usa

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Testdata\USAQuery> $usa

```php
usa(\Grafana\Foundation\Cog\Builder $usa)
```

### <span class="badge object-method"></span> withNil

```php
withNil(bool $withNil)
```

## See also

 * <span class="badge object-type-class"></span> [Dataquery](./object-Dataquery.md)
