<?php

namespace Grafana\Foundation\Cloudwatch;

class MetricStat implements \JsonSerializable
{
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
     * @deprecated use statistic
     * @var array<string>|null
     */
    public ?array $statistics;

    /**
     * @param string|null $region
     * @param string|null $namespace
     * @param string|null $metricName
     * @param array<string, string|array<string>>|null $dimensions
     * @param bool|null $matchExact
     * @param string|null $period
     * @param string|null $accountId
     * @param string|null $statistic
     * @param array<string>|null $statistics
     */
    public function __construct(?string $region = null, ?string $namespace = null, ?string $metricName = null, ?array $dimensions = null, ?bool $matchExact = null, ?string $period = null, ?string $accountId = null, ?string $statistic = null, ?array $statistics = null)
    {
        $this->region = $region ?: "";
        $this->namespace = $namespace ?: "";
        $this->metricName = $metricName;
        $this->dimensions = $dimensions ?: [];
        $this->matchExact = $matchExact;
        $this->period = $period;
        $this->accountId = $accountId;
        $this->statistic = $statistic;
        $this->statistics = $statistics;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{region?: string, namespace?: string, metricName?: string, dimensions?: array<string, string|array<string>>, matchExact?: bool, period?: string, accountId?: string, statistic?: string, statistics?: array<string>} $inputData */
        $data = $inputData;
        return new self(
            region: $data["region"] ?? null,
            namespace: $data["namespace"] ?? null,
            metricName: $data["metricName"] ?? null,
            dimensions: $data["dimensions"] ?? null,
            matchExact: $data["matchExact"] ?? null,
            period: $data["period"] ?? null,
            accountId: $data["accountId"] ?? null,
            statistic: $data["statistic"] ?? null,
            statistics: $data["statistics"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "region" => $this->region,
            "namespace" => $this->namespace,
            "dimensions" => $this->dimensions,
        ];
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
        if (isset($this->statistics)) {
            $data["statistics"] = $this->statistics;
        }
        return $data;
    }
}
