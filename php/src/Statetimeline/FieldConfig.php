<?php

namespace Grafana\Foundation\Statetimeline;

class FieldConfig implements \JsonSerializable
{
    public ?int $lineWidth;

    public ?\Grafana\Foundation\Common\AxisPlacement $axisPlacement;

    public ?\Grafana\Foundation\Common\AxisColorMode $axisColorMode;

    public ?string $axisLabel;

    public ?float $axisWidth;

    public ?float $axisSoftMin;

    public ?float $axisSoftMax;

    public ?bool $axisGridShow;

    public ?\Grafana\Foundation\Common\ScaleDistributionConfig $scaleDistribution;

    public ?bool $axisCenteredZero;

    public ?\Grafana\Foundation\Common\HideSeriesConfig $hideFrom;

    public ?int $fillOpacity;

    public ?bool $axisBorderShow;

    /**
     * @param int|null $lineWidth
     * @param \Grafana\Foundation\Common\AxisPlacement|null $axisPlacement
     * @param \Grafana\Foundation\Common\AxisColorMode|null $axisColorMode
     * @param string|null $axisLabel
     * @param float|null $axisWidth
     * @param float|null $axisSoftMin
     * @param float|null $axisSoftMax
     * @param bool|null $axisGridShow
     * @param \Grafana\Foundation\Common\ScaleDistributionConfig|null $scaleDistribution
     * @param bool|null $axisCenteredZero
     * @param \Grafana\Foundation\Common\HideSeriesConfig|null $hideFrom
     * @param int|null $fillOpacity
     * @param bool|null $axisBorderShow
     */
    public function __construct(?int $lineWidth = null, ?\Grafana\Foundation\Common\AxisPlacement $axisPlacement = null, ?\Grafana\Foundation\Common\AxisColorMode $axisColorMode = null, ?string $axisLabel = null, ?float $axisWidth = null, ?float $axisSoftMin = null, ?float $axisSoftMax = null, ?bool $axisGridShow = null, ?\Grafana\Foundation\Common\ScaleDistributionConfig $scaleDistribution = null, ?bool $axisCenteredZero = null, ?\Grafana\Foundation\Common\HideSeriesConfig $hideFrom = null, ?int $fillOpacity = null, ?bool $axisBorderShow = null)
    {
        $this->lineWidth = $lineWidth;
        $this->axisPlacement = $axisPlacement;
        $this->axisColorMode = $axisColorMode;
        $this->axisLabel = $axisLabel;
        $this->axisWidth = $axisWidth;
        $this->axisSoftMin = $axisSoftMin;
        $this->axisSoftMax = $axisSoftMax;
        $this->axisGridShow = $axisGridShow;
        $this->scaleDistribution = $scaleDistribution;
        $this->axisCenteredZero = $axisCenteredZero;
        $this->hideFrom = $hideFrom;
        $this->fillOpacity = $fillOpacity;
        $this->axisBorderShow = $axisBorderShow;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{lineWidth?: int, axisPlacement?: string, axisColorMode?: string, axisLabel?: string, axisWidth?: float, axisSoftMin?: float, axisSoftMax?: float, axisGridShow?: bool, scaleDistribution?: mixed, axisCenteredZero?: bool, hideFrom?: mixed, fillOpacity?: int, axisBorderShow?: bool} $inputData */
        $data = $inputData;
        return new self(
            lineWidth: $data["lineWidth"] ?? null,
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
            hideFrom: isset($data["hideFrom"]) ? (function($input) {
    	/** @var array{tooltip?: bool, legend?: bool, viz?: bool} */
    $val = $input;
    	return \Grafana\Foundation\Common\HideSeriesConfig::fromArray($val);
    })($data["hideFrom"]) : null,
            fillOpacity: $data["fillOpacity"] ?? null,
            axisBorderShow: $data["axisBorderShow"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        if (isset($this->lineWidth)) {
            $data->lineWidth = $this->lineWidth;
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
        if (isset($this->axisCenteredZero)) {
            $data->axisCenteredZero = $this->axisCenteredZero;
        }
        if (isset($this->hideFrom)) {
            $data->hideFrom = $this->hideFrom;
        }
        if (isset($this->fillOpacity)) {
            $data->fillOpacity = $this->fillOpacity;
        }
        if (isset($this->axisBorderShow)) {
            $data->axisBorderShow = $this->axisBorderShow;
        }
        return $data;
    }
}
