<?php

namespace Grafana\Foundation\Trend;

/**
 * Identical to timeseries... except it does not have timezone settings
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Trend\Options>
 */
class OptionsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Trend\Options $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Trend\Options();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Trend\Options
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\VizLegendOptions> $legend
     */
    public function legend(\Grafana\Foundation\Cog\Builder $legend): static
    {
        $legendResource = $legend->build();
        $this->internal->legend = $legendResource;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\VizTooltipOptions> $tooltip
     */
    public function tooltip(\Grafana\Foundation\Cog\Builder $tooltip): static
    {
        $tooltipResource = $tooltip->build();
        $this->internal->tooltip = $tooltipResource;
    
        return $this;
    }

    /**
     * Name of the x field to use (defaults to first number)
     */
    public function xField(string $xField): static
    {
        $this->internal->xField = $xField;
    
        return $this;
    }

}
