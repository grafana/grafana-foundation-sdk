// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as testdata from '../testdata';

export class StreamingQueryBuilder implements cog.Builder<testdata.StreamingQuery> {
    protected readonly internal: testdata.StreamingQuery;

    constructor() {
        this.internal = testdata.defaultStreamingQuery();
    }

    /**
     * Builds the object.
     */
    build(): testdata.StreamingQuery {
        return this.internal;
    }

    type(type: "signal" | "logs" | "fetch"): this {
        this.internal.type = type;
        return this;
    }

    speed(speed: number): this {
        this.internal.speed = speed;
        return this;
    }

    spread(spread: number): this {
        this.internal.spread = spread;
        return this;
    }

    noise(noise: number): this {
        this.internal.noise = noise;
        return this;
    }

    bands(bands: number): this {
        this.internal.bands = bands;
        return this;
    }

    url(url: string): this {
        this.internal.url = url;
        return this;
    }
}
