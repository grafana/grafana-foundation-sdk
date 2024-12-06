<?php

namespace Grafana\Foundation\Piechart;

class Options implements \JsonSerializable
{
    public \Grafana\Foundation\Piechart\PieChartType $pieType;

    /**
     * @var array<\Grafana\Foundation\Piechart\PieChartLabels>|null
     */
    public ?array $displayLabels;

    public \Grafana\Foundation\Common\VizTooltipOptions $tooltip;

    public \Grafana\Foundation\Common\ReduceDataOptions $reduceOptions;

    public ?\Grafana\Foundation\Common\VizTextDisplayOptions $text;

    public \Grafana\Foundation\Piechart\PieChartLegendOptions $legend;

    public \Grafana\Foundation\Common\VizOrientation $orientation;

    /**
     * @param \Grafana\Foundation\Piechart\PieChartType|null $pieType
     * @param array<\Grafana\Foundation\Piechart\PieChartLabels>|null $displayLabels
     * @param \Grafana\Foundation\Common\VizTooltipOptions|null $tooltip
     * @param \Grafana\Foundation\Common\ReduceDataOptions|null $reduceOptions
     * @param \Grafana\Foundation\Common\VizTextDisplayOptions|null $text
     * @param \Grafana\Foundation\Piechart\PieChartLegendOptions|null $legend
     * @param \Grafana\Foundation\Common\VizOrientation|null $orientation
     */
    public function __construct(?\Grafana\Foundation\Piechart\PieChartType $pieType = null, ?array $displayLabels = null, ?\Grafana\Foundation\Common\VizTooltipOptions $tooltip = null, ?\Grafana\Foundation\Common\ReduceDataOptions $reduceOptions = null, ?\Grafana\Foundation\Common\VizTextDisplayOptions $text = null, ?\Grafana\Foundation\Piechart\PieChartLegendOptions $legend = null, ?\Grafana\Foundation\Common\VizOrientation $orientation = null)
    {
        $this->pieType = $pieType ?: \Grafana\Foundation\Piechart\PieChartType::Pie();
        $this->displayLabels = $displayLabels;
        $this->tooltip = $tooltip ?: new \Grafana\Foundation\Common\VizTooltipOptions();
        $this->reduceOptions = $reduceOptions ?: new \Grafana\Foundation\Common\ReduceDataOptions();
        $this->text = $text;
        $this->legend = $legend ?: new \Grafana\Foundation\Piechart\PieChartLegendOptions();
        $this->orientation = $orientation ?: \Grafana\Foundation\Common\VizOrientation::Auto();
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{pieType?: string, displayLabels?: array<string>, tooltip?: mixed, reduceOptions?: mixed, text?: mixed, legend?: mixed, orientation?: string} $inputData */
        $data = $inputData;
        return new self(
            pieType: isset($data["pieType"]) ? (function($input) { return \Grafana\Foundation\Piechart\PieChartType::fromValue($input); })($data["pieType"]) : null,
            displayLabels: array_filter(array_map((function($input) { return \Grafana\Foundation\Piechart\PieChartLabels::fromValue($input); }), $data["displayLabels"] ?? [])),
            tooltip: isset($data["tooltip"]) ? (function($input) {
    	/** @var array{mode?: string, sort?: string, maxWidth?: float, maxHeight?: float} */
    $val = $input;
    	return \Grafana\Foundation\Common\VizTooltipOptions::fromArray($val);
    })($data["tooltip"]) : null,
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
            legend: isset($data["legend"]) ? (function($input) {
    	/** @var array{values?: array<string>, displayMode?: string, placement?: string, showLegend?: bool, asTable?: bool, isVisible?: bool, sortBy?: string, sortDesc?: bool, width?: float, calcs?: array<string>} */
    $val = $input;
    	return \Grafana\Foundation\Piechart\PieChartLegendOptions::fromArray($val);
    })($data["legend"]) : null,
            orientation: isset($data["orientation"]) ? (function($input) { return \Grafana\Foundation\Common\VizOrientation::fromValue($input); })($data["orientation"]) : null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "pieType" => $this->pieType,
            "tooltip" => $this->tooltip,
            "reduceOptions" => $this->reduceOptions,
            "legend" => $this->legend,
            "orientation" => $this->orientation,
        ];
        if (isset($this->displayLabels)) {
            $data["displayLabels"] = $this->displayLabels;
        }
        if (isset($this->text)) {
            $data["text"] = $this->text;
        }
        return $data;
    }
}
