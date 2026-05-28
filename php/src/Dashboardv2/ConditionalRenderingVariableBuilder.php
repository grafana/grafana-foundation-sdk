<?php

namespace Grafana\Foundation\Dashboardv2;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\ConditionalRenderingVariableKind>
 */
class ConditionalRenderingVariableBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2\ConditionalRenderingVariableKind $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2\ConditionalRenderingVariableKind();
    $this->internal->kind = "ConditionalRenderingVariable";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2\ConditionalRenderingVariableKind
     */
    public function build()
    {
        return $this->internal;
    }

    public function variable(string $variable): static
    {
        $this->internal->spec->variable = $variable;
    
        return $this;
    }

    public function operator(\Grafana\Foundation\Dashboardv2\ConditionalRenderingVariableSpecOperator $operator): static
    {
        $this->internal->spec->operator = $operator;
    
        return $this;
    }

    public function value(string $value): static
    {
        $this->internal->spec->value = $value;
    
        return $this;
    }

}
