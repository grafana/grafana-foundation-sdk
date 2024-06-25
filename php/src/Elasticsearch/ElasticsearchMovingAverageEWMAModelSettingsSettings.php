<?php

namespace Grafana\Foundation\Elasticsearch;

class ElasticsearchMovingAverageEWMAModelSettingsSettings implements \JsonSerializable
{
    public ?string $alpha;

    /**
     * @param string|null $alpha
     */
    public function __construct(?string $alpha = null)
    {
        $this->alpha = $alpha;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{alpha?: string} $inputData */
        $data = $inputData;
        return new self(
            alpha: $data["alpha"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
        ];
        if (isset($this->alpha)) {
            $data["alpha"] = $this->alpha;
        }
        return $data;
    }
}
