---
title: <span class="badge object-type-interface"></span> Resource
---
# <span class="badge object-type-interface"></span> Resource

## Definition

```typescript
export interface Resource {
	group: string;
	kind: string;
	// The set of resources
	// +listType=set
	names: string[];
}

```
## See also

 * <span class="badge builder"></span> [ResourceBuilder](./builder-ResourceBuilder.md)
