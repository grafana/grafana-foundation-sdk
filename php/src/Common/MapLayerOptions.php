<?php

namespace Grafana\Foundation\Common;

class MapLayerOptions implements \JsonSerializable
{
    public string $type;

    /**
     * configured unique display name
     */
    public string $name;

    /**
     * Custom options depending on the type
     * @var mixed|null
     */
    public $config;

    /**
     * Common method to define geometry fields
     */
    public ?\Grafana\Foundation\Common\FrameGeometrySource $location;

    /**
     * Defines a frame MatcherConfig that may filter data for the given layer
     * @var mixed|null
     */
    public $filterData;

    /**
     * Common properties:
     * https://openlayers.org/en/latest/apidoc/module-ol_layer_Base-BaseLayer.html
     * Layer opacity (0-1)
     */
    public ?int $opacity;

    /**
     * Check tooltip (defaults to true)
     */
    public ?bool $tooltip;

    /**
     * @param string|null $type
     * @param string|null $name
     * @param mixed|null $config
     * @param \Grafana\Foundation\Common\FrameGeometrySource|null $location
     * @param mixed|null $filterData
     * @param int|null $opacity
     * @param bool|null $tooltip
     */
    public function __construct(?string $type = null, ?string $name = null,  $config = null, ?\Grafana\Foundation\Common\FrameGeometrySource $location = null,  $filterData = null, ?int $opacity = null, ?bool $tooltip = null)
    {
        $this->type = $type ?: "";
        $this->name = $name ?: "";
        $this->config = $config;
        $this->location = $location;
        $this->filterData = $filterData;
        $this->opacity = $opacity;
        $this->tooltip = $tooltip;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{type?: string, name?: string, config?: mixed, location?: mixed, filterData?: mixed, opacity?: int, tooltip?: bool} $inputData */
        $data = $inputData;
        return new self(
            type: $data["type"] ?? null,
            name: $data["name"] ?? null,
            config: $data["config"] ?? null,
            location: isset($data["location"]) ? (function($input) {
    	/** @var array{mode?: string, geohash?: string, latitude?: string, longitude?: string, wkt?: string, lookup?: string, gazetteer?: string} */
    $val = $input;
    	return \Grafana\Foundation\Common\FrameGeometrySource::fromArray($val);
    })($data["location"]) : null,
            filterData: $data["filterData"] ?? null,
            opacity: $data["opacity"] ?? null,
            tooltip: $data["tooltip"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "type" => $this->type,
            "name" => $this->name,
        ];
        if (isset($this->config)) {
            $data["config"] = $this->config;
        }
        if (isset($this->location)) {
            $data["location"] = $this->location;
        }
        if (isset($this->filterData)) {
            $data["filterData"] = $this->filterData;
        }
        if (isset($this->opacity)) {
            $data["opacity"] = $this->opacity;
        }
        if (isset($this->tooltip)) {
            $data["tooltip"] = $this->tooltip;
        }
        return $data;
    }
}
