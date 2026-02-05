<?php

namespace Grafana\Foundation\Common;

/**
 * TODO docs
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\OptionsWithTooltip>
 */
class OptionsWithTooltipBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Common\OptionsWithTooltip $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Common\OptionsWithTooltip();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Common\OptionsWithTooltip
     */
    public function build()
    {
        return $this->internal;
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

}
