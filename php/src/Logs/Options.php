<?php

namespace Grafana\Foundation\Logs;

class Options implements \JsonSerializable
{
    public bool $showLabels;

    public bool $showCommonLabels;

    public bool $showTime;

    public bool $wrapLogMessage;

    public bool $prettifyLogMessage;

    public bool $enableLogDetails;

    public \Grafana\Foundation\Common\LogsSortOrder $sortOrder;

    public \Grafana\Foundation\Common\LogsDedupStrategy $dedupStrategy;

    /**
     * @param bool|null $showLabels
     * @param bool|null $showCommonLabels
     * @param bool|null $showTime
     * @param bool|null $wrapLogMessage
     * @param bool|null $prettifyLogMessage
     * @param bool|null $enableLogDetails
     * @param \Grafana\Foundation\Common\LogsSortOrder|null $sortOrder
     * @param \Grafana\Foundation\Common\LogsDedupStrategy|null $dedupStrategy
     */
    public function __construct(?bool $showLabels = null, ?bool $showCommonLabels = null, ?bool $showTime = null, ?bool $wrapLogMessage = null, ?bool $prettifyLogMessage = null, ?bool $enableLogDetails = null, ?\Grafana\Foundation\Common\LogsSortOrder $sortOrder = null, ?\Grafana\Foundation\Common\LogsDedupStrategy $dedupStrategy = null)
    {
        $this->showLabels = $showLabels ?: false;
        $this->showCommonLabels = $showCommonLabels ?: false;
        $this->showTime = $showTime ?: false;
        $this->wrapLogMessage = $wrapLogMessage ?: false;
        $this->prettifyLogMessage = $prettifyLogMessage ?: false;
        $this->enableLogDetails = $enableLogDetails ?: false;
        $this->sortOrder = $sortOrder ?: \Grafana\Foundation\Common\LogsSortOrder::Descending();
        $this->dedupStrategy = $dedupStrategy ?: \Grafana\Foundation\Common\LogsDedupStrategy::None();
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{showLabels?: bool, showCommonLabels?: bool, showTime?: bool, wrapLogMessage?: bool, prettifyLogMessage?: bool, enableLogDetails?: bool, sortOrder?: string, dedupStrategy?: string} $inputData */
        $data = $inputData;
        return new self(
            showLabels: $data["showLabels"] ?? null,
            showCommonLabels: $data["showCommonLabels"] ?? null,
            showTime: $data["showTime"] ?? null,
            wrapLogMessage: $data["wrapLogMessage"] ?? null,
            prettifyLogMessage: $data["prettifyLogMessage"] ?? null,
            enableLogDetails: $data["enableLogDetails"] ?? null,
            sortOrder: isset($data["sortOrder"]) ? (function($input) { return \Grafana\Foundation\Common\LogsSortOrder::fromValue($input); })($data["sortOrder"]) : null,
            dedupStrategy: isset($data["dedupStrategy"]) ? (function($input) { return \Grafana\Foundation\Common\LogsDedupStrategy::fromValue($input); })($data["dedupStrategy"]) : null,
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
            "wrapLogMessage" => $this->wrapLogMessage,
            "prettifyLogMessage" => $this->prettifyLogMessage,
            "enableLogDetails" => $this->enableLogDetails,
            "sortOrder" => $this->sortOrder,
            "dedupStrategy" => $this->dedupStrategy,
        ];
        return $data;
    }
}
