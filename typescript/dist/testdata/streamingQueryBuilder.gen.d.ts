import * as cog from '../cog';
import * as testdata from '../testdata';
export declare class StreamingQueryBuilder implements cog.Builder<testdata.StreamingQuery> {
    protected readonly internal: testdata.StreamingQuery;
    constructor();
    /**
     * Builds the object.
     */
    build(): testdata.StreamingQuery;
    bands(bands: number): this;
    noise(noise: number): this;
    speed(speed: number): this;
    spread(spread: number): this;
    type(type: "fetch" | "logs" | "signal" | "traces"): this;
    url(url: string): this;
}
