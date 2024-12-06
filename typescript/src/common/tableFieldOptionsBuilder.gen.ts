// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as common from '../common';

// Field options for each field within a table (e.g 10, "The String", 64.20, etc.)
// Generally defines alignment, filtering capabilties, display options, etc.
export class TableFieldOptionsBuilder implements cog.Builder<common.TableFieldOptions> {
    protected readonly internal: common.TableFieldOptions;

    constructor() {
        this.internal = common.defaultTableFieldOptions();
    }

    /**
     * Builds the object.
     */
    build(): common.TableFieldOptions {
        return this.internal;
    }

    width(width: number): this {
        this.internal.width = width;
        return this;
    }

    minWidth(minWidth: number): this {
        this.internal.minWidth = minWidth;
        return this;
    }

    align(align: common.FieldTextAlignment): this {
        this.internal.align = align;
        return this;
    }

    // This field is deprecated in favor of using cellOptions
    displayMode(displayMode: common.TableCellDisplayMode): this {
        this.internal.displayMode = displayMode;
        return this;
    }

    cellOptions(cellOptions: common.TableCellOptions): this {
        this.internal.cellOptions = cellOptions;
        return this;
    }

    // ?? default is missing or false ??
    hidden(hidden: boolean): this {
        this.internal.hidden = hidden;
        return this;
    }

    inspect(inspect: boolean): this {
        this.internal.inspect = inspect;
        return this;
    }

    filterable(filterable: boolean): this {
        this.internal.filterable = filterable;
        return this;
    }

    // Hides any header for a column, useful for columns that show some static content or buttons.
    hideHeader(hideHeader: boolean): this {
        this.internal.hideHeader = hideHeader;
        return this;
    }
}
