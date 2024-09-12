<?php

namespace Grafana\Foundation\Alerting;

class RecordRule implements \JsonSerializable
{
    public string $from;

    public string $metric;

    /**
     * @param string|null $from
     * @param string|null $metric
     */
    public function __construct(?string $from = null, ?string $metric = null)
    {
        $this->from = $from ?: "";
        $this->metric = $metric ?: "";
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{from?: string, metric?: string} $inputData */
        $data = $inputData;
        return new self(
            from: $data["from"] ?? null,
            metric: $data["metric"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "from" => $this->from,
            "metric" => $this->metric,
        ];
        return $data;
    }
}
