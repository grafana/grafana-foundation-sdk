<?php

namespace Grafana\Foundation\Gauge;

class Options implements \JsonSerializable
{
    public bool $showThresholdLabels;

    public bool $showThresholdMarkers;

    public \Grafana\Foundation\Common\BarGaugeSizing $sizing;

    public int $minVizWidth;

    public \Grafana\Foundation\Common\ReduceDataOptions $reduceOptions;

    public ?\Grafana\Foundation\Common\VizTextDisplayOptions $text;

    public int $minVizHeight;

    public \Grafana\Foundation\Common\VizOrientation $orientation;

    /**
     * @param bool|null $showThresholdLabels
     * @param bool|null $showThresholdMarkers
     * @param \Grafana\Foundation\Common\BarGaugeSizing|null $sizing
     * @param int|null $minVizWidth
     * @param \Grafana\Foundation\Common\ReduceDataOptions|null $reduceOptions
     * @param \Grafana\Foundation\Common\VizTextDisplayOptions|null $text
     * @param int|null $minVizHeight
     * @param \Grafana\Foundation\Common\VizOrientation|null $orientation
     */
    public function __construct(?bool $showThresholdLabels = null, ?bool $showThresholdMarkers = null, ?\Grafana\Foundation\Common\BarGaugeSizing $sizing = null, ?int $minVizWidth = null, ?\Grafana\Foundation\Common\ReduceDataOptions $reduceOptions = null, ?\Grafana\Foundation\Common\VizTextDisplayOptions $text = null, ?int $minVizHeight = null, ?\Grafana\Foundation\Common\VizOrientation $orientation = null)
    {
        $this->showThresholdLabels = $showThresholdLabels ?: false;
        $this->showThresholdMarkers = $showThresholdMarkers ?: true;
        $this->sizing = $sizing ?: \Grafana\Foundation\Common\BarGaugeSizing::auto();
        $this->minVizWidth = $minVizWidth ?: 75;
        $this->reduceOptions = $reduceOptions ?: new \Grafana\Foundation\Common\ReduceDataOptions();
        $this->text = $text;
        $this->minVizHeight = $minVizHeight ?: 75;
        $this->orientation = $orientation ?: \Grafana\Foundation\Common\VizOrientation::Auto();
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{showThresholdLabels?: bool, showThresholdMarkers?: bool, sizing?: string, minVizWidth?: int, reduceOptions?: mixed, text?: mixed, minVizHeight?: int, orientation?: string} $inputData */
        $data = $inputData;
        return new self(
            showThresholdLabels: $data["showThresholdLabels"] ?? null,
            showThresholdMarkers: $data["showThresholdMarkers"] ?? null,
            sizing: isset($data["sizing"]) ? (function($input) { return \Grafana\Foundation\Common\BarGaugeSizing::fromValue($input); })($data["sizing"]) : null,
            minVizWidth: $data["minVizWidth"] ?? null,
            reduceOptions: isset($data["reduceOptions"]) ? (function($input) {
    	/** @var array{values?: bool, limit?: float, calcs?: array<string>, fields?: string} */
    $val = $input;
    	return \Grafana\Foundation\Common\ReduceDataOptions::fromArray($val);
    })($data["reduceOptions"]) : null,
            text: isset($data["text"]) ? (function($input) {
    	/** @var array{titleSize?: float, valueSize?: float} */
    $val = $input;
    	return \Grafana\Foundation\Common\VizTextDisplayOptions::fromArray($val);
    })($data["text"]) : null,
            minVizHeight: $data["minVizHeight"] ?? null,
            orientation: isset($data["orientation"]) ? (function($input) { return \Grafana\Foundation\Common\VizOrientation::fromValue($input); })($data["orientation"]) : null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "showThresholdLabels" => $this->showThresholdLabels,
            "showThresholdMarkers" => $this->showThresholdMarkers,
            "sizing" => $this->sizing,
            "minVizWidth" => $this->minVizWidth,
            "reduceOptions" => $this->reduceOptions,
            "minVizHeight" => $this->minVizHeight,
            "orientation" => $this->orientation,
        ];
        if (isset($this->text)) {
            $data["text"] = $this->text;
        }
        return $data;
    }
}
