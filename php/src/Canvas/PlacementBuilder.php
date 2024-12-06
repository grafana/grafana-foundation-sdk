<?php

namespace Grafana\Foundation\Canvas;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Canvas\Placement>
 */
class PlacementBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Canvas\Placement $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Canvas\Placement();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Canvas\Placement
     */
    public function build()
    {
        return $this->internal;
    }

    public function top(float $top): static
    {
        $this->internal->top = $top;
    
        return $this;
    }
    public function left(float $left): static
    {
        $this->internal->left = $left;
    
        return $this;
    }
    public function right(float $right): static
    {
        $this->internal->right = $right;
    
        return $this;
    }
    public function bottom(float $bottom): static
    {
        $this->internal->bottom = $bottom;
    
        return $this;
    }
    public function width(float $width): static
    {
        $this->internal->width = $width;
    
        return $this;
    }
    public function height(float $height): static
    {
        $this->internal->height = $height;
    
        return $this;
    }
    public function rotation(float $rotation): static
    {
        $this->internal->rotation = $rotation;
    
        return $this;
    }

}
