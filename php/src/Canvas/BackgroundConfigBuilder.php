<?php

namespace Grafana\Foundation\Canvas;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Canvas\BackgroundConfig>
 */
class BackgroundConfigBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Canvas\BackgroundConfig $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Canvas\BackgroundConfig();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Canvas\BackgroundConfig
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
    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\ResourceDimensionConfig> $image
     */
    public function image(\Grafana\Foundation\Cog\Builder $image): static
    {
        $imageResource = $image->build();
        $this->internal->image = $imageResource;
    
        return $this;
    }
    public function size(\Grafana\Foundation\Canvas\BackgroundImageSize $size): static
    {
        $this->internal->size = $size;
    
        return $this;
    }

}
