---
title: <span class="badge object-type-struct"></span> LibrarypanelLibraryPanelModel
---
# <span class="badge object-type-struct"></span> LibrarypanelLibraryPanelModel

## Definition

```go
type LibrarypanelLibraryPanelModel struct {
    // The panel plugin type id. This is used to find the plugin to display the panel.
    Type string `json:"type"`
    // The version of the plugin that is used for this panel. This is used to find the plugin to display the panel and to migrate old panel configs.
    PluginVersion *string `json:"pluginVersion,omitempty"`
    // Depends on the panel plugin. See the plugin documentation for details.
    Targets []cog/variants.Dataquery `json:"targets,omitempty"`
    // Panel title.
    Title *string `json:"title,omitempty"`
    // Panel description.
    Description *string `json:"description,omitempty"`
    // Whether to display the panel without a background.
    Transparent *bool `json:"transparent,omitempty"`
    // The datasource used in all targets.
    Datasource *dashboard.DataSourceRef `json:"datasource,omitempty"`
    // Panel links.
    Links []dashboard.DashboardLink `json:"links,omitempty"`
    // Name of template variable to repeat for.
    Repeat *string `json:"repeat,omitempty"`
    // Direction to repeat in if 'repeat' is set.
    // `h` for horizontal, `v` for vertical.
    RepeatDirection *librarypanel.LibraryPanelRepeatDirection `json:"repeatDirection,omitempty"`
    // Option for repeated panels that controls max items per row
    // Only relevant for horizontally repeated panels
    MaxPerRow *float64 `json:"maxPerRow,omitempty"`
    // The maximum number of data points that the panel queries are retrieving.
    MaxDataPoints *float64 `json:"maxDataPoints,omitempty"`
    // List of transformations that are applied to the panel data before rendering.
    // When there are multiple transformations, Grafana applies them in the order they are listed.
    // Each transformation creates a result set that then passes on to the next transformation in the processing pipeline.
    Transformations []dashboard.DataTransformerConfig `json:"transformations,omitempty"`
    // The min time interval setting defines a lower limit for the $__interval and $__interval_ms variables.
    // This value must be formatted as a number followed by a valid time
    // identifier like: "40s", "3d", etc.
    // See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options
    Interval *string `json:"interval,omitempty"`
    // Overrides the relative time range for individual panels,
    // which causes them to be different than what is selected in
    // the dashboard time picker in the top-right corner of the dashboard. You can use this to show metrics from different
    // time periods or days on the same dashboard.
    // The value is formatted as time operation like: `now-5m` (Last 5 minutes), `now/d` (the day so far),
    // `now-5d/d`(Last 5 days), `now/w` (This week so far), `now-2y/y` (Last 2 years).
    // Note: Panel time overrides have no effect when the dashboard’s time range is absolute.
    // See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options
    TimeFrom *string `json:"timeFrom,omitempty"`
    // Overrides the time range for individual panels by shifting its start and end relative to the time picker.
    // For example, you can shift the time range for the panel to be two hours earlier than the dashboard time picker setting `2h`.
    // Note: Panel time overrides have no effect when the dashboard’s time range is absolute.
    // See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options
    TimeShift *string `json:"timeShift,omitempty"`
    // Controls if the timeFrom or timeShift overrides are shown in the panel header
    HideTimeOverride *bool `json:"hideTimeOverride,omitempty"`
    // Sets panel queries cache timeout.
    CacheTimeout *string `json:"cacheTimeout,omitempty"`
    // Overrides the data source configured time-to-live for a query cache item in milliseconds
    QueryCachingTTL *float64 `json:"queryCachingTTL,omitempty"`
    // It depends on the panel plugin. They are specified by the Options field in panel plugin schemas.
    Options any `json:"options,omitempty"`
    // Field options allow you to change how the data is displayed in your visualizations.
    FieldConfig *dashboard.FieldConfigSource `json:"fieldConfig,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSON

UnmarshalJSON implements a custom JSON unmarshalling logic to decode `LibrarypanelLibraryPanelModel` from JSON.

```go
func (librarypanelLibraryPanelModel *LibrarypanelLibraryPanelModel) UnmarshalJSON(raw []byte) error
```

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `LibrarypanelLibraryPanelModel` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (librarypanelLibraryPanelModel *LibrarypanelLibraryPanelModel) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `LibrarypanelLibraryPanelModel` objects.

```go
func (librarypanelLibraryPanelModel *LibrarypanelLibraryPanelModel) Equals(other LibrarypanelLibraryPanelModel) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `LibrarypanelLibraryPanelModel` fields for violations and returns them.

```go
func (librarypanelLibraryPanelModel *LibrarypanelLibraryPanelModel) Validate() error
```

## See also

 * <span class="badge builder"></span> [LibrarypanelLibraryPanelModelBuilder](./builder-LibrarypanelLibraryPanelModelBuilder.md)
