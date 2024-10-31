<?php

namespace Grafana\Foundation\Common;

/**
 * Colored text cell options
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\TableColorTextCellOptions>
 */
class TableColorTextCellOptionsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Common\TableColorTextCellOptions $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Common\TableColorTextCellOptions();
    $this->internal->type = "color-text";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Common\TableColorTextCellOptions
     */
    public function build()
    {
        return $this->internal;
    }

    public function wrapText(bool $wrapText): static
    {
        $this->internal->wrapText = $wrapText;
    
        return $this;
    }

}
