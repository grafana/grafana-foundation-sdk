---
title: <span class="badge object-type-interface"></span> Options
---
# <span class="badge object-type-interface"></span> Options

## Definition

```typescript
export interface Options {
	// Represents the index of the selected frame
	frameIndex: number;
	// Controls whether the panel should show the header
	showHeader: boolean;
	// Controls whether the header should show icons for the column types
	showTypeIcons?: boolean;
	// Used to control row sorting
	sortBy?: common.TableSortByFieldState[];
	// Controls footer options
	footer?: common.TableFooterOptions;
	// Controls the height of the rows
	cellHeight?: common.TableCellHeight;
}

```
## Methods

No methods.
