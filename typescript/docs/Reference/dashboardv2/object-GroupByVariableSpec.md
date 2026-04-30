---
title: <span class="badge object-type-interface"></span> GroupByVariableSpec
---
# <span class="badge object-type-interface"></span> GroupByVariableSpec

GroupBy variable specification

## Definition

```typescript
export interface GroupByVariableSpec {
	name: string;
	defaultValue?: dashboardv2.VariableOption;
	current: dashboardv2.VariableOption;
	options: dashboardv2.VariableOption[];
	multi: boolean;
	label?: string;
	hide: dashboardv2.VariableHide;
	skipUrlSync: boolean;
	description?: string;
	origin?: dashboardv2.ControlSourceRef;
}

```
