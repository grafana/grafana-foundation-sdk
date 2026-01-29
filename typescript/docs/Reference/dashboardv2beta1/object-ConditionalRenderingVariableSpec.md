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
## See also

 * <span class="badge builder"></span> [ConditionalRenderingVariableSpecBuilder](./builder-ConditionalRenderingVariableSpecBuilder.md)
