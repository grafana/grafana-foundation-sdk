import * as cog from '../cog';
import * as xychart from '../xychart';
export declare class MatcherConfigBuilder implements cog.Builder<xychart.MatcherConfig> {
    protected readonly internal: xychart.MatcherConfig;
    constructor();
    /**
     * Builds the object.
     */
    build(): xychart.MatcherConfig;
    id(id: string): this;
    options(options: any): this;
}
