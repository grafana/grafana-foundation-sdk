<?php

namespace Grafana\Foundation\Common;

/**
 * TODO docs
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\VizTooltipOptions>
 */
class VizTooltipOptionsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Common\VizTooltipOptions $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Common\VizTooltipOptions();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Common\VizTooltipOptions
     */
    public function build()
    {
        return $this->internal;
    }

    public function mode(\Grafana\Foundation\Common\TooltipDisplayMode $mode): static
    {
        $this->internal->mode = $mode;
    
        return $this;
    }
    public function sort(\Grafana\Foundation\Common\SortOrder $sort): static
    {
        $this->internal->sort = $sort;
    
        return $this;
    }

}
