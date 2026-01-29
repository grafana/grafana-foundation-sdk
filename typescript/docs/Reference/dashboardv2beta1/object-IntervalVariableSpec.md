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
	current: dashboardv2beta1.VariableOption;
	options: dashboardv2beta1.VariableOption[];
	auto: boolean;
	auto_min: string;
	auto_count: number;
	refresh: "onTimeRangeChanged";
	label?: string;
	hide: dashboardv2beta1.VariableHide;
	skipUrlSync: boolean;
	description?: string;
}

```
