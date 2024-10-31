---
title: <span class="badge object-type-class"></span> Dataquery
---
# <span class="badge object-type-class"></span> Dataquery

## Definition

```php
class Dataquery implements \JsonSerializable, \Grafana\Foundation\Cog\Dataquery
{
    /**
     * Alias pattern
     */
    public ?string $alias;

    /**
     * Lucene query
     */
    public ?string $query;

    /**
     * Name of time field
     */
    public ?string $timeField;

    /**
     * List of bucket aggregations
     * @var array<\Grafana\Foundation\Elasticsearch\DateHistogram|\Grafana\Foundation\Elasticsearch\Histogram|\Grafana\Foundation\Elasticsearch\Terms|\Grafana\Foundation\Elasticsearch\Filters|\Grafana\Foundation\Elasticsearch\GeoHashGrid|\Grafana\Foundation\Elasticsearch\Nested>|null
     */
    public ?array $bucketAggs;

    /**
     * List of metric aggregations
     * @var array<\Grafana\Foundation\Elasticsearch\Count|\Grafana\Foundation\Elasticsearch\MovingAverage|\Grafana\Foundation\Elasticsearch\Derivative|\Grafana\Foundation\Elasticsearch\CumulativeSum|\Grafana\Foundation\Elasticsearch\BucketScript|\Grafana\Foundation\Elasticsearch\SerialDiff|\Grafana\Foundation\Elasticsearch\RawData|\Grafana\Foundation\Elasticsearch\RawDocument|\Grafana\Foundation\Elasticsearch\UniqueCount|\Grafana\Foundation\Elasticsearch\Percentiles|\Grafana\Foundation\Elasticsearch\ExtendedStats|\Grafana\Foundation\Elasticsearch\Min|\Grafana\Foundation\Elasticsearch\Max|\Grafana\Foundation\Elasticsearch\Sum|\Grafana\Foundation\Elasticsearch\Average|\Grafana\Foundation\Elasticsearch\MovingFunction|\Grafana\Foundation\Elasticsearch\Logs|\Grafana\Foundation\Elasticsearch\Rate|\Grafana\Foundation\Elasticsearch\TopMetrics>|null
     */
    public ?array $metrics;

    /**
     * A unique identifier for the query within the list of targets.
     * In server side expressions, the refId is used as a variable name to identify results.
     * By default, the UI will assign A->Z; however setting meaningful names may be useful.
     */
    public string $refId;

    /**
     * If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
     */
    public ?bool $hide;

    /**
     * Specify the query flavor
     * TODO make this required and give it a default
     */
    public ?string $queryType;

    /**
     * For mixed data sources the selected datasource is on the query level.
     * For non mixed scenarios this is undefined.
     * TODO find a better way to do this ^ that's friendly to schema
     * TODO this shouldn't be unknown but DataSourceRef | null
     */
    public ?\Grafana\Foundation\Dashboard\DataSourceRef $datasource;

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
