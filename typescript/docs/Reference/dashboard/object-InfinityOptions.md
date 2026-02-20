---
title: <span class="badge object-type-interface"></span> InfinityOptions
---
# <span class="badge object-type-interface"></span> InfinityOptions

Infinity options

## Definition

```typescript
export interface InfinityOptions {
	method: dashboard.HttpRequestMethod;
	url: string;
	body?: string;
	// These are 2D arrays of strings, each representing a key-value pair
	// We are defining them this way because we can't generate a go struct that
	// that would have exactly two strings in each sub-array
	queryParams?: string[][];
	headers?: string[][];
	datasourceUid: string;
}

```
## See also

 * <span class="badge builder"></span> [InfinityOptionsBuilder](./builder-InfinityOptionsBuilder.md)
