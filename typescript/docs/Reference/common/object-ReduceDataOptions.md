---
title: <span class="badge object-type-interface"></span> ReduceDataOptions
---
# <span class="badge object-type-interface"></span> ReduceDataOptions

TODO docs

## Definition

```typescript
export interface ReduceDataOptions {
	// If true show each row value
	values?: boolean;
	// if showing all values limit
	limit?: number;
	// When !values, pick one value for the whole field
	calcs: string[];
	// Which fields to show.  By default this is only numeric fields
	fields?: string;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [ReduceDataOptionsBuilder](./builder-ReduceDataOptionsBuilder.md)
