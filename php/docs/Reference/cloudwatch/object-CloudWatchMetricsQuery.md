---
title: <span class="badge object-type-class"></span> CloudWatchMetricsQuery
---
# <span class="badge object-type-class"></span> CloudWatchMetricsQuery

Shape of a CloudWatch Metrics query

## Definition

```php
class CloudWatchMetricsQuery implements \JsonSerializable, \Grafana\Foundation\Cog\Dataquery
{
    /**
     * Whether a query is a Metrics, Logs, or Annotations query
     */
    public \Grafana\Foundation\Cloudwatch\CloudWatchQueryMode $queryMode;

    /**
     * Whether to use a metric search or metric query. Metric query is referred to as "Metrics Insights" in the AWS console.
     */
    public ?\Grafana\Foundation\Cloudwatch\MetricQueryType $metricQueryType;

    /**
     * Whether to use the query builder or code editor to create the query
     */
    public ?\Grafana\Foundation\Cloudwatch\MetricEditorMode $metricEditorMode;

    /**
     * ID can be used to reference other queries in math expressions. The ID can include numbers, letters, and underscore, and must start with a lowercase letter.
     */
    public string $id;

    /**
     * Deprecated: use label
     * @deprecated use label
     */
    public ?string $alias;

    /**
     * Change the time series legend names using dynamic labels. See https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/graph-dynamic-labels.html for more details.
     */
    public ?string $label;

    /**
     * Math expression query
     */
    public ?string $expression;

    /**
     * When the metric query type is `metricQueryType` is set to `Query`, this field is used to specify the query string.
     */
    public ?string $sqlExpression;

    /**
     * A unique identifier for the query within the list of targets.
     * In server side expressions, the refId is used as a variable name to identify results.
     * By default, the UI will assign A->Z; however setting meaningful names may be useful.
     */
    public string $refId;

    /**
     * true if query is disabled (ie should not be returned to the dashboard)
     * Note this does not always imply that the query should not be executed since
     * the results from a hidden query may be used as the input to other queries (SSE etc)
     */
    public ?bool $hide;

    /**
     * Specify the query flavor
     * TODO make this required and give it a default
     */
    public ?string $queryType;

    /**
     * AWS region to query for the metric
     */
    public string $region;

    /**
     * A namespace is a container for CloudWatch metrics. Metrics in different namespaces are isolated from each other, so that metrics from different applications are not mistakenly aggregated into the same statistics. For example, Amazon EC2 uses the AWS/EC2 namespace.
     */
    public string $namespace;

    /**
     * Name of the metric
     */
    public ?string $metricName;

    /**
     * The dimensions of the metric
     * @var array<string, string|array<string>>
     */
    public array $dimensions;

    /**
     * Only show metrics that exactly match all defined dimension names.
     */
    public ?bool $matchExact;

    /**
     * The length of time associated with a specific Amazon CloudWatch statistic. Can be specified by a number of seconds, 'auto', or as a duration string e.g. '15m' being 15 minutes
     */
    public ?string $period;

    /**
     * The ID of the AWS account to query for the metric, specifying `all` will query all accounts that the monitoring account is permitted to query.
     */
    public ?string $accountId;

    /**
     * Metric data aggregations over specified periods of time. For detailed definitions of the statistics supported by CloudWatch, see https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Statistics-definitions.html.
     */
    public ?string $statistic;

    /**
     * When the metric query type is `metricQueryType` is set to `Query` and the `metricEditorMode` is set to `Builder`, this field is used to build up an object representation of a SQL query.
     */
    public ?\Grafana\Foundation\Cloudwatch\SQLExpression $sql;

    /**
     * For mixed data sources the selected datasource is on the query level.
     * For non mixed scenarios this is undefined.
     * TODO find a better way to do this ^ that's friendly to schema
     * TODO this shouldn't be unknown but DataSourceRef | null
     */
    public ?\Grafana\Foundation\Dashboard\DataSourceRef $datasource;

    /**
     * @deprecated use statistic
     * @var array<string>|null
     */
    public ?array $statistics;

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

 * <span class="badge builder"></span> [CloudWatchMetricsQueryBuilder](./builder-CloudWatchMetricsQueryBuilder.md)
