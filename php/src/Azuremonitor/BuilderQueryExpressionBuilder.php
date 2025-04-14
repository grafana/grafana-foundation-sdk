<?php

namespace Grafana\Foundation\Azuremonitor;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\BuilderQueryExpression>
 */
class BuilderQueryExpressionBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Azuremonitor\BuilderQueryExpression $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Azuremonitor\BuilderQueryExpression();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Azuremonitor\BuilderQueryExpression
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\BuilderQueryEditorPropertyExpression> $from
     */
    public function from(\Grafana\Foundation\Cog\Builder $from): static
    {
        $fromResource = $from->build();
        $this->internal->from = $fromResource;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\BuilderQueryEditorColumnsExpression> $columns
     */
    public function columns(\Grafana\Foundation\Cog\Builder $columns): static
    {
        $columnsResource = $columns->build();
        $this->internal->columns = $columnsResource;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\BuilderQueryEditorWhereExpressionArray> $where
     */
    public function where(\Grafana\Foundation\Cog\Builder $where): static
    {
        $whereResource = $where->build();
        $this->internal->where = $whereResource;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\BuilderQueryEditorReduceExpressionArray> $reduce
     */
    public function reduce(\Grafana\Foundation\Cog\Builder $reduce): static
    {
        $reduceResource = $reduce->build();
        $this->internal->reduce = $reduceResource;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\BuilderQueryEditorGroupByExpressionArray> $groupBy
     */
    public function groupBy(\Grafana\Foundation\Cog\Builder $groupBy): static
    {
        $groupByResource = $groupBy->build();
        $this->internal->groupBy = $groupByResource;
    
        return $this;
    }

    public function limit(int $limit): static
    {
        $this->internal->limit = $limit;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\BuilderQueryEditorOrderByExpressionArray> $orderBy
     */
    public function orderBy(\Grafana\Foundation\Cog\Builder $orderBy): static
    {
        $orderByResource = $orderBy->build();
        $this->internal->orderBy = $orderByResource;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\BuilderQueryEditorWhereExpressionArray> $fuzzySearch
     */
    public function fuzzySearch(\Grafana\Foundation\Cog\Builder $fuzzySearch): static
    {
        $fuzzySearchResource = $fuzzySearch->build();
        $this->internal->fuzzySearch = $fuzzySearchResource;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\BuilderQueryEditorWhereExpressionArray> $timeFilter
     */
    public function timeFilter(\Grafana\Foundation\Cog\Builder $timeFilter): static
    {
        $timeFilterResource = $timeFilter->build();
        $this->internal->timeFilter = $timeFilterResource;
    
        return $this;
    }

}
