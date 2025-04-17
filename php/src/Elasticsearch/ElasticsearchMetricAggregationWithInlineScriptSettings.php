<?php

namespace Grafana\Foundation\Elasticsearch;

class ElasticsearchMetricAggregationWithInlineScriptSettings implements \JsonSerializable
{
    /**
     * @var string|\Grafana\Foundation\Elasticsearch\ElasticsearchInlineScript
     */
    public $script;

    /**
     * @param string|\Grafana\Foundation\Elasticsearch\ElasticsearchInlineScript|null $script
     */
    public function __construct( $script = null)
    {
        $this->script = $script ?: "";
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{script?: string|mixed} $inputData */
        $data = $inputData;
        return new self(
            script: isset($data["script"]) ? (function($input) {
        switch (true) {
        case is_string($input):
            return $input;
        default:
            /** @var array{inline?: string} $input */
            return \Grafana\Foundation\Elasticsearch\ElasticsearchInlineScript::fromArray($input);
    }
    })($data["script"]) : null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->script = $this->script;
        return $data;
    }
}
