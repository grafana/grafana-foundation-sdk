---
title: <span class="badge object-type-interface"></span> AnnotationContainer
---
# <span class="badge object-type-interface"></span> AnnotationContainer

Contains the list of annotations that are associated with the dashboard.

Annotations are used to overlay event markers and overlay event tags on graphs.

Grafana comes with a native annotation store and the ability to add annotation events directly from the graph panel or via the HTTP API.

See https://grafana.com/docs/grafana/latest/dashboards/build-dashboards/annotate-visualizations/

## Definition

```typescript
export interface AnnotationContainer {
	// List of annotations
	list?: dashboard.AnnotationQuery[];
}

```
## Methods

No methods.
