<?php

namespace Grafana\Foundation\Common;

/**
 * TODO docs
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\OptionsWithLegend>
 */
class OptionsWithLegendBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Common\OptionsWithLegend $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Common\OptionsWithLegend();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Common\OptionsWithLegend
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

}
