<?php

namespace Grafana\Foundation\Elasticsearch;

class BucketScript implements \JsonSerializable
{
    public \Grafana\Foundation\Elasticsearch\MetricAggregationType $type;

    /**
     * @var array<\Grafana\Foundation\Elasticsearch\PipelineVariable>|null
     */
    public ?array $pipelineVariables;

    public string $id;

    public ?\Grafana\Foundation\Elasticsearch\ElasticsearchBucketScriptSettings $settings;

    public ?bool $hide;

    /**
     * @param array<\Grafana\Foundation\Elasticsearch\PipelineVariable>|null $pipelineVariables
     * @param string|null $id
     * @param \Grafana\Foundation\Elasticsearch\ElasticsearchBucketScriptSettings|null $settings
     * @param bool|null $hide
     */
    public function __construct(?array $pipelineVariables = null, ?string $id = null, ?\Grafana\Foundation\Elasticsearch\ElasticsearchBucketScriptSettings $settings = null, ?bool $hide = null)
    {
        $this->type = \Grafana\Foundation\Elasticsearch\MetricAggregationType::count();
        $this->pipelineVariables = $pipelineVariables;
        $this->id = $id ?: "";
        $this->settings = $settings;
        $this->hide = $hide;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{type?: "bucket_script", pipelineVariables?: array<mixed>, id?: string, settings?: mixed, hide?: bool} $inputData */
        $data = $inputData;
        return new self(
            pipelineVariables: array_filter(array_map((function($input) {
    	/** @var array{name?: string, pipelineAgg?: string} */
    $val = $input;
    	return \Grafana\Foundation\Elasticsearch\PipelineVariable::fromArray($val);
    }), $data["pipelineVariables"] ?? [])),
            id: $data["id"] ?? null,
            settings: isset($data["settings"]) ? (function($input) {
    	/** @var array{script?: string|mixed} */
    $val = $input;
    	return \Grafana\Foundation\Elasticsearch\ElasticsearchBucketScriptSettings::fromArray($val);
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
        if (isset($this->pipelineVariables)) {
            $data["pipelineVariables"] = $this->pipelineVariables;
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
