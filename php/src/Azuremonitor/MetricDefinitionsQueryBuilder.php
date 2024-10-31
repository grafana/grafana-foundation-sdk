<?php

namespace Grafana\Foundation\Azuremonitor;

/**
 * @deprecated Use MetricNamespaceQuery instead
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\MetricDefinitionsQuery>
 */
class MetricDefinitionsQueryBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Azuremonitor\MetricDefinitionsQuery $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Azuremonitor\MetricDefinitionsQuery();
    $this->internal->kind = "MetricDefinitionsQuery";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Azuremonitor\MetricDefinitionsQuery
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
