import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
export declare class ConditionalRenderingTimeRangeSizeSpecBuilder implements cog.Builder<dashboardv2beta1.ConditionalRenderingTimeRangeSizeSpec> {
    protected readonly internal: dashboardv2beta1.ConditionalRenderingTimeRangeSizeSpec;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.ConditionalRenderingTimeRangeSizeSpec;
    value(value: string): this;
}
