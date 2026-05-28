<?php

namespace Grafana\Foundation\Dashboardv2;

/**
 * Define the MetricFindValue type
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\MetricFindValue>
 */
class MetricFindValueBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2\MetricFindValue $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2\MetricFindValue();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2\MetricFindValue
     */
    public function build()
    {
        return $this->internal;
    }

    public function text(string $text): static
    {
        $this->internal->text = $text;
    
        return $this;
    }

    /**
     * @param string|float $value
     */
    public function value( $value): static
    {
        $this->internal->value = $value;
    
        return $this;
    }

    public function group(string $group): static
    {
        $this->internal->group = $group;
    
        return $this;
    }

    public function expandable(bool $expandable): static
    {
        $this->internal->expandable = $expandable;
    
        return $this;
    }

}
