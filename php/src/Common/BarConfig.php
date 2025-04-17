<?php

namespace Grafana\Foundation\Common;

/**
 * TODO docs
 */
class BarConfig implements \JsonSerializable
{
    public ?\Grafana\Foundation\Common\BarAlignment $barAlignment;

    public ?float $barWidthFactor;

    public ?float $barMaxWidth;

    /**
     * @param \Grafana\Foundation\Common\BarAlignment|null $barAlignment
     * @param float|null $barWidthFactor
     * @param float|null $barMaxWidth
     */
    public function __construct(?\Grafana\Foundation\Common\BarAlignment $barAlignment = null, ?float $barWidthFactor = null, ?float $barMaxWidth = null)
    {
        $this->barAlignment = $barAlignment;
        $this->barWidthFactor = $barWidthFactor;
        $this->barMaxWidth = $barMaxWidth;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{barAlignment?: int, barWidthFactor?: float, barMaxWidth?: float} $inputData */
        $data = $inputData;
        return new self(
            barAlignment: isset($data["barAlignment"]) ? (function($input) { return \Grafana\Foundation\Common\BarAlignment::fromValue($input); })($data["barAlignment"]) : null,
            barWidthFactor: $data["barWidthFactor"] ?? null,
            barMaxWidth: $data["barMaxWidth"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        if (isset($this->barAlignment)) {
            $data->barAlignment = $this->barAlignment;
        }
        if (isset($this->barWidthFactor)) {
            $data->barWidthFactor = $this->barWidthFactor;
        }
        if (isset($this->barMaxWidth)) {
            $data->barMaxWidth = $this->barMaxWidth;
        }
        return $data;
    }
}
