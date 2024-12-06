<?php

namespace Grafana\Foundation\Elasticsearch;

class ElasticsearchDateHistogramSettings implements \JsonSerializable
{
    public ?string $interval;

    public ?string $minDocCount;

    public ?string $trimEdges;

    public ?string $offset;

    public ?string $timeZone;

    /**
     * @param string|null $interval
     * @param string|null $minDocCount
     * @param string|null $trimEdges
     * @param string|null $offset
     * @param string|null $timeZone
     */
    public function __construct(?string $interval = null, ?string $minDocCount = null, ?string $trimEdges = null, ?string $offset = null, ?string $timeZone = null)
    {
        $this->interval = $interval;
        $this->minDocCount = $minDocCount;
        $this->trimEdges = $trimEdges;
        $this->offset = $offset;
        $this->timeZone = $timeZone;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{interval?: string, min_doc_count?: string, trimEdges?: string, offset?: string, timeZone?: string} $inputData */
        $data = $inputData;
        return new self(
            interval: $data["interval"] ?? null,
            minDocCount: $data["min_doc_count"] ?? null,
            trimEdges: $data["trimEdges"] ?? null,
            offset: $data["offset"] ?? null,
            timeZone: $data["timeZone"] ?? null,
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
        if (isset($this->trimEdges)) {
            $data["trimEdges"] = $this->trimEdges;
        }
        if (isset($this->offset)) {
            $data["offset"] = $this->offset;
        }
        if (isset($this->timeZone)) {
            $data["timeZone"] = $this->timeZone;
        }
        return $data;
    }
}
