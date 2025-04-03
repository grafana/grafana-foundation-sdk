<?php

namespace Grafana\Foundation\Candlestick;

class Options implements \JsonSerializable
{
    /**
     * Sets which dimensions are used for the visualization
     */
    public \Grafana\Foundation\Candlestick\VizDisplayMode $mode;

    /**
     * Sets the style of the candlesticks
     */
    public \Grafana\Foundation\Candlestick\CandleStyle $candleStyle;

    /**
     * Sets the color strategy for the candlesticks
     */
    public \Grafana\Foundation\Candlestick\ColorStrategy $colorStrategy;

    /**
     * Map fields to appropriate dimension
     */
    public \Grafana\Foundation\Candlestick\CandlestickFieldMap $fields;

    /**
     * Set which colors are used when the price movement is up or down
     */
    public \Grafana\Foundation\Candlestick\CandlestickColors $colors;

    public \Grafana\Foundation\Common\VizLegendOptions $legend;

    public \Grafana\Foundation\Common\VizTooltipOptions $tooltip;

    /**
     * When enabled, all fields will be sent to the graph
     */
    public ?bool $includeAllFields;

    /**
     * @param \Grafana\Foundation\Candlestick\VizDisplayMode|null $mode
     * @param \Grafana\Foundation\Candlestick\CandleStyle|null $candleStyle
     * @param \Grafana\Foundation\Candlestick\ColorStrategy|null $colorStrategy
     * @param \Grafana\Foundation\Candlestick\CandlestickFieldMap|null $fields
     * @param \Grafana\Foundation\Candlestick\CandlestickColors|null $colors
     * @param \Grafana\Foundation\Common\VizLegendOptions|null $legend
     * @param \Grafana\Foundation\Common\VizTooltipOptions|null $tooltip
     * @param bool|null $includeAllFields
     */
    public function __construct(?\Grafana\Foundation\Candlestick\VizDisplayMode $mode = null, ?\Grafana\Foundation\Candlestick\CandleStyle $candleStyle = null, ?\Grafana\Foundation\Candlestick\ColorStrategy $colorStrategy = null, ?\Grafana\Foundation\Candlestick\CandlestickFieldMap $fields = null, ?\Grafana\Foundation\Candlestick\CandlestickColors $colors = null, ?\Grafana\Foundation\Common\VizLegendOptions $legend = null, ?\Grafana\Foundation\Common\VizTooltipOptions $tooltip = null, ?bool $includeAllFields = null)
    {
        $this->mode = $mode ?: \Grafana\Foundation\Candlestick\VizDisplayMode::candlesVolume();
        $this->candleStyle = $candleStyle ?: \Grafana\Foundation\Candlestick\CandleStyle::candles();
        $this->colorStrategy = $colorStrategy ?: \Grafana\Foundation\Candlestick\ColorStrategy::openClose();
        $this->fields = $fields ?: new \Grafana\Foundation\Candlestick\CandlestickFieldMap();
        $this->colors = $colors ?: new \Grafana\Foundation\Candlestick\CandlestickColors(down: "red", flat: "gray", up: "green");
        $this->legend = $legend ?: new \Grafana\Foundation\Common\VizLegendOptions();
        $this->tooltip = $tooltip ?: new \Grafana\Foundation\Common\VizTooltipOptions();
        $this->includeAllFields = $includeAllFields;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{mode?: string, candleStyle?: string, colorStrategy?: string, fields?: mixed, colors?: mixed, legend?: mixed, tooltip?: mixed, includeAllFields?: bool} $inputData */
        $data = $inputData;
        return new self(
            mode: isset($data["mode"]) ? (function($input) { return \Grafana\Foundation\Candlestick\VizDisplayMode::fromValue($input); })($data["mode"]) : null,
            candleStyle: isset($data["candleStyle"]) ? (function($input) { return \Grafana\Foundation\Candlestick\CandleStyle::fromValue($input); })($data["candleStyle"]) : null,
            colorStrategy: isset($data["colorStrategy"]) ? (function($input) { return \Grafana\Foundation\Candlestick\ColorStrategy::fromValue($input); })($data["colorStrategy"]) : null,
            fields: isset($data["fields"]) ? (function($input) {
    	/** @var array{open?: string, high?: string, low?: string, close?: string, volume?: string} */
    $val = $input;
    	return \Grafana\Foundation\Candlestick\CandlestickFieldMap::fromArray($val);
    })($data["fields"]) : null,
            colors: isset($data["colors"]) ? (function($input) {
    	/** @var array{up?: string, down?: string, flat?: string} */
    $val = $input;
    	return \Grafana\Foundation\Candlestick\CandlestickColors::fromArray($val);
    })($data["colors"]) : null,
            legend: isset($data["legend"]) ? (function($input) {
    	/** @var array{displayMode?: string, placement?: string, showLegend?: bool, asTable?: bool, isVisible?: bool, sortBy?: string, sortDesc?: bool, width?: float, calcs?: array<string>} */
    $val = $input;
    	return \Grafana\Foundation\Common\VizLegendOptions::fromArray($val);
    })($data["legend"]) : null,
            tooltip: isset($data["tooltip"]) ? (function($input) {
    	/** @var array{mode?: string, sort?: string, maxWidth?: float, maxHeight?: float, hideZeros?: bool} */
    $val = $input;
    	return \Grafana\Foundation\Common\VizTooltipOptions::fromArray($val);
    })($data["tooltip"]) : null,
            includeAllFields: $data["includeAllFields"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "mode" => $this->mode,
            "candleStyle" => $this->candleStyle,
            "colorStrategy" => $this->colorStrategy,
            "fields" => $this->fields,
            "colors" => $this->colors,
            "legend" => $this->legend,
            "tooltip" => $this->tooltip,
        ];
        if (isset($this->includeAllFields)) {
            $data["includeAllFields"] = $this->includeAllFields;
        }
        return $data;
    }
}
