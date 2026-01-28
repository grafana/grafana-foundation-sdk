<?php

namespace Grafana\Foundation\Gauge;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Gauge\Options>
 */
class OptionsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Gauge\Options $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Gauge\Options();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Gauge\Options
     */
    public function build()
    {
        return $this->internal;
    }

    public function showThresholdLabels(bool $showThresholdLabels): static
    {
        $this->internal->showThresholdLabels = $showThresholdLabels;
    
        return $this;
    }

    public function showThresholdMarkers(bool $showThresholdMarkers): static
    {
        $this->internal->showThresholdMarkers = $showThresholdMarkers;
    
        return $this;
    }

    public function sizing(\Grafana\Foundation\Common\BarGaugeSizing $sizing): static
    {
        $this->internal->sizing = $sizing;
    
        return $this;
    }

    public function minVizWidth(int $minVizWidth): static
    {
        $this->internal->minVizWidth = $minVizWidth;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\ReduceDataOptions> $reduceOptions
     */
    public function reduceOptions(\Grafana\Foundation\Cog\Builder $reduceOptions): static
    {
        $reduceOptionsResource = $reduceOptions->build();
        $this->internal->reduceOptions = $reduceOptionsResource;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\VizTextDisplayOptions> $text
     */
    public function text(\Grafana\Foundation\Cog\Builder $text): static
    {
        $textResource = $text->build();
        $this->internal->text = $textResource;
    
        return $this;
    }

    public function minVizHeight(int $minVizHeight): static
    {
        $this->internal->minVizHeight = $minVizHeight;
    
        return $this;
    }

    public function orientation(\Grafana\Foundation\Common\VizOrientation $orientation): static
    {
        $this->internal->orientation = $orientation;
    
        return $this;
    }

}
