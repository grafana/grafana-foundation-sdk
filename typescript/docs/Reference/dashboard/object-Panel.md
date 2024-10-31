---
title: <span class="badge object-type-interface"></span> Panel
---
# <span class="badge object-type-interface"></span> Panel

Dashboard panels are the basic visualization building blocks.

## Definition

```typescript
export interface Panel {
	// The panel plugin type id. This is used to find the plugin to display the panel.
	type: string;
	// Unique identifier of the panel. Generated by Grafana when creating a new panel. It must be unique within a dashboard, but not globally.
	id?: number;
	// The version of the plugin that is used for this panel. This is used to find the plugin to display the panel and to migrate old panel configs.
	pluginVersion?: string;
	// Tags for the panel.
	tags?: string[];
	// Depends on the panel plugin. See the plugin documentation for details.
	targets?: cog.Dataquery[];
	// Panel title.
	title?: string;
	// Panel description.
	description?: string;
	// Whether to display the panel without a background.
	transparent: boolean;
	// The datasource used in all targets.
	datasource?: dashboard.DataSourceRef;
	// Grid position.
	gridPos?: dashboard.GridPos;
	// Panel links.
	links?: dashboard.DashboardLink[];
	// Name of template variable to repeat for.
	repeat?: string;
	// Direction to repeat in if 'repeat' is set.
	// `h` for horizontal, `v` for vertical.
	repeatDirection?: "h" | "v";
	// Id of the repeating panel.
	repeatPanelId?: number;
	// The maximum number of data points that the panel queries are retrieving.
	maxDataPoints?: number;
	// List of transformations that are applied to the panel data before rendering.
	// When there are multiple transformations, Grafana applies them in the order they are listed.
	// Each transformation creates a result set that then passes on to the next transformation in the processing pipeline.
	transformations: dashboard.DataTransformerConfig[];
	// The min time interval setting defines a lower limit for the $__interval and $__interval_ms variables.
	// This value must be formatted as a number followed by a valid time
	// identifier like: "40s", "3d", etc.
	// See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options
	interval?: string;
	// Overrides the relative time range for individual panels,
	// which causes them to be different than what is selected in
	// the dashboard time picker in the top-right corner of the dashboard. You can use this to show metrics from different
	// time periods or days on the same dashboard.
	// The value is formatted as time operation like: `now-5m` (Last 5 minutes), `now/d` (the day so far),
	// `now-5d/d`(Last 5 days), `now/w` (This week so far), `now-2y/y` (Last 2 years).
	// Note: Panel time overrides have no effect when the dashboard’s time range is absolute.
	// See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options
	timeFrom?: string;
	// Overrides the time range for individual panels by shifting its start and end relative to the time picker.
	// For example, you can shift the time range for the panel to be two hours earlier than the dashboard time picker setting `2h`.
	// Note: Panel time overrides have no effect when the dashboard’s time range is absolute.
	// See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options
	timeShift?: string;
	// Dynamically load the panel
	libraryPanel?: dashboard.LibraryPanelRef;
	// It depends on the panel plugin. They are specified by the Options field in panel plugin schemas.
	options: any;
	// Field options allow you to change how the data is displayed in your visualizations.
	fieldConfig: dashboard.FieldConfigSource;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [PanelBuilder](./builder-PanelBuilder.md)
 * <span class="badge builder"></span> [heatmap.PanelBuilder](../heatmap/builder-PanelBuilder.md)
 * <span class="badge builder"></span> [table.PanelBuilder](../table/builder-PanelBuilder.md)
 * <span class="badge builder"></span> [text.PanelBuilder](../text/builder-PanelBuilder.md)
 * <span class="badge builder"></span> [bargauge.PanelBuilder](../bargauge/builder-PanelBuilder.md)
 * <span class="badge builder"></span> [candlestick.PanelBuilder](../candlestick/builder-PanelBuilder.md)
 * <span class="badge builder"></span> [gauge.PanelBuilder](../gauge/builder-PanelBuilder.md)
 * <span class="badge builder"></span> [statushistory.PanelBuilder](../statushistory/builder-PanelBuilder.md)
 * <span class="badge builder"></span> [geomap.PanelBuilder](../geomap/builder-PanelBuilder.md)
 * <span class="badge builder"></span> [barchart.PanelBuilder](../barchart/builder-PanelBuilder.md)
 * <span class="badge builder"></span> [histogram.PanelBuilder](../histogram/builder-PanelBuilder.md)
 * <span class="badge builder"></span> [annotationslist.PanelBuilder](../annotationslist/builder-PanelBuilder.md)
 * <span class="badge builder"></span> [news.PanelBuilder](../news/builder-PanelBuilder.md)
 * <span class="badge builder"></span> [stat.PanelBuilder](../stat/builder-PanelBuilder.md)
 * <span class="badge builder"></span> [timeseries.PanelBuilder](../timeseries/builder-PanelBuilder.md)
 * <span class="badge builder"></span> [trend.PanelBuilder](../trend/builder-PanelBuilder.md)
 * <span class="badge builder"></span> [xychart.PanelBuilder](../xychart/builder-PanelBuilder.md)
 * <span class="badge builder"></span> [datagrid.PanelBuilder](../datagrid/builder-PanelBuilder.md)
 * <span class="badge builder"></span> [logs.PanelBuilder](../logs/builder-PanelBuilder.md)
 * <span class="badge builder"></span> [alertgroups.PanelBuilder](../alertgroups/builder-PanelBuilder.md)
 * <span class="badge builder"></span> [nodegraph.PanelBuilder](../nodegraph/builder-PanelBuilder.md)
 * <span class="badge builder"></span> [piechart.PanelBuilder](../piechart/builder-PanelBuilder.md)
 * <span class="badge builder"></span> [debug.PanelBuilder](../debug/builder-PanelBuilder.md)
 * <span class="badge builder"></span> [dashboardlist.PanelBuilder](../dashboardlist/builder-PanelBuilder.md)
 * <span class="badge builder"></span> [statetimeline.PanelBuilder](../statetimeline/builder-PanelBuilder.md)
 * <span class="badge builder"></span> [canvas.PanelBuilder](../canvas/builder-PanelBuilder.md)