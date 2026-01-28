<?php

namespace Grafana\Foundation\Xychart;

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
    $this->internal->group = "xychart";
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
     * Calculate min max per field
     */
    public function fieldMinMax(bool $fieldMinMax): static
    {
        $this->internal->spec->fieldConfig->defaults->fieldMinMax = $fieldMinMax;
    
        return $this;
    }

    /**
     * How null values should be handled when calculating field stats
     * "null" - Include null values, "connected" - Ignore nulls, "null as zero" - Treat nulls as zero
     */
    public function nullValueMode(\Grafana\Foundation\Dashboardv2beta1\NullValueMode $nullValueMode): static
    {
        $this->internal->spec->fieldConfig->defaults->nullValueMode = $nullValueMode;
    
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

    public function show(\Grafana\Foundation\Xychart\XYShowMode $show): static
    {    
        if ($this->internal->spec->fieldConfig->defaults->custom === null) {
            $this->internal->spec->fieldConfig->defaults->custom = new \Grafana\Foundation\Xychart\FieldConfig();
        }
        assert($this->internal->spec->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Xychart\FieldConfig);
        $this->internal->spec->fieldConfig->defaults->custom->show = $show;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Xychart\XychartFieldConfigPointSize> $pointSize
     */
    public function pointSize(\Grafana\Foundation\Cog\Builder $pointSize): static
    {    
        if ($this->internal->spec->fieldConfig->defaults->custom === null) {
            $this->internal->spec->fieldConfig->defaults->custom = new \Grafana\Foundation\Xychart\FieldConfig();
        }
        assert($this->internal->spec->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Xychart\FieldConfig);
        $pointSizeResource = $pointSize->build();
        $this->internal->spec->fieldConfig->defaults->custom->pointSize = $pointSizeResource;
    
        return $this;
    }

    public function pointShape(\Grafana\Foundation\Xychart\PointShape $pointShape): static
    {    
        if ($this->internal->spec->fieldConfig->defaults->custom === null) {
            $this->internal->spec->fieldConfig->defaults->custom = new \Grafana\Foundation\Xychart\FieldConfig();
        }
        assert($this->internal->spec->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Xychart\FieldConfig);
        $this->internal->spec->fieldConfig->defaults->custom->pointShape = $pointShape;
    
        return $this;
    }

    public function pointStrokeWidth(int $pointStrokeWidth): static
    {
        if (!($pointStrokeWidth >= 0)) {
            throw new \ValueError('$pointStrokeWidth must be >= 0');
        }    
        if ($this->internal->spec->fieldConfig->defaults->custom === null) {
            $this->internal->spec->fieldConfig->defaults->custom = new \Grafana\Foundation\Xychart\FieldConfig();
        }
        assert($this->internal->spec->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Xychart\FieldConfig);
        $this->internal->spec->fieldConfig->defaults->custom->pointStrokeWidth = $pointStrokeWidth;
    
        return $this;
    }

    public function fillOpacity(int $fillOpacity): static
    {
        if (!($fillOpacity <= 100)) {
            throw new \ValueError('$fillOpacity must be <= 100');
        }    
        if ($this->internal->spec->fieldConfig->defaults->custom === null) {
            $this->internal->spec->fieldConfig->defaults->custom = new \Grafana\Foundation\Xychart\FieldConfig();
        }
        assert($this->internal->spec->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Xychart\FieldConfig);
        $this->internal->spec->fieldConfig->defaults->custom->fillOpacity = $fillOpacity;
    
        return $this;
    }

    public function lineWidth(int $lineWidth): static
    {
        if (!($lineWidth >= 0)) {
            throw new \ValueError('$lineWidth must be >= 0');
        }    
        if ($this->internal->spec->fieldConfig->defaults->custom === null) {
            $this->internal->spec->fieldConfig->defaults->custom = new \Grafana\Foundation\Xychart\FieldConfig();
        }
        assert($this->internal->spec->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Xychart\FieldConfig);
        $this->internal->spec->fieldConfig->defaults->custom->lineWidth = $lineWidth;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\HideSeriesConfig> $hideFrom
     */
    public function hideFrom(\Grafana\Foundation\Cog\Builder $hideFrom): static
    {    
        if ($this->internal->spec->fieldConfig->defaults->custom === null) {
            $this->internal->spec->fieldConfig->defaults->custom = new \Grafana\Foundation\Xychart\FieldConfig();
        }
        assert($this->internal->spec->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Xychart\FieldConfig);
        $hideFromResource = $hideFrom->build();
        $this->internal->spec->fieldConfig->defaults->custom->hideFrom = $hideFromResource;
    
        return $this;
    }

    public function axisPlacement(\Grafana\Foundation\Common\AxisPlacement $axisPlacement): static
    {    
        if ($this->internal->spec->fieldConfig->defaults->custom === null) {
            $this->internal->spec->fieldConfig->defaults->custom = new \Grafana\Foundation\Xychart\FieldConfig();
        }
        assert($this->internal->spec->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Xychart\FieldConfig);
        $this->internal->spec->fieldConfig->defaults->custom->axisPlacement = $axisPlacement;
    
        return $this;
    }

    public function axisColorMode(\Grafana\Foundation\Common\AxisColorMode $axisColorMode): static
    {    
        if ($this->internal->spec->fieldConfig->defaults->custom === null) {
            $this->internal->spec->fieldConfig->defaults->custom = new \Grafana\Foundation\Xychart\FieldConfig();
        }
        assert($this->internal->spec->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Xychart\FieldConfig);
        $this->internal->spec->fieldConfig->defaults->custom->axisColorMode = $axisColorMode;
    
        return $this;
    }

    public function axisLabel(string $axisLabel): static
    {    
        if ($this->internal->spec->fieldConfig->defaults->custom === null) {
            $this->internal->spec->fieldConfig->defaults->custom = new \Grafana\Foundation\Xychart\FieldConfig();
        }
        assert($this->internal->spec->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Xychart\FieldConfig);
        $this->internal->spec->fieldConfig->defaults->custom->axisLabel = $axisLabel;
    
        return $this;
    }

    public function axisWidth(float $axisWidth): static
    {    
        if ($this->internal->spec->fieldConfig->defaults->custom === null) {
            $this->internal->spec->fieldConfig->defaults->custom = new \Grafana\Foundation\Xychart\FieldConfig();
        }
        assert($this->internal->spec->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Xychart\FieldConfig);
        $this->internal->spec->fieldConfig->defaults->custom->axisWidth = $axisWidth;
    
        return $this;
    }

    public function axisSoftMin(float $axisSoftMin): static
    {    
        if ($this->internal->spec->fieldConfig->defaults->custom === null) {
            $this->internal->spec->fieldConfig->defaults->custom = new \Grafana\Foundation\Xychart\FieldConfig();
        }
        assert($this->internal->spec->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Xychart\FieldConfig);
        $this->internal->spec->fieldConfig->defaults->custom->axisSoftMin = $axisSoftMin;
    
        return $this;
    }

    public function axisSoftMax(float $axisSoftMax): static
    {    
        if ($this->internal->spec->fieldConfig->defaults->custom === null) {
            $this->internal->spec->fieldConfig->defaults->custom = new \Grafana\Foundation\Xychart\FieldConfig();
        }
        assert($this->internal->spec->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Xychart\FieldConfig);
        $this->internal->spec->fieldConfig->defaults->custom->axisSoftMax = $axisSoftMax;
    
        return $this;
    }

    public function axisGridShow(bool $axisGridShow): static
    {    
        if ($this->internal->spec->fieldConfig->defaults->custom === null) {
            $this->internal->spec->fieldConfig->defaults->custom = new \Grafana\Foundation\Xychart\FieldConfig();
        }
        assert($this->internal->spec->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Xychart\FieldConfig);
        $this->internal->spec->fieldConfig->defaults->custom->axisGridShow = $axisGridShow;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\ScaleDistributionConfig> $scaleDistribution
     */
    public function scaleDistribution(\Grafana\Foundation\Cog\Builder $scaleDistribution): static
    {    
        if ($this->internal->spec->fieldConfig->defaults->custom === null) {
            $this->internal->spec->fieldConfig->defaults->custom = new \Grafana\Foundation\Xychart\FieldConfig();
        }
        assert($this->internal->spec->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Xychart\FieldConfig);
        $scaleDistributionResource = $scaleDistribution->build();
        $this->internal->spec->fieldConfig->defaults->custom->scaleDistribution = $scaleDistributionResource;
    
        return $this;
    }

    public function axisCenteredZero(bool $axisCenteredZero): static
    {    
        if ($this->internal->spec->fieldConfig->defaults->custom === null) {
            $this->internal->spec->fieldConfig->defaults->custom = new \Grafana\Foundation\Xychart\FieldConfig();
        }
        assert($this->internal->spec->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Xychart\FieldConfig);
        $this->internal->spec->fieldConfig->defaults->custom->axisCenteredZero = $axisCenteredZero;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\LineStyle> $lineStyle
     */
    public function lineStyle(\Grafana\Foundation\Cog\Builder $lineStyle): static
    {    
        if ($this->internal->spec->fieldConfig->defaults->custom === null) {
            $this->internal->spec->fieldConfig->defaults->custom = new \Grafana\Foundation\Xychart\FieldConfig();
        }
        assert($this->internal->spec->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Xychart\FieldConfig);
        $lineStyleResource = $lineStyle->build();
        $this->internal->spec->fieldConfig->defaults->custom->lineStyle = $lineStyleResource;
    
        return $this;
    }

    public function axisBorderShow(bool $axisBorderShow): static
    {    
        if ($this->internal->spec->fieldConfig->defaults->custom === null) {
            $this->internal->spec->fieldConfig->defaults->custom = new \Grafana\Foundation\Xychart\FieldConfig();
        }
        assert($this->internal->spec->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Xychart\FieldConfig);
        $this->internal->spec->fieldConfig->defaults->custom->axisBorderShow = $axisBorderShow;
    
        return $this;
    }

    public function mapping(\Grafana\Foundation\Xychart\SeriesMapping $mapping): static
    {    
        if ($this->internal->spec->options === null) {
            $this->internal->spec->options = new \Grafana\Foundation\Xychart\Options();
        }
        assert($this->internal->spec->options instanceof \Grafana\Foundation\Xychart\Options);
        $this->internal->spec->options->mapping = $mapping;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\VizLegendOptions> $legend
     */
    public function legend(\Grafana\Foundation\Cog\Builder $legend): static
    {    
        if ($this->internal->spec->options === null) {
            $this->internal->spec->options = new \Grafana\Foundation\Xychart\Options();
        }
        assert($this->internal->spec->options instanceof \Grafana\Foundation\Xychart\Options);
        $legendResource = $legend->build();
        $this->internal->spec->options->legend = $legendResource;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\VizTooltipOptions> $tooltip
     */
    public function tooltip(\Grafana\Foundation\Cog\Builder $tooltip): static
    {    
        if ($this->internal->spec->options === null) {
            $this->internal->spec->options = new \Grafana\Foundation\Xychart\Options();
        }
        assert($this->internal->spec->options instanceof \Grafana\Foundation\Xychart\Options);
        $tooltipResource = $tooltip->build();
        $this->internal->spec->options->tooltip = $tooltipResource;
    
        return $this;
    }

    /**
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Xychart\XYSeriesConfig>> $series
     */
    public function series(array $series): static
    {    
        if ($this->internal->spec->options === null) {
            $this->internal->spec->options = new \Grafana\Foundation\Xychart\Options();
        }
        assert($this->internal->spec->options instanceof \Grafana\Foundation\Xychart\Options);
            $seriesResources = [];
            foreach ($series as $r1) {
                    $seriesResources[] = $r1->build();
            }
        $this->internal->spec->options->series = $seriesResources;
    
        return $this;
    }

}
