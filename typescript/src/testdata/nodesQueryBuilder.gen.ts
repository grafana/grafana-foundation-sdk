// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as testdata from '../testdata';

export class NodesQueryBuilder implements cog.Builder<testdata.NodesQuery> {
    protected readonly internal: testdata.NodesQuery;

    constructor() {
        this.internal = testdata.defaultNodesQuery();
    }

    /**
     * Builds the object.
     */
    build(): testdata.NodesQuery {
        return this.internal;
    }

    type(type: "random" | "response" | "random edges"): this {
        this.internal.type = type;
        return this;
    }

    count(count: number): this {
        this.internal.count = count;
        return this;
    }
}
