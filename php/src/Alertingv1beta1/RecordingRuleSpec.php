<?php

namespace Grafana\Foundation\Alertingv1beta1;

class RecordingRuleSpec implements \JsonSerializable
{
    public string $title;

    public ?bool $paused;

    public \Grafana\Foundation\Alertingv1beta1\IntervalTrigger $trigger;

    /**
     * @var array<string, string>|null
     */
    public ?array $labels;

    public string $metric;

    /**
     * @var array<string, \Grafana\Foundation\Alertingv1beta1\Expression>
     */
    public array $expressions;

    public string $targetDatasourceUID;

    /**
     * @param string|null $title
     * @param bool|null $paused
     * @param \Grafana\Foundation\Alertingv1beta1\IntervalTrigger|null $trigger
     * @param array<string, string>|null $labels
     * @param string|null $metric
     * @param array<string, \Grafana\Foundation\Alertingv1beta1\Expression>|null $expressions
     * @param string|null $targetDatasourceUID
     */
    public function __construct(?string $title = null, ?bool $paused = null, ?\Grafana\Foundation\Alertingv1beta1\IntervalTrigger $trigger = null, ?array $labels = null, ?string $metric = null, ?array $expressions = null, ?string $targetDatasourceUID = null)
    {
        $this->title = $title ?: "";
        $this->paused = $paused;
        $this->trigger = $trigger ?: new \Grafana\Foundation\Alertingv1beta1\IntervalTrigger();
        $this->labels = $labels;
        $this->metric = $metric ?: "";
        $this->expressions = $expressions ?: [];
        $this->targetDatasourceUID = $targetDatasourceUID ?: "";
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{title?: string, paused?: bool, trigger?: mixed, labels?: array<string, string>, metric?: string, expressions?: array<string, mixed>, targetDatasourceUID?: string} $inputData */
        $data = $inputData;
        return new self(
            title: $data["title"] ?? null,
            paused: $data["paused"] ?? null,
            trigger: isset($data["trigger"]) ? (function($input) {
    	/** @var array{interval?: string} */
    $val = $input;
    	return \Grafana\Foundation\Alertingv1beta1\IntervalTrigger::fromArray($val);
    })($data["trigger"]) : null,
            labels: $data["labels"] ?? null,
            metric: $data["metric"] ?? null,
            expressions: isset($data["expressions"]) ? (function($input) {
        /** @var array<string, \Grafana\Foundation\Alertingv1beta1\Expression> $results */
        $results = [];
        foreach ($input as $key => $val) {
            $results[$key] = isset($val) ? (function($input) {
    	/** @var array{queryType?: string, relativeTimeRange?: mixed, datasourceUID?: string, model?: mixed, source?: bool} */
    $val = $input;
    	return \Grafana\Foundation\Alertingv1beta1\Expression::fromArray($val);
    })($val) : null;
        }
        return array_filter($results);
    })($data["expressions"]) : null,
            targetDatasourceUID: $data["targetDatasourceUID"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->title = $this->title;
        $data->trigger = $this->trigger;
        $data->metric = $this->metric;
        $data->expressions = $this->expressions;
        $data->targetDatasourceUID = $this->targetDatasourceUID;
        if (isset($this->paused)) {
            $data->paused = $this->paused;
        }
        if (isset($this->labels)) {
            $data->labels = $this->labels;
        }
        return $data;
    }
}
