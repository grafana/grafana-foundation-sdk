<?php

namespace Grafana\Foundation\Elasticsearch;

class PipelineMetricAggregationWithMultipleBucketPaths implements \JsonSerializable
{
    /**
     * @var array<\Grafana\Foundation\Elasticsearch\PipelineVariable>|null
     */
    public ?array $pipelineVariables;

    /**
     * @var string|\Grafana\Foundation\Elasticsearch\PipelineMetricAggregationType
     */
    public $type;

    public string $id;

    public ?bool $hide;

    /**
     * @param array<\Grafana\Foundation\Elasticsearch\PipelineVariable>|null $pipelineVariables
     * @param string|\Grafana\Foundation\Elasticsearch\PipelineMetricAggregationType|null $type
     * @param string|null $id
     * @param bool|null $hide
     */
    public function __construct(?array $pipelineVariables = null,  $type = null, ?string $id = null, ?bool $hide = null)
    {
        $this->pipelineVariables = $pipelineVariables;
        $this->type = $type ?: "count";
        $this->id = $id ?: "";
        $this->hide = $hide;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{pipelineVariables?: array<mixed>, type?: string|string, id?: string, hide?: bool} $inputData */
        $data = $inputData;
        return new self(
            pipelineVariables: array_filter(array_map((function($input) {
    	/** @var array{name?: string, pipelineAgg?: string} */
    $val = $input;
    	return \Grafana\Foundation\Elasticsearch\PipelineVariable::fromArray($val);
    }), $data["pipelineVariables"] ?? [])),
            type: isset($data["type"]) ? (function($input) {
        switch (true) {
        case is_string($input):
            return $input;
        default:
            /** @var string $input */
            return \Grafana\Foundation\Elasticsearch\PipelineMetricAggregationType::fromValue($input);
    }
    })($data["type"]) : null,
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
        if (isset($this->pipelineVariables)) {
            $data["pipelineVariables"] = $this->pipelineVariables;
        }
        if (isset($this->hide)) {
            $data["hide"] = $this->hide;
        }
        return $data;
    }
}
