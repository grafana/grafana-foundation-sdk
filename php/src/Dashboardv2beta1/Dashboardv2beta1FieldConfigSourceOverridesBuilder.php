<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1FieldConfigSourceOverrides>
 */
class Dashboardv2beta1FieldConfigSourceOverridesBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1FieldConfigSourceOverrides $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1FieldConfigSourceOverrides();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1FieldConfigSourceOverrides
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

    public function matcher(\Grafana\Foundation\Dashboardv2beta1\MatcherConfig $matcher): static
    {
        $this->internal->matcher = $matcher;
    
        return $this;
    }

    /**
     * @param array<\Grafana\Foundation\Dashboardv2beta1\DynamicConfigValue> $properties
     */
    public function properties(array $properties): static
    {
        $this->internal->properties = $properties;
    
        return $this;
    }

}
