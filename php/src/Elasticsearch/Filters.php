<?php

namespace Grafana\Foundation\Elasticsearch;

class Filters implements \JsonSerializable
{
    public string $id;

    public string $type;

    public ?\Grafana\Foundation\Elasticsearch\ElasticsearchFiltersSettings $settings;

    /**
     * @param string|null $id
     * @param \Grafana\Foundation\Elasticsearch\ElasticsearchFiltersSettings|null $settings
     */
    public function __construct(?string $id = null, ?\Grafana\Foundation\Elasticsearch\ElasticsearchFiltersSettings $settings = null)
    {
        $this->id = $id ?: "";
        $this->type = "filters";
    
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
            settings: isset($data["settings"]) ? (function($input) {
    	/** @var array{filters?: array<mixed>} */
    $val = $input;
    	return \Grafana\Foundation\Elasticsearch\ElasticsearchFiltersSettings::fromArray($val);
    })($data["settings"]) : null,
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
