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
	current: dashboardv2.VariableOption;
	options: dashboardv2.VariableOption[];
	multi: boolean;
	includeAll: boolean;
	allValue?: string;
	label?: string;
	hide: dashboardv2.VariableHide;
	skipUrlSync: boolean;
	description?: string;
	allowCustomValue: boolean;
	valuesFormat?: "csv" | "json";
	origin?: dashboardv2.ControlSourceRef;
}

```
