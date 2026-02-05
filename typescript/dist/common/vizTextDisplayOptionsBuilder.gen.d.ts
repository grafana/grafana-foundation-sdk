import * as cog from '../cog';
import * as common from '../common';
export declare class VizTextDisplayOptionsBuilder implements cog.Builder<common.VizTextDisplayOptions> {
    protected readonly internal: common.VizTextDisplayOptions;
    constructor();
    /**
     * Builds the object.
     */
    build(): common.VizTextDisplayOptions;
    titleSize(titleSize: number): this;
    valueSize(valueSize: number): this;
    percentSize(percentSize: number): this;
}
