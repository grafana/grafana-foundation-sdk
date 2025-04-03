<?php

namespace Grafana\Foundation\Expr;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Expr\ExprTypeClassicConditionsConditionsReducer>
 */
class ExprTypeClassicConditionsConditionsReducerBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Expr\ExprTypeClassicConditionsConditionsReducer $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Expr\ExprTypeClassicConditionsConditionsReducer();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Expr\ExprTypeClassicConditionsConditionsReducer
     */
    public function build()
    {
        return $this->internal;
    }

    public function type(string $type): static
    {
        $this->internal->type = $type;
    
        return $this;
    }

}
