<?php

namespace Grafana\Foundation\Elasticsearch;

class ElasticsearchPercentilesSettings implements \JsonSerializable
{
    /**
     * @var string|\Grafana\Foundation\Elasticsearch\ElasticsearchInlineScript
     */
    public $script;

    public ?string $missing;

    /**
     * @var array<string>|null
     */
    public ?array $percents;

    /**
     * @param string|\Grafana\Foundation\Elasticsearch\ElasticsearchInlineScript|null $script
     * @param string|null $missing
     * @param array<string>|null $percents
     */
    public function __construct( $script = null, ?string $missing = null, ?array $percents = null)
    {
        $this->script = $script ?: "";
        $this->missing = $missing;
        $this->percents = $percents;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{script?: string|mixed, missing?: string, percents?: array<string>} $inputData */
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
            missing: $data["missing"] ?? null,
            percents: $data["percents"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->script = $this->script;
        if (isset($this->missing)) {
            $data->missing = $this->missing;
        }
        if (isset($this->percents)) {
            $data->percents = $this->percents;
        }
        return $data;
    }
}
