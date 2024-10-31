// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as common from '../common';

// TODO docs
export class ReduceDataOptionsBuilder implements cog.Builder<common.ReduceDataOptions> {
    protected readonly internal: common.ReduceDataOptions;

    constructor() {
        this.internal = common.defaultReduceDataOptions();
    }

    /**
     * Builds the object.
     */
    build(): common.ReduceDataOptions {
        return this.internal;
    }

    // If true show each row value
    values(values: boolean): this {
        this.internal.values = values;
        return this;
    }

    // if showing all values limit
    limit(limit: number): this {
        this.internal.limit = limit;
        return this;
    }

    // When !values, pick one value for the whole field
    calcs(calcs: string[]): this {
        this.internal.calcs = calcs;
        return this;
    }

    // Which fields to show.  By default this is only numeric fields
    fields(fields: string): this {
        this.internal.fields = fields;
        return this;
    }
}
