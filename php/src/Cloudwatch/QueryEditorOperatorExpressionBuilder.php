<?php

namespace Grafana\Foundation\Cloudwatch;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Cloudwatch\QueryEditorOperatorExpression>
 */
class QueryEditorOperatorExpressionBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Cloudwatch\QueryEditorOperatorExpression $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Cloudwatch\QueryEditorOperatorExpression();
    $this->internal->type = "operator";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Cloudwatch\QueryEditorOperatorExpression
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Cloudwatch\QueryEditorProperty> $property
     */
    public function property(\Grafana\Foundation\Cog\Builder $property): static
    {
        $propertyResource = $property->build();
        $this->internal->property = $propertyResource;
    
        return $this;
    }
    /**
     * TS type is operator: QueryEditorOperator<QueryEditorOperatorValueType>, extended in veneer
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Cloudwatch\QueryEditorOperator> $operator
     */
    public function operator(\Grafana\Foundation\Cog\Builder $operator): static
    {
        $operatorResource = $operator->build();
        $this->internal->operator = $operatorResource;
    
        return $this;
    }

}
