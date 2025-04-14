---
title: <span class="badge object-type-interface"></span> GeoHashGrid
---
# <span class="badge object-type-interface"></span> GeoHashGrid

## Definition

```typescript
export interface GeoHashGrid {
	field?: string;
	id: string;
	type: elasticsearch.BucketAggregationType.GeohashGrid;
	settings?: {
		precision?: string;
	};
}

```
## See also

 * <span class="badge builder"></span> [GeoHashGridBuilder](./builder-GeoHashGridBuilder.md)
