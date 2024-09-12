<?php

namespace Grafana\Foundation\Heatmap;

/**
 * Controls cell value options
 */
class CellValues implements \JsonSerializable
{
    /**
     * Controls the cell value unit
     */
    public ?string $unit;

    /**
     * Controls the number of decimals for cell values
     */
    public ?float $decimals;

    /**
     * @param string|null $unit
     * @param float|null $decimals
     */
    public function __construct(?string $unit = null, ?float $decimals = null)
    {
        $this->unit = $unit;
        $this->decimals = $decimals;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{unit?: string, decimals?: float} $inputData */
        $data = $inputData;
        return new self(
            unit: $data["unit"] ?? null,
            decimals: $data["decimals"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
        ];
        if (isset($this->unit)) {
            $data["unit"] = $this->unit;
        }
        if (isset($this->decimals)) {
            $data["decimals"] = $this->decimals;
        }
        return $data;
    }
}
