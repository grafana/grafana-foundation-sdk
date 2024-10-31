// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as common from '../common';

// Footer options
export class TableFooterOptionsBuilder implements cog.Builder<common.TableFooterOptions> {
    protected readonly internal: common.TableFooterOptions;

    constructor() {
        this.internal = common.defaultTableFooterOptions();
    }

    /**
     * Builds the object.
     */
    build(): common.TableFooterOptions {
        return this.internal;
    }

    show(show: boolean): this {
        this.internal.show = show;
        return this;
    }

    // actually 1 value
    reducer(reducer: string[]): this {
        this.internal.reducer = reducer;
        return this;
    }

    fields(fields: string[]): this {
        this.internal.fields = fields;
        return this;
    }

    enablePagination(enablePagination: boolean): this {
        this.internal.enablePagination = enablePagination;
        return this;
    }

    countRows(countRows: boolean): this {
        this.internal.countRows = countRows;
        return this;
    }
}
