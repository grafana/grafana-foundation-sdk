<?php

namespace Grafana\Foundation\Cloudwatch;

/**
 * Shape of a CloudWatch Metrics query
 */
class CloudWatchMetricsQuery implements \JsonSerializable, \Grafana\Foundation\Cog\Dataquery
{
    /**
     * Whether a query is a Metrics, Logs, or Annotations query
     */
    public \Grafana\Foundation\Cloudwatch\CloudWatchQueryMode $queryMode;

    /**
     * Whether to use a metric search or metric insights query
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
     * When the metric query type is set to `Insights`, this field is used to specify the query string.
     */
    public ?string $sqlExpression;

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
     * When the metric query type is set to `Insights` and the `metricEditorMode` is set to `Builder`, this field is used to build up an object representation of a SQL query.
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

    /**
     * @param \Grafana\Foundation\Cloudwatch\CloudWatchQueryMode|null $queryMode
     * @param \Grafana\Foundation\Cloudwatch\MetricQueryType|null $metricQueryType
     * @param \Grafana\Foundation\Cloudwatch\MetricEditorMode|null $metricEditorMode
     * @param string|null $id
     * @param string|null $alias
     * @param string|null $label
     * @param string|null $expression
     * @param string|null $sqlExpression
     * @param string|null $refId
     * @param bool|null $hide
     * @param string|null $queryType
     * @param string|null $region
     * @param string|null $namespace
     * @param string|null $metricName
     * @param array<string, string|array<string>>|null $dimensions
     * @param bool|null $matchExact
     * @param string|null $period
     * @param string|null $accountId
     * @param string|null $statistic
     * @param \Grafana\Foundation\Cloudwatch\SQLExpression|null $sql
     * @param \Grafana\Foundation\Dashboard\DataSourceRef|null $datasource
     * @param array<string>|null $statistics
     */
    public function __construct(?\Grafana\Foundation\Cloudwatch\CloudWatchQueryMode $queryMode = null, ?\Grafana\Foundation\Cloudwatch\MetricQueryType $metricQueryType = null, ?\Grafana\Foundation\Cloudwatch\MetricEditorMode $metricEditorMode = null, ?string $id = null, ?string $alias = null, ?string $label = null, ?string $expression = null, ?string $sqlExpression = null, ?string $refId = null, ?bool $hide = null, ?string $queryType = null, ?string $region = null, ?string $namespace = null, ?string $metricName = null, ?array $dimensions = null, ?bool $matchExact = null, ?string $period = null, ?string $accountId = null, ?string $statistic = null, ?\Grafana\Foundation\Cloudwatch\SQLExpression $sql = null, ?\Grafana\Foundation\Dashboard\DataSourceRef $datasource = null, ?array $statistics = null)
    {
        $this->queryMode = $queryMode ?: \Grafana\Foundation\Cloudwatch\CloudWatchQueryMode::metrics();
        $this->metricQueryType = $metricQueryType;
        $this->metricEditorMode = $metricEditorMode;
        $this->id = $id ?: "";
        $this->alias = $alias;
        $this->label = $label;
        $this->expression = $expression;
        $this->sqlExpression = $sqlExpression;
        $this->refId = $refId ?: "";
        $this->hide = $hide;
        $this->queryType = $queryType;
        $this->region = $region ?: "";
        $this->namespace = $namespace ?: "";
        $this->metricName = $metricName;
        $this->dimensions = $dimensions ?: [];
        $this->matchExact = $matchExact;
        $this->period = $period;
        $this->accountId = $accountId;
        $this->statistic = $statistic;
        $this->sql = $sql;
        $this->datasource = $datasource;
        $this->statistics = $statistics;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{queryMode?: string, metricQueryType?: int, metricEditorMode?: int, id?: string, alias?: string, label?: string, expression?: string, sqlExpression?: string, refId?: string, hide?: bool, queryType?: string, region?: string, namespace?: string, metricName?: string, dimensions?: array<string, string|array<string>>, matchExact?: bool, period?: string, accountId?: string, statistic?: string, sql?: mixed, datasource?: mixed, statistics?: array<string>} $inputData */
        $data = $inputData;
        return new self(
            queryMode: isset($data["queryMode"]) ? (function($input) { return \Grafana\Foundation\Cloudwatch\CloudWatchQueryMode::fromValue($input); })($data["queryMode"]) : null,
            metricQueryType: isset($data["metricQueryType"]) ? (function($input) { return \Grafana\Foundation\Cloudwatch\MetricQueryType::fromValue($input); })($data["metricQueryType"]) : null,
            metricEditorMode: isset($data["metricEditorMode"]) ? (function($input) { return \Grafana\Foundation\Cloudwatch\MetricEditorMode::fromValue($input); })($data["metricEditorMode"]) : null,
            id: $data["id"] ?? null,
            alias: $data["alias"] ?? null,
            label: $data["label"] ?? null,
            expression: $data["expression"] ?? null,
            sqlExpression: $data["sqlExpression"] ?? null,
            refId: $data["refId"] ?? null,
            hide: $data["hide"] ?? null,
            queryType: $data["queryType"] ?? null,
            region: $data["region"] ?? null,
            namespace: $data["namespace"] ?? null,
            metricName: $data["metricName"] ?? null,
            dimensions: $data["dimensions"] ?? null,
            matchExact: $data["matchExact"] ?? null,
            period: $data["period"] ?? null,
            accountId: $data["accountId"] ?? null,
            statistic: $data["statistic"] ?? null,
            sql: isset($data["sql"]) ? (function($input) {
    	/** @var array{select?: mixed, from?: mixed|mixed, where?: mixed, groupBy?: mixed, orderBy?: mixed, orderByDirection?: string, limit?: int} */
    $val = $input;
    	return \Grafana\Foundation\Cloudwatch\SQLExpression::fromArray($val);
    })($data["sql"]) : null,
            datasource: isset($data["datasource"]) ? (function($input) {
    	/** @var array{type?: string, uid?: string} */
    $val = $input;
    	return \Grafana\Foundation\Dashboard\DataSourceRef::fromArray($val);
    })($data["datasource"]) : null,
            statistics: $data["statistics"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "queryMode" => $this->queryMode,
            "id" => $this->id,
            "refId" => $this->refId,
            "region" => $this->region,
            "namespace" => $this->namespace,
            "dimensions" => $this->dimensions,
        ];
        if (isset($this->metricQueryType)) {
            $data["metricQueryType"] = $this->metricQueryType;
        }
        if (isset($this->metricEditorMode)) {
            $data["metricEditorMode"] = $this->metricEditorMode;
        }
        if (isset($this->alias)) {
            $data["alias"] = $this->alias;
        }
        if (isset($this->label)) {
            $data["label"] = $this->label;
        }
        if (isset($this->expression)) {
            $data["expression"] = $this->expression;
        }
        if (isset($this->sqlExpression)) {
            $data["sqlExpression"] = $this->sqlExpression;
        }
        if (isset($this->hide)) {
            $data["hide"] = $this->hide;
        }
        if (isset($this->queryType)) {
            $data["queryType"] = $this->queryType;
        }
        if (isset($this->metricName)) {
            $data["metricName"] = $this->metricName;
        }
        if (isset($this->matchExact)) {
            $data["matchExact"] = $this->matchExact;
        }
        if (isset($this->period)) {
            $data["period"] = $this->period;
        }
        if (isset($this->accountId)) {
            $data["accountId"] = $this->accountId;
        }
        if (isset($this->statistic)) {
            $data["statistic"] = $this->statistic;
        }
        if (isset($this->sql)) {
            $data["sql"] = $this->sql;
        }
        if (isset($this->datasource)) {
            $data["datasource"] = $this->datasource;
        }
        if (isset($this->statistics)) {
            $data["statistics"] = $this->statistics;
        }
        return $data;
    }

    public function dataqueryType(): string
    {
        return "cloudwatch";
    }
}
