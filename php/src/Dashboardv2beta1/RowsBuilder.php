<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\RowsLayoutKind>
 */
class RowsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\RowsLayoutKind $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\RowsLayoutKind();
    $this->internal->kind = "RowsLayout";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\RowsLayoutKind
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
        $this->internal->spec->rows = $rowsResources;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\RowsLayoutRowKind> $row
     */
    public function row(\Grafana\Foundation\Cog\Builder $row): static
    {
        $rowResource = $row->build();
        $this->internal->spec->rows[] = $rowResource;
    
        return $this;
    }

}
