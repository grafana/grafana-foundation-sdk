<?php

namespace Grafana\Foundation\Timeseries;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Timeseries\Options>
 */
class OptionsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Timeseries\Options $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Timeseries\Options();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Timeseries\Options
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * @param array<string> $timezone
     */
    public function timezone(array $timezone): static
    {
        $this->internal->timezone = $timezone;
    
        return $this;
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

    public function orientation(\Grafana\Foundation\Common\VizOrientation $orientation): static
    {
        $this->internal->orientation = $orientation;
    
        return $this;
    }

}
