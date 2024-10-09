<?php

namespace Grafana\Foundation\Cloudwatch;

/**
 * Shape of a CloudWatch Annotation query
 * TS type is CloudWatchDefaultQuery = Omit<CloudWatchLogsQuery, 'queryMode'> & CloudWatchMetricsQuery, declared in veneer
 * #CloudWatchDefaultQuery: #CloudWatchLogsQuery & #CloudWatchMetricsQuery @cuetsy(kind="type")
 */
class CloudWatchAnnotationQuery implements \JsonSerializable, \Grafana\Foundation\Cog\Dataquery
{
    /**
     * Whether a query is a Metrics, Logs, or Annotations query
     */
    public \Grafana\Foundation\Cloudwatch\CloudWatchQueryMode $queryMode;

    /**
     * Enable matching on the prefix of the action name or alarm name, specify the prefixes with actionPrefix and/or alarmNamePrefix
     */
    public ?bool $prefixMatching;

    /**
     * Use this parameter to filter the results of the operation to only those alarms
     * that use a certain alarm action. For example, you could specify the ARN of
     * an SNS topic to find all alarms that send notifications to that topic.
     * e.g. `arn:aws:sns:us-east-1:123456789012:my-app-` would match `arn:aws:sns:us-east-1:123456789012:my-app-action`
     * but not match `arn:aws:sns:us-east-1:123456789012:your-app-action`
     */
    public ?string $actionPrefix;

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
     * An alarm name prefix. If you specify this parameter, you receive information
     * about all alarms that have names that start with this prefix.
     * e.g. `my-team-service-` would match `my-team-service-high-cpu` but not match `your-team-service-high-cpu`
     */
    public ?string $alarmNamePrefix;

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
     * @param bool|null $prefixMatching
     * @param string|null $actionPrefix
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
     * @param string|null $alarmNamePrefix
     * @param \Grafana\Foundation\Dashboard\DataSourceRef|null $datasource
     * @param array<string>|null $statistics
     */
    public function __construct(?\Grafana\Foundation\Cloudwatch\CloudWatchQueryMode $queryMode = null, ?bool $prefixMatching = null, ?string $actionPrefix = null, ?string $refId = null, ?bool $hide = null, ?string $queryType = null, ?string $region = null, ?string $namespace = null, ?string $metricName = null, ?array $dimensions = null, ?bool $matchExact = null, ?string $period = null, ?string $accountId = null, ?string $statistic = null, ?string $alarmNamePrefix = null, ?\Grafana\Foundation\Dashboard\DataSourceRef $datasource = null, ?array $statistics = null)
    {
        $this->queryMode = $queryMode ?: \Grafana\Foundation\Cloudwatch\CloudWatchQueryMode::annotations();
        $this->prefixMatching = $prefixMatching;
        $this->actionPrefix = $actionPrefix;
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
        $this->alarmNamePrefix = $alarmNamePrefix;
        $this->datasource = $datasource;
        $this->statistics = $statistics;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{queryMode?: string, prefixMatching?: bool, actionPrefix?: string, refId?: string, hide?: bool, queryType?: string, region?: string, namespace?: string, metricName?: string, dimensions?: array<string, string|array<string>>, matchExact?: bool, period?: string, accountId?: string, statistic?: string, alarmNamePrefix?: string, datasource?: mixed, statistics?: array<string>} $inputData */
        $data = $inputData;
        return new self(
            queryMode: isset($data["queryMode"]) ? (function($input) { return \Grafana\Foundation\Cloudwatch\CloudWatchQueryMode::fromValue($input); })($data["queryMode"]) : null,
            prefixMatching: $data["prefixMatching"] ?? null,
            actionPrefix: $data["actionPrefix"] ?? null,
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
            alarmNamePrefix: $data["alarmNamePrefix"] ?? null,
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
            "refId" => $this->refId,
            "region" => $this->region,
            "namespace" => $this->namespace,
            "dimensions" => $this->dimensions,
        ];
        if (isset($this->prefixMatching)) {
            $data["prefixMatching"] = $this->prefixMatching;
        }
        if (isset($this->actionPrefix)) {
            $data["actionPrefix"] = $this->actionPrefix;
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
        if (isset($this->alarmNamePrefix)) {
            $data["alarmNamePrefix"] = $this->alarmNamePrefix;
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
