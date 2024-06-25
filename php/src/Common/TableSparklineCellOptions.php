<?php

namespace Grafana\Foundation\Common;

/**
 * Sparkline cell options
 */
class TableSparklineCellOptions implements \JsonSerializable
{
    public string $type;

    public ?\Grafana\Foundation\Common\GraphDrawStyle $drawStyle;

    public ?\Grafana\Foundation\Common\GraphGradientMode $gradientMode;

    public ?\Grafana\Foundation\Common\GraphThresholdsStyleConfig $thresholdsStyle;

    public ?string $lineColor;

    public ?float $lineWidth;

    public ?\Grafana\Foundation\Common\LineInterpolation $lineInterpolation;

    public ?\Grafana\Foundation\Common\LineStyle $lineStyle;

    public ?string $fillColor;

    public ?float $fillOpacity;

    public ?\Grafana\Foundation\Common\VisibilityMode $showPoints;

    public ?float $pointSize;

    public ?string $pointColor;

    public ?\Grafana\Foundation\Common\AxisPlacement $axisPlacement;

    public ?\Grafana\Foundation\Common\AxisColorMode $axisColorMode;

    public ?string $axisLabel;

    public ?float $axisWidth;

    public ?float $axisSoftMin;

    public ?float $axisSoftMax;

    public ?bool $axisGridShow;

    public ?\Grafana\Foundation\Common\ScaleDistributionConfig $scaleDistribution;

    public ?bool $axisCenteredZero;

    public ?\Grafana\Foundation\Common\BarAlignment $barAlignment;

    public ?float $barWidthFactor;

    public ?\Grafana\Foundation\Common\StackingConfig $stacking;

    public ?\Grafana\Foundation\Common\HideSeriesConfig $hideFrom;

    public ?bool $hideValue;

    public ?\Grafana\Foundation\Common\GraphTransform $transform;

    /**
     * Indicate if null values should be treated as gaps or connected.
     * When the value is a number, it represents the maximum delta in the
     * X axis that should be considered connected.  For timeseries, this is milliseconds
     * @var bool|float|null
     */
    public $spanNulls;

    public ?string $fillBelowTo;

    public ?string $pointSymbol;

    public ?bool $axisBorderShow;

    public ?float $barMaxWidth;

    /**
     * @param \Grafana\Foundation\Common\GraphDrawStyle|null $drawStyle
     * @param \Grafana\Foundation\Common\GraphGradientMode|null $gradientMode
     * @param \Grafana\Foundation\Common\GraphThresholdsStyleConfig|null $thresholdsStyle
     * @param string|null $lineColor
     * @param float|null $lineWidth
     * @param \Grafana\Foundation\Common\LineInterpolation|null $lineInterpolation
     * @param \Grafana\Foundation\Common\LineStyle|null $lineStyle
     * @param string|null $fillColor
     * @param float|null $fillOpacity
     * @param \Grafana\Foundation\Common\VisibilityMode|null $showPoints
     * @param float|null $pointSize
     * @param string|null $pointColor
     * @param \Grafana\Foundation\Common\AxisPlacement|null $axisPlacement
     * @param \Grafana\Foundation\Common\AxisColorMode|null $axisColorMode
     * @param string|null $axisLabel
     * @param float|null $axisWidth
     * @param float|null $axisSoftMin
     * @param float|null $axisSoftMax
     * @param bool|null $axisGridShow
     * @param \Grafana\Foundation\Common\ScaleDistributionConfig|null $scaleDistribution
     * @param bool|null $axisCenteredZero
     * @param \Grafana\Foundation\Common\BarAlignment|null $barAlignment
     * @param float|null $barWidthFactor
     * @param \Grafana\Foundation\Common\StackingConfig|null $stacking
     * @param \Grafana\Foundation\Common\HideSeriesConfig|null $hideFrom
     * @param bool|null $hideValue
     * @param \Grafana\Foundation\Common\GraphTransform|null $transform
     * @param bool|float|null $spanNulls
     * @param string|null $fillBelowTo
     * @param string|null $pointSymbol
     * @param bool|null $axisBorderShow
     * @param float|null $barMaxWidth
     */
    public function __construct(?\Grafana\Foundation\Common\GraphDrawStyle $drawStyle = null, ?\Grafana\Foundation\Common\GraphGradientMode $gradientMode = null, ?\Grafana\Foundation\Common\GraphThresholdsStyleConfig $thresholdsStyle = null, ?string $lineColor = null, ?float $lineWidth = null, ?\Grafana\Foundation\Common\LineInterpolation $lineInterpolation = null, ?\Grafana\Foundation\Common\LineStyle $lineStyle = null, ?string $fillColor = null, ?float $fillOpacity = null, ?\Grafana\Foundation\Common\VisibilityMode $showPoints = null, ?float $pointSize = null, ?string $pointColor = null, ?\Grafana\Foundation\Common\AxisPlacement $axisPlacement = null, ?\Grafana\Foundation\Common\AxisColorMode $axisColorMode = null, ?string $axisLabel = null, ?float $axisWidth = null, ?float $axisSoftMin = null, ?float $axisSoftMax = null, ?bool $axisGridShow = null, ?\Grafana\Foundation\Common\ScaleDistributionConfig $scaleDistribution = null, ?bool $axisCenteredZero = null, ?\Grafana\Foundation\Common\BarAlignment $barAlignment = null, ?float $barWidthFactor = null, ?\Grafana\Foundation\Common\StackingConfig $stacking = null, ?\Grafana\Foundation\Common\HideSeriesConfig $hideFrom = null, ?bool $hideValue = null, ?\Grafana\Foundation\Common\GraphTransform $transform = null,  $spanNulls = null, ?string $fillBelowTo = null, ?string $pointSymbol = null, ?bool $axisBorderShow = null, ?float $barMaxWidth = null)
    {
        $this->type = "sparkline";
    
        $this->drawStyle = $drawStyle;
        $this->gradientMode = $gradientMode;
        $this->thresholdsStyle = $thresholdsStyle;
        $this->lineColor = $lineColor;
        $this->lineWidth = $lineWidth;
        $this->lineInterpolation = $lineInterpolation;
        $this->lineStyle = $lineStyle;
        $this->fillColor = $fillColor;
        $this->fillOpacity = $fillOpacity;
        $this->showPoints = $showPoints;
        $this->pointSize = $pointSize;
        $this->pointColor = $pointColor;
        $this->axisPlacement = $axisPlacement;
        $this->axisColorMode = $axisColorMode;
        $this->axisLabel = $axisLabel;
        $this->axisWidth = $axisWidth;
        $this->axisSoftMin = $axisSoftMin;
        $this->axisSoftMax = $axisSoftMax;
        $this->axisGridShow = $axisGridShow;
        $this->scaleDistribution = $scaleDistribution;
        $this->axisCenteredZero = $axisCenteredZero;
        $this->barAlignment = $barAlignment;
        $this->barWidthFactor = $barWidthFactor;
        $this->stacking = $stacking;
        $this->hideFrom = $hideFrom;
        $this->hideValue = $hideValue;
        $this->transform = $transform;
        $this->spanNulls = $spanNulls;
        $this->fillBelowTo = $fillBelowTo;
        $this->pointSymbol = $pointSymbol;
        $this->axisBorderShow = $axisBorderShow;
        $this->barMaxWidth = $barMaxWidth;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{type?: string, drawStyle?: string, gradientMode?: string, thresholdsStyle?: mixed, lineColor?: string, lineWidth?: float, lineInterpolation?: string, lineStyle?: mixed, fillColor?: string, fillOpacity?: float, showPoints?: string, pointSize?: float, pointColor?: string, axisPlacement?: string, axisColorMode?: string, axisLabel?: string, axisWidth?: float, axisSoftMin?: float, axisSoftMax?: float, axisGridShow?: bool, scaleDistribution?: mixed, axisCenteredZero?: bool, barAlignment?: int, barWidthFactor?: float, stacking?: mixed, hideFrom?: mixed, hideValue?: bool, transform?: string, spanNulls?: bool|float, fillBelowTo?: string, pointSymbol?: string, axisBorderShow?: bool, barMaxWidth?: float} $inputData */
        $data = $inputData;
        return new self(
            drawStyle: isset($data["drawStyle"]) ? (function($input) { return \Grafana\Foundation\Common\GraphDrawStyle::fromValue($input); })($data["drawStyle"]) : null,
            gradientMode: isset($data["gradientMode"]) ? (function($input) { return \Grafana\Foundation\Common\GraphGradientMode::fromValue($input); })($data["gradientMode"]) : null,
            thresholdsStyle: isset($data["thresholdsStyle"]) ? (function($input) {
    	/** @var array{mode?: string} */
    $val = $input;
    	return \Grafana\Foundation\Common\GraphThresholdsStyleConfig::fromArray($val);
    })($data["thresholdsStyle"]) : null,
            lineColor: $data["lineColor"] ?? null,
            lineWidth: $data["lineWidth"] ?? null,
            lineInterpolation: isset($data["lineInterpolation"]) ? (function($input) { return \Grafana\Foundation\Common\LineInterpolation::fromValue($input); })($data["lineInterpolation"]) : null,
            lineStyle: isset($data["lineStyle"]) ? (function($input) {
    	/** @var array{fill?: string, dash?: array<float>} */
    $val = $input;
    	return \Grafana\Foundation\Common\LineStyle::fromArray($val);
    })($data["lineStyle"]) : null,
            fillColor: $data["fillColor"] ?? null,
            fillOpacity: $data["fillOpacity"] ?? null,
            showPoints: isset($data["showPoints"]) ? (function($input) { return \Grafana\Foundation\Common\VisibilityMode::fromValue($input); })($data["showPoints"]) : null,
            pointSize: $data["pointSize"] ?? null,
            pointColor: $data["pointColor"] ?? null,
            axisPlacement: isset($data["axisPlacement"]) ? (function($input) { return \Grafana\Foundation\Common\AxisPlacement::fromValue($input); })($data["axisPlacement"]) : null,
            axisColorMode: isset($data["axisColorMode"]) ? (function($input) { return \Grafana\Foundation\Common\AxisColorMode::fromValue($input); })($data["axisColorMode"]) : null,
            axisLabel: $data["axisLabel"] ?? null,
            axisWidth: $data["axisWidth"] ?? null,
            axisSoftMin: $data["axisSoftMin"] ?? null,
            axisSoftMax: $data["axisSoftMax"] ?? null,
            axisGridShow: $data["axisGridShow"] ?? null,
            scaleDistribution: isset($data["scaleDistribution"]) ? (function($input) {
    	/** @var array{type?: string, log?: float, linearThreshold?: float} */
    $val = $input;
    	return \Grafana\Foundation\Common\ScaleDistributionConfig::fromArray($val);
    })($data["scaleDistribution"]) : null,
            axisCenteredZero: $data["axisCenteredZero"] ?? null,
            barAlignment: isset($data["barAlignment"]) ? (function($input) { return \Grafana\Foundation\Common\BarAlignment::fromValue($input); })($data["barAlignment"]) : null,
            barWidthFactor: $data["barWidthFactor"] ?? null,
            stacking: isset($data["stacking"]) ? (function($input) {
    	/** @var array{mode?: string, group?: string} */
    $val = $input;
    	return \Grafana\Foundation\Common\StackingConfig::fromArray($val);
    })($data["stacking"]) : null,
            hideFrom: isset($data["hideFrom"]) ? (function($input) {
    	/** @var array{tooltip?: bool, legend?: bool, viz?: bool} */
    $val = $input;
    	return \Grafana\Foundation\Common\HideSeriesConfig::fromArray($val);
    })($data["hideFrom"]) : null,
            hideValue: $data["hideValue"] ?? null,
            transform: isset($data["transform"]) ? (function($input) { return \Grafana\Foundation\Common\GraphTransform::fromValue($input); })($data["transform"]) : null,
            spanNulls: isset($data["spanNulls"]) ? (function($input) {
        switch (true) {
        case is_bool($input):
            return $input;
        case is_float($input):
            return $input;
        default:
            throw new \ValueError('incorrect value for disjunction');
    }
    })($data["spanNulls"]) : null,
            fillBelowTo: $data["fillBelowTo"] ?? null,
            pointSymbol: $data["pointSymbol"] ?? null,
            axisBorderShow: $data["axisBorderShow"] ?? null,
            barMaxWidth: $data["barMaxWidth"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "type" => $this->type,
        ];
        if (isset($this->drawStyle)) {
            $data["drawStyle"] = $this->drawStyle;
        }
        if (isset($this->gradientMode)) {
            $data["gradientMode"] = $this->gradientMode;
        }
        if (isset($this->thresholdsStyle)) {
            $data["thresholdsStyle"] = $this->thresholdsStyle;
        }
        if (isset($this->lineColor)) {
            $data["lineColor"] = $this->lineColor;
        }
        if (isset($this->lineWidth)) {
            $data["lineWidth"] = $this->lineWidth;
        }
        if (isset($this->lineInterpolation)) {
            $data["lineInterpolation"] = $this->lineInterpolation;
        }
        if (isset($this->lineStyle)) {
            $data["lineStyle"] = $this->lineStyle;
        }
        if (isset($this->fillColor)) {
            $data["fillColor"] = $this->fillColor;
        }
        if (isset($this->fillOpacity)) {
            $data["fillOpacity"] = $this->fillOpacity;
        }
        if (isset($this->showPoints)) {
            $data["showPoints"] = $this->showPoints;
        }
        if (isset($this->pointSize)) {
            $data["pointSize"] = $this->pointSize;
        }
        if (isset($this->pointColor)) {
            $data["pointColor"] = $this->pointColor;
        }
        if (isset($this->axisPlacement)) {
            $data["axisPlacement"] = $this->axisPlacement;
        }
        if (isset($this->axisColorMode)) {
            $data["axisColorMode"] = $this->axisColorMode;
        }
        if (isset($this->axisLabel)) {
            $data["axisLabel"] = $this->axisLabel;
        }
        if (isset($this->axisWidth)) {
            $data["axisWidth"] = $this->axisWidth;
        }
        if (isset($this->axisSoftMin)) {
            $data["axisSoftMin"] = $this->axisSoftMin;
        }
        if (isset($this->axisSoftMax)) {
            $data["axisSoftMax"] = $this->axisSoftMax;
        }
        if (isset($this->axisGridShow)) {
            $data["axisGridShow"] = $this->axisGridShow;
        }
        if (isset($this->scaleDistribution)) {
            $data["scaleDistribution"] = $this->scaleDistribution;
        }
        if (isset($this->axisCenteredZero)) {
            $data["axisCenteredZero"] = $this->axisCenteredZero;
        }
        if (isset($this->barAlignment)) {
            $data["barAlignment"] = $this->barAlignment;
        }
        if (isset($this->barWidthFactor)) {
            $data["barWidthFactor"] = $this->barWidthFactor;
        }
        if (isset($this->stacking)) {
            $data["stacking"] = $this->stacking;
        }
        if (isset($this->hideFrom)) {
            $data["hideFrom"] = $this->hideFrom;
        }
        if (isset($this->hideValue)) {
            $data["hideValue"] = $this->hideValue;
        }
        if (isset($this->transform)) {
            $data["transform"] = $this->transform;
        }
        if (isset($this->spanNulls)) {
            $data["spanNulls"] = $this->spanNulls;
        }
        if (isset($this->fillBelowTo)) {
            $data["fillBelowTo"] = $this->fillBelowTo;
        }
        if (isset($this->pointSymbol)) {
            $data["pointSymbol"] = $this->pointSymbol;
        }
        if (isset($this->axisBorderShow)) {
            $data["axisBorderShow"] = $this->axisBorderShow;
        }
        if (isset($this->barMaxWidth)) {
            $data["barMaxWidth"] = $this->barMaxWidth;
        }
        return $data;
    }
}
