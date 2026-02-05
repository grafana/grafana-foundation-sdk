import * as cog from '../cog';
import * as xychart from '../xychart';
export declare class XYSeriesConfigBuilder implements cog.Builder<xychart.XYSeriesConfig> {
    protected readonly internal: xychart.XYSeriesConfig;
    constructor();
    /**
     * Builds the object.
     */
    build(): xychart.XYSeriesConfig;
    name(name: {
        fixed?: string;
    }): this;
    frame(frame: {
        matcher: xychart.MatcherConfig;
    }): this;
    x(x: {
        matcher: xychart.MatcherConfig;
    }): this;
    y(y: {
        matcher: xychart.MatcherConfig;
    }): this;
    color(color: {
        matcher: xychart.MatcherConfig;
    }): this;
    size(size: {
        matcher: xychart.MatcherConfig;
    }): this;
}
