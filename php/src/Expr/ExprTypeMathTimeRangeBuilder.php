<?php

namespace Grafana\Foundation\Expr;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Expr\ExprTypeMathTimeRange>
 */
class ExprTypeMathTimeRangeBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Expr\ExprTypeMathTimeRange $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Expr\ExprTypeMathTimeRange();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Expr\ExprTypeMathTimeRange
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * From is the start time of the query.
     */
    public function from(string $from): static
    {
        $this->internal->from = $from;
    
        return $this;
    }
    /**
     * To is the end time of the query.
     */
    public function to(string $to): static
    {
        $this->internal->to = $to;
    
        return $this;
    }

}
