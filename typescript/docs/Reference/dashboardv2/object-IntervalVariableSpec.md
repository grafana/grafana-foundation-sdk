---
title: <span class="badge object-type-interface"></span> IntervalVariableSpec
---
# <span class="badge object-type-interface"></span> IntervalVariableSpec

Interval variable specification

## Definition

```typescript
export interface IntervalVariableSpec {
	name: string;
	query: string;
	current: dashboardv2.VariableOption;
	options: dashboardv2.VariableOption[];
	auto: boolean;
	auto_min: string;
	auto_count: number;
	refresh: "onTimeRangeChanged";
	label?: string;
	hide: dashboardv2.VariableHide;
	skipUrlSync: boolean;
	description?: string;
	origin?: dashboardv2.ControlSourceRef;
}

```
