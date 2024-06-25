<?php

namespace Grafana\Foundation\Common;

/**
 * Gauge cell options
 */
class TableBarGaugeCellOptions implements \JsonSerializable
{
    public string $type;

    public ?\Grafana\Foundation\Common\BarGaugeDisplayMode $mode;

    public ?\Grafana\Foundation\Common\BarGaugeValueMode $valueDisplayMode;

    /**
     * @param \Grafana\Foundation\Common\BarGaugeDisplayMode|null $mode
     * @param \Grafana\Foundation\Common\BarGaugeValueMode|null $valueDisplayMode
     */
    public function __construct(?\Grafana\Foundation\Common\BarGaugeDisplayMode $mode = null, ?\Grafana\Foundation\Common\BarGaugeValueMode $valueDisplayMode = null)
    {
        $this->type = "gauge";
    
        $this->mode = $mode;
        $this->valueDisplayMode = $valueDisplayMode;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{type?: string, mode?: string, valueDisplayMode?: string} $inputData */
        $data = $inputData;
        return new self(
            mode: isset($data["mode"]) ? (function($input) { return \Grafana\Foundation\Common\BarGaugeDisplayMode::fromValue($input); })($data["mode"]) : null,
            valueDisplayMode: isset($data["valueDisplayMode"]) ? (function($input) { return \Grafana\Foundation\Common\BarGaugeValueMode::fromValue($input); })($data["valueDisplayMode"]) : null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "type" => $this->type,
        ];
        if (isset($this->mode)) {
            $data["mode"] = $this->mode;
        }
        if (isset($this->valueDisplayMode)) {
            $data["valueDisplayMode"] = $this->valueDisplayMode;
        }
        return $data;
    }
}
