---
title: <span class="badge object-type-interface"></span> AdHocFilterWithLabels
---
# <span class="badge object-type-interface"></span> AdHocFilterWithLabels

Define the AdHocFilterWithLabels type

## Definition

```typescript
export interface AdHocFilterWithLabels {
	key: string;
	operator: string;
	value: string;
	values?: string[];
	keyLabel?: string;
	valueLabels?: string[];
	forceEdit?: boolean;
	origin?: "dashboard";
	// @deprecated
	condition?: string;
}

```
## See also

 * <span class="badge builder"></span> [AdHocFilterWithLabelsBuilder](./builder-AdHocFilterWithLabelsBuilder.md)
