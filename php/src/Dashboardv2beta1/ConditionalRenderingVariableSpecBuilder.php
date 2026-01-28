<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingVariableSpec>
 */
class ConditionalRenderingVariableSpecBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingVariableSpec $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingVariableSpec();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingVariableSpec
     */
    public function build()
    {
        return $this->internal;
    }

    public function variable(string $variable): static
    {
        $this->internal->variable = $variable;
    
        return $this;
    }

    public function operator(\Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingVariableSpecOperator $operator): static
    {
        $this->internal->operator = $operator;
    
        return $this;
    }

    public function value(string $value): static
    {
        $this->internal->value = $value;
    
        return $this;
    }

}
