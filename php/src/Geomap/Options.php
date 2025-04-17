<?php

namespace Grafana\Foundation\Geomap;

class Options implements \JsonSerializable
{
    public \Grafana\Foundation\Geomap\MapViewConfig $view;

    public \Grafana\Foundation\Geomap\ControlsOptions $controls;

    public \Grafana\Foundation\Common\MapLayerOptions $basemap;

    /**
     * @var array<\Grafana\Foundation\Common\MapLayerOptions>
     */
    public array $layers;

    public \Grafana\Foundation\Geomap\TooltipOptions $tooltip;

    /**
     * @param \Grafana\Foundation\Geomap\MapViewConfig|null $view
     * @param \Grafana\Foundation\Geomap\ControlsOptions|null $controls
     * @param \Grafana\Foundation\Common\MapLayerOptions|null $basemap
     * @param array<\Grafana\Foundation\Common\MapLayerOptions>|null $layers
     * @param \Grafana\Foundation\Geomap\TooltipOptions|null $tooltip
     */
    public function __construct(?\Grafana\Foundation\Geomap\MapViewConfig $view = null, ?\Grafana\Foundation\Geomap\ControlsOptions $controls = null, ?\Grafana\Foundation\Common\MapLayerOptions $basemap = null, ?array $layers = null, ?\Grafana\Foundation\Geomap\TooltipOptions $tooltip = null)
    {
        $this->view = $view ?: new \Grafana\Foundation\Geomap\MapViewConfig();
        $this->controls = $controls ?: new \Grafana\Foundation\Geomap\ControlsOptions();
        $this->basemap = $basemap ?: new \Grafana\Foundation\Common\MapLayerOptions();
        $this->layers = $layers ?: [];
        $this->tooltip = $tooltip ?: new \Grafana\Foundation\Geomap\TooltipOptions();
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{view?: mixed, controls?: mixed, basemap?: mixed, layers?: array<mixed>, tooltip?: mixed} $inputData */
        $data = $inputData;
        return new self(
            view: isset($data["view"]) ? (function($input) {
    	/** @var array{id?: string, lat?: int, lon?: int, zoom?: int, minZoom?: int, maxZoom?: int, padding?: int, allLayers?: bool, lastOnly?: bool, layer?: string, shared?: bool} */
    $val = $input;
    	return \Grafana\Foundation\Geomap\MapViewConfig::fromArray($val);
    })($data["view"]) : null,
            controls: isset($data["controls"]) ? (function($input) {
    	/** @var array{showZoom?: bool, mouseWheelZoom?: bool, showAttribution?: bool, showScale?: bool, showDebug?: bool, showMeasure?: bool} */
    $val = $input;
    	return \Grafana\Foundation\Geomap\ControlsOptions::fromArray($val);
    })($data["controls"]) : null,
            basemap: isset($data["basemap"]) ? (function($input) {
    	/** @var array{type?: string, name?: string, config?: mixed, location?: mixed, filterData?: mixed, opacity?: int, tooltip?: bool} */
    $val = $input;
    	return \Grafana\Foundation\Common\MapLayerOptions::fromArray($val);
    })($data["basemap"]) : null,
            layers: array_filter(array_map((function($input) {
    	/** @var array{type?: string, name?: string, config?: mixed, location?: mixed, filterData?: mixed, opacity?: int, tooltip?: bool} */
    $val = $input;
    	return \Grafana\Foundation\Common\MapLayerOptions::fromArray($val);
    }), $data["layers"] ?? [])),
            tooltip: isset($data["tooltip"]) ? (function($input) {
    	/** @var array{mode?: string} */
    $val = $input;
    	return \Grafana\Foundation\Geomap\TooltipOptions::fromArray($val);
    })($data["tooltip"]) : null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->view = $this->view;
        $data->controls = $this->controls;
        $data->basemap = $this->basemap;
        $data->layers = $this->layers;
        $data->tooltip = $this->tooltip;
        return $data;
    }
}
