---
title: <span class="badge object-type-interface"></span> Options
---
# <span class="badge object-type-interface"></span> Options

## Definition

```typescript
export interface Options {
	showControls: boolean;
	showTime: boolean;
	wrapLogMessage: boolean;
	enableLogDetails: boolean;
	syntaxHighlighting: boolean;
	sortOrder: common.LogsSortOrder;
	dedupStrategy: common.LogsDedupStrategy;
	grammar?: any;
	enableInfiniteScrolling?: boolean;
	onLogOptionsChange?: any;
	onNewLogsReceived?: any;
}

```
