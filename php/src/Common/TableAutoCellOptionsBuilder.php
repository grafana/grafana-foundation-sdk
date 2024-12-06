<?php

namespace Grafana\Foundation\Common;

/**
 * Auto mode table cell options
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\TableAutoCellOptions>
 */
class TableAutoCellOptionsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Common\TableAutoCellOptions $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Common\TableAutoCellOptions();
    $this->internal->type = "auto";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Common\TableAutoCellOptions
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
