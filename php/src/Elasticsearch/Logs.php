<?php

namespace Grafana\Foundation\Elasticsearch;

class Logs implements \JsonSerializable
{
    public string $type;

    public string $id;

    public ?\Grafana\Foundation\Elasticsearch\ElasticsearchLogsSettings $settings;

    public ?bool $hide;

    /**
     * @param string|null $id
     * @param \Grafana\Foundation\Elasticsearch\ElasticsearchLogsSettings|null $settings
     * @param bool|null $hide
     */
    public function __construct(?string $id = null, ?\Grafana\Foundation\Elasticsearch\ElasticsearchLogsSettings $settings = null, ?bool $hide = null)
    {
        $this->type = "logs";
    
        $this->id = $id ?: "";
        $this->settings = $settings;
        $this->hide = $hide;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{type?: string, id?: string, settings?: mixed, hide?: bool} $inputData */
        $data = $inputData;
        return new self(
            id: $data["id"] ?? null,
            settings: isset($data["settings"]) ? (function($input) {
    	/** @var array{limit?: string} */
    $val = $input;
    	return \Grafana\Foundation\Elasticsearch\ElasticsearchLogsSettings::fromArray($val);
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
        if (isset($this->settings)) {
            $data["settings"] = $this->settings;
        }
        if (isset($this->hide)) {
            $data["hide"] = $this->hide;
        }
        return $data;
    }
}
