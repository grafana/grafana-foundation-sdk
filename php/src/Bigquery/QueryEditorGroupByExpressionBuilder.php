<?php

namespace Grafana\Foundation\Bigquery;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Bigquery\QueryEditorGroupByExpression>
 */
class QueryEditorGroupByExpressionBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Bigquery\QueryEditorGroupByExpression $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Bigquery\QueryEditorGroupByExpression();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Bigquery\QueryEditorGroupByExpression
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Bigquery\QueryEditorProperty> $property
     */
    public function property(\Grafana\Foundation\Cog\Builder $property): static
    {
        $propertyResource = $property->build();
        $this->internal->property = $propertyResource;
    
        return $this;
    }

}
