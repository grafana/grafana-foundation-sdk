---
title: <span class="badge object-type-interface"></span> MetricAggregationWithInlineScript
---
# <span class="badge object-type-interface"></span> MetricAggregationWithInlineScript

## Definition

```typescript
export interface MetricAggregationWithInlineScript {
	settings?: {
		script?: elasticsearch.InlineScript;
	};
	type: elasticsearch.MetricAggregationType;
	id: string;
	hide?: boolean;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [MetricAggregationWithInlineScriptBuilder](./builder-MetricAggregationWithInlineScriptBuilder.md)
