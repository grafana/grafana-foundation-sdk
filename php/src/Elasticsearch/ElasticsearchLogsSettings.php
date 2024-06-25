<?php

namespace Grafana\Foundation\Elasticsearch;

class ElasticsearchLogsSettings implements \JsonSerializable
{
    public ?string $limit;

    /**
     * @param string|null $limit
     */
    public function __construct(?string $limit = null)
    {
        $this->limit = $limit;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{limit?: string} $inputData */
        $data = $inputData;
        return new self(
            limit: $data["limit"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
        ];
        if (isset($this->limit)) {
            $data["limit"] = $this->limit;
        }
        return $data;
    }
}
