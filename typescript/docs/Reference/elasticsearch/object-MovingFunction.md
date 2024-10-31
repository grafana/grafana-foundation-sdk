---
title: <span class="badge object-type-interface"></span> MovingFunction
---
# <span class="badge object-type-interface"></span> MovingFunction

## Definition

```typescript
export interface MovingFunction {
	pipelineAgg?: string;
	field?: string;
	type: "moving_fn";
	id: string;
	settings?: {
		window?: string;
		script?: elasticsearch.InlineScript;
		shift?: string;
	};
	hide?: boolean;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [MovingFunctionBuilder](./builder-MovingFunctionBuilder.md)
