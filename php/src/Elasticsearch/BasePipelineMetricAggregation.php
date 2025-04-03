<?php

namespace Grafana\Foundation\Elasticsearch;

class BasePipelineMetricAggregation implements \JsonSerializable
{
    public ?string $pipelineAgg;

    public ?string $field;

    public string $type;

    public string $id;

    public ?bool $hide;

    /**
     * @param string|null $pipelineAgg
     * @param string|null $field
     * @param string|null $type
     * @param string|null $id
     * @param bool|null $hide
     */
    public function __construct(?string $pipelineAgg = null, ?string $field = null, ?string $type = null, ?string $id = null, ?bool $hide = null)
    {
        $this->pipelineAgg = $pipelineAgg;
        $this->field = $field;
        $this->type = $type ?: "";
        $this->id = $id ?: "";
        $this->hide = $hide;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{pipelineAgg?: string, field?: string, type?: string, id?: string, hide?: bool} $inputData */
        $data = $inputData;
        return new self(
            pipelineAgg: $data["pipelineAgg"] ?? null,
            field: $data["field"] ?? null,
            type: $data["type"] ?? null,
            id: $data["id"] ?? null,
            hide: $data["hide"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "type" => $this->type,
            "id" => $this->id,
        ];
        if (isset($this->pipelineAgg)) {
            $data["pipelineAgg"] = $this->pipelineAgg;
        }
        if (isset($this->field)) {
            $data["field"] = $this->field;
        }
        if (isset($this->hide)) {
            $data["hide"] = $this->hide;
        }
        return $data;
    }
}
