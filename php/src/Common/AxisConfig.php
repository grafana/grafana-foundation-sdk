<?php

namespace Grafana\Foundation\Common;

/**
 * TODO docs
 */
class AxisConfig implements \JsonSerializable
{
    public ?\Grafana\Foundation\Common\AxisPlacement $axisPlacement;

    public ?\Grafana\Foundation\Common\AxisColorMode $axisColorMode;

    public ?string $axisLabel;

    public ?float $axisWidth;

    public ?float $axisSoftMin;

    public ?float $axisSoftMax;

    public ?bool $axisGridShow;

    public ?\Grafana\Foundation\Common\ScaleDistributionConfig $scaleDistribution;

    public ?bool $axisCenteredZero;

    public ?bool $axisBorderShow;

    /**
     * @param \Grafana\Foundation\Common\AxisPlacement|null $axisPlacement
     * @param \Grafana\Foundation\Common\AxisColorMode|null $axisColorMode
     * @param string|null $axisLabel
     * @param float|null $axisWidth
     * @param float|null $axisSoftMin
     * @param float|null $axisSoftMax
     * @param bool|null $axisGridShow
     * @param \Grafana\Foundation\Common\ScaleDistributionConfig|null $scaleDistribution
     * @param bool|null $axisCenteredZero
     * @param bool|null $axisBorderShow
     */
    public function __construct(?\Grafana\Foundation\Common\AxisPlacement $axisPlacement = null, ?\Grafana\Foundation\Common\AxisColorMode $axisColorMode = null, ?string $axisLabel = null, ?float $axisWidth = null, ?float $axisSoftMin = null, ?float $axisSoftMax = null, ?bool $axisGridShow = null, ?\Grafana\Foundation\Common\ScaleDistributionConfig $scaleDistribution = null, ?bool $axisCenteredZero = null, ?bool $axisBorderShow = null)
    {
        $this->axisPlacement = $axisPlacement;
        $this->axisColorMode = $axisColorMode;
        $this->axisLabel = $axisLabel;
        $this->axisWidth = $axisWidth;
        $this->axisSoftMin = $axisSoftMin;
        $this->axisSoftMax = $axisSoftMax;
        $this->axisGridShow = $axisGridShow;
        $this->scaleDistribution = $scaleDistribution;
        $this->axisCenteredZero = $axisCenteredZero;
        $this->axisBorderShow = $axisBorderShow;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{axisPlacement?: string, axisColorMode?: string, axisLabel?: string, axisWidth?: float, axisSoftMin?: float, axisSoftMax?: float, axisGridShow?: bool, scaleDistribution?: mixed, axisCenteredZero?: bool, axisBorderShow?: bool} $inputData */
        $data = $inputData;
        return new self(
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
        if (isset($this->axisBorderShow)) {
            $data["axisBorderShow"] = $this->axisBorderShow;
        }
        return $data;
    }
}
