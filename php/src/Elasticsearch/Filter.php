<?php

namespace Grafana\Foundation\Elasticsearch;

class Filter implements \JsonSerializable
{
    public string $query;

    public string $label;

    /**
     * @param string|null $query
     * @param string|null $label
     */
    public function __construct(?string $query = null, ?string $label = null)
    {
        $this->query = $query ?: "";
        $this->label = $label ?: "";
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{query?: string, label?: string} $inputData */
        $data = $inputData;
        return new self(
            query: $data["query"] ?? null,
            label: $data["label"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "query" => $this->query,
            "label" => $this->label,
        ];
        return $data;
    }
}
