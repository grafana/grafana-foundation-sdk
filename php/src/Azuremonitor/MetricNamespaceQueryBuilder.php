<?php

namespace Grafana\Foundation\Azuremonitor;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\MetricNamespaceQuery>
 */
class MetricNamespaceQueryBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Azuremonitor\MetricNamespaceQuery $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Azuremonitor\MetricNamespaceQuery();
    $this->internal->kind = "MetricNamespaceQuery";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Azuremonitor\MetricNamespaceQuery
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
    public function metricNamespace(string $metricNamespace): static
    {
        $this->internal->metricNamespace = $metricNamespace;
    
        return $this;
    }
    public function resourceName(string $resourceName): static
    {
        $this->internal->resourceName = $resourceName;
    
        return $this;
    }

}
