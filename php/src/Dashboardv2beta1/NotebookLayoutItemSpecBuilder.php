<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * One ordered item in a notebook layout. `element` references either a CellKind
 * (markdown/code content) or a V2PanelKind in the notebook's elements map. `source`
 * records who authored the cell; `collapsed` hides the body in the UI.
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\NotebookLayoutItemSpec>
 */
class NotebookLayoutItemSpecBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\NotebookLayoutItemSpec $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\NotebookLayoutItemSpec();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\NotebookLayoutItemSpec
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\ElementReference> $element
     */
    public function element(\Grafana\Foundation\Cog\Builder $element): static
    {
        $elementResource = $element->build();
        $this->internal->element = $elementResource;
    
        return $this;
    }

    public function source(\Grafana\Foundation\Dashboardv2beta1\NotebookLayoutItemSpecSource $source): static
    {
        $this->internal->source = $source;
    
        return $this;
    }

    public function collapsed(bool $collapsed): static
    {
        $this->internal->collapsed = $collapsed;
    
        return $this;
    }

}
