---
title: <span class="badge object-type-interface"></span> Options
---
# <span class="badge object-type-interface"></span> Options

## Definition

```typescript
export interface Options {
	showTime: boolean;
	wrapLogMessage: boolean;
	enableLogDetails: boolean;
	sortOrder: common.LogsSortOrder;
	dedupStrategy: common.LogsDedupStrategy;
	enableInfiniteScrolling?: boolean;
	onNewLogsReceived?: any;
}

```
## See also

 * <span class="badge builder"></span> [OptionsBuilder](./builder-OptionsBuilder.md)
