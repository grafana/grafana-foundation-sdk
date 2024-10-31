<?php

namespace Grafana\Foundation\Expr;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Expr\ExprTypeThresholdConditions>
 */
class ExprTypeThresholdConditionsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Expr\ExprTypeThresholdConditions $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Expr\ExprTypeThresholdConditions();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Expr\ExprTypeThresholdConditions
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Expr\ExprTypeThresholdConditionsEvaluator> $evaluator
     */
    public function evaluator(\Grafana\Foundation\Cog\Builder $evaluator): static
    {
        $evaluatorResource = $evaluator->build();
        $this->internal->evaluator = $evaluatorResource;
    
        return $this;
    }
    /**
     * @param mixed $loadedDimensions
     */
    public function loadedDimensions( $loadedDimensions): static
    {
        $this->internal->loadedDimensions = $loadedDimensions;
    
        return $this;
    }
    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Expr\ExprTypeThresholdConditionsUnloadEvaluator> $unloadEvaluator
     */
    public function unloadEvaluator(\Grafana\Foundation\Cog\Builder $unloadEvaluator): static
    {
        $unloadEvaluatorResource = $unloadEvaluator->build();
        $this->internal->unloadEvaluator = $unloadEvaluatorResource;
    
        return $this;
    }

}
