import * as cog from '../cog';
import * as testdata from '../testdata';
export declare class CSVWaveBuilder implements cog.Builder<testdata.CSVWave> {
    protected readonly internal: testdata.CSVWave;
    constructor();
    /**
     * Builds the object.
     */
    build(): testdata.CSVWave;
    labels(labels: string): this;
    name(name: string): this;
    timeStep(timeStep: number): this;
    valuesCSV(valuesCSV: string): this;
}
