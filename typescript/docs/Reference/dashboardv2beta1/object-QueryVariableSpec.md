---
title: <span class="badge object-type-interface"></span> QueryVariableSpec
---
# <span class="badge object-type-interface"></span> QueryVariableSpec

Query variable specification

## Definition

```typescript
export interface QueryVariableSpec {
	name: string;
	current: dashboardv2beta1.VariableOption;
	label?: string;
	hide: dashboardv2beta1.VariableHide;
	refresh: dashboardv2beta1.VariableRefresh;
	skipUrlSync: boolean;
	description?: string;
	query: dashboardv2beta1.DataQueryKind;
	regex: string;
	regexApplyTo?: dashboardv2beta1.VariableRegexApplyTo;
	sort: dashboardv2beta1.VariableSort;
	definition?: string;
	options: dashboardv2beta1.VariableOption[];
	multi: boolean;
	includeAll: boolean;
	allValue?: string;
	placeholder?: string;
	allowCustomValue: boolean;
	staticOptions?: dashboardv2beta1.VariableOption[];
	staticOptionsOrder?: "before" | "after" | "sorted";
}

```
