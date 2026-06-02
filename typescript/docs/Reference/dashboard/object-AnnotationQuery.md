---
title: <span class="badge object-type-interface"></span> AnnotationQuery
---
# <span class="badge object-type-interface"></span> AnnotationQuery

TODO docs

FROM: AnnotationQuery in grafana-data/src/types/annotations.ts

## Definition

```typescript
export interface AnnotationQuery {
	// Name of annotation.
	name: string;
	// Datasource where the annotations data is
	datasource?: common.DataSourceRef;
	// When enabled the annotation query is issued with every dashboard refresh
	enable: boolean;
	// Annotation queries can be toggled on or off at the top of the dashboard.
	// When hide is true, the toggle is not shown in the dashboard.
	hide?: boolean;
	// Color to use for the annotation event markers
	iconColor: string;
	// Filters to apply when fetching annotations
	filter?: dashboard.AnnotationPanelFilter;
	// TODO.. this should just be a normal query target
	target?: cog.Dataquery;
	// TODO -- this should not exist here, it is based on the --grafana-- datasource
	type?: string;
	// Set to 1 for the standard annotation query all dashboards have by default.
	builtIn?: number;
	// Placement can be used to display the annotation query somewhere else on the dashboard other than the default location.
	placement?: dashboard.AnnotationQueryPlacement.InControlsMenu;
	expr?: string;
	// Format for Prometheus annotation text. Label values can be interpolated with templates like {{instance}}.
	textFormat?: string;
	// Format for Prometheus and Loki annotation titles. Label values can be interpolated with templates like {{instance}}.
	titleFormat?: string;
	// Comma-separated label keys used as annotation tags.
	tagKeys?: string;
	// Legacy Prometheus annotation query step interval.
	step?: string;
	// Use the Prometheus series value as the annotation timestamp.
	useValueForTime?: boolean;
	// Mappings define how to convert data frame fields to annotation event fields.
	mappings?: Record<string, dashboard.AnnotationEventFieldMapping>;
}

```
## See also

 * <span class="badge builder"></span> [AnnotationQueryBuilder](./builder-AnnotationQueryBuilder.md)
