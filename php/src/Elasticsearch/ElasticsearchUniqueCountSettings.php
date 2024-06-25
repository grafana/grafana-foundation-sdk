<?php

namespace Grafana\Foundation\Elasticsearch;

class ElasticsearchUniqueCountSettings implements \JsonSerializable
{
    public ?string $precisionThreshold;

    public ?string $missing;

    /**
     * @param string|null $precisionThreshold
     * @param string|null $missing
     */
    public function __construct(?string $precisionThreshold = null, ?string $missing = null)
    {
        $this->precisionThreshold = $precisionThreshold;
        $this->missing = $missing;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{precision_threshold?: string, missing?: string} $inputData */
        $data = $inputData;
        return new self(
            precisionThreshold: $data["precision_threshold"] ?? null,
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
        if (isset($this->precisionThreshold)) {
            $data["precision_threshold"] = $this->precisionThreshold;
        }
        if (isset($this->missing)) {
            $data["missing"] = $this->missing;
        }
        return $data;
    }
}
