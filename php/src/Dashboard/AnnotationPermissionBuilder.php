<?php

namespace Grafana\Foundation\Dashboard;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\AnnotationPermission>
 */
class AnnotationPermissionBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboard\AnnotationPermission $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboard\AnnotationPermission();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboard\AnnotationPermission
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\AnnotationActions> $dashboard
     */
    public function dashboard(\Grafana\Foundation\Cog\Builder $dashboard): static
    {
        $dashboardResource = $dashboard->build();
        $this->internal->dashboard = $dashboardResource;
    
        return $this;
    }
    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\AnnotationActions> $organization
     */
    public function organization(\Grafana\Foundation\Cog\Builder $organization): static
    {
        $organizationResource = $organization->build();
        $this->internal->organization = $organizationResource;
    
        return $this;
    }

}
