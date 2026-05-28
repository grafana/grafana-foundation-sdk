---
title: <span class="badge object-type-interface"></span> QueryVariableSpec
---
# <span class="badge object-type-interface"></span> QueryVariableSpec

Query variable specification

## Definition

```typescript
export interface QueryVariableSpec {
	name: string;
	current: dashboardv2.VariableOption;
	label?: string;
	hide: dashboardv2.VariableHide;
	refresh: dashboardv2.VariableRefresh;
	skipUrlSync: boolean;
	description?: string;
	query: dashboardv2.DataQueryKind;
	regex: string;
	regexApplyTo?: dashboardv2.VariableRegexApplyTo;
	sort: dashboardv2.VariableSort;
	definition?: string;
	options: dashboardv2.VariableOption[];
	multi: boolean;
	includeAll: boolean;
	allValue?: string;
	placeholder?: string;
	allowCustomValue: boolean;
	staticOptions?: dashboardv2.VariableOption[];
	staticOptionsOrder?: "before" | "after" | "sorted";
	origin?: dashboardv2.ControlSourceRef;
}

```
