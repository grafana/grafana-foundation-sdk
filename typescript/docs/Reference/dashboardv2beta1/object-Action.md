---
title: <span class="badge object-type-interface"></span> Action
---
# <span class="badge object-type-interface"></span> Action

## Definition

```typescript
export interface Action {
	type: dashboardv2beta1.ActionType;
	title: string;
	fetch?: dashboardv2beta1.FetchOptions;
	infinity?: dashboardv2beta1.InfinityOptions;
	confirmation?: string;
	oneClick?: boolean;
	variables?: dashboardv2beta1.ActionVariable[];
	style?: {
		backgroundColor?: string;
	};
}

```
## See also

 * <span class="badge builder"></span> [ActionBuilder](./builder-ActionBuilder.md)
