<?php

namespace Grafana\Foundation\Heatmap;

/**
 * Configuration options for the yAxis
 */
class YAxisConfig implements \JsonSerializable
{
    /**
     * Sets the yAxis unit
     */
    public ?string $unit;

    /**
     * Reverses the yAxis
     */
    public ?bool $reverse;

    /**
     * Controls the number of decimals for yAxis values
     */
    public ?float $decimals;

    /**
     * Sets the minimum value for the yAxis
     */
    public ?float $min;

    public ?\Grafana\Foundation\Common\AxisPlacement $axisPlacement;

    public ?\Grafana\Foundation\Common\AxisColorMode $axisColorMode;

    public ?string $axisLabel;

    public ?float $axisWidth;

    public ?float $axisSoftMin;

    public ?float $axisSoftMax;

    public ?bool $axisGridShow;

    public ?\Grafana\Foundation\Common\ScaleDistributionConfig $scaleDistribution;

    /**
     * Sets the maximum value for the yAxis
     */
    public ?float $max;

    public ?bool $axisCenteredZero;

    /**
     * @param string|null $unit
     * @param bool|null $reverse
     * @param float|null $decimals
     * @param float|null $min
     * @param \Grafana\Foundation\Common\AxisPlacement|null $axisPlacement
     * @param \Grafana\Foundation\Common\AxisColorMode|null $axisColorMode
     * @param string|null $axisLabel
     * @param float|null $axisWidth
     * @param float|null $axisSoftMin
     * @param float|null $axisSoftMax
     * @param bool|null $axisGridShow
     * @param \Grafana\Foundation\Common\ScaleDistributionConfig|null $scaleDistribution
     * @param float|null $max
     * @param bool|null $axisCenteredZero
     */
    public function __construct(?string $unit = null, ?bool $reverse = null, ?float $decimals = null, ?float $min = null, ?\Grafana\Foundation\Common\AxisPlacement $axisPlacement = null, ?\Grafana\Foundation\Common\AxisColorMode $axisColorMode = null, ?string $axisLabel = null, ?float $axisWidth = null, ?float $axisSoftMin = null, ?float $axisSoftMax = null, ?bool $axisGridShow = null, ?\Grafana\Foundation\Common\ScaleDistributionConfig $scaleDistribution = null, ?float $max = null, ?bool $axisCenteredZero = null)
    {
        $this->unit = $unit;
        $this->reverse = $reverse;
        $this->decimals = $decimals;
        $this->min = $min;
        $this->axisPlacement = $axisPlacement;
        $this->axisColorMode = $axisColorMode;
        $this->axisLabel = $axisLabel;
        $this->axisWidth = $axisWidth;
        $this->axisSoftMin = $axisSoftMin;
        $this->axisSoftMax = $axisSoftMax;
        $this->axisGridShow = $axisGridShow;
        $this->scaleDistribution = $scaleDistribution;
        $this->max = $max;
        $this->axisCenteredZero = $axisCenteredZero;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{unit?: string, reverse?: bool, decimals?: float, min?: float, axisPlacement?: string, axisColorMode?: string, axisLabel?: string, axisWidth?: float, axisSoftMin?: float, axisSoftMax?: float, axisGridShow?: bool, scaleDistribution?: mixed, max?: float, axisCenteredZero?: bool} $inputData */
        $data = $inputData;
        return new self(
            unit: $data["unit"] ?? null,
            reverse: $data["reverse"] ?? null,
            decimals: $data["decimals"] ?? null,
            min: $data["min"] ?? null,
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
            max: $data["max"] ?? null,
            axisCenteredZero: $data["axisCenteredZero"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        if (isset($this->unit)) {
            $data->unit = $this->unit;
        }
        if (isset($this->reverse)) {
            $data->reverse = $this->reverse;
        }
        if (isset($this->decimals)) {
            $data->decimals = $this->decimals;
        }
        if (isset($this->min)) {
            $data->min = $this->min;
        }
        if (isset($this->axisPlacement)) {
            $data->axisPlacement = $this->axisPlacement;
        }
        if (isset($this->axisColorMode)) {
            $data->axisColorMode = $this->axisColorMode;
        }
        if (isset($this->axisLabel)) {
            $data->axisLabel = $this->axisLabel;
        }
        if (isset($this->axisWidth)) {
            $data->axisWidth = $this->axisWidth;
        }
        if (isset($this->axisSoftMin)) {
            $data->axisSoftMin = $this->axisSoftMin;
        }
        if (isset($this->axisSoftMax)) {
            $data->axisSoftMax = $this->axisSoftMax;
        }
        if (isset($this->axisGridShow)) {
            $data->axisGridShow = $this->axisGridShow;
        }
        if (isset($this->scaleDistribution)) {
            $data->scaleDistribution = $this->scaleDistribution;
        }
        if (isset($this->max)) {
            $data->max = $this->max;
        }
        if (isset($this->axisCenteredZero)) {
            $data->axisCenteredZero = $this->axisCenteredZero;
        }
        return $data;
    }
}
