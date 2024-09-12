<?php

namespace Grafana\Foundation\Candlestick;

class CandlestickColors implements \JsonSerializable
{
    public string $up;

    public string $down;

    public string $flat;

    /**
     * @param string|null $up
     * @param string|null $down
     * @param string|null $flat
     */
    public function __construct(?string $up = null, ?string $down = null, ?string $flat = null)
    {
        $this->up = $up ?: "green";
        $this->down = $down ?: "red";
        $this->flat = $flat ?: "gray";
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{up?: string, down?: string, flat?: string} $inputData */
        $data = $inputData;
        return new self(
            up: $data["up"] ?? null,
            down: $data["down"] ?? null,
            flat: $data["flat"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "up" => $this->up,
            "down" => $this->down,
            "flat" => $this->flat,
        ];
        return $data;
    }
}
