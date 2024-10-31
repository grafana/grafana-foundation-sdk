---
title: <span class="badge object-type-interface"></span> StreamingQuery
---
# <span class="badge object-type-interface"></span> StreamingQuery

## Definition

```typescript
export interface StreamingQuery {
	type: "signal" | "logs" | "fetch" | "traces";
	speed: number;
	spread: number;
	noise: number;
	bands?: number;
	url?: string;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [StreamingQueryBuilder](./builder-StreamingQueryBuilder.md)
