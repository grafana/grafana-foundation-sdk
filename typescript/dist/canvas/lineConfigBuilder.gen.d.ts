import * as cog from '../cog';
import * as canvas from '../canvas';
import * as common from '../common';
export declare class LineConfigBuilder implements cog.Builder<canvas.LineConfig> {
    protected readonly internal: canvas.LineConfig;
    constructor();
    /**
     * Builds the object.
     */
    build(): canvas.LineConfig;
    color(color: cog.Builder<common.ColorDimensionConfig>): this;
    width(width: number): this;
    radius(radius: number): this;
}
