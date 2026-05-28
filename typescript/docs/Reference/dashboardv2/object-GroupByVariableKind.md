---
title: <span class="badge object-type-interface"></span> GroupByVariableKind
---
# <span class="badge object-type-interface"></span> GroupByVariableKind

Group variable kind

## Definition

```typescript
export interface GroupByVariableKind {
	kind: "GroupByVariable";
	group: string;
	labels?: Record<string, string>;
	datasource?: {
		name?: string;
	};
	spec: dashboardv2.GroupByVariableSpec;
}

```
## See also

 * <span class="badge builder"></span> [GroupByVariableBuilder](./builder-GroupByVariableBuilder.md)
