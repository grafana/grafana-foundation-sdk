// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as testdata from '../testdata';

export class USAQueryBuilder implements cog.Builder<testdata.USAQuery> {
    protected readonly internal: testdata.USAQuery;

    constructor() {
        this.internal = testdata.defaultUSAQuery();
    }

    /**
     * Builds the object.
     */
    build(): testdata.USAQuery {
        return this.internal;
    }

    mode(mode: string): this {
        this.internal.mode = mode;
        return this;
    }

    period(period: string): this {
        this.internal.period = period;
        return this;
    }

    fields(fields: string[]): this {
        this.internal.fields = fields;
        return this;
    }

    states(states: string[]): this {
        this.internal.states = states;
        return this;
    }
}
