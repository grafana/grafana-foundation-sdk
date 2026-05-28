---
title: <span class="badge object-type-interface"></span> TextVariableSpec
---
# <span class="badge object-type-interface"></span> TextVariableSpec

Text variable specification

## Definition

```typescript
export interface TextVariableSpec {
	name: string;
	current: dashboardv2.VariableOption;
	query: string;
	label?: string;
	hide: dashboardv2.VariableHide;
	skipUrlSync: boolean;
	description?: string;
	origin?: dashboardv2.ControlSourceRef;
}

```
