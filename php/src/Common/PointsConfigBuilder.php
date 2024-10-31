<?php

namespace Grafana\Foundation\Common;

/**
 * TODO docs
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\PointsConfig>
 */
class PointsConfigBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Common\PointsConfig $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Common\PointsConfig();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Common\PointsConfig
     */
    public function build()
    {
        return $this->internal;
    }

    public function showPoints(\Grafana\Foundation\Common\VisibilityMode $showPoints): static
    {
        $this->internal->showPoints = $showPoints;
    
        return $this;
    }
    public function pointSize(float $pointSize): static
    {
        $this->internal->pointSize = $pointSize;
    
        return $this;
    }
    public function pointColor(string $pointColor): static
    {
        $this->internal->pointColor = $pointColor;
    
        return $this;
    }
    public function pointSymbol(string $pointSymbol): static
    {
        $this->internal->pointSymbol = $pointSymbol;
    
        return $this;
    }

}
