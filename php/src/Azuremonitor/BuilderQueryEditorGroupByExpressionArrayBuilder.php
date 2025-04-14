<?php

namespace Grafana\Foundation\Azuremonitor;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\BuilderQueryEditorGroupByExpressionArray>
 */
class BuilderQueryEditorGroupByExpressionArrayBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Azuremonitor\BuilderQueryEditorGroupByExpressionArray $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Azuremonitor\BuilderQueryEditorGroupByExpressionArray();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Azuremonitor\BuilderQueryEditorGroupByExpressionArray
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\BuilderQueryEditorGroupByExpression>> $expressions
     */
    public function expressions(array $expressions): static
    {
            $expressionsResources = [];
            foreach ($expressions as $r1) {
                    $expressionsResources[] = $r1->build();
            }
        $this->internal->expressions = $expressionsResources;
    
        return $this;
    }

    public function type(\Grafana\Foundation\Azuremonitor\BuilderQueryEditorExpressionType $type): static
    {
        $this->internal->type = $type;
    
        return $this;
    }

}
