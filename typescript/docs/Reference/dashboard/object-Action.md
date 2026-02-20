---
title: <span class="badge object-type-interface"></span> Action
---
# <span class="badge object-type-interface"></span> Action

Dashboard action

## Definition

```typescript
export interface Action {
	type: dashboard.ActionType;
	title: string;
	fetch?: dashboard.FetchOptions;
	infinity?: dashboard.InfinityOptions;
	confirmation?: string;
	oneClick?: boolean;
	variables?: dashboard.ActionVariable[];
	style?: {
		backgroundColor?: string;
	};
}

```
## See also

 * <span class="badge builder"></span> [ActionBuilder](./builder-ActionBuilder.md)
