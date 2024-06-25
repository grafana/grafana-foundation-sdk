<?php

namespace Grafana\Foundation\Expr;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Expr\ExprTypeThresholdTimeRange>
 */
class ExprTypeThresholdTimeRangeBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Expr\ExprTypeThresholdTimeRange $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Expr\ExprTypeThresholdTimeRange();
    }

    /**
     * @return \Grafana\Foundation\Expr\ExprTypeThresholdTimeRange
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
