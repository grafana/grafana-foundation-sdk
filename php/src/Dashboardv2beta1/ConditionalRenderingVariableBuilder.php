<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingVariableKind>
 */
class ConditionalRenderingVariableBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingVariableKind $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingVariableKind();
    $this->internal->kind = "ConditionalRenderingVariable";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingVariableKind
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingVariableSpec> $spec
     */
    public function spec(\Grafana\Foundation\Cog\Builder $spec): static
    {
        $specResource = $spec->build();
        $this->internal->spec = $specResource;
    
        return $this;
    }

    public function variable(string $variable): static
    {
        $this->internal->spec->variable = $variable;
    
        return $this;
    }

    public function operator(\Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingVariableSpecOperator $operator): static
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
