<?php

namespace Grafana\Foundation\Elasticsearch;

class ElasticsearchMovingFunctionSettings implements \JsonSerializable
{
    public ?string $window;

    /**
     * @var string|\Grafana\Foundation\Elasticsearch\ElasticsearchInlineScript
     */
    public $script;

    public ?string $shift;

    /**
     * @param string|null $window
     * @param string|\Grafana\Foundation\Elasticsearch\ElasticsearchInlineScript|null $script
     * @param string|null $shift
     */
    public function __construct(?string $window = null,  $script = null, ?string $shift = null)
    {
        $this->window = $window;
        $this->script = $script ?: "";
        $this->shift = $shift;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{window?: string, script?: string|mixed, shift?: string} $inputData */
        $data = $inputData;
        return new self(
            window: $data["window"] ?? null,
            script: isset($data["script"]) ? (function($input) {
        switch (true) {
        case is_string($input):
            return $input;
        default:
            /** @var array{inline?: string} $input */
            return \Grafana\Foundation\Elasticsearch\ElasticsearchInlineScript::fromArray($input);
    }
    })($data["script"]) : null,
            shift: $data["shift"] ?? null,
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
        if (isset($this->window)) {
            $data["window"] = $this->window;
        }
        if (isset($this->shift)) {
            $data["shift"] = $this->shift;
        }
        return $data;
    }
}
