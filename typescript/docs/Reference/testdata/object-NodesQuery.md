---
title: <span class="badge object-type-interface"></span> NodesQuery
---
# <span class="badge object-type-interface"></span> NodesQuery

## Definition

```typescript
export interface NodesQuery {
	count?: number;
	seed?: number;
	// Possible enum values:
	//  - `"random"` 
	//  - `"random edges"` 
	//  - `"response_medium"` 
	//  - `"response_small"` 
	//  - `"feature_showcase"` 
	type?: "random" | "random edges" | "response_medium" | "response_small" | "feature_showcase";
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [NodesQueryBuilder](./builder-NodesQueryBuilder.md)
