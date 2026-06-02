---
title: <span class="badge object-type-interface"></span> ResourceGraphQuery
---
# <span class="badge object-type-interface"></span> ResourceGraphQuery

## Definition

```typescript
export interface ResourceGraphQuery {
	// Azure Resource Graph KQL query to be executed.
	query?: string;
	// Specifies the format results should be returned as. Defaults to table.
	resultFormat?: string;
}

```
## See also

 * <span class="badge builder"></span> [ResourceGraphQueryBuilder](./builder-ResourceGraphQueryBuilder.md)
