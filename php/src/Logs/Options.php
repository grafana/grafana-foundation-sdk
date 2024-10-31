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
    public $logRowMenuIconsBefore;

    /**
     * @var mixed|null
     */
    public $logRowMenuIconsAfter;

    /**
     * @var array<string>|null
     */
    public ?array $displayedFields;

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
     * @param mixed|null $onClickFilterLabel
     * @param mixed|null $onClickFilterOutLabel
     * @param mixed|null $isFilterLabelActive
     * @param mixed|null $onClickFilterString
     * @param mixed|null $onClickFilterOutString
     * @param mixed|null $onClickShowField
     * @param mixed|null $onClickHideField
     * @param mixed|null $logRowMenuIconsBefore
     * @param mixed|null $logRowMenuIconsAfter
     * @param array<string>|null $displayedFields
     */
    public function __construct(?bool $showLabels = null, ?bool $showCommonLabels = null, ?bool $showTime = null, ?bool $showLogContextToggle = null, ?bool $wrapLogMessage = null, ?bool $prettifyLogMessage = null, ?bool $enableLogDetails = null, ?\Grafana\Foundation\Common\LogsSortOrder $sortOrder = null, ?\Grafana\Foundation\Common\LogsDedupStrategy $dedupStrategy = null,  $onClickFilterLabel = null,  $onClickFilterOutLabel = null,  $isFilterLabelActive = null,  $onClickFilterString = null,  $onClickFilterOutString = null,  $onClickShowField = null,  $onClickHideField = null,  $logRowMenuIconsBefore = null,  $logRowMenuIconsAfter = null, ?array $displayedFields = null)
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
        $this->onClickFilterLabel = $onClickFilterLabel;
        $this->onClickFilterOutLabel = $onClickFilterOutLabel;
        $this->isFilterLabelActive = $isFilterLabelActive;
        $this->onClickFilterString = $onClickFilterString;
        $this->onClickFilterOutString = $onClickFilterOutString;
        $this->onClickShowField = $onClickShowField;
        $this->onClickHideField = $onClickHideField;
        $this->logRowMenuIconsBefore = $logRowMenuIconsBefore;
        $this->logRowMenuIconsAfter = $logRowMenuIconsAfter;
        $this->displayedFields = $displayedFields;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{showLabels?: bool, showCommonLabels?: bool, showTime?: bool, showLogContextToggle?: bool, wrapLogMessage?: bool, prettifyLogMessage?: bool, enableLogDetails?: bool, sortOrder?: string, dedupStrategy?: string, onClickFilterLabel?: mixed, onClickFilterOutLabel?: mixed, isFilterLabelActive?: mixed, onClickFilterString?: mixed, onClickFilterOutString?: mixed, onClickShowField?: mixed, onClickHideField?: mixed, logRowMenuIconsBefore?: mixed, logRowMenuIconsAfter?: mixed, displayedFields?: array<string>} $inputData */
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
            onClickFilterLabel: $data["onClickFilterLabel"] ?? null,
            onClickFilterOutLabel: $data["onClickFilterOutLabel"] ?? null,
            isFilterLabelActive: $data["isFilterLabelActive"] ?? null,
            onClickFilterString: $data["onClickFilterString"] ?? null,
            onClickFilterOutString: $data["onClickFilterOutString"] ?? null,
            onClickShowField: $data["onClickShowField"] ?? null,
            onClickHideField: $data["onClickHideField"] ?? null,
            logRowMenuIconsBefore: $data["logRowMenuIconsBefore"] ?? null,
            logRowMenuIconsAfter: $data["logRowMenuIconsAfter"] ?? null,
            displayedFields: $data["displayedFields"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "showLabels" => $this->showLabels,
            "showCommonLabels" => $this->showCommonLabels,
            "showTime" => $this->showTime,
            "showLogContextToggle" => $this->showLogContextToggle,
            "wrapLogMessage" => $this->wrapLogMessage,
            "prettifyLogMessage" => $this->prettifyLogMessage,
            "enableLogDetails" => $this->enableLogDetails,
            "sortOrder" => $this->sortOrder,
            "dedupStrategy" => $this->dedupStrategy,
        ];
        if (isset($this->onClickFilterLabel)) {
            $data["onClickFilterLabel"] = $this->onClickFilterLabel;
        }
        if (isset($this->onClickFilterOutLabel)) {
            $data["onClickFilterOutLabel"] = $this->onClickFilterOutLabel;
        }
        if (isset($this->isFilterLabelActive)) {
            $data["isFilterLabelActive"] = $this->isFilterLabelActive;
        }
        if (isset($this->onClickFilterString)) {
            $data["onClickFilterString"] = $this->onClickFilterString;
        }
        if (isset($this->onClickFilterOutString)) {
            $data["onClickFilterOutString"] = $this->onClickFilterOutString;
        }
        if (isset($this->onClickShowField)) {
            $data["onClickShowField"] = $this->onClickShowField;
        }
        if (isset($this->onClickHideField)) {
            $data["onClickHideField"] = $this->onClickHideField;
        }
        if (isset($this->logRowMenuIconsBefore)) {
            $data["logRowMenuIconsBefore"] = $this->logRowMenuIconsBefore;
        }
        if (isset($this->logRowMenuIconsAfter)) {
            $data["logRowMenuIconsAfter"] = $this->logRowMenuIconsAfter;
        }
        if (isset($this->displayedFields)) {
            $data["displayedFields"] = $this->displayedFields;
        }
        return $data;
    }
}
