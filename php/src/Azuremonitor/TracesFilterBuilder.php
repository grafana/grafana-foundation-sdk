<?php

namespace Grafana\Foundation\Azuremonitor;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\TracesFilter>
 */
class TracesFilterBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Azuremonitor\TracesFilter $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Azuremonitor\TracesFilter();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Azuremonitor\TracesFilter
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
