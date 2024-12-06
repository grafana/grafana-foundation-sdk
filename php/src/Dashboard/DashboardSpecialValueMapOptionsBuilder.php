<?php

namespace Grafana\Foundation\Dashboard;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\DashboardSpecialValueMapOptions>
 */
class DashboardSpecialValueMapOptionsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboard\DashboardSpecialValueMapOptions $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboard\DashboardSpecialValueMapOptions();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboard\DashboardSpecialValueMapOptions
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Special value to match against
     */
    public function match(\Grafana\Foundation\Dashboard\SpecialValueMatch $match): static
    {
        $this->internal->match = $match;
    
        return $this;
    }
    /**
     * Config to apply when the value matches the special value
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\ValueMappingResult> $result
     */
    public function result(\Grafana\Foundation\Cog\Builder $result): static
    {
        $resultResource = $result->build();
        $this->internal->result = $resultResource;
    
        return $this;
    }

}
