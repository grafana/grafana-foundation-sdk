<?php

namespace Grafana\Foundation\Geomap;

class ControlsOptions implements \JsonSerializable
{
    /**
     * Zoom (upper left)
     */
    public ?bool $showZoom;

    /**
     * let the mouse wheel zoom
     */
    public ?bool $mouseWheelZoom;

    /**
     * Lower right
     */
    public ?bool $showAttribution;

    /**
     * Scale options
     */
    public ?bool $showScale;

    /**
     * Show debug
     */
    public ?bool $showDebug;

    /**
     * Show measure
     */
    public ?bool $showMeasure;

    /**
     * @param bool|null $showZoom
     * @param bool|null $mouseWheelZoom
     * @param bool|null $showAttribution
     * @param bool|null $showScale
     * @param bool|null $showDebug
     * @param bool|null $showMeasure
     */
    public function __construct(?bool $showZoom = null, ?bool $mouseWheelZoom = null, ?bool $showAttribution = null, ?bool $showScale = null, ?bool $showDebug = null, ?bool $showMeasure = null)
    {
        $this->showZoom = $showZoom;
        $this->mouseWheelZoom = $mouseWheelZoom;
        $this->showAttribution = $showAttribution;
        $this->showScale = $showScale;
        $this->showDebug = $showDebug;
        $this->showMeasure = $showMeasure;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{showZoom?: bool, mouseWheelZoom?: bool, showAttribution?: bool, showScale?: bool, showDebug?: bool, showMeasure?: bool} $inputData */
        $data = $inputData;
        return new self(
            showZoom: $data["showZoom"] ?? null,
            mouseWheelZoom: $data["mouseWheelZoom"] ?? null,
            showAttribution: $data["showAttribution"] ?? null,
            showScale: $data["showScale"] ?? null,
            showDebug: $data["showDebug"] ?? null,
            showMeasure: $data["showMeasure"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        if (isset($this->showZoom)) {
            $data->showZoom = $this->showZoom;
        }
        if (isset($this->mouseWheelZoom)) {
            $data->mouseWheelZoom = $this->mouseWheelZoom;
        }
        if (isset($this->showAttribution)) {
            $data->showAttribution = $this->showAttribution;
        }
        if (isset($this->showScale)) {
            $data->showScale = $this->showScale;
        }
        if (isset($this->showDebug)) {
            $data->showDebug = $this->showDebug;
        }
        if (isset($this->showMeasure)) {
            $data->showMeasure = $this->showMeasure;
        }
        return $data;
    }
}
