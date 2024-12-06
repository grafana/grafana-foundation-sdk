<?php

namespace Grafana\Foundation\Common;

/**
 * Json view cell options
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\TableImageCellOptions>
 */
class TableImageCellOptionsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Common\TableImageCellOptions $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Common\TableImageCellOptions();
    $this->internal->type = "image";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Common\TableImageCellOptions
     */
    public function build()
    {
        return $this->internal;
    }

    public function alt(string $alt): static
    {
        $this->internal->alt = $alt;
    
        return $this;
    }
    public function title(string $title): static
    {
        $this->internal->title = $title;
    
        return $this;
    }

}
