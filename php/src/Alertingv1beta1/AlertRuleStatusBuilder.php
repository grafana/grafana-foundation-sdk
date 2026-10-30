<?php

namespace Grafana\Foundation\Alertingv1beta1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alertingv1beta1\AlertRuleStatus>
 */
class AlertRuleStatusBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Alertingv1beta1\AlertRuleStatus $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Alertingv1beta1\AlertRuleStatus();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Alertingv1beta1\AlertRuleStatus
     */
    public function build()
    {
        return $this->internal;
    }

    public function health(\Grafana\Foundation\Alertingv1beta1\AlertRuleHealth $health): static
    {
        $this->internal->health = $health;
    
        return $this;
    }

    public function state(\Grafana\Foundation\Alertingv1beta1\AlertRuleState $state): static
    {
        $this->internal->state = $state;
    
        return $this;
    }

    public function stateReason(\Grafana\Foundation\Alertingv1beta1\AlertRuleStateReason $stateReason): static
    {
        $this->internal->stateReason = $stateReason;
    
        return $this;
    }

    public function lastEvaluationTime(string $lastEvaluationTime): static
    {
        $this->internal->lastEvaluationTime = $lastEvaluationTime;
    
        return $this;
    }

    /**
     * duration of the last evaluation in seconds
     */
    public function evaluationDuration(float $evaluationDuration): static
    {
        $this->internal->evaluationDuration = $evaluationDuration;
    
        return $this;
    }

    public function lastError(string $lastError): static
    {
        $this->internal->lastError = $lastError;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alertingv1beta1\AlertRuleInstanceTotals> $totals
     */
    public function totals(\Grafana\Foundation\Cog\Builder $totals): static
    {
        $totalsResource = $totals->build();
        $this->internal->totals = $totalsResource;
    
        return $this;
    }

}
