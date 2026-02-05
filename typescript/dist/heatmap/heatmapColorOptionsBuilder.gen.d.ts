import * as cog from '../cog';
import * as heatmap from '../heatmap';
export declare class HeatmapColorOptionsBuilder implements cog.Builder<heatmap.HeatmapColorOptions> {
    protected readonly internal: heatmap.HeatmapColorOptions;
    constructor();
    /**
     * Builds the object.
     */
    build(): heatmap.HeatmapColorOptions;
    mode(mode: heatmap.HeatmapColorMode): this;
    scheme(scheme: string): this;
    fill(fill: string): this;
    scale(scale: heatmap.HeatmapColorScale): this;
    exponent(exponent: number): this;
    steps(steps: number): this;
    reverse(reverse: boolean): this;
    min(min: number): this;
    max(max: number): this;
}
