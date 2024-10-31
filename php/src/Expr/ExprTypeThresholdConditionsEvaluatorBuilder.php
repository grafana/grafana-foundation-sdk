<?php

namespace Grafana\Foundation\Expr;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Expr\ExprTypeThresholdConditionsEvaluator>
 */
class ExprTypeThresholdConditionsEvaluatorBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Expr\ExprTypeThresholdConditionsEvaluator $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Expr\ExprTypeThresholdConditionsEvaluator();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Expr\ExprTypeThresholdConditionsEvaluator
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * @param array<float> $params
     */
    public function params(array $params): static
    {
        $this->internal->params = $params;
    
        return $this;
    }
    /**
     * e.g. "gt"
     */
    public function type(\Grafana\Foundation\Expr\TypeThresholdType $type): static
    {
        $this->internal->type = $type;
    
        return $this;
    }

}
