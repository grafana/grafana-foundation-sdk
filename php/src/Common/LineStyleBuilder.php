<?php

namespace Grafana\Foundation\Common;

/**
 * TODO docs
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\LineStyle>
 */
class LineStyleBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Common\LineStyle $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Common\LineStyle();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Common\LineStyle
     */
    public function build()
    {
        return $this->internal;
    }

    public function fill(\Grafana\Foundation\Common\LineStyleFill $fill): static
    {
        $this->internal->fill = $fill;
    
        return $this;
    }
    /**
     * @param array<float> $dash
     */
    public function dash(array $dash): static
    {
        $this->internal->dash = $dash;
    
        return $this;
    }

}
