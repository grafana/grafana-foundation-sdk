<?php

namespace Grafana\Foundation\Piechart;

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
    $this->internal->group = "piechart";
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

    public function pieType(\Grafana\Foundation\Piechart\PieChartType $pieType): static
    {    
        if ($this->internal->spec->options === null) {
            $this->internal->spec->options = new \Grafana\Foundation\Piechart\Options();
        }
        assert($this->internal->spec->options instanceof \Grafana\Foundation\Piechart\Options);
        $this->internal->spec->options->pieType = $pieType;
    
        return $this;
    }

    /**
     * @param array<\Grafana\Foundation\Piechart\PieChartLabels> $displayLabels
     */
    public function displayLabels(array $displayLabels): static
    {    
        if ($this->internal->spec->options === null) {
            $this->internal->spec->options = new \Grafana\Foundation\Piechart\Options();
        }
        assert($this->internal->spec->options instanceof \Grafana\Foundation\Piechart\Options);
        $this->internal->spec->options->displayLabels = $displayLabels;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\VizTooltipOptions> $tooltip
     */
    public function tooltip(\Grafana\Foundation\Cog\Builder $tooltip): static
    {    
        if ($this->internal->spec->options === null) {
            $this->internal->spec->options = new \Grafana\Foundation\Piechart\Options();
        }
        assert($this->internal->spec->options instanceof \Grafana\Foundation\Piechart\Options);
        $tooltipResource = $tooltip->build();
        $this->internal->spec->options->tooltip = $tooltipResource;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\ReduceDataOptions> $reduceOptions
     */
    public function reduceOptions(\Grafana\Foundation\Cog\Builder $reduceOptions): static
    {    
        if ($this->internal->spec->options === null) {
            $this->internal->spec->options = new \Grafana\Foundation\Piechart\Options();
        }
        assert($this->internal->spec->options instanceof \Grafana\Foundation\Piechart\Options);
        $reduceOptionsResource = $reduceOptions->build();
        $this->internal->spec->options->reduceOptions = $reduceOptionsResource;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\VizTextDisplayOptions> $text
     */
    public function text(\Grafana\Foundation\Cog\Builder $text): static
    {    
        if ($this->internal->spec->options === null) {
            $this->internal->spec->options = new \Grafana\Foundation\Piechart\Options();
        }
        assert($this->internal->spec->options instanceof \Grafana\Foundation\Piechart\Options);
        $textResource = $text->build();
        $this->internal->spec->options->text = $textResource;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Piechart\PieChartLegendOptions> $legend
     */
    public function legend(\Grafana\Foundation\Cog\Builder $legend): static
    {    
        if ($this->internal->spec->options === null) {
            $this->internal->spec->options = new \Grafana\Foundation\Piechart\Options();
        }
        assert($this->internal->spec->options instanceof \Grafana\Foundation\Piechart\Options);
        $legendResource = $legend->build();
        $this->internal->spec->options->legend = $legendResource;
    
        return $this;
    }

    public function orientation(\Grafana\Foundation\Common\VizOrientation $orientation): static
    {    
        if ($this->internal->spec->options === null) {
            $this->internal->spec->options = new \Grafana\Foundation\Piechart\Options();
        }
        assert($this->internal->spec->options instanceof \Grafana\Foundation\Piechart\Options);
        $this->internal->spec->options->orientation = $orientation;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\HideSeriesConfig> $hideFrom
     */
    public function hideFrom(\Grafana\Foundation\Cog\Builder $hideFrom): static
    {    
        if ($this->internal->spec->fieldConfig->defaults->custom === null) {
            $this->internal->spec->fieldConfig->defaults->custom = new \Grafana\Foundation\Piechart\FieldConfig();
        }
        assert($this->internal->spec->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Piechart\FieldConfig);
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
