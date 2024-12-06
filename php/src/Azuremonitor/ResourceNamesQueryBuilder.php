<?php

namespace Grafana\Foundation\Azuremonitor;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\ResourceNamesQuery>
 */
class ResourceNamesQueryBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Azuremonitor\ResourceNamesQuery $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Azuremonitor\ResourceNamesQuery();
    $this->internal->kind = "ResourceNamesQuery";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Azuremonitor\ResourceNamesQuery
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

}
