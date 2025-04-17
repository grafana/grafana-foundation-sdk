<?php

namespace Grafana\Foundation\Elasticsearch;

class GeoHashGrid implements \JsonSerializable
{
    public ?string $field;

    public string $id;

    public \Grafana\Foundation\Elasticsearch\BucketAggregationType $type;

    public ?\Grafana\Foundation\Elasticsearch\ElasticsearchGeoHashGridSettings $settings;

    /**
     * @param string|null $field
     * @param string|null $id
     * @param \Grafana\Foundation\Elasticsearch\ElasticsearchGeoHashGridSettings|null $settings
     */
    public function __construct(?string $field = null, ?string $id = null, ?\Grafana\Foundation\Elasticsearch\ElasticsearchGeoHashGridSettings $settings = null)
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
        /** @var array{field?: string, id?: string, type?: "geohash_grid", settings?: mixed} $inputData */
        $data = $inputData;
        return new self(
            field: $data["field"] ?? null,
            id: $data["id"] ?? null,
            settings: isset($data["settings"]) ? (function($input) {
    	/** @var array{precision?: string} */
    $val = $input;
    	return \Grafana\Foundation\Elasticsearch\ElasticsearchGeoHashGridSettings::fromArray($val);
    })($data["settings"]) : null,
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
