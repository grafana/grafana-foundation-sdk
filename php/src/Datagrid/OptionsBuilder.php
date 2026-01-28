<?php

namespace Grafana\Foundation\Datagrid;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Datagrid\Options>
 */
class OptionsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Datagrid\Options $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Datagrid\Options();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Datagrid\Options
     */
    public function build()
    {
        return $this->internal;
    }

    public function selectedSeries(int $selectedSeries): static
    {
        if (!($selectedSeries >= 0)) {
            throw new \ValueError('$selectedSeries must be >= 0');
        }
        $this->internal->selectedSeries = $selectedSeries;
    
        return $this;
    }

}
