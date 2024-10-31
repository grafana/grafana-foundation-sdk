<?php

namespace Grafana\Foundation\Cloudwatch;

/**
 * Shape of a CloudWatch Metrics query
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Cloudwatch\CloudWatchMetricsQuery>
 */
class CloudWatchMetricsQueryBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Cloudwatch\CloudWatchMetricsQuery $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Cloudwatch\CloudWatchMetricsQuery();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Cloudwatch\CloudWatchMetricsQuery
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Whether a query is a Metrics, Logs, or Annotations query
     */
    public function queryMode(\Grafana\Foundation\Cloudwatch\CloudWatchQueryMode $queryMode): static
    {
        $this->internal->queryMode = $queryMode;
    
        return $this;
    }
    /**
     * Whether to use a metric search or metric query. Metric query is referred to as "Metrics Insights" in the AWS console.
     */
    public function metricQueryType(\Grafana\Foundation\Cloudwatch\MetricQueryType $metricQueryType): static
    {
        $this->internal->metricQueryType = $metricQueryType;
    
        return $this;
    }
    /**
     * Whether to use the query builder or code editor to create the query
     */
    public function metricEditorMode(\Grafana\Foundation\Cloudwatch\MetricEditorMode $metricEditorMode): static
    {
        $this->internal->metricEditorMode = $metricEditorMode;
    
        return $this;
    }
    /**
     * ID can be used to reference other queries in math expressions. The ID can include numbers, letters, and underscore, and must start with a lowercase letter.
     */
    public function id(string $id): static
    {
        $this->internal->id = $id;
    
        return $this;
    }
    /**
     * Deprecated: use label
     * @deprecated use label
     */
    public function alias(string $alias): static
    {
        $this->internal->alias = $alias;
    
        return $this;
    }
    /**
     * Change the time series legend names using dynamic labels. See https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/graph-dynamic-labels.html for more details.
     */
    public function label(string $label): static
    {
        $this->internal->label = $label;
    
        return $this;
    }
    /**
     * Math expression query
     */
    public function expression(string $expression): static
    {
        $this->internal->expression = $expression;
    
        return $this;
    }
    /**
     * When the metric query type is `metricQueryType` is set to `Query`, this field is used to specify the query string.
     */
    public function sqlExpression(string $sqlExpression): static
    {
        $this->internal->sqlExpression = $sqlExpression;
    
        return $this;
    }
    /**
     * A unique identifier for the query within the list of targets.
     * In server side expressions, the refId is used as a variable name to identify results.
     * By default, the UI will assign A->Z; however setting meaningful names may be useful.
     */
    public function refId(string $refId): static
    {
        $this->internal->refId = $refId;
    
        return $this;
    }
    /**
     * If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
     */
    public function hide(bool $hide): static
    {
        $this->internal->hide = $hide;
    
        return $this;
    }
    /**
     * Specify the query flavor
     * TODO make this required and give it a default
     */
    public function queryType(string $queryType): static
    {
        $this->internal->queryType = $queryType;
    
        return $this;
    }
    /**
     * AWS region to query for the metric
     */
    public function region(string $region): static
    {
        $this->internal->region = $region;
    
        return $this;
    }
    /**
     * A namespace is a container for CloudWatch metrics. Metrics in different namespaces are isolated from each other, so that metrics from different applications are not mistakenly aggregated into the same statistics. For example, Amazon EC2 uses the AWS/EC2 namespace.
     */
    public function namespace(string $namespace): static
    {
        $this->internal->namespace = $namespace;
    
        return $this;
    }
    /**
     * Name of the metric
     */
    public function metricName(string $metricName): static
    {
        $this->internal->metricName = $metricName;
    
        return $this;
    }
    /**
     * The dimensions of the metric
     * @param array<string, string|array<string>> $dimensions
     */
    public function dimensions(array $dimensions): static
    {
        $this->internal->dimensions = $dimensions;
    
        return $this;
    }
    /**
     * Only show metrics that exactly match all defined dimension names.
     */
    public function matchExact(bool $matchExact): static
    {
        $this->internal->matchExact = $matchExact;
    
        return $this;
    }
    /**
     * The length of time associated with a specific Amazon CloudWatch statistic. Can be specified by a number of seconds, 'auto', or as a duration string e.g. '15m' being 15 minutes
     */
    public function period(string $period): static
    {
        $this->internal->period = $period;
    
        return $this;
    }
    /**
     * The ID of the AWS account to query for the metric, specifying `all` will query all accounts that the monitoring account is permitted to query.
     */
    public function accountId(string $accountId): static
    {
        $this->internal->accountId = $accountId;
    
        return $this;
    }
    /**
     * Metric data aggregations over specified periods of time. For detailed definitions of the statistics supported by CloudWatch, see https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Statistics-definitions.html.
     */
    public function statistic(string $statistic): static
    {
        $this->internal->statistic = $statistic;
    
        return $this;
    }
    /**
     * When the metric query type is `metricQueryType` is set to `Query` and the `metricEditorMode` is set to `Builder`, this field is used to build up an object representation of a SQL query.
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Cloudwatch\SQLExpression> $sql
     */
    public function sql(\Grafana\Foundation\Cog\Builder $sql): static
    {
        $sqlResource = $sql->build();
        $this->internal->sql = $sqlResource;
    
        return $this;
    }
    /**
     * For mixed data sources the selected datasource is on the query level.
     * For non mixed scenarios this is undefined.
     * TODO find a better way to do this ^ that's friendly to schema
     * TODO this shouldn't be unknown but DataSourceRef | null
     */
    public function datasource(\Grafana\Foundation\Dashboard\DataSourceRef $datasource): static
    {
        $this->internal->datasource = $datasource;
    
        return $this;
    }
    /**
     * @deprecated use statistic
     * @param array<string> $statistics
     */
    public function statistics(array $statistics): static
    {
        $this->internal->statistics = $statistics;
    
        return $this;
    }

}
