<?php

namespace Grafana\Foundation\Statushistory;

class FieldConfig implements \JsonSerializable
{
    public ?int $lineWidth;

    public ?\Grafana\Foundation\Common\HideSeriesConfig $hideFrom;

    public ?int $fillOpacity;

    /**
     * @param int|null $lineWidth
     * @param \Grafana\Foundation\Common\HideSeriesConfig|null $hideFrom
     * @param int|null $fillOpacity
     */
    public function __construct(?int $lineWidth = null, ?\Grafana\Foundation\Common\HideSeriesConfig $hideFrom = null, ?int $fillOpacity = null)
    {
        $this->lineWidth = $lineWidth;
        $this->hideFrom = $hideFrom;
        $this->fillOpacity = $fillOpacity;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{lineWidth?: int, hideFrom?: mixed, fillOpacity?: int} $inputData */
        $data = $inputData;
        return new self(
            lineWidth: $data["lineWidth"] ?? null,
            hideFrom: isset($data["hideFrom"]) ? (function($input) {
    	/** @var array{tooltip?: bool, legend?: bool, viz?: bool} */
    $val = $input;
    	return \Grafana\Foundation\Common\HideSeriesConfig::fromArray($val);
    })($data["hideFrom"]) : null,
            fillOpacity: $data["fillOpacity"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        if (isset($this->lineWidth)) {
            $data->lineWidth = $this->lineWidth;
        }
        if (isset($this->hideFrom)) {
            $data->hideFrom = $this->hideFrom;
        }
        if (isset($this->fillOpacity)) {
            $data->fillOpacity = $this->fillOpacity;
        }
        return $data;
    }
}
