<?php

namespace Grafana\Foundation\Xychart;

class FieldConfig implements \JsonSerializable
{
    public ?\Grafana\Foundation\Xychart\XYShowMode $show;

    public ?\Grafana\Foundation\Xychart\XychartFieldConfigPointSize $pointSize;

    public ?\Grafana\Foundation\Xychart\PointShape $pointShape;

    public ?int $pointStrokeWidth;

    public ?int $fillOpacity;

    public ?int $lineWidth;

    public ?\Grafana\Foundation\Common\HideSeriesConfig $hideFrom;

    public ?\Grafana\Foundation\Common\AxisPlacement $axisPlacement;

    public ?\Grafana\Foundation\Common\AxisColorMode $axisColorMode;

    public ?string $axisLabel;

    public ?float $axisWidth;

    public ?float $axisSoftMin;

    public ?float $axisSoftMax;

    public ?bool $axisGridShow;

    public ?\Grafana\Foundation\Common\ScaleDistributionConfig $scaleDistribution;

    public ?bool $axisCenteredZero;

    public ?\Grafana\Foundation\Common\LineStyle $lineStyle;

    public ?bool $axisBorderShow;

    /**
     * @param \Grafana\Foundation\Xychart\XYShowMode|null $show
     * @param \Grafana\Foundation\Xychart\XychartFieldConfigPointSize|null $pointSize
     * @param \Grafana\Foundation\Xychart\PointShape|null $pointShape
     * @param int|null $pointStrokeWidth
     * @param int|null $fillOpacity
     * @param int|null $lineWidth
     * @param \Grafana\Foundation\Common\HideSeriesConfig|null $hideFrom
     * @param \Grafana\Foundation\Common\AxisPlacement|null $axisPlacement
     * @param \Grafana\Foundation\Common\AxisColorMode|null $axisColorMode
     * @param string|null $axisLabel
     * @param float|null $axisWidth
     * @param float|null $axisSoftMin
     * @param float|null $axisSoftMax
     * @param bool|null $axisGridShow
     * @param \Grafana\Foundation\Common\ScaleDistributionConfig|null $scaleDistribution
     * @param bool|null $axisCenteredZero
     * @param \Grafana\Foundation\Common\LineStyle|null $lineStyle
     * @param bool|null $axisBorderShow
     */
    public function __construct(?\Grafana\Foundation\Xychart\XYShowMode $show = null, ?\Grafana\Foundation\Xychart\XychartFieldConfigPointSize $pointSize = null, ?\Grafana\Foundation\Xychart\PointShape $pointShape = null, ?int $pointStrokeWidth = null, ?int $fillOpacity = null, ?int $lineWidth = null, ?\Grafana\Foundation\Common\HideSeriesConfig $hideFrom = null, ?\Grafana\Foundation\Common\AxisPlacement $axisPlacement = null, ?\Grafana\Foundation\Common\AxisColorMode $axisColorMode = null, ?string $axisLabel = null, ?float $axisWidth = null, ?float $axisSoftMin = null, ?float $axisSoftMax = null, ?bool $axisGridShow = null, ?\Grafana\Foundation\Common\ScaleDistributionConfig $scaleDistribution = null, ?bool $axisCenteredZero = null, ?\Grafana\Foundation\Common\LineStyle $lineStyle = null, ?bool $axisBorderShow = null)
    {
        $this->show = $show;
        $this->pointSize = $pointSize;
        $this->pointShape = $pointShape;
        $this->pointStrokeWidth = $pointStrokeWidth;
        $this->fillOpacity = $fillOpacity;
        $this->lineWidth = $lineWidth;
        $this->hideFrom = $hideFrom;
        $this->axisPlacement = $axisPlacement;
        $this->axisColorMode = $axisColorMode;
        $this->axisLabel = $axisLabel;
        $this->axisWidth = $axisWidth;
        $this->axisSoftMin = $axisSoftMin;
        $this->axisSoftMax = $axisSoftMax;
        $this->axisGridShow = $axisGridShow;
        $this->scaleDistribution = $scaleDistribution;
        $this->axisCenteredZero = $axisCenteredZero;
        $this->lineStyle = $lineStyle;
        $this->axisBorderShow = $axisBorderShow;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{show?: string, pointSize?: mixed, pointShape?: string, pointStrokeWidth?: int, fillOpacity?: int, lineWidth?: int, hideFrom?: mixed, axisPlacement?: string, axisColorMode?: string, axisLabel?: string, axisWidth?: float, axisSoftMin?: float, axisSoftMax?: float, axisGridShow?: bool, scaleDistribution?: mixed, axisCenteredZero?: bool, lineStyle?: mixed, axisBorderShow?: bool} $inputData */
        $data = $inputData;
        return new self(
            show: isset($data["show"]) ? (function($input) { return \Grafana\Foundation\Xychart\XYShowMode::fromValue($input); })($data["show"]) : null,
            pointSize: isset($data["pointSize"]) ? (function($input) {
    	/** @var array{fixed?: int, min?: int, max?: int} */
    $val = $input;
    	return \Grafana\Foundation\Xychart\XychartFieldConfigPointSize::fromArray($val);
    })($data["pointSize"]) : null,
            pointShape: isset($data["pointShape"]) ? (function($input) { return \Grafana\Foundation\Xychart\PointShape::fromValue($input); })($data["pointShape"]) : null,
            pointStrokeWidth: $data["pointStrokeWidth"] ?? null,
            fillOpacity: $data["fillOpacity"] ?? null,
            lineWidth: $data["lineWidth"] ?? null,
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
            axisCenteredZero: $data["axisCenteredZero"] ?? null,
            lineStyle: isset($data["lineStyle"]) ? (function($input) {
    	/** @var array{fill?: string, dash?: array<float>} */
    $val = $input;
    	return \Grafana\Foundation\Common\LineStyle::fromArray($val);
    })($data["lineStyle"]) : null,
            axisBorderShow: $data["axisBorderShow"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
        ];
        if (isset($this->show)) {
            $data["show"] = $this->show;
        }
        if (isset($this->pointSize)) {
            $data["pointSize"] = $this->pointSize;
        }
        if (isset($this->pointShape)) {
            $data["pointShape"] = $this->pointShape;
        }
        if (isset($this->pointStrokeWidth)) {
            $data["pointStrokeWidth"] = $this->pointStrokeWidth;
        }
        if (isset($this->fillOpacity)) {
            $data["fillOpacity"] = $this->fillOpacity;
        }
        if (isset($this->lineWidth)) {
            $data["lineWidth"] = $this->lineWidth;
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
        if (isset($this->axisCenteredZero)) {
            $data["axisCenteredZero"] = $this->axisCenteredZero;
        }
        if (isset($this->lineStyle)) {
            $data["lineStyle"] = $this->lineStyle;
        }
        if (isset($this->axisBorderShow)) {
            $data["axisBorderShow"] = $this->axisBorderShow;
        }
        return $data;
    }
}
