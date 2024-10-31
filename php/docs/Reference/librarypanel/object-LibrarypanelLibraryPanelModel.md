---
title: <span class="badge object-type-class"></span> LibrarypanelLibraryPanelModel
---
# <span class="badge object-type-class"></span> LibrarypanelLibraryPanelModel

## Definition

```php
class LibrarypanelLibraryPanelModel implements \JsonSerializable
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
     * Tags for the panel.
     * @var array<string>|null
     */
    public ?array $tags;

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
    public ?\Grafana\Foundation\Librarypanel\LibraryPanelRepeatDirection $repeatDirection;

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
     * It depends on the panel plugin. They are specified by the Options field in panel plugin schemas.
     * @var mixed|null
     */
    public $options;

    /**
     * Field options allow you to change how the data is displayed in your visualizations.
     */
    public ?\Grafana\Foundation\Dashboard\FieldConfigSource $fieldConfig;

}
```
## Methods

### <span class="badge object-method"></span> fromArray

Builds this object from an array.

This function is meant to be used with the return value of `json_decode($json, true)`.

```php
static fromArray(array $inputData)
```

### <span class="badge object-method"></span> jsonSerialize

Returns the data representing this object, preparing it for JSON serialization with `json_encode()`.

```php
jsonSerialize()
```

## See also

 * <span class="badge builder"></span> [LibrarypanelLibraryPanelModelBuilder](./builder-LibrarypanelLibraryPanelModelBuilder.md)
