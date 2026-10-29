<?php

namespace Grafana\Foundation\Alertingv1beta1;

class AlertRuleSpec implements \JsonSerializable
{
    public string $title;

    public ?bool $paused;

    public \Grafana\Foundation\Alertingv1beta1\IntervalTrigger $trigger;

    /**
     * @var array<string, string>|null
     */
    public ?array $labels;

    /**
     * @var array<string, string>|null
     */
    public ?array $annotations;

    public ?string $for;

    public ?string $keepFiringFor;

    public ?int $missingSeriesEvalsToResolve;

    public \Grafana\Foundation\Alertingv1beta1\NoDataState $noDataState;

    public \Grafana\Foundation\Alertingv1beta1\ExecErrState $execErrState;

    /**
     * @var \Grafana\Foundation\Alertingv1beta1\SimplifiedRouting|\Grafana\Foundation\Alertingv1beta1\NamedRoutingTree
     */
    public $notificationSettings;

    /**
     * @var array<string, \Grafana\Foundation\Alertingv1beta1\Expression>
     */
    public array $expressions;

    public ?\Grafana\Foundation\Alertingv1beta1\PanelRef $panelRef;

    /**
     * @param string|null $title
     * @param bool|null $paused
     * @param \Grafana\Foundation\Alertingv1beta1\IntervalTrigger|null $trigger
     * @param array<string, string>|null $labels
     * @param array<string, string>|null $annotations
     * @param string|null $for
     * @param string|null $keepFiringFor
     * @param int|null $missingSeriesEvalsToResolve
     * @param \Grafana\Foundation\Alertingv1beta1\NoDataState|null $noDataState
     * @param \Grafana\Foundation\Alertingv1beta1\ExecErrState|null $execErrState
     * @param \Grafana\Foundation\Alertingv1beta1\SimplifiedRouting|\Grafana\Foundation\Alertingv1beta1\NamedRoutingTree|null $notificationSettings
     * @param array<string, \Grafana\Foundation\Alertingv1beta1\Expression>|null $expressions
     * @param \Grafana\Foundation\Alertingv1beta1\PanelRef|null $panelRef
     */
    public function __construct(?string $title = null, ?bool $paused = null, ?\Grafana\Foundation\Alertingv1beta1\IntervalTrigger $trigger = null, ?array $labels = null, ?array $annotations = null, ?string $for = null, ?string $keepFiringFor = null, ?int $missingSeriesEvalsToResolve = null, ?\Grafana\Foundation\Alertingv1beta1\NoDataState $noDataState = null, ?\Grafana\Foundation\Alertingv1beta1\ExecErrState $execErrState = null,  $notificationSettings = null, ?array $expressions = null, ?\Grafana\Foundation\Alertingv1beta1\PanelRef $panelRef = null)
    {
        $this->title = $title ?: "";
        $this->paused = $paused;
        $this->trigger = $trigger ?: new \Grafana\Foundation\Alertingv1beta1\IntervalTrigger();
        $this->labels = $labels;
        $this->annotations = $annotations;
        $this->for = $for;
        $this->keepFiringFor = $keepFiringFor;
        $this->missingSeriesEvalsToResolve = $missingSeriesEvalsToResolve;
        $this->noDataState = $noDataState ?: \Grafana\Foundation\Alertingv1beta1\NoDataState::noData();
        $this->execErrState = $execErrState ?: \Grafana\Foundation\Alertingv1beta1\ExecErrState::error();
        $this->notificationSettings = $notificationSettings ?: new \Grafana\Foundation\Alertingv1beta1\SimplifiedRouting();
        $this->expressions = $expressions ?: [];
        $this->panelRef = $panelRef;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{title?: string, paused?: bool, trigger?: mixed, labels?: array<string, string>, annotations?: array<string, string>, for?: string, keepFiringFor?: string, missingSeriesEvalsToResolve?: int, noDataState?: string, execErrState?: string, notificationSettings?: mixed|mixed, expressions?: array<string, mixed>, panelRef?: mixed} $inputData */
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
            annotations: $data["annotations"] ?? null,
            for: $data["for"] ?? null,
            keepFiringFor: $data["keepFiringFor"] ?? null,
            missingSeriesEvalsToResolve: $data["missingSeriesEvalsToResolve"] ?? null,
            noDataState: isset($data["noDataState"]) ? (function($input) { return \Grafana\Foundation\Alertingv1beta1\NoDataState::fromValue($input); })($data["noDataState"]) : null,
            execErrState: isset($data["execErrState"]) ? (function($input) { return \Grafana\Foundation\Alertingv1beta1\ExecErrState::fromValue($input); })($data["execErrState"]) : null,
            notificationSettings: isset($data["notificationSettings"]) ? (function($input) {
        \assert(is_array($input), 'expected disjunction value to be an array');
        /** @var array<string, mixed> $input */
        switch ($input["type"]) {
        case "NamedRoutingTree":
            return NamedRoutingTree::fromArray($input);
        case "SimplifiedRouting":
            return SimplifiedRouting::fromArray($input);
        default:
            throw new \ValueError('can not parse disjunction from array');
    }
    })($data["notificationSettings"]) : null,
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
            panelRef: isset($data["panelRef"]) ? (function($input) {
    	/** @var array{dashboardUID?: string, panelID?: int} */
    $val = $input;
    	return \Grafana\Foundation\Alertingv1beta1\PanelRef::fromArray($val);
    })($data["panelRef"]) : null,
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
        $data->noDataState = $this->noDataState;
        $data->execErrState = $this->execErrState;
        $data->notificationSettings = $this->notificationSettings;
        $data->expressions = $this->expressions;
        if (isset($this->paused)) {
            $data->paused = $this->paused;
        }
        if (isset($this->labels)) {
            $data->labels = $this->labels;
        }
        if (isset($this->annotations)) {
            $data->annotations = $this->annotations;
        }
        if (isset($this->for)) {
            $data->for = $this->for;
        }
        if (isset($this->keepFiringFor)) {
            $data->keepFiringFor = $this->keepFiringFor;
        }
        if (isset($this->missingSeriesEvalsToResolve)) {
            $data->missingSeriesEvalsToResolve = $this->missingSeriesEvalsToResolve;
        }
        if (isset($this->panelRef)) {
            $data->panelRef = $this->panelRef;
        }
        return $data;
    }
}
