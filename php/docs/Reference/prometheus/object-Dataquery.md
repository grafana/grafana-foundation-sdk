---
title: <span class="badge object-type-class"></span> Dataquery
---
# <span class="badge object-type-class"></span> Dataquery

## Definition

```php
class Dataquery implements \JsonSerializable, \Grafana\Foundation\Cog\Dataquery
{
    /**
     * Additional Ad-hoc filters that take precedence over Scope on conflict.
     * @var array<\Grafana\Foundation\Prometheus\AdhocFilters>|null
     */
    public ?array $adhocFilters;

    /**
     * The datasource
     */
    public ?\Grafana\Foundation\Common\DataSourceRef $datasource;

    /**
     * what we should show in the editor
     * Possible enum values:
     *  - `"builder"` 
     *  - `"code"` 
     */
    public ?\Grafana\Foundation\Prometheus\QueryEditorMode $editorMode;

    /**
     * Execute an additional query to identify interesting raw samples relevant for the given expr
     */
    public ?bool $exemplar;

    /**
     * The actual expression/query that will be evaluated by Prometheus
     */
    public string $expr;

    /**
     * The response format
     * Possible enum values:
     *  - `"time_series"` 
     *  - `"table"` 
     *  - `"heatmap"` 
     */
    public ?\Grafana\Foundation\Prometheus\PromQueryFormat $format;

    /**
     * Group By parameters to apply to aggregate expressions in the query
     * @var array<string>|null
     */
    public ?array $groupByKeys;

    /**
     * true if query is disabled (ie should not be returned to the dashboard)
     * NOTE: this does not always imply that the query should not be executed since
     * the results from a hidden query may be used as the input to other queries (SSE etc)
     */
    public ?bool $hide;

    /**
     * Returns only the latest value that Prometheus has scraped for the requested time series
     */
    public ?bool $instant;

    /**
     * Used to specify how many times to divide max data points by. We use max data points under query options
     * See https://github.com/grafana/grafana/issues/48081
     * Deprecated: use interval
     */
    public ?int $intervalFactor;

    /**
     * Interval is the suggested duration between time points in a time series query.
     * NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
     * from the interval required to fill a pixels in the visualization
     */
    public ?float $intervalMs;

    /**
     * Series name override or template. Ex. {{hostname}} will be replaced with label value for hostname
     */
    public ?string $legendFormat;

    /**
     * MaxDataPoints is the maximum number of data points that should be returned from a time series query.
     * NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
     * from the number of pixels visible in a visualization
     */
    public ?int $maxDataPoints;

    /**
     * QueryType is an optional identifier for the type of query.
     * It can be used to distinguish different types of queries.
     */
    public ?string $queryType;

    /**
     * Returns a Range vector, comprised of a set of time series containing a range of data points over time for each time series
     */
    public ?bool $range;

    /**
     * RefID is the unique identifier of the query, set by the frontend call.
     */
    public ?string $refId;

    /**
     * Optionally define expected query result behavior
     */
    public ?\Grafana\Foundation\Prometheus\ResultAssertions $resultAssertions;

    /**
     * A set of filters applied to apply to the query
     * @var array<\Grafana\Foundation\Prometheus\Scopes>|null
     */
    public ?array $scopes;

    /**
     * TimeRange represents the query range
     * NOTE: unlike generic /ds/query, we can now send explicit time values in each query
     * NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
     */
    public ?\Grafana\Foundation\Prometheus\TimeRange $timeRange;

    /**
     * An additional lower limit for the step parameter of the Prometheus query and for the
     * `$__interval` and `$__rate_interval` variables.
     */
    public ?string $interval;

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
