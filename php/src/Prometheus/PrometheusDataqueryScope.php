<?php

namespace Grafana\Foundation\Prometheus;

class PrometheusDataqueryScope implements \JsonSerializable
{
    public string $matchers;

    /**
     * @param string|null $matchers
     */
    public function __construct(?string $matchers = null)
    {
        $this->matchers = $matchers ?: "";
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{matchers?: string} $inputData */
        $data = $inputData;
        return new self(
            matchers: $data["matchers"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "matchers" => $this->matchers,
        ];
        return $data;
    }
}
