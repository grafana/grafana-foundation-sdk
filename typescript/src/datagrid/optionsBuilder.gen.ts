// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as datagrid from '../datagrid';

export class OptionsBuilder implements cog.Builder<datagrid.Options> {
    protected readonly internal: datagrid.Options;

    constructor() {
        this.internal = datagrid.defaultOptions();
    }

    /**
     * Builds the object.
     */
    build(): datagrid.Options {
        return this.internal;
    }

    selectedSeries(selectedSeries: number): this {
        if (!(selectedSeries >= 0)) {
            throw new Error("selectedSeries must be >= 0");
        }
        this.internal.selectedSeries = selectedSeries;
        return this;
    }
}

