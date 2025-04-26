---
title: <span class="badge object-type-interface"></span> Options
---
# <span class="badge object-type-interface"></span> Options

## Definition

```typescript
export interface Options {
	showLabels: boolean;
	showCommonLabels: boolean;
	showTime: boolean;
	showLogContextToggle: boolean;
	showControls?: boolean;
	controlsStorageKey?: string;
	wrapLogMessage: boolean;
	prettifyLogMessage: boolean;
	enableLogDetails: boolean;
	sortOrder: common.LogsSortOrder;
	dedupStrategy: common.LogsDedupStrategy;
	enableInfiniteScrolling?: boolean;
	// TODO: figure out how to define callbacks
	onClickFilterLabel?: any;
	onClickFilterOutLabel?: any;
	isFilterLabelActive?: any;
	onClickFilterString?: any;
	onClickFilterOutString?: any;
	onClickShowField?: any;
	onClickHideField?: any;
	onLogOptionsChange?: any;
	logRowMenuIconsBefore?: any;
	logRowMenuIconsAfter?: any;
	onNewLogsReceived?: any;
	displayedFields?: string[];
}

```
