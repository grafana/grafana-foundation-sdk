<?php

namespace Grafana\Foundation\Alertingv1beta1;

class AlertRuleStatus implements \JsonSerializable
{
    public ?\Grafana\Foundation\Alertingv1beta1\AlertRuleHealth $health;

    public ?\Grafana\Foundation\Alertingv1beta1\AlertRuleState $state;

    public ?\Grafana\Foundation\Alertingv1beta1\AlertRuleStateReason $stateReason;

    public ?string $lastEvaluationTime;

    /**
     * duration of the last evaluation in seconds
     */
    public ?float $evaluationDuration;

    public ?string $lastError;

    public ?\Grafana\Foundation\Alertingv1beta1\AlertRuleInstanceTotals $totals;

    /**
     * @param \Grafana\Foundation\Alertingv1beta1\AlertRuleHealth|null $health
     * @param \Grafana\Foundation\Alertingv1beta1\AlertRuleState|null $state
     * @param \Grafana\Foundation\Alertingv1beta1\AlertRuleStateReason|null $stateReason
     * @param string|null $lastEvaluationTime
     * @param float|null $evaluationDuration
     * @param string|null $lastError
     * @param \Grafana\Foundation\Alertingv1beta1\AlertRuleInstanceTotals|null $totals
     */
    public function __construct(?\Grafana\Foundation\Alertingv1beta1\AlertRuleHealth $health = null, ?\Grafana\Foundation\Alertingv1beta1\AlertRuleState $state = null, ?\Grafana\Foundation\Alertingv1beta1\AlertRuleStateReason $stateReason = null, ?string $lastEvaluationTime = null, ?float $evaluationDuration = null, ?string $lastError = null, ?\Grafana\Foundation\Alertingv1beta1\AlertRuleInstanceTotals $totals = null)
    {
        $this->health = $health;
        $this->state = $state;
        $this->stateReason = $stateReason;
        $this->lastEvaluationTime = $lastEvaluationTime;
        $this->evaluationDuration = $evaluationDuration;
        $this->lastError = $lastError;
        $this->totals = $totals;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{health?: string, state?: string, stateReason?: string, lastEvaluationTime?: string, evaluationDuration?: float, lastError?: string, totals?: mixed} $inputData */
        $data = $inputData;
        return new self(
            health: isset($data["health"]) ? (function($input) { return \Grafana\Foundation\Alertingv1beta1\AlertRuleHealth::fromValue($input); })($data["health"]) : null,
            state: isset($data["state"]) ? (function($input) { return \Grafana\Foundation\Alertingv1beta1\AlertRuleState::fromValue($input); })($data["state"]) : null,
            stateReason: isset($data["stateReason"]) ? (function($input) { return \Grafana\Foundation\Alertingv1beta1\AlertRuleStateReason::fromValue($input); })($data["stateReason"]) : null,
            lastEvaluationTime: $data["lastEvaluationTime"] ?? null,
            evaluationDuration: $data["evaluationDuration"] ?? null,
            lastError: $data["lastError"] ?? null,
            totals: isset($data["totals"]) ? (function($input) {
    	/** @var array{healthy?: int, firing?: int, pending?: int, recovering?: int, nodata?: int, error?: int} */
    $val = $input;
    	return \Grafana\Foundation\Alertingv1beta1\AlertRuleInstanceTotals::fromArray($val);
    })($data["totals"]) : null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        if (isset($this->health)) {
            $data->health = $this->health;
        }
        if (isset($this->state)) {
            $data->state = $this->state;
        }
        if (isset($this->stateReason)) {
            $data->stateReason = $this->stateReason;
        }
        if (isset($this->lastEvaluationTime)) {
            $data->lastEvaluationTime = $this->lastEvaluationTime;
        }
        if (isset($this->evaluationDuration)) {
            $data->evaluationDuration = $this->evaluationDuration;
        }
        if (isset($this->lastError)) {
            $data->lastError = $this->lastError;
        }
        if (isset($this->totals)) {
            $data->totals = $this->totals;
        }
        return $data;
    }
}
