import * as cog from '../cog';
import * as common from '../common';
export declare class StackableFieldConfigBuilder implements cog.Builder<common.StackableFieldConfig> {
    protected readonly internal: common.StackableFieldConfig;
    constructor();
    /**
     * Builds the object.
     */
    build(): common.StackableFieldConfig;
    stacking(stacking: cog.Builder<common.StackingConfig>): this;
}
