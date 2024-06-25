<?php

namespace Grafana\Foundation\Candlestick;

class CandlestickFieldMap implements \JsonSerializable
{
    /**
     * Corresponds to the starting value of the given period
     */
    public ?string $open;

    /**
     * Corresponds to the highest value of the given period
     */
    public ?string $high;

    /**
     * Corresponds to the lowest value of the given period
     */
    public ?string $low;

    /**
     * Corresponds to the final (end) value of the given period
     */
    public ?string $close;

    /**
     * Corresponds to the sample count in the given period. (e.g. number of trades)
     */
    public ?string $volume;

    /**
     * @param string|null $open
     * @param string|null $high
     * @param string|null $low
     * @param string|null $close
     * @param string|null $volume
     */
    public function __construct(?string $open = null, ?string $high = null, ?string $low = null, ?string $close = null, ?string $volume = null)
    {
        $this->open = $open;
        $this->high = $high;
        $this->low = $low;
        $this->close = $close;
        $this->volume = $volume;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{open?: string, high?: string, low?: string, close?: string, volume?: string} $inputData */
        $data = $inputData;
        return new self(
            open: $data["open"] ?? null,
            high: $data["high"] ?? null,
            low: $data["low"] ?? null,
            close: $data["close"] ?? null,
            volume: $data["volume"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
        ];
        if (isset($this->open)) {
            $data["open"] = $this->open;
        }
        if (isset($this->high)) {
            $data["high"] = $this->high;
        }
        if (isset($this->low)) {
            $data["low"] = $this->low;
        }
        if (isset($this->close)) {
            $data["close"] = $this->close;
        }
        if (isset($this->volume)) {
            $data["volume"] = $this->volume;
        }
        return $data;
    }
}
