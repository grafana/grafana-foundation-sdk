<?php

namespace Grafana\Foundation\Xychart;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Xychart\ScatterSeriesConfig>
 */
class ScatterSeriesConfigBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Xychart\ScatterSeriesConfig $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Xychart\ScatterSeriesConfig();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Xychart\ScatterSeriesConfig
     */
    public function build()
    {
        return $this->internal;
    }

    public function x(string $x): static
    {
        $this->internal->x = $x;
    
        return $this;
    }
    public function y(string $y): static
    {
        $this->internal->y = $y;
    
        return $this;
    }
    public function show(\Grafana\Foundation\Xychart\ScatterShow $show): static
    {
        $this->internal->show = $show;
    
        return $this;
    }
    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\ScaleDimensionConfig> $pointSize
     */
    public function pointSize(\Grafana\Foundation\Cog\Builder $pointSize): static
    {
        $pointSizeResource = $pointSize->build();
        $this->internal->pointSize = $pointSizeResource;
    
        return $this;
    }
    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\ColorDimensionConfig> $pointColor
     */
    public function pointColor(\Grafana\Foundation\Cog\Builder $pointColor): static
    {
        $pointColorResource = $pointColor->build();
        $this->internal->pointColor = $pointColorResource;
    
        return $this;
    }
    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\ColorDimensionConfig> $lineColor
     */
    public function lineColor(\Grafana\Foundation\Cog\Builder $lineColor): static
    {
        $lineColorResource = $lineColor->build();
        $this->internal->lineColor = $lineColorResource;
    
        return $this;
    }
    public function lineWidth(int $lineWidth): static
    {
        if (!($lineWidth >= 0)) {
            throw new \ValueError('$lineWidth must be >= 0');
        }
        $this->internal->lineWidth = $lineWidth;
    
        return $this;
    }
    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\LineStyle> $lineStyle
     */
    public function lineStyle(\Grafana\Foundation\Cog\Builder $lineStyle): static
    {
        $lineStyleResource = $lineStyle->build();
        $this->internal->lineStyle = $lineStyleResource;
    
        return $this;
    }
    public function label(\Grafana\Foundation\Common\VisibilityMode $label): static
    {
        $this->internal->label = $label;
    
        return $this;
    }
    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\HideSeriesConfig> $hideFrom
     */
    public function hideFrom(\Grafana\Foundation\Cog\Builder $hideFrom): static
    {
        $hideFromResource = $hideFrom->build();
        $this->internal->hideFrom = $hideFromResource;
    
        return $this;
    }
    public function axisPlacement(\Grafana\Foundation\Common\AxisPlacement $axisPlacement): static
    {
        $this->internal->axisPlacement = $axisPlacement;
    
        return $this;
    }
    public function axisColorMode(\Grafana\Foundation\Common\AxisColorMode $axisColorMode): static
    {
        $this->internal->axisColorMode = $axisColorMode;
    
        return $this;
    }
    public function axisLabel(string $axisLabel): static
    {
        $this->internal->axisLabel = $axisLabel;
    
        return $this;
    }
    public function axisWidth(float $axisWidth): static
    {
        $this->internal->axisWidth = $axisWidth;
    
        return $this;
    }
    public function axisSoftMin(float $axisSoftMin): static
    {
        $this->internal->axisSoftMin = $axisSoftMin;
    
        return $this;
    }
    public function axisSoftMax(float $axisSoftMax): static
    {
        $this->internal->axisSoftMax = $axisSoftMax;
    
        return $this;
    }
    public function axisGridShow(bool $axisGridShow): static
    {
        $this->internal->axisGridShow = $axisGridShow;
    
        return $this;
    }
    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\ScaleDistributionConfig> $scaleDistribution
     */
    public function scaleDistribution(\Grafana\Foundation\Cog\Builder $scaleDistribution): static
    {
        $scaleDistributionResource = $scaleDistribution->build();
        $this->internal->scaleDistribution = $scaleDistributionResource;
    
        return $this;
    }
    public function name(string $name): static
    {
        $this->internal->name = $name;
    
        return $this;
    }
    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\TextDimensionConfig> $labelValue
     */
    public function labelValue(\Grafana\Foundation\Cog\Builder $labelValue): static
    {
        $labelValueResource = $labelValue->build();
        $this->internal->labelValue = $labelValueResource;
    
        return $this;
    }
    public function axisCenteredZero(bool $axisCenteredZero): static
    {
        $this->internal->axisCenteredZero = $axisCenteredZero;
    
        return $this;
    }

}
