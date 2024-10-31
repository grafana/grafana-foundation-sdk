---
title: <span class="badge object-type-class"></span> CloudMonitoringQuery
---
# <span class="badge object-type-class"></span> CloudMonitoringQuery

## Definition

```php
class CloudMonitoringQuery implements \JsonSerializable, \Grafana\Foundation\Cog\Dataquery
{
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
     * Aliases can be set to modify the legend labels. e.g. {{metric.label.xxx}}. See docs for more detail.
     */
    public ?string $aliasBy;

    /**
     * GCM query type.
     * queryType: #QueryType
     * Time Series List sub-query properties.
     */
    public ?\Grafana\Foundation\Googlecloudmonitoring\TimeSeriesList $timeSeriesList;

    /**
     * Time Series sub-query properties.
     */
    public ?\Grafana\Foundation\Googlecloudmonitoring\TimeSeriesQuery $timeSeriesQuery;

    /**
     * SLO sub-query properties.
     */
    public ?\Grafana\Foundation\Googlecloudmonitoring\SLOQuery $sloQuery;

    /**
     * PromQL sub-query properties.
     */
    public ?\Grafana\Foundation\Googlecloudmonitoring\PromQLQuery $promQLQuery;

    /**
     * For mixed data sources the selected datasource is on the query level.
     * For non mixed scenarios this is undefined.
     * TODO find a better way to do this ^ that's friendly to schema
     * TODO this shouldn't be unknown but DataSourceRef | null
     */
    public ?\Grafana\Foundation\Dashboard\DataSourceRef $datasource;

    /**
     * Time interval in milliseconds.
     */
    public ?float $intervalMs;

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

 * <span class="badge builder"></span> [CloudMonitoringQueryBuilder](./builder-CloudMonitoringQueryBuilder.md)
