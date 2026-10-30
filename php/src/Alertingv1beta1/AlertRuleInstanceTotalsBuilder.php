<?php

namespace Grafana\Foundation\Alertingv1beta1;

/**
 * Count of alert instances per state. error also counts instances whose evaluation
 * errored but were mapped to another state via execErrState, so it can overlap.
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alertingv1beta1\AlertRuleInstanceTotals>
 */
class AlertRuleInstanceTotalsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Alertingv1beta1\AlertRuleInstanceTotals $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Alertingv1beta1\AlertRuleInstanceTotals();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Alertingv1beta1\AlertRuleInstanceTotals
     */
    public function build()
    {
        return $this->internal;
    }

    public function healthy(int $healthy): static
    {
        $this->internal->healthy = $healthy;
    
        return $this;
    }

    public function firing(int $firing): static
    {
        $this->internal->firing = $firing;
    
        return $this;
    }

    public function pending(int $pending): static
    {
        $this->internal->pending = $pending;
    
        return $this;
    }

    public function recovering(int $recovering): static
    {
        $this->internal->recovering = $recovering;
    
        return $this;
    }

    public function nodata(int $nodata): static
    {
        $this->internal->nodata = $nodata;
    
        return $this;
    }

    public function error(int $error): static
    {
        $this->internal->error = $error;
    
        return $this;
    }

}
