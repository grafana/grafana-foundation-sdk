<?php

namespace Grafana\Foundation\Expr;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Expr\ExprTypeThresholdConditionsUnloadEvaluator>
 */
class ExprTypeThresholdConditionsUnloadEvaluatorBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Expr\ExprTypeThresholdConditionsUnloadEvaluator $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Expr\ExprTypeThresholdConditionsUnloadEvaluator();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Expr\ExprTypeThresholdConditionsUnloadEvaluator
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
