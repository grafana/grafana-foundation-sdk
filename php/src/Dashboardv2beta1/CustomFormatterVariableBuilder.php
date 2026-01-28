<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * Custom formatter variable
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\CustomFormatterVariable>
 */
class CustomFormatterVariableBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\CustomFormatterVariable $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\CustomFormatterVariable();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\CustomFormatterVariable
     */
    public function build()
    {
        return $this->internal;
    }

    public function name(string $name): static
    {
        $this->internal->name = $name;
    
        return $this;
    }

    public function type(\Grafana\Foundation\Dashboardv2beta1\VariableType $type): static
    {
        $this->internal->type = $type;
    
        return $this;
    }

    public function multi(bool $multi): static
    {
        $this->internal->multi = $multi;
    
        return $this;
    }

    public function includeAll(bool $includeAll): static
    {
        $this->internal->includeAll = $includeAll;
    
        return $this;
    }

}
