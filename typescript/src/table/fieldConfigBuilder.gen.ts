// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as table from '../table';
import * as common from '../common';

export class FieldConfigBuilder implements cog.Builder<table.FieldConfig> {
    protected readonly internal: table.FieldConfig;

    constructor() {
        this.internal = table.defaultFieldConfig();
    }

    /**
     * Builds the object.
     */
    build(): table.FieldConfig {
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

