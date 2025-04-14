<?php

namespace Grafana\Foundation\Azuremonitor;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\BuilderQueryEditorWhereExpressionArray>
 */
class BuilderQueryEditorWhereExpressionArrayBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Azuremonitor\BuilderQueryEditorWhereExpressionArray $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Azuremonitor\BuilderQueryEditorWhereExpressionArray();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Azuremonitor\BuilderQueryEditorWhereExpressionArray
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\BuilderQueryEditorWhereExpression>> $expressions
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
