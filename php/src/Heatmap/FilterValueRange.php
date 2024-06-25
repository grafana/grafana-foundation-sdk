<?php

namespace Grafana\Foundation\Heatmap;

/**
 * Controls the value filter range
 */
class FilterValueRange implements \JsonSerializable
{
    /**
     * Sets the filter range to values less than or equal to the given value
     */
    public ?float $le;

    /**
     * Sets the filter range to values greater than or equal to the given value
     */
    public ?float $ge;

    /**
     * @param float|null $le
     * @param float|null $ge
     */
    public function __construct(?float $le = null, ?float $ge = null)
    {
        $this->le = $le;
        $this->ge = $ge;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{le?: float, ge?: float} $inputData */
        $data = $inputData;
        return new self(
            le: $data["le"] ?? null,
            ge: $data["ge"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
        ];
        if (isset($this->le)) {
            $data["le"] = $this->le;
        }
        if (isset($this->ge)) {
            $data["ge"] = $this->ge;
        }
        return $data;
    }
}
