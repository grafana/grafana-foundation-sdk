<?php

namespace Grafana\Foundation\Logsnew;

class Options implements \JsonSerializable
{
    public bool $showTime;

    public bool $wrapLogMessage;

    public bool $enableLogDetails;

    public \Grafana\Foundation\Common\LogsSortOrder $sortOrder;

    public \Grafana\Foundation\Common\LogsDedupStrategy $dedupStrategy;

    public ?bool $enableInfiniteScrolling;

    /**
     * @var mixed|null
     */
    public $onNewLogsReceived;

    /**
     * @param bool|null $showTime
     * @param bool|null $wrapLogMessage
     * @param bool|null $enableLogDetails
     * @param \Grafana\Foundation\Common\LogsSortOrder|null $sortOrder
     * @param \Grafana\Foundation\Common\LogsDedupStrategy|null $dedupStrategy
     * @param bool|null $enableInfiniteScrolling
     * @param mixed|null $onNewLogsReceived
     */
    public function __construct(?bool $showTime = null, ?bool $wrapLogMessage = null, ?bool $enableLogDetails = null, ?\Grafana\Foundation\Common\LogsSortOrder $sortOrder = null, ?\Grafana\Foundation\Common\LogsDedupStrategy $dedupStrategy = null, ?bool $enableInfiniteScrolling = null,  $onNewLogsReceived = null)
    {
        $this->showTime = $showTime ?: false;
        $this->wrapLogMessage = $wrapLogMessage ?: false;
        $this->enableLogDetails = $enableLogDetails ?: false;
        $this->sortOrder = $sortOrder ?: \Grafana\Foundation\Common\LogsSortOrder::Descending();
        $this->dedupStrategy = $dedupStrategy ?: \Grafana\Foundation\Common\LogsDedupStrategy::None();
        $this->enableInfiniteScrolling = $enableInfiniteScrolling;
        $this->onNewLogsReceived = $onNewLogsReceived;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{showTime?: bool, wrapLogMessage?: bool, enableLogDetails?: bool, sortOrder?: string, dedupStrategy?: string, enableInfiniteScrolling?: bool, onNewLogsReceived?: mixed} $inputData */
        $data = $inputData;
        return new self(
            showTime: $data["showTime"] ?? null,
            wrapLogMessage: $data["wrapLogMessage"] ?? null,
            enableLogDetails: $data["enableLogDetails"] ?? null,
            sortOrder: isset($data["sortOrder"]) ? (function($input) { return \Grafana\Foundation\Common\LogsSortOrder::fromValue($input); })($data["sortOrder"]) : null,
            dedupStrategy: isset($data["dedupStrategy"]) ? (function($input) { return \Grafana\Foundation\Common\LogsDedupStrategy::fromValue($input); })($data["dedupStrategy"]) : null,
            enableInfiniteScrolling: $data["enableInfiniteScrolling"] ?? null,
            onNewLogsReceived: $data["onNewLogsReceived"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->showTime = $this->showTime;
        $data->wrapLogMessage = $this->wrapLogMessage;
        $data->enableLogDetails = $this->enableLogDetails;
        $data->sortOrder = $this->sortOrder;
        $data->dedupStrategy = $this->dedupStrategy;
        if (isset($this->enableInfiniteScrolling)) {
            $data->enableInfiniteScrolling = $this->enableInfiniteScrolling;
        }
        if (isset($this->onNewLogsReceived)) {
            $data->onNewLogsReceived = $this->onNewLogsReceived;
        }
        return $data;
    }
}
