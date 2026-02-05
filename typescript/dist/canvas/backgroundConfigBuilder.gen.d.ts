import * as cog from '../cog';
import * as canvas from '../canvas';
import * as common from '../common';
export declare class BackgroundConfigBuilder implements cog.Builder<canvas.BackgroundConfig> {
    protected readonly internal: canvas.BackgroundConfig;
    constructor();
    /**
     * Builds the object.
     */
    build(): canvas.BackgroundConfig;
    color(color: cog.Builder<common.ColorDimensionConfig>): this;
    image(image: cog.Builder<common.ResourceDimensionConfig>): this;
    size(size: canvas.BackgroundImageSize): this;
}
