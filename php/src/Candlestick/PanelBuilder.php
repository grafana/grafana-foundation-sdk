<?php

namespace Grafana\Foundation\Candlestick;

/**
 * Dashboard panels are the basic visualization building blocks.
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\Panel>
 */
class PanelBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboard\Panel $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboard\Panel();
    $this->internal->type = "candlestick";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboard\Panel
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Unique identifier of the panel. Generated by Grafana when creating a new panel. It must be unique within a dashboard, but not globally.
     */
    public function id(int $id): static
    {
        $this->internal->id = $id;
    
        return $this;
    }
    /**
     * Depends on the panel plugin. See the plugin documentation for details.
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Cog\Dataquery>> $targets
     */
    public function targets(array $targets): static
    {
            $targetsResources = [];
            foreach ($targets as $r1) {
                    $targetsResources[] = $r1->build();
            }
        $this->internal->targets = $targetsResources;
    
        return $this;
    }
    /**
     * Depends on the panel plugin. See the plugin documentation for details.
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Cog\Dataquery> $target
     */
    public function withTarget(\Grafana\Foundation\Cog\Builder $target): static
    {    
        if ($this->internal->targets === null) {
            $this->internal->targets = [];
        }
        
        $targetResource = $target->build();
        $this->internal->targets[] = $targetResource;
    
        return $this;
    }
    /**
     * Panel title.
     */
    public function title(string $title): static
    {
        $this->internal->title = $title;
    
        return $this;
    }
    /**
     * Panel description.
     */
    public function description(string $description): static
    {
        $this->internal->description = $description;
    
        return $this;
    }
    /**
     * Whether to display the panel without a background.
     */
    public function transparent(bool $transparent): static
    {
        $this->internal->transparent = $transparent;
    
        return $this;
    }
    /**
     * The datasource used in all targets.
     */
    public function datasource(\Grafana\Foundation\Dashboard\DataSourceRef $datasource): static
    {
        $this->internal->datasource = $datasource;
    
        return $this;
    }
    /**
     * Grid position.
     */
    public function gridPos(\Grafana\Foundation\Dashboard\GridPos $gridPos): static
    {
        $this->internal->gridPos = $gridPos;
    
        return $this;
    }
    /**
     * Panel height. The height is the number of rows from the top edge of the panel.
     */
    public function height(int $h): static
    {
        if (!($h > 0)) {
            throw new \ValueError('$h must be > 0');
        }    
        if ($this->internal->gridPos === null) {
            $this->internal->gridPos = new \Grafana\Foundation\Dashboard\GridPos();
        }
        assert($this->internal->gridPos instanceof \Grafana\Foundation\Dashboard\GridPos);
        $this->internal->gridPos->h = $h;
    
        return $this;
    }
    /**
     * Panel width. The width is the number of columns from the left edge of the panel.
     */
    public function span(int $w): static
    {
        if (!($w > 0)) {
            throw new \ValueError('$w must be > 0');
        }
        if (!($w <= 24)) {
            throw new \ValueError('$w must be <= 24');
        }    
        if ($this->internal->gridPos === null) {
            $this->internal->gridPos = new \Grafana\Foundation\Dashboard\GridPos();
        }
        assert($this->internal->gridPos instanceof \Grafana\Foundation\Dashboard\GridPos);
        $this->internal->gridPos->w = $w;
    
        return $this;
    }
    /**
     * Panel links.
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\DashboardLink>> $links
     */
    public function links(array $links): static
    {
            $linksResources = [];
            foreach ($links as $r1) {
                    $linksResources[] = $r1->build();
            }
        $this->internal->links = $linksResources;
    
        return $this;
    }
    /**
     * Name of template variable to repeat for.
     */
    public function repeat(string $repeat): static
    {
        $this->internal->repeat = $repeat;
    
        return $this;
    }
    /**
     * Direction to repeat in if 'repeat' is set.
     * `h` for horizontal, `v` for vertical.
     */
    public function repeatDirection(\Grafana\Foundation\Dashboard\PanelRepeatDirection $repeatDirection): static
    {
        $this->internal->repeatDirection = $repeatDirection;
    
        return $this;
    }
    /**
     * The maximum number of data points that the panel queries are retrieving.
     */
    public function maxDataPoints(float $maxDataPoints): static
    {
        $this->internal->maxDataPoints = $maxDataPoints;
    
        return $this;
    }
    /**
     * List of transformations that are applied to the panel data before rendering.
     * When there are multiple transformations, Grafana applies them in the order they are listed.
     * Each transformation creates a result set that then passes on to the next transformation in the processing pipeline.
     * @param array<\Grafana\Foundation\Dashboard\DataTransformerConfig> $transformations
     */
    public function transformations(array $transformations): static
    {
        $this->internal->transformations = $transformations;
    
        return $this;
    }
    /**
     * List of transformations that are applied to the panel data before rendering.
     * When there are multiple transformations, Grafana applies them in the order they are listed.
     * Each transformation creates a result set that then passes on to the next transformation in the processing pipeline.
     */
    public function withTransformation(\Grafana\Foundation\Dashboard\DataTransformerConfig $transformation): static
    {
        $this->internal->transformations[] = $transformation;
    
        return $this;
    }
    /**
     * The min time interval setting defines a lower limit for the $__interval and $__interval_ms variables.
     * This value must be formatted as a number followed by a valid time
     * identifier like: "40s", "3d", etc.
     * See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options
     */
    public function interval(string $interval): static
    {
        $this->internal->interval = $interval;
    
        return $this;
    }
    /**
     * Overrides the relative time range for individual panels,
     * which causes them to be different than what is selected in
     * the dashboard time picker in the top-right corner of the dashboard. You can use this to show metrics from different
     * time periods or days on the same dashboard.
     * The value is formatted as time operation like: `now-5m` (Last 5 minutes), `now/d` (the day so far),
     * `now-5d/d`(Last 5 days), `now/w` (This week so far), `now-2y/y` (Last 2 years).
     * Note: Panel time overrides have no effect when the dashboard’s time range is absolute.
     * See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options
     */
    public function timeFrom(string $timeFrom): static
    {
        $this->internal->timeFrom = $timeFrom;
    
        return $this;
    }
    /**
     * Overrides the time range for individual panels by shifting its start and end relative to the time picker.
     * For example, you can shift the time range for the panel to be two hours earlier than the dashboard time picker setting `2h`.
     * Note: Panel time overrides have no effect when the dashboard’s time range is absolute.
     * See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options
     */
    public function timeShift(string $timeShift): static
    {
        $this->internal->timeShift = $timeShift;
    
        return $this;
    }
    /**
     * Dynamically load the panel
     */
    public function libraryPanel(\Grafana\Foundation\Dashboard\LibraryPanelRef $libraryPanel): static
    {
        $this->internal->libraryPanel = $libraryPanel;
    
        return $this;
    }
    /**
     * The display value for this field.  This supports template variables blank is auto
     */
    public function displayName(string $displayName): static
    {
        $this->internal->fieldConfig->defaults->displayName = $displayName;
    
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
        $this->internal->fieldConfig->defaults->unit = $unit;
    
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
        $this->internal->fieldConfig->defaults->decimals = $decimals;
    
        return $this;
    }
    /**
     * The minimum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.
     */
    public function min(float $min): static
    {
        $this->internal->fieldConfig->defaults->min = $min;
    
        return $this;
    }
    /**
     * The maximum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.
     */
    public function max(float $max): static
    {
        $this->internal->fieldConfig->defaults->max = $max;
    
        return $this;
    }
    /**
     * Convert input values into a display string
     * @param array<\Grafana\Foundation\Dashboard\ValueMap|\Grafana\Foundation\Dashboard\RangeMap|\Grafana\Foundation\Dashboard\RegexMap|\Grafana\Foundation\Dashboard\SpecialValueMap> $mappings
     */
    public function mappings(array $mappings): static
    {
        $this->internal->fieldConfig->defaults->mappings = $mappings;
    
        return $this;
    }
    /**
     * Map numeric values to states
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\ThresholdsConfig> $thresholds
     */
    public function thresholds(\Grafana\Foundation\Cog\Builder $thresholds): static
    {
        $thresholdsResource = $thresholds->build();
        $this->internal->fieldConfig->defaults->thresholds = $thresholdsResource;
    
        return $this;
    }
    /**
     * Panel color configuration
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\FieldColor> $color
     */
    public function colorScheme(\Grafana\Foundation\Cog\Builder $color): static
    {
        $colorResource = $color->build();
        $this->internal->fieldConfig->defaults->color = $colorResource;
    
        return $this;
    }
    /**
     * The behavior when clicking on a result
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\DashboardLink>> $links
     */
    public function dataLinks(array $links): static
    {
            $linksResources = [];
            foreach ($links as $r1) {
                    $linksResources[] = $r1->build();
            }
        $this->internal->fieldConfig->defaults->links = $linksResources;
    
        return $this;
    }
    /**
     * Alternative to empty string
     */
    public function noValue(string $noValue): static
    {
        $this->internal->fieldConfig->defaults->noValue = $noValue;
    
        return $this;
    }
    /**
     * Overrides are the options applied to specific fields overriding the defaults.
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\DashboardFieldConfigSourceOverrides>> $overrides
     */
    public function overrides(array $overrides): static
    {
            $overridesResources = [];
            foreach ($overrides as $r1) {
                    $overridesResources[] = $r1->build();
            }
        $this->internal->fieldConfig->overrides = $overridesResources;
    
        return $this;
    }
    /**
     * Overrides are the options applied to specific fields overriding the defaults.
     * @param array<\Grafana\Foundation\Dashboard\DynamicConfigValue> $properties
     */
    public function withOverride(\Grafana\Foundation\Dashboard\MatcherConfig $matcher,array $properties): static
    {
        $this->internal->fieldConfig->overrides[] = new \Grafana\Foundation\Dashboard\DashboardFieldConfigSourceOverrides(
            matcher: $matcher,
            properties: $properties,
        );
    
        return $this;
    }
    /**
     * Adds override rules for a specific field, referred to by its name.
     * @param array<\Grafana\Foundation\Dashboard\DynamicConfigValue> $properties
     */
    public function overrideByName(string $name,array $properties): static
    {
        $this->internal->fieldConfig->overrides[] = new \Grafana\Foundation\Dashboard\DashboardFieldConfigSourceOverrides(
            matcher: new \Grafana\Foundation\Dashboard\MatcherConfig(
            id: "byName",
            options: $name,
        ),
            properties: $properties,
        );
    
        return $this;
    }
    /**
     * Adds override rules for the fields whose name match the given regexp.
     * @param array<\Grafana\Foundation\Dashboard\DynamicConfigValue> $properties
     */
    public function overrideByRegexp(string $regexp,array $properties): static
    {
        $this->internal->fieldConfig->overrides[] = new \Grafana\Foundation\Dashboard\DashboardFieldConfigSourceOverrides(
            matcher: new \Grafana\Foundation\Dashboard\MatcherConfig(
            id: "byRegexp",
            options: $regexp,
        ),
            properties: $properties,
        );
    
        return $this;
    }
    /**
     * Adds override rules for all the fields of the given type.
     * @param array<\Grafana\Foundation\Dashboard\DynamicConfigValue> $properties
     */
    public function overrideByFieldType(string $fieldType,array $properties): static
    {
        $this->internal->fieldConfig->overrides[] = new \Grafana\Foundation\Dashboard\DashboardFieldConfigSourceOverrides(
            matcher: new \Grafana\Foundation\Dashboard\MatcherConfig(
            id: "byType",
            options: $fieldType,
        ),
            properties: $properties,
        );
    
        return $this;
    }
    /**
     * @param array<\Grafana\Foundation\Dashboard\DynamicConfigValue> $properties
     */
    public function overrideByQuery(string $queryRefId,array $properties): static
    {
        $this->internal->fieldConfig->overrides[] = new \Grafana\Foundation\Dashboard\DashboardFieldConfigSourceOverrides(
            matcher: new \Grafana\Foundation\Dashboard\MatcherConfig(
            id: "byFrameRefID",
            options: $queryRefId,
        ),
            properties: $properties,
        );
    
        return $this;
    }
    /**
     * Sets which dimensions are used for the visualization
     */
    public function mode(\Grafana\Foundation\Candlestick\VizDisplayMode $mode): static
    {    
        if ($this->internal->options === null) {
            $this->internal->options = new \Grafana\Foundation\Candlestick\Options();
        }
        assert($this->internal->options instanceof \Grafana\Foundation\Candlestick\Options);
        $this->internal->options->mode = $mode;
    
        return $this;
    }
    /**
     * Sets the style of the candlesticks
     */
    public function candleStyle(\Grafana\Foundation\Candlestick\CandleStyle $candleStyle): static
    {    
        if ($this->internal->options === null) {
            $this->internal->options = new \Grafana\Foundation\Candlestick\Options();
        }
        assert($this->internal->options instanceof \Grafana\Foundation\Candlestick\Options);
        $this->internal->options->candleStyle = $candleStyle;
    
        return $this;
    }
    /**
     * Sets the color strategy for the candlesticks
     */
    public function colorStrategy(\Grafana\Foundation\Candlestick\ColorStrategy $colorStrategy): static
    {    
        if ($this->internal->options === null) {
            $this->internal->options = new \Grafana\Foundation\Candlestick\Options();
        }
        assert($this->internal->options instanceof \Grafana\Foundation\Candlestick\Options);
        $this->internal->options->colorStrategy = $colorStrategy;
    
        return $this;
    }
    /**
     * Map fields to appropriate dimension
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Candlestick\CandlestickFieldMap> $fields
     */
    public function fields(\Grafana\Foundation\Cog\Builder $fields): static
    {    
        if ($this->internal->options === null) {
            $this->internal->options = new \Grafana\Foundation\Candlestick\Options();
        }
        assert($this->internal->options instanceof \Grafana\Foundation\Candlestick\Options);
        $fieldsResource = $fields->build();
        $this->internal->options->fields = $fieldsResource;
    
        return $this;
    }
    /**
     * Set which colors are used when the price movement is up or down
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Candlestick\CandlestickColors> $colors
     */
    public function colors(\Grafana\Foundation\Cog\Builder $colors): static
    {    
        if ($this->internal->options === null) {
            $this->internal->options = new \Grafana\Foundation\Candlestick\Options();
        }
        assert($this->internal->options instanceof \Grafana\Foundation\Candlestick\Options);
        $colorsResource = $colors->build();
        $this->internal->options->colors = $colorsResource;
    
        return $this;
    }
    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\VizLegendOptions> $legend
     */
    public function legend(\Grafana\Foundation\Cog\Builder $legend): static
    {    
        if ($this->internal->options === null) {
            $this->internal->options = new \Grafana\Foundation\Candlestick\Options();
        }
        assert($this->internal->options instanceof \Grafana\Foundation\Candlestick\Options);
        $legendResource = $legend->build();
        $this->internal->options->legend = $legendResource;
    
        return $this;
    }
    /**
     * When enabled, all fields will be sent to the graph
     */
    public function includeAllFields(bool $includeAllFields): static
    {    
        if ($this->internal->options === null) {
            $this->internal->options = new \Grafana\Foundation\Candlestick\Options();
        }
        assert($this->internal->options instanceof \Grafana\Foundation\Candlestick\Options);
        $this->internal->options->includeAllFields = $includeAllFields;
    
        return $this;
    }
    public function drawStyle(\Grafana\Foundation\Common\GraphDrawStyle $drawStyle): static
    {    
        if ($this->internal->fieldConfig->defaults->custom === null) {
            $this->internal->fieldConfig->defaults->custom = new \Grafana\Foundation\Candlestick\FieldConfig();
        }
        assert($this->internal->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Candlestick\FieldConfig);
        $this->internal->fieldConfig->defaults->custom->drawStyle = $drawStyle;
    
        return $this;
    }
    public function gradientMode(\Grafana\Foundation\Common\GraphGradientMode $gradientMode): static
    {    
        if ($this->internal->fieldConfig->defaults->custom === null) {
            $this->internal->fieldConfig->defaults->custom = new \Grafana\Foundation\Candlestick\FieldConfig();
        }
        assert($this->internal->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Candlestick\FieldConfig);
        $this->internal->fieldConfig->defaults->custom->gradientMode = $gradientMode;
    
        return $this;
    }
    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\GraphThresholdsStyleConfig> $thresholdsStyle
     */
    public function thresholdsStyle(\Grafana\Foundation\Cog\Builder $thresholdsStyle): static
    {    
        if ($this->internal->fieldConfig->defaults->custom === null) {
            $this->internal->fieldConfig->defaults->custom = new \Grafana\Foundation\Candlestick\FieldConfig();
        }
        assert($this->internal->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Candlestick\FieldConfig);
        $thresholdsStyleResource = $thresholdsStyle->build();
        $this->internal->fieldConfig->defaults->custom->thresholdsStyle = $thresholdsStyleResource;
    
        return $this;
    }
    public function lineColor(string $lineColor): static
    {    
        if ($this->internal->fieldConfig->defaults->custom === null) {
            $this->internal->fieldConfig->defaults->custom = new \Grafana\Foundation\Candlestick\FieldConfig();
        }
        assert($this->internal->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Candlestick\FieldConfig);
        $this->internal->fieldConfig->defaults->custom->lineColor = $lineColor;
    
        return $this;
    }
    public function lineWidth(float $lineWidth): static
    {    
        if ($this->internal->fieldConfig->defaults->custom === null) {
            $this->internal->fieldConfig->defaults->custom = new \Grafana\Foundation\Candlestick\FieldConfig();
        }
        assert($this->internal->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Candlestick\FieldConfig);
        $this->internal->fieldConfig->defaults->custom->lineWidth = $lineWidth;
    
        return $this;
    }
    public function lineInterpolation(\Grafana\Foundation\Common\LineInterpolation $lineInterpolation): static
    {    
        if ($this->internal->fieldConfig->defaults->custom === null) {
            $this->internal->fieldConfig->defaults->custom = new \Grafana\Foundation\Candlestick\FieldConfig();
        }
        assert($this->internal->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Candlestick\FieldConfig);
        $this->internal->fieldConfig->defaults->custom->lineInterpolation = $lineInterpolation;
    
        return $this;
    }
    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\LineStyle> $lineStyle
     */
    public function lineStyle(\Grafana\Foundation\Cog\Builder $lineStyle): static
    {    
        if ($this->internal->fieldConfig->defaults->custom === null) {
            $this->internal->fieldConfig->defaults->custom = new \Grafana\Foundation\Candlestick\FieldConfig();
        }
        assert($this->internal->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Candlestick\FieldConfig);
        $lineStyleResource = $lineStyle->build();
        $this->internal->fieldConfig->defaults->custom->lineStyle = $lineStyleResource;
    
        return $this;
    }
    public function fillColor(string $fillColor): static
    {    
        if ($this->internal->fieldConfig->defaults->custom === null) {
            $this->internal->fieldConfig->defaults->custom = new \Grafana\Foundation\Candlestick\FieldConfig();
        }
        assert($this->internal->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Candlestick\FieldConfig);
        $this->internal->fieldConfig->defaults->custom->fillColor = $fillColor;
    
        return $this;
    }
    public function fillOpacity(float $fillOpacity): static
    {    
        if ($this->internal->fieldConfig->defaults->custom === null) {
            $this->internal->fieldConfig->defaults->custom = new \Grafana\Foundation\Candlestick\FieldConfig();
        }
        assert($this->internal->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Candlestick\FieldConfig);
        $this->internal->fieldConfig->defaults->custom->fillOpacity = $fillOpacity;
    
        return $this;
    }
    public function showPoints(\Grafana\Foundation\Common\VisibilityMode $showPoints): static
    {    
        if ($this->internal->fieldConfig->defaults->custom === null) {
            $this->internal->fieldConfig->defaults->custom = new \Grafana\Foundation\Candlestick\FieldConfig();
        }
        assert($this->internal->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Candlestick\FieldConfig);
        $this->internal->fieldConfig->defaults->custom->showPoints = $showPoints;
    
        return $this;
    }
    public function pointSize(float $pointSize): static
    {    
        if ($this->internal->fieldConfig->defaults->custom === null) {
            $this->internal->fieldConfig->defaults->custom = new \Grafana\Foundation\Candlestick\FieldConfig();
        }
        assert($this->internal->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Candlestick\FieldConfig);
        $this->internal->fieldConfig->defaults->custom->pointSize = $pointSize;
    
        return $this;
    }
    public function pointColor(string $pointColor): static
    {    
        if ($this->internal->fieldConfig->defaults->custom === null) {
            $this->internal->fieldConfig->defaults->custom = new \Grafana\Foundation\Candlestick\FieldConfig();
        }
        assert($this->internal->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Candlestick\FieldConfig);
        $this->internal->fieldConfig->defaults->custom->pointColor = $pointColor;
    
        return $this;
    }
    public function axisPlacement(\Grafana\Foundation\Common\AxisPlacement $axisPlacement): static
    {    
        if ($this->internal->fieldConfig->defaults->custom === null) {
            $this->internal->fieldConfig->defaults->custom = new \Grafana\Foundation\Candlestick\FieldConfig();
        }
        assert($this->internal->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Candlestick\FieldConfig);
        $this->internal->fieldConfig->defaults->custom->axisPlacement = $axisPlacement;
    
        return $this;
    }
    public function axisColorMode(\Grafana\Foundation\Common\AxisColorMode $axisColorMode): static
    {    
        if ($this->internal->fieldConfig->defaults->custom === null) {
            $this->internal->fieldConfig->defaults->custom = new \Grafana\Foundation\Candlestick\FieldConfig();
        }
        assert($this->internal->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Candlestick\FieldConfig);
        $this->internal->fieldConfig->defaults->custom->axisColorMode = $axisColorMode;
    
        return $this;
    }
    public function axisLabel(string $axisLabel): static
    {    
        if ($this->internal->fieldConfig->defaults->custom === null) {
            $this->internal->fieldConfig->defaults->custom = new \Grafana\Foundation\Candlestick\FieldConfig();
        }
        assert($this->internal->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Candlestick\FieldConfig);
        $this->internal->fieldConfig->defaults->custom->axisLabel = $axisLabel;
    
        return $this;
    }
    public function axisWidth(float $axisWidth): static
    {    
        if ($this->internal->fieldConfig->defaults->custom === null) {
            $this->internal->fieldConfig->defaults->custom = new \Grafana\Foundation\Candlestick\FieldConfig();
        }
        assert($this->internal->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Candlestick\FieldConfig);
        $this->internal->fieldConfig->defaults->custom->axisWidth = $axisWidth;
    
        return $this;
    }
    public function axisSoftMin(float $axisSoftMin): static
    {    
        if ($this->internal->fieldConfig->defaults->custom === null) {
            $this->internal->fieldConfig->defaults->custom = new \Grafana\Foundation\Candlestick\FieldConfig();
        }
        assert($this->internal->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Candlestick\FieldConfig);
        $this->internal->fieldConfig->defaults->custom->axisSoftMin = $axisSoftMin;
    
        return $this;
    }
    public function axisSoftMax(float $axisSoftMax): static
    {    
        if ($this->internal->fieldConfig->defaults->custom === null) {
            $this->internal->fieldConfig->defaults->custom = new \Grafana\Foundation\Candlestick\FieldConfig();
        }
        assert($this->internal->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Candlestick\FieldConfig);
        $this->internal->fieldConfig->defaults->custom->axisSoftMax = $axisSoftMax;
    
        return $this;
    }
    public function axisGridShow(bool $axisGridShow): static
    {    
        if ($this->internal->fieldConfig->defaults->custom === null) {
            $this->internal->fieldConfig->defaults->custom = new \Grafana\Foundation\Candlestick\FieldConfig();
        }
        assert($this->internal->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Candlestick\FieldConfig);
        $this->internal->fieldConfig->defaults->custom->axisGridShow = $axisGridShow;
    
        return $this;
    }
    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\ScaleDistributionConfig> $scaleDistribution
     */
    public function scaleDistribution(\Grafana\Foundation\Cog\Builder $scaleDistribution): static
    {    
        if ($this->internal->fieldConfig->defaults->custom === null) {
            $this->internal->fieldConfig->defaults->custom = new \Grafana\Foundation\Candlestick\FieldConfig();
        }
        assert($this->internal->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Candlestick\FieldConfig);
        $scaleDistributionResource = $scaleDistribution->build();
        $this->internal->fieldConfig->defaults->custom->scaleDistribution = $scaleDistributionResource;
    
        return $this;
    }
    public function barAlignment(\Grafana\Foundation\Common\BarAlignment $barAlignment): static
    {    
        if ($this->internal->fieldConfig->defaults->custom === null) {
            $this->internal->fieldConfig->defaults->custom = new \Grafana\Foundation\Candlestick\FieldConfig();
        }
        assert($this->internal->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Candlestick\FieldConfig);
        $this->internal->fieldConfig->defaults->custom->barAlignment = $barAlignment;
    
        return $this;
    }
    public function barWidthFactor(float $barWidthFactor): static
    {    
        if ($this->internal->fieldConfig->defaults->custom === null) {
            $this->internal->fieldConfig->defaults->custom = new \Grafana\Foundation\Candlestick\FieldConfig();
        }
        assert($this->internal->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Candlestick\FieldConfig);
        $this->internal->fieldConfig->defaults->custom->barWidthFactor = $barWidthFactor;
    
        return $this;
    }
    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\StackingConfig> $stacking
     */
    public function stacking(\Grafana\Foundation\Cog\Builder $stacking): static
    {    
        if ($this->internal->fieldConfig->defaults->custom === null) {
            $this->internal->fieldConfig->defaults->custom = new \Grafana\Foundation\Candlestick\FieldConfig();
        }
        assert($this->internal->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Candlestick\FieldConfig);
        $stackingResource = $stacking->build();
        $this->internal->fieldConfig->defaults->custom->stacking = $stackingResource;
    
        return $this;
    }
    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\HideSeriesConfig> $hideFrom
     */
    public function hideFrom(\Grafana\Foundation\Cog\Builder $hideFrom): static
    {    
        if ($this->internal->fieldConfig->defaults->custom === null) {
            $this->internal->fieldConfig->defaults->custom = new \Grafana\Foundation\Candlestick\FieldConfig();
        }
        assert($this->internal->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Candlestick\FieldConfig);
        $hideFromResource = $hideFrom->build();
        $this->internal->fieldConfig->defaults->custom->hideFrom = $hideFromResource;
    
        return $this;
    }
    public function transform(\Grafana\Foundation\Common\GraphTransform $transform): static
    {    
        if ($this->internal->fieldConfig->defaults->custom === null) {
            $this->internal->fieldConfig->defaults->custom = new \Grafana\Foundation\Candlestick\FieldConfig();
        }
        assert($this->internal->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Candlestick\FieldConfig);
        $this->internal->fieldConfig->defaults->custom->transform = $transform;
    
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
        if ($this->internal->fieldConfig->defaults->custom === null) {
            $this->internal->fieldConfig->defaults->custom = new \Grafana\Foundation\Candlestick\FieldConfig();
        }
        assert($this->internal->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Candlestick\FieldConfig);
        $this->internal->fieldConfig->defaults->custom->spanNulls = $spanNulls;
    
        return $this;
    }
    public function fillBelowTo(string $fillBelowTo): static
    {    
        if ($this->internal->fieldConfig->defaults->custom === null) {
            $this->internal->fieldConfig->defaults->custom = new \Grafana\Foundation\Candlestick\FieldConfig();
        }
        assert($this->internal->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Candlestick\FieldConfig);
        $this->internal->fieldConfig->defaults->custom->fillBelowTo = $fillBelowTo;
    
        return $this;
    }
    public function pointSymbol(string $pointSymbol): static
    {    
        if ($this->internal->fieldConfig->defaults->custom === null) {
            $this->internal->fieldConfig->defaults->custom = new \Grafana\Foundation\Candlestick\FieldConfig();
        }
        assert($this->internal->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Candlestick\FieldConfig);
        $this->internal->fieldConfig->defaults->custom->pointSymbol = $pointSymbol;
    
        return $this;
    }
    public function axisCenteredZero(bool $axisCenteredZero): static
    {    
        if ($this->internal->fieldConfig->defaults->custom === null) {
            $this->internal->fieldConfig->defaults->custom = new \Grafana\Foundation\Candlestick\FieldConfig();
        }
        assert($this->internal->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Candlestick\FieldConfig);
        $this->internal->fieldConfig->defaults->custom->axisCenteredZero = $axisCenteredZero;
    
        return $this;
    }
    public function barMaxWidth(float $barMaxWidth): static
    {    
        if ($this->internal->fieldConfig->defaults->custom === null) {
            $this->internal->fieldConfig->defaults->custom = new \Grafana\Foundation\Candlestick\FieldConfig();
        }
        assert($this->internal->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Candlestick\FieldConfig);
        $this->internal->fieldConfig->defaults->custom->barMaxWidth = $barMaxWidth;
    
        return $this;
    }
    /**
     * @param bool|int $insertNulls
     */
    public function insertNulls( $insertNulls): static
    {    
        if ($this->internal->fieldConfig->defaults->custom === null) {
            $this->internal->fieldConfig->defaults->custom = new \Grafana\Foundation\Candlestick\FieldConfig();
        }
        assert($this->internal->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Candlestick\FieldConfig);
        $this->internal->fieldConfig->defaults->custom->insertNulls = $insertNulls;
    
        return $this;
    }

}
