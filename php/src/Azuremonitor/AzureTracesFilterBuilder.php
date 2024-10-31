<?php

namespace Grafana\Foundation\Azuremonitor;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\AzureTracesFilter>
 */
class AzureTracesFilterBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Azuremonitor\AzureTracesFilter $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Azuremonitor\AzureTracesFilter();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Azuremonitor\AzureTracesFilter
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Property name, auto-populated based on available traces.
     */
    public function property(string $property): static
    {
        $this->internal->property = $property;
    
        return $this;
    }
    /**
     * Comparison operator to use. Either equals or not equals.
     */
    public function operation(string $operation): static
    {
        $this->internal->operation = $operation;
    
        return $this;
    }
    /**
     * Values to filter by.
     * @param array<string> $filters
     */
    public function filters(array $filters): static
    {
        $this->internal->filters = $filters;
    
        return $this;
    }

}
