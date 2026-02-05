import * as cog from '../cog';
import * as common from '../common';
export declare class ScaleDimensionConfigBuilder implements cog.Builder<common.ScaleDimensionConfig> {
    protected readonly internal: common.ScaleDimensionConfig;
    constructor();
    /**
     * Builds the object.
     */
    build(): common.ScaleDimensionConfig;
    min(min: number): this;
    max(max: number): this;
    fixed(fixed: number): this;
    field(field: string): this;
    mode(mode: common.ScaleDimensionMode): this;
}
