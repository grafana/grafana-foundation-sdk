<?php

namespace Grafana\Foundation\Nodegraph;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Nodegraph\ArcOption>
 */
class ArcOptionBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Nodegraph\ArcOption $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Nodegraph\ArcOption();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Nodegraph\ArcOption
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Field from which to get the value. Values should be less than 1, representing fraction of a circle.
     */
    public function field(string $field): static
    {
        $this->internal->field = $field;
    
        return $this;
    }
    /**
     * The color of the arc.
     */
    public function color(string $color): static
    {
        $this->internal->color = $color;
    
        return $this;
    }

}
