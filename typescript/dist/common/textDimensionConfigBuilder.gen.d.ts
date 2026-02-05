import * as cog from '../cog';
import * as common from '../common';
export declare class TextDimensionConfigBuilder implements cog.Builder<common.TextDimensionConfig> {
    protected readonly internal: common.TextDimensionConfig;
    constructor();
    /**
     * Builds the object.
     */
    build(): common.TextDimensionConfig;
    mode(mode: common.TextDimensionMode): this;
    field(field: string): this;
    fixed(fixed: string): this;
}
