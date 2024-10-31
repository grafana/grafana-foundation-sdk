<?php

namespace Grafana\Foundation\Cloudwatch;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Cloudwatch\QueryEditorArrayExpression>
 */
class QueryEditorArrayExpressionBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Cloudwatch\QueryEditorArrayExpression $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Cloudwatch\QueryEditorArrayExpression();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Cloudwatch\QueryEditorArrayExpression
     */
    public function build()
    {
        return $this->internal;
    }

    public function type(\Grafana\Foundation\Cloudwatch\QueryEditorArrayExpressionType $type): static
    {
        $this->internal->type = $type;
    
        return $this;
    }
    /**
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Cloudwatch\QueryEditorArrayExpression>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Cloudwatch\QueryEditorPropertyExpression>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Cloudwatch\QueryEditorGroupByExpression>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Cloudwatch\QueryEditorFunctionExpression>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Cloudwatch\QueryEditorFunctionParameterExpression>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Cloudwatch\QueryEditorOperatorExpression>> $expressions
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
