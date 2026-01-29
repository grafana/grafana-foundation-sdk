<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * FIXME: should we introduce this? --- Variable value option
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\VariableValueOption>
 */
class VariableValueOptionBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\VariableValueOption $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\VariableValueOption();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\VariableValueOption
     */
    public function build()
    {
        return $this->internal;
    }

    public function label(string $label): static
    {
        $this->internal->label = $label;
    
        return $this;
    }

    /**
     * @param string|bool|float|\Grafana\Foundation\Dashboardv2beta1\CustomVariableValue $value
     */
    public function value( $value): static
    {
        $this->internal->value = $value;
    
        return $this;
    }

    public function group(string $group): static
    {
        $this->internal->group = $group;
    
        return $this;
    }

}
