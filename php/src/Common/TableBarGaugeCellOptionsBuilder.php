<?php

namespace Grafana\Foundation\Common;

/**
 * Gauge cell options
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\TableBarGaugeCellOptions>
 */
class TableBarGaugeCellOptionsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Common\TableBarGaugeCellOptions $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Common\TableBarGaugeCellOptions();
    $this->internal->type = "gauge";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Common\TableBarGaugeCellOptions
     */
    public function build()
    {
        return $this->internal;
    }

    public function mode(\Grafana\Foundation\Common\BarGaugeDisplayMode $mode): static
    {
        $this->internal->mode = $mode;
    
        return $this;
    }
    public function valueDisplayMode(\Grafana\Foundation\Common\BarGaugeValueMode $valueDisplayMode): static
    {
        $this->internal->valueDisplayMode = $valueDisplayMode;
    
        return $this;
    }

}
