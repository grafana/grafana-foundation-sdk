<?php

namespace Grafana\Foundation\Azuremonitor;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\AzureMetricDimension>
 */
class AzureMetricDimensionBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Azuremonitor\AzureMetricDimension $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Azuremonitor\AzureMetricDimension();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Azuremonitor\AzureMetricDimension
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Name of Dimension to be filtered on.
     */
    public function dimension(string $dimension): static
    {
        $this->internal->dimension = $dimension;
    
        return $this;
    }
    /**
     * String denoting the filter operation. Supports 'eq' - equals,'ne' - not equals, 'sw' - starts with. Note that some dimensions may not support all operators.
     */
    public function operator(string $operator): static
    {
        $this->internal->operator = $operator;
    
        return $this;
    }
    /**
     * Values to match with the filter.
     * @param array<string> $filters
     */
    public function filters(array $filters): static
    {
        $this->internal->filters = $filters;
    
        return $this;
    }
    /**
     * @deprecated filter is deprecated in favour of filters to support multiselect.
     */
    public function filter(string $filter): static
    {
        $this->internal->filter = $filter;
    
        return $this;
    }

}
