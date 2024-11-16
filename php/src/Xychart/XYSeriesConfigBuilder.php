<?php

namespace Grafana\Foundation\Xychart;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Xychart\XYSeriesConfig>
 */
class XYSeriesConfigBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Xychart\XYSeriesConfig $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Xychart\XYSeriesConfig();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Xychart\XYSeriesConfig
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Xychart\XychartXYSeriesConfigName> $name
     */
    public function name(\Grafana\Foundation\Cog\Builder $name): static
    {
        $nameResource = $name->build();
        $this->internal->name = $nameResource;
    
        return $this;
    }
    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Xychart\XychartXYSeriesConfigFrame> $frame
     */
    public function frame(\Grafana\Foundation\Cog\Builder $frame): static
    {
        $frameResource = $frame->build();
        $this->internal->frame = $frameResource;
    
        return $this;
    }
    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Xychart\XychartXYSeriesConfigX> $x
     */
    public function x(\Grafana\Foundation\Cog\Builder $x): static
    {
        $xResource = $x->build();
        $this->internal->x = $xResource;
    
        return $this;
    }
    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Xychart\XychartXYSeriesConfigY> $y
     */
    public function y(\Grafana\Foundation\Cog\Builder $y): static
    {
        $yResource = $y->build();
        $this->internal->y = $yResource;
    
        return $this;
    }
    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Xychart\XychartXYSeriesConfigColor> $color
     */
    public function color(\Grafana\Foundation\Cog\Builder $color): static
    {
        $colorResource = $color->build();
        $this->internal->color = $colorResource;
    
        return $this;
    }
    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Xychart\XychartXYSeriesConfigSize> $size
     */
    public function size(\Grafana\Foundation\Cog\Builder $size): static
    {
        $sizeResource = $size->build();
        $this->internal->size = $sizeResource;
    
        return $this;
    }

}
