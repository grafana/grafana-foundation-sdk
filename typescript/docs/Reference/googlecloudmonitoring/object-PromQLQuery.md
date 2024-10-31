---
title: <span class="badge object-type-interface"></span> PromQLQuery
---
# <span class="badge object-type-interface"></span> PromQLQuery

PromQL sub-query properties.

## Definition

```typescript
export interface PromQLQuery {
	// GCP project to execute the query against.
	projectName: string;
	// PromQL expression/query to be executed.
	expr: string;
	// PromQL min step
	step: string;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [PromQLQueryBuilder](./builder-PromQLQueryBuilder.md)
