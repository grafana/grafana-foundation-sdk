// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as testdata from '../testdata';

export class PulseWaveQueryBuilder implements cog.Builder<testdata.PulseWaveQuery> {
    protected readonly internal: testdata.PulseWaveQuery;

    constructor() {
        this.internal = testdata.defaultPulseWaveQuery();
    }

    /**
     * Builds the object.
     */
    build(): testdata.PulseWaveQuery {
        return this.internal;
    }

    offCount(offCount: number): this {
        this.internal.offCount = offCount;
        return this;
    }

    offValue(offValue: number): this {
        this.internal.offValue = offValue;
        return this;
    }

    onCount(onCount: number): this {
        this.internal.onCount = onCount;
        return this;
    }

    onValue(onValue: number): this {
        this.internal.onValue = onValue;
        return this;
    }

    timeStep(timeStep: number): this {
        this.internal.timeStep = timeStep;
        return this;
    }
}
