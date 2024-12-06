<?php

namespace Grafana\Foundation\Heatmap;

/**
 * Controls cell value options
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Heatmap\CellValues>
 */
class CellValuesBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Heatmap\CellValues $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Heatmap\CellValues();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Heatmap\CellValues
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Controls the cell value unit
     */
    public function unit(string $unit): static
    {
        $this->internal->unit = $unit;
    
        return $this;
    }
    /**
     * Controls the number of decimals for cell values
     */
    public function decimals(float $decimals): static
    {
        $this->internal->decimals = $decimals;
    
        return $this;
    }

}
