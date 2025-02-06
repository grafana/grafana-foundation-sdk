<?php

namespace Grafana\Foundation\Statetimeline;

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
    $this->internal->type = "state-timeline";
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
     * Option for repeated panels that controls max items per row
     * Only relevant for horizontally repeated panels
     */
    public function maxPerRow(float $maxPerRow): static
    {
        $this->internal->maxPerRow = $maxPerRow;
    
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
        if ($this->internal->transformations === null) {
            $this->internal->transformations = [];
        }
        
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
     * Controls if the timeFrom or timeShift overrides are shown in the panel header
     */
    public function hideTimeOverride(bool $hideTimeOverride): static
    {
        $this->internal->hideTimeOverride = $hideTimeOverride;
    
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
     * Sets panel queries cache timeout.
     */
    public function cacheTimeout(string $cacheTimeout): static
    {
        $this->internal->cacheTimeout = $cacheTimeout;
    
        return $this;
    }
    /**
     * Overrides the data source configured time-to-live for a query cache item in milliseconds
     */
    public function queryCachingTTL(float $queryCachingTTL): static
    {
        $this->internal->queryCachingTTL = $queryCachingTTL;
    
        return $this;
    }
    /**
     * The display value for this field.  This supports template variables blank is auto
     */
    public function displayName(string $displayName): static
    {    
        if ($this->internal->fieldConfig === null) {
            $this->internal->fieldConfig = new \Grafana\Foundation\Dashboard\FieldConfigSource();
        }
        assert($this->internal->fieldConfig instanceof \Grafana\Foundation\Dashboard\FieldConfigSource);
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
        if ($this->internal->fieldConfig === null) {
            $this->internal->fieldConfig = new \Grafana\Foundation\Dashboard\FieldConfigSource();
        }
        assert($this->internal->fieldConfig instanceof \Grafana\Foundation\Dashboard\FieldConfigSource);
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
        if ($this->internal->fieldConfig === null) {
            $this->internal->fieldConfig = new \Grafana\Foundation\Dashboard\FieldConfigSource();
        }
        assert($this->internal->fieldConfig instanceof \Grafana\Foundation\Dashboard\FieldConfigSource);
        $this->internal->fieldConfig->defaults->decimals = $decimals;
    
        return $this;
    }
    /**
     * The minimum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.
     */
    public function min(float $min): static
    {    
        if ($this->internal->fieldConfig === null) {
            $this->internal->fieldConfig = new \Grafana\Foundation\Dashboard\FieldConfigSource();
        }
        assert($this->internal->fieldConfig instanceof \Grafana\Foundation\Dashboard\FieldConfigSource);
        $this->internal->fieldConfig->defaults->min = $min;
    
        return $this;
    }
    /**
     * The maximum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.
     */
    public function max(float $max): static
    {    
        if ($this->internal->fieldConfig === null) {
            $this->internal->fieldConfig = new \Grafana\Foundation\Dashboard\FieldConfigSource();
        }
        assert($this->internal->fieldConfig instanceof \Grafana\Foundation\Dashboard\FieldConfigSource);
        $this->internal->fieldConfig->defaults->max = $max;
    
        return $this;
    }
    /**
     * Convert input values into a display string
     * @param array<\Grafana\Foundation\Dashboard\ValueMap|\Grafana\Foundation\Dashboard\RangeMap|\Grafana\Foundation\Dashboard\RegexMap|\Grafana\Foundation\Dashboard\SpecialValueMap> $mappings
     */
    public function mappings(array $mappings): static
    {    
        if ($this->internal->fieldConfig === null) {
            $this->internal->fieldConfig = new \Grafana\Foundation\Dashboard\FieldConfigSource();
        }
        assert($this->internal->fieldConfig instanceof \Grafana\Foundation\Dashboard\FieldConfigSource);
        $this->internal->fieldConfig->defaults->mappings = $mappings;
    
        return $this;
    }
    /**
     * Map numeric values to states
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\ThresholdsConfig> $thresholds
     */
    public function thresholds(\Grafana\Foundation\Cog\Builder $thresholds): static
    {    
        if ($this->internal->fieldConfig === null) {
            $this->internal->fieldConfig = new \Grafana\Foundation\Dashboard\FieldConfigSource();
        }
        assert($this->internal->fieldConfig instanceof \Grafana\Foundation\Dashboard\FieldConfigSource);
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
        if ($this->internal->fieldConfig === null) {
            $this->internal->fieldConfig = new \Grafana\Foundation\Dashboard\FieldConfigSource();
        }
        assert($this->internal->fieldConfig instanceof \Grafana\Foundation\Dashboard\FieldConfigSource);
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
        if ($this->internal->fieldConfig === null) {
            $this->internal->fieldConfig = new \Grafana\Foundation\Dashboard\FieldConfigSource();
        }
        assert($this->internal->fieldConfig instanceof \Grafana\Foundation\Dashboard\FieldConfigSource);
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
        if ($this->internal->fieldConfig === null) {
            $this->internal->fieldConfig = new \Grafana\Foundation\Dashboard\FieldConfigSource();
        }
        assert($this->internal->fieldConfig instanceof \Grafana\Foundation\Dashboard\FieldConfigSource);
        $this->internal->fieldConfig->defaults->noValue = $noValue;
    
        return $this;
    }
    /**
     * Overrides are the options applied to specific fields overriding the defaults.
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\DashboardFieldConfigSourceOverrides>> $overrides
     */
    public function overrides(array $overrides): static
    {    
        if ($this->internal->fieldConfig === null) {
            $this->internal->fieldConfig = new \Grafana\Foundation\Dashboard\FieldConfigSource();
        }
        assert($this->internal->fieldConfig instanceof \Grafana\Foundation\Dashboard\FieldConfigSource);
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
        if ($this->internal->fieldConfig === null) {
            $this->internal->fieldConfig = new \Grafana\Foundation\Dashboard\FieldConfigSource();
        }
        assert($this->internal->fieldConfig instanceof \Grafana\Foundation\Dashboard\FieldConfigSource);
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
        if ($this->internal->fieldConfig === null) {
            $this->internal->fieldConfig = new \Grafana\Foundation\Dashboard\FieldConfigSource();
        }
        assert($this->internal->fieldConfig instanceof \Grafana\Foundation\Dashboard\FieldConfigSource);
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
        if ($this->internal->fieldConfig === null) {
            $this->internal->fieldConfig = new \Grafana\Foundation\Dashboard\FieldConfigSource();
        }
        assert($this->internal->fieldConfig instanceof \Grafana\Foundation\Dashboard\FieldConfigSource);
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
        if ($this->internal->fieldConfig === null) {
            $this->internal->fieldConfig = new \Grafana\Foundation\Dashboard\FieldConfigSource();
        }
        assert($this->internal->fieldConfig instanceof \Grafana\Foundation\Dashboard\FieldConfigSource);
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
        if ($this->internal->fieldConfig === null) {
            $this->internal->fieldConfig = new \Grafana\Foundation\Dashboard\FieldConfigSource();
        }
        assert($this->internal->fieldConfig instanceof \Grafana\Foundation\Dashboard\FieldConfigSource);
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
     * Show timeline values on chart
     */
    public function showValue(\Grafana\Foundation\Common\VisibilityMode $showValue): static
    {    
        if ($this->internal->options === null) {
            $this->internal->options = new \Grafana\Foundation\Statetimeline\Options();
        }
        assert($this->internal->options instanceof \Grafana\Foundation\Statetimeline\Options);
        $this->internal->options->showValue = $showValue;
    
        return $this;
    }
    /**
     * Controls the row height
     */
    public function rowHeight(float $rowHeight): static
    {
        if (!($rowHeight <= 1)) {
            throw new \ValueError('$rowHeight must be <= 1');
        }    
        if ($this->internal->options === null) {
            $this->internal->options = new \Grafana\Foundation\Statetimeline\Options();
        }
        assert($this->internal->options instanceof \Grafana\Foundation\Statetimeline\Options);
        $this->internal->options->rowHeight = $rowHeight;
    
        return $this;
    }
    /**
     * Merge equal consecutive values
     */
    public function mergeValues(bool $mergeValues): static
    {    
        if ($this->internal->options === null) {
            $this->internal->options = new \Grafana\Foundation\Statetimeline\Options();
        }
        assert($this->internal->options instanceof \Grafana\Foundation\Statetimeline\Options);
        $this->internal->options->mergeValues = $mergeValues;
    
        return $this;
    }
    /**
     * Controls value alignment on the timelines
     */
    public function alignValue(\Grafana\Foundation\Common\TimelineValueAlignment $alignValue): static
    {    
        if ($this->internal->options === null) {
            $this->internal->options = new \Grafana\Foundation\Statetimeline\Options();
        }
        assert($this->internal->options instanceof \Grafana\Foundation\Statetimeline\Options);
        $this->internal->options->alignValue = $alignValue;
    
        return $this;
    }
    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\VizLegendOptions> $legend
     */
    public function legend(\Grafana\Foundation\Cog\Builder $legend): static
    {    
        if ($this->internal->options === null) {
            $this->internal->options = new \Grafana\Foundation\Statetimeline\Options();
        }
        assert($this->internal->options instanceof \Grafana\Foundation\Statetimeline\Options);
        $legendResource = $legend->build();
        $this->internal->options->legend = $legendResource;
    
        return $this;
    }
    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\VizTooltipOptions> $tooltip
     */
    public function tooltip(\Grafana\Foundation\Cog\Builder $tooltip): static
    {    
        if ($this->internal->options === null) {
            $this->internal->options = new \Grafana\Foundation\Statetimeline\Options();
        }
        assert($this->internal->options instanceof \Grafana\Foundation\Statetimeline\Options);
        $tooltipResource = $tooltip->build();
        $this->internal->options->tooltip = $tooltipResource;
    
        return $this;
    }
    /**
     * @param array<string> $timezone
     */
    public function timezone(array $timezone): static
    {    
        if ($this->internal->options === null) {
            $this->internal->options = new \Grafana\Foundation\Statetimeline\Options();
        }
        assert($this->internal->options instanceof \Grafana\Foundation\Statetimeline\Options);
        $this->internal->options->timezone = $timezone;
    
        return $this;
    }
    /**
     * Enables pagination when > 0
     */
    public function perPage(float $perPage): static
    {
        if (!($perPage >= 1)) {
            throw new \ValueError('$perPage must be >= 1');
        }    
        if ($this->internal->options === null) {
            $this->internal->options = new \Grafana\Foundation\Statetimeline\Options();
        }
        assert($this->internal->options instanceof \Grafana\Foundation\Statetimeline\Options);
        $this->internal->options->perPage = $perPage;
    
        return $this;
    }
    public function lineWidth(int $lineWidth): static
    {
        if (!($lineWidth <= 10)) {
            throw new \ValueError('$lineWidth must be <= 10');
        }    
        if ($this->internal->fieldConfig === null) {
            $this->internal->fieldConfig = new \Grafana\Foundation\Dashboard\FieldConfigSource();
        }
        assert($this->internal->fieldConfig instanceof \Grafana\Foundation\Dashboard\FieldConfigSource);    
        if ($this->internal->fieldConfig->defaults->custom === null) {
            $this->internal->fieldConfig->defaults->custom = new \Grafana\Foundation\Statetimeline\FieldConfig();
        }
        assert($this->internal->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Statetimeline\FieldConfig);
        $this->internal->fieldConfig->defaults->custom->lineWidth = $lineWidth;
    
        return $this;
    }
    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\HideSeriesConfig> $hideFrom
     */
    public function hideFrom(\Grafana\Foundation\Cog\Builder $hideFrom): static
    {    
        if ($this->internal->fieldConfig === null) {
            $this->internal->fieldConfig = new \Grafana\Foundation\Dashboard\FieldConfigSource();
        }
        assert($this->internal->fieldConfig instanceof \Grafana\Foundation\Dashboard\FieldConfigSource);    
        if ($this->internal->fieldConfig->defaults->custom === null) {
            $this->internal->fieldConfig->defaults->custom = new \Grafana\Foundation\Statetimeline\FieldConfig();
        }
        assert($this->internal->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Statetimeline\FieldConfig);
        $hideFromResource = $hideFrom->build();
        $this->internal->fieldConfig->defaults->custom->hideFrom = $hideFromResource;
    
        return $this;
    }
    public function fillOpacity(int $fillOpacity): static
    {
        if (!($fillOpacity <= 100)) {
            throw new \ValueError('$fillOpacity must be <= 100');
        }    
        if ($this->internal->fieldConfig === null) {
            $this->internal->fieldConfig = new \Grafana\Foundation\Dashboard\FieldConfigSource();
        }
        assert($this->internal->fieldConfig instanceof \Grafana\Foundation\Dashboard\FieldConfigSource);    
        if ($this->internal->fieldConfig->defaults->custom === null) {
            $this->internal->fieldConfig->defaults->custom = new \Grafana\Foundation\Statetimeline\FieldConfig();
        }
        assert($this->internal->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Statetimeline\FieldConfig);
        $this->internal->fieldConfig->defaults->custom->fillOpacity = $fillOpacity;
    
        return $this;
    }

}
