<?php

namespace Grafana\Foundation\Logs;

class Options implements \JsonSerializable
{
    public bool $showLabels;

    public bool $showCommonLabels;

    public bool $showTime;

    public bool $showLogContextToggle;

    public ?bool $showControls;

    public ?string $controlsStorageKey;

    public bool $wrapLogMessage;

    public bool $prettifyLogMessage;

    public bool $enableLogDetails;

    public \Grafana\Foundation\Common\LogsSortOrder $sortOrder;

    public \Grafana\Foundation\Common\LogsDedupStrategy $dedupStrategy;

    public ?bool $enableInfiniteScrolling;

    /**
     * TODO: figure out how to define callbacks
     * @var mixed|null
     */
    public $onClickFilterLabel;

    /**
     * @var mixed|null
     */
    public $onClickFilterOutLabel;

    /**
     * @var mixed|null
     */
    public $isFilterLabelActive;

    /**
     * @var mixed|null
     */
    public $onClickFilterString;

    /**
     * @var mixed|null
     */
    public $onClickFilterOutString;

    /**
     * @var mixed|null
     */
    public $onClickShowField;

    /**
     * @var mixed|null
     */
    public $onClickHideField;

    /**
     * @var mixed|null
     */
    public $onLogOptionsChange;

    /**
     * @var mixed|null
     */
    public $logRowMenuIconsBefore;

    /**
     * @var mixed|null
     */
    public $logRowMenuIconsAfter;

    /**
     * @var mixed|null
     */
    public $onNewLogsReceived;

    /**
     * @var array<string>|null
     */
    public ?array $displayedFields;

    /**
     * @param bool|null $showLabels
     * @param bool|null $showCommonLabels
     * @param bool|null $showTime
     * @param bool|null $showLogContextToggle
     * @param bool|null $showControls
     * @param string|null $controlsStorageKey
     * @param bool|null $wrapLogMessage
     * @param bool|null $prettifyLogMessage
     * @param bool|null $enableLogDetails
     * @param \Grafana\Foundation\Common\LogsSortOrder|null $sortOrder
     * @param \Grafana\Foundation\Common\LogsDedupStrategy|null $dedupStrategy
     * @param bool|null $enableInfiniteScrolling
     * @param mixed|null $onClickFilterLabel
     * @param mixed|null $onClickFilterOutLabel
     * @param mixed|null $isFilterLabelActive
     * @param mixed|null $onClickFilterString
     * @param mixed|null $onClickFilterOutString
     * @param mixed|null $onClickShowField
     * @param mixed|null $onClickHideField
     * @param mixed|null $onLogOptionsChange
     * @param mixed|null $logRowMenuIconsBefore
     * @param mixed|null $logRowMenuIconsAfter
     * @param mixed|null $onNewLogsReceived
     * @param array<string>|null $displayedFields
     */
    public function __construct(?bool $showLabels = null, ?bool $showCommonLabels = null, ?bool $showTime = null, ?bool $showLogContextToggle = null, ?bool $showControls = null, ?string $controlsStorageKey = null, ?bool $wrapLogMessage = null, ?bool $prettifyLogMessage = null, ?bool $enableLogDetails = null, ?\Grafana\Foundation\Common\LogsSortOrder $sortOrder = null, ?\Grafana\Foundation\Common\LogsDedupStrategy $dedupStrategy = null, ?bool $enableInfiniteScrolling = null,  $onClickFilterLabel = null,  $onClickFilterOutLabel = null,  $isFilterLabelActive = null,  $onClickFilterString = null,  $onClickFilterOutString = null,  $onClickShowField = null,  $onClickHideField = null,  $onLogOptionsChange = null,  $logRowMenuIconsBefore = null,  $logRowMenuIconsAfter = null,  $onNewLogsReceived = null, ?array $displayedFields = null)
    {
        $this->showLabels = $showLabels ?: false;
        $this->showCommonLabels = $showCommonLabels ?: false;
        $this->showTime = $showTime ?: false;
        $this->showLogContextToggle = $showLogContextToggle ?: false;
        $this->showControls = $showControls;
        $this->controlsStorageKey = $controlsStorageKey;
        $this->wrapLogMessage = $wrapLogMessage ?: false;
        $this->prettifyLogMessage = $prettifyLogMessage ?: false;
        $this->enableLogDetails = $enableLogDetails ?: false;
        $this->sortOrder = $sortOrder ?: \Grafana\Foundation\Common\LogsSortOrder::Descending();
        $this->dedupStrategy = $dedupStrategy ?: \Grafana\Foundation\Common\LogsDedupStrategy::None();
        $this->enableInfiniteScrolling = $enableInfiniteScrolling;
        $this->onClickFilterLabel = $onClickFilterLabel;
        $this->onClickFilterOutLabel = $onClickFilterOutLabel;
        $this->isFilterLabelActive = $isFilterLabelActive;
        $this->onClickFilterString = $onClickFilterString;
        $this->onClickFilterOutString = $onClickFilterOutString;
        $this->onClickShowField = $onClickShowField;
        $this->onClickHideField = $onClickHideField;
        $this->onLogOptionsChange = $onLogOptionsChange;
        $this->logRowMenuIconsBefore = $logRowMenuIconsBefore;
        $this->logRowMenuIconsAfter = $logRowMenuIconsAfter;
        $this->onNewLogsReceived = $onNewLogsReceived;
        $this->displayedFields = $displayedFields;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{showLabels?: bool, showCommonLabels?: bool, showTime?: bool, showLogContextToggle?: bool, showControls?: bool, controlsStorageKey?: string, wrapLogMessage?: bool, prettifyLogMessage?: bool, enableLogDetails?: bool, sortOrder?: string, dedupStrategy?: string, enableInfiniteScrolling?: bool, onClickFilterLabel?: mixed, onClickFilterOutLabel?: mixed, isFilterLabelActive?: mixed, onClickFilterString?: mixed, onClickFilterOutString?: mixed, onClickShowField?: mixed, onClickHideField?: mixed, onLogOptionsChange?: mixed, logRowMenuIconsBefore?: mixed, logRowMenuIconsAfter?: mixed, onNewLogsReceived?: mixed, displayedFields?: array<string>} $inputData */
        $data = $inputData;
        return new self(
            showLabels: $data["showLabels"] ?? null,
            showCommonLabels: $data["showCommonLabels"] ?? null,
            showTime: $data["showTime"] ?? null,
            showLogContextToggle: $data["showLogContextToggle"] ?? null,
            showControls: $data["showControls"] ?? null,
            controlsStorageKey: $data["controlsStorageKey"] ?? null,
            wrapLogMessage: $data["wrapLogMessage"] ?? null,
            prettifyLogMessage: $data["prettifyLogMessage"] ?? null,
            enableLogDetails: $data["enableLogDetails"] ?? null,
            sortOrder: isset($data["sortOrder"]) ? (function($input) { return \Grafana\Foundation\Common\LogsSortOrder::fromValue($input); })($data["sortOrder"]) : null,
            dedupStrategy: isset($data["dedupStrategy"]) ? (function($input) { return \Grafana\Foundation\Common\LogsDedupStrategy::fromValue($input); })($data["dedupStrategy"]) : null,
            enableInfiniteScrolling: $data["enableInfiniteScrolling"] ?? null,
            onClickFilterLabel: $data["onClickFilterLabel"] ?? null,
            onClickFilterOutLabel: $data["onClickFilterOutLabel"] ?? null,
            isFilterLabelActive: $data["isFilterLabelActive"] ?? null,
            onClickFilterString: $data["onClickFilterString"] ?? null,
            onClickFilterOutString: $data["onClickFilterOutString"] ?? null,
            onClickShowField: $data["onClickShowField"] ?? null,
            onClickHideField: $data["onClickHideField"] ?? null,
            onLogOptionsChange: $data["onLogOptionsChange"] ?? null,
            logRowMenuIconsBefore: $data["logRowMenuIconsBefore"] ?? null,
            logRowMenuIconsAfter: $data["logRowMenuIconsAfter"] ?? null,
            onNewLogsReceived: $data["onNewLogsReceived"] ?? null,
            displayedFields: $data["displayedFields"] ?? null,
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
        if (isset($this->showControls)) {
            $data->showControls = $this->showControls;
        }
        if (isset($this->controlsStorageKey)) {
            $data->controlsStorageKey = $this->controlsStorageKey;
        }
        if (isset($this->enableInfiniteScrolling)) {
            $data->enableInfiniteScrolling = $this->enableInfiniteScrolling;
        }
        if (isset($this->onClickFilterLabel)) {
            $data->onClickFilterLabel = $this->onClickFilterLabel;
        }
        if (isset($this->onClickFilterOutLabel)) {
            $data->onClickFilterOutLabel = $this->onClickFilterOutLabel;
        }
        if (isset($this->isFilterLabelActive)) {
            $data->isFilterLabelActive = $this->isFilterLabelActive;
        }
        if (isset($this->onClickFilterString)) {
            $data->onClickFilterString = $this->onClickFilterString;
        }
        if (isset($this->onClickFilterOutString)) {
            $data->onClickFilterOutString = $this->onClickFilterOutString;
        }
        if (isset($this->onClickShowField)) {
            $data->onClickShowField = $this->onClickShowField;
        }
        if (isset($this->onClickHideField)) {
            $data->onClickHideField = $this->onClickHideField;
        }
        if (isset($this->onLogOptionsChange)) {
            $data->onLogOptionsChange = $this->onLogOptionsChange;
        }
        if (isset($this->logRowMenuIconsBefore)) {
            $data->logRowMenuIconsBefore = $this->logRowMenuIconsBefore;
        }
        if (isset($this->logRowMenuIconsAfter)) {
            $data->logRowMenuIconsAfter = $this->logRowMenuIconsAfter;
        }
        if (isset($this->onNewLogsReceived)) {
            $data->onNewLogsReceived = $this->onNewLogsReceived;
        }
        if (isset($this->displayedFields)) {
            $data->displayedFields = $this->displayedFields;
        }
        return $data;
    }
}
