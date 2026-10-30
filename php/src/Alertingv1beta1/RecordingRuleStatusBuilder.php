<?php

namespace Grafana\Foundation\Alertingv1beta1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alertingv1beta1\RecordingRuleStatus>
 */
class RecordingRuleStatusBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Alertingv1beta1\RecordingRuleStatus $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Alertingv1beta1\RecordingRuleStatus();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Alertingv1beta1\RecordingRuleStatus
     */
    public function build()
    {
        return $this->internal;
    }

    public function health(\Grafana\Foundation\Alertingv1beta1\RecordingRuleHealth $health): static
    {
        $this->internal->health = $health;
    
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

}
