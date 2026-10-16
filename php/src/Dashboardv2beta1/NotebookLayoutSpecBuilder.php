<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\NotebookLayoutSpec>
 */
class NotebookLayoutSpecBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\NotebookLayoutSpec $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\NotebookLayoutSpec();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\NotebookLayoutSpec
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\NotebookLayoutItemKind>> $cells
     */
    public function cells(array $cells): static
    {
            $cellsResources = [];
            foreach ($cells as $r1) {
                    $cellsResources[] = $r1->build();
            }
        $this->internal->cells = $cellsResources;
    
        return $this;
    }

}
