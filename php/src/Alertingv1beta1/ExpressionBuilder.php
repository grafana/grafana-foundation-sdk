<?php

namespace Grafana\Foundation\Alertingv1beta1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alertingv1beta1\Expression>
 */
class ExpressionBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Alertingv1beta1\Expression $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Alertingv1beta1\Expression();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Alertingv1beta1\Expression
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * The type of query if this is a query expression
     */
    public function queryType(string $queryType): static
    {
        $this->internal->queryType = $queryType;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alertingv1beta1\RelativeTimeRange> $relativeTimeRange
     */
    public function relativeTimeRange(\Grafana\Foundation\Cog\Builder $relativeTimeRange): static
    {
        $relativeTimeRangeResource = $relativeTimeRange->build();
        $this->internal->relativeTimeRange = $relativeTimeRangeResource;
    
        return $this;
    }

    /**
     * The UID of the datasource to run this expression against. If omitted, the expression will be run against the `__expr__` datasource
     */
    public function datasourceUID(string $datasourceUID): static
    {
        if (!preg_match("/^[a-zA-Z0-9_-]+$/", $datasourceUID)) {
            throw new \ValueError('$datasourceUID must match regex ^[a-zA-Z0-9_-]+$');
        }
        $this->internal->datasourceUID = $datasourceUID;
    
        return $this;
    }

    /**
     * @param mixed $model
     */
    public function model( $model): static
    {
        $this->internal->model = $model;
    
        return $this;
    }

    /**
     * Used to mark the expression to be used as the final source for the rule evaluation
     * Only one expression in a rule can be marked as the source
     * For AlertRules, this is the expression that will be evaluated against the alerting condition
     * For RecordingRules, this is the expression that will be recorded
     */
    public function source(bool $source): static
    {
        $this->internal->source = $source;
    
        return $this;
    }

}
