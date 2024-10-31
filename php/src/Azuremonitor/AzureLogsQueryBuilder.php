<?php

namespace Grafana\Foundation\Azuremonitor;

/**
 * Azure Monitor Logs sub-query properties
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\AzureLogsQuery>
 */
class AzureLogsQueryBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Azuremonitor\AzureLogsQuery $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Azuremonitor\AzureLogsQuery();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Azuremonitor\AzureLogsQuery
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * KQL query to be executed.
     */
    public function query(string $query): static
    {
        $this->internal->query = $query;
    
        return $this;
    }
    /**
     * Specifies the format results should be returned as.
     */
    public function resultFormat(\Grafana\Foundation\Azuremonitor\ResultFormat $resultFormat): static
    {
        $this->internal->resultFormat = $resultFormat;
    
        return $this;
    }
    /**
     * Array of resource URIs to be queried.
     * @param array<string> $resources
     */
    public function resources(array $resources): static
    {
        $this->internal->resources = $resources;
    
        return $this;
    }
    /**
     * If set to true the intersection of time ranges specified in the query and Grafana will be used. Otherwise the query time ranges will be used. Defaults to false
     */
    public function intersectTime(bool $intersectTime): static
    {
        $this->internal->intersectTime = $intersectTime;
    
        return $this;
    }
    /**
     * Workspace ID. This was removed in Grafana 8, but remains for backwards compat
     */
    public function workspace(string $workspace): static
    {
        $this->internal->workspace = $workspace;
    
        return $this;
    }
    /**
     * @deprecated Use resources instead
     */
    public function resource(string $resource): static
    {
        $this->internal->resource = $resource;
    
        return $this;
    }

}
