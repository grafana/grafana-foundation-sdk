<?php

namespace Grafana\Foundation\Bigquery;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Bigquery\QueryEditorFunctionExpression>
 */
class QueryEditorFunctionExpressionBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Bigquery\QueryEditorFunctionExpression $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Bigquery\QueryEditorFunctionExpression();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Bigquery\QueryEditorFunctionExpression
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
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Bigquery\QueryEditorFunctionParameterExpression>> $parameters
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
