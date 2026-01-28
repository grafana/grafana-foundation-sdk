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
	hide: dashboardv2beta1.VariableHide;
	skipUrlSync: boolean;
	description?: string;
}

```
## See also

 * <span class="badge builder"></span> [SwitchVariableSpecBuilder](./builder-SwitchVariableSpecBuilder.md)
