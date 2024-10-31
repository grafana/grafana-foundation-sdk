<?php

namespace Grafana\Foundation\Librarypanel;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Librarypanel\LibrarypanelLibraryPanelModel>
 */
class LibrarypanelLibraryPanelModelBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Librarypanel\LibrarypanelLibraryPanelModel $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Librarypanel\LibrarypanelLibraryPanelModel();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Librarypanel\LibrarypanelLibraryPanelModel
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * The panel plugin type id. This is used to find the plugin to display the panel.
     */
    public function type(string $type): static
    {
        if (!(strlen($type) >= 1)) {
            throw new \ValueError('strlen($type) must be >= 1');
        }
        $this->internal->type = $type;
    
        return $this;
    }
    /**
     * The version of the plugin that is used for this panel. This is used to find the plugin to display the panel and to migrate old panel configs.
     */
    public function pluginVersion(string $pluginVersion): static
    {
        $this->internal->pluginVersion = $pluginVersion;
    
        return $this;
    }
    /**
     * Tags for the panel.
     * @param array<string> $tags
     */
    public function tags(array $tags): static
    {
        $this->internal->tags = $tags;
    
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
    public function repeatDirection(\Grafana\Foundation\Librarypanel\LibraryPanelRepeatDirection $repeatDirection): static
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
     * It depends on the panel plugin. They are specified by the Options field in panel plugin schemas.
     * @param mixed $options
     */
    public function options( $options): static
    {
        $this->internal->options = $options;
    
        return $this;
    }
    /**
     * Field options allow you to change how the data is displayed in your visualizations.
     */
    public function fieldConfig(\Grafana\Foundation\Dashboard\FieldConfigSource $fieldConfig): static
    {
        $this->internal->fieldConfig = $fieldConfig;
    
        return $this;
    }

}
