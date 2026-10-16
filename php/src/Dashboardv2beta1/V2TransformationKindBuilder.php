<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * Dashboard v2 shape: the transformation ID moved from `kind` to `group`.
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\V2TransformationKind>
 */
class V2TransformationKindBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\V2TransformationKind $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\V2TransformationKind();
    $this->internal->kind = "Transformation";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\V2TransformationKind
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * The group is the transformation ID
     */
    public function group(string $group): static
    {
        $this->internal->group = $group;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\V2TransformationSpec> $spec
     */
    public function spec(\Grafana\Foundation\Cog\Builder $spec): static
    {
        $specResource = $spec->build();
        $this->internal->spec = $specResource;
    
        return $this;
    }

}
