<?php

namespace Grafana\Foundation\Heatmap;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\VizConfigKind>
 */
class VisualizationBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\VizConfigKind $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\VizConfigKind();
    $this->internal->kind = "VizConfig";
    $this->internal->group = "heatmap";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\VizConfigKind
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * The display value for this field.  This supports template variables blank is auto
     */
    public function displayName(string $displayName): static
    {
        $this->internal->spec->fieldConfig->defaults->displayName = $displayName;
    
        return $this;
    }

    /**
     * This can be used by data sources that return and explicit naming structure for values and labels
     * When this property is configured, this value is used rather than the default naming strategy.
     */
    public function displayNameFromDS(string $displayNameFromDS): static
    {
        $this->internal->spec->fieldConfig->defaults->displayNameFromDS = $displayNameFromDS;
    
        return $this;
    }

    /**
     * Human readable field metadata
     */
    public function description(string $description): static
    {
        $this->internal->spec->fieldConfig->defaults->description = $description;
    
        return $this;
    }

    /**
     * An explicit path to the field in the datasource.  When the frame meta includes a path,
     * This will default to `${frame.meta.path}/${field.name}
     * 
     * When defined, this value can be used as an identifier within the datasource scope, and
     * may be used to update the results
     */
    public function path(string $path): static
    {
        $this->internal->spec->fieldConfig->defaults->path = $path;
    
        return $this;
    }

    /**
     * Unit a field should use. The unit you select is applied to all fields except time.
     * You can use the units ID availables in Grafana or a custom unit.
     * Available units in Grafana: https://github.com/grafana/grafana/blob/main/packages/grafana-data/src/valueFormats/categories.ts
     * As custom unit, you can use the following formats:
     * `suffix:<suffix>` for custom unit that should go after value.
     * `prefix:<prefix>` for custom unit that should go before value.
     * `time:<format>` For custom date time formats type for example `time:YYYY-MM-DD`.
     * `si:<base scale><unit characters>` for custom SI units. For example: `si: mF`. This one is a bit more advanced as you can specify both a unit and the source data scale. So if your source data is represented as milli (thousands of) something prefix the unit with that SI scale character.
     * `count:<unit>` for a custom count unit.
     * `currency:<unit>` for custom a currency unit.
     */
    public function unit(string $unit): static
    {
        $this->internal->spec->fieldConfig->defaults->unit = $unit;
    
        return $this;
    }

    /**
     * Specify the number of decimals Grafana includes in the rendered value.
     * If you leave this field blank, Grafana automatically truncates the number of decimals based on the value.
     * For example 1.1234 will display as 1.12 and 100.456 will display as 100.
     * To display all decimals, set the unit to `String`.
     */
    public function decimals(float $decimals): static
    {
        $this->internal->spec->fieldConfig->defaults->decimals = $decimals;
    
        return $this;
    }

    /**
     * The minimum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.
     */
    public function min(float $min): static
    {
        $this->internal->spec->fieldConfig->defaults->min = $min;
    
        return $this;
    }

    /**
     * The maximum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.
     */
    public function max(float $max): static
    {
        $this->internal->spec->fieldConfig->defaults->max = $max;
    
        return $this;
    }

    /**
     * Convert input values into a display string
     * @param array<\Grafana\Foundation\Dashboardv2beta1\ValueMap|\Grafana\Foundation\Dashboardv2beta1\RangeMap|\Grafana\Foundation\Dashboardv2beta1\RegexMap|\Grafana\Foundation\Dashboardv2beta1\SpecialValueMap> $mappings
     */
    public function mappings(array $mappings): static
    {
        $this->internal->spec->fieldConfig->defaults->mappings = $mappings;
    
        return $this;
    }

    /**
     * Map numeric values to states
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\ThresholdsConfig> $thresholds
     */
    public function thresholds(\Grafana\Foundation\Cog\Builder $thresholds): static
    {
        $thresholdsResource = $thresholds->build();
        $this->internal->spec->fieldConfig->defaults->thresholds = $thresholdsResource;
    
        return $this;
    }

    /**
     * Panel color configuration
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\FieldColor> $color
     */
    public function colorScheme(\Grafana\Foundation\Cog\Builder $color): static
    {
        $colorResource = $color->build();
        $this->internal->spec->fieldConfig->defaults->color = $colorResource;
    
        return $this;
    }

    /**
     * The behavior when clicking on a result
     * @param array<mixed> $links
     */
    public function dataLinks(array $links): static
    {
        $this->internal->spec->fieldConfig->defaults->links = $links;
    
        return $this;
    }

    /**
     * Define interactive HTTP requests that can be triggered from data visualizations.
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\Action>> $actions
     */
    public function actions(array $actions): static
    {
            $actionsResources = [];
            foreach ($actions as $r1) {
                    $actionsResources[] = $r1->build();
            }
        $this->internal->spec->fieldConfig->defaults->actions = $actionsResources;
    
        return $this;
    }

    /**
     * Alternative to empty string
     */
    public function noValue(string $noValue): static
    {
        $this->internal->spec->fieldConfig->defaults->noValue = $noValue;
    
        return $this;
    }

    /**
     * Overrides are the options applied to specific fields overriding the defaults.
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1FieldConfigSourceOverrides>> $overrides
     */
    public function overrides(array $overrides): static
    {
            $overridesResources = [];
            foreach ($overrides as $r1) {
                    $overridesResources[] = $r1->build();
            }
        $this->internal->spec->fieldConfig->overrides = $overridesResources;
    
        return $this;
    }

    /**
     * Overrides are the options applied to specific fields overriding the defaults.
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1FieldConfigSourceOverrides> $override
     */
    public function override(\Grafana\Foundation\Cog\Builder $override): static
    {
        $overrideResource = $override->build();
        $this->internal->spec->fieldConfig->overrides[] = $overrideResource;
    
        return $this;
    }

    /**
     * Controls if the heatmap should be calculated from data
     */
    public function calculate(bool $calculate): static
    {    
        if ($this->internal->spec->options === null) {
            $this->internal->spec->options = new \Grafana\Foundation\Heatmap\Options();
        }
        assert($this->internal->spec->options instanceof \Grafana\Foundation\Heatmap\Options);
        $this->internal->spec->options->calculate = $calculate;
    
        return $this;
    }

    /**
     * Calculation options for the heatmap
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\HeatmapCalculationOptions> $calculation
     */
    public function calculation(\Grafana\Foundation\Cog\Builder $calculation): static
    {    
        if ($this->internal->spec->options === null) {
            $this->internal->spec->options = new \Grafana\Foundation\Heatmap\Options();
        }
        assert($this->internal->spec->options instanceof \Grafana\Foundation\Heatmap\Options);
        $calculationResource = $calculation->build();
        $this->internal->spec->options->calculation = $calculationResource;
    
        return $this;
    }

    /**
     * Controls the color options
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Heatmap\HeatmapColorOptions> $color
     */
    public function color(\Grafana\Foundation\Cog\Builder $color): static
    {    
        if ($this->internal->spec->options === null) {
            $this->internal->spec->options = new \Grafana\Foundation\Heatmap\Options();
        }
        assert($this->internal->spec->options instanceof \Grafana\Foundation\Heatmap\Options);
        $colorResource = $color->build();
        $this->internal->spec->options->color = $colorResource;
    
        return $this;
    }

    /**
     * Filters values between a given range
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Heatmap\FilterValueRange> $filterValues
     */
    public function filterValues(\Grafana\Foundation\Cog\Builder $filterValues): static
    {    
        if ($this->internal->spec->options === null) {
            $this->internal->spec->options = new \Grafana\Foundation\Heatmap\Options();
        }
        assert($this->internal->spec->options instanceof \Grafana\Foundation\Heatmap\Options);
        $filterValuesResource = $filterValues->build();
        $this->internal->spec->options->filterValues = $filterValuesResource;
    
        return $this;
    }

    /**
     * Controls tick alignment and value name when not calculating from data
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Heatmap\RowsHeatmapOptions> $rowsFrame
     */
    public function rowsFrame(\Grafana\Foundation\Cog\Builder $rowsFrame): static
    {    
        if ($this->internal->spec->options === null) {
            $this->internal->spec->options = new \Grafana\Foundation\Heatmap\Options();
        }
        assert($this->internal->spec->options instanceof \Grafana\Foundation\Heatmap\Options);
        $rowsFrameResource = $rowsFrame->build();
        $this->internal->spec->options->rowsFrame = $rowsFrameResource;
    
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
        if ($this->internal->spec->options === null) {
            $this->internal->spec->options = new \Grafana\Foundation\Heatmap\Options();
        }
        assert($this->internal->spec->options instanceof \Grafana\Foundation\Heatmap\Options);
        $this->internal->spec->options->showValue = $showValue;
    
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
        if ($this->internal->spec->options === null) {
            $this->internal->spec->options = new \Grafana\Foundation\Heatmap\Options();
        }
        assert($this->internal->spec->options instanceof \Grafana\Foundation\Heatmap\Options);
        $this->internal->spec->options->cellGap = $cellGap;
    
        return $this;
    }

    /**
     * Controls cell radius
     */
    public function cellRadius(float $cellRadius): static
    {    
        if ($this->internal->spec->options === null) {
            $this->internal->spec->options = new \Grafana\Foundation\Heatmap\Options();
        }
        assert($this->internal->spec->options instanceof \Grafana\Foundation\Heatmap\Options);
        $this->internal->spec->options->cellRadius = $cellRadius;
    
        return $this;
    }

    /**
     * Controls cell value unit
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Heatmap\CellValues> $cellValues
     */
    public function cellValues(\Grafana\Foundation\Cog\Builder $cellValues): static
    {    
        if ($this->internal->spec->options === null) {
            $this->internal->spec->options = new \Grafana\Foundation\Heatmap\Options();
        }
        assert($this->internal->spec->options instanceof \Grafana\Foundation\Heatmap\Options);
        $cellValuesResource = $cellValues->build();
        $this->internal->spec->options->cellValues = $cellValuesResource;
    
        return $this;
    }

    /**
     * Controls yAxis placement
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Heatmap\YAxisConfig> $yAxis
     */
    public function yAxis(\Grafana\Foundation\Cog\Builder $yAxis): static
    {    
        if ($this->internal->spec->options === null) {
            $this->internal->spec->options = new \Grafana\Foundation\Heatmap\Options();
        }
        assert($this->internal->spec->options instanceof \Grafana\Foundation\Heatmap\Options);
        $yAxisResource = $yAxis->build();
        $this->internal->spec->options->yAxis = $yAxisResource;
    
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
        if ($this->internal->spec->options === null) {
            $this->internal->spec->options = new \Grafana\Foundation\Heatmap\Options();
        }
        assert($this->internal->spec->options instanceof \Grafana\Foundation\Heatmap\Options);
        $legendResource = $legend->build();
        $this->internal->spec->options->legend = $legendResource;
    
        return $this;
    }

    /**
     * Controls tooltip options
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Heatmap\HeatmapTooltip> $tooltip
     */
    public function tooltip(\Grafana\Foundation\Cog\Builder $tooltip): static
    {    
        if ($this->internal->spec->options === null) {
            $this->internal->spec->options = new \Grafana\Foundation\Heatmap\Options();
        }
        assert($this->internal->spec->options instanceof \Grafana\Foundation\Heatmap\Options);
        $tooltipResource = $tooltip->build();
        $this->internal->spec->options->tooltip = $tooltipResource;
    
        return $this;
    }

    /**
     * Controls exemplar options
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Heatmap\ExemplarConfig> $exemplars
     */
    public function exemplars(\Grafana\Foundation\Cog\Builder $exemplars): static
    {    
        if ($this->internal->spec->options === null) {
            $this->internal->spec->options = new \Grafana\Foundation\Heatmap\Options();
        }
        assert($this->internal->spec->options instanceof \Grafana\Foundation\Heatmap\Options);
        $exemplarsResource = $exemplars->build();
        $this->internal->spec->options->exemplars = $exemplarsResource;
    
        return $this;
    }

    /**
     * Controls which axis to allow selection on
     */
    public function selectionMode(\Grafana\Foundation\Heatmap\HeatmapSelectionMode $selectionMode): static
    {    
        if ($this->internal->spec->options === null) {
            $this->internal->spec->options = new \Grafana\Foundation\Heatmap\Options();
        }
        assert($this->internal->spec->options instanceof \Grafana\Foundation\Heatmap\Options);
        $this->internal->spec->options->selectionMode = $selectionMode;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\ScaleDistributionConfig> $scaleDistribution
     */
    public function scaleDistribution(\Grafana\Foundation\Cog\Builder $scaleDistribution): static
    {    
        if ($this->internal->spec->fieldConfig->defaults->custom === null) {
            $this->internal->spec->fieldConfig->defaults->custom = new \Grafana\Foundation\Heatmap\FieldConfig();
        }
        assert($this->internal->spec->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Heatmap\FieldConfig);
        $scaleDistributionResource = $scaleDistribution->build();
        $this->internal->spec->fieldConfig->defaults->custom->scaleDistribution = $scaleDistributionResource;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\HideSeriesConfig> $hideFrom
     */
    public function hideFrom(\Grafana\Foundation\Cog\Builder $hideFrom): static
    {    
        if ($this->internal->spec->fieldConfig->defaults->custom === null) {
            $this->internal->spec->fieldConfig->defaults->custom = new \Grafana\Foundation\Heatmap\FieldConfig();
        }
        assert($this->internal->spec->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Heatmap\FieldConfig);
        $hideFromResource = $hideFrom->build();
        $this->internal->spec->fieldConfig->defaults->custom->hideFrom = $hideFromResource;
    
        return $this;
    }

    /**
     * Adds override rules for a specific field, referred to by its name.
     * @param array<\Grafana\Foundation\Dashboardv2beta1\DynamicConfigValue> $properties
     */
    public function overrideByName(string $name, array $properties): static
    {
        $this->internal->spec->fieldConfig->overrides[] = new \Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1FieldConfigSourceOverrides(
            matcher: new \Grafana\Foundation\Dashboardv2beta1\MatcherConfig(
            id: "byName",
            options: $name,
        ),
            properties: $properties,
        );
    
        return $this;
    }

    /**
     * Adds override rules for the fields whose name match the given regexp.
     * @param array<\Grafana\Foundation\Dashboardv2beta1\DynamicConfigValue> $properties
     */
    public function overrideByRegexp(string $regexp, array $properties): static
    {
        $this->internal->spec->fieldConfig->overrides[] = new \Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1FieldConfigSourceOverrides(
            matcher: new \Grafana\Foundation\Dashboardv2beta1\MatcherConfig(
            id: "byRegexp",
            options: $regexp,
        ),
            properties: $properties,
        );
    
        return $this;
    }

    /**
     * Adds override rules for all the fields of the given type.
     * @param array<\Grafana\Foundation\Dashboardv2beta1\DynamicConfigValue> $properties
     */
    public function overrideByFieldType(string $fieldType, array $properties): static
    {
        $this->internal->spec->fieldConfig->overrides[] = new \Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1FieldConfigSourceOverrides(
            matcher: new \Grafana\Foundation\Dashboardv2beta1\MatcherConfig(
            id: "byType",
            options: $fieldType,
        ),
            properties: $properties,
        );
    
        return $this;
    }

    /**
     * @param array<\Grafana\Foundation\Dashboardv2beta1\DynamicConfigValue> $properties
     */
    public function overrideByQuery(string $queryRefId, array $properties): static
    {
        $this->internal->spec->fieldConfig->overrides[] = new \Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1FieldConfigSourceOverrides(
            matcher: new \Grafana\Foundation\Dashboardv2beta1\MatcherConfig(
            id: "byFrameRefID",
            options: $queryRefId,
        ),
            properties: $properties,
        );
    
        return $this;
    }

}
