---
title: <span class="badge object-type-interface"></span> AdhocVariableSpec
---
# <span class="badge object-type-interface"></span> AdhocVariableSpec

Adhoc variable specification

## Definition

```typescript
export interface AdhocVariableSpec {
	name: string;
	baseFilters: dashboardv2.AdHocFilterWithLabels[];
	filters: dashboardv2.AdHocFilterWithLabels[];
	defaultKeys: dashboardv2.MetricFindValue[];
	label?: string;
	hide: dashboardv2.VariableHide;
	skipUrlSync: boolean;
	description?: string;
	allowCustomValue: boolean;
	// Whether the group-by operator is enabled in the ad hoc filter combobox.
	enableGroupBy?: boolean;
	origin?: dashboardv2.ControlSourceRef;
}

```
