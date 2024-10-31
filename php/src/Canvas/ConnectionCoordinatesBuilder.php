<?php

namespace Grafana\Foundation\Canvas;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Canvas\ConnectionCoordinates>
 */
class ConnectionCoordinatesBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Canvas\ConnectionCoordinates $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Canvas\ConnectionCoordinates();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Canvas\ConnectionCoordinates
     */
    public function build()
    {
        return $this->internal;
    }

    public function x(float $x): static
    {
        $this->internal->x = $x;
    
        return $this;
    }
    public function y(float $y): static
    {
        $this->internal->y = $y;
    
        return $this;
    }

}
