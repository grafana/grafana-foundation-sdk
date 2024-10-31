<?php

namespace Grafana\Foundation\Geomap;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Geomap\TooltipOptions>
 */
class TooltipOptionsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Geomap\TooltipOptions $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Geomap\TooltipOptions();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Geomap\TooltipOptions
     */
    public function build()
    {
        return $this->internal;
    }

    public function mode(\Grafana\Foundation\Geomap\TooltipMode $mode): static
    {
        $this->internal->mode = $mode;
    
        return $this;
    }

}
