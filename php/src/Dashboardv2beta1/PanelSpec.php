<?php

namespace Grafana\Foundation\Dashboardv2beta1;

class PanelSpec implements \JsonSerializable
{
    public float $id;

    public string $title;

    public string $description;

    /**
     * @var array<\Grafana\Foundation\Dashboardv2beta1\DataLink>
     */
    public array $links;

    public \Grafana\Foundation\Dashboardv2beta1\QueryGroupKind $data;

    public \Grafana\Foundation\Dashboardv2beta1\VizConfigKind $vizConfig;

    public ?bool $transparent;

    /**
     * @param float|null $id
     * @param string|null $title
     * @param string|null $description
     * @param array<\Grafana\Foundation\Dashboardv2beta1\DataLink>|null $links
     * @param \Grafana\Foundation\Dashboardv2beta1\QueryGroupKind|null $data
     * @param \Grafana\Foundation\Dashboardv2beta1\VizConfigKind|null $vizConfig
     * @param bool|null $transparent
     */
    public function __construct(?float $id = null, ?string $title = null, ?string $description = null, ?array $links = null, ?\Grafana\Foundation\Dashboardv2beta1\QueryGroupKind $data = null, ?\Grafana\Foundation\Dashboardv2beta1\VizConfigKind $vizConfig = null, ?bool $transparent = null)
    {
        $this->id = $id ?: 0;
        $this->title = $title ?: "";
        $this->description = $description ?: "";
        $this->links = $links ?: [];
        $this->data = $data ?: new \Grafana\Foundation\Dashboardv2beta1\QueryGroupKind();
        $this->vizConfig = $vizConfig ?: new \Grafana\Foundation\Dashboardv2beta1\VizConfigKind();
        $this->transparent = $transparent;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{id?: float, title?: string, description?: string, links?: array<mixed>, data?: mixed, vizConfig?: mixed, transparent?: bool} $inputData */
        $data = $inputData;
        return new self(
            id: $data["id"] ?? null,
            title: $data["title"] ?? null,
            description: $data["description"] ?? null,
            links: array_filter(array_map((function($input) {
    	/** @var array{title?: string, url?: string, targetBlank?: bool} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\DataLink::fromArray($val);
    }), $data["links"] ?? [])),
            data: isset($data["data"]) ? (function($input) {
    	/** @var array{kind?: string, spec?: mixed} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\QueryGroupKind::fromArray($val);
    })($data["data"]) : null,
            vizConfig: isset($data["vizConfig"]) ? (function($input) {
    	/** @var array{kind?: string, group?: string, version?: string, spec?: mixed} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\VizConfigKind::fromArray($val);
    })($data["vizConfig"]) : null,
            transparent: $data["transparent"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->id = $this->id;
        $data->title = $this->title;
        $data->description = $this->description;
        $data->links = $this->links;
        $data->data = $this->data;
        $data->vizConfig = $this->vizConfig;
        if (isset($this->transparent)) {
            $data->transparent = $this->transparent;
        }
        return $data;
    }
}
