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
	datasource: dashboard.DataSourceRef;
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
	target?: dashboard.AnnotationTarget;
	// TODO -- this should not exist here, it is based on the --grafana-- datasource
	type?: string;
	// Set to 1 for the standard annotation query all dashboards have by default.
	builtIn?: number;
	expr?: string;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [AnnotationQueryBuilder](./builder-AnnotationQueryBuilder.md)
