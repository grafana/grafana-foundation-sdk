<?php

namespace Grafana\Foundation\Elasticsearch;

class BaseBucketAggregation implements \JsonSerializable
{
    public string $id;

    public \Grafana\Foundation\Elasticsearch\BucketAggregationType $type;

    /**
     * @var mixed|null
     */
    public $settings;

    /**
     * @param string|null $id
     * @param \Grafana\Foundation\Elasticsearch\BucketAggregationType|null $type
     * @param mixed|null $settings
     */
    public function __construct(?string $id = null, ?\Grafana\Foundation\Elasticsearch\BucketAggregationType $type = null,  $settings = null)
    {
        $this->id = $id ?: "";
        $this->type = $type ?: \Grafana\Foundation\Elasticsearch\BucketAggregationType::Terms();
        $this->settings = $settings;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{id?: string, type?: string, settings?: mixed} $inputData */
        $data = $inputData;
        return new self(
            id: $data["id"] ?? null,
            type: isset($data["type"]) ? (function($input) { return \Grafana\Foundation\Elasticsearch\BucketAggregationType::fromValue($input); })($data["type"]) : null,
            settings: $data["settings"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "id" => $this->id,
            "type" => $this->type,
        ];
        if (isset($this->settings)) {
            $data["settings"] = $this->settings;
        }
        return $data;
    }
}
