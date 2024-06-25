<?php

namespace Grafana\Foundation\Xychart;

class ScatterSeriesConfig implements \JsonSerializable
{
    public ?string $x;

    public ?string $y;

    public ?\Grafana\Foundation\Xychart\ScatterShow $show;

    public ?\Grafana\Foundation\Common\ScaleDimensionConfig $pointSize;

    public ?\Grafana\Foundation\Common\ColorDimensionConfig $pointColor;

    public ?\Grafana\Foundation\Common\ColorDimensionConfig $lineColor;

    public ?int $lineWidth;

    public ?\Grafana\Foundation\Common\LineStyle $lineStyle;

    public ?\Grafana\Foundation\Common\VisibilityMode $label;

    public ?\Grafana\Foundation\Common\HideSeriesConfig $hideFrom;

    public ?\Grafana\Foundation\Common\AxisPlacement $axisPlacement;

    public ?\Grafana\Foundation\Common\AxisColorMode $axisColorMode;

    public ?string $axisLabel;

    public ?float $axisWidth;

    public ?float $axisSoftMin;

    public ?float $axisSoftMax;

    public ?bool $axisGridShow;

    public ?\Grafana\Foundation\Common\ScaleDistributionConfig $scaleDistribution;

    public ?string $name;

    public ?\Grafana\Foundation\Common\TextDimensionConfig $labelValue;

    public ?bool $axisCenteredZero;

    /**
     * @param string|null $x
     * @param string|null $y
     * @param \Grafana\Foundation\Xychart\ScatterShow|null $show
     * @param \Grafana\Foundation\Common\ScaleDimensionConfig|null $pointSize
     * @param \Grafana\Foundation\Common\ColorDimensionConfig|null $pointColor
     * @param \Grafana\Foundation\Common\ColorDimensionConfig|null $lineColor
     * @param int|null $lineWidth
     * @param \Grafana\Foundation\Common\LineStyle|null $lineStyle
     * @param \Grafana\Foundation\Common\VisibilityMode|null $label
     * @param \Grafana\Foundation\Common\HideSeriesConfig|null $hideFrom
     * @param \Grafana\Foundation\Common\AxisPlacement|null $axisPlacement
     * @param \Grafana\Foundation\Common\AxisColorMode|null $axisColorMode
     * @param string|null $axisLabel
     * @param float|null $axisWidth
     * @param float|null $axisSoftMin
     * @param float|null $axisSoftMax
     * @param bool|null $axisGridShow
     * @param \Grafana\Foundation\Common\ScaleDistributionConfig|null $scaleDistribution
     * @param string|null $name
     * @param \Grafana\Foundation\Common\TextDimensionConfig|null $labelValue
     * @param bool|null $axisCenteredZero
     */
    public function __construct(?string $x = null, ?string $y = null, ?\Grafana\Foundation\Xychart\ScatterShow $show = null, ?\Grafana\Foundation\Common\ScaleDimensionConfig $pointSize = null, ?\Grafana\Foundation\Common\ColorDimensionConfig $pointColor = null, ?\Grafana\Foundation\Common\ColorDimensionConfig $lineColor = null, ?int $lineWidth = null, ?\Grafana\Foundation\Common\LineStyle $lineStyle = null, ?\Grafana\Foundation\Common\VisibilityMode $label = null, ?\Grafana\Foundation\Common\HideSeriesConfig $hideFrom = null, ?\Grafana\Foundation\Common\AxisPlacement $axisPlacement = null, ?\Grafana\Foundation\Common\AxisColorMode $axisColorMode = null, ?string $axisLabel = null, ?float $axisWidth = null, ?float $axisSoftMin = null, ?float $axisSoftMax = null, ?bool $axisGridShow = null, ?\Grafana\Foundation\Common\ScaleDistributionConfig $scaleDistribution = null, ?string $name = null, ?\Grafana\Foundation\Common\TextDimensionConfig $labelValue = null, ?bool $axisCenteredZero = null)
    {
        $this->x = $x;
        $this->y = $y;
        $this->show = $show;
        $this->pointSize = $pointSize;
        $this->pointColor = $pointColor;
        $this->lineColor = $lineColor;
        $this->lineWidth = $lineWidth;
        $this->lineStyle = $lineStyle;
        $this->label = $label;
        $this->hideFrom = $hideFrom;
        $this->axisPlacement = $axisPlacement;
        $this->axisColorMode = $axisColorMode;
        $this->axisLabel = $axisLabel;
        $this->axisWidth = $axisWidth;
        $this->axisSoftMin = $axisSoftMin;
        $this->axisSoftMax = $axisSoftMax;
        $this->axisGridShow = $axisGridShow;
        $this->scaleDistribution = $scaleDistribution;
        $this->name = $name;
        $this->labelValue = $labelValue;
        $this->axisCenteredZero = $axisCenteredZero;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{x?: string, y?: string, show?: string, pointSize?: mixed, pointColor?: mixed, lineColor?: mixed, lineWidth?: int, lineStyle?: mixed, label?: string, hideFrom?: mixed, axisPlacement?: string, axisColorMode?: string, axisLabel?: string, axisWidth?: float, axisSoftMin?: float, axisSoftMax?: float, axisGridShow?: bool, scaleDistribution?: mixed, name?: string, labelValue?: mixed, axisCenteredZero?: bool} $inputData */
        $data = $inputData;
        return new self(
            x: $data["x"] ?? null,
            y: $data["y"] ?? null,
            show: isset($data["show"]) ? (function($input) { return \Grafana\Foundation\Xychart\ScatterShow::fromValue($input); })($data["show"]) : null,
            pointSize: isset($data["pointSize"]) ? (function($input) {
    	/** @var array{min?: float, max?: float, fixed?: float, field?: string, mode?: string} */
    $val = $input;
    	return \Grafana\Foundation\Common\ScaleDimensionConfig::fromArray($val);
    })($data["pointSize"]) : null,
            pointColor: isset($data["pointColor"]) ? (function($input) {
    	/** @var array{fixed?: string, field?: string} */
    $val = $input;
    	return \Grafana\Foundation\Common\ColorDimensionConfig::fromArray($val);
    })($data["pointColor"]) : null,
            lineColor: isset($data["lineColor"]) ? (function($input) {
    	/** @var array{fixed?: string, field?: string} */
    $val = $input;
    	return \Grafana\Foundation\Common\ColorDimensionConfig::fromArray($val);
    })($data["lineColor"]) : null,
            lineWidth: $data["lineWidth"] ?? null,
            lineStyle: isset($data["lineStyle"]) ? (function($input) {
    	/** @var array{fill?: string, dash?: array<float>} */
    $val = $input;
    	return \Grafana\Foundation\Common\LineStyle::fromArray($val);
    })($data["lineStyle"]) : null,
            label: isset($data["label"]) ? (function($input) { return \Grafana\Foundation\Common\VisibilityMode::fromValue($input); })($data["label"]) : null,
            hideFrom: isset($data["hideFrom"]) ? (function($input) {
    	/** @var array{tooltip?: bool, legend?: bool, viz?: bool} */
    $val = $input;
    	return \Grafana\Foundation\Common\HideSeriesConfig::fromArray($val);
    })($data["hideFrom"]) : null,
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
            name: $data["name"] ?? null,
            labelValue: isset($data["labelValue"]) ? (function($input) {
    	/** @var array{mode?: string, field?: string, fixed?: string} */
    $val = $input;
    	return \Grafana\Foundation\Common\TextDimensionConfig::fromArray($val);
    })($data["labelValue"]) : null,
            axisCenteredZero: $data["axisCenteredZero"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
        ];
        if (isset($this->x)) {
            $data["x"] = $this->x;
        }
        if (isset($this->y)) {
            $data["y"] = $this->y;
        }
        if (isset($this->show)) {
            $data["show"] = $this->show;
        }
        if (isset($this->pointSize)) {
            $data["pointSize"] = $this->pointSize;
        }
        if (isset($this->pointColor)) {
            $data["pointColor"] = $this->pointColor;
        }
        if (isset($this->lineColor)) {
            $data["lineColor"] = $this->lineColor;
        }
        if (isset($this->lineWidth)) {
            $data["lineWidth"] = $this->lineWidth;
        }
        if (isset($this->lineStyle)) {
            $data["lineStyle"] = $this->lineStyle;
        }
        if (isset($this->label)) {
            $data["label"] = $this->label;
        }
        if (isset($this->hideFrom)) {
            $data["hideFrom"] = $this->hideFrom;
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
        if (isset($this->name)) {
            $data["name"] = $this->name;
        }
        if (isset($this->labelValue)) {
            $data["labelValue"] = $this->labelValue;
        }
        if (isset($this->axisCenteredZero)) {
            $data["axisCenteredZero"] = $this->axisCenteredZero;
        }
        return $data;
    }
}
