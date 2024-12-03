<?php

namespace Grafana\Foundation\Bigquery;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Bigquery\SQLExpression>
 */
class SQLExpressionBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Bigquery\SQLExpression $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Bigquery\SQLExpression();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Bigquery\SQLExpression
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Bigquery\QueryEditorFunctionExpression>> $columns
     */
    public function columns(array $columns): static
    {
            $columnsResources = [];
            foreach ($columns as $r1) {
                    $columnsResources[] = $r1->build();
            }
        $this->internal->columns = $columnsResources;
    
        return $this;
    }
    public function from(string $from): static
    {
        $this->internal->from = $from;
    
        return $this;
    }
    /**
     * whereJsonTree?: _
     */
    public function whereString(string $whereString): static
    {
        $this->internal->whereString = $whereString;
    
        return $this;
    }
    /**
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Bigquery\QueryEditorGroupByExpression>> $groupBy
     */
    public function groupBy(array $groupBy): static
    {
            $groupByResources = [];
            foreach ($groupBy as $r1) {
                    $groupByResources[] = $r1->build();
            }
        $this->internal->groupBy = $groupByResources;
    
        return $this;
    }
    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Bigquery\QueryEditorPropertyExpression> $orderBy
     */
    public function orderBy(\Grafana\Foundation\Cog\Builder $orderBy): static
    {
        $orderByResource = $orderBy->build();
        $this->internal->orderBy = $orderByResource;
    
        return $this;
    }
    public function orderByDirection(\Grafana\Foundation\Bigquery\OrderByDirection $orderByDirection): static
    {
        $this->internal->orderByDirection = $orderByDirection;
    
        return $this;
    }
    public function limit(int $limit): static
    {
        $this->internal->limit = $limit;
    
        return $this;
    }
    public function offset(int $offset): static
    {
        $this->internal->offset = $offset;
    
        return $this;
    }

}
