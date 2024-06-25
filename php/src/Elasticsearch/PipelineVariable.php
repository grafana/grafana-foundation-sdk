<?php

namespace Grafana\Foundation\Elasticsearch;

class PipelineVariable implements \JsonSerializable
{
    public string $name;

    public string $pipelineAgg;

    /**
     * @param string|null $name
     * @param string|null $pipelineAgg
     */
    public function __construct(?string $name = null, ?string $pipelineAgg = null)
    {
        $this->name = $name ?: "";
        $this->pipelineAgg = $pipelineAgg ?: "";
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{name?: string, pipelineAgg?: string} $inputData */
        $data = $inputData;
        return new self(
            name: $data["name"] ?? null,
            pipelineAgg: $data["pipelineAgg"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "name" => $this->name,
            "pipelineAgg" => $this->pipelineAgg,
        ];
        return $data;
    }
}
