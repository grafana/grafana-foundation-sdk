<?php

namespace Grafana\Foundation\Alertingv1beta1;

class RuleSequenceSpec implements \JsonSerializable
{
    public \Grafana\Foundation\Alertingv1beta1\IntervalTrigger $trigger;

    /**
     * @var array<\Grafana\Foundation\Alertingv1beta1\RuleRef>
     */
    public array $recordingRules;

    /**
     * @var array<\Grafana\Foundation\Alertingv1beta1\RuleRef>|null
     */
    public ?array $alertRules;

    /**
     * @param \Grafana\Foundation\Alertingv1beta1\IntervalTrigger|null $trigger
     * @param array<\Grafana\Foundation\Alertingv1beta1\RuleRef>|null $recordingRules
     * @param array<\Grafana\Foundation\Alertingv1beta1\RuleRef>|null $alertRules
     */
    public function __construct(?\Grafana\Foundation\Alertingv1beta1\IntervalTrigger $trigger = null, ?array $recordingRules = null, ?array $alertRules = null)
    {
        $this->trigger = $trigger ?: new \Grafana\Foundation\Alertingv1beta1\IntervalTrigger();
        $this->recordingRules = $recordingRules ?: [];
        $this->alertRules = $alertRules;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{trigger?: mixed, recordingRules?: array<mixed>, alertRules?: array<mixed>} $inputData */
        $data = $inputData;
        return new self(
            trigger: isset($data["trigger"]) ? (function($input) {
    	/** @var array{interval?: string} */
    $val = $input;
    	return \Grafana\Foundation\Alertingv1beta1\IntervalTrigger::fromArray($val);
    })($data["trigger"]) : null,
            recordingRules: array_filter(array_map((function($input) {
    	/** @var array{name?: string} */
    $val = $input;
    	return \Grafana\Foundation\Alertingv1beta1\RuleRef::fromArray($val);
    }), $data["recordingRules"] ?? [])),
            alertRules: array_filter(array_map((function($input) {
    	/** @var array{name?: string} */
    $val = $input;
    	return \Grafana\Foundation\Alertingv1beta1\RuleRef::fromArray($val);
    }), $data["alertRules"] ?? [])),
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->trigger = $this->trigger;
        $data->recordingRules = $this->recordingRules;
        if (isset($this->alertRules)) {
            $data->alertRules = $this->alertRules;
        }
        return $data;
    }
}
