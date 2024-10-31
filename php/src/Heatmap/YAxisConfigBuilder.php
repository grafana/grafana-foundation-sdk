<?php

namespace Grafana\Foundation\Heatmap;

/**
 * Configuration options for the yAxis
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Heatmap\YAxisConfig>
 */
class YAxisConfigBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Heatmap\YAxisConfig $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Heatmap\YAxisConfig();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Heatmap\YAxisConfig
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Sets the yAxis unit
     */
    public function unit(string $unit): static
    {
        $this->internal->unit = $unit;
    
        return $this;
    }
    /**
     * Reverses the yAxis
     */
    public function reverse(bool $reverse): static
    {
        $this->internal->reverse = $reverse;
    
        return $this;
    }
    /**
     * Controls the number of decimals for yAxis values
     */
    public function decimals(float $decimals): static
    {
        $this->internal->decimals = $decimals;
    
        return $this;
    }
    /**
     * Sets the minimum value for the yAxis
     */
    public function min(float $min): static
    {
        $this->internal->min = $min;
    
        return $this;
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
    /**
     * Sets the maximum value for the yAxis
     */
    public function max(float $max): static
    {
        $this->internal->max = $max;
    
        return $this;
    }
    public function axisBorderShow(bool $axisBorderShow): static
    {
        $this->internal->axisBorderShow = $axisBorderShow;
    
        return $this;
    }

}
