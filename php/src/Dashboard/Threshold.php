<?php

namespace Grafana\Foundation\Dashboard;

/**
 * User-defined value for a metric that triggers visual changes in a panel when this value is met or exceeded
 * They are used to conditionally style and color visualizations based on query results , and can be applied to most visualizations.
 */
class Threshold implements \JsonSerializable
{
    /**
     * Value represents a specified metric for the threshold, which triggers a visual change in the dashboard when this value is met or exceeded.
     * Nulls currently appear here when serializing -Infinity to JSON.
     */
    public ?float $value;

    /**
     * Color represents the color of the visual change that will occur in the dashboard when the threshold value is met or exceeded.
     */
    public string $color;

    /**
     * @param float|null $value
     * @param string|null $color
     */
    public function __construct(?float $value = null, ?string $color = null)
    {
        $this->value = $value;
        $this->color = $color ?: "";
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{value?: float, color?: string} $inputData */
        $data = $inputData;
        return new self(
            value: $data["value"] ?? null,
            color: $data["color"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "color" => $this->color,
        ];
        if (isset($this->value)) {
            $data["value"] = $this->value;
        }
        return $data;
    }
}
