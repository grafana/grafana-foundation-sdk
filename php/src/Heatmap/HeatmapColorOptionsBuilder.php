<?php

namespace Grafana\Foundation\Heatmap;

/**
 * Controls various color options
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Heatmap\HeatmapColorOptions>
 */
class HeatmapColorOptionsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Heatmap\HeatmapColorOptions $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Heatmap\HeatmapColorOptions();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Heatmap\HeatmapColorOptions
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Sets the color mode
     */
    public function mode(\Grafana\Foundation\Heatmap\HeatmapColorMode $mode): static
    {
        $this->internal->mode = $mode;
    
        return $this;
    }
    /**
     * Controls the color scheme used
     */
    public function scheme(string $scheme): static
    {
        $this->internal->scheme = $scheme;
    
        return $this;
    }
    /**
     * Controls the color fill when in opacity mode
     */
    public function fill(string $fill): static
    {
        $this->internal->fill = $fill;
    
        return $this;
    }
    /**
     * Controls the color scale
     */
    public function scale(\Grafana\Foundation\Heatmap\HeatmapColorScale $scale): static
    {
        $this->internal->scale = $scale;
    
        return $this;
    }
    /**
     * Controls the exponent when scale is set to exponential
     */
    public function exponent(float $exponent): static
    {
        $this->internal->exponent = $exponent;
    
        return $this;
    }
    /**
     * Controls the number of color steps
     */
    public function steps(int $steps): static
    {
        if (!($steps >= 2)) {
            throw new \ValueError('$steps must be >= 2');
        }
        if (!($steps <= 128)) {
            throw new \ValueError('$steps must be <= 128');
        }
        $this->internal->steps = $steps;
    
        return $this;
    }
    /**
     * Reverses the color scheme
     */
    public function reverse(bool $reverse): static
    {
        $this->internal->reverse = $reverse;
    
        return $this;
    }
    /**
     * Sets the minimum value for the color scale
     */
    public function min(float $min): static
    {
        $this->internal->min = $min;
    
        return $this;
    }
    /**
     * Sets the maximum value for the color scale
     */
    public function max(float $max): static
    {
        $this->internal->max = $max;
    
        return $this;
    }

}
