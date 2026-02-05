import * as cog from '../cog';
import * as testdata from '../testdata';
export declare class PulseWaveQueryBuilder implements cog.Builder<testdata.PulseWaveQuery> {
    protected readonly internal: testdata.PulseWaveQuery;
    constructor();
    /**
     * Builds the object.
     */
    build(): testdata.PulseWaveQuery;
    offCount(offCount: number): this;
    offValue(offValue: number): this;
    onCount(onCount: number): this;
    onValue(onValue: number): this;
    timeStep(timeStep: number): this;
}
