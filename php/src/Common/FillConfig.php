<?php

namespace Grafana\Foundation\Common;

/**
 * TODO docs
 */
class FillConfig implements \JsonSerializable
{
    public ?string $fillColor;

    public ?float $fillOpacity;

    public ?string $fillBelowTo;

    /**
     * @param string|null $fillColor
     * @param float|null $fillOpacity
     * @param string|null $fillBelowTo
     */
    public function __construct(?string $fillColor = null, ?float $fillOpacity = null, ?string $fillBelowTo = null)
    {
        $this->fillColor = $fillColor;
        $this->fillOpacity = $fillOpacity;
        $this->fillBelowTo = $fillBelowTo;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{fillColor?: string, fillOpacity?: float, fillBelowTo?: string} $inputData */
        $data = $inputData;
        return new self(
            fillColor: $data["fillColor"] ?? null,
            fillOpacity: $data["fillOpacity"] ?? null,
            fillBelowTo: $data["fillBelowTo"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
        ];
        if (isset($this->fillColor)) {
            $data["fillColor"] = $this->fillColor;
        }
        if (isset($this->fillOpacity)) {
            $data["fillOpacity"] = $this->fillOpacity;
        }
        if (isset($this->fillBelowTo)) {
            $data["fillBelowTo"] = $this->fillBelowTo;
        }
        return $data;
    }
}
