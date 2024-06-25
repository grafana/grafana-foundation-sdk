<?php

namespace Grafana\Foundation\Elasticsearch;

class Derivative implements \JsonSerializable
{
    public ?string $pipelineAgg;

    public ?string $field;

    public string $type;

    public string $id;

    public ?\Grafana\Foundation\Elasticsearch\ElasticsearchDerivativeSettings $settings;

    public ?bool $hide;

    /**
     * @param string|null $pipelineAgg
     * @param string|null $field
     * @param string|null $id
     * @param \Grafana\Foundation\Elasticsearch\ElasticsearchDerivativeSettings|null $settings
     * @param bool|null $hide
     */
    public function __construct(?string $pipelineAgg = null, ?string $field = null, ?string $id = null, ?\Grafana\Foundation\Elasticsearch\ElasticsearchDerivativeSettings $settings = null, ?bool $hide = null)
    {
        $this->pipelineAgg = $pipelineAgg;
        $this->field = $field;
        $this->type = "derivative";
    
        $this->id = $id ?: "";
        $this->settings = $settings;
        $this->hide = $hide;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{pipelineAgg?: string, field?: string, type?: string, id?: string, settings?: mixed, hide?: bool} $inputData */
        $data = $inputData;
        return new self(
            pipelineAgg: $data["pipelineAgg"] ?? null,
            field: $data["field"] ?? null,
            id: $data["id"] ?? null,
            settings: isset($data["settings"]) ? (function($input) {
    	/** @var array{unit?: string} */
    $val = $input;
    	return \Grafana\Foundation\Elasticsearch\ElasticsearchDerivativeSettings::fromArray($val);
    })($data["settings"]) : null,
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
