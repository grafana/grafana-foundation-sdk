<?php

namespace Grafana\Foundation\Azuremonitor;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\AzureMonitorResource>
 */
class AzureMonitorResourceBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Azuremonitor\AzureMonitorResource $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Azuremonitor\AzureMonitorResource();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Azuremonitor\AzureMonitorResource
     */
    public function build()
    {
        return $this->internal;
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
    public function region(string $region): static
    {
        $this->internal->region = $region;
    
        return $this;
    }

}
