---
title: <span class="badge object-type-interface"></span> StreamingQuery
---
# <span class="badge object-type-interface"></span> StreamingQuery

## Definition

```typescript
export interface StreamingQuery {
	bands?: number;
	noise: number;
	speed: number;
	spread: number;
	// Possible enum values:
	//  - `"fetch"` 
	//  - `"logs"` 
	//  - `"signal"` 
	//  - `"traces"` 
	type: "fetch" | "logs" | "signal" | "traces";
	url?: string;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [StreamingQueryBuilder](./builder-StreamingQueryBuilder.md)
