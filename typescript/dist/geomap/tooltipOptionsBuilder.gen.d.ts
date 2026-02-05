import * as cog from '../cog';
import * as geomap from '../geomap';
export declare class TooltipOptionsBuilder implements cog.Builder<geomap.TooltipOptions> {
    protected readonly internal: geomap.TooltipOptions;
    constructor();
    /**
     * Builds the object.
     */
    build(): geomap.TooltipOptions;
    mode(mode: geomap.TooltipMode): this;
}
