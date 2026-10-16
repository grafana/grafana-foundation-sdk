<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * The notebook's own panel chain. It is a copy of the dashboard one down to the transformation,
 * which follows the dashboard v2 shape rather than the v2beta1 shape in this package. The chain has
 * to be forked rather than shared because PanelKind reaches TransformationKind through
 * QueryGroupKind, and those three are what Dashboard v2beta1 serves. Everything the chain does not
 * change (DataLink, VizConfigKind, PanelQueryKind, QueryOptionsSpec) stays shared.
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\V2PanelKind>
 */
class V2PanelKindBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\V2PanelKind $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\V2PanelKind();
    $this->internal->kind = "Panel";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\V2PanelKind
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\V2PanelSpec> $spec
     */
    public function spec(\Grafana\Foundation\Cog\Builder $spec): static
    {
        $specResource = $spec->build();
        $this->internal->spec = $specResource;
    
        return $this;
    }

}
