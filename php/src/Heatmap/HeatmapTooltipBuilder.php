<?php

namespace Grafana\Foundation\Heatmap;

/**
 * Controls tooltip options
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Heatmap\HeatmapTooltip>
 */
class HeatmapTooltipBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Heatmap\HeatmapTooltip $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Heatmap\HeatmapTooltip();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Heatmap\HeatmapTooltip
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Controls how the tooltip is shown
     */
    public function mode(\Grafana\Foundation\Common\TooltipDisplayMode $mode): static
    {
        $this->internal->mode = $mode;
    
        return $this;
    }
    public function maxHeight(float $maxHeight): static
    {
        $this->internal->maxHeight = $maxHeight;
    
        return $this;
    }
    public function maxWidth(float $maxWidth): static
    {
        $this->internal->maxWidth = $maxWidth;
    
        return $this;
    }
    /**
     * Controls if the tooltip shows a histogram of the y-axis values
     */
    public function yHistogram(bool $yHistogram): static
    {
        $this->internal->yHistogram = $yHistogram;
    
        return $this;
    }
    /**
     * Controls if the tooltip shows a color scale in header
     */
    public function showColorScale(bool $showColorScale): static
    {
        $this->internal->showColorScale = $showColorScale;
    
        return $this;
    }

}
