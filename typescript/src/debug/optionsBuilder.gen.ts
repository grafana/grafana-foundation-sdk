// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as debug from '../debug';

export class OptionsBuilder implements cog.Builder<debug.Options> {
    protected readonly internal: debug.Options;

    constructor() {
        this.internal = debug.defaultOptions();
    }

    /**
     * Builds the object.
     */
    build(): debug.Options {
        return this.internal;
    }

    mode(mode: debug.DebugMode): this {
        this.internal.mode = mode;
        return this;
    }

    counters(counters: cog.Builder<debug.UpdateConfig>): this {
        const countersResource = counters.build();
        this.internal.counters = countersResource;
        return this;
    }
}

