---
title: <span class="badge object-type-interface"></span> AnnotationTarget
---
# <span class="badge object-type-interface"></span> AnnotationTarget

TODO: this should be a regular DataQuery that depends on the selected dashboard

these match the properties of the "grafana" datasouce that is default in most dashboards

## Definition

```typescript
export interface AnnotationTarget {
	// Only required/valid for the grafana datasource...
	// but code+tests is already depending on it so hard to change
	limit: number;
	// Only required/valid for the grafana datasource...
	// but code+tests is already depending on it so hard to change
	matchAny: boolean;
	// Only required/valid for the grafana datasource...
	// but code+tests is already depending on it so hard to change
	tags: string[];
	// Only required/valid for the grafana datasource...
	// but code+tests is already depending on it so hard to change
	type: string;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [AnnotationTargetBuilder](./builder-AnnotationTargetBuilder.md)
