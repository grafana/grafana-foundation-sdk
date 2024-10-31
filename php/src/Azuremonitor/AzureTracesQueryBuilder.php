<?php

namespace Grafana\Foundation\Azuremonitor;

/**
 * Application Insights Traces sub-query properties
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\AzureTracesQuery>
 */
class AzureTracesQueryBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Azuremonitor\AzureTracesQuery $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Azuremonitor\AzureTracesQuery();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Azuremonitor\AzureTracesQuery
     */
    public function build()
    {
        return $this->internal;
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
     * Operation ID. Used only for Traces queries.
     */
    public function operationId(string $operationId): static
    {
        $this->internal->operationId = $operationId;
    
        return $this;
    }
    /**
     * Types of events to filter by.
     * @param array<string> $traceTypes
     */
    public function traceTypes(array $traceTypes): static
    {
        $this->internal->traceTypes = $traceTypes;
    
        return $this;
    }
    /**
     * Filters for property values.
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\AzureTracesFilter>> $filters
     */
    public function filters(array $filters): static
    {
            $filtersResources = [];
            foreach ($filters as $r1) {
                    $filtersResources[] = $r1->build();
            }
        $this->internal->filters = $filtersResources;
    
        return $this;
    }
    /**
     * KQL query to be executed.
     */
    public function query(string $query): static
    {
        $this->internal->query = $query;
    
        return $this;
    }

}
