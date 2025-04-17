<?php

namespace Grafana\Foundation\Librarypanel;

/**
 * Dashboard panels are the basic visualization building blocks.
 */
class PanelModel implements \JsonSerializable
{
    /**
     * The panel plugin type id. This is used to find the plugin to display the panel.
     */
    public string $type;

    /**
     * The version of the plugin that is used for this panel. This is used to find the plugin to display the panel and to migrate old panel configs.
     */
    public ?string $pluginVersion;

    /**
     * Depends on the panel plugin. See the plugin documentation for details.
     * @var array<\Grafana\Foundation\Cog\Dataquery>|null
     */
    public ?array $targets;

    /**
     * Panel title.
     */
    public ?string $title;

    /**
     * Panel description.
     */
    public ?string $description;

    /**
     * Whether to display the panel without a background.
     */
    public ?bool $transparent;

    /**
     * The datasource used in all targets.
     */
    public ?\Grafana\Foundation\Dashboard\DataSourceRef $datasource;

    /**
     * Panel links.
     * @var array<\Grafana\Foundation\Dashboard\DashboardLink>|null
     */
    public ?array $links;

    /**
     * Name of template variable to repeat for.
     */
    public ?string $repeat;

    /**
     * Direction to repeat in if 'repeat' is set.
     * `h` for horizontal, `v` for vertical.
     */
    public ?\Grafana\Foundation\Librarypanel\PanelModelRepeatDirection $repeatDirection;

    /**
     * Option for repeated panels that controls max items per row
     * Only relevant for horizontally repeated panels
     */
    public ?float $maxPerRow;

    /**
     * The maximum number of data points that the panel queries are retrieving.
     */
    public ?float $maxDataPoints;

    /**
     * List of transformations that are applied to the panel data before rendering.
     * When there are multiple transformations, Grafana applies them in the order they are listed.
     * Each transformation creates a result set that then passes on to the next transformation in the processing pipeline.
     * @var array<\Grafana\Foundation\Dashboard\DataTransformerConfig>|null
     */
    public ?array $transformations;

    /**
     * The min time interval setting defines a lower limit for the $__interval and $__interval_ms variables.
     * This value must be formatted as a number followed by a valid time
     * identifier like: "40s", "3d", etc.
     * See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options
     */
    public ?string $interval;

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
    public ?string $timeFrom;

    /**
     * Overrides the time range for individual panels by shifting its start and end relative to the time picker.
     * For example, you can shift the time range for the panel to be two hours earlier than the dashboard time picker setting `2h`.
     * Note: Panel time overrides have no effect when the dashboard’s time range is absolute.
     * See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options
     */
    public ?string $timeShift;

    /**
     * Controls if the timeFrom or timeShift overrides are shown in the panel header
     */
    public ?bool $hideTimeOverride;

    /**
     * Sets panel queries cache timeout.
     */
    public ?string $cacheTimeout;

    /**
     * Overrides the data source configured time-to-live for a query cache item in milliseconds
     */
    public ?float $queryCachingTTL;

    /**
     * It depends on the panel plugin. They are specified by the Options field in panel plugin schemas.
     * @var mixed|null
     */
    public $options;

    /**
     * Field options allow you to change how the data is displayed in your visualizations.
     */
    public ?\Grafana\Foundation\Dashboard\FieldConfigSource $fieldConfig;

    /**
     * @param string|null $type
     * @param string|null $pluginVersion
     * @param array<\Grafana\Foundation\Cog\Dataquery>|null $targets
     * @param string|null $title
     * @param string|null $description
     * @param bool|null $transparent
     * @param \Grafana\Foundation\Dashboard\DataSourceRef|null $datasource
     * @param array<\Grafana\Foundation\Dashboard\DashboardLink>|null $links
     * @param string|null $repeat
     * @param \Grafana\Foundation\Librarypanel\PanelModelRepeatDirection|null $repeatDirection
     * @param float|null $maxPerRow
     * @param float|null $maxDataPoints
     * @param array<\Grafana\Foundation\Dashboard\DataTransformerConfig>|null $transformations
     * @param string|null $interval
     * @param string|null $timeFrom
     * @param string|null $timeShift
     * @param bool|null $hideTimeOverride
     * @param string|null $cacheTimeout
     * @param float|null $queryCachingTTL
     * @param mixed|null $options
     * @param \Grafana\Foundation\Dashboard\FieldConfigSource|null $fieldConfig
     */
    public function __construct(?string $type = null, ?string $pluginVersion = null, ?array $targets = null, ?string $title = null, ?string $description = null, ?bool $transparent = null, ?\Grafana\Foundation\Dashboard\DataSourceRef $datasource = null, ?array $links = null, ?string $repeat = null, ?\Grafana\Foundation\Librarypanel\PanelModelRepeatDirection $repeatDirection = null, ?float $maxPerRow = null, ?float $maxDataPoints = null, ?array $transformations = null, ?string $interval = null, ?string $timeFrom = null, ?string $timeShift = null, ?bool $hideTimeOverride = null, ?string $cacheTimeout = null, ?float $queryCachingTTL = null,  $options = null, ?\Grafana\Foundation\Dashboard\FieldConfigSource $fieldConfig = null)
    {
        $this->type = $type ?: "";
        $this->pluginVersion = $pluginVersion;
        $this->targets = $targets;
        $this->title = $title;
        $this->description = $description;
        $this->transparent = $transparent;
        $this->datasource = $datasource;
        $this->links = $links;
        $this->repeat = $repeat;
        $this->repeatDirection = $repeatDirection;
        $this->maxPerRow = $maxPerRow;
        $this->maxDataPoints = $maxDataPoints;
        $this->transformations = $transformations;
        $this->interval = $interval;
        $this->timeFrom = $timeFrom;
        $this->timeShift = $timeShift;
        $this->hideTimeOverride = $hideTimeOverride;
        $this->cacheTimeout = $cacheTimeout;
        $this->queryCachingTTL = $queryCachingTTL;
        $this->options = $options;
        $this->fieldConfig = $fieldConfig;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{type?: string, pluginVersion?: string, targets?: array<mixed>, title?: string, description?: string, transparent?: bool, datasource?: mixed, links?: array<mixed>, repeat?: string, repeatDirection?: string, maxPerRow?: float, maxDataPoints?: float, transformations?: array<mixed>, interval?: string, timeFrom?: string, timeShift?: string, hideTimeOverride?: bool, cacheTimeout?: string, queryCachingTTL?: float, options?: mixed, fieldConfig?: mixed} $inputData */
        $data = $inputData;
        return new self(
            type: $data["type"] ?? null,
            pluginVersion: $data["pluginVersion"] ?? null,
            targets: isset($data["targets"]) ? (function ($in) {
    	/** @var array{datasource?: array{type?: mixed}} $in */
        $hint = (isset($in["datasource"], $in["datasource"]["type"]) && is_string($in["datasource"]["type"])) ? $in["datasource"]["type"] : "";
    /** @var array<array<string, mixed>> $in */
        return \Grafana\Foundation\Cog\Runtime::get()->dataqueriesFromArray($in, $hint);
    })($data["targets"]) : null,
            title: $data["title"] ?? null,
            description: $data["description"] ?? null,
            transparent: $data["transparent"] ?? null,
            datasource: isset($data["datasource"]) ? (function($input) {
    	/** @var array{type?: string, uid?: string} */
    $val = $input;
    	return \Grafana\Foundation\Dashboard\DataSourceRef::fromArray($val);
    })($data["datasource"]) : null,
            links: array_filter(array_map((function($input) {
    	/** @var array{title?: string, type?: string, icon?: string, tooltip?: string, url?: string, tags?: array<string>, asDropdown?: bool, targetBlank?: bool, includeVars?: bool, keepTime?: bool} */
    $val = $input;
    	return \Grafana\Foundation\Dashboard\DashboardLink::fromArray($val);
    }), $data["links"] ?? [])),
            repeat: $data["repeat"] ?? null,
            repeatDirection: isset($data["repeatDirection"]) ? (function($input) { return \Grafana\Foundation\Librarypanel\PanelModelRepeatDirection::fromValue($input); })($data["repeatDirection"]) : null,
            maxPerRow: $data["maxPerRow"] ?? null,
            maxDataPoints: $data["maxDataPoints"] ?? null,
            transformations: array_filter(array_map((function($input) {
    	/** @var array{id?: string, disabled?: bool, filter?: mixed, topic?: string, options?: mixed} */
    $val = $input;
    	return \Grafana\Foundation\Dashboard\DataTransformerConfig::fromArray($val);
    }), $data["transformations"] ?? [])),
            interval: $data["interval"] ?? null,
            timeFrom: $data["timeFrom"] ?? null,
            timeShift: $data["timeShift"] ?? null,
            hideTimeOverride: $data["hideTimeOverride"] ?? null,
            cacheTimeout: $data["cacheTimeout"] ?? null,
            queryCachingTTL: $data["queryCachingTTL"] ?? null,
            options: $data["options"] ?? null,
            fieldConfig: isset($data["fieldConfig"]) ? (function($input) {
    	/** @var array{defaults?: mixed, overrides?: array<mixed>} */
    $val = $input;
    	return \Grafana\Foundation\Dashboard\FieldConfigSource::fromArray($val);
    })($data["fieldConfig"]) : null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->type = $this->type;
        if (isset($this->pluginVersion)) {
            $data->pluginVersion = $this->pluginVersion;
        }
        if (isset($this->targets)) {
            $data->targets = $this->targets;
        }
        if (isset($this->title)) {
            $data->title = $this->title;
        }
        if (isset($this->description)) {
            $data->description = $this->description;
        }
        if (isset($this->transparent)) {
            $data->transparent = $this->transparent;
        }
        if (isset($this->datasource)) {
            $data->datasource = $this->datasource;
        }
        if (isset($this->links)) {
            $data->links = $this->links;
        }
        if (isset($this->repeat)) {
            $data->repeat = $this->repeat;
        }
        if (isset($this->repeatDirection)) {
            $data->repeatDirection = $this->repeatDirection;
        }
        if (isset($this->maxPerRow)) {
            $data->maxPerRow = $this->maxPerRow;
        }
        if (isset($this->maxDataPoints)) {
            $data->maxDataPoints = $this->maxDataPoints;
        }
        if (isset($this->transformations)) {
            $data->transformations = $this->transformations;
        }
        if (isset($this->interval)) {
            $data->interval = $this->interval;
        }
        if (isset($this->timeFrom)) {
            $data->timeFrom = $this->timeFrom;
        }
        if (isset($this->timeShift)) {
            $data->timeShift = $this->timeShift;
        }
        if (isset($this->hideTimeOverride)) {
            $data->hideTimeOverride = $this->hideTimeOverride;
        }
        if (isset($this->cacheTimeout)) {
            $data->cacheTimeout = $this->cacheTimeout;
        }
        if (isset($this->queryCachingTTL)) {
            $data->queryCachingTTL = $this->queryCachingTTL;
        }
        if (isset($this->options)) {
            $data->options = $this->options;
        }
        if (isset($this->fieldConfig)) {
            $data->fieldConfig = $this->fieldConfig;
        }
        return $data;
    }
}
