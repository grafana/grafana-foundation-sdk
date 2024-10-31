---
title: <span class="badge object-type-interface"></span> MetricDefinitionsQuery
---
# <span class="badge object-type-interface"></span> MetricDefinitionsQuery

@deprecated Use MetricNamespaceQuery instead

## Definition

```typescript
export interface MetricDefinitionsQuery {
	rawQuery?: string;
	kind: "MetricDefinitionsQuery";
	subscription: string;
	resourceGroup: string;
	metricNamespace?: string;
	resourceName?: string;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [MetricDefinitionsQueryBuilder](./builder-MetricDefinitionsQueryBuilder.md)
