<?php

namespace Grafana\Foundation\Barchart;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Barchart\Options>
 */
class OptionsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Barchart\Options $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Barchart\Options();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Barchart\Options
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Manually select which field from the dataset to represent the x field.
     */
    public function xField(string $xField): static
    {
        $this->internal->xField = $xField;
    
        return $this;
    }

    /**
     * Use the color value for a sibling field to color each bar value.
     */
    public function colorByField(string $colorByField): static
    {
        $this->internal->colorByField = $colorByField;
    
        return $this;
    }

    /**
     * Controls the orientation of the bar chart, either vertical or horizontal.
     */
    public function orientation(\Grafana\Foundation\Common\VizOrientation $orientation): static
    {
        $this->internal->orientation = $orientation;
    
        return $this;
    }

    /**
     * Controls the radius of each bar.
     */
    public function barRadius(float $barRadius): static
    {
        if (!($barRadius >= 0)) {
            throw new \ValueError('$barRadius must be >= 0');
        }
        if (!($barRadius <= 0.5)) {
            throw new \ValueError('$barRadius must be <= 0.5');
        }
        $this->internal->barRadius = $barRadius;
    
        return $this;
    }

    /**
     * Controls the rotation of the x axis labels.
     */
    public function xTickLabelRotation(int $xTickLabelRotation): static
    {
        if (!($xTickLabelRotation >= -90)) {
            throw new \ValueError('$xTickLabelRotation must be >= -90');
        }
        if (!($xTickLabelRotation <= 90)) {
            throw new \ValueError('$xTickLabelRotation must be <= 90');
        }
        $this->internal->xTickLabelRotation = $xTickLabelRotation;
    
        return $this;
    }

    /**
     * Sets the max length that a label can have before it is truncated.
     */
    public function xTickLabelMaxLength(int $xTickLabelMaxLength): static
    {
        if (!($xTickLabelMaxLength >= 0)) {
            throw new \ValueError('$xTickLabelMaxLength must be >= 0');
        }
        $this->internal->xTickLabelMaxLength = $xTickLabelMaxLength;
    
        return $this;
    }

    /**
     * Controls the spacing between x axis labels.
     * negative values indicate backwards skipping behavior
     */
    public function xTickLabelSpacing(int $xTickLabelSpacing): static
    {
        $this->internal->xTickLabelSpacing = $xTickLabelSpacing;
    
        return $this;
    }

    /**
     * Controls whether bars are stacked or not, either normally or in percent mode.
     */
    public function stacking(\Grafana\Foundation\Common\StackingMode $stacking): static
    {
        $this->internal->stacking = $stacking;
    
        return $this;
    }

    /**
     * This controls whether values are shown on top or to the left of bars.
     */
    public function showValue(\Grafana\Foundation\Common\VisibilityMode $showValue): static
    {
        $this->internal->showValue = $showValue;
    
        return $this;
    }

    /**
     * Controls the width of bars. 1 = Max width, 0 = Min width.
     */
    public function barWidth(float $barWidth): static
    {
        if (!($barWidth >= 0)) {
            throw new \ValueError('$barWidth must be >= 0');
        }
        if (!($barWidth <= 1)) {
            throw new \ValueError('$barWidth must be <= 1');
        }
        $this->internal->barWidth = $barWidth;
    
        return $this;
    }

    /**
     * Controls the width of groups. 1 = max with, 0 = min width.
     */
    public function groupWidth(float $groupWidth): static
    {
        if (!($groupWidth >= 0)) {
            throw new \ValueError('$groupWidth must be >= 0');
        }
        if (!($groupWidth <= 1)) {
            throw new \ValueError('$groupWidth must be <= 1');
        }
        $this->internal->groupWidth = $groupWidth;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\VizLegendOptions> $legend
     */
    public function legend(\Grafana\Foundation\Cog\Builder $legend): static
    {
        $legendResource = $legend->build();
        $this->internal->legend = $legendResource;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\VizTooltipOptions> $tooltip
     */
    public function tooltip(\Grafana\Foundation\Cog\Builder $tooltip): static
    {
        $tooltipResource = $tooltip->build();
        $this->internal->tooltip = $tooltipResource;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\VizTextDisplayOptions> $text
     */
    public function text(\Grafana\Foundation\Cog\Builder $text): static
    {
        $textResource = $text->build();
        $this->internal->text = $textResource;
    
        return $this;
    }

    /**
     * Enables mode which highlights the entire bar area and shows tooltip when cursor
     * hovers over highlighted area
     */
    public function fullHighlight(bool $fullHighlight): static
    {
        $this->internal->fullHighlight = $fullHighlight;
    
        return $this;
    }

}
