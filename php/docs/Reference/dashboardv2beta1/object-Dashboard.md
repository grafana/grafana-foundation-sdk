---
title: <span class="badge object-type-class"></span> Dashboard
---
# <span class="badge object-type-class"></span> Dashboard

## Definition

```php
class Dashboard implements \JsonSerializable
{
    /**
     * @var array<\Grafana\Foundation\Dashboardv2beta1\AnnotationQueryKind>
     */
    public array $annotations;

    /**
     * Configuration of dashboard cursor sync behavior.
     * "Off" for no shared crosshair or tooltip (default).
     * "Crosshair" for shared crosshair.
     * "Tooltip" for shared crosshair AND shared tooltip.
     */
    public \Grafana\Foundation\Dashboardv2beta1\DashboardCursorSync $cursorSync;

    /**
     * Description of dashboard.
     */
    public ?string $description;

    /**
     * Whether a dashboard is editable or not.
     */
    public ?bool $editable;

    /**
     * @var array<string, \Grafana\Foundation\Dashboardv2beta1\PanelKind|\Grafana\Foundation\Dashboardv2beta1\LibraryPanelKind>
     */
    public array $elements;

    /**
     * @var \Grafana\Foundation\Dashboardv2beta1\GridLayoutKind|\Grafana\Foundation\Dashboardv2beta1\RowsLayoutKind|\Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutKind|\Grafana\Foundation\Dashboardv2beta1\TabsLayoutKind
     */
    public $layout;

    /**
     * Links with references to other dashboards or external websites.
     * @var array<\Grafana\Foundation\Dashboardv2beta1\DashboardLink>
     */
    public array $links;

    /**
     * When set to true, the dashboard will redraw panels at an interval matching the pixel width.
     * This will keep data "moving left" regardless of the query refresh rate. This setting helps
     * avoid dashboards presenting stale live data.
     */
    public ?bool $liveNow;

    /**
     * When set to true, the dashboard will load all panels in the dashboard when it's loaded.
     */
    public bool $preload;

    /**
     * Plugins only. The version of the dashboard installed together with the plugin.
     * This is used to determine if the dashboard should be updated when the plugin is updated.
     */
    public ?int $revision;

    /**
     * Tags associated with dashboard.
     * @var array<string>
     */
    public array $tags;

    public \Grafana\Foundation\Dashboardv2beta1\TimeSettingsSpec $timeSettings;

    /**
     * Title of dashboard.
     */
    public string $title;

    /**
     * Configured template variables.
     * @var array<\Grafana\Foundation\Dashboardv2beta1\QueryVariableKind|\Grafana\Foundation\Dashboardv2beta1\TextVariableKind|\Grafana\Foundation\Dashboardv2beta1\ConstantVariableKind|\Grafana\Foundation\Dashboardv2beta1\DatasourceVariableKind|\Grafana\Foundation\Dashboardv2beta1\IntervalVariableKind|\Grafana\Foundation\Dashboardv2beta1\CustomVariableKind|\Grafana\Foundation\Dashboardv2beta1\GroupByVariableKind|\Grafana\Foundation\Dashboardv2beta1\AdhocVariableKind|\Grafana\Foundation\Dashboardv2beta1\SwitchVariableKind>
     */
    public array $variables;

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

 * <span class="badge builder"></span> [DashboardBuilder](./builder-DashboardBuilder.md)
