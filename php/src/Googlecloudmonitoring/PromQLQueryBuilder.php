<?php

namespace Grafana\Foundation\Googlecloudmonitoring;

/**
 * PromQL sub-query properties.
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Googlecloudmonitoring\PromQLQuery>
 */
class PromQLQueryBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Googlecloudmonitoring\PromQLQuery $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Googlecloudmonitoring\PromQLQuery();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Googlecloudmonitoring\PromQLQuery
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * GCP project to execute the query against.
     */
    public function projectName(string $projectName): static
    {
        $this->internal->projectName = $projectName;
    
        return $this;
    }
    /**
     * PromQL expression/query to be executed.
     */
    public function expr(string $expr): static
    {
        $this->internal->expr = $expr;
    
        return $this;
    }
    /**
     * PromQL min step
     */
    public function step(string $step): static
    {
        $this->internal->step = $step;
    
        return $this;
    }

}
