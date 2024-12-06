<?php

namespace Grafana\Foundation\Expr;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Expr\ExprTypeClassicConditionsConditionsQuery>
 */
class ExprTypeClassicConditionsConditionsQueryBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Expr\ExprTypeClassicConditionsConditionsQuery $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Expr\ExprTypeClassicConditionsConditionsQuery();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Expr\ExprTypeClassicConditionsConditionsQuery
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * @param array<string> $params
     */
    public function params(array $params): static
    {
        $this->internal->params = $params;
    
        return $this;
    }

}
