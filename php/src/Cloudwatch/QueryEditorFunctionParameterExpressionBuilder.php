<?php

namespace Grafana\Foundation\Cloudwatch;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Cloudwatch\QueryEditorFunctionParameterExpression>
 */
class QueryEditorFunctionParameterExpressionBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Cloudwatch\QueryEditorFunctionParameterExpression $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Cloudwatch\QueryEditorFunctionParameterExpression();
    $this->internal->type = "functionParameter";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Cloudwatch\QueryEditorFunctionParameterExpression
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
