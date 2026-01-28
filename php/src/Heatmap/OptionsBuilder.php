<?php

namespace Grafana\Foundation\Heatmap;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Heatmap\Options>
 */
class OptionsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Heatmap\Options $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Heatmap\Options();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Heatmap\Options
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Controls if the heatmap should be calculated from data
     */
    public function calculate(bool $calculate): static
    {
        $this->internal->calculate = $calculate;
    
        return $this;
    }

    /**
     * Calculation options for the heatmap
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\HeatmapCalculationOptions> $calculation
     */
    public function calculation(\Grafana\Foundation\Cog\Builder $calculation): static
    {
        $calculationResource = $calculation->build();
        $this->internal->calculation = $calculationResource;
    
        return $this;
    }

    /**
     * Controls the color options
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Heatmap\HeatmapColorOptions> $color
     */
    public function color(\Grafana\Foundation\Cog\Builder $color): static
    {
        $colorResource = $color->build();
        $this->internal->color = $colorResource;
    
        return $this;
    }

    /**
     * Filters values between a given range
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Heatmap\FilterValueRange> $filterValues
     */
    public function filterValues(\Grafana\Foundation\Cog\Builder $filterValues): static
    {
        $filterValuesResource = $filterValues->build();
        $this->internal->filterValues = $filterValuesResource;
    
        return $this;
    }

    /**
     * Controls tick alignment and value name when not calculating from data
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Heatmap\RowsHeatmapOptions> $rowsFrame
     */
    public function rowsFrame(\Grafana\Foundation\Cog\Builder $rowsFrame): static
    {
        $rowsFrameResource = $rowsFrame->build();
        $this->internal->rowsFrame = $rowsFrameResource;
    
        return $this;
    }

    /**
     * | *{
     * 	layout: ui.HeatmapCellLayout & "auto" // TODO: fix after remove when https://github.com/grafana/cuetsy/issues/74 is fixed
     * }
     * Controls the display of the value in the cell
     */
    public function showValue(\Grafana\Foundation\Common\VisibilityMode $showValue): static
    {
        $this->internal->showValue = $showValue;
    
        return $this;
    }

    /**
     * Controls gap between cells
     */
    public function cellGap(int $cellGap): static
    {
        if (!($cellGap <= 25)) {
            throw new \ValueError('$cellGap must be <= 25');
        }
        $this->internal->cellGap = $cellGap;
    
        return $this;
    }

    /**
     * Controls cell radius
     */
    public function cellRadius(float $cellRadius): static
    {
        $this->internal->cellRadius = $cellRadius;
    
        return $this;
    }

    /**
     * Controls cell value unit
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Heatmap\CellValues> $cellValues
     */
    public function cellValues(\Grafana\Foundation\Cog\Builder $cellValues): static
    {
        $cellValuesResource = $cellValues->build();
        $this->internal->cellValues = $cellValuesResource;
    
        return $this;
    }

    /**
     * Controls yAxis placement
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Heatmap\YAxisConfig> $yAxis
     */
    public function yAxis(\Grafana\Foundation\Cog\Builder $yAxis): static
    {
        $yAxisResource = $yAxis->build();
        $this->internal->yAxis = $yAxisResource;
    
        return $this;
    }

    /**
     * | *{
     * 	axisPlacement: ui.AxisPlacement & "left" // TODO: fix after remove when https://github.com/grafana/cuetsy/issues/74 is fixed
     * }
     * Controls legend options
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Heatmap\HeatmapLegend> $legend
     */
    public function legend(\Grafana\Foundation\Cog\Builder $legend): static
    {
        $legendResource = $legend->build();
        $this->internal->legend = $legendResource;
    
        return $this;
    }

    /**
     * Controls tooltip options
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Heatmap\HeatmapTooltip> $tooltip
     */
    public function tooltip(\Grafana\Foundation\Cog\Builder $tooltip): static
    {
        $tooltipResource = $tooltip->build();
        $this->internal->tooltip = $tooltipResource;
    
        return $this;
    }

    /**
     * Controls exemplar options
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Heatmap\ExemplarConfig> $exemplars
     */
    public function exemplars(\Grafana\Foundation\Cog\Builder $exemplars): static
    {
        $exemplarsResource = $exemplars->build();
        $this->internal->exemplars = $exemplarsResource;
    
        return $this;
    }

    /**
     * Controls which axis to allow selection on
     */
    public function selectionMode(\Grafana\Foundation\Heatmap\HeatmapSelectionMode $selectionMode): static
    {
        $this->internal->selectionMode = $selectionMode;
    
        return $this;
    }

}
