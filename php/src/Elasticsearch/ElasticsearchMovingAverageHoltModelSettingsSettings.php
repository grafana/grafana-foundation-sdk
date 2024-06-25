<?php

namespace Grafana\Foundation\Elasticsearch;

class ElasticsearchMovingAverageHoltModelSettingsSettings implements \JsonSerializable
{
    public ?string $alpha;

    public ?string $beta;

    /**
     * @param string|null $alpha
     * @param string|null $beta
     */
    public function __construct(?string $alpha = null, ?string $beta = null)
    {
        $this->alpha = $alpha;
        $this->beta = $beta;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{alpha?: string, beta?: string} $inputData */
        $data = $inputData;
        return new self(
            alpha: $data["alpha"] ?? null,
            beta: $data["beta"] ?? null,
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
        if (isset($this->beta)) {
            $data["beta"] = $this->beta;
        }
        return $data;
    }
}
