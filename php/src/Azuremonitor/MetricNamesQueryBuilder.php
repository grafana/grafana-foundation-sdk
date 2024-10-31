<?php

namespace Grafana\Foundation\Azuremonitor;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\MetricNamesQuery>
 */
class MetricNamesQueryBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Azuremonitor\MetricNamesQuery $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Azuremonitor\MetricNamesQuery();
    $this->internal->kind = "MetricNamesQuery";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Azuremonitor\MetricNamesQuery
     */
    public function build()
    {
        return $this->internal;
    }

    public function rawQuery(string $rawQuery): static
    {
        $this->internal->rawQuery = $rawQuery;
    
        return $this;
    }
    public function subscription(string $subscription): static
    {
        $this->internal->subscription = $subscription;
    
        return $this;
    }
    public function resourceGroup(string $resourceGroup): static
    {
        $this->internal->resourceGroup = $resourceGroup;
    
        return $this;
    }
    public function resourceName(string $resourceName): static
    {
        $this->internal->resourceName = $resourceName;
    
        return $this;
    }
    public function metricNamespace(string $metricNamespace): static
    {
        $this->internal->metricNamespace = $metricNamespace;
    
        return $this;
    }

}
