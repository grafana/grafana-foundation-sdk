// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as prometheus from '../prometheus';

export class ResultAssertionsBuilder implements cog.Builder<prometheus.ResultAssertions> {
    protected readonly internal: prometheus.ResultAssertions;

    constructor() {
        this.internal = prometheus.defaultResultAssertions();
    }

    /**
     * Builds the object.
     */
    build(): prometheus.ResultAssertions {
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
    type(type: prometheus.ResultAssertionsType): this {
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

