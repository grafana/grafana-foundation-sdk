---
title: <span class="badge object-type-interface"></span> TextVariableSpec
---
# <span class="badge object-type-interface"></span> TextVariableSpec

Text variable specification

## Definition

```typescript
export interface TextVariableSpec {
	name: string;
	current: dashboardv2beta1.VariableOption;
	query: string;
	label?: string;
	hide: dashboardv2beta1.VariableHide;
	skipUrlSync: boolean;
	description?: string;
}

```
