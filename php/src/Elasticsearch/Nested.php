<?php

namespace Grafana\Foundation\Elasticsearch;

class Nested implements \JsonSerializable
{
    public ?string $field;

    public string $id;

    public \Grafana\Foundation\Elasticsearch\BucketAggregationType $type;

    /**
     * @var mixed|null
     */
    public $settings;

    /**
     * @param string|null $field
     * @param string|null $id
     * @param mixed|null $settings
     */
    public function __construct(?string $field = null, ?string $id = null,  $settings = null)
    {
        $this->field = $field;
        $this->id = $id ?: "";
        $this->type = \Grafana\Foundation\Elasticsearch\BucketAggregationType::terms();
        $this->settings = $settings;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{field?: string, id?: string, type?: "nested", settings?: mixed} $inputData */
        $data = $inputData;
        return new self(
            field: $data["field"] ?? null,
            id: $data["id"] ?? null,
            settings: $data["settings"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->id = $this->id;
        $data->type = $this->type;
        if (isset($this->field)) {
            $data->field = $this->field;
        }
        if (isset($this->settings)) {
            $data->settings = $this->settings;
        }
        return $data;
    }
}
