<?php

namespace App\Custom;

use Grafana\Foundation\Cog;

/**
 * @implements Cog\Builder<Query>
 */
class QueryBuilder implements Cog\Builder
{
    protected Query $internal;

    public function __construct(string $query)
    {
    	$this->internal = new Query(expr: $query);
    }

    /**
     * @return Query
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * The actual expression/query that will be evaluated by Prometheus
     */
    public function expr(string $expr): static
    {
        $this->internal->expr = $expr;
        return $this;
    }

    /**
     * A unique identifier for the query within the list of targets.
     * In server side expressions, the refId is used as a variable name to identify results.
     * By default, the UI will assign A->Z; however setting meaningful names may be useful.
     */
    public function refId(string $refId): static
    {
        $this->internal->refId = $refId;
        return $this;
    }

    /**
     * If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
     */
    public function hide(bool $hide): static
    {
        $this->internal->hide = $hide;
        return $this;
    }
}