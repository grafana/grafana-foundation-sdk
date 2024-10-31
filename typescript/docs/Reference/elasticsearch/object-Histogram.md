---
title: <span class="badge object-type-interface"></span> Histogram
---
# <span class="badge object-type-interface"></span> Histogram

## Definition

```typescript
export interface Histogram {
	field?: string;
	id: string;
	type: "histogram";
	settings?: {
		interval?: string;
		min_doc_count?: string;
	};
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [HistogramBuilder](./builder-HistogramBuilder.md)
