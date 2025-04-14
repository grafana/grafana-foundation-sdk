<?php

namespace Grafana\Foundation\Azuremonitor;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\BuilderQueryEditorReduceExpression>
 */
class BuilderQueryEditorReduceExpressionBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Azuremonitor\BuilderQueryEditorReduceExpression $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Azuremonitor\BuilderQueryEditorReduceExpression();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Azuremonitor\BuilderQueryEditorReduceExpression
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\BuilderQueryEditorProperty> $property
     */
    public function property(\Grafana\Foundation\Cog\Builder $property): static
    {
        $propertyResource = $property->build();
        $this->internal->property = $propertyResource;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\BuilderQueryEditorProperty> $reduce
     */
    public function reduce(\Grafana\Foundation\Cog\Builder $reduce): static
    {
        $reduceResource = $reduce->build();
        $this->internal->reduce = $reduceResource;
    
        return $this;
    }

    /**
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\BuilderQueryEditorFunctionParameterExpression>> $parameters
     */
    public function parameters(array $parameters): static
    {
            $parametersResources = [];
            foreach ($parameters as $r1) {
                    $parametersResources[] = $r1->build();
            }
        $this->internal->parameters = $parametersResources;
    
        return $this;
    }

    public function focus(bool $focus): static
    {
        $this->internal->focus = $focus;
    
        return $this;
    }

}
