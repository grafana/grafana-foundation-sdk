<?php

namespace Grafana\Foundation\Dashboardv2;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\Dashboardv2SpecialValueMapOptions>
 */
class Dashboardv2SpecialValueMapOptionsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2\Dashboardv2SpecialValueMapOptions $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2\Dashboardv2SpecialValueMapOptions();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2\Dashboardv2SpecialValueMapOptions
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Special value to match against
     */
    public function match(\Grafana\Foundation\Dashboardv2\SpecialValueMatch $match): static
    {
        $this->internal->match = $match;
    
        return $this;
    }

    /**
     * Config to apply when the value matches the special value
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\ValueMappingResult> $result
     */
    public function result(\Grafana\Foundation\Cog\Builder $result): static
    {
        $resultResource = $result->build();
        $this->internal->result = $resultResource;
    
        return $this;
    }

}
