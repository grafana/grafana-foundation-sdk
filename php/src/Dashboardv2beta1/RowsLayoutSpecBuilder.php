<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\RowsLayoutSpec>
 */
class RowsLayoutSpecBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\RowsLayoutSpec $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\RowsLayoutSpec();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\RowsLayoutSpec
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\RowsLayoutRowKind>> $rows
     */
    public function rows(array $rows): static
    {
            $rowsResources = [];
            foreach ($rows as $r1) {
                    $rowsResources[] = $r1->build();
            }
        $this->internal->rows = $rowsResources;
    
        return $this;
    }

}
