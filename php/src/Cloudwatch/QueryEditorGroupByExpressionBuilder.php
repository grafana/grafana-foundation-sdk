<?php

namespace Grafana\Foundation\Cloudwatch;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Cloudwatch\QueryEditorGroupByExpression>
 */
class QueryEditorGroupByExpressionBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Cloudwatch\QueryEditorGroupByExpression $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Cloudwatch\QueryEditorGroupByExpression();
    $this->internal->type = "groupBy";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Cloudwatch\QueryEditorGroupByExpression
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

}
