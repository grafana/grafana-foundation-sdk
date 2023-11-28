// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as testdata from '../testdata';

export class SimulationQueryBuilder implements cog.OptionsBuilder<testdata.SimulationQuery> {
    private readonly internal: testdata.SimulationQuery;

    constructor() {
        this.internal = testdata.defaultSimulationQuery();
    }

    build(): testdata.SimulationQuery {
        return this.internal;
    }

    key(key: {
	type: string;
	tick: number;
	uid?: string;
}): this {
        this.internal.key = key;
        return this;
    }

    config(config: any): this {
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
