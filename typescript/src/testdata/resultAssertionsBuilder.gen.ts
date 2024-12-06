// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as testdata from '../testdata';

export class ResultAssertionsBuilder implements cog.Builder<testdata.ResultAssertions> {
    protected readonly internal: testdata.ResultAssertions;

    constructor() {
        this.internal = testdata.defaultResultAssertions();
    }

    /**
     * Builds the object.
     */
    build(): testdata.ResultAssertions {
        return this.internal;
    }

    // Maximum frame count
    maxFrames(maxFrames: number): this {
        this.internal.maxFrames = maxFrames;
        return this;
    }

    // Type asserts that the frame matches a known type structure.
    // Possible enum values:
    //  - `""` 
    //  - `"timeseries-wide"` 
    //  - `"timeseries-long"` 
    //  - `"timeseries-many"` 
    //  - `"timeseries-multi"` 
    //  - `"directory-listing"` 
    //  - `"table"` 
    //  - `"numeric-wide"` 
    //  - `"numeric-multi"` 
    //  - `"numeric-long"` 
    //  - `"log-lines"` 
    type(type: "" | "timeseries-wide" | "timeseries-long" | "timeseries-many" | "timeseries-multi" | "directory-listing" | "table" | "numeric-wide" | "numeric-multi" | "numeric-long" | "log-lines"): this {
        this.internal.type = type;
        return this;
    }

    // TypeVersion is the version of the Type property. Versions greater than 0.0 correspond to the dataplane
    // contract documentation https://grafana.github.io/dataplane/contract/.
    typeVersion(typeVersion: number[]): this {
        this.internal.typeVersion = typeVersion;
        return this;
    }
}
