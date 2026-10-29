<?php

namespace Grafana\Foundation\Alertingv0alpha1;

class RuleSequenceSpec implements \JsonSerializable
{
    public \Grafana\Foundation\Alertingv0alpha1\IntervalTrigger $trigger;

    /**
     * @var array<\Grafana\Foundation\Alertingv0alpha1\RuleRef>
     */
    public array $recordingRules;

    /**
     * @var array<\Grafana\Foundation\Alertingv0alpha1\RuleRef>|null
     */
    public ?array $alertingRules;

    /**
     * @param \Grafana\Foundation\Alertingv0alpha1\IntervalTrigger|null $trigger
     * @param array<\Grafana\Foundation\Alertingv0alpha1\RuleRef>|null $recordingRules
     * @param array<\Grafana\Foundation\Alertingv0alpha1\RuleRef>|null $alertingRules
     */
    public function __construct(?\Grafana\Foundation\Alertingv0alpha1\IntervalTrigger $trigger = null, ?array $recordingRules = null, ?array $alertingRules = null)
    {
        $this->trigger = $trigger ?: new \Grafana\Foundation\Alertingv0alpha1\IntervalTrigger();
        $this->recordingRules = $recordingRules ?: [];
        $this->alertingRules = $alertingRules;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{trigger?: mixed, recordingRules?: array<mixed>, alertingRules?: array<mixed>} $inputData */
        $data = $inputData;
        return new self(
            trigger: isset($data["trigger"]) ? (function($input) {
    	/** @var array{interval?: string} */
    $val = $input;
    	return \Grafana\Foundation\Alertingv0alpha1\IntervalTrigger::fromArray($val);
    })($data["trigger"]) : null,
            recordingRules: array_filter(array_map((function($input) {
    	/** @var array{name?: string} */
    $val = $input;
    	return \Grafana\Foundation\Alertingv0alpha1\RuleRef::fromArray($val);
    }), $data["recordingRules"] ?? [])),
            alertingRules: array_filter(array_map((function($input) {
    	/** @var array{name?: string} */
    $val = $input;
    	return \Grafana\Foundation\Alertingv0alpha1\RuleRef::fromArray($val);
    }), $data["alertingRules"] ?? [])),
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
        if (isset($this->alertingRules)) {
            $data->alertingRules = $this->alertingRules;
        }
        return $data;
    }
}
