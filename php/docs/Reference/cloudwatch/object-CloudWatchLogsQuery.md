---
title: <span class="badge object-type-class"></span> CloudWatchLogsQuery
---
# <span class="badge object-type-class"></span> CloudWatchLogsQuery

Shape of a CloudWatch Logs query

## Definition

```php
class CloudWatchLogsQuery implements \JsonSerializable, \Grafana\Foundation\Cog\Dataquery
{
    /**
     * Whether a query is a Metrics, Logs, or Annotations query
     */
    public \Grafana\Foundation\Cloudwatch\CloudWatchQueryMode $queryMode;

    public string $id;

    /**
     * AWS region to query for the logs
     */
    public string $region;

    /**
     * The CloudWatch Logs Insights query to execute
     */
    public ?string $expression;

    /**
     * Fields to group the results by, this field is automatically populated whenever the query is updated
     * @var array<string>|null
     */
    public ?array $statsGroups;

    /**
     * Log groups to query
     * @var array<\Grafana\Foundation\Cloudwatch\LogGroup>|null
     */
    public ?array $logGroups;

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
     * @deprecated use logGroups
     * @var array<string>|null
     */
    public ?array $logGroupNames;

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

 * <span class="badge builder"></span> [CloudWatchLogsQueryBuilder](./builder-CloudWatchLogsQueryBuilder.md)
