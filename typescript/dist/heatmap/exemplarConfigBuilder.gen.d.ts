import * as cog from '../cog';
import * as heatmap from '../heatmap';
export declare class ExemplarConfigBuilder implements cog.Builder<heatmap.ExemplarConfig> {
    protected readonly internal: heatmap.ExemplarConfig;
    constructor();
    /**
     * Builds the object.
     */
    build(): heatmap.ExemplarConfig;
    color(color: string): this;
}
