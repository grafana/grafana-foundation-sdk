<?php

namespace Grafana\Foundation\Canvas;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Canvas\LineConfig>
 */
class LineConfigBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Canvas\LineConfig $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Canvas\LineConfig();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Canvas\LineConfig
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\ColorDimensionConfig> $color
     */
    public function color(\Grafana\Foundation\Cog\Builder $color): static
    {
        $colorResource = $color->build();
        $this->internal->color = $colorResource;
    
        return $this;
    }
    public function width(float $width): static
    {
        $this->internal->width = $width;
    
        return $this;
    }

}
