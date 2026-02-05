import * as cog from '../cog';
import * as common from '../common';
export declare class BarConfigBuilder implements cog.Builder<common.BarConfig> {
    protected readonly internal: common.BarConfig;
    constructor();
    /**
     * Builds the object.
     */
    build(): common.BarConfig;
    barAlignment(barAlignment: common.BarAlignment): this;
    barWidthFactor(barWidthFactor: number): this;
    barMaxWidth(barMaxWidth: number): this;
}
