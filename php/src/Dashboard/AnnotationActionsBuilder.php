<?php

namespace Grafana\Foundation\Dashboard;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\AnnotationActions>
 */
class AnnotationActionsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboard\AnnotationActions $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboard\AnnotationActions();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboard\AnnotationActions
     */
    public function build()
    {
        return $this->internal;
    }

    public function canAdd(bool $canAdd): static
    {
        $this->internal->canAdd = $canAdd;
    
        return $this;
    }
    public function canDelete(bool $canDelete): static
    {
        $this->internal->canDelete = $canDelete;
    
        return $this;
    }
    public function canEdit(bool $canEdit): static
    {
        $this->internal->canEdit = $canEdit;
    
        return $this;
    }

}
