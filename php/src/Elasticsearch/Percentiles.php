<?php

namespace Grafana\Foundation\Elasticsearch;

class Percentiles implements \JsonSerializable
{
    public string $type;

    public ?string $field;

    public string $id;

    public ?\Grafana\Foundation\Elasticsearch\ElasticsearchPercentilesSettings $settings;

    public ?bool $hide;

    /**
     * @param string|null $field
     * @param string|null $id
     * @param \Grafana\Foundation\Elasticsearch\ElasticsearchPercentilesSettings|null $settings
     * @param bool|null $hide
     */
    public function __construct(?string $field = null, ?string $id = null, ?\Grafana\Foundation\Elasticsearch\ElasticsearchPercentilesSettings $settings = null, ?bool $hide = null)
    {
        $this->type = "percentiles";
    
        $this->field = $field;
        $this->id = $id ?: "";
        $this->settings = $settings;
        $this->hide = $hide;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{type?: string, field?: string, id?: string, settings?: mixed, hide?: bool} $inputData */
        $data = $inputData;
        return new self(
            field: $data["field"] ?? null,
            id: $data["id"] ?? null,
            settings: isset($data["settings"]) ? (function($input) {
    	/** @var array{script?: string|mixed, missing?: string, percents?: array<string>} */
    $val = $input;
    	return \Grafana\Foundation\Elasticsearch\ElasticsearchPercentilesSettings::fromArray($val);
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
        if (isset($this->field)) {
            $data["field"] = $this->field;
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
