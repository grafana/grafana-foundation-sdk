<?php

namespace Grafana\Foundation\Elasticsearch;

class ElasticsearchExtendedStatsSettings implements \JsonSerializable
{
    /**
     * @var string|\Grafana\Foundation\Elasticsearch\ElasticsearchInlineScript
     */
    public $script;

    public ?string $missing;

    public ?string $sigma;

    /**
     * @param string|\Grafana\Foundation\Elasticsearch\ElasticsearchInlineScript|null $script
     * @param string|null $missing
     * @param string|null $sigma
     */
    public function __construct( $script = null, ?string $missing = null, ?string $sigma = null)
    {
        $this->script = $script ?: "";
        $this->missing = $missing;
        $this->sigma = $sigma;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{script?: string|mixed, missing?: string, sigma?: string} $inputData */
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
            sigma: $data["sigma"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "script" => $this->script,
        ];
        if (isset($this->missing)) {
            $data["missing"] = $this->missing;
        }
        if (isset($this->sigma)) {
            $data["sigma"] = $this->sigma;
        }
        return $data;
    }
}
