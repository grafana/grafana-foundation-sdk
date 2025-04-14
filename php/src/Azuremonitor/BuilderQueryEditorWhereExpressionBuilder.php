<?php

namespace Grafana\Foundation\Azuremonitor;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\BuilderQueryEditorWhereExpression>
 */
class BuilderQueryEditorWhereExpressionBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Azuremonitor\BuilderQueryEditorWhereExpression $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Azuremonitor\BuilderQueryEditorWhereExpression();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Azuremonitor\BuilderQueryEditorWhereExpression
     */
    public function build()
    {
        return $this->internal;
    }

    public function type(\Grafana\Foundation\Azuremonitor\BuilderQueryEditorExpressionType $type): static
    {
        $this->internal->type = $type;
    
        return $this;
    }

    /**
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\BuilderQueryEditorWhereExpressionItems>> $expressions
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

}
