<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\GridLayoutItemKind>
 */
class GridItemBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\GridLayoutItemKind $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\GridLayoutItemKind();
    $this->internal->kind = "GridLayoutItem";
    $this->internal->spec->element->kind = "ElementReference";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\GridLayoutItemKind
     */
    public function build()
    {
        return $this->internal;
    }

    public function x(int $x): static
    {
        $this->internal->spec->x = $x;
    
        return $this;
    }

    public function y(int $y): static
    {
        $this->internal->spec->y = $y;
    
        return $this;
    }

    public function width(int $width): static
    {
        $this->internal->spec->width = $width;
    
        return $this;
    }

    public function height(int $height): static
    {
        $this->internal->spec->height = $height;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\RepeatOptions> $repeat
     */
    public function repeat(\Grafana\Foundation\Cog\Builder $repeat): static
    {
        $repeatResource = $repeat->build();
        $this->internal->spec->repeat = $repeatResource;
    
        return $this;
    }

    public function name(string $name): static
    {
        $this->internal->spec->element->name = $name;
    
        return $this;
    }

}
