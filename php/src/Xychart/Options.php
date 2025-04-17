<?php

namespace Grafana\Foundation\Xychart;

class Options implements \JsonSerializable
{
    public ?\Grafana\Foundation\Xychart\SeriesMapping $seriesMapping;

    /**
     * Table Mode (auto)
     */
    public \Grafana\Foundation\Xychart\XYDimensionConfig $dims;

    public \Grafana\Foundation\Common\VizLegendOptions $legend;

    public \Grafana\Foundation\Common\VizTooltipOptions $tooltip;

    /**
     * Manual Mode
     * @var array<\Grafana\Foundation\Xychart\ScatterSeriesConfig>
     */
    public array $series;

    /**
     * @param \Grafana\Foundation\Xychart\SeriesMapping|null $seriesMapping
     * @param \Grafana\Foundation\Xychart\XYDimensionConfig|null $dims
     * @param \Grafana\Foundation\Common\VizLegendOptions|null $legend
     * @param \Grafana\Foundation\Common\VizTooltipOptions|null $tooltip
     * @param array<\Grafana\Foundation\Xychart\ScatterSeriesConfig>|null $series
     */
    public function __construct(?\Grafana\Foundation\Xychart\SeriesMapping $seriesMapping = null, ?\Grafana\Foundation\Xychart\XYDimensionConfig $dims = null, ?\Grafana\Foundation\Common\VizLegendOptions $legend = null, ?\Grafana\Foundation\Common\VizTooltipOptions $tooltip = null, ?array $series = null)
    {
        $this->seriesMapping = $seriesMapping;
        $this->dims = $dims ?: new \Grafana\Foundation\Xychart\XYDimensionConfig();
        $this->legend = $legend ?: new \Grafana\Foundation\Common\VizLegendOptions();
        $this->tooltip = $tooltip ?: new \Grafana\Foundation\Common\VizTooltipOptions();
        $this->series = $series ?: [];
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{seriesMapping?: string, dims?: mixed, legend?: mixed, tooltip?: mixed, series?: array<mixed>} $inputData */
        $data = $inputData;
        return new self(
            seriesMapping: isset($data["seriesMapping"]) ? (function($input) { return \Grafana\Foundation\Xychart\SeriesMapping::fromValue($input); })($data["seriesMapping"]) : null,
            dims: isset($data["dims"]) ? (function($input) {
    	/** @var array{frame?: int, x?: string, exclude?: array<string>} */
    $val = $input;
    	return \Grafana\Foundation\Xychart\XYDimensionConfig::fromArray($val);
    })($data["dims"]) : null,
            legend: isset($data["legend"]) ? (function($input) {
    	/** @var array{displayMode?: string, placement?: string, showLegend?: bool, asTable?: bool, isVisible?: bool, sortBy?: string, sortDesc?: bool, width?: float, calcs?: array<string>} */
    $val = $input;
    	return \Grafana\Foundation\Common\VizLegendOptions::fromArray($val);
    })($data["legend"]) : null,
            tooltip: isset($data["tooltip"]) ? (function($input) {
    	/** @var array{mode?: string, sort?: string} */
    $val = $input;
    	return \Grafana\Foundation\Common\VizTooltipOptions::fromArray($val);
    })($data["tooltip"]) : null,
            series: array_filter(array_map((function($input) {
    	/** @var array{x?: string, y?: string, show?: string, pointSize?: mixed, pointColor?: mixed, lineColor?: mixed, lineWidth?: int, lineStyle?: mixed, label?: string, hideFrom?: mixed, axisPlacement?: string, axisColorMode?: string, axisLabel?: string, axisWidth?: float, axisSoftMin?: float, axisSoftMax?: float, axisGridShow?: bool, scaleDistribution?: mixed, axisCenteredZero?: bool, name?: string, labelValue?: mixed, axisBorderShow?: bool} */
    $val = $input;
    	return \Grafana\Foundation\Xychart\ScatterSeriesConfig::fromArray($val);
    }), $data["series"] ?? [])),
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->dims = $this->dims;
        $data->legend = $this->legend;
        $data->tooltip = $this->tooltip;
        $data->series = $this->series;
        if (isset($this->seriesMapping)) {
            $data->seriesMapping = $this->seriesMapping;
        }
        return $data;
    }
}
