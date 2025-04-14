---
title: <span class="badge object-type-class"></span> Dataquery
---
# <span class="badge object-type-class"></span> Dataquery

## Definition

```php
class Dataquery implements \JsonSerializable, \Grafana\Foundation\Cog\Dataquery
{
    public ?string $alias;

    /**
     * Used for live query
     */
    public ?string $channel;

    public ?string $csvContent;

    public ?string $csvFileName;

    /**
     * @var array<\Grafana\Foundation\Testdata\CSVWave>|null
     */
    public ?array $csvWave;

    /**
     * The datasource
     */
    public ?\Grafana\Foundation\Dashboard\DataSourceRef $datasource;

    /**
     * Drop percentage (the chance we will lose a point 0-100)
     */
    public ?float $dropPercent;

    /**
     * Possible enum values:
     *  - `"plugin"` 
     *  - `"downstream"` 
     */
    public ?\Grafana\Foundation\Testdata\DataqueryErrorSource $errorSource;

    /**
     * Possible enum values:
     *  - `"frontend_exception"` 
     *  - `"frontend_observable"` 
     *  - `"server_panic"` 
     */
    public ?\Grafana\Foundation\Testdata\DataqueryErrorType $errorType;

    public ?bool $flamegraphDiff;

    /**
     * true if query is disabled (ie should not be returned to the dashboard)
     * NOTE: this does not always imply that the query should not be executed since
     * the results from a hidden query may be used as the input to other queries (SSE etc)
     */
    public ?bool $hide;

    /**
     * Interval is the suggested duration between time points in a time series query.
     * NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
     * from the interval required to fill a pixels in the visualization
     */
    public ?float $intervalMs;

    public ?string $labels;

    public ?bool $levelColumn;

    public ?int $lines;

    public ?float $max;

    /**
     * MaxDataPoints is the maximum number of data points that should be returned from a time series query.
     * NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
     * from the number of pixels visible in a visualization
     */
    public ?int $maxDataPoints;

    public ?float $min;

    public ?\Grafana\Foundation\Testdata\NodesQuery $nodes;

    public ?float $noise;

    /**
     * @var array<array<mixed>>|null
     */
    public ?array $points;

    public ?\Grafana\Foundation\Testdata\PulseWaveQuery $pulseWave;

    /**
     * QueryType is an optional identifier for the type of query.
     * It can be used to distinguish different types of queries.
     */
    public ?string $queryType;

    public ?string $rawFrameContent;

    /**
     * RefID is the unique identifier of the query, set by the frontend call.
     */
    public ?string $refId;

    /**
     * Optionally define expected query result behavior
     */
    public ?\Grafana\Foundation\Testdata\ResultAssertions $resultAssertions;

    /**
     * Possible enum values:
     *  - `"annotations"` 
     *  - `"arrow"` 
     *  - `"csv_content"` 
     *  - `"csv_file"` 
     *  - `"csv_metric_values"` 
     *  - `"datapoints_outside_range"` 
     *  - `"error_with_source"` 
     *  - `"exponential_heatmap_bucket_data"` 
     *  - `"flame_graph"` 
     *  - `"grafana_api"` 
     *  - `"linear_heatmap_bucket_data"` 
     *  - `"live"` 
     *  - `"logs"` 
     *  - `"manual_entry"` 
     *  - `"no_data_points"` 
     *  - `"node_graph"` 
     *  - `"predictable_csv_wave"` 
     *  - `"predictable_pulse"` 
     *  - `"random_walk"` 
     *  - `"random_walk_table"` 
     *  - `"random_walk_with_error"` 
     *  - `"raw_frame"` 
     *  - `"server_error_500"` 
     *  - `"simulation"` 
     *  - `"slow_query"` 
     *  - `"streaming_client"` 
     *  - `"table_static"` 
     *  - `"trace"` 
     *  - `"usa"` 
     *  - `"variables-query"` 
     */
    public ?\Grafana\Foundation\Testdata\DataqueryScenarioId $scenarioId;

    public ?int $seriesCount;

    public ?\Grafana\Foundation\Testdata\SimulationQuery $sim;

    public ?int $spanCount;

    public ?float $spread;

    public ?float $startValue;

    public ?\Grafana\Foundation\Testdata\StreamingQuery $stream;

    /**
     * common parameter used by many query types
     */
    public ?string $stringInput;

    /**
     * TimeRange represents the query range
     * NOTE: unlike generic /ds/query, we can now send explicit time values in each query
     * NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
     */
    public ?\Grafana\Foundation\Testdata\TimeRange $timeRange;

    public ?\Grafana\Foundation\Testdata\USAQuery $usa;

    public ?bool $withNil;

}
```
## Methods

### <span class="badge object-method"></span> fromArray

Builds this object from an array.

This function is meant to be used with the return value of `json_decode($json, true)`.

```php
static fromArray(array $inputData)
```

### <span class="badge object-method"></span> jsonSerialize

Returns the data representing this object, preparing it for JSON serialization with `json_encode()`.

```php
jsonSerialize()
```

### <span class="badge object-method"></span> dataqueryType

Returns the type of this dataquery object.

```php
dataqueryType()
```

## See also

 * <span class="badge builder"></span> [DataqueryBuilder](./builder-DataqueryBuilder.md)
