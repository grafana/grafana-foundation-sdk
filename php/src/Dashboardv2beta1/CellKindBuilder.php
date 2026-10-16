<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * A cell holds non-panel narrative content (markdown text, code) in a notebook layout.
 * Panel cells are not represented here — they reuse V2PanelKind.
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\CellKind>
 */
class CellKindBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\CellKind $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\CellKind();
    $this->internal->kind = "Cell";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\CellKind
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\CellSpec> $spec
     */
    public function spec(\Grafana\Foundation\Cog\Builder $spec): static
    {
        $specResource = $spec->build();
        $this->internal->spec = $specResource;
    
        return $this;
    }

}
