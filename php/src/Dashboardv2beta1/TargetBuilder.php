<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\PanelQueryKind>
 */
class TargetBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\PanelQueryKind $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\PanelQueryKind();
    $this->internal->kind = "PanelQuery";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\PanelQueryKind
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\DataQueryKind> $query
     */
    public function query(\Grafana\Foundation\Cog\Builder $query): static
    {
        $queryResource = $query->build();
        $this->internal->spec->query = $queryResource;
    
        return $this;
    }

    public function refId(string $refId): static
    {
        $this->internal->spec->refId = $refId;
    
        return $this;
    }

    public function hidden(bool $hidden): static
    {
        $this->internal->spec->hidden = $hidden;
    
        return $this;
    }

}
