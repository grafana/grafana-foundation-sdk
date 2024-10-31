<?php

namespace Grafana\Foundation\Heatmap;

/**
 * Controls legend options
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Heatmap\HeatmapLegend>
 */
class HeatmapLegendBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Heatmap\HeatmapLegend $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Heatmap\HeatmapLegend();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Heatmap\HeatmapLegend
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Controls if the legend is shown
     */
    public function show(bool $show): static
    {
        $this->internal->show = $show;
    
        return $this;
    }

}
