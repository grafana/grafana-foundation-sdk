<?php

namespace Grafana\Foundation\Cloudwatch;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Cloudwatch\QueryEditorFunctionExpression>
 */
class QueryEditorFunctionExpressionBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Cloudwatch\QueryEditorFunctionExpression $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Cloudwatch\QueryEditorFunctionExpression();
    $this->internal->type = "function";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Cloudwatch\QueryEditorFunctionExpression
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
    /**
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Cloudwatch\QueryEditorFunctionParameterExpression>> $parameters
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

}
