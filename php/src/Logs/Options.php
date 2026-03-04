<?php

namespace Grafana\Foundation\Logs;

class Options implements \JsonSerializable
{
    public bool $showLabels;

    public bool $showCommonLabels;

    public bool $showTime;

    public bool $showLogContextToggle;

    public bool $wrapLogMessage;

    public bool $prettifyLogMessage;

    public bool $enableLogDetails;

    public \Grafana\Foundation\Common\LogsSortOrder $sortOrder;

    public \Grafana\Foundation\Common\LogsDedupStrategy $dedupStrategy;

    public ?bool $enableInfiniteScrolling;

    /**
     * Display controls to jump to the last or first log line, and filters by log level.
     */
    public ?bool $showControls;

    /**
     * Show a component to manage the displayed fields from the logs.
     */
    public ?bool $showFieldSelector;

    /**
     * Use a predefined coloring scheme to highlight relevant parts of the log lines.
     */
    public ?bool $syntaxHighlighting;

    public ?\Grafana\Foundation\Logs\OptionsFontSize $fontSize;

    public ?\Grafana\Foundation\Logs\OptionsDetailsMode $detailsMode;

    /**
     * @param bool|null $showLabels
     * @param bool|null $showCommonLabels
     * @param bool|null $showTime
     * @param bool|null $showLogContextToggle
     * @param bool|null $wrapLogMessage
     * @param bool|null $prettifyLogMessage
     * @param bool|null $enableLogDetails
     * @param \Grafana\Foundation\Common\LogsSortOrder|null $sortOrder
     * @param \Grafana\Foundation\Common\LogsDedupStrategy|null $dedupStrategy
     * @param bool|null $enableInfiniteScrolling
     * @param bool|null $showControls
     * @param bool|null $showFieldSelector
     * @param bool|null $syntaxHighlighting
     * @param \Grafana\Foundation\Logs\OptionsFontSize|null $fontSize
     * @param \Grafana\Foundation\Logs\OptionsDetailsMode|null $detailsMode
     */
    public function __construct(?bool $showLabels = null, ?bool $showCommonLabels = null, ?bool $showTime = null, ?bool $showLogContextToggle = null, ?bool $wrapLogMessage = null, ?bool $prettifyLogMessage = null, ?bool $enableLogDetails = null, ?\Grafana\Foundation\Common\LogsSortOrder $sortOrder = null, ?\Grafana\Foundation\Common\LogsDedupStrategy $dedupStrategy = null, ?bool $enableInfiniteScrolling = null, ?bool $showControls = null, ?bool $showFieldSelector = null, ?bool $syntaxHighlighting = null, ?\Grafana\Foundation\Logs\OptionsFontSize $fontSize = null, ?\Grafana\Foundation\Logs\OptionsDetailsMode $detailsMode = null)
    {
        $this->showLabels = $showLabels ?: false;
        $this->showCommonLabels = $showCommonLabels ?: false;
        $this->showTime = $showTime ?: false;
        $this->showLogContextToggle = $showLogContextToggle ?: false;
        $this->wrapLogMessage = $wrapLogMessage ?: false;
        $this->prettifyLogMessage = $prettifyLogMessage ?: false;
        $this->enableLogDetails = $enableLogDetails ?: false;
        $this->sortOrder = $sortOrder ?: \Grafana\Foundation\Common\LogsSortOrder::Descending();
        $this->dedupStrategy = $dedupStrategy ?: \Grafana\Foundation\Common\LogsDedupStrategy::None();
        $this->enableInfiniteScrolling = $enableInfiniteScrolling;
        $this->showControls = $showControls;
        $this->showFieldSelector = $showFieldSelector;
        $this->syntaxHighlighting = $syntaxHighlighting;
        $this->fontSize = $fontSize;
        $this->detailsMode = $detailsMode;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{showLabels?: bool, showCommonLabels?: bool, showTime?: bool, showLogContextToggle?: bool, wrapLogMessage?: bool, prettifyLogMessage?: bool, enableLogDetails?: bool, sortOrder?: string, dedupStrategy?: string, enableInfiniteScrolling?: bool, showControls?: bool, showFieldSelector?: bool, syntaxHighlighting?: bool, fontSize?: string, detailsMode?: string} $inputData */
        $data = $inputData;
        return new self(
            showLabels: $data["showLabels"] ?? null,
            showCommonLabels: $data["showCommonLabels"] ?? null,
            showTime: $data["showTime"] ?? null,
            showLogContextToggle: $data["showLogContextToggle"] ?? null,
            wrapLogMessage: $data["wrapLogMessage"] ?? null,
            prettifyLogMessage: $data["prettifyLogMessage"] ?? null,
            enableLogDetails: $data["enableLogDetails"] ?? null,
            sortOrder: isset($data["sortOrder"]) ? (function($input) { return \Grafana\Foundation\Common\LogsSortOrder::fromValue($input); })($data["sortOrder"]) : null,
            dedupStrategy: isset($data["dedupStrategy"]) ? (function($input) { return \Grafana\Foundation\Common\LogsDedupStrategy::fromValue($input); })($data["dedupStrategy"]) : null,
            enableInfiniteScrolling: $data["enableInfiniteScrolling"] ?? null,
            showControls: $data["showControls"] ?? null,
            showFieldSelector: $data["showFieldSelector"] ?? null,
            syntaxHighlighting: $data["syntaxHighlighting"] ?? null,
            fontSize: isset($data["fontSize"]) ? (function($input) { return \Grafana\Foundation\Logs\OptionsFontSize::fromValue($input); })($data["fontSize"]) : null,
            detailsMode: isset($data["detailsMode"]) ? (function($input) { return \Grafana\Foundation\Logs\OptionsDetailsMode::fromValue($input); })($data["detailsMode"]) : null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->showLabels = $this->showLabels;
        $data->showCommonLabels = $this->showCommonLabels;
        $data->showTime = $this->showTime;
        $data->showLogContextToggle = $this->showLogContextToggle;
        $data->wrapLogMessage = $this->wrapLogMessage;
        $data->prettifyLogMessage = $this->prettifyLogMessage;
        $data->enableLogDetails = $this->enableLogDetails;
        $data->sortOrder = $this->sortOrder;
        $data->dedupStrategy = $this->dedupStrategy;
        if (isset($this->enableInfiniteScrolling)) {
            $data->enableInfiniteScrolling = $this->enableInfiniteScrolling;
        }
        if (isset($this->showControls)) {
            $data->showControls = $this->showControls;
        }
        if (isset($this->showFieldSelector)) {
            $data->showFieldSelector = $this->showFieldSelector;
        }
        if (isset($this->syntaxHighlighting)) {
            $data->syntaxHighlighting = $this->syntaxHighlighting;
        }
        if (isset($this->fontSize)) {
            $data->fontSize = $this->fontSize;
        }
        if (isset($this->detailsMode)) {
            $data->detailsMode = $this->detailsMode;
        }
        return $data;
    }
}
