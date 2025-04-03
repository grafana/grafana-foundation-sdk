<?php

namespace Grafana\Foundation\Bargauge;

class Options implements \JsonSerializable
{
    public \Grafana\Foundation\Common\BarGaugeDisplayMode $displayMode;

    public \Grafana\Foundation\Common\BarGaugeValueMode $valueMode;

    public \Grafana\Foundation\Common\BarGaugeNamePlacement $namePlacement;

    public bool $showUnfilled;

    public \Grafana\Foundation\Common\BarGaugeSizing $sizing;

    public int $minVizWidth;

    public int $minVizHeight;

    public \Grafana\Foundation\Common\VizLegendOptions $legend;

    public \Grafana\Foundation\Common\ReduceDataOptions $reduceOptions;

    public ?\Grafana\Foundation\Common\VizTextDisplayOptions $text;

    public int $maxVizHeight;

    public \Grafana\Foundation\Common\VizOrientation $orientation;

    /**
     * @param \Grafana\Foundation\Common\BarGaugeDisplayMode|null $displayMode
     * @param \Grafana\Foundation\Common\BarGaugeValueMode|null $valueMode
     * @param \Grafana\Foundation\Common\BarGaugeNamePlacement|null $namePlacement
     * @param bool|null $showUnfilled
     * @param \Grafana\Foundation\Common\BarGaugeSizing|null $sizing
     * @param int|null $minVizWidth
     * @param int|null $minVizHeight
     * @param \Grafana\Foundation\Common\VizLegendOptions|null $legend
     * @param \Grafana\Foundation\Common\ReduceDataOptions|null $reduceOptions
     * @param \Grafana\Foundation\Common\VizTextDisplayOptions|null $text
     * @param int|null $maxVizHeight
     * @param \Grafana\Foundation\Common\VizOrientation|null $orientation
     */
    public function __construct(?\Grafana\Foundation\Common\BarGaugeDisplayMode $displayMode = null, ?\Grafana\Foundation\Common\BarGaugeValueMode $valueMode = null, ?\Grafana\Foundation\Common\BarGaugeNamePlacement $namePlacement = null, ?bool $showUnfilled = null, ?\Grafana\Foundation\Common\BarGaugeSizing $sizing = null, ?int $minVizWidth = null, ?int $minVizHeight = null, ?\Grafana\Foundation\Common\VizLegendOptions $legend = null, ?\Grafana\Foundation\Common\ReduceDataOptions $reduceOptions = null, ?\Grafana\Foundation\Common\VizTextDisplayOptions $text = null, ?int $maxVizHeight = null, ?\Grafana\Foundation\Common\VizOrientation $orientation = null)
    {
        $this->displayMode = $displayMode ?: \Grafana\Foundation\Common\BarGaugeDisplayMode::gradient();
        $this->valueMode = $valueMode ?: \Grafana\Foundation\Common\BarGaugeValueMode::color();
        $this->namePlacement = $namePlacement ?: \Grafana\Foundation\Common\BarGaugeNamePlacement::auto();
        $this->showUnfilled = $showUnfilled ?: true;
        $this->sizing = $sizing ?: \Grafana\Foundation\Common\BarGaugeSizing::auto();
        $this->minVizWidth = $minVizWidth ?: 8;
        $this->minVizHeight = $minVizHeight ?: 16;
        $this->legend = $legend ?: new \Grafana\Foundation\Common\VizLegendOptions();
        $this->reduceOptions = $reduceOptions ?: new \Grafana\Foundation\Common\ReduceDataOptions();
        $this->text = $text;
        $this->maxVizHeight = $maxVizHeight ?: 300;
        $this->orientation = $orientation ?: \Grafana\Foundation\Common\VizOrientation::Auto();
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{displayMode?: string, valueMode?: string, namePlacement?: string, showUnfilled?: bool, sizing?: string, minVizWidth?: int, minVizHeight?: int, legend?: mixed, reduceOptions?: mixed, text?: mixed, maxVizHeight?: int, orientation?: string} $inputData */
        $data = $inputData;
        return new self(
            displayMode: isset($data["displayMode"]) ? (function($input) { return \Grafana\Foundation\Common\BarGaugeDisplayMode::fromValue($input); })($data["displayMode"]) : null,
            valueMode: isset($data["valueMode"]) ? (function($input) { return \Grafana\Foundation\Common\BarGaugeValueMode::fromValue($input); })($data["valueMode"]) : null,
            namePlacement: isset($data["namePlacement"]) ? (function($input) { return \Grafana\Foundation\Common\BarGaugeNamePlacement::fromValue($input); })($data["namePlacement"]) : null,
            showUnfilled: $data["showUnfilled"] ?? null,
            sizing: isset($data["sizing"]) ? (function($input) { return \Grafana\Foundation\Common\BarGaugeSizing::fromValue($input); })($data["sizing"]) : null,
            minVizWidth: $data["minVizWidth"] ?? null,
            minVizHeight: $data["minVizHeight"] ?? null,
            legend: isset($data["legend"]) ? (function($input) {
    	/** @var array{displayMode?: string, placement?: string, showLegend?: bool, asTable?: bool, isVisible?: bool, sortBy?: string, sortDesc?: bool, width?: float, calcs?: array<string>} */
    $val = $input;
    	return \Grafana\Foundation\Common\VizLegendOptions::fromArray($val);
    })($data["legend"]) : null,
            reduceOptions: isset($data["reduceOptions"]) ? (function($input) {
    	/** @var array{values?: bool, limit?: float, calcs?: array<string>, fields?: string} */
    $val = $input;
    	return \Grafana\Foundation\Common\ReduceDataOptions::fromArray($val);
    })($data["reduceOptions"]) : null,
            text: isset($data["text"]) ? (function($input) {
    	/** @var array{titleSize?: float, valueSize?: float, percentSize?: float} */
    $val = $input;
    	return \Grafana\Foundation\Common\VizTextDisplayOptions::fromArray($val);
    })($data["text"]) : null,
            maxVizHeight: $data["maxVizHeight"] ?? null,
            orientation: isset($data["orientation"]) ? (function($input) { return \Grafana\Foundation\Common\VizOrientation::fromValue($input); })($data["orientation"]) : null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "displayMode" => $this->displayMode,
            "valueMode" => $this->valueMode,
            "namePlacement" => $this->namePlacement,
            "showUnfilled" => $this->showUnfilled,
            "sizing" => $this->sizing,
            "minVizWidth" => $this->minVizWidth,
            "minVizHeight" => $this->minVizHeight,
            "legend" => $this->legend,
            "reduceOptions" => $this->reduceOptions,
            "maxVizHeight" => $this->maxVizHeight,
            "orientation" => $this->orientation,
        ];
        if (isset($this->text)) {
            $data["text"] = $this->text;
        }
        return $data;
    }
}
