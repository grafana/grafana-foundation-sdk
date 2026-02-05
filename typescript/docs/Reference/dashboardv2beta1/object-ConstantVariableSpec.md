---
title: <span class="badge object-type-interface"></span> ConstantVariableSpec
---
# <span class="badge object-type-interface"></span> ConstantVariableSpec

Constant variable specification

## Definition

```typescript
export interface ConstantVariableSpec {
	name: string;
	query: string;
	current: dashboardv2beta1.VariableOption;
	label?: string;
	hide: dashboardv2beta1.VariableHide;
	skipUrlSync: boolean;
	description?: string;
}

```
