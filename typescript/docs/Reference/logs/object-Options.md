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
	wrapLogMessage: boolean;
	prettifyLogMessage: boolean;
	enableLogDetails: boolean;
	sortOrder: common.LogsSortOrder;
	dedupStrategy: common.LogsDedupStrategy;
	enableInfiniteScrolling?: boolean;
	// Display controls to jump to the last or first log line, and filters by log level.
	showControls?: boolean;
	// Show a component to manage the displayed fields from the logs.
	showFieldSelector?: boolean;
	// Use a predefined coloring scheme to highlight relevant parts of the log lines.
	syntaxHighlighting?: boolean;
	fontSize?: "default" | "small";
	detailsMode?: "inline" | "sidebar";
}

```
