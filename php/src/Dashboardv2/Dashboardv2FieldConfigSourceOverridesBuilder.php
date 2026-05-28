<?php

namespace Grafana\Foundation\Dashboardv2;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\Dashboardv2FieldConfigSourceOverrides>
 */
class Dashboardv2FieldConfigSourceOverridesBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2\Dashboardv2FieldConfigSourceOverrides $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2\Dashboardv2FieldConfigSourceOverrides();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2\Dashboardv2FieldConfigSourceOverrides
     */
    public function build()
    {
        return $this->internal;
    }

    public function systemRef(string $systemRef): static
    {
        $this->internal->systemRef = $systemRef;
    
        return $this;
    }

    public function matcher(\Grafana\Foundation\Dashboardv2\MatcherConfig $matcher): static
    {
        $this->internal->matcher = $matcher;
    
        return $this;
    }

    /**
     * @param array<\Grafana\Foundation\Dashboardv2\DynamicConfigValue> $properties
     */
    public function properties(array $properties): static
    {
        $this->internal->properties = $properties;
    
        return $this;
    }

}
