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

For mixed data sources the selected datasource is on the query level.

For non mixed scenarios this is undefined.

TODO find a better way to do this ^ that's friendly to schema

TODO this shouldn't be unknown but DataSourceRef | null

```php
datasource(\Grafana\Foundation\Dashboard\DataSourceRef $datasource)
```

### <span class="badge object-method"></span> dropPercent

Drop percentage (the chance we will lose a point 0-100)

```php
dropPercent(float $dropPercent)
```

### <span class="badge object-method"></span> errorType

```php
errorType(\Grafana\Foundation\Testdata\DataqueryErrorType $errorType)
```

### <span class="badge object-method"></span> flamegraphDiff

```php
flamegraphDiff(bool $flamegraphDiff)
```

### <span class="badge object-method"></span> hide

true if query is disabled (ie should not be returned to the dashboard)

Note this does not always imply that the query should not be executed since

the results from a hidden query may be used as the input to other queries (SSE etc)

```php
hide(bool $hide)
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

### <span class="badge object-method"></span> nodes

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Testdata\NodesQuery> $nodes

```php
nodes(\Grafana\Foundation\Cog\Builder $nodes)
```

### <span class="badge object-method"></span> points

@param array<array<string|int>> $points

```php
points(array $points)
```

### <span class="badge object-method"></span> pulseWave

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Testdata\PulseWaveQuery> $pulseWave

```php
pulseWave(\Grafana\Foundation\Cog\Builder $pulseWave)
```

### <span class="badge object-method"></span> queryType

Specify the query flavor

TODO make this required and give it a default

```php
queryType(string $queryType)
```

### <span class="badge object-method"></span> rawFrameContent

```php
rawFrameContent(string $rawFrameContent)
```

### <span class="badge object-method"></span> refId

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```php
refId(string $refId)
```

### <span class="badge object-method"></span> scenarioId

```php
scenarioId(\Grafana\Foundation\Testdata\TestDataQueryType $scenarioId)
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

### <span class="badge object-method"></span> stream

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Testdata\StreamingQuery> $stream

```php
stream(\Grafana\Foundation\Cog\Builder $stream)
```

### <span class="badge object-method"></span> stringInput

```php
stringInput(string $stringInput)
```

### <span class="badge object-method"></span> usa

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Testdata\USAQuery> $usa

```php
usa(\Grafana\Foundation\Cog\Builder $usa)
```

## See also

 * <span class="badge object-type-class"></span> [Dataquery](./object-Dataquery.md)
