---
title: <span class="badge object-type-interface"></span> Histogram
---
# <span class="badge object-type-interface"></span> Histogram

## Definition

```typescript
export interface Histogram {
	field?: string;
	id: string;
	type: elasticsearch.BucketAggregationType.Histogram;
	settings?: {
		interval?: string;
		min_doc_count?: string;
	};
}

```
## See also

 * <span class="badge builder"></span> [HistogramBuilder](./builder-HistogramBuilder.md)
