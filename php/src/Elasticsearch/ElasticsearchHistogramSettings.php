<?php

namespace Grafana\Foundation\Elasticsearch;

class ElasticsearchHistogramSettings implements \JsonSerializable
{
    public ?string $interval;

    public ?string $minDocCount;

    /**
     * @param string|null $interval
     * @param string|null $minDocCount
     */
    public function __construct(?string $interval = null, ?string $minDocCount = null)
    {
        $this->interval = $interval;
        $this->minDocCount = $minDocCount;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{interval?: string, min_doc_count?: string} $inputData */
        $data = $inputData;
        return new self(
            interval: $data["interval"] ?? null,
            minDocCount: $data["min_doc_count"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
        ];
        if (isset($this->interval)) {
            $data["interval"] = $this->interval;
        }
        if (isset($this->minDocCount)) {
            $data["min_doc_count"] = $this->minDocCount;
        }
        return $data;
    }
}
