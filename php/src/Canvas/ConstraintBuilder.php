<?php

namespace Grafana\Foundation\Canvas;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Canvas\Constraint>
 */
class ConstraintBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Canvas\Constraint $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Canvas\Constraint();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Canvas\Constraint
     */
    public function build()
    {
        return $this->internal;
    }

    public function horizontal(\Grafana\Foundation\Canvas\HorizontalConstraint $horizontal): static
    {
        $this->internal->horizontal = $horizontal;
    
        return $this;
    }
    public function vertical(\Grafana\Foundation\Canvas\VerticalConstraint $vertical): static
    {
        $this->internal->vertical = $vertical;
    
        return $this;
    }

}
