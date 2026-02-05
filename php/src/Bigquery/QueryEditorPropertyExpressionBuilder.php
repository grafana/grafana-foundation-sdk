<?php

namespace Grafana\Foundation\Bigquery;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Bigquery\QueryEditorPropertyExpression>
 */
class QueryEditorPropertyExpressionBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Bigquery\QueryEditorPropertyExpression $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Bigquery\QueryEditorPropertyExpression();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Bigquery\QueryEditorPropertyExpression
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
