<?php

namespace Grafana\Foundation\Prometheus;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Prometheus\PrometheusDataqueryScope>
 */
class PrometheusDataqueryScopeBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Prometheus\PrometheusDataqueryScope $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Prometheus\PrometheusDataqueryScope();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Prometheus\PrometheusDataqueryScope
     */
    public function build()
    {
        return $this->internal;
    }

    public function matchers(string $matchers): static
    {
        $this->internal->matchers = $matchers;
    
        return $this;
    }

}
