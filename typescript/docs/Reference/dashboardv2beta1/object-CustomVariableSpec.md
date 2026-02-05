---
title: <span class="badge object-type-interface"></span> CustomVariableSpec
---
# <span class="badge object-type-interface"></span> CustomVariableSpec

Custom variable specification

## Definition

```typescript
export interface CustomVariableSpec {
	name: string;
	query: string;
	current: dashboardv2beta1.VariableOption;
	options: dashboardv2beta1.VariableOption[];
	multi: boolean;
	includeAll: boolean;
	allValue?: string;
	label?: string;
	hide: dashboardv2beta1.VariableHide;
	skipUrlSync: boolean;
	description?: string;
	allowCustomValue: boolean;
}

```
