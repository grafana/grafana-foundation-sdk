// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as testdata from '../testdata';

export class SimulationQueryBuilder implements cog.Builder<testdata.SimulationQuery> {
    protected readonly internal: testdata.SimulationQuery;

    constructor() {
        this.internal = testdata.defaultSimulationQuery();
    }

    /**
     * Builds the object.
     */
    build(): testdata.SimulationQuery {
        return this.internal;
    }

    key(key: cog.Builder<testdata.Key>): this {
        const keyResource = key.build();
        this.internal.key = keyResource;
        return this;
    }

    config(config: Record<string, any>): this {
        this.internal.config = config;
        return this;
    }

    stream(stream: boolean): this {
        this.internal.stream = stream;
        return this;
    }

    last(last: boolean): this {
        this.internal.last = last;
        return this;
    }
}
