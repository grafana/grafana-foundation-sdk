<?php

namespace Grafana\Foundation\Testdata;

class NodesQuery implements \JsonSerializable
{
    public ?int $count;

    public ?int $seed;

    /**
     * Possible enum values:
     *  - `"random"` 
     *  - `"random edges"` 
     *  - `"response_medium"` 
     *  - `"response_small"` 
     *  - `"feature_showcase"` 
     */
    public ?\Grafana\Foundation\Testdata\NodesQueryType $type;

    /**
     * @param int|null $count
     * @param int|null $seed
     * @param \Grafana\Foundation\Testdata\NodesQueryType|null $type
     */
    public function __construct(?int $count = null, ?int $seed = null, ?\Grafana\Foundation\Testdata\NodesQueryType $type = null)
    {
        $this->count = $count;
        $this->seed = $seed;
        $this->type = $type;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{count?: int, seed?: int, type?: string} $inputData */
        $data = $inputData;
        return new self(
            count: $data["count"] ?? null,
            seed: $data["seed"] ?? null,
            type: isset($data["type"]) ? (function($input) { return \Grafana\Foundation\Testdata\NodesQueryType::fromValue($input); })($data["type"]) : null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
        ];
        if (isset($this->count)) {
            $data["count"] = $this->count;
        }
        if (isset($this->seed)) {
            $data["seed"] = $this->seed;
        }
        if (isset($this->type)) {
            $data["type"] = $this->type;
        }
        return $data;
    }
}
