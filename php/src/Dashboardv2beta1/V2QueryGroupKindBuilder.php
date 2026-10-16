<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\V2QueryGroupKind>
 */
class V2QueryGroupKindBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\V2QueryGroupKind $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\V2QueryGroupKind();
    $this->internal->kind = "QueryGroup";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\V2QueryGroupKind
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\V2QueryGroupSpec> $spec
     */
    public function spec(\Grafana\Foundation\Cog\Builder $spec): static
    {
        $specResource = $spec->build();
        $this->internal->spec = $specResource;
    
        return $this;
    }

}
