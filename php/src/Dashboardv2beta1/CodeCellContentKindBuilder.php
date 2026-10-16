<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\CodeCellContentKind>
 */
class CodeCellContentKindBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\CodeCellContentKind $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\CodeCellContentKind();
    $this->internal->kind = "Code";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\CodeCellContentKind
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\CodeCellContentSpec> $spec
     */
    public function spec(\Grafana\Foundation\Cog\Builder $spec): static
    {
        $specResource = $spec->build();
        $this->internal->spec = $specResource;
    
        return $this;
    }

}
