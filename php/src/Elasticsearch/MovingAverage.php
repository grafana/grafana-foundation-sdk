<?php

namespace Grafana\Foundation\Elasticsearch;

/**
 * #MovingAverage's settings are overridden in types.ts
 */
class MovingAverage implements \JsonSerializable
{
    public ?string $pipelineAgg;

    public ?string $field;

    public string $type;

    public string $id;

    /**
     * @var array<string, mixed>|null
     */
    public ?array $settings;

    public ?bool $hide;

    /**
     * @param string|null $pipelineAgg
     * @param string|null $field
     * @param string|null $id
     * @param array<string, mixed>|null $settings
     * @param bool|null $hide
     */
    public function __construct(?string $pipelineAgg = null, ?string $field = null, ?string $id = null, ?array $settings = null, ?bool $hide = null)
    {
        $this->pipelineAgg = $pipelineAgg;
        $this->field = $field;
        $this->type = "moving_avg";
    
        $this->id = $id ?: "";
        $this->settings = $settings;
        $this->hide = $hide;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{pipelineAgg?: string, field?: string, type?: string, id?: string, settings?: array<string, mixed>, hide?: bool} $inputData */
        $data = $inputData;
        return new self(
            pipelineAgg: $data["pipelineAgg"] ?? null,
            field: $data["field"] ?? null,
            id: $data["id"] ?? null,
            settings: $data["settings"] ?? null,
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
        if (isset($this->settings)) {
            $data["settings"] = $this->settings;
        }
        if (isset($this->hide)) {
            $data["hide"] = $this->hide;
        }
        return $data;
    }
}
