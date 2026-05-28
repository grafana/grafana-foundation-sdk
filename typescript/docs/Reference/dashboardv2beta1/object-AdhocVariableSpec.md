---
title: <span class="badge object-type-interface"></span> AdhocVariableSpec
---
# <span class="badge object-type-interface"></span> AdhocVariableSpec

Adhoc variable specification

## Definition

```typescript
export interface AdhocVariableSpec {
	name: string;
	baseFilters: dashboardv2beta1.AdHocFilterWithLabels[];
	filters: dashboardv2beta1.AdHocFilterWithLabels[];
	defaultKeys: dashboardv2beta1.MetricFindValue[];
	label?: string;
	hide: dashboardv2beta1.VariableHide;
	skipUrlSync: boolean;
	description?: string;
	allowCustomValue: boolean;
	// Whether the group-by operator is enabled in the ad hoc filter combobox.
	enableGroupBy?: boolean;
	origin?: dashboardv2beta1.ControlSourceRef;
}

```
