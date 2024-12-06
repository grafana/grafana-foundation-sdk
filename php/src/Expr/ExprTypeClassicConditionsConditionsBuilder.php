<?php

namespace Grafana\Foundation\Expr;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Expr\ExprTypeClassicConditionsConditions>
 */
class ExprTypeClassicConditionsConditionsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Expr\ExprTypeClassicConditionsConditions $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Expr\ExprTypeClassicConditionsConditions();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Expr\ExprTypeClassicConditionsConditions
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Expr\ExprTypeClassicConditionsConditionsEvaluator> $evaluator
     */
    public function evaluator(\Grafana\Foundation\Cog\Builder $evaluator): static
    {
        $evaluatorResource = $evaluator->build();
        $this->internal->evaluator = $evaluatorResource;
    
        return $this;
    }
    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Expr\ExprTypeClassicConditionsConditionsOperator> $operator
     */
    public function operator(\Grafana\Foundation\Cog\Builder $operator): static
    {
        $operatorResource = $operator->build();
        $this->internal->operator = $operatorResource;
    
        return $this;
    }
    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Expr\ExprTypeClassicConditionsConditionsQuery> $query
     */
    public function query(\Grafana\Foundation\Cog\Builder $query): static
    {
        $queryResource = $query->build();
        $this->internal->query = $queryResource;
    
        return $this;
    }
    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Expr\ExprTypeClassicConditionsConditionsReducer> $reducer
     */
    public function reducer(\Grafana\Foundation\Cog\Builder $reducer): static
    {
        $reducerResource = $reducer->build();
        $this->internal->reducer = $reducerResource;
    
        return $this;
    }

}
