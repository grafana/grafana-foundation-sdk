<?php

namespace Grafana\Foundation\Cloudwatch;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Cloudwatch\SQLExpression>
 */
class SQLExpressionBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Cloudwatch\SQLExpression $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Cloudwatch\SQLExpression();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Cloudwatch\SQLExpression
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * SELECT part of the SQL expression
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Cloudwatch\QueryEditorFunctionExpression> $select
     */
    public function select(\Grafana\Foundation\Cog\Builder $select): static
    {
        $selectResource = $select->build();
        $this->internal->select = $selectResource;
    
        return $this;
    }
    /**
     * FROM part of the SQL expression
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Cloudwatch\QueryEditorPropertyExpression>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Cloudwatch\QueryEditorFunctionExpression> $from
     */
    public function from( $from): static
    {
        $fromResource = $from->build();
        $this->internal->from = $fromResource;
    
        return $this;
    }
    /**
     * WHERE part of the SQL expression
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Cloudwatch\QueryEditorArrayExpression> $where
     */
    public function where(\Grafana\Foundation\Cog\Builder $where): static
    {
        $whereResource = $where->build();
        $this->internal->where = $whereResource;
    
        return $this;
    }
    /**
     * GROUP BY part of the SQL expression
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Cloudwatch\QueryEditorArrayExpression> $groupBy
     */
    public function groupBy(\Grafana\Foundation\Cog\Builder $groupBy): static
    {
        $groupByResource = $groupBy->build();
        $this->internal->groupBy = $groupByResource;
    
        return $this;
    }
    /**
     * ORDER BY part of the SQL expression
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Cloudwatch\QueryEditorFunctionExpression> $orderBy
     */
    public function orderBy(\Grafana\Foundation\Cog\Builder $orderBy): static
    {
        $orderByResource = $orderBy->build();
        $this->internal->orderBy = $orderByResource;
    
        return $this;
    }
    /**
     * The sort order of the SQL expression, `ASC` or `DESC`
     */
    public function orderByDirection(string $orderByDirection): static
    {
        $this->internal->orderByDirection = $orderByDirection;
    
        return $this;
    }
    /**
     * LIMIT part of the SQL expression
     */
    public function limit(int $limit): static
    {
        $this->internal->limit = $limit;
    
        return $this;
    }

}
