<?php

namespace Grafana\Foundation\Canvas;

class ConnectionCoordinates implements \JsonSerializable
{
    public float $x;

    public float $y;

    /**
     * @param float|null $x
     * @param float|null $y
     */
    public function __construct(?float $x = null, ?float $y = null)
    {
        $this->x = $x ?: 0;
        $this->y = $y ?: 0;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{x?: float, y?: float} $inputData */
        $data = $inputData;
        return new self(
            x: $data["x"] ?? null,
            y: $data["y"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->x = $this->x;
        $data->y = $this->y;
        return $data;
    }
}
