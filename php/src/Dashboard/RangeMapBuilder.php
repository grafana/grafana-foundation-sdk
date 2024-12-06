<?php

namespace Grafana\Foundation\Dashboard;

/**
 * Maps numerical ranges to a display text and color.
 * For example, if a value is within a certain range, you can configure a range value mapping to display Low or High rather than the number.
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\RangeMap>
 */
class RangeMapBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboard\RangeMap $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboard\RangeMap();
    $this->internal->type = "range";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboard\RangeMap
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Range to match against and the result to apply when the value is within the range
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\DashboardRangeMapOptions> $options
     */
    public function options(\Grafana\Foundation\Cog\Builder $options): static
    {
        $optionsResource = $options->build();
        $this->internal->options = $optionsResource;
    
        return $this;
    }

}
