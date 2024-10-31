<?php

namespace Grafana\Foundation\Elasticsearch;

class ElasticsearchTopMetricsSettings implements \JsonSerializable
{
    public ?string $order;

    public ?string $orderBy;

    /**
     * @var array<string>|null
     */
    public ?array $metrics;

    /**
     * @param string|null $order
     * @param string|null $orderBy
     * @param array<string>|null $metrics
     */
    public function __construct(?string $order = null, ?string $orderBy = null, ?array $metrics = null)
    {
        $this->order = $order;
        $this->orderBy = $orderBy;
        $this->metrics = $metrics;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{order?: string, orderBy?: string, metrics?: array<string>} $inputData */
        $data = $inputData;
        return new self(
            order: $data["order"] ?? null,
            orderBy: $data["orderBy"] ?? null,
            metrics: $data["metrics"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
        ];
        if (isset($this->order)) {
            $data["order"] = $this->order;
        }
        if (isset($this->orderBy)) {
            $data["orderBy"] = $this->orderBy;
        }
        if (isset($this->metrics)) {
            $data["metrics"] = $this->metrics;
        }
        return $data;
    }
}
