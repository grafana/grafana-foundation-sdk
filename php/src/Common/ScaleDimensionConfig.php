<?php

namespace Grafana\Foundation\Common;

class ScaleDimensionConfig implements \JsonSerializable
{
    public float $min;

    public float $max;

    public ?float $fixed;

    /**
     * fixed: T -- will be added by each element
     */
    public ?string $field;

    /**
     * | *"linear"
     */
    public ?\Grafana\Foundation\Common\ScaleDimensionMode $mode;

    /**
     * @param float|null $min
     * @param float|null $max
     * @param float|null $fixed
     * @param string|null $field
     * @param \Grafana\Foundation\Common\ScaleDimensionMode|null $mode
     */
    public function __construct(?float $min = null, ?float $max = null, ?float $fixed = null, ?string $field = null, ?\Grafana\Foundation\Common\ScaleDimensionMode $mode = null)
    {
        $this->min = $min ?: 0;
        $this->max = $max ?: 0;
        $this->fixed = $fixed;
        $this->field = $field;
        $this->mode = $mode;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{min?: float, max?: float, fixed?: float, field?: string, mode?: string} $inputData */
        $data = $inputData;
        return new self(
            min: $data["min"] ?? null,
            max: $data["max"] ?? null,
            fixed: $data["fixed"] ?? null,
            field: $data["field"] ?? null,
            mode: isset($data["mode"]) ? (function($input) { return \Grafana\Foundation\Common\ScaleDimensionMode::fromValue($input); })($data["mode"]) : null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "min" => $this->min,
            "max" => $this->max,
        ];
        if (isset($this->fixed)) {
            $data["fixed"] = $this->fixed;
        }
        if (isset($this->field)) {
            $data["field"] = $this->field;
        }
        if (isset($this->mode)) {
            $data["mode"] = $this->mode;
        }
        return $data;
    }
}
