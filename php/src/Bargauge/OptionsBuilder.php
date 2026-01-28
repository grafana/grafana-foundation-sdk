<?php

namespace Grafana\Foundation\Bargauge;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Bargauge\Options>
 */
class OptionsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Bargauge\Options $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Bargauge\Options();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Bargauge\Options
     */
    public function build()
    {
        return $this->internal;
    }

    public function displayMode(\Grafana\Foundation\Common\BarGaugeDisplayMode $displayMode): static
    {
        $this->internal->displayMode = $displayMode;
    
        return $this;
    }

    public function valueMode(\Grafana\Foundation\Common\BarGaugeValueMode $valueMode): static
    {
        $this->internal->valueMode = $valueMode;
    
        return $this;
    }

    public function namePlacement(\Grafana\Foundation\Common\BarGaugeNamePlacement $namePlacement): static
    {
        $this->internal->namePlacement = $namePlacement;
    
        return $this;
    }

    public function showUnfilled(bool $showUnfilled): static
    {
        $this->internal->showUnfilled = $showUnfilled;
    
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

    public function minVizHeight(int $minVizHeight): static
    {
        $this->internal->minVizHeight = $minVizHeight;
    
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

    public function maxVizHeight(int $maxVizHeight): static
    {
        $this->internal->maxVizHeight = $maxVizHeight;
    
        return $this;
    }

    public function orientation(\Grafana\Foundation\Common\VizOrientation $orientation): static
    {
        $this->internal->orientation = $orientation;
    
        return $this;
    }

}
