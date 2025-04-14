<?php

namespace Grafana\Foundation\Bigquery;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Bigquery\QueryEditorFunctionParameterExpression>
 */
class QueryEditorFunctionParameterExpressionBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Bigquery\QueryEditorFunctionParameterExpression $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Bigquery\QueryEditorFunctionParameterExpression();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Bigquery\QueryEditorFunctionParameterExpression
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

}
