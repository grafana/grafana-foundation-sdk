<?php

namespace Grafana\Foundation\Elasticsearch;

class ElasticsearchSerialDiffSettings implements \JsonSerializable
{
    public ?string $lag;

    /**
     * @param string|null $lag
     */
    public function __construct(?string $lag = null)
    {
        $this->lag = $lag;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{lag?: string} $inputData */
        $data = $inputData;
        return new self(
            lag: $data["lag"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
        ];
        if (isset($this->lag)) {
            $data["lag"] = $this->lag;
        }
        return $data;
    }
}
