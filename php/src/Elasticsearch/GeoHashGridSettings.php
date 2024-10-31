<?php

namespace Grafana\Foundation\Elasticsearch;

class GeoHashGridSettings implements \JsonSerializable
{
    public ?string $precision;

    /**
     * @param string|null $precision
     */
    public function __construct(?string $precision = null)
    {
        $this->precision = $precision;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{precision?: string} $inputData */
        $data = $inputData;
        return new self(
            precision: $data["precision"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
        ];
        if (isset($this->precision)) {
            $data["precision"] = $this->precision;
        }
        return $data;
    }
}
