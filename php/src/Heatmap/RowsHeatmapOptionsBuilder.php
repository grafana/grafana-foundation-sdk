<?php

namespace Grafana\Foundation\Heatmap;

/**
 * Controls frame rows options
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Heatmap\RowsHeatmapOptions>
 */
class RowsHeatmapOptionsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Heatmap\RowsHeatmapOptions $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Heatmap\RowsHeatmapOptions();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Heatmap\RowsHeatmapOptions
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Sets the name of the cell when not calculating from data
     */
    public function value(string $value): static
    {
        $this->internal->value = $value;
    
        return $this;
    }
    /**
     * Controls tick alignment when not calculating from data
     */
    public function layout(\Grafana\Foundation\Common\HeatmapCellLayout $layout): static
    {
        $this->internal->layout = $layout;
    
        return $this;
    }

}
