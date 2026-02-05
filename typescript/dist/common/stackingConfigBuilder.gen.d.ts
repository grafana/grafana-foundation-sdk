import * as cog from '../cog';
import * as common from '../common';
export declare class StackingConfigBuilder implements cog.Builder<common.StackingConfig> {
    protected readonly internal: common.StackingConfig;
    constructor();
    /**
     * Builds the object.
     */
    build(): common.StackingConfig;
    mode(mode: common.StackingMode): this;
    group(group: string): this;
}
