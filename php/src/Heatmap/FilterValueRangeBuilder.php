<?php

namespace Grafana\Foundation\Heatmap;

/**
 * Controls the value filter range
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Heatmap\FilterValueRange>
 */
class FilterValueRangeBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Heatmap\FilterValueRange $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Heatmap\FilterValueRange();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Heatmap\FilterValueRange
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Sets the filter range to values less than or equal to the given value
     */
    public function le(float $le): static
    {
        $this->internal->le = $le;
    
        return $this;
    }
    /**
     * Sets the filter range to values greater than or equal to the given value
     */
    public function ge(float $ge): static
    {
        $this->internal->ge = $ge;
    
        return $this;
    }

}
