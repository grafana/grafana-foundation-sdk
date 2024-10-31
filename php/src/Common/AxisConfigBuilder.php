<?php

namespace Grafana\Foundation\Common;

/**
 * TODO docs
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\AxisConfig>
 */
class AxisConfigBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Common\AxisConfig $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Common\AxisConfig();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Common\AxisConfig
     */
    public function build()
    {
        return $this->internal;
    }

    public function axisPlacement(\Grafana\Foundation\Common\AxisPlacement $axisPlacement): static
    {
        $this->internal->axisPlacement = $axisPlacement;
    
        return $this;
    }
    public function axisColorMode(\Grafana\Foundation\Common\AxisColorMode $axisColorMode): static
    {
        $this->internal->axisColorMode = $axisColorMode;
    
        return $this;
    }
    public function axisLabel(string $axisLabel): static
    {
        $this->internal->axisLabel = $axisLabel;
    
        return $this;
    }
    public function axisWidth(float $axisWidth): static
    {
        $this->internal->axisWidth = $axisWidth;
    
        return $this;
    }
    public function axisSoftMin(float $axisSoftMin): static
    {
        $this->internal->axisSoftMin = $axisSoftMin;
    
        return $this;
    }
    public function axisSoftMax(float $axisSoftMax): static
    {
        $this->internal->axisSoftMax = $axisSoftMax;
    
        return $this;
    }
    public function axisGridShow(bool $axisGridShow): static
    {
        $this->internal->axisGridShow = $axisGridShow;
    
        return $this;
    }
    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\ScaleDistributionConfig> $scaleDistribution
     */
    public function scaleDistribution(\Grafana\Foundation\Cog\Builder $scaleDistribution): static
    {
        $scaleDistributionResource = $scaleDistribution->build();
        $this->internal->scaleDistribution = $scaleDistributionResource;
    
        return $this;
    }
    public function axisCenteredZero(bool $axisCenteredZero): static
    {
        $this->internal->axisCenteredZero = $axisCenteredZero;
    
        return $this;
    }

}
