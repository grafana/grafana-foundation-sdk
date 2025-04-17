<?php

namespace Grafana\Foundation\Xychart;

class Options implements \JsonSerializable
{
    public \Grafana\Foundation\Xychart\SeriesMapping $mapping;

    public \Grafana\Foundation\Common\VizLegendOptions $legend;

    public \Grafana\Foundation\Common\VizTooltipOptions $tooltip;

    /**
     * @var array<\Grafana\Foundation\Xychart\XYSeriesConfig>
     */
    public array $series;

    /**
     * @param \Grafana\Foundation\Xychart\SeriesMapping|null $mapping
     * @param \Grafana\Foundation\Common\VizLegendOptions|null $legend
     * @param \Grafana\Foundation\Common\VizTooltipOptions|null $tooltip
     * @param array<\Grafana\Foundation\Xychart\XYSeriesConfig>|null $series
     */
    public function __construct(?\Grafana\Foundation\Xychart\SeriesMapping $mapping = null, ?\Grafana\Foundation\Common\VizLegendOptions $legend = null, ?\Grafana\Foundation\Common\VizTooltipOptions $tooltip = null, ?array $series = null)
    {
        $this->mapping = $mapping ?: \Grafana\Foundation\Xychart\SeriesMapping::Auto();
        $this->legend = $legend ?: new \Grafana\Foundation\Common\VizLegendOptions();
        $this->tooltip = $tooltip ?: new \Grafana\Foundation\Common\VizTooltipOptions();
        $this->series = $series ?: [];
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{mapping?: string, legend?: mixed, tooltip?: mixed, series?: array<mixed>} $inputData */
        $data = $inputData;
        return new self(
            mapping: isset($data["mapping"]) ? (function($input) { return \Grafana\Foundation\Xychart\SeriesMapping::fromValue($input); })($data["mapping"]) : null,
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
            series: array_filter(array_map((function($input) {
    	/** @var array{name?: mixed, frame?: mixed, x?: mixed, y?: mixed, color?: mixed, size?: mixed} */
    $val = $input;
    	return \Grafana\Foundation\Xychart\XYSeriesConfig::fromArray($val);
    }), $data["series"] ?? [])),
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->mapping = $this->mapping;
        $data->legend = $this->legend;
        $data->tooltip = $this->tooltip;
        $data->series = $this->series;
        return $data;
    }
}
