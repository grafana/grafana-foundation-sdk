import * as cog from '../cog';
import * as common from '../common';
export declare class BaseDimensionConfigBuilder implements cog.Builder<common.BaseDimensionConfig> {
    protected readonly internal: common.BaseDimensionConfig;
    constructor();
    /**
     * Builds the object.
     */
    build(): common.BaseDimensionConfig;
    field(field: string): this;
}
