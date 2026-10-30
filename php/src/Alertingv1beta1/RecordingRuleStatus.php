<?php

namespace Grafana\Foundation\Alertingv1beta1;

class RecordingRuleStatus implements \JsonSerializable
{
    public ?\Grafana\Foundation\Alertingv1beta1\RecordingRuleHealth $health;

    public ?string $lastEvaluationTime;

    /**
     * duration of the last evaluation in seconds
     */
    public ?float $evaluationDuration;

    public ?string $lastError;

    /**
     * @param \Grafana\Foundation\Alertingv1beta1\RecordingRuleHealth|null $health
     * @param string|null $lastEvaluationTime
     * @param float|null $evaluationDuration
     * @param string|null $lastError
     */
    public function __construct(?\Grafana\Foundation\Alertingv1beta1\RecordingRuleHealth $health = null, ?string $lastEvaluationTime = null, ?float $evaluationDuration = null, ?string $lastError = null)
    {
        $this->health = $health;
        $this->lastEvaluationTime = $lastEvaluationTime;
        $this->evaluationDuration = $evaluationDuration;
        $this->lastError = $lastError;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{health?: string, lastEvaluationTime?: string, evaluationDuration?: float, lastError?: string} $inputData */
        $data = $inputData;
        return new self(
            health: isset($data["health"]) ? (function($input) { return \Grafana\Foundation\Alertingv1beta1\RecordingRuleHealth::fromValue($input); })($data["health"]) : null,
            lastEvaluationTime: $data["lastEvaluationTime"] ?? null,
            evaluationDuration: $data["evaluationDuration"] ?? null,
            lastError: $data["lastError"] ?? null,
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
        if (isset($this->lastEvaluationTime)) {
            $data->lastEvaluationTime = $this->lastEvaluationTime;
        }
        if (isset($this->evaluationDuration)) {
            $data->evaluationDuration = $this->evaluationDuration;
        }
        if (isset($this->lastError)) {
            $data->lastError = $this->lastError;
        }
        return $data;
    }
}
