import * as cog from '../cog';
import * as common from '../common';
export declare class HideableFieldConfigBuilder implements cog.Builder<common.HideableFieldConfig> {
    protected readonly internal: common.HideableFieldConfig;
    constructor();
    /**
     * Builds the object.
     */
    build(): common.HideableFieldConfig;
    hideFrom(hideFrom: cog.Builder<common.HideSeriesConfig>): this;
}
