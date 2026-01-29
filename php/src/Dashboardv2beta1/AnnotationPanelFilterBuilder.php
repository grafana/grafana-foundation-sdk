<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\AnnotationPanelFilter>
 */
class AnnotationPanelFilterBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\AnnotationPanelFilter $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\AnnotationPanelFilter();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\AnnotationPanelFilter
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Should the specified panels be included or excluded
     */
    public function exclude(bool $exclude): static
    {
        $this->internal->exclude = $exclude;
    
        return $this;
    }

    /**
     * Panel IDs that should be included or excluded
     * @param array<int> $ids
     */
    public function ids(array $ids): static
    {
        $this->internal->ids = $ids;
    
        return $this;
    }

}
