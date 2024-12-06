// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as testdata from '../testdata';

export class CSVWaveBuilder implements cog.Builder<testdata.CSVWave> {
    protected readonly internal: testdata.CSVWave;

    constructor() {
        this.internal = testdata.defaultCSVWave();
    }

    /**
     * Builds the object.
     */
    build(): testdata.CSVWave {
        return this.internal;
    }

    labels(labels: string): this {
        this.internal.labels = labels;
        return this;
    }

    name(name: string): this {
        this.internal.name = name;
        return this;
    }

    timeStep(timeStep: number): this {
        this.internal.timeStep = timeStep;
        return this;
    }

    valuesCSV(valuesCSV: string): this {
        this.internal.valuesCSV = valuesCSV;
        return this;
    }
}
