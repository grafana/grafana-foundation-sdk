<?php

namespace Grafana\Foundation\Barchart;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Barchart\FieldConfig>
 */
class FieldConfigBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Barchart\FieldConfig $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Barchart\FieldConfig();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Barchart\FieldConfig
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Controls line width of the bars.
     */
    public function lineWidth(int $lineWidth): static
    {
        if (!($lineWidth <= 10)) {
            throw new \ValueError('$lineWidth must be <= 10');
        }
        $this->internal->lineWidth = $lineWidth;
    
        return $this;
    }

    /**
     * Controls the fill opacity of the bars.
     */
    public function fillOpacity(int $fillOpacity): static
    {
        if (!($fillOpacity <= 100)) {
            throw new \ValueError('$fillOpacity must be <= 100');
        }
        $this->internal->fillOpacity = $fillOpacity;
    
        return $this;
    }

    /**
     * Set the mode of the gradient fill. Fill gradient is based on the line color. To change the color, use the standard color scheme field option.
     * Gradient appearance is influenced by the Fill opacity setting.
     */
    public function gradientMode(\Grafana\Foundation\Common\GraphGradientMode $gradientMode): static
    {
        $this->internal->gradientMode = $gradientMode;
    
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
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\HideSeriesConfig> $hideFrom
     */
    public function hideFrom(\Grafana\Foundation\Cog\Builder $hideFrom): static
    {
        $hideFromResource = $hideFrom->build();
        $this->internal->hideFrom = $hideFromResource;
    
        return $this;
    }

    /**
     * Threshold rendering
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\GraphThresholdsStyleConfig> $thresholdsStyle
     */
    public function thresholdsStyle(\Grafana\Foundation\Cog\Builder $thresholdsStyle): static
    {
        $thresholdsStyleResource = $thresholdsStyle->build();
        $this->internal->thresholdsStyle = $thresholdsStyleResource;
    
        return $this;
    }

    public function axisBorderShow(bool $axisBorderShow): static
    {
        $this->internal->axisBorderShow = $axisBorderShow;
    
        return $this;
    }

}
