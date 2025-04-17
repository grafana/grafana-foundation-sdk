<?php

namespace Grafana\Foundation\Dashboard;

class DashboardDashboardTime implements \JsonSerializable
{
    public string $from;

    public string $to;

    /**
     * @param string|null $from
     * @param string|null $to
     */
    public function __construct(?string $from = null, ?string $to = null)
    {
        $this->from = $from ?: "now-6h";
        $this->to = $to ?: "now";
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{from?: string, to?: string} $inputData */
        $data = $inputData;
        return new self(
            from: $data["from"] ?? null,
            to: $data["to"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->from = $this->from;
        $data->to = $this->to;
        return $data;
    }
}
