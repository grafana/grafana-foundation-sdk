<?php

namespace Grafana\Foundation\Elasticsearch;

class ElasticsearchGeoHashGridSettings implements \JsonSerializable
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
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        if (isset($this->precision)) {
            $data->precision = $this->precision;
        }
        return $data;
    }
}
