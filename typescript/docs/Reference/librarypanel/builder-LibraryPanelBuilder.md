---
title: <span class="badge builder"></span> LibraryPanelBuilder
---
# <span class="badge builder"></span> LibraryPanelBuilder

## Constructor

```typescript
new LibraryPanelBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> description

Panel description

```typescript
description(description: string)
```

### <span class="badge object-method"></span> folderUid

Folder UID

```typescript
folderUid(folderUid: string)
```

### <span class="badge object-method"></span> meta

Object storage metadata

```typescript
meta(meta: cog.Builder<librarypanel.LibraryElementDTOMeta>)
```

### <span class="badge object-method"></span> model

TODO: should be the same panel schema defined in dashboard

Typescript: Omit<Panel, 'gridPos' | 'id' | 'libraryPanel'>;

```typescript
model(model: {
	// The panel plugin type id. This is used to find the plugin to display the panel.
	type: string;
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
	// It depends on the panel plugin. They are specified by the Options field in panel plugin schemas.
	options: any;
	// Field options allow you to change how the data is displayed in your visualizations.
	fieldConfig: dashboard.FieldConfigSource;
})
```

### <span class="badge object-method"></span> name

Panel name (also saved in the model)

```typescript
name(name: string)
```

### <span class="badge object-method"></span> schemaVersion

Dashboard version when this was saved (zero if unknown)

```typescript
schemaVersion(schemaVersion: number)
```

### <span class="badge object-method"></span> type

The panel type (from inside the model)

```typescript
type(type: string)
```

### <span class="badge object-method"></span> uid

Library element UID

```typescript
uid(uid: string)
```

### <span class="badge object-method"></span> version

panel version, incremented each time the dashboard is updated.

```typescript
version(version: number)
```

## See also

 * <span class="badge object-type-interface"></span> [LibraryPanel](./object-LibraryPanel.md)
