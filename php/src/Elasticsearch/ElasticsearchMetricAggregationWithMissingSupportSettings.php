<?php

namespace Grafana\Foundation\Elasticsearch;

class ElasticsearchMetricAggregationWithMissingSupportSettings implements \JsonSerializable
{
    public ?string $missing;

    /**
     * @param string|null $missing
     */
    public function __construct(?string $missing = null)
    {
        $this->missing = $missing;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{missing?: string} $inputData */
        $data = $inputData;
        return new self(
            missing: $data["missing"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
        ];
        if (isset($this->missing)) {
            $data["missing"] = $this->missing;
        }
        return $data;
    }
}
