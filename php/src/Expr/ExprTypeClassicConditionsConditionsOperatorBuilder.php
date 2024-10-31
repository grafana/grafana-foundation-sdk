<?php

namespace Grafana\Foundation\Expr;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Expr\ExprTypeClassicConditionsConditionsOperator>
 */
class ExprTypeClassicConditionsConditionsOperatorBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Expr\ExprTypeClassicConditionsConditionsOperator $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Expr\ExprTypeClassicConditionsConditionsOperator();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Expr\ExprTypeClassicConditionsConditionsOperator
     */
    public function build()
    {
        return $this->internal;
    }

    public function type(\Grafana\Foundation\Expr\TypeClassicConditionsType $type): static
    {
        $this->internal->type = $type;
    
        return $this;
    }

}
