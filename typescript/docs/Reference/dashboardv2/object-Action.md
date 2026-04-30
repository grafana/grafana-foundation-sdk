---
title: <span class="badge object-type-interface"></span> Action
---
# <span class="badge object-type-interface"></span> Action

## Definition

```typescript
export interface Action {
	type: dashboardv2.ActionType;
	title: string;
	fetch?: dashboardv2.FetchOptions;
	infinity?: dashboardv2.InfinityOptions;
	confirmation?: string;
	oneClick?: boolean;
	variables?: dashboardv2.ActionVariable[];
	style?: {
		backgroundColor?: string;
	};
}

```
## See also

 * <span class="badge builder"></span> [ActionBuilder](./builder-ActionBuilder.md)
