<?php

namespace Grafana\Foundation\Azuremonitor;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\BuilderQueryEditorColumnsExpression>
 */
class BuilderQueryEditorColumnsExpressionBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Azuremonitor\BuilderQueryEditorColumnsExpression $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Azuremonitor\BuilderQueryEditorColumnsExpression();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Azuremonitor\BuilderQueryEditorColumnsExpression
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * @param array<string> $columns
     */
    public function columns(array $columns): static
    {
        $this->internal->columns = $columns;
    
        return $this;
    }

    public function type(\Grafana\Foundation\Azuremonitor\BuilderQueryEditorExpressionType $type): static
    {
        $this->internal->type = $type;
    
        return $this;
    }

}
