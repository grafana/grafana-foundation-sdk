---
title: <span class="badge object-type-interface"></span> GroupByVariableSpec
---
# <span class="badge object-type-interface"></span> GroupByVariableSpec

GroupBy variable specification

## Definition

```typescript
export interface GroupByVariableSpec {
	name: string;
	defaultValue?: dashboardv2beta1.VariableOption;
	current: dashboardv2beta1.VariableOption;
	options: dashboardv2beta1.VariableOption[];
	multi: boolean;
	label?: string;
	hide: dashboardv2beta1.VariableHide;
	skipUrlSync: boolean;
	description?: string;
}

```
