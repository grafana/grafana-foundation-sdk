<?php

namespace Grafana\Foundation\Dashboard;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\DashboardFieldConfigSourceOverrides>
 */
class DashboardFieldConfigSourceOverridesBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboard\DashboardFieldConfigSourceOverrides $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboard\DashboardFieldConfigSourceOverrides();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboard\DashboardFieldConfigSourceOverrides
     */
    public function build()
    {
        return $this->internal;
    }

    public function matcher(\Grafana\Foundation\Dashboard\MatcherConfig $matcher): static
    {
        $this->internal->matcher = $matcher;
    
        return $this;
    }
    /**
     * @param array<\Grafana\Foundation\Dashboard\DynamicConfigValue> $properties
     */
    public function properties(array $properties): static
    {
        $this->internal->properties = $properties;
    
        return $this;
    }

}
