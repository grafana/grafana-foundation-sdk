<?php

namespace Grafana\Foundation\Common;

/**
 * TODO docs
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\LineConfig>
 */
class LineConfigBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Common\LineConfig $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Common\LineConfig();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Common\LineConfig
     */
    public function build()
    {
        return $this->internal;
    }

    public function lineColor(string $lineColor): static
    {
        $this->internal->lineColor = $lineColor;
    
        return $this;
    }
    public function lineWidth(float $lineWidth): static
    {
        $this->internal->lineWidth = $lineWidth;
    
        return $this;
    }
    public function lineInterpolation(\Grafana\Foundation\Common\LineInterpolation $lineInterpolation): static
    {
        $this->internal->lineInterpolation = $lineInterpolation;
    
        return $this;
    }
    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\LineStyle> $lineStyle
     */
    public function lineStyle(\Grafana\Foundation\Cog\Builder $lineStyle): static
    {
        $lineStyleResource = $lineStyle->build();
        $this->internal->lineStyle = $lineStyleResource;
    
        return $this;
    }
    /**
     * Indicate if null values should be treated as gaps or connected.
     * When the value is a number, it represents the maximum delta in the
     * X axis that should be considered connected.  For timeseries, this is milliseconds
     * @param bool|float $spanNulls
     */
    public function spanNulls( $spanNulls): static
    {
        $this->internal->spanNulls = $spanNulls;
    
        return $this;
    }

}
