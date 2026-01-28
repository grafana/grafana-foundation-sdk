<?php

namespace Grafana\Foundation\Xychart;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Xychart\FieldConfig>
 */
class FieldConfigBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Xychart\FieldConfig $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Xychart\FieldConfig();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Xychart\FieldConfig
     */
    public function build()
    {
        return $this->internal;
    }

    public function show(\Grafana\Foundation\Xychart\XYShowMode $show): static
    {
        $this->internal->show = $show;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Xychart\XychartFieldConfigPointSize> $pointSize
     */
    public function pointSize(\Grafana\Foundation\Cog\Builder $pointSize): static
    {
        $pointSizeResource = $pointSize->build();
        $this->internal->pointSize = $pointSizeResource;
    
        return $this;
    }

    public function pointShape(\Grafana\Foundation\Xychart\PointShape $pointShape): static
    {
        $this->internal->pointShape = $pointShape;
    
        return $this;
    }

    public function pointStrokeWidth(int $pointStrokeWidth): static
    {
        if (!($pointStrokeWidth >= 0)) {
            throw new \ValueError('$pointStrokeWidth must be >= 0');
        }
        $this->internal->pointStrokeWidth = $pointStrokeWidth;
    
        return $this;
    }

    public function fillOpacity(int $fillOpacity): static
    {
        if (!($fillOpacity <= 100)) {
            throw new \ValueError('$fillOpacity must be <= 100');
        }
        $this->internal->fillOpacity = $fillOpacity;
    
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

    public function axisCenteredZero(bool $axisCenteredZero): static
    {
        $this->internal->axisCenteredZero = $axisCenteredZero;
    
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

    public function axisBorderShow(bool $axisBorderShow): static
    {
        $this->internal->axisBorderShow = $axisBorderShow;
    
        return $this;
    }

}
