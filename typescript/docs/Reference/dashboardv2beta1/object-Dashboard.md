---
title: <span class="badge object-type-interface"></span> Dashboard
---
# <span class="badge object-type-interface"></span> Dashboard

## Definition

```typescript
export interface Dashboard {
	annotations: dashboardv2beta1.AnnotationQueryKind[];
	// Configuration of dashboard cursor sync behavior.
	// "Off" for no shared crosshair or tooltip (default).
	// "Crosshair" for shared crosshair.
	// "Tooltip" for shared crosshair AND shared tooltip.
	cursorSync: dashboardv2beta1.DashboardCursorSync;
	// Description of dashboard.
	description?: string;
	// Whether a dashboard is editable or not.
	editable?: boolean;
	elements: Record<string, dashboardv2beta1.Element>;
	layout: dashboardv2beta1.GridLayoutKind | dashboardv2beta1.RowsLayoutKind | dashboardv2beta1.AutoGridLayoutKind | dashboardv2beta1.TabsLayoutKind;
	// Links with references to other dashboards or external websites.
	links: dashboardv2beta1.DashboardLink[];
	// When set to true, the dashboard will redraw panels at an interval matching the pixel width.
	// This will keep data "moving left" regardless of the query refresh rate. This setting helps
	// avoid dashboards presenting stale live data.
	liveNow?: boolean;
	// When set to true, the dashboard will load all panels in the dashboard when it's loaded.
	preload: boolean;
	// Plugins only. The version of the dashboard installed together with the plugin.
	// This is used to determine if the dashboard should be updated when the plugin is updated.
	revision?: number;
	// Tags associated with dashboard.
	tags: string[];
	timeSettings: dashboardv2beta1.TimeSettingsSpec;
	// Title of dashboard.
	title: string;
	// Configured template variables.
	variables: dashboardv2beta1.VariableKind[];
}

```
## See also

 * <span class="badge builder"></span> [DashboardBuilder](./builder-DashboardBuilder.md)
