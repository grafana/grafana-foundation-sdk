<?php

namespace Grafana\Foundation\Common;

/**
 * Sparkline cell options
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\TableSparklineCellOptions>
 */
class TableSparklineCellOptionsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Common\TableSparklineCellOptions $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Common\TableSparklineCellOptions();
    $this->internal->type = "sparkline";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Common\TableSparklineCellOptions
     */
    public function build()
    {
        return $this->internal;
    }

    public function drawStyle(\Grafana\Foundation\Common\GraphDrawStyle $drawStyle): static
    {
        $this->internal->drawStyle = $drawStyle;
    
        return $this;
    }
    public function gradientMode(\Grafana\Foundation\Common\GraphGradientMode $gradientMode): static
    {
        $this->internal->gradientMode = $gradientMode;
    
        return $this;
    }
    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\GraphThresholdsStyleConfig> $thresholdsStyle
     */
    public function thresholdsStyle(\Grafana\Foundation\Cog\Builder $thresholdsStyle): static
    {
        $thresholdsStyleResource = $thresholdsStyle->build();
        $this->internal->thresholdsStyle = $thresholdsStyleResource;
    
        return $this;
    }
    public function lineColor(string $lineColor): static
    {
        $this->internal->lineColor = $lineColor;
    
        return $this;
    }
    public function lineWidth(float $lineWidth): static
    {
        $this->internal->lineWidth = $lineWidth;
    
        return $this;
    }
    public function lineInterpolation(\Grafana\Foundation\Common\LineInterpolation $lineInterpolation): static
    {
        $this->internal->lineInterpolation = $lineInterpolation;
    
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
    public function fillColor(string $fillColor): static
    {
        $this->internal->fillColor = $fillColor;
    
        return $this;
    }
    public function fillOpacity(float $fillOpacity): static
    {
        $this->internal->fillOpacity = $fillOpacity;
    
        return $this;
    }
    public function showPoints(\Grafana\Foundation\Common\VisibilityMode $showPoints): static
    {
        $this->internal->showPoints = $showPoints;
    
        return $this;
    }
    public function pointSize(float $pointSize): static
    {
        $this->internal->pointSize = $pointSize;
    
        return $this;
    }
    public function pointColor(string $pointColor): static
    {
        $this->internal->pointColor = $pointColor;
    
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
    public function barAlignment(\Grafana\Foundation\Common\BarAlignment $barAlignment): static
    {
        $this->internal->barAlignment = $barAlignment;
    
        return $this;
    }
    public function barWidthFactor(float $barWidthFactor): static
    {
        $this->internal->barWidthFactor = $barWidthFactor;
    
        return $this;
    }
    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\StackingConfig> $stacking
     */
    public function stacking(\Grafana\Foundation\Cog\Builder $stacking): static
    {
        $stackingResource = $stacking->build();
        $this->internal->stacking = $stackingResource;
    
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
    public function transform(\Grafana\Foundation\Common\GraphTransform $transform): static
    {
        $this->internal->transform = $transform;
    
        return $this;
    }
    /**
     * Indicate if null values should be treated as gaps or connected.
     * When the value is a number, it represents the maximum delta in the
     * X axis that should be considered connected.  For timeseries, this is milliseconds
     * @param bool|float $spanNulls
     */
    public function spanNulls( $spanNulls): static
    {
        $this->internal->spanNulls = $spanNulls;
    
        return $this;
    }
    public function fillBelowTo(string $fillBelowTo): static
    {
        $this->internal->fillBelowTo = $fillBelowTo;
    
        return $this;
    }
    public function pointSymbol(string $pointSymbol): static
    {
        $this->internal->pointSymbol = $pointSymbol;
    
        return $this;
    }
    public function axisCenteredZero(bool $axisCenteredZero): static
    {
        $this->internal->axisCenteredZero = $axisCenteredZero;
    
        return $this;
    }
    public function barMaxWidth(float $barMaxWidth): static
    {
        $this->internal->barMaxWidth = $barMaxWidth;
    
        return $this;
    }

}
