<?php

namespace Grafana\Foundation\Dashboard;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\AnnotationPanelFilter>
 */
class AnnotationPanelFilterBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboard\AnnotationPanelFilter $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboard\AnnotationPanelFilter();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboard\AnnotationPanelFilter
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
