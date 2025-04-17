<?php

namespace Grafana\Foundation\Logsnew;

class Options implements \JsonSerializable
{
    public bool $showControls;

    public bool $showTime;

    public bool $wrapLogMessage;

    public bool $enableLogDetails;

    public bool $syntaxHighlighting;

    public \Grafana\Foundation\Common\LogsSortOrder $sortOrder;

    public \Grafana\Foundation\Common\LogsDedupStrategy $dedupStrategy;

    /**
     * @var mixed|null
     */
    public $grammar;

    public ?bool $enableInfiniteScrolling;

    /**
     * @var mixed|null
     */
    public $onLogOptionsChange;

    /**
     * @var mixed|null
     */
    public $onNewLogsReceived;

    /**
     * @param bool|null $showControls
     * @param bool|null $showTime
     * @param bool|null $wrapLogMessage
     * @param bool|null $enableLogDetails
     * @param bool|null $syntaxHighlighting
     * @param \Grafana\Foundation\Common\LogsSortOrder|null $sortOrder
     * @param \Grafana\Foundation\Common\LogsDedupStrategy|null $dedupStrategy
     * @param mixed|null $grammar
     * @param bool|null $enableInfiniteScrolling
     * @param mixed|null $onLogOptionsChange
     * @param mixed|null $onNewLogsReceived
     */
    public function __construct(?bool $showControls = null, ?bool $showTime = null, ?bool $wrapLogMessage = null, ?bool $enableLogDetails = null, ?bool $syntaxHighlighting = null, ?\Grafana\Foundation\Common\LogsSortOrder $sortOrder = null, ?\Grafana\Foundation\Common\LogsDedupStrategy $dedupStrategy = null,  $grammar = null, ?bool $enableInfiniteScrolling = null,  $onLogOptionsChange = null,  $onNewLogsReceived = null)
    {
        $this->showControls = $showControls ?: false;
        $this->showTime = $showTime ?: false;
        $this->wrapLogMessage = $wrapLogMessage ?: false;
        $this->enableLogDetails = $enableLogDetails ?: false;
        $this->syntaxHighlighting = $syntaxHighlighting ?: false;
        $this->sortOrder = $sortOrder ?: \Grafana\Foundation\Common\LogsSortOrder::Descending();
        $this->dedupStrategy = $dedupStrategy ?: \Grafana\Foundation\Common\LogsDedupStrategy::None();
        $this->grammar = $grammar;
        $this->enableInfiniteScrolling = $enableInfiniteScrolling;
        $this->onLogOptionsChange = $onLogOptionsChange;
        $this->onNewLogsReceived = $onNewLogsReceived;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{showControls?: bool, showTime?: bool, wrapLogMessage?: bool, enableLogDetails?: bool, syntaxHighlighting?: bool, sortOrder?: string, dedupStrategy?: string, grammar?: mixed, enableInfiniteScrolling?: bool, onLogOptionsChange?: mixed, onNewLogsReceived?: mixed} $inputData */
        $data = $inputData;
        return new self(
            showControls: $data["showControls"] ?? null,
            showTime: $data["showTime"] ?? null,
            wrapLogMessage: $data["wrapLogMessage"] ?? null,
            enableLogDetails: $data["enableLogDetails"] ?? null,
            syntaxHighlighting: $data["syntaxHighlighting"] ?? null,
            sortOrder: isset($data["sortOrder"]) ? (function($input) { return \Grafana\Foundation\Common\LogsSortOrder::fromValue($input); })($data["sortOrder"]) : null,
            dedupStrategy: isset($data["dedupStrategy"]) ? (function($input) { return \Grafana\Foundation\Common\LogsDedupStrategy::fromValue($input); })($data["dedupStrategy"]) : null,
            grammar: $data["grammar"] ?? null,
            enableInfiniteScrolling: $data["enableInfiniteScrolling"] ?? null,
            onLogOptionsChange: $data["onLogOptionsChange"] ?? null,
            onNewLogsReceived: $data["onNewLogsReceived"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->showControls = $this->showControls;
        $data->showTime = $this->showTime;
        $data->wrapLogMessage = $this->wrapLogMessage;
        $data->enableLogDetails = $this->enableLogDetails;
        $data->syntaxHighlighting = $this->syntaxHighlighting;
        $data->sortOrder = $this->sortOrder;
        $data->dedupStrategy = $this->dedupStrategy;
        if (isset($this->grammar)) {
            $data->grammar = $this->grammar;
        }
        if (isset($this->enableInfiniteScrolling)) {
            $data->enableInfiniteScrolling = $this->enableInfiniteScrolling;
        }
        if (isset($this->onLogOptionsChange)) {
            $data->onLogOptionsChange = $this->onLogOptionsChange;
        }
        if (isset($this->onNewLogsReceived)) {
            $data->onNewLogsReceived = $this->onNewLogsReceived;
        }
        return $data;
    }
}
