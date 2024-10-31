---
title: <span class="badge object-type-interface"></span> TableFieldOptions
---
# <span class="badge object-type-interface"></span> TableFieldOptions

Field options for each field within a table (e.g 10, "The String", 64.20, etc.)

Generally defines alignment, filtering capabilties, display options, etc.

## Definition

```typescript
export interface TableFieldOptions {
	width?: number;
	minWidth?: number;
	align: common.FieldTextAlignment;
	// This field is deprecated in favor of using cellOptions
	displayMode?: common.TableCellDisplayMode;
	cellOptions?: common.TableCellOptions;
	// ?? default is missing or false ??
	hidden?: boolean;
	inspect: boolean;
	filterable?: boolean;
	// Hides any header for a column, useful for columns that show some static content or buttons.
	hideHeader?: boolean;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [TableFieldOptionsBuilder](./builder-TableFieldOptionsBuilder.md)
