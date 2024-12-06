<?php

namespace Grafana\Foundation\Common;

/**
 * TODO docs
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\VizTextDisplayOptions>
 */
class VizTextDisplayOptionsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Common\VizTextDisplayOptions $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Common\VizTextDisplayOptions();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Common\VizTextDisplayOptions
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Explicit title text size
     */
    public function titleSize(float $titleSize): static
    {
        $this->internal->titleSize = $titleSize;
    
        return $this;
    }
    /**
     * Explicit value text size
     */
    public function valueSize(float $valueSize): static
    {
        $this->internal->valueSize = $valueSize;
    
        return $this;
    }

}
