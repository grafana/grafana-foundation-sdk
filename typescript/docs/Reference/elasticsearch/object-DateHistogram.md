---
title: <span class="badge object-type-interface"></span> DateHistogram
---
# <span class="badge object-type-interface"></span> DateHistogram

## Definition

```typescript
export interface DateHistogram {
	field?: string;
	id: string;
	type: "date_histogram";
	settings?: {
		interval?: string;
		min_doc_count?: string;
		trimEdges?: string;
		offset?: string;
		timeZone?: string;
	};
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [DateHistogramBuilder](./builder-DateHistogramBuilder.md)
