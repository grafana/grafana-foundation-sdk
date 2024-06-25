<?php

namespace Grafana\Foundation\Bargauge;

class Options implements \JsonSerializable
{
    public \Grafana\Foundation\Common\BarGaugeDisplayMode $displayMode;

    public \Grafana\Foundation\Common\BarGaugeValueMode $valueMode;

    public bool $showUnfilled;

    public int $minVizWidth;

    public \Grafana\Foundation\Common\ReduceDataOptions $reduceOptions;

    public ?\Grafana\Foundation\Common\VizTextDisplayOptions $text;

    public int $minVizHeight;

    public \Grafana\Foundation\Common\VizOrientation $orientation;

    /**
     * @param \Grafana\Foundation\Common\BarGaugeDisplayMode|null $displayMode
     * @param \Grafana\Foundation\Common\BarGaugeValueMode|null $valueMode
     * @param bool|null $showUnfilled
     * @param int|null $minVizWidth
     * @param \Grafana\Foundation\Common\ReduceDataOptions|null $reduceOptions
     * @param \Grafana\Foundation\Common\VizTextDisplayOptions|null $text
     * @param int|null $minVizHeight
     * @param \Grafana\Foundation\Common\VizOrientation|null $orientation
     */
    public function __construct(?\Grafana\Foundation\Common\BarGaugeDisplayMode $displayMode = null, ?\Grafana\Foundation\Common\BarGaugeValueMode $valueMode = null, ?bool $showUnfilled = null, ?int $minVizWidth = null, ?\Grafana\Foundation\Common\ReduceDataOptions $reduceOptions = null, ?\Grafana\Foundation\Common\VizTextDisplayOptions $text = null, ?int $minVizHeight = null, ?\Grafana\Foundation\Common\VizOrientation $orientation = null)
    {
        $this->displayMode = $displayMode ?: \Grafana\Foundation\Common\BarGaugeDisplayMode::gradient();
        $this->valueMode = $valueMode ?: \Grafana\Foundation\Common\BarGaugeValueMode::color();
        $this->showUnfilled = $showUnfilled ?: true;
        $this->minVizWidth = $minVizWidth ?: 0;
        $this->reduceOptions = $reduceOptions ?: new \Grafana\Foundation\Common\ReduceDataOptions();
        $this->text = $text;
        $this->minVizHeight = $minVizHeight ?: 10;
        $this->orientation = $orientation ?: \Grafana\Foundation\Common\VizOrientation::Auto();
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{displayMode?: string, valueMode?: string, showUnfilled?: bool, minVizWidth?: int, reduceOptions?: mixed, text?: mixed, minVizHeight?: int, orientation?: string} $inputData */
        $data = $inputData;
        return new self(
            displayMode: isset($data["displayMode"]) ? (function($input) { return \Grafana\Foundation\Common\BarGaugeDisplayMode::fromValue($input); })($data["displayMode"]) : null,
            valueMode: isset($data["valueMode"]) ? (function($input) { return \Grafana\Foundation\Common\BarGaugeValueMode::fromValue($input); })($data["valueMode"]) : null,
            showUnfilled: $data["showUnfilled"] ?? null,
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
            "displayMode" => $this->displayMode,
            "valueMode" => $this->valueMode,
            "showUnfilled" => $this->showUnfilled,
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
