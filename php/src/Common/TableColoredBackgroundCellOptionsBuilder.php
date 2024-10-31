<?php

namespace Grafana\Foundation\Common;

/**
 * Colored background cell options
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\TableColoredBackgroundCellOptions>
 */
class TableColoredBackgroundCellOptionsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Common\TableColoredBackgroundCellOptions $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Common\TableColoredBackgroundCellOptions();
    $this->internal->type = "color-background";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Common\TableColoredBackgroundCellOptions
     */
    public function build()
    {
        return $this->internal;
    }

    public function mode(\Grafana\Foundation\Common\TableCellBackgroundDisplayMode $mode): static
    {
        $this->internal->mode = $mode;
    
        return $this;
    }

}
