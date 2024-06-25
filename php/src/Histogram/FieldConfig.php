<?php

namespace Grafana\Foundation\Histogram;

class FieldConfig implements \JsonSerializable
{
    /**
     * Controls line width of the bars.
     */
    public ?int $lineWidth;

    /**
     * Controls the fill opacity of the bars.
     */
    public ?int $fillOpacity;

    public ?\Grafana\Foundation\Common\AxisPlacement $axisPlacement;

    public ?\Grafana\Foundation\Common\AxisColorMode $axisColorMode;

    public ?string $axisLabel;

    public ?float $axisWidth;

    public ?float $axisSoftMin;

    public ?float $axisSoftMax;

    public ?bool $axisGridShow;

    public ?\Grafana\Foundation\Common\ScaleDistributionConfig $scaleDistribution;

    public ?\Grafana\Foundation\Common\HideSeriesConfig $hideFrom;

    /**
     * Set the mode of the gradient fill. Fill gradient is based on the line color. To change the color, use the standard color scheme field option.
     * Gradient appearance is influenced by the Fill opacity setting.
     */
    public ?\Grafana\Foundation\Common\GraphGradientMode $gradientMode;

    public ?bool $axisCenteredZero;

    /**
     * @param int|null $lineWidth
     * @param int|null $fillOpacity
     * @param \Grafana\Foundation\Common\AxisPlacement|null $axisPlacement
     * @param \Grafana\Foundation\Common\AxisColorMode|null $axisColorMode
     * @param string|null $axisLabel
     * @param float|null $axisWidth
     * @param float|null $axisSoftMin
     * @param float|null $axisSoftMax
     * @param bool|null $axisGridShow
     * @param \Grafana\Foundation\Common\ScaleDistributionConfig|null $scaleDistribution
     * @param \Grafana\Foundation\Common\HideSeriesConfig|null $hideFrom
     * @param \Grafana\Foundation\Common\GraphGradientMode|null $gradientMode
     * @param bool|null $axisCenteredZero
     */
    public function __construct(?int $lineWidth = null, ?int $fillOpacity = null, ?\Grafana\Foundation\Common\AxisPlacement $axisPlacement = null, ?\Grafana\Foundation\Common\AxisColorMode $axisColorMode = null, ?string $axisLabel = null, ?float $axisWidth = null, ?float $axisSoftMin = null, ?float $axisSoftMax = null, ?bool $axisGridShow = null, ?\Grafana\Foundation\Common\ScaleDistributionConfig $scaleDistribution = null, ?\Grafana\Foundation\Common\HideSeriesConfig $hideFrom = null, ?\Grafana\Foundation\Common\GraphGradientMode $gradientMode = null, ?bool $axisCenteredZero = null)
    {
        $this->lineWidth = $lineWidth;
        $this->fillOpacity = $fillOpacity;
        $this->axisPlacement = $axisPlacement;
        $this->axisColorMode = $axisColorMode;
        $this->axisLabel = $axisLabel;
        $this->axisWidth = $axisWidth;
        $this->axisSoftMin = $axisSoftMin;
        $this->axisSoftMax = $axisSoftMax;
        $this->axisGridShow = $axisGridShow;
        $this->scaleDistribution = $scaleDistribution;
        $this->hideFrom = $hideFrom;
        $this->gradientMode = $gradientMode;
        $this->axisCenteredZero = $axisCenteredZero;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{lineWidth?: int, fillOpacity?: int, axisPlacement?: string, axisColorMode?: string, axisLabel?: string, axisWidth?: float, axisSoftMin?: float, axisSoftMax?: float, axisGridShow?: bool, scaleDistribution?: mixed, hideFrom?: mixed, gradientMode?: string, axisCenteredZero?: bool} $inputData */
        $data = $inputData;
        return new self(
            lineWidth: $data["lineWidth"] ?? null,
            fillOpacity: $data["fillOpacity"] ?? null,
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
            hideFrom: isset($data["hideFrom"]) ? (function($input) {
    	/** @var array{tooltip?: bool, legend?: bool, viz?: bool} */
    $val = $input;
    	return \Grafana\Foundation\Common\HideSeriesConfig::fromArray($val);
    })($data["hideFrom"]) : null,
            gradientMode: isset($data["gradientMode"]) ? (function($input) { return \Grafana\Foundation\Common\GraphGradientMode::fromValue($input); })($data["gradientMode"]) : null,
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
        if (isset($this->lineWidth)) {
            $data["lineWidth"] = $this->lineWidth;
        }
        if (isset($this->fillOpacity)) {
            $data["fillOpacity"] = $this->fillOpacity;
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
        if (isset($this->hideFrom)) {
            $data["hideFrom"] = $this->hideFrom;
        }
        if (isset($this->gradientMode)) {
            $data["gradientMode"] = $this->gradientMode;
        }
        if (isset($this->axisCenteredZero)) {
            $data["axisCenteredZero"] = $this->axisCenteredZero;
        }
        return $data;
    }
}
