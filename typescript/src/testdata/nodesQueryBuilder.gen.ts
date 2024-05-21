// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as testdata from '../testdata';

export class NodesQueryBuilder implements cog.Builder<testdata.NodesQuery> {
    protected readonly internal: testdata.NodesQuery;

    constructor() {
        this.internal = testdata.defaultNodesQuery();
    }

    build(): testdata.NodesQuery {
        return this.internal;
    }

    count(count: number): this {
        this.internal.count = count;
        return this;
    }

    seed(seed: number): this {
        this.internal.seed = seed;
        return this;
    }

    // Possible enum values:
    //  - `"random"` 
    //  - `"random edges"` 
    //  - `"response_medium"` 
    //  - `"response_small"` 
    //  - `"feature_showcase"` 
    type(type: "random" | "random edges" | "response_medium" | "response_small" | "feature_showcase"): this {
        this.internal.type = type;
        return this;
    }
}
