<?php

namespace Grafana\Foundation\Expr;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Expr\ExprTypeClassicConditionsConditionsEvaluator>
 */
class ExprTypeClassicConditionsConditionsEvaluatorBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Expr\ExprTypeClassicConditionsConditionsEvaluator $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Expr\ExprTypeClassicConditionsConditionsEvaluator();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Expr\ExprTypeClassicConditionsConditionsEvaluator
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
    public function type(string $type): static
    {
        $this->internal->type = $type;
    
        return $this;
    }

}
