---
title: <span class="badge object-type-interface"></span> ConditionalRenderingVariableSpec
---
# <span class="badge object-type-interface"></span> ConditionalRenderingVariableSpec

## Definition

```typescript
export interface ConditionalRenderingVariableSpec {
	variable: string;
	operator: "equals" | "notEquals" | "matches" | "notMatches";
	value: string;
}

```
