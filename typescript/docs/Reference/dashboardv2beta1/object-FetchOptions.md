---
title: <span class="badge object-type-interface"></span> FetchOptions
---
# <span class="badge object-type-interface"></span> FetchOptions

## Definition

```typescript
export interface FetchOptions {
	method: dashboardv2beta1.HttpRequestMethod;
	url: string;
	body?: string;
	// These are 2D arrays of strings, each representing a key-value pair
	// We are defining them this way because we can't generate a go struct that
	// that would have exactly two strings in each sub-array
	queryParams?: string[][];
	headers?: string[][];
}

```
## See also

 * <span class="badge builder"></span> [FetchOptionsBuilder](./builder-FetchOptionsBuilder.md)
