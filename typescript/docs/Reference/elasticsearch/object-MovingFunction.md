---
title: <span class="badge object-type-interface"></span> MovingFunction
---
# <span class="badge object-type-interface"></span> MovingFunction

## Definition

```typescript
export interface MovingFunction {
	pipelineAgg?: string;
	field?: string;
	type: unknown;
	id: string;
	settings?: {
		window?: string;
		script?: elasticsearch.InlineScript;
		shift?: string;
	};
	hide?: boolean;
}

```
## See also

 * <span class="badge builder"></span> [MovingFunctionBuilder](./builder-MovingFunctionBuilder.md)
