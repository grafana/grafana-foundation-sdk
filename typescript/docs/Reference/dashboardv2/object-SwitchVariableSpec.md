---
title: <span class="badge object-type-interface"></span> SwitchVariableSpec
---
# <span class="badge object-type-interface"></span> SwitchVariableSpec

## Definition

```typescript
export interface SwitchVariableSpec {
	name: string;
	current: string;
	enabledValue: string;
	disabledValue: string;
	label?: string;
	hide: dashboardv2.VariableHide;
	skipUrlSync: boolean;
	description?: string;
	origin?: dashboardv2.ControlSourceRef;
}

```
