<?php

namespace Grafana\Foundation\Common;

/**
 * TODO docs
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\SingleStatBaseOptions>
 */
class SingleStatBaseOptionsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Common\SingleStatBaseOptions $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Common\SingleStatBaseOptions();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Common\SingleStatBaseOptions
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\ReduceDataOptions> $reduceOptions
     */
    public function reduceOptions(\Grafana\Foundation\Cog\Builder $reduceOptions): static
    {
        $reduceOptionsResource = $reduceOptions->build();
        $this->internal->reduceOptions = $reduceOptionsResource;
    
        return $this;
    }
    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\VizTextDisplayOptions> $text
     */
    public function text(\Grafana\Foundation\Cog\Builder $text): static
    {
        $textResource = $text->build();
        $this->internal->text = $textResource;
    
        return $this;
    }
    public function orientation(\Grafana\Foundation\Common\VizOrientation $orientation): static
    {
        $this->internal->orientation = $orientation;
    
        return $this;
    }

}
