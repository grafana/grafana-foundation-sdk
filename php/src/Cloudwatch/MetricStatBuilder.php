<?php

namespace Grafana\Foundation\Cloudwatch;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Cloudwatch\MetricStat>
 */
class MetricStatBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Cloudwatch\MetricStat $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Cloudwatch\MetricStat();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Cloudwatch\MetricStat
     */
    public function build()
    {
        return $this->internal;
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
     * @deprecated use statistic
     * @param array<string> $statistics
     */
    public function statistics(array $statistics): static
    {
        $this->internal->statistics = $statistics;
    
        return $this;
    }

}
