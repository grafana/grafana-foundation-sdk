<?php

namespace Grafana\Foundation\Heatmap;

class Options implements \JsonSerializable
{
    /**
     * Controls if the heatmap should be calculated from data
     */
    public ?bool $calculate;

    /**
     * Calculation options for the heatmap
     */
    public ?\Grafana\Foundation\Common\HeatmapCalculationOptions $calculation;

    /**
     * Controls the color options
     */
    public \Grafana\Foundation\Heatmap\HeatmapColorOptions $color;

    /**
     * Filters values between a given range
     */
    public ?\Grafana\Foundation\Heatmap\FilterValueRange $filterValues;

    /**
     * Controls tick alignment and value name when not calculating from data
     */
    public ?\Grafana\Foundation\Heatmap\RowsHeatmapOptions $rowsFrame;

    /**
     * | *{
     * 	layout: ui.HeatmapCellLayout & "auto" // TODO: fix after remove when https://github.com/grafana/cuetsy/issues/74 is fixed
     * }
     * Controls the display of the value in the cell
     */
    public \Grafana\Foundation\Common\VisibilityMode $showValue;

    /**
     * Controls gap between cells
     */
    public ?int $cellGap;

    /**
     * Controls cell radius
     */
    public ?float $cellRadius;

    /**
     * Controls cell value unit
     */
    public ?\Grafana\Foundation\Heatmap\CellValues $cellValues;

    /**
     * Controls yAxis placement
     */
    public \Grafana\Foundation\Heatmap\YAxisConfig $yAxis;

    /**
     * | *{
     * 	axisPlacement: ui.AxisPlacement & "left" // TODO: fix after remove when https://github.com/grafana/cuetsy/issues/74 is fixed
     * }
     * Controls legend options
     */
    public \Grafana\Foundation\Heatmap\HeatmapLegend $legend;

    /**
     * Controls tooltip options
     */
    public \Grafana\Foundation\Heatmap\HeatmapTooltip $tooltip;

    /**
     * Controls exemplar options
     */
    public \Grafana\Foundation\Heatmap\ExemplarConfig $exemplars;

    /**
     * @param bool|null $calculate
     * @param \Grafana\Foundation\Common\HeatmapCalculationOptions|null $calculation
     * @param \Grafana\Foundation\Heatmap\HeatmapColorOptions|null $color
     * @param \Grafana\Foundation\Heatmap\FilterValueRange|null $filterValues
     * @param \Grafana\Foundation\Heatmap\RowsHeatmapOptions|null $rowsFrame
     * @param \Grafana\Foundation\Common\VisibilityMode|null $showValue
     * @param int|null $cellGap
     * @param float|null $cellRadius
     * @param \Grafana\Foundation\Heatmap\CellValues|null $cellValues
     * @param \Grafana\Foundation\Heatmap\YAxisConfig|null $yAxis
     * @param \Grafana\Foundation\Heatmap\HeatmapLegend|null $legend
     * @param \Grafana\Foundation\Heatmap\HeatmapTooltip|null $tooltip
     * @param \Grafana\Foundation\Heatmap\ExemplarConfig|null $exemplars
     */
    public function __construct(?bool $calculate = null, ?\Grafana\Foundation\Common\HeatmapCalculationOptions $calculation = null, ?\Grafana\Foundation\Heatmap\HeatmapColorOptions $color = null, ?\Grafana\Foundation\Heatmap\FilterValueRange $filterValues = null, ?\Grafana\Foundation\Heatmap\RowsHeatmapOptions $rowsFrame = null, ?\Grafana\Foundation\Common\VisibilityMode $showValue = null, ?int $cellGap = null, ?float $cellRadius = null, ?\Grafana\Foundation\Heatmap\CellValues $cellValues = null, ?\Grafana\Foundation\Heatmap\YAxisConfig $yAxis = null, ?\Grafana\Foundation\Heatmap\HeatmapLegend $legend = null, ?\Grafana\Foundation\Heatmap\HeatmapTooltip $tooltip = null, ?\Grafana\Foundation\Heatmap\ExemplarConfig $exemplars = null)
    {
        $this->calculate = $calculate;
        $this->calculation = $calculation;
        $this->color = $color ?: new \Grafana\Foundation\Heatmap\HeatmapColorOptions(exponent: 0.5, fill: "dark-orange", reverse: false, scheme: "Oranges", steps: 64);
        $this->filterValues = $filterValues;
        $this->rowsFrame = $rowsFrame;
        $this->showValue = $showValue ?: \Grafana\Foundation\Common\VisibilityMode::auto();
        $this->cellGap = $cellGap;
        $this->cellRadius = $cellRadius;
        $this->cellValues = $cellValues;
        $this->yAxis = $yAxis ?: new \Grafana\Foundation\Heatmap\YAxisConfig();
        $this->legend = $legend ?: new \Grafana\Foundation\Heatmap\HeatmapLegend(show: true);
        $this->tooltip = $tooltip ?: new \Grafana\Foundation\Heatmap\HeatmapTooltip();
        $this->exemplars = $exemplars ?: new \Grafana\Foundation\Heatmap\ExemplarConfig(color: "rgba(255,0,255,0.7)");
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{calculate?: bool, calculation?: mixed, color?: mixed, filterValues?: mixed, rowsFrame?: mixed, showValue?: string, cellGap?: int, cellRadius?: float, cellValues?: mixed, yAxis?: mixed, legend?: mixed, tooltip?: mixed, exemplars?: mixed} $inputData */
        $data = $inputData;
        return new self(
            calculate: $data["calculate"] ?? null,
            calculation: isset($data["calculation"]) ? (function($input) {
    	/** @var array{xBuckets?: mixed, yBuckets?: mixed} */
    $val = $input;
    	return \Grafana\Foundation\Common\HeatmapCalculationOptions::fromArray($val);
    })($data["calculation"]) : null,
            color: isset($data["color"]) ? (function($input) {
    	/** @var array{mode?: string, scheme?: string, fill?: string, scale?: string, exponent?: float, steps?: int, reverse?: bool, min?: float, max?: float} */
    $val = $input;
    	return \Grafana\Foundation\Heatmap\HeatmapColorOptions::fromArray($val);
    })($data["color"]) : null,
            filterValues: isset($data["filterValues"]) ? (function($input) {
    	/** @var array{le?: float, ge?: float} */
    $val = $input;
    	return \Grafana\Foundation\Heatmap\FilterValueRange::fromArray($val);
    })($data["filterValues"]) : null,
            rowsFrame: isset($data["rowsFrame"]) ? (function($input) {
    	/** @var array{value?: string, layout?: string} */
    $val = $input;
    	return \Grafana\Foundation\Heatmap\RowsHeatmapOptions::fromArray($val);
    })($data["rowsFrame"]) : null,
            showValue: isset($data["showValue"]) ? (function($input) { return \Grafana\Foundation\Common\VisibilityMode::fromValue($input); })($data["showValue"]) : null,
            cellGap: $data["cellGap"] ?? null,
            cellRadius: $data["cellRadius"] ?? null,
            cellValues: isset($data["cellValues"]) ? (function($input) {
    	/** @var array{unit?: string, decimals?: float} */
    $val = $input;
    	return \Grafana\Foundation\Heatmap\CellValues::fromArray($val);
    })($data["cellValues"]) : null,
            yAxis: isset($data["yAxis"]) ? (function($input) {
    	/** @var array{unit?: string, reverse?: bool, decimals?: float, min?: float, axisPlacement?: string, axisColorMode?: string, axisLabel?: string, axisWidth?: float, axisSoftMin?: float, axisSoftMax?: float, axisGridShow?: bool, scaleDistribution?: mixed, axisCenteredZero?: bool, max?: float, axisBorderShow?: bool} */
    $val = $input;
    	return \Grafana\Foundation\Heatmap\YAxisConfig::fromArray($val);
    })($data["yAxis"]) : null,
            legend: isset($data["legend"]) ? (function($input) {
    	/** @var array{show?: bool} */
    $val = $input;
    	return \Grafana\Foundation\Heatmap\HeatmapLegend::fromArray($val);
    })($data["legend"]) : null,
            tooltip: isset($data["tooltip"]) ? (function($input) {
    	/** @var array{mode?: string, yHistogram?: bool, showColorScale?: bool} */
    $val = $input;
    	return \Grafana\Foundation\Heatmap\HeatmapTooltip::fromArray($val);
    })($data["tooltip"]) : null,
            exemplars: isset($data["exemplars"]) ? (function($input) {
    	/** @var array{color?: string} */
    $val = $input;
    	return \Grafana\Foundation\Heatmap\ExemplarConfig::fromArray($val);
    })($data["exemplars"]) : null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "color" => $this->color,
            "showValue" => $this->showValue,
            "yAxis" => $this->yAxis,
            "legend" => $this->legend,
            "tooltip" => $this->tooltip,
            "exemplars" => $this->exemplars,
        ];
        if (isset($this->calculate)) {
            $data["calculate"] = $this->calculate;
        }
        if (isset($this->calculation)) {
            $data["calculation"] = $this->calculation;
        }
        if (isset($this->filterValues)) {
            $data["filterValues"] = $this->filterValues;
        }
        if (isset($this->rowsFrame)) {
            $data["rowsFrame"] = $this->rowsFrame;
        }
        if (isset($this->cellGap)) {
            $data["cellGap"] = $this->cellGap;
        }
        if (isset($this->cellRadius)) {
            $data["cellRadius"] = $this->cellRadius;
        }
        if (isset($this->cellValues)) {
            $data["cellValues"] = $this->cellValues;
        }
        return $data;
    }
}
